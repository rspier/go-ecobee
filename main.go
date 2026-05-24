// Copyright © 2017 Google LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	goflag "flag"

	"github.com/rspier/go-ecobee/cmd"
)

func main() {
	// Parse an empty slice to mark the standard Go flag package as parsed,
	// preventing glog from implicitly calling flag.Parse() and failing on Cobra's arguments.
	_ = goflag.CommandLine.Parse([]string{})
	cmd.Execute()
}
