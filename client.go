package tourmanager

import "net/http"

const (
	// Version is this client library's release version, reported in the
	// default User-Agent header.
	Version = "0.1.0"

	defaultBaseURL   = "https://tours.api.rezkit.app/"
	defaultUserAgent = "rezkit-tour-manager-go/" + Version
)

// Client is a client for the RezKit Tour Manager API. Create one with
// [New]. The zero value is not ready to use.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	userAgent  string
}

// Option configures a [Client], for use with [New].
type Option func(*Client)

// WithBaseURL overrides the default API base URL. Mainly useful for
// pointing the client at a mock server in tests, or a non-production
// environment.
func WithBaseURL(rawURL string) Option {
	return func(c *Client) { c.baseURL = rawURL }
}

// WithHTTPClient supplies a custom *http.Client, for example to configure
// a custom http.RoundTripper (retries, tracing, proxying). The client does
// not set a default timeout on requests; use a context deadline per call
// (recommended), or set http.Client.Timeout here.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// WithUserAgent overrides the default User-Agent header sent with every
// request.
func WithUserAgent(ua string) Option {
	return func(c *Client) { c.userAgent = ua }
}

// New creates a [Client] authenticating with the given static API key.
//
//	client := tourmanager.New(os.Getenv("REZKIT_API_KEY"))
func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
		baseURL:    defaultBaseURL,
		userAgent:  defaultUserAgent,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
