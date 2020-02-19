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

package controllers

import (
	"context"

	"github.com/pingcap/chaos-mesh/api/v1alpha1"
	"github.com/pingcap/chaos-mesh/controllers/networkchaos"
	"github.com/pingcap/chaos-mesh/pkg/utils"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NetworkChaosReconciler reconciles a NetworkChaos object
type NetworkChaosReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=pingcap.com,resources=networkchaos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pingcap.com,resources=networkchaos/status,verbs=get;update;patch

// Reconcile reconciles a NetworkChaos resource
func (r *NetworkChaosReconciler) Reconcile(req ctrl.Request) (result ctrl.Result, err error) {
	logger := r.Log.WithValues("reconciler", "networkchaos")

	reconciler := networkchaos.Reconciler{
		Client: r.Client,
		Log:    logger,
	}

	chaos := &v1alpha1.NetworkChaos{}
	if err := r.Get(context.Background(), req.NamespacedName, chaos); err != nil {
		r.Log.Error(err, "unable to get network chaos")
		return ctrl.Result{}, nil
	}

	result, err = reconciler.Reconcile(req, chaos)
	if !chaos.IsDeleted() {
		if err != nil {
			r.Recorder.Event(chaos, v1.EventTypeWarning, utils.EventChaosInjectFailed, err.Error())
		} else {
			r.Recorder.Event(chaos, v1.EventTypeNormal, utils.EventChaosStarted, "")
		}
	} else {
		if err != nil {
			r.Recorder.Event(chaos, v1.EventTypeWarning, utils.EventChaosRecoverFailed, err.Error())
		} else {
			r.Recorder.Event(chaos, v1.EventTypeNormal, utils.EventChaosCompleted, "")
		}
	}

	return result, nil

}

func (r *NetworkChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.NetworkChaos{}).
		Complete(r)
}
