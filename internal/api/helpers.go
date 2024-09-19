package api

import (
    "net/http"
    "fmt"
    "io"
)


func (c *TechniClient) executeRequest(req *http.Request) ([]byte, error) {
    var respBody []byte
    resp, err := c.c.Do(req)
    if err != nil {
        return respBody, fmt.Errorf("Failed execute request to login: %v", err)
    }
    
    respBody, err = io.ReadAll(resp.Body)
    if err != nil {
        return respBody, fmt.Errorf("Failed to read response body: %v", err)
    }

    if resp.StatusCode > 299 {
        return respBody, fmt.Errorf("Received non-normal status code: %d, %s", resp.StatusCode, string(respBody))
    }

    return respBody, nil
}
