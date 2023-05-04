package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ItemAddFromURLRequest struct {
	URL              string            `json:"url"`
	Name             string            `json:"name"`
	Website          string            `json:"website,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	Annotation       string            `json:"annotation,omitempty"`
	ModificationTime int64             `json:"modificationTime,omitempty"`
	FolderID         string            `json:"folderId,omitempty"`
	Headers          map[string]string `json:"headers,omitempty"`
}

type ItemAddFromURLResponse struct {
	Status string `json:"status"`
}

func (c *Client) AddItemFromURL(request ItemAddFromURLRequest) (*ItemAddFromURLResponse, error) {
	url := fmt.Sprintf("%s/api/item/addFromURL", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemAddFromURLResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
