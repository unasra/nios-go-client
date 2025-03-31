/*
IbClient

OpenAPI 3.x.x specification for the IbClient API

API version: 3.0.0
Contact: jkhatri@infoblox.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"github.com/unasra/nios-go-client/internal"
	"github.com/unasra/nios-go-client/option"
)

const serviceBasePath = "/wapi/v2.12.3"

// APIClient manages communication with the IbClient 3.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*internal.APIClient

	// API Services
	RecordaAPI RecordaAPI
	RecordaaaaAPI RecordaaaaAPI
}

// NewAPIClient creates a new API client.
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithNIOSHostUrl(string) sets the URL for NIOS Portal.
// - WithNIOSAuth(string) sets the NIOSAuth for accessing the NIOS Portal.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultTags(map[string]string) sets the tags the client can set by default for objects that has tags support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	cfg := internal.NewConfiguration()
	for _, o := range options {
		o(cfg)
	}

	c := &APIClient{}
	c.APIClient = internal.NewAPIClient(serviceBasePath, cfg)

	// API Services
	c.RecordaAPI = (*RecordaAPIService)(&c.Common)
	c.RecordaaaaAPI = (*RecordaaaaAPIService)(&c.Common)

	return c
}
