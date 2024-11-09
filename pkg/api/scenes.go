package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ActivateScene activates a specified scene in Home Assistant
func (c *Client) ActivateScene(sceneID string) error {
	url := fmt.Sprintf("%s/api/services/scene/turn_on", c.BaseURL)
	body := map[string]string{"entity_id": sceneID}
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
		return fmt.Errorf("failed to activate scene: %s", resp.Status)
	}

	return nil
}
