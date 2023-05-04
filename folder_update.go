package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type FolderUpdateRequest struct {
	FolderID       string `json:"folderId"`
	NewName        string `json:"newName"`
	NewDescription string `json:"newDescription,omitempty"`
	NewColor       string `json:"newColor,omitempty"`
}

type FolderUpdateResponse struct {
	Status string     `json:"status"`
	Data   FolderData `json:"data"`
}

func (c *Client) UpdateFolder(folderID, newName, newDescription, newColor string) (*FolderUpdateResponse, error) {
	url := fmt.Sprintf("%s/api/folder/update", c.BaseURL)
	requestData := FolderUpdateRequest{
		FolderID:       folderID,
		NewName:        newName,
		NewDescription: newDescription,
		NewColor:       newColor,
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
	var response FolderUpdateResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
