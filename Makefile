# Set DEBUGGER=1 to build debug symbols
LDFLAGS = $(if $(IMG_LDFLAGS),$(IMG_LDFLAGS),$(if $(DEBUGGER),,-s -w) $(shell ./hack/version.sh))
DOCKER_REGISTRY ?= "localhost:5000"

# SET DOCKER_REGISTRY to change the docker registry
DOCKER_REGISTRY_PREFIX := $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/,)
DOCKER_BUILD_ARGS := --build-arg HTTP_PROXY=${HTTP_PROXY} --build-arg HTTPS_PROXY=${HTTPS_PROXY} --build-arg UI=${UI} --build-arg SWAGGER=${SWAGGER} --build-arg LDFLAGS="${LDFLAGS}" --build-arg CRATES_MIRROR="${CRATES_MIRROR}"

GOVER_MAJOR := $(shell go version | sed -E -e "s/.*go([0-9]+)[.]([0-9]+).*/\1/")
GOVER_MINOR := $(shell go version | sed -E -e "s/.*go([0-9]+)[.]([0-9]+).*/\2/")
GO111 := $(shell [ $(GOVER_MAJOR) -gt 1 ] || [ $(GOVER_MAJOR) -eq 1 ] && [ $(GOVER_MINOR) -ge 11 ]; echo $$?)

IMAGE_TAG := $(if $(IMAGE_TAG),$(IMAGE_TAG),latest)

ROOT=$(shell pwd)
OUTPUT_BIN=$(ROOT)/output/bin
KUSTOMIZE_BIN=$(OUTPUT_BIN)/kustomize
KUBEBUILDER_BIN=$(OUTPUT_BIN)/kubebuilder
KUBECTL_BIN=$(OUTPUT_BIN)/kubectl
HELM_BIN=$(OUTPUT_BIN)/helm

ifeq ($(GO111), 1)
$(error Please upgrade your Go compiler to 1.11 or higher version)
endif

# Enable GO111MODULE=on explicitly, disable it with GO111MODULE=off when necessary.
export GO111MODULE := on
GOOS   := $(if $(GOOS),$(GOOS),"")
GOARCH := $(if $(GOARCH),$(GOARCH),"")
GOENV  := GO15VENDOREXPERIMENT="1" CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH)
CGOENV := GO15VENDOREXPERIMENT="1" CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH)
GO     := $(GOENV) go
CGO    := $(CGOENV) go
GOTEST := USE_EXISTING_CLUSTER=false NO_PROXY="${NO_PROXY},testhost" go test
SHELL  := bash

PACKAGE_LIST := echo $$(go list ./... | grep -vE "chaos-mesh/test|pkg/ptrace|zz_generated|vendor") github.com/chaos-mesh/chaos-mesh/api/v1alpha1

# no version conversion
CRD_OPTIONS ?= "crd:trivialVersions=true,preserveUnknownFields=false,crdVersions=v1"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

GO_BUILD_CACHE ?= $(ROOT)/.cache/chaos-mesh

BUILD_TAGS ?=

ifeq ($(SWAGGER),1)
	BUILD_TAGS += swagger_server
endif

ifeq ($(UI),1)
	BUILD_TAGS += ui_server
endif

CLEAN_TARGETS :=

all: yaml image
go_build_cache_directory:  $(GO_BUILD_CACHE)/chaos-mesh-gobuild $(GO_BUILD_CACHE)/chaos-mesh-gopath
$(GO_BUILD_CACHE)/chaos-mesh-gobuild:
	@mkdir -p $(GO_BUILD_CACHE)/chaos-mesh-gobuild
$(GO_BUILD_CACHE)/chaos-mesh-gopath:
	@mkdir -p $(GO_BUILD_CACHE)/chaos-mesh-gopath

test-utils: timer multithread_tracee

timer:
	$(GO) build -ldflags '$(LDFLAGS)' -o bin/test/timer ./test/cmd/timer/*.go

multithread_tracee: test/cmd/multithread_tracee/main.c
	cc test/cmd/multithread_tracee/main.c -lpthread -O2 -o ./bin/test/multithread_tracee

mockgen:
	GO111MODULE=on go get github.com/golang/mock/mockgen@v1.5.0

generate-mock: mockgen
	go generate ./pkg/workflow

swagger_spec:
ifeq (${SWAGGER},1)
	hack/generate_swagger_spec.sh
endif

yarn_dependencies:
ifeq (${UI},1)
	cd ui &&\
	yarn install --frozen-lockfile
endif

ui: yarn_dependencies
ifeq (${UI},1)
	cd ui &&\
	yarn build
	hack/embed_ui_assets.sh
endif

watchmaker:
	$(CGO) build -ldflags '$(LDFLAGS)' -o bin/watchmaker ./cmd/watchmaker/...

# Build chaosctl
chaosctl:
	$(GO) build -ldflags '$(LDFLAGS)' -o bin/chaosctl ./cmd/chaosctl/*.go

# Build schedule-migration
schedule-migration:
	$(GO) build -ldflags '$(LDFLAGS)' -o bin/schedule-migration ./tools/schedule-migration/*.go

schedule-migration.tar.gz: schedule-migration
	cp ./bin/schedule-migration ./schedule-migration
	cp ./tools/schedule-migration/migrate.sh ./migrate.sh
	tar -czvf schedule-migration.tar.gz schedule-migration migrate.sh
	rm ./migrate.sh
	rm ./schedule-migration

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	$(GO) run ./cmd/controller-manager/main.go

NAMESPACE ?= chaos-testing
# Install CRDs into a cluster
install: manifests
	$(HELM_BIN) upgrade --install chaos-mesh helm/chaos-mesh --namespace=${NAMESPACE} --set registry=${DOCKER_REGISTRY} --set dnsServer.create=true --set dashboard.create=true;

clean:
	rm -rf docs/docs.go $(CLEAN_TARGETS)

boilerplate:
	./hack/verify-boilerplate.sh

image: image-chaos-daemon image-chaos-mesh image-chaos-dashboard $(if $(DEBUGGER), image-chaos-dlv)

e2e-image: image-e2e-helper

GO_TARGET_PHONY :=

BINARIES :=

define COMPILE_GO_TEMPLATE
ifeq ($(IN_DOCKER),1)

$(1): $(4)
ifeq ($(3),1)
	$(CGO) build -ldflags "$(LDFLAGS)" -tags "${BUILD_TAGS}" -o $(1) $(2)
else
	$(GO) build -ldflags "$(LDFLAGS)" -tags "${BUILD_TAGS}" -o $(1) $(2)
endif

endif
GO_TARGET_PHONY += $(1)
endef

BUILD_INDOCKER_ARG := --env IN_DOCKER=1 --volume $(ROOT):/mnt --user $(shell id -u):$(shell id -g)

ifneq ($(GO_BUILD_CACHE),)
	BUILD_INDOCKER_ARG += --volume $(GO_BUILD_CACHE)/chaos-mesh-gopath:/tmp/go
	BUILD_INDOCKER_ARG += --volume $(GO_BUILD_CACHE)/chaos-mesh-gobuild:/tmp/go-build
endif

define BUILD_IN_DOCKER_TEMPLATE
CLEAN_TARGETS += $(2)
ifneq ($(IN_DOCKER),1)

$(2): image-build-env go_build_cache_directory
	@DOCKER_ID=$$$$(docker run -d \
		$(BUILD_INDOCKER_ARG) \
		${DOCKER_REGISTRY_PREFIX}pingcap/build-env:${IMAGE_TAG} \
		sleep infinity); \
	docker exec --workdir /mnt/ \
		--env IMG_LDFLAGS="${LDFLAGS}" \
		--env UI=${UI} --env SWAGGER=${SWAGGER} \
		$$$$DOCKER_ID /usr/bin/make $(2) && \
	docker rm -f $$$$DOCKER_ID
endif

image-$(1)-dependencies := $(image-$(1)-dependencies) $(2)
BINARIES := $(BINARIES) $(2)
endef

enter-buildenv: image-build-env go_build_cache_directory
	@docker run -it \
		$(BUILD_INDOCKER_ARG) \
		${DOCKER_REGISTRY_PREFIX}pingcap/build-env:${IMAGE_TAG} \
		bash

enter-devenv: image-dev-env go_build_cache_directory
	@docker run -it \
		$(BUILD_INDOCKER_ARG) \
		${DOCKER_REGISTRY_PREFIX}pingcap/dev-env:${IMAGE_TAG} \
		bash

ifeq ($(IN_DOCKER),1)
images/chaos-daemon/bin/pause: hack/pause.c
	cc ./hack/pause.c -o images/chaos-daemon/bin/pause
endif
$(eval $(call BUILD_IN_DOCKER_TEMPLATE,chaos-daemon,images/chaos-daemon/bin/pause))

$(eval $(call BUILD_IN_DOCKER_TEMPLATE,chaos-daemon,images/chaos-daemon/bin/chaos-daemon))
$(eval $(call COMPILE_GO_TEMPLATE,images/chaos-daemon/bin/chaos-daemon,./cmd/chaos-daemon/main.go,1))

$(eval $(call BUILD_IN_DOCKER_TEMPLATE,chaos-dashboard,images/chaos-dashboard/bin/chaos-dashboard))
$(eval $(call COMPILE_GO_TEMPLATE,images/chaos-dashboard/bin/chaos-dashboard,./cmd/chaos-dashboard/main.go,1,ui swagger_spec))

$(eval $(call BUILD_IN_DOCKER_TEMPLATE,chaos-mesh,images/chaos-mesh/bin/chaos-controller-manager))
$(eval $(call COMPILE_GO_TEMPLATE,images/chaos-mesh/bin/chaos-controller-manager,./cmd/chaos-controller-manager/main.go,0))

prepare-install: all docker-push docker-push-dns-server

prepare-e2e: e2e-image docker-push-e2e

GINKGO_FLAGS ?=
e2e: e2e-build
	./e2e-test/image/e2e/bin/ginkgo ${GINKGO_FLAGS} ./e2e-test/image/e2e/bin/e2e.test -- --e2e-image ${DOCKER_REGISTRY_PREFIX}pingcap/e2e-helper:${IMAGE_TAG}

image-chaos-mesh-e2e-dependencies += e2e-test/image/e2e/manifests e2e-test/image/e2e/chaos-mesh e2e-build
CLEAN_TARGETS += e2e-test/image/e2e/manifests e2e-test/image/e2e/chaos-mesh

e2e-build: e2e-test/image/e2e/bin/ginkgo e2e-test/image/e2e/bin/e2e.test

e2e-test/image/e2e/manifests: manifests
	rm -rf e2e-test/image/e2e/manifests
	cp -r manifests e2e-test/image/e2e

e2e-test/image/e2e/chaos-mesh: helm/chaos-mesh
	rm -rf e2e-test/image/e2e/chaos-mesh
	cp -r helm/chaos-mesh e2e-test/image/e2e

define IMAGE_TEMPLATE
CLEAN_TARGETS += $(2)/.dockerbuilt

image-$(1): $(2)/.dockerbuilt

$(2)/.dockerbuilt:$(image-$(1)-dependencies) $(2)/Dockerfile
ifeq ($(DOCKER_CACHE),1)

ifneq ($(DISABLE_CACHE_FROM),1)
	DOCKER_BUILDKIT=1 DOCKER_CLI_EXPERIMENTAL=enabled docker buildx build --load --cache-to type=local,dest=$(DOCKER_CACHE_DIR)/image-$(1) --cache-from type=local,src=$(DOCKER_CACHE_DIR)/image-$(1) -t ${DOCKER_REGISTRY_PREFIX}pingcap/$(1):${IMAGE_TAG} ${DOCKER_BUILD_ARGS} $(2)
else
	DOCKER_BUILDKIT=1 DOCKER_CLI_EXPERIMENTAL=enabled docker buildx build --load --cache-to type=local,dest=$(DOCKER_CACHE_DIR)/image-$(1) -t ${DOCKER_REGISTRY_PREFIX}pingcap/$(1):${IMAGE_TAG} ${DOCKER_BUILD_ARGS} $(2)
endif

else ifneq ($(TARGET_PLATFORM),)
	DOCKER_BUILDKIT=1 docker buildx build --load --platform linux/$(TARGET_PLATFORM) -t ${DOCKER_REGISTRY_PREFIX}pingcap/$(1):${IMAGE_TAG} --build-arg TARGET_PLATFORM=$(TARGET_PLATFORM) ${DOCKER_BUILD_ARGS} $(2)
else
	DOCKER_BUILDKIT=1 docker build -t ${DOCKER_REGISTRY_PREFIX}pingcap/$(1):${IMAGE_TAG} ${DOCKER_BUILD_ARGS} $(2)
endif
	touch $(2)/.dockerbuilt
endef

$(eval $(call IMAGE_TEMPLATE,chaos-daemon,images/chaos-daemon))
$(eval $(call IMAGE_TEMPLATE,chaos-mesh,images/chaos-mesh))
$(eval $(call IMAGE_TEMPLATE,chaos-dashboard,images/chaos-dashboard))
$(eval $(call IMAGE_TEMPLATE,build-env,images/build-env))
$(eval $(call IMAGE_TEMPLATE,dev-env,images/dev-env))
$(eval $(call IMAGE_TEMPLATE,e2e-helper,e2e-test/cmd/e2e_helper))
$(eval $(call IMAGE_TEMPLATE,chaos-mesh-e2e,e2e-test/image/e2e))
$(eval $(call IMAGE_TEMPLATE,chaos-kernel,images/chaos-kernel))
$(eval $(call IMAGE_TEMPLATE,chaos-jvm,images/chaos-jvm))
$(eval $(call IMAGE_TEMPLATE,chaos-dlv,images/chaos-dlv))

binary: $(BINARIES)

docker-push:
	docker push "${DOCKER_REGISTRY_PREFIX}pingcap/chaos-mesh:${IMAGE_TAG}"
	docker push "${DOCKER_REGISTRY_PREFIX}pingcap/chaos-dashboard:${IMAGE_TAG}"
	docker push "${DOCKER_REGISTRY_PREFIX}pingcap/chaos-daemon:${IMAGE_TAG}"

docker-push-e2e:
	docker push "${DOCKER_REGISTRY_PREFIX}pingcap/e2e-helper:${IMAGE_TAG}"

# the version of dns server should keep consistent with helm
DNS_SERVER_VERSION ?= v0.2.0
docker-push-dns-server:
	docker pull pingcap/coredns:${DNS_SERVER_VERSION}
	docker tag pingcap/coredns:${DNS_SERVER_VERSION} "${DOCKER_REGISTRY_PREFIX}pingcap/coredns:${DNS_SERVER_VERSION}"
	docker push "${DOCKER_REGISTRY_PREFIX}pingcap/coredns:${DNS_SERVER_VERSION}"

docker-push-chaos-kernel:
	docker push "${DOCKER_REGISTRY_PREFIX}pingcap/chaos-kernel:${IMAGE_TAG}"

bin/chaos-builder:
	$(CGOENV) go build -ldflags '$(LDFLAGS)' -o bin/chaos-builder ./cmd/chaos-builder/...

chaos-build: bin/chaos-builder
	bin/chaos-builder

define RUN_IN_DEV_ENV_TEMPLATE
ifeq ($$(IN_DOCKER),1)
$(1):$(2)
$($(1)-make)
else
$(1):image-dev-env go_build_cache_directory
	@docker run $$$$(if [ -t 0 ] ;then echo -n "-it";fi) --rm --workdir /mnt/ \
		--cap-add=sys_ptrace \
		--env CI="${CI}" \
		$(BUILD_INDOCKER_ARG) \
		${DOCKER_REGISTRY_PREFIX}pingcap/dev-env:${IMAGE_TAG} \
		/usr/bin/make $(1)
endif
endef

define proto-make
	for dir in pkg/chaosdaemon pkg/chaoskernel ; do\
		protoc -I $$$$dir/pb $$$$dir/pb/*.proto --go_out=plugins=grpc:$$$$dir/pb --go_out=./$$$$dir/pb ;\
	done
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,proto))

define manifests/crd.yaml-make
	kustomize build config/default > manifests/crd.yaml
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,manifests/crd.yaml,config))

define manifests/crd-v1beta1.yaml-make
	mkdir -p ./output
	cp -Tr ./config ./output/config-v1beta1
	cd ./api/v1alpha1 ;\
		controller-gen "crd:trivialVersions=true,preserveUnknownFields=false,crdVersions=v1beta1" rbac:roleName=manager-role paths="./..." output:crd:artifacts:config=../../output/config-v1beta1/crd/bases ;
	kustomize build output/config-v1beta1/default > manifests/crd-v1beta1.yaml
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,manifests/crd-v1beta1.yaml,config))
yaml: manifests/crd.yaml manifests/crd-v1beta1.yaml

define config-make
	cd ./api/v1alpha1 ;\
		controller-gen $(CRD_OPTIONS) rbac:roleName=manager-role paths="./..." output:crd:artifacts:config=../../config/crd/bases ;\
		controller-gen $(CRD_OPTIONS) rbac:roleName=manager-role paths="./..." output:crd:artifacts:config=../../helm/chaos-mesh/crds ;
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,config))

define lint-make
	revive -formatter friendly -config revive.toml $$$$($$(PACKAGE_LIST))
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,lint))

define failpoint-enable-make
	find $$$$PWD/* -type d | grep -vE "(\.git|bin|\.cache|ui)" | xargs failpoint-ctl enable
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,failpoint-enable))

define failpoint-disable-make
	find $$$$PWD/* -type d | grep -vE "(\.git|bin|\.cache|ui)" | xargs failpoint-ctl disable
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,failpoint-disable))

define groupimports-make
	goimports -w -l -local github.com/chaos-mesh/chaos-mesh $$$$(find . -type f -name '*.go' -not -path '**/zz_generated.*.go' -not -path './.cache/**')
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,groupimports))

define fmt-make
	$(CGO) fmt $$$$(go list ./... | grep -v 'zz_generated.*.go')
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,fmt,groupimports))

define vet-make
	$(CGOENV) go vet ./...
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,vet))

define tidy-make
	@echo "go mod tidy"
	GO111MODULE=on go mod tidy
	git diff -U --exit-code go.mod go.sum
	cd api/v1alpha1; GO111MODULE=on go mod tidy; git diff -U --exit-code go.mod go.sum
	cd e2e-test; GO111MODULE=on go mod tidy; git diff -U --exit-code go.mod go.sum
	cd e2e-test/cmd/e2e_helper; GO111MODULE=on go mod tidy; git diff -U --exit-code go.mod go.sum
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,tidy,clean))

define generate-ctrl-make
	go generate ./pkg/ctrlserver/graph
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,generate-ctrl))

define generate-make
	cd ./api/v1alpha1 ;\
		controller-gen object:headerFile=../../hack/boilerplate/boilerplate.generatego.txt paths="./..." ;
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,generate,chaos-build,generate-ctrl))
check: fmt vet boilerplate lint generate yaml tidy install.sh

CLEAN_TARGETS+=e2e-test/image/e2e/bin/ginkgo
define e2e-test/image/e2e/bin/ginkgo-make
	cd e2e-test && $(GO) build -ldflags "$(LDFLAGS)" -tags "${BUILD_TAGS}" -o image/e2e/bin/ginkgo github.com/onsi/ginkgo/ginkgo
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,e2e-test/image/e2e/bin/ginkgo))

CLEAN_TARGETS+=e2e-test/image/e2e/bin/e2e.test
define e2e-test/image/e2e/bin/e2e.test-make
	cd e2e-test && $(GO) test -c  -o ./image/e2e/bin/e2e.test ./e2e
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,e2e-test/image/e2e/bin/e2e.test,e2e-test/e2e/**/*.go))

# Run tests
CLEAN_TARGETS += cover.out cover.out.tmp
define test-make
	$(GOTEST) -p 1 $$$$($$(PACKAGE_LIST)) -coverprofile cover.out.tmp
	cat cover.out.tmp | grep -v "_generated.deepcopy.go" > cover.out
	make failpoint-disable
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,test,failpoint-enable generate generate-mock manifests test-utils))

define gosec-scan-make
	gosec ./api/... ./controllers/... ./pkg/... || echo "*** sec-scan failed: known-issues ***"
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,gosec-scan))

define coverage-make
ifeq ("$(CI)", "1")
	@bash <(curl -s https://codecov.io/bash) -f cover.out -t $(CODECOV_TOKEN)
else
	mkdir -p cover
	gocov convert cover.out > cover.json
	gocov-xml < cover.json > cover.xml
	gocov-html < cover.json > cover/index.html
endif
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,coverage))

define install.sh-make
	./hack/update_install_script.sh
endef
$(eval $(call RUN_IN_DEV_ENV_TEMPLATE,install.sh))

.PHONY: all clean test install manifests groupimports fmt vet tidy image \
	docker-push lint generate config mockgen generate-mock \
	install.sh $(GO_TARGET_PHONY) \
	manager chaosfs chaosdaemon chaos-dashboard \
	dashboard dashboard-server-frontend gosec-scan \
	failpoint-enable failpoint-disable \
	proto bin/chaos-builder go_build_cache_directory schedule-migration enter-buildenv enter-devenv \
	manifests/crd.yaml
