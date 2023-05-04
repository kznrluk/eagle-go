package eagle

import (
	"encoding/json"
)

type FolderListResentData struct {
	ID                   string        `json:"id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	Children             []interface{} `json:"children"`
	ModificationTime     int64         `json:"modificationTime"`
	Tags                 []string      `json:"tags"`
	ImageCount           int           `json:"imageCount"`
	DescendantImageCount int           `json:"descendantImageCount"`
	Pinyin               string        `json:"pinyin"`
	ExtendTags           []string      `json:"extendTags"`
	Password             string        `json:"password,omitempty"`
	PasswordTips         string        `json:"passwordTips,omitempty"`
	Images               []interface{} `json:"images,omitempty"`
	IsExpand             bool          `json:"isExpand,omitempty"`
	NewFolderName        string        `json:"newFolderName,omitempty"`
	ImagesMappings       interface{}   `json:"imagesMappings,omitempty"`
}

type FolderListRecentResponse struct {
	Status string                 `json:"status"`
	Data   []FolderListResentData `json:"data"`
}

func (c *Client) GetFolderListRecent() (*FolderListRecentResponse, error) {
	url := c.BaseURL + "/api/folder/listRecent"
	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response FolderListRecentResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
