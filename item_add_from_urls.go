package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ItemsAddFromURLsRequest struct {
	Items    []ItemAddFromURLRequest `json:"items"`
	FolderID string                  `json:"folderId,omitempty"`
}

type ItemsAddFromURLsResponse struct {
	Status string `json:"status"`
}

func (c *Client) AddItemsFromURLs(request ItemsAddFromURLsRequest) (*ItemsAddFromURLsResponse, error) {
	url := fmt.Sprintf("%s/api/item/addFromURLs", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemsAddFromURLsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
