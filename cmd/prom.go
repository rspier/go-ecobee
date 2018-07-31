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
	"log"
	"strconv"

	"github.com/golang/glog"
	"github.com/rspier/go-ecobee/ecobee"
	"github.com/spf13/cobra"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
)

var pushGateway string

// promCmd represents the status command
var promCmd = &cobra.Command{
	Use:   "prompush",
	Short: "Push thermostat status to prometheus.",
	Long:  "Push thermostat status to prometheus.",
	Run: func(cmd *cobra.Command, args []string) {
		checkRequiredFlags()
		c := client()

		tsm, err := c.GetThermostatSummary(
			ecobee.Selection{
				SelectionType:          "thermostats",
				SelectionMatch:         thermostat,
				IncludeEquipmentStatus: true,
			})
		if err != nil {
			glog.Exitf("error retrieving thermostat summary for %s: %v", thermostat, err)
		}

		var ts ecobee.ThermostatSummary
		var ok bool

		if ts, ok = tsm[thermostat]; !ok {
			glog.Exitf("thermostat %s missing from ThermostatSummary", thermostat)
		}

		t, err := c.GetThermostat(thermostat)
		if err != nil {
			glog.Exitf("error retrieving thermostat %s: %v", thermostat, err)
		}

		if pushGateway == "" {
			glog.Exit("required flag --pushgateway missing")
		}

		promPush(c, &ts, t)
	},
}

func init() {
	RootCmd.AddCommand(promCmd)
	promCmd.Flags().StringVarP(&pushGateway, "pushgateway", "p", "", "URL of prometheus push gateway")
}

func promPush(c *ecobee.Client, ts *ecobee.ThermostatSummary, t *ecobee.Thermostat) {

	gauges := []struct {
		name string
		val  float64
	}{
		{"fan", boolToFloat(ts.EquipmentStatus.Fan)},
		{"comp_cool1", boolToFloat(ts.EquipmentStatus.CompCool1)},
		{"comp_cool2", boolToFloat(ts.EquipmentStatus.CompCool2)},

		{"aux_heat1", boolToFloat(ts.EquipmentStatus.AuxHeat1)},
		{"aux_heat2", boolToFloat(ts.EquipmentStatus.AuxHeat2)},
		{"aux_heat3", boolToFloat(ts.EquipmentStatus.AuxHeat3)},

		{"desired_heat", float64(t.Runtime.DesiredHeat) / 10.0},
		{"desired_cool", float64(t.Runtime.DesiredCool) / 10.0},
		{"temperature", float64(t.Runtime.ActualTemperature) / 10.0},
	}
	for _, i := range gauges {
		g := promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: i.name,
				Help: i.name,
			},
		)
		g.Set(i.val)
	}

	sensorTemp := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sensor_temperature",
			Help: "Description",
		},
		[]string{"name"},
	)
	sensorOccupied := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sensor_occupied",
			Help: "Description",
		},
		[]string{"name"},
	)
	for _, s := range t.RemoteSensors {
		for _, c := range s.Capability {
			if c.Type == "temperature" {
				t, err := strconv.ParseFloat(c.Value, 64)
				if err == nil {
					g, _ := sensorTemp.GetMetricWithLabelValues(s.Name)
					g.Set(t / 10.0)
				}
			}
			if c.Type == "occupancy" {
				g, _ := sensorOccupied.GetMetricWithLabelValues(s.Name)
				g.Set(stringBoolToFloat(c.Value))
			}
		}
	}

	err := push.New(pushGateway, "ecobee").Grouping("instance", "home").Gatherer(prometheus.DefaultGatherer).Push()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

}
