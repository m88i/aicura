package nexus

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"go.uber.org/zap"
)

const (
	userAgent              = "acuri-nexus-client"
	applicationJsonContent = "application/json"
	defaultNewAPIVersion   = "beta"
	apiPath                = "service/rest"
)

type service struct {
	client *Client
}

// Client base structure for the Nexus Client
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	username   string
	password   string
	apiVersion string
	shared     service // same instance shared among all services
	logger     *zap.SugaredLogger

	UserService *UserService
}

// ClientBuilder fluent API to build a new Nexus Client
type ClientBuilder struct {
	client *Client
}

// APIVersion sets the API version for the new Nexus Service API. Default to `beta`.
func (b *ClientBuilder) APIVersion(version string) *ClientBuilder {
	b.client.apiVersion = version
	return b
}

// WithCredentials defines the credentials to be used on each request to the Nexus Server.
func (b *ClientBuilder) WithCredentials(username, password string) *ClientBuilder {
	b.client.username = username
	b.client.password = password
	return b
}

// WithHTTPClient defines a custom `http.Client` reference
func (b *ClientBuilder) WithHTTPClient(httpClient *http.Client) *ClientBuilder {
	b.client.httpClient = httpClient
	return b
}

func (b *ClientBuilder) Verbose() *ClientBuilder {
	b.client.logger = getLogger(true)
	return b
}

// Build returns the new Nexus Client
func (b *ClientBuilder) Build() *Client {
	if b.client.httpClient == nil {
		b.client.httpClient = http.DefaultClient
	}
	if b.client.logger == nil {
		b.client.logger = getLogger(false)
	}
	return b.client
}

func NewClient(baseURL string) *ClientBuilder {
	c := &Client{}

	serverURL, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	c.baseURL = serverURL.ResolveReference(&url.URL{Path: apiPath})
	c.apiVersion = defaultNewAPIVersion
	c.shared.client = c

	// services builder
	c.UserService = (*UserService)(&c.shared)

	return &ClientBuilder{client: c}
}

// NewDefaultClient creates a new raw, straight forward Nexus Client. For a more customizable client, use `NewClient` instead
func NewDefaultClient(baseURL string) *Client {
	return NewClient(baseURL).Build()
}

func (c *Client) newRequest(method, apiPath string, body interface{}) (*http.Request, error) {
	pathReq := path.Join(c.baseURL.Path, apiPath)
	urlReq := c.baseURL.ResolveReference(&url.URL{Path: pathReq})
	c.logger.Debugf("Making a '%s' request to %s", method, urlReq)

	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, urlReq.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", applicationJsonContent)
	}
	req.Header.Set("Accept", applicationJsonContent)
	req.Header.Set("User-Agent", userAgent)
	if len(c.username) > 0 && len(c.password) > 0 {
		req.SetBasicAuth(c.username, c.password)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) appendVersion(apiPath string) string {
	return path.Join(c.apiVersion, apiPath)
}
