package eagle

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ItemUpdateRequest is the request struct to update an item.
type ItemUpdateRequest struct {
	ID         string   `json:"id"`
	Tags       []string `json:"tags,omitempty"`
	Annotation string   `json:"annotation,omitempty"`
	URL        string   `json:"url,omitempty"`
	Star       int      `json:"star,omitempty"`
}

// ItemUpdateResponse holds the response from the UpdateItem function.
type ItemUpdateResponse struct {
	Status string   `json:"status"`
	Data   ItemData `json:"data"`
}

// UpdateItem updates specified fields of an item.
func (c *Client) UpdateItem(request ItemUpdateRequest) (*ItemUpdateResponse, error) {
	url := fmt.Sprintf("%s/api/item/update", c.BaseURL)

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemUpdateResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
