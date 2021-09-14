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

package curl

import (
	"fmt"
	"net/http"
)

func Render(request RequestFlags) (Commands, error) {
	// TODO: validation of request
	result := []string{"curl", "-i", "-s"}

	// follow the request
	if request.Follow {
		result = append(result, "-L")
	}

	if request.Method != http.MethodGet {
		result = append(result, "-X", request.Method)
	}

	if len(request.Body) > 0 {
		result = append(result, "-d", fmt.Sprintf("%s", request.Body))
	}

	if request.JsonContent {
		if request.Header == nil {
			request.Header = http.Header{}
		}
		request.Header["Content-Type"] = []string{"application/json"}
	}

	for key, values := range request.Header {
		for _, value := range values {
			result = append(result, "-H", fmt.Sprintf("%s: %s", key, value))
		}
	}

	result = append(result, request.URL)

	return result, nil
}
