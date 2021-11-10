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

package finalizer

func RemoveFromFinalizer(finalizers []string, key string) []string {
	slice := make([]string, 0, len(finalizers))
	for _, f := range finalizers {
		if f != key {
			slice = append(slice, f)
		}
	}

	return slice
}

func InsertFinalizer(finalizers []string, finalizer string) []string {
	exist := false

	for _, f := range finalizers {
		if f == finalizer {
			exist = true
		}
	}

	if exist {
		return finalizers
	}
	return append(finalizers, finalizer)
}

func ContainsFinalizer(finalizers []string, finalizer string) bool {
	for _, f := range finalizers {
		if f == finalizer {
			return true
		}
	}

	return false
}
