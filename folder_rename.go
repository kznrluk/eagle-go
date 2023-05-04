package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type FolderRenameRequest struct {
	FolderID string `json:"folderId"`
	NewName  string `json:"newName"`
}

type FolderRenameResponse struct {
	Status string     `json:"status"`
	Data   FolderData `json:"data"`
}

func (c *Client) RenameFolder(folderID, newName string) (*FolderRenameResponse, error) {
	url := fmt.Sprintf("%s/api/folder/rename", c.BaseURL)
	requestData := FolderRenameRequest{
		FolderID: folderID,
		NewName:  newName,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response FolderRenameResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
