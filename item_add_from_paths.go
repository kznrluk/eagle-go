package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ItemsAddFromPathsRequest struct {
	Items    []ItemAddFromPathRequest `json:"items"`
	FolderID string                   `json:"folderId,omitempty"`
}

type ItemsAddFromPathsResponse struct {
	Status string `json:"status"`
}

func (c *Client) AddItemsFromPaths(request ItemsAddFromPathsRequest) (*ItemsAddFromPathsResponse, error) {
	url := fmt.Sprintf("%s/api/item/addFromPaths", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemsAddFromPathsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
