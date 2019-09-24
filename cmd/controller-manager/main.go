// Copyright 2019 PingCAP, Inc.
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

package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/pingcap/chaos-operator/pkg/client/clientset/versioned"
	informers "github.com/pingcap/chaos-operator/pkg/client/informers/externalversions"
	"github.com/pingcap/chaos-operator/pkg/controller"
	"github.com/pingcap/chaos-operator/pkg/controller/podchaos"
	"github.com/pingcap/chaos-operator/pkg/manager"
	"github.com/pingcap/chaos-operator/pkg/signals"
	"github.com/pingcap/chaos-operator/pkg/version"
	"github.com/robfig/cron/v3"

	"golang.org/x/sync/errgroup"

	"k8s.io/apiserver/pkg/util/logs"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	printVersion bool
	pprofPort    string
)

func init() {
	flag.BoolVar(&printVersion, "version", false, "print version information and exit")
	flag.StringVar(&pprofPort, "pprof", "10080", "controller manager pprof port")
	flag.DurationVar(&controller.ResyncDuration, "resync-duration", time.Duration(30*time.Second), "resync time of informer")

	flag.Parse()
}

func main() {
	version.PrintVersionInfo()

	if printVersion {
		os.Exit(0)
	}

	logs.InitLogs()
	defer logs.FlushLogs()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := rest.InClusterConfig()
	if err != nil {
		glog.Fatalf("failed to get config: %v", err)
	}

	kubeCli, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("failed to get kubernetes Clientset: %v", err)
	}

	cli, err := versioned.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("failed to create Clientset: %v", err)
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeCli, controller.ResyncDuration)
	informerFactory := informers.NewSharedInformerFactory(cli, controller.ResyncDuration)

	glog.Info("Starting cron engine")
	cronEngine := cron.New()
	cronEngine.Start()

	managerBase := manager.NewManagerBase(cronEngine)

	podChaosController := podchaos.NewController(
		kubeCli, cli,
		kubeInformerFactory,
		informerFactory,
		managerBase)

	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)
	informerFactory.Start(stopCh)

	g := errgroup.Group{}

	g.Go(func() error {
		return podChaosController.Run(stopCh)
	})

	g.Go(func() error {
		return http.ListenAndServe(fmt.Sprintf(":%s", pprofPort), nil)
	})

	if err := g.Wait(); err != nil {
		glog.Fatal(err)
	}
}
