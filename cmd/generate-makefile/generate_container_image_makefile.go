// Copyright 2023 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/pkg/errors"
)

// containerImageGeneratedMkTemplate is the template for the file container-image.generated.mk, use containerImageGeneratedMkOptions as the context.
const containerImageGeneratedMkTemplate = `# Generated by ./cmd/generate-makefile. DO NOT EDIT.

.PHONY: image
image-all: {{- range .Recipes }} image-{{.ImageName}} {{- end }} ## Build all container images

{{ .Content -}}

.PHONY: clean-image-built
clean-image-built:
{{- range .Recipes }}
	rm -f {{ .SourcePath }}/.dockerbuilt
{{- end }}

`

// containerImageGeneratedMkTemplate is the template for one target, use containerImageRecipeOptions as the context.
const containerImageRecipeTemplate = `.PHONY: image-{{ .ImageName }}
image-{{ .ImageName }}:{{ .SourcePath }}/.dockerbuilt ## {{ .Comment }}

{{ .SourcePath }}/.dockerbuilt: SHELL=bash
{{ .SourcePath }}/.dockerbuilt: {{ StringsJoin .DependencyTargets " " }} {{ .SourcePath }}/Dockerfile
	$(ROOT)/build/build_image.py {{ .ImageName }} {{ .SourcePath }}
	touch {{ .SourcePath }}/.dockerbuilt

`

func renderContainerImageGeneratedMk() error {
	targetFile, err := os.OpenFile("container-image.generated.mk", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, "open file container-image.generated.mk")
	}
	recipeTemplate, err := template.New("container-image.generated.mk recipe").Funcs(defaultFuncMap).Parse(containerImageRecipeTemplate)
	if err != nil {
		return errors.Wrap(err, "parse container-image.generated.mk recipe")
	}
	var buffer bytes.Buffer
	for _, recipe := range containerImageRecipes {
		err := recipeTemplate.Execute(&buffer, recipe)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("render recipe in container-image.generated.mk, recipe: image-%s", recipe.ImageName))
		}
	}

	containerImageTemplate, err := template.New("container-image.generated.mk").Funcs(defaultFuncMap).Parse(containerImageGeneratedMkTemplate)
	if err != nil {
		return errors.Wrap(err, "parse container-image.generated.mk template")
	}
	err = containerImageTemplate.Execute(targetFile, containerImageGeneratedMkOptions{
		Recipes: containerImageRecipes,
		Content: buffer.String(),
	})
	if err != nil {
		return errors.Wrap(err, "render container-image.generated.mk")
	}

	return nil
}

type containerImageRecipeOptions struct {
	ImageName  string
	SourcePath string
	// DependencyTargets are the targets that this target depends on.
	DependencyTargets []string
	// Comment is the comment for the target, do not need to include the leading `##`
	Comment string
}

type containerImageGeneratedMkOptions struct {
	Recipes []containerImageRecipeOptions
	Content string
}

var containerImageRecipes = []containerImageRecipeOptions{
	{
		ImageName:  "chaos-daemon",
		SourcePath: "images/chaos-daemon",
		DependencyTargets: []string{
			"images/chaos-daemon/bin/chaos-daemon",
			"images/chaos-daemon/bin/pause",
			"images/chaos-daemon/bin/cdh",
		},
		Comment: "Build container image for chaos-daemon, ghcr.io/chaos-mesh/chaos-daemon:latest",
	}, {
		ImageName:         "chaos-mesh",
		SourcePath:        "images/chaos-mesh",
		DependencyTargets: []string{"images/chaos-mesh/bin/chaos-controller-manager"},
		Comment:           "Build container image for chaos-mesh, ghcr.io/chaos-mesh/chaos-mesh:latest",
	}, {
		ImageName:         "chaos-dashboard",
		SourcePath:        "images/chaos-dashboard",
		DependencyTargets: []string{"images/chaos-dashboard/bin/chaos-dashboard"},
		Comment:           "Build container image for chaos-dashboard, ghcr.io/chaos-mesh/chaos-dashboard:latest",
	}, {
		ImageName:         "build-env",
		SourcePath:        "images/build-env",
		DependencyTargets: nil,
		Comment:           "Build container image for build-env, ghcr.io/chaos-mesh/build-env:latest",
	}, {
		ImageName:         "dev-env",
		SourcePath:        "images/dev-env",
		DependencyTargets: nil,
		Comment:           "Build container image for build-env, ghcr.io/chaos-mesh/dev-env:latest",
	}, {
		ImageName:         "e2e-helper",
		SourcePath:        "e2e-test/cmd/e2e_helper",
		DependencyTargets: nil,
		Comment:           "Build container image for e2e-helper",
	}, {
		ImageName:  "chaos-mesh-e2e",
		SourcePath: "e2e-test/image/e2e",
		DependencyTargets: []string{
			"e2e-test/image/e2e/manifests",
			"e2e-test/image/e2e/chaos-mesh",
			"e2e-build",
		},
		Comment: "Build container image for running e2e tests",
	}, {
		ImageName:         "chaos-kernel",
		SourcePath:        "images/chaos-kernel",
		DependencyTargets: nil,
		Comment:           "Build container image for chaos-kernel, ghcr.io/chaos-mesh/chaos-kernel:latest",
	}, {
		ImageName:         "chaos-dlv",
		SourcePath:        "images/chaos-dlv",
		DependencyTargets: nil,
		Comment:           "Build container image for chaos-dlv",
	},
}
