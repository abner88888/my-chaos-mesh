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

package physicalmachinechaos

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/common"
)

type Impl struct {
	client.Client
	Log logr.Logger
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("apply physical machine chaos")

	physicalMachinechaos := obj.(*v1alpha1.PhysicalMachineChaos)
	addresses := records[index].Id
	addressArray := strings.Split(addresses, ",")

	// for example, physicalMachinechaos.Spec.Action is 'network-delay', action is 'network', subAction is 'delay'
	actions := strings.SplitN(string(physicalMachinechaos.Spec.Action), "-", 2)
	if len(actions) != 2 {
		err := errors.New("action invalid")
		return v1alpha1.NotInjected, err
	}
	action, subAction := actions[0], actions[1]
	physicalMachinechaos.Spec.ExpInfo.Action = subAction

	for _, address := range addressArray {
		url := fmt.Sprintf("%s/api/attack/%s", address, action)

		/*
			var objmap map[string]interface{}
			err := json.Unmarshal([]byte(physicalMachinechaos.Spec.ExpInfo), &objmap)
			if err != nil {
				impl.Log.Error(err, "fail to unmarshal experiment info")
				return v1alpha1.NotInjected, err
			}
			objmap["uid"] = physicalMachinechaos.Spec.UID
		*/

		expInfo, err := json.Marshal(physicalMachinechaos.Spec.ExpInfo)
		if err != nil {
			impl.Log.Error(err, "fail to marshal experiment info")
			return v1alpha1.NotInjected, err
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(expInfo))
		if err != nil {
			impl.Log.Error(err, "fail to generate HTTP request")
			return v1alpha1.NotInjected, err
		}
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			impl.Log.Error(err, "do HTTP request")
			return v1alpha1.NotInjected, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			impl.Log.Error(err, "read HTTP response body")
			return v1alpha1.NotInjected, err
		}
		impl.Log.Info("HTTP response", "address", address, "status", resp.Status, "body", string(body))

		if resp.StatusCode != http.StatusOK {
			err = errors.New("HTTP status is not OK")
			impl.Log.Error(err, "")
			return v1alpha1.NotInjected, err
		}
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("recover physical machine chaos")

	physicalMachinechaos := obj.(*v1alpha1.PhysicalMachineChaos)
	addresses := records[index].Id

	addressArray := strings.Split(addresses, ",")
	// TODO: do this in goroutine
	for _, address := range addressArray {
		url := fmt.Sprintf("%s/api/attack/%s", address, physicalMachinechaos.Spec.ExpInfo.UID)

		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			impl.Log.Error(err, "fail to generate HTTP request")
			return v1alpha1.Injected, err
		}
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			impl.Log.Error(err, "do http request")
			return v1alpha1.Injected, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			impl.Log.Error(err, "read HTTP response body")
			return v1alpha1.Injected, err
		}
		impl.Log.Info("HTTP response", "url", url, "status", resp.Status, "body", string(body))
		if resp.StatusCode != http.StatusOK {
			err = errors.New("HTTP status is not OK")
			impl.Log.Error(err, "")
			return v1alpha1.Injected, err
		}
	}

	return v1alpha1.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger) *common.ChaosImplPair {
	return &common.ChaosImplPair{
		Name:   "physicalmachinechaos",
		Object: &v1alpha1.PhysicalMachineChaos{},
		Impl: &Impl{
			Client: c,
			Log:    log.WithName("physicalmachinechaos"),
		},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
