package eagle

import (
	"encoding/json"
	"fmt"
)

type ItemThumbnailResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func (c *Client) GetItemThumbnail(itemID string) (*ItemThumbnailResponse, error) {
	url := fmt.Sprintf("%s/api/item/thumbnail?id=%s", c.BaseURL, itemID)

	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemThumbnailResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
