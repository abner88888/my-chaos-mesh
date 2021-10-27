// Copyright 2021 Chaos Mesh Authors.
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

package label

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector/generic"
)

const Name = "label"

type labelSelector struct {
	selector labels.Selector
}

var _ generic.Selector = &labelSelector{}

func (s *labelSelector) ListOption() client.ListOption {
	return client.MatchingLabelsSelector{Selector: s.selector}
}

func (s *labelSelector) ListFunc(_ client.Reader) generic.ListFunc {
	return nil
}

func (s *labelSelector) Match(_ client.Object) bool {

	return true
}

func New(spec v1alpha1.GenericSelectorSpec, _ generic.Option) (generic.Selector, error) {
	labelSelectors := spec.LabelSelectors
	expressions := spec.ExpressionSelectors

	if len(labelSelectors) == 0 && len(expressions) == 0 {
		return &labelSelector{}, nil
	}
	metav1Ls := &metav1.LabelSelector{
		MatchLabels:      labelSelectors,
		MatchExpressions: expressions,
	}
	ls, err := metav1.LabelSelectorAsSelector(metav1Ls)
	if err != nil {
		return nil, err
	}
	return &labelSelector{selector: ls}, nil
}
