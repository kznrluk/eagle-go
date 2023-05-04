package eagle

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// ItemListRequest is the request struct for the GetItems function.
type ItemListRequest struct {
	Limit   int
	Offset  int
	OrderBy string
	Keyword string
	Ext     string
	Tags    []string
	Folders []string
}

// ItemListResponse is the response struct for the GetItems function.
type ItemListResponse struct {
	Status string     `json:"status"`
	Data   []ItemData `json:"data"`
}

// GetItems retrieves a list of items based on filter conditions.
func (c *Client) GetItems(request ItemListRequest) (*ItemListResponse, error) {
	baseURL := fmt.Sprintf("%s/api/item/list", c.BaseURL)

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("limit", fmt.Sprint(request.Limit))
	q.Set("offset", fmt.Sprint(request.Offset))
	q.Set("orderBy", request.OrderBy)
	q.Set("keyword", request.Keyword)
	q.Set("ext", request.Ext)
	q.Set("tags", strings.Join(request.Tags, ","))
	q.Set("folders", strings.Join(request.Folders, ","))
	u.RawQuery = q.Encode()

	resp, err := c.doRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ItemListResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
