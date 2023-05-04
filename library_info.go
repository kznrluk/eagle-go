package eagle

import (
	"encoding/json"
	"fmt"
)

// LibraryInfoData holds the information about the library.
type LibraryInfoData struct {
	Folders            []FolderData `json:"folders"`
	SmartFolders       []FolderData `json:"smartFolders"`
	QuickAccess        []FolderData `json:"quickAccess"`
	TagsGroups         []FolderData `json:"tagsGroups"`
	ModificationTime   int64        `json:"modificationTime"`
	ApplicationVersion string       `json:"applicationVersion"`
}

// LibraryGetInfoResponse represents the response of the '/api/library/info' endpoint.
type LibraryGetInfoResponse struct {
	Status string          `json:"status"`
	Data   LibraryInfoData `json:"data"`
}

// GetLibraryInfo sends a request to '/api/library/info' endpoint and returns the library information.
func (c *Client) GetLibraryInfo() (*LibraryGetInfoResponse, error) {
	url := fmt.Sprintf("%s/api/library/info", c.BaseURL)

	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response LibraryGetInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
