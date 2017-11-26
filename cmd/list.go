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

package cmd

import (
	"fmt"
	"log"

	"github.com/rspier/go-ecobee/ecobee"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered thermostats.",
	Long:  `Lists all thermostats associated with your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		list(client())
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func list(c *ecobee.Client) {
	s := ecobee.Selection{
		SelectionType: "registered",
	}
	ts, err := c.GetThermostats(s)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range ts {
		fmt.Printf("%v: %v\n", t.Identifier, t.Name)
	}
}
