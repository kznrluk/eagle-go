package eagle

import (
	"encoding/json"
	"fmt"
)

// ItemInfoResponse holds the response from the GetItemInfo function.
type ItemInfoResponse struct {
	Status string   `json:"status"`
	Data   ItemData `json:"data"`
}

// GetItemInfo retrieves the properties of a specified item.
func (c *Client) GetItemInfo(itemID string) (*ItemInfoResponse, error) {
	url := fmt.Sprintf("%s/api/item/info?id=%s", c.BaseURL, itemID)

	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
