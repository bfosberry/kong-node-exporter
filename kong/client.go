package kong

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	timeout = 500 * time.Millisecond
)

// Client provides an interface for interacting with the Kong Admin API
type Client interface {
	GetStatus() (*Status, error)
}

type httpClient struct {
	URL    string
	client *http.Client
}

// NewHTTPClient creates a new Kong client via HTTP using the host and port
// provided to establish a connection
func NewHTTPClient(scheme, host string, port int) Client {
	url := fmt.Sprintf("%s://%s:%d/status", scheme, host, port)
	return &httpClient{
		URL: url,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *httpClient) GetStatus() (*Status, error) {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	status := new(Status)

	err = json.Unmarshal(bodyBytes, status)
	return status, err
}
