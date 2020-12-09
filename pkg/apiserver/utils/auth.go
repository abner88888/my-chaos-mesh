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

package utils

import (
	//"fmt"
	"net/http"
	//"strings"

	"github.com/gin-gonic/gin"
	//"github.com/joomcode/errorx"
	//apierrors "k8s.io/apimachinery/pkg/api/errors"
	"github.com/chaos-mesh/chaos-mesh/pkg/clientpool"

	authorizationv1 "k8s.io/api/authorization/v1"
	//authorizationv1client "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

func CanListChaos(c *gin.Context, namespace string) bool {
	authCli, err := clientpool.ExtractTokenAndGetAuthClient(c.Request.Header)
	if err != nil {
		_ = c.Error(ErrInvalidRequest.WrapWithNoMessage(err))
		return false
	}

	sar := &authorizationv1.SelfSubjectAccessReview{
		Spec: authorizationv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Namespace: namespace,
				Verb:      "list",
				Group:     "chaos-mesh.org",
				Resource:  "*",
			},
		},
	}

	response, err := authCli.SelfSubjectAccessReviews().Create(sar)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(ErrInternalServer.NewWithNoMessage())
		return false
	}

	if !response.Status.Allowed {
		c.Status(http.StatusInternalServerError)
		if len(namespace) == 0 {
			_ = c.Error(ErrNoClusterPrivilege.New("can't list chaos experiments in the cluster"))
		} else {
			_ = c.Error(ErrNoNamespacePrivilege.New("can't list chaos experiments in namespace %s", namespace))
		}
	}

	return true
}
