package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// LibrarySwitchRequest represents the request data for '/api/library/switch' endpoint.
type LibrarySwitchRequest struct {
	LibraryPath string `json:"libraryPath"`
}

// LibrarySwitchResponse represents the response of the '/api/library/switch' endpoint.
type LibrarySwitchResponse struct {
	Status string `json:"status"`
}

// SwitchLibrary sends a request to '/api/library/switch' endpoint and switches the current library.
func (c *Client) SwitchLibrary(request LibrarySwitchRequest) (*LibrarySwitchResponse, error) {
	url := fmt.Sprintf("%s/api/library/switch", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response LibrarySwitchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
