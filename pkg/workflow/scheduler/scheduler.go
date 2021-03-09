// Copyright 2020 Chaos Mesh Authors.
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

package scheduler

import (
	"context"

	"github.com/chaos-mesh/chaos-mesh/pkg/workflow/model/template"
)

/*  interface Scheduler follows the iterator-pattern, it provides templates which need instantiate. */
type Scheduler interface {
	ScheduleNext(ctx context.Context) (nextTemplates []template.Template, parentNodeName string, err error)
}
