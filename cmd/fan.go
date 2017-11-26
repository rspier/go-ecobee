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
	"time"

	"github.com/spf13/cobra"
)

var (
	fanDuration time.Duration
)

// fanCmd represents the fan command
var fanCmd = &cobra.Command{
	Use:   "fan",
	Short: "Run the fan.",
	Long:  `Run the fan for a specified time period.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlags()
		c := client()

		// should get current temperatures and use that for the fan temp
		err := c.RunFan(thermostat, fanDuration)
		if err != nil {
			log.Fatalf("RunFan error: %v", err)
		}
		fmt.Printf("Running fan for %s\n", fanDuration.String())

	},
}

func init() {
	RootCmd.AddCommand(fanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	fanCmd.Flags().DurationVarP(&fanDuration, "duration", "", 1*time.Hour, "duration")
}
