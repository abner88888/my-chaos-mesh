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
//

package v1alpha1

const (
	// TemplateTypeLabelKey is a label that represents the template type.
	TemplateTypeLabelKey = "template.chaos-mesh.org/type"
	// ManagedByLabelKey is a label that represents the tool being used
	// to manage the operation of the object.
	ManagedByLabelKey = "app.kubernetes.io/managed-by"

	// TemplateNameAnnotationKey is an annotation that represents
	// the real name of the template.
	TemplateNameAnnotationKey = "template.chaos-mesh.org/name"
	// TemplateDescriptionAnnotationKey is an annotation that represents
	// the description of the template.
	TemplateDescriptionAnnotationKey = "template.chaos-mesh.org/description"

	// PrefixStatusCheckTemplate is the prefix of the name of a StatusCheckTemplate.
	PrefixStatusCheckTemplate = "template-status-check"
)

// StatusCheckTemplate represents a template of status check.
// A statusCheckTemplate would save in the ConfigMap named `template-status-check-<template-name>`.
// +kubebuilder:object:generate=false
type StatusCheckTemplate struct {
	StatusCheckSpec
}
