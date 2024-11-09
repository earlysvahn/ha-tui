package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchSwitches retrieves a list of all switches from Home Assistant
func (c *Client) FetchSwitches() ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/states", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch switches: %s", resp.Status)
	}

	// Parse response to find switches
	var states []struct {
		EntityID string `json:"entity_id"`
		State    string `json:"state"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&states); err != nil {
		return nil, err
	}

	// Filter for switches
	var switches []string
	for _, state := range states {
		if len(state.EntityID) > 7 && state.EntityID[:7] == "switch." {
			switches = append(switches, state.EntityID)
		}
	}

	return switches, nil
}

// TurnOnSwitch sends a request to turn on a specified switch
func (c *Client) TurnOnSwitch(switchID string) error {
	url := fmt.Sprintf("%s/api/services/switch/turn_on", c.BaseURL)
	body := map[string]string{"entity_id": switchID}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to turn on switch: %s", resp.Status)
	}

	return nil
}

// TurnOffSwitch sends a request to turn off a specified switch
func (c *Client) TurnOffSwitch(switchID string) error {
	url := fmt.Sprintf("%s/api/services/switch/turn_off", c.BaseURL)
	body := map[string]string{"entity_id": switchID}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to turn off switch: %s", resp.Status)
	}

	return nil
}
