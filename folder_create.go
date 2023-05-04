package eagle

// POST Create a folder. The created folder will be put at the bottom of the folder list of the current library.
// https://api.eagle.cool/folder/create

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type FolderCreateRequest struct {
	FolderName string `json:"folderName"`
	Parent     string `json:"parent,omitempty"`
}

type FolderCreateResponse struct {
	Status string `json:"status"`
	Data   struct {
		ID               string                 `json:"id"`
		Name             string                 `json:"name"`
		Images           []interface{}          `json:"images"`
		Folders          []interface{}          `json:"folders"`
		ModificationTime int64                  `json:"modificationTime"`
		ImagesMappings   map[string]interface{} `json:"imagesMappings"`
		Tags             []interface{}          `json:"tags"`
		Children         []interface{}          `json:"children"`
		IsExpand         bool                   `json:"isExpand"`
	} `json:"data"`
}

func (c *Client) CreateFolder(request FolderCreateRequest) (*FolderCreateResponse, error) {
	url := fmt.Sprintf("%s/api/folder/create", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response FolderCreateResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
