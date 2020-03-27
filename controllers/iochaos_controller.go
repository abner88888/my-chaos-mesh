// Copyright 2020 PingCAP, Inc.
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
	"github.com/pingcap/chaos-mesh/controllers/iochaos"
	"github.com/pingcap/chaos-mesh/pkg/utils"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// IoChaosReconciler reconciles a IoChaos object
type IoChaosReconciler struct {
	client.Client
	record.EventRecorder
	Log logr.Logger
}

// +kubebuilder:rbac:groups=chaosmesh.pingcap.com,resources=iochaos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=chaosmesh.pingcap.com,resources=iochaos/status,verbs=get;update;patch

// Reconcile reconciles an IOChaos resource
func (r *IoChaosReconciler) Reconcile(req ctrl.Request) (result ctrl.Result, err error) {
	logger := r.Log.WithValues("iochaos", req.NamespacedName)

	reconciler := iochaos.Reconciler{
		Client:        r.Client,
		EventRecorder: r.EventRecorder,
		Log:           logger,
	}
	chaos := &v1alpha1.IoChaos{}
	if err := r.Get(context.Background(), req.NamespacedName, chaos); err != nil {
		r.Log.Error(err, "unable to get iochaos")
		return ctrl.Result{}, nil
	}

	result, err = reconciler.Reconcile(req, chaos)
	if err != nil {
		if !chaos.IsDeleted() {
			r.Event(chaos, v1.EventTypeWarning, utils.EventChaosInjectFailed, err.Error())
		} else {
			r.Event(chaos, v1.EventTypeWarning, utils.EventChaosRecoverFailed, err.Error())
		}
	}
	return result, nil
}

func (r *IoChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.IoChaos{}).
		Complete(r)
}
