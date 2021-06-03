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

package provider

import (
	"github.com/go-logr/logr"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	ccfg "github.com/chaos-mesh/chaos-mesh/controllers/config"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = v1alpha1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func NewScheme() *runtime.Scheme {
	return scheme
}

func NewOption(logger logr.Logger) *ctrl.Options {
	setupLog := logger.WithName("setup")

	options := ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: ccfg.ControllerCfg.MetricsAddr,
		LeaderElection:     ccfg.ControllerCfg.EnableLeaderElection,
		Port:               9443,
	}

	if ccfg.ControllerCfg.ClusterScoped {
		setupLog.Info("Chaos controller manager is running in cluster scoped mode.")
		// will not specific a certain namespace
	} else {
		setupLog.Info("Chaos controller manager is running in namespace scoped mode.", "targetNamespace", ccfg.ControllerCfg.TargetNamespace)
		options.Namespace = ccfg.ControllerCfg.TargetNamespace
	}

	return &options
}

func NewConfig() *rest.Config {
	return ctrl.GetConfigOrDie()
}

func NewManager(options *ctrl.Options, cfg *rest.Config) (ctrl.Manager, error) {
	if ccfg.ControllerCfg.QPS > 0 {
		cfg.QPS = ccfg.ControllerCfg.QPS
	}
	if ccfg.ControllerCfg.Burst > 0 {
		cfg.Burst = ccfg.ControllerCfg.Burst
	}

	return ctrl.NewManager(cfg, *options)
}

func NewAuthCli(cfg *rest.Config) (*authorizationv1.AuthorizationV1Client, error) {

	if ccfg.ControllerCfg.QPS > 0 {
		cfg.QPS = ccfg.ControllerCfg.QPS
	}
	if ccfg.ControllerCfg.Burst > 0 {
		cfg.Burst = ccfg.ControllerCfg.Burst
	}

	return authorizationv1.NewForConfig(cfg)
}

func NewClient(mgr ctrl.Manager) client.Client {
	return mgr.GetClient()
}

func NewLogger() logr.Logger {
	return ctrl.Log
}

func NewReader(mgr ctrl.Manager) client.Reader {
	return mgr.GetAPIReader()
}
