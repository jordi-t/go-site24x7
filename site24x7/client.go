package site24x7

import (
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// ClientConfig provides an interface for the client configuration
type authenticator interface {
	GetAuthToken() (string, error)
}

// Client represents a client to the Site24x7 API.
type Client struct {
	auth       authenticator
	httpClient *http.Client
	Users      *UserService
	apiDomain  string
}

// NewClient returns a Site24x7 client.
func NewClient(auth authenticator, apiDomain string, hc ...http.Client) *Client {

	var httpClient http.Client
	if hc == nil {
		httpClient = http.Client{
			Timeout: time.Second * 10,
		}
	} else {
		httpClient = hc[0]
	}

	c := &Client{
		auth:       auth,
		httpClient: &httpClient,
		apiDomain:  apiDomain,
	}

	c.Users = &UserService{client: c}

	return c
}

// Get does an HTTP-GET on a given Site24x7 API endpoint.
func (c *Client) get(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.apiDomain+endpoint, nil)
	if err != nil {
		return nil, err
	}

	authToken, err := c.auth.GetAuthToken()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get token")
	}
	req.Header.Set("Accept", "application/json; version=2.0")
	req.Header.Set("Authorization", "Zoho-oauthtoken "+authToken)

	res, err := c.httpClient.Do(req)

	return res, err
}
