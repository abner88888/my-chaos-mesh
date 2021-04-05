// Copyright 2019 Chaos Mesh Authors.
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

package twophase

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	ctx "github.com/chaos-mesh/chaos-mesh/pkg/router/context"
	"github.com/chaos-mesh/chaos-mesh/pkg/router/endpoint"
	sch "github.com/chaos-mesh/chaos-mesh/pkg/scheduler"

	ctrl "sigs.k8s.io/controller-runtime"
)

const emptyString = ""

// Reconciler for the twophase reconciler
type Reconciler struct {
	endpoint.Endpoint
	ctx.Context
}

// NewReconciler would create reconciler for twophase controller
func NewReconciler(req ctrl.Request, e endpoint.Endpoint, ctx ctx.Context) *Reconciler {
	ctx.Log = ctx.Log.WithName(req.NamespacedName.String())

	return &Reconciler{
		Endpoint: e,
		Context:  ctx,
	}
}

// Reconcile is twophase reconcile implement
func (r *Reconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	var err error
	now := time.Now()

	r.Log.Info("Reconciling a two phase chaos", "name", req.Name, "namespace", req.Namespace, "time", time.Now())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_chaos := r.Object()
	if err = r.Client.Get(ctx, req.NamespacedName, _chaos); err != nil {
		r.Log.Error(err, "unable to get chaos")
		return ctrl.Result{}, err
	}
	chaos := _chaos.(v1alpha1.InnerSchedulerObject)

	status := chaos.GetStatus()

	targetPhase := status.Experiment.Phase

	if !chaos.GetNextRecover().IsZero() && chaos.GetNextRecover().Before(now) {
		targetPhase = v1alpha1.ExperimentPhaseWaiting
	}

	if chaos.GetNextStart().Before(now) {
		targetPhase = v1alpha1.ExperimentPhaseRunning
	}

	// For standard cron(* * * * *), judge the phase when start
	if chaos.GetNextStart().IsZero() && chaos.GetNextRecover().IsZero() {
		phase, nextStart, err := ifStartAtWaitingPhase(chaos, now, r.Log)
		if err != nil {
			return ctrl.Result{}, err
		}
		if phase == v1alpha1.ExperimentPhaseWaiting {
			targetPhase = phase
			chaos.SetNextStart(*nextStart)
			duration, err := chaos.GetDuration()
			if err != nil {
				return ctrl.Result{}, err
			}
			chaos.SetNextRecover(nextStart.Add(*duration))
		}
	}

	if chaos.IsPaused() {
		targetPhase = v1alpha1.ExperimentPhasePaused
	}

	// TODO: find a better way to solve the pause and resume problem.
	// Or pause is a bad design for the scheduler :(
	if !chaos.IsPaused() && status.Experiment.Phase == v1alpha1.ExperimentPhasePaused {
		// Running and Waiting has the same logic for resuming
		targetPhase = v1alpha1.ExperimentPhaseRunning
	}

	if chaos.IsDeleted() {
		targetPhase = v1alpha1.ExperimentPhaseFinished
	}

	r.Log.Info("decide target phase", "target phase", targetPhase)

	machine := chaosStateMachine{
		Chaos:      chaos,
		Req:        req,
		Reconciler: r,
	}
	err = machine.Into(ctx, targetPhase, now)
	if err != nil {
		r.Log.Error(err, "fail to step into the phase", "target phase", targetPhase)
		return ctrl.Result{}, err
	}

	// the reconciliation of Finished and Paused resource shouldn't be triggered by time
	if chaos.GetStatus().Experiment.Phase == v1alpha1.ExperimentPhaseFinished ||
		chaos.GetStatus().Experiment.Phase == v1alpha1.ExperimentPhasePaused {
		return ctrl.Result{}, nil
	}

	requeueAfter, err := calcRequeueAfterTime(chaos, now)
	if err != nil {
		r.Log.Error(err, "unexpected time", "now", now, "nextStart", chaos.GetNextStart(), "nextRecover", chaos.GetNextRecover())

		// will not return error and retry
		// because nothing will be better with retrying
		return ctrl.Result{}, nil
	}

	r.Log.Info("requeue", "requeue after", requeueAfter)
	if requeueAfter == time.Duration(0) {
		return ctrl.Result{Requeue: true}, nil
	}

	return ctrl.Result{
		RequeueAfter: requeueAfter,
	}, nil
}

func calcRequeueAfterTime(chaos v1alpha1.InnerSchedulerObject, now time.Time) (time.Duration, error) {
	requeueAfter := time.Duration(math.MaxInt64)
	// requeueAfter = min(filter([nextRecoverAfter, nextStartAfter], >0))
	nextRecoverAfter := chaos.GetNextRecover().Sub(now)
	nextStartAfter := chaos.GetNextStart().Sub(now)

	// If the `nextRecoverAfter` is less zero, we reset `nextRecoverAfter` to zero,
	// which represents that the chaos action should be recovered immediately.
	// `chaos.GetNextRecover()` ignores millisecond field, which will cause the `nextRecoverAfter` is less zero
	// when the duration is zero or less 1 second.
	if nextRecoverAfter < 0 {
		nextRecoverAfter = 0 * time.Second
	}

	if requeueAfter > nextRecoverAfter {
		requeueAfter = nextRecoverAfter
	}

	if nextStartAfter > 0 && requeueAfter > nextStartAfter {
		requeueAfter = nextStartAfter
	}

	var err error
	if requeueAfter == math.MaxInt64 {
		err = errors.Errorf("unexpected behavior, now is greater than nextRecover and nextStart")
	}

	return requeueAfter, err
}

func ifStartAtWaitingPhase(chaos v1alpha1.InnerSchedulerObject, now time.Time, log logr.Logger) (v1alpha1.ExperimentPhase, *time.Time, error) {
	duration, err := chaos.GetDuration()
	if err != nil {
		log.Error(err, "failed to get chaos duration")
		return "", nil, err
	}
	if duration == nil {
		zero := 0 * time.Second
		duration = &zero
	}

	scheduler := chaos.GetScheduler()
	if scheduler == nil {
		log.Info("Scheduler should be defined currently")
		return "", nil, fmt.Errorf("misdefined scheduler")
	}

	lastStart, err := sch.LastTime(*scheduler, now)
	if err != nil {
		log.Error(err, "failed to get the last start time")
		return "", nil, err
	}
	if lastStart.Add(*duration).Before(now) {
		nextStart, err := nextTime(*scheduler, now)
		if err != nil {
			log.Error(err, "failed to get the next start time")
			return "", nil, err
		}
		return v1alpha1.ExperimentPhaseWaiting, nextStart, nil
	}
	return "", nil, nil
}

func nextTime(spec v1alpha1.SchedulerSpec, now time.Time) (*time.Time, error) {
	scheduler, err := cron.ParseStandard(spec.Cron)
	if err != nil {
		return nil, fmt.Errorf("fail to parse runner rule %s, %v", spec.Cron, err)
	}

	next := scheduler.Next(now)
	return &next, nil
}
