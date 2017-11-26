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
	"strings"

	"github.com/spf13/cobra"
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Display a message on the thermostat.",
	Long:  `Display a message on the thermostat`,
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlags()
		c := client()

		m := strings.Join(args, " ")
		err := c.SendMessage(thermostat, m)
		if err != nil {
			log.Fatalf("SendMessage error: %v", err)
		}
		fmt.Printf("Successfully sent message: %q\n", m)

	},
}

func init() {
	RootCmd.AddCommand(messageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// messageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// messageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
