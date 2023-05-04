package eagle

import (
	"encoding/json"
)

type FolderListResponse struct {
	Status string `json:"status"`
	Data   []struct {
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
	} `json:"data"`
}

func (c *Client) GetFolderList() (*FolderListResponse, error) {
	url := c.BaseURL + "/api/folder/list"
	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response FolderListResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
