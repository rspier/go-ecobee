package ecobee

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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/golang/glog"
)

const thermostatAPIURL = `https://api.ecobee.com/1/thermostat`

func (c *Client) UpdateThermostat(utr UpdateThermostatRequest) error {
	j, err := json.Marshal(&utr)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}

	glog.V(1).Infof("UpdateThermostat request: %s", j)

	// everything below here can be factored out into a common POST func
	resp, err := c.Post(thermostatAPIURL, "application/json", bytes.NewReader(j))
	if err != nil {
		return fmt.Errorf("error on post request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %v", err)
	}
	resp.Body.Close()

	var s UpdateThermostatResponse
	if err = json.Unmarshal(body, &s); err != nil {
		return fmt.Errorf("error unmarshalling json: %v", err)
	}

	glog.V(1).Infof("UpdateThermostat response: %+v", s)

	if s.Status.Code == 0 {
		return nil
	}
	return fmt.Errorf("API error: %s", s.Status.Message)
}

func (c *Client) GetThermostat(thermostatID string) (*Thermostat, error) {
	// TODO: Consider factoring the generation of Selection out into
	// something else to make it more convenient to toggle the IncludeX
	// flags?
	s := Selection{
		SelectionType:  "thermostats",
		SelectionMatch: thermostatID,

		IncludeAlerts:   false,
		IncludeEvents:   true,
		IncludeProgram:  true,
		IncludeRuntime:  true,
		IncludeSettings: false,
		IncludeSensors:  true,
	}
	thermostats, err := c.GetThermostats(s)
	if err != nil {
		return nil, err
	} else if len(thermostats) != 1 {
		return nil, fmt.Errorf("got %d thermostats, wanted 1", len(thermostats))
	}
	return &thermostats[0], nil
}

func (c *Client) GetThermostats(selection Selection) ([]Thermostat, error) {
	req := GetThermostatsRequest{
		Selection: selection,
	}
	j, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json: %v", err)
	}

	// everything below here can be factored out into a common GET func
	resp, err := c.Get(fmt.Sprintf("%s?json=%s", thermostatAPIURL, j))
	if err != nil {
		return nil, fmt.Errorf("error on post request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}
	resp.Body.Close()

	var r GetThermostatsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %v", err)
	}

	glog.V(1).Infof("GetThermostats response: %s", r)

	if r.Status.Code != 0 {
		return nil, fmt.Errorf("api error %d: %v", r.Status.Code, r.Status.Message)
	}
	return r.ThermostatList, nil
}
