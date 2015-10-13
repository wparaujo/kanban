package gitlab

import (
	"golang.org/x/oauth2"
	"net/http"
)

type Config struct {
	BasePath string
	Domain   string
	Oauth2   *oauth2.Config
}

type GitlabContext struct {
	client *http.Client
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page string `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage string `url:"per_page,omitempty"`
}

type Transport struct {
	Token string
	Base  http.RoundTripper
}

var (
	cfg *Config
)

// New gitlab api client
func NewEngine(c *Config) {
	cfg = c
}

// AuthCodeURL returns a URL to OAuth 2.0 provider's consent page
// that asks for permissions for the required scopes explicitly.
func AuthCodeURL() string {
	return cfg.Oauth2.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// Exchange is
func Exchange(c string) (*oauth2.Token, error) {
	return cfg.Oauth2.Exchange(oauth2.NoContext, c)
}

// NewContext
func NewContext(t *oauth2.Token, pt string) *GitlabContext {
	if pt != "" {
		return &GitlabContext{
			client: &http.Client{
				Transport: &Transport{
					Base:  http.DefaultTransport,
					Token: pt,
				},
			},
		}
	}

	return &GitlabContext{
		client: cfg.Oauth2.Client(oauth2.NoContext, t),
	}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Private-Token", t.Token)
	return t.Base.RoundTrip(req)
}