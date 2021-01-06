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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/chaos-mesh/chaos-mesh/pkg/workflow/engine/model/workflow"
)

// +kubebuilder:object:root=true

// Workflow is the schema for the Chaos Mesh Workflow API
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the workflow
	Spec WorkflowSpec `json:"spec"`

	// +optional
	// Status represents the workflow status
	Status WorkflowStatus `json:"status"`
}

type WorkflowSpec struct {
	Entry     string     `json:"entry"`
	Templates []Template `json:"templates"`
}

type WorkflowStatus struct {
	Phase workflow.WorkflowPhase `json:"phase"`
	Nodes map[string]Node        `json:"nodes"`
}
