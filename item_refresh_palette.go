package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ItemRefreshPaletteRequest struct {
	ID string `json:"id"`
}

type ItemRefreshPaletteResponse struct {
	Status string `json:"status"`
}

func (c *Client) RefreshItemPalette(request ItemRefreshPaletteRequest) (*ItemRefreshPaletteResponse, error) {
	url := fmt.Sprintf("%s/api/item/refreshPalette", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemRefreshPaletteResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
