package client

import (
	"github.com/unasra/nios-go-client/dnsdata"
	"github.com/unasra/nios-go-client/option"
)

// APIClient is an aggregation of different BloxOne API clients.
type APIClient struct {
	//RecordaAPI *dnsdata.RecordaAPI
	DNSDataAPI *dnsdata.APIClient
}

// NewAPIClient creates a new BloxOne API Client.
// This is an aggregation of different BloxOne API clients.
// The following clients are available:
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithCSPUrl(string) sets the URL for BloxOne Cloud Services Portal.
// - WithAPIKey(string) sets the APIKey for accessing the BloxOne API.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultTags(map[string]string) sets the tags the client can set by default for objects that has tags support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	return &APIClient{
		DNSDataAPI: dnsdata.NewAPIClient(options...),
	}
}
