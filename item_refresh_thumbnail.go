package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ItemRefreshThumbnailRequest struct {
	ID string `json:"id"`
}

type ItemRefreshThumbnailResponse struct {
	Status string `json:"status"`
}

func (c *Client) RefreshItemThumbnail(request ItemRefreshThumbnailRequest) (*ItemRefreshThumbnailResponse, error) {
	url := fmt.Sprintf("%s/api/item/refreshThumbnail", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemRefreshThumbnailResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
