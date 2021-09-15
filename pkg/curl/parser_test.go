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

import "fmt"

func ExampleParseCommands() {
	cmd := []string{"curl", "-i", "-s", "-L", "https://github.com/chaos-mesh/chaos-mesh"}
	flags, err := ParseCommands(cmd)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v", flags)
	// Output: {Method:GET URL:https://github.com/chaos-mesh/chaos-mesh Header:map[] Body: FollowLocation:true JsonContent:false}
}

func ExampleParseCommands_withCustomHeader() {
	cmd := []string{"curl", "-i", "-s", "-L", "-H", "User-Agent: Go-http-client/1.1", "https://github.com/chaos-mesh/chaos-mesh"}
	flags, err := ParseCommands(cmd)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v", flags)
	// Output: {Method:GET URL:https://github.com/chaos-mesh/chaos-mesh Header:map[User-Agent:[Go-http-client/1.1]] Body: FollowLocation:true JsonContent:false}
}

func ExampleParseCommands_postJson() {
	cmd := []string{"curl", "-i", "-s", "-X", "POST", "-d", "{\"foo\": \"bar\"}", "-H", "Content-Type: application/json", "https://jsonplaceholder.typicode.com/posts"}
	flags, err := ParseCommands(cmd)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v", flags)
	// Output: {Method:POST URL:https://jsonplaceholder.typicode.com/posts Header:map[Content-Type:[application/json]] Body:{"foo": "bar"} FollowLocation:false JsonContent:true}
}
