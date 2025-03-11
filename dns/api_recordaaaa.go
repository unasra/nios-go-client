/*
IbClient

OpenAPI 3.x.x specification for the IbClient API

API version: 3.0.0
Contact: jkhatri@infoblox.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/unasra/nios-go-client/internal"
)

type RecordaaaaAPI interface {
	/*
		Get Method for Get

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordaaaaAPIGetRequest
	*/
	Get(ctx context.Context) RecordaaaaAPIGetRequest

	// GetExecute executes the request
	//  @return []RecordAAAA
	GetExecute(r RecordaaaaAPIGetRequest) ([]RecordAAAA, *http.Response, error)
	/*
		Post Method for Post

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RecordaaaaAPIPostRequest
	*/
	Post(ctx context.Context) RecordaaaaAPIPostRequest

	// PostExecute executes the request
	//  @return string
	PostExecute(r RecordaaaaAPIPostRequest) (string, *http.Response, error)
	/*
		RecordaaaaReferenceDelete Method for RecordaaaaReferenceDelete

		Delete the record:aaaa resource

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param recordaaaaReference Enter the reference for record:aaaa
		@return RecordaaaaAPIRecordaaaaReferenceDeleteRequest
	*/
	RecordaaaaReferenceDelete(ctx context.Context, recordaaaaReference string) RecordaaaaAPIRecordaaaaReferenceDeleteRequest

	// RecordaaaaReferenceDeleteExecute executes the request
	//  @return string
	RecordaaaaReferenceDeleteExecute(r RecordaaaaAPIRecordaaaaReferenceDeleteRequest) (string, *http.Response, error)
	/*
		RecordaaaaReferenceGet Method for RecordaaaaReferenceGet

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param recordaaaaReference Enter the reference for record:aaaa
		@return RecordaaaaAPIRecordaaaaReferenceGetRequest
	*/
	RecordaaaaReferenceGet(ctx context.Context, recordaaaaReference string) RecordaaaaAPIRecordaaaaReferenceGetRequest

	// RecordaaaaReferenceGetExecute executes the request
	//  @return RecordAAAA
	RecordaaaaReferenceGetExecute(r RecordaaaaAPIRecordaaaaReferenceGetRequest) (*RecordAAAA, *http.Response, error)
	/*
		RecordaaaaReferencePut Method for RecordaaaaReferencePut

		Update the record:aaaa resource

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param recordaaaaReference Enter the reference for record:aaaa
		@return RecordaaaaAPIRecordaaaaReferencePutRequest
	*/
	RecordaaaaReferencePut(ctx context.Context, recordaaaaReference string) RecordaaaaAPIRecordaaaaReferencePutRequest

	// RecordaaaaReferencePutExecute executes the request
	//  @return string
	RecordaaaaReferencePutExecute(r RecordaaaaAPIRecordaaaaReferencePutRequest) (string, *http.Response, error)
}

// RecordaaaaAPIService RecordaaaaAPI service
type RecordaaaaAPIService internal.Service

type RecordaaaaAPIGetRequest struct {
	ctx              context.Context
	ApiService       RecordaaaaAPI
	returnFields     *string
	returnFields2    *string
	maxResults       *int32
	returnAsObject   *int32
	paging           *int32
	pageId           *string
	proxySearch      *string
	schema           *string
	schemaVersion    *int32
	getDoc           *int32
	schemaSearchable *int32
	inheritance      *bool
	comment          *string
	creator          *string
	ddnsPrincipal    *string
	ipv4addr         *string
	name             *string
	reclaimable      *string
	view             *string
	zone             *string
}

// Enter the field names followed by comma
func (r RecordaaaaAPIGetRequest) ReturnFields(returnFields string) RecordaaaaAPIGetRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordaaaaAPIGetRequest) ReturnFields2(returnFields2 string) RecordaaaaAPIGetRequest {
	r.returnFields2 = &returnFields2
	return r
}

// Enter the number of results to be fetched
func (r RecordaaaaAPIGetRequest) MaxResults(maxResults int32) RecordaaaaAPIGetRequest {
	r.maxResults = &maxResults
	return r
}

// Select 1 if result is required as an object
func (r RecordaaaaAPIGetRequest) ReturnAsObject(returnAsObject int32) RecordaaaaAPIGetRequest {
	r.returnAsObject = &returnAsObject
	return r
}

// Select 1 if paging is required. If SET, _max_results and _return_as_object must be entered.
func (r RecordaaaaAPIGetRequest) Paging(paging int32) RecordaaaaAPIGetRequest {
	r.paging = &paging
	return r
}

// Enter the page ID for fetching the next page
func (r RecordaaaaAPIGetRequest) PageId(pageId string) RecordaaaaAPIGetRequest {
	r.pageId = &pageId
	return r
}

// If set to GM, the request is redirected to Grid master for processing. If set to LOCAL, the request is processed locally. This option is applicable only on vConnector grid members. The default is LOCAL.
func (r RecordaaaaAPIGetRequest) ProxySearch(proxySearch string) RecordaaaaAPIGetRequest {
	r.proxySearch = &proxySearch
	return r
}

// If this option is specified, a WAPI schema will be returned
func (r RecordaaaaAPIGetRequest) Schema(schema string) RecordaaaaAPIGetRequest {
	r.schema = &schema
	return r
}

// If this option is specified, a WAPI schema of particular version will be returned. If options is omitted, schema version is assumed to be 1
func (r RecordaaaaAPIGetRequest) SchemaVersion(schemaVersion int32) RecordaaaaAPIGetRequest {
	r.schemaVersion = &schemaVersion
	return r
}

// When set to 1, it returns the documentation of the object.Applicable only when _schema_version is 2
func (r RecordaaaaAPIGetRequest) GetDoc(getDoc int32) RecordaaaaAPIGetRequest {
	r.getDoc = &getDoc
	return r
}

// If this option is specified, search only fields will also be returned. Applicable only when _schema_version is 2
func (r RecordaaaaAPIGetRequest) SchemaSearchable(schemaSearchable int32) RecordaaaaAPIGetRequest {
	r.schemaSearchable = &schemaSearchable
	return r
}

// If this option is set to True, fields which support inheritance, will display data properly.
func (r RecordaaaaAPIGetRequest) Inheritance(inheritance bool) RecordaaaaAPIGetRequest {
	r.inheritance = &inheritance
	return r
}

// Enter the value of the field
func (r RecordaaaaAPIGetRequest) Comment(comment string) RecordaaaaAPIGetRequest {
	r.comment = &comment
	return r
}

// Enter the value of the field
func (r RecordaaaaAPIGetRequest) Creator(creator string) RecordaaaaAPIGetRequest {
	r.creator = &creator
	return r
}

// Enter the value of the field
func (r RecordaaaaAPIGetRequest) DdnsPrincipal(ddnsPrincipal string) RecordaaaaAPIGetRequest {
	r.ddnsPrincipal = &ddnsPrincipal
	return r
}

// Enter the value of the field
func (r RecordaaaaAPIGetRequest) Ipv4addr(ipv4addr string) RecordaaaaAPIGetRequest {
	r.ipv4addr = &ipv4addr
	return r
}

// Enter the value of the field
func (r RecordaaaaAPIGetRequest) Name(name string) RecordaaaaAPIGetRequest {
	r.name = &name
	return r
}

// Enter the value of the field
func (r RecordaaaaAPIGetRequest) Reclaimable(reclaimable string) RecordaaaaAPIGetRequest {
	r.reclaimable = &reclaimable
	return r
}

// Enter the value of the field
func (r RecordaaaaAPIGetRequest) View(view string) RecordaaaaAPIGetRequest {
	r.view = &view
	return r
}

// Enter the value of the field
func (r RecordaaaaAPIGetRequest) Zone(zone string) RecordaaaaAPIGetRequest {
	r.zone = &zone
	return r
}

func (r RecordaaaaAPIGetRequest) Execute() ([]RecordAAAA, *http.Response, error) {
	return r.ApiService.GetExecute(r)
}

/*
Get Method for Get

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordaaaaAPIGetRequest
*/
func (a *RecordaaaaAPIService) Get(ctx context.Context) RecordaaaaAPIGetRequest {
	return RecordaaaaAPIGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []RecordAAAA
func (a *RecordaaaaAPIService) GetExecute(r RecordaaaaAPIGetRequest) ([]RecordAAAA, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue []RecordAAAA
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordaaaaAPIService.Get")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "")
	}
	if r.returnFields2 != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFields2, "")
	}
	if r.maxResults != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_max_results", r.maxResults, "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "")
	} else {
		var defaultValue int32 = 0
		r.returnAsObject = &defaultValue
	}
	if r.paging != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_paging", r.paging, "")
	} else {
		var defaultValue int32 = 0
		r.paging = &defaultValue
	}
	if r.pageId != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_id", r.pageId, "")
	}
	if r.proxySearch != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_proxy_search", r.proxySearch, "")
	}
	if r.schema != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_schema", r.schema, "")
	}
	if r.schemaVersion != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_schema_version", r.schemaVersion, "")
	}
	if r.getDoc != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_get_doc", r.getDoc, "")
	}
	if r.schemaSearchable != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_schema_searchable", r.schemaSearchable, "")
	}
	if r.inheritance != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inheritance", r.inheritance, "")
	}
	if r.comment != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "comment", r.comment, "")
	}
	if r.creator != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "creator", r.creator, "")
	}
	if r.ddnsPrincipal != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "ddns_principal", r.ddnsPrincipal, "")
	}
	if r.ipv4addr != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "ipv4addr", r.ipv4addr, "")
	}
	if r.name != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.reclaimable != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "reclaimable", r.reclaimable, "")
	}
	if r.view != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "view", r.view, "")
	}
	if r.zone != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "zone", r.zone, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordaaaaAPIPostRequest struct {
	ctx               context.Context
	ApiService        RecordaaaaAPI
	recordAAAARequest *RecordAAAARequest
	returnFields      *string
}

// Enter the request body here
func (r RecordaaaaAPIPostRequest) RecordAAAARequest(recordAAAARequest RecordAAAARequest) RecordaaaaAPIPostRequest {
	r.recordAAAARequest = &recordAAAARequest
	return r
}

// Enter the field names followed by comma
func (r RecordaaaaAPIPostRequest) ReturnFields(returnFields string) RecordaaaaAPIPostRequest {
	r.returnFields = &returnFields
	return r
}

func (r RecordaaaaAPIPostRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PostExecute(r)
}

/*
Post Method for Post

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordaaaaAPIPostRequest
*/
func (a *RecordaaaaAPIService) Post(ctx context.Context) RecordaaaaAPIPostRequest {
	return RecordaaaaAPIPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return string
func (a *RecordaaaaAPIService) PostExecute(r RecordaaaaAPIPostRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue string
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordaaaaAPIService.Post")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordAAAARequest == nil {
		return localVarReturnValue, nil, internal.ReportError("recordAAAARequest is required and must be specified")
	}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.recordAAAARequest
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordaaaaAPIRecordaaaaReferenceDeleteRequest struct {
	ctx                 context.Context
	ApiService          RecordaaaaAPI
	recordaaaaReference string
}

func (r RecordaaaaAPIRecordaaaaReferenceDeleteRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.RecordaaaaReferenceDeleteExecute(r)
}

/*
RecordaaaaReferenceDelete Method for RecordaaaaReferenceDelete

Delete the record:aaaa resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recordaaaaReference Enter the reference for record:aaaa
	@return RecordaaaaAPIRecordaaaaReferenceDeleteRequest
*/
func (a *RecordaaaaAPIService) RecordaaaaReferenceDelete(ctx context.Context, recordaaaaReference string) RecordaaaaAPIRecordaaaaReferenceDeleteRequest {
	return RecordaaaaAPIRecordaaaaReferenceDeleteRequest{
		ApiService:          a,
		ctx:                 ctx,
		recordaaaaReference: recordaaaaReference,
	}
}

// Execute executes the request
//
//	@return string
func (a *RecordaaaaAPIService) RecordaaaaReferenceDeleteExecute(r RecordaaaaAPIRecordaaaaReferenceDeleteRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue string
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordaaaaAPIService.RecordaaaaReferenceDelete")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa/{record:aaaa_reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"record:aaaa_reference"+"}", url.PathEscape(internal.ParameterValueToString(r.recordaaaaReference, "recordaaaaReference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordaaaaAPIRecordaaaaReferenceGetRequest struct {
	ctx                 context.Context
	ApiService          RecordaaaaAPI
	recordaaaaReference string
	returnFields        *string
	returnFields2       *string
	returnAsObject      *int32
}

// Enter the field names followed by comma
func (r RecordaaaaAPIRecordaaaaReferenceGetRequest) ReturnFields(returnFields string) RecordaaaaAPIRecordaaaaReferenceGetRequest {
	r.returnFields = &returnFields
	return r
}

// Enter the field names followed by comma, this returns the required fields along with the default fields
func (r RecordaaaaAPIRecordaaaaReferenceGetRequest) ReturnFields2(returnFields2 string) RecordaaaaAPIRecordaaaaReferenceGetRequest {
	r.returnFields2 = &returnFields2
	return r
}

// Select 1 if result is required as an object
func (r RecordaaaaAPIRecordaaaaReferenceGetRequest) ReturnAsObject(returnAsObject int32) RecordaaaaAPIRecordaaaaReferenceGetRequest {
	r.returnAsObject = &returnAsObject
	return r
}

func (r RecordaaaaAPIRecordaaaaReferenceGetRequest) Execute() (*RecordAAAA, *http.Response, error) {
	return r.ApiService.RecordaaaaReferenceGetExecute(r)
}

/*
RecordaaaaReferenceGet Method for RecordaaaaReferenceGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recordaaaaReference Enter the reference for record:aaaa
	@return RecordaaaaAPIRecordaaaaReferenceGetRequest
*/
func (a *RecordaaaaAPIService) RecordaaaaReferenceGet(ctx context.Context, recordaaaaReference string) RecordaaaaAPIRecordaaaaReferenceGetRequest {
	return RecordaaaaAPIRecordaaaaReferenceGetRequest{
		ApiService:          a,
		ctx:                 ctx,
		recordaaaaReference: recordaaaaReference,
	}
}

// Execute executes the request
//
//	@return RecordAAAA
func (a *RecordaaaaAPIService) RecordaaaaReferenceGetExecute(r RecordaaaaAPIRecordaaaaReferenceGetRequest) (*RecordAAAA, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *RecordAAAA
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordaaaaAPIService.RecordaaaaReferenceGet")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa/{record:aaaa_reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"record:aaaa_reference"+"}", url.PathEscape(internal.ParameterValueToString(r.recordaaaaReference, "recordaaaaReference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "")
	}
	if r.returnFields2 != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields+", r.returnFields2, "")
	}
	if r.returnAsObject != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_as_object", r.returnAsObject, "")
	} else {
		var defaultValue int32 = 0
		r.returnAsObject = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type RecordaaaaAPIRecordaaaaReferencePutRequest struct {
	ctx                 context.Context
	ApiService          RecordaaaaAPI
	recordaaaaReference string
	recordAAAARequest   *RecordAAAARequest
	returnFields        *string
}

// Enter the request body here
func (r RecordaaaaAPIRecordaaaaReferencePutRequest) RecordAAAARequest(recordAAAARequest RecordAAAARequest) RecordaaaaAPIRecordaaaaReferencePutRequest {
	r.recordAAAARequest = &recordAAAARequest
	return r
}

// Enter the field names followed by comma
func (r RecordaaaaAPIRecordaaaaReferencePutRequest) ReturnFields(returnFields string) RecordaaaaAPIRecordaaaaReferencePutRequest {
	r.returnFields = &returnFields
	return r
}

func (r RecordaaaaAPIRecordaaaaReferencePutRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.RecordaaaaReferencePutExecute(r)
}

/*
RecordaaaaReferencePut Method for RecordaaaaReferencePut

Update the record:aaaa resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recordaaaaReference Enter the reference for record:aaaa
	@return RecordaaaaAPIRecordaaaaReferencePutRequest
*/
func (a *RecordaaaaAPIService) RecordaaaaReferencePut(ctx context.Context, recordaaaaReference string) RecordaaaaAPIRecordaaaaReferencePutRequest {
	return RecordaaaaAPIRecordaaaaReferencePutRequest{
		ApiService:          a,
		ctx:                 ctx,
		recordaaaaReference: recordaaaaReference,
	}
}

// Execute executes the request
//
//	@return string
func (a *RecordaaaaAPIService) RecordaaaaReferencePutExecute(r RecordaaaaAPIRecordaaaaReferencePutRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue string
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordaaaaAPIService.RecordaaaaReferencePut")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/record:aaaa/{record:aaaa_reference}"
	localVarPath = strings.Replace(localVarPath, "{"+"record:aaaa_reference"+"}", url.PathEscape(internal.ParameterValueToString(r.recordaaaaReference, "recordaaaaReference")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordAAAARequest == nil {
		return localVarReturnValue, nil, internal.ReportError("recordAAAARequest is required and must be specified")
	}

	if r.returnFields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_return_fields", r.returnFields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.recordAAAARequest
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
