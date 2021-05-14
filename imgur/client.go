// Package imgur provides a very simple wrapper around parts of the
// Imgur REST API.
package imgur

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	// BaseURL is the unchanging section of the API URL.
	BaseURL = "https://api.imgur.com/3"
)

// Client provides methods for fetching data from the API.
type Client struct {
	client *http.Client
	id     string
}

// NewClient returns a new Client that uses the given user ID for
// unauthenticated fetches.
func NewClient(id string, options ...ClientOption) *Client {
	c := Client{
		id:     id,
		client: http.DefaultClient,
	}

	for _, option := range options {
		option((*clientConfig)(&c))
	}

	return &c
}

// get fetches data from the given endpoint and puts it into data.
func (c *Client) get(data *response, endpoint ...interface{}) error {
	var urlBuilder strings.Builder
	urlBuilder.WriteString(BaseURL)
	for _, part := range endpoint {
		urlBuilder.WriteByte('/')
		urlBuilder.WriteString(url.PathEscape(fmt.Sprint(part)))
	}

	req, err := http.NewRequest(http.MethodGet, urlBuilder.String(), nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %v", c.id))

	rsp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("perform request: %w", err)
	}
	defer rsp.Body.Close()

	buf, err := io.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	err = json.Unmarshal(buf, data)
	if err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}

	return nil
}

// response is the standard, basic response format that the API returns.
type response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Status  int         `json:"status"`
}

// clientConfig is just a redefinition of Client to prevent manual
// usage of ClientOptions.
type clientConfig Client

// A ClientOption provides optional configuration for a Client.
type ClientOption func(*clientConfig)

// WithHTTPClient specifies the *http.Client to use for requests.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(config *clientConfig) {
		config.client = c
	}
}
