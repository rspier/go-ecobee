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

	"github.com/spf13/cobra"
)

var (
	resumeAll bool
)

// resumeCmd represents the resume command
var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume the normally scheduled program.",
	Long:  `Resume the normally scheduled program, releasing any holds.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlags()
		c := client()

		err := c.ResumeProgram(thermostat, resumeAll)
		if err != nil {
			log.Fatalf("ResumeProgram error: %v", err)
		}
		fmt.Printf("Successfully resumed program\n")

	},
}

func init() {
	RootCmd.AddCommand(resumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	resumeCmd.Flags().BoolVar(&resumeAll, "all", false, "Resume all programs???")
}
