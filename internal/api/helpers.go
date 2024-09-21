package api

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"
)

type Response struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func (c *TechniClient) executeRequest(req *http.Request) ([]byte, error) {
	var respBody []byte

	resp, err := c.c.Do(req)
	if err != nil {
		return respBody, fmt.Errorf("Failed to execute request: %v", err)
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return respBody, fmt.Errorf("Failed to read response body: %v", err)
	}

	if resp.StatusCode > 299 {
		return respBody, fmt.Errorf("Received non-normal status code: %d, %s", resp.StatusCode, string(respBody))
	}

	err = checkResponseStatus(respBody)

	return respBody, err
}

func checkResponseStatus(respBody []byte) error {
	var r *Response

	err := json.Unmarshal(respBody, &r)
	if err != nil {
		return fmt.Errorf("Error unmarshaling technitium response: %v", err)
	}

	if r.Status != "ok" {
		return fmt.Errorf("%s", r.ErrorMessage)
	}

	return nil
}
