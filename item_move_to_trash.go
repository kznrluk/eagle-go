package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ItemMoveToTrashRequest struct {
	ItemIDs []string `json:"itemIds"`
}

type ItemMoveToTrashResponse struct {
	Status string `json:"status"`
}

func (c *Client) MoveItemsToTrash(request ItemMoveToTrashRequest) (*ItemMoveToTrashResponse, error) {
	url := fmt.Sprintf("%s/api/item/moveToTrash", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemMoveToTrashResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
