package api

import (
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL:    baseURL,
		Token:      token,
		HTTPClient: &http.Client{},
	}
}

// Simple function to test connection with Home Assistant
func (c *Client) GetStatus() (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/status", c.BaseURL), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get status: %s", resp.Status)
	}

	return "Connected to Home Assistant!", nil
}
