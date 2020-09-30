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
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/tedpearson/go-ecobee/ecobee"
	"github.com/spf13/cobra"
)

var (
	heat, cool float64
	duration   time.Duration
	relativeRe = regexp.MustCompile(`^[+-]\d+$`)
)

// holdCmd represents the hold command
var holdCmd = &cobra.Command{
	Use:   "hold [optional relative temp]",
	Short: "Program a hold",
	Long:  `Set a hold status on the thermostat to keep the temperature between the specified heat and cool points.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlags()
		c := client()

		if heat == 0 && cool == 0 && len(args) > 0 {
			if !relativeRe.MatchString(args[0]) {
				glog.Exitf("Invalid relative temperature: %q", args[0])
			}

			rel, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				glog.Exitf("Invalid relative temperature: %q", args[0])
			}

			if math.Abs(rel) > 2 {
				glog.Exitf("Maximum relative temperature 2, got %.0f", rel)
			}

			t, err := c.GetThermostat(thermostat)
			if err != nil {
				glog.Exitf("error retrieving thermostat %s: %v", thermostat, err)
			}

			heat = rel + float64(t.Runtime.DesiredHeat/10.0)
			cool = rel + float64(t.Runtime.DesiredCool/10.0)
		}

		setHold(c, heat, cool, duration)
	},
}

func init() {
	RootCmd.AddCommand(holdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// holdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// holdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	holdCmd.Flags().Float64VarP(&heat, "heat", "", 0, "heat temp")
	holdCmd.Flags().Float64VarP(&cool, "cool", "", 0, "cool temp")
	holdCmd.Flags().DurationVarP(&duration, "duration", "", 1*time.Hour, "duration")
}

func setHold(c *ecobee.Client, heat, cool float64, duration time.Duration) {

	err := c.HoldTemp(thermostat, heat, cool, duration)
	if err != nil {
		glog.Exitf("HoldTemp error: %v", err)
	}
	fmt.Printf("Successfully held temperature between %0.1f and %0.1f for %v\n", heat, cool, duration)
}
