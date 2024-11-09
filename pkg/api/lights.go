package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) FetchLights() ([]string, error) {
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
		return nil, fmt.Errorf("failed to fetch lights: %s", resp.Status)
	}

	var states []struct {
		EntityID string `json:"entity_id"`
		State    string `json:"state"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&states); err != nil {
		return nil, err
	}

	var lights []string
	for _, state := range states {
		if len(state.EntityID) > 6 && state.EntityID[:6] == "light." {
			lights = append(lights, state.EntityID)
		}
	}

	return lights, nil
}

func (c *Client) TurnOffLight(lightID string) error {
	url := fmt.Sprintf("%s/api/services/light/turn_off", c.BaseURL)
	body := map[string]string{"entity_id": lightID}
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
		return fmt.Errorf("failed to turn off light: %s", resp.Status)
	}

	return nil
}

// SetLightBrightness sets the brightness of a specified light
func (c *Client) SetLightBrightness(lightID string, brightness int) error {
	url := fmt.Sprintf("%s/api/services/light/turn_on", c.BaseURL)
	body := map[string]interface{}{
		"entity_id":  lightID,
		"brightness": brightness,
	}
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
		return fmt.Errorf("failed to set light brightness: %s", resp.Status)
	}

	return nil
}
