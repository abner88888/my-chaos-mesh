// Copyright Chaos Mesh Authors.
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

package ttlcontroller

import (
	"github.com/chaos-mesh/chaos-mesh/pkg/dashboard/core"
	"github.com/go-logr/logr"
)

func Bootstrap(experiment core.ExperimentStore, event core.EventStore, schedule core.ScheduleStore, workflow core.WorkflowStore, ttlc *TTLconfig, logger logr.Logger) *Controller {
	return NewController(experiment, event, schedule, workflow, ttlc, logger.WithName("ttlcontroller"))
}
