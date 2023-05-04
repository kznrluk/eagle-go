package eagle

import (
	"io"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient() *Client {
	return &Client{BaseURL: "http://localhost:41595"} // default port
}

func (c *Client) doRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
