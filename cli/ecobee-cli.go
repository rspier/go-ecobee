package main

// Copyright 2017 Google Inc.
//
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

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rspier/go-ecobee/ecobee"
)

var (
	authCache      = flag.String("authFile", authCacheFile, "Where to store the OAuth tokens.")
	appID          = flag.String("appid", "", "Ecobee API Application ID")
	command        = flag.String("command", "", "which command")
	coolFlag       = flag.Float64("cool", 0, "cool temp")
	duration       = flag.Duration("duration", 1*time.Hour, "duration of event")
	heatFlag       = flag.Float64("heat", 0, "heat temp")
	message        = flag.String("message", "", "message to send")
	thermostatFlag = flag.String("thermostat", "", "thermostat id")
)

const (
	authCacheFile  = "${HOME}/.go-ecobee-cli-authcache"
	thermostatFile = "${HOME}/.go-ecobee-cli-thermostat"
)

func showStatus(c *ecobee.Client) {
	t, err := c.GetThermostat(*thermostatFlag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current Settings (%s): %.1f - %.1f.  Fan: %s\n",
		strings.Title(t.Program.CurrentClimateRef),
		float64(t.Runtime.DesiredHeat/10.0),
		float64(t.Runtime.DesiredCool/10.0),
		t.Runtime.DesiredFanMode)

	ev := t.Events[0]
	if ev.Running && ev.Type == "hold" {
		fmt.Printf("Holding at %.1f - %.1f (Fan: %s) until %s %s\n",
			float64(ev.HeatHoldTemp/10.0),
			float64(ev.CoolHoldTemp/10.0),
			ev.Fan,
			ev.EndDate, ev.EndTime)
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

func listThermostats(c *ecobee.Client) {
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

func getThermostat() string {
	if *thermostatFlag != "" {
		return *thermostatFlag
	}
	t, err := ioutil.ReadFile(os.ExpandEnv(thermostatFile))
	if err != nil {
		return *thermostatFlag
	}
	return strings.TrimSpace(string(t))
}

func main() {
	flag.Parse()

	thermostat := getThermostat()
	if *command != "list" && thermostat == "" {
		log.Fatal("required flag --thermostat missing")
	}

	if *appID == "" {
		log.Fatal("required flag --appId missing")
	}

	client := ecobee.NewClient(*appID, os.ExpandEnv(*authCache))

	switch *command {
	case "list":
		listThermostats(client)

	case "fan":
		// should get current temperatures and use that for the fan temp
		err := client.RunFan(thermostat, *duration)
		if err != nil {
			log.Fatalf("RunFan error: %v", err)
		}
		fmt.Printf("Running fan for %d minutes\n", *duration)

	case "hold":
		err := client.HoldTemp(thermostat, *heatFlag, *coolFlag, *duration)
		if err != nil {
			log.Fatalf("HoldTemp error: %v", err)
		}
		fmt.Printf("Successfully held temperature\n")

	case "status":
		showStatus(client)

	case "message":
		err := client.SendMessage(thermostat, *message)
		if err != nil {
			log.Fatalf("SendMessage error: %v", err)
		}
		fmt.Printf("Successfully sent message: %q\n", *message)

	default:
		log.Fatalf("invalid command: %q.\nCommands are: list, fan, hold, status, message", *command)
	}

}
