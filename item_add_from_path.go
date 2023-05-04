package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ItemAddFromPathRequest struct {
	Path       string   `json:"path"`
	Name       string   `json:"name"`
	Website    string   `json:"website,omitempty"`
	Annotation string   `json:"annotation,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	FolderID   string   `json:"folderId,omitempty"`
}

type ItemAddFromPathResponse struct {
	Status string `json:"status"`
}

func (c *Client) AddItemFromPath(request ItemAddFromPathRequest) (*ItemAddFromPathResponse, error) {
	u := fmt.Sprintf("%s/api/item/addFromPath", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", u, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemAddFromPathResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
