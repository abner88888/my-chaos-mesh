// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"context"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	"github.com/chaos-mesh/chaos-mesh/controllers/config"
	ccfg "github.com/chaos-mesh/chaos-mesh/controllers/config"
	chaosdaemonclient "github.com/chaos-mesh/chaos-mesh/pkg/chaosdaemon/client"
	grpcUtils "github.com/chaos-mesh/chaos-mesh/pkg/grpc"
	"github.com/chaos-mesh/chaos-mesh/pkg/mock"
)

var log = ctrl.Log.WithName("controller-chaos-daemon-client-utils")
var cachedClient client.Client

func init() {
	cfg := ctrl.GetConfigOrDie()

	mapper, err := apiutil.NewDynamicRESTMapper(cfg)
	if err != nil {
		log.Error(err, "Failed to get API Group-Resources")

		return
	}

	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)

	// Create the cache for the cached read client and registering informers
	cache, err := cache.New(cfg, cache.Options{Scheme: scheme, Mapper: mapper, Resync: nil, Namespace: ccfg.ControllerCfg.Namespace})
	if err != nil {
		log.Error(err, "Failed to setup cache")

		return
	}
	// TODO: store the channel and use it to stop
	go cache.Start(make(chan struct{}))

	c, err := client.New(cfg, client.Options{Scheme: scheme, Mapper: mapper})
	if err != nil {
		log.Error(err, "Failed to setup client")

		return
	}

	cachedClient = &client.DelegatingClient{
		Reader: &client.DelegatingReader{
			CacheReader:  cache,
			ClientReader: c,
		},
		Writer:       c,
		StatusClient: c,
	}
}

func FindDaemonIP(ctx context.Context, c client.Reader, pod *v1.Pod) (string, error) {
	// TODO: use cached client to get the chaos-daemon
	// It will take several steps to use the namespace scoped cached client
	// 1. Upgrade controller-runtime to a high enough version (which supports scoped cache)
	// 2. Create a cached client in deployment namespace to get the endpoint
	// 3. Limit controller-manager in the target namespace

	nodeName := pod.Spec.NodeName
	log.Info("Creating client to chaos-daemon", "node", nodeName)

	ns := config.ControllerCfg.Namespace
	var endpoints v1.Endpoints
	err := c.Get(ctx, types.NamespacedName{
		Namespace: ns,
		Name:      "chaos-daemon",
	}, &endpoints)
	if err != nil {
		return "", err
	}

	daemonIP := findIPOnEndpoints(&endpoints, nodeName)
	if len(daemonIP) == 0 {
		return "", errors.Errorf("cannot find daemonIP on node %s in related Endpoints %v", nodeName, endpoints)
	}

	return daemonIP, nil
}

func findIPOnEndpoints(e *v1.Endpoints, nodeName string) string {
	for _, subset := range e.Subsets {
		for _, addr := range subset.Addresses {
			if addr.NodeName != nil && *addr.NodeName == nodeName {
				return addr.IP
			}
		}
	}

	return ""
}

// NewChaosDaemonClient would create ChaosDaemonClient
func NewChaosDaemonClient(ctx context.Context, c client.Reader, pod *v1.Pod) (chaosdaemonclient.ChaosDaemonClientInterface, error) {
	if cli := mock.On("MockChaosDaemonClient"); cli != nil {
		return cli.(chaosdaemonclient.ChaosDaemonClientInterface), nil
	}
	if err := mock.On("NewChaosDaemonClientError"); err != nil {
		return nil, err.(error)
	}

	daemonIP, err := FindDaemonIP(ctx, c, pod)
	if err != nil {
		return nil, err
	}
	builder := grpcUtils.Builder(daemonIP, config.ControllerCfg.ChaosDaemonPort).WithDefaultTimeout()
	if config.ControllerCfg.TLSConfig.ChaosMeshCACert != "" {
		builder.TLSFromFile(config.ControllerCfg.TLSConfig.ChaosMeshCACert, config.ControllerCfg.TLSConfig.ChaosDaemonClientCert, config.ControllerCfg.TLSConfig.ChaosDaemonClientKey)
	} else {
		builder.Insecure()
	}
	cc, err := builder.Build()
	if err != nil {
		return nil, err
	}
	return chaosdaemonclient.New(cc), nil
}
