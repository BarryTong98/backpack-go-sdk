package service

import (
	"encoding/json"

	"backpack-trade-bot/types"
)

func (c *BackpackClient) Status() (*types.Status, error) {
	params := make(map[string]string)

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/status", params)
	if err != nil {
		return nil, err
	}

	var status *types.Status
	err = json.Unmarshal(resp, &status)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (c *BackpackClient) Ping() (string, error) {
	params := make(map[string]string)

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/ping", params)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func (c *BackpackClient) GetSystemTime() (uint64, error) {
	params := make(map[string]string)

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/time", params)
	if err != nil {
		return 0, err
	}
	var time uint64
	err = json.Unmarshal(resp, &time)
	if err != nil {
		return 0, err
	}

	return time, nil
}
