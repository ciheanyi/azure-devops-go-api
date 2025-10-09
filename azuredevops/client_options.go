package azuredevops

import (
	"net/http"
)

// ClientOptionFunc can be used customize a new AzureDevops API client.
type ClientOptionFunc func(*Client)

// WithHTTPClient can be used to configure a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) {
		c.client = httpClient
	}
}

// WithCAEEnabled enables Continuous Access Evaluation support
func WithCAEEnabled() ClientOptionFunc {
	return func(c *Client) {
		c.clientCapabilities = append(c.clientCapabilities, caeClientCapability)
	}
}

// WithTokenRefreshHandler sets a token refresh handler for CAE
func WithTokenRefreshHandler(handler TokenRefreshHandler) ClientOptionFunc {
	return func(c *Client) {
		c.tokenRefreshHandler = handler
	}
}

// WithClientCapabilities sets custom client capabilities
func WithClientCapabilities(capabilities []string) ClientOptionFunc {
	return func(c *Client) {
		c.clientCapabilities = capabilities
	}
}
