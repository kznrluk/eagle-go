package eagle

import (
	"encoding/json"
	"fmt"
)

// LibraryHistoryData holds the list of library paths.
type LibraryHistoryData []string

// LibraryGetHistoryResponse represents the response of the '/api/library/history' endpoint.
type LibraryGetHistoryResponse struct {
	Status string             `json:"status"`
	Data   LibraryHistoryData `json:"data"`
}

// GetLibraryHistory sends a request to '/api/library/history' endpoint and returns the library history.
func (c *Client) GetLibraryHistory() (*LibraryGetHistoryResponse, error) {
	url := fmt.Sprintf("%s/api/library/history", c.BaseURL)

	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response LibraryGetHistoryResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
