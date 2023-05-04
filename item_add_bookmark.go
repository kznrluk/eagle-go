package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ItemAddBookmarkRequest struct {
	URL              string   `json:"url"`
	Name             string   `json:"name"`
	Base64           string   `json:"base64,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	ModificationTime int64    `json:"modificationTime,omitempty"`
	FolderID         string   `json:"folderId,omitempty"`
}

type ItemAddBookmarkResponse struct {
	Status string `json:"status"`
}

func (c *Client) AddItemBookmark(request ItemAddBookmarkRequest) (*ItemAddBookmarkResponse, error) {
	url := fmt.Sprintf("%s/api/item/addBookmark", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemAddBookmarkResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
