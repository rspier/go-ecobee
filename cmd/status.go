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
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/rspier/go-ecobee/ecobee"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display current thermostat status.",
	Long:  "Display current thermostat status.",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlags()
		c := client()
		showStatus(c)
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}

func showStatus(c *ecobee.Client) {
	tsm, err := c.GetThermostatSummary(
		ecobee.Selection{
			SelectionType:          "thermostats",
			SelectionMatch:         thermostat,
			IncludeEquipmentStatus: true,
		})

	var ts ecobee.ThermostatSummary
	var ok bool

	if ts, ok = tsm[thermostat]; !ok {
		glog.Exitf("thermostat %s missing from ThermostatSummary", thermostat)
	}

	t, err := c.GetThermostat(thermostat)
	if err != nil {
		glog.Exitf("error retrieving thermostat %s: %v", thermostat, err)
	}

	running := formatEquipmentStatus(ts)

	fmt.Printf("Current Settings (%s): %.1f - %.1f.  Fan: %s%s\n",
		strings.Title(t.Program.CurrentClimateRef),
		float64(t.Runtime.DesiredHeat/10.0),
		float64(t.Runtime.DesiredCool/10.0),
		t.Runtime.DesiredFanMode,
		running)

	ev := t.Events[0]
	if ev.Running && ev.Type == "hold" {
		fmt.Printf("Holding at %.1f - %.1f (Fan: %s) until %s %s\n",
			float64(ev.HeatHoldTemp/10.0),
			float64(ev.CoolHoldTemp/10.0),
			ev.Fan,
			ev.EndDate,
			ev.EndTime)
	}

	fmt.Printf("Temperature: %.1f\n", float64(t.Runtime.ActualTemperature/10.0))

	for _, s := range t.RemoteSensors {
		var temp, occ string
		for _, c := range s.Capability {
			if c.Type == "temperature" {
				t, err := strconv.ParseFloat(c.Value, 64)
				if err == nil {
					temp = fmt.Sprintf("%.1f", t/10.0)
				}
			}
			if c.Type == "occupancy" {
				if c.Value == "true" {
					occ = "occupied"
				}
			}
		}
		var inuse string
		if s.InUse {
			inuse = "*"
		}
		fmt.Printf("  %s%s: %s %s\n", s.Name, inuse, temp, occ)
	}

}

func formatEquipmentStatus(ts ecobee.ThermostatSummary) string {
	eqs := ""
	if ts.EquipmentStatus.Fan {
		eqs = " (running)"
	}
	if ts.EquipmentStatus.CompCool1 {
		eqs += " Cool"
	}
	if ts.EquipmentStatus.CompCool2 {
		eqs += " Cool2"
	}
	if ts.EquipmentStatus.AuxHeat1 {
		eqs += " Heat"
	}
	if ts.EquipmentStatus.AuxHeat2 {
		eqs += " Heat2"
	}
	if ts.EquipmentStatus.AuxHeat3 {
		eqs += " Heat3"
	}
	return eqs
}
