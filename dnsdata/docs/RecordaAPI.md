# RecordaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WapiV2123RecordaGet**](RecordaAPI.md#WapiV2123RecordaGet) | **Get** /wapi/v2.12.3/record:a | List all A records
[**WapiV2123RecordaPost**](RecordaAPI.md#WapiV2123RecordaPost) | **Post** /wapi/v2.12.3/record:a | 
[**WapiV2123RecordaRecordaReferenceDelete**](RecordaAPI.md#WapiV2123RecordaRecordaReferenceDelete) | **Delete** /wapi/v2.12.3/record:a/{record:a_reference} | 
[**WapiV2123RecordaRecordaReferenceGet**](RecordaAPI.md#WapiV2123RecordaRecordaReferenceGet) | **Get** /wapi/v2.12.3/record:a/{record:a_reference} | 
[**WapiV2123RecordaRecordaReferencePut**](RecordaAPI.md#WapiV2123RecordaRecordaReferencePut) | **Put** /wapi/v2.12.3/record:a/{record:a_reference} | 



## WapiV2123RecordaGet

> RecordA WapiV2123RecordaGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Comment(comment).Creator(creator).DdnsPrincipal(ddnsPrincipal).Ipv4addr(ipv4addr).Name(name).Reclaimable(reclaimable).View(view).Zone(zone).Execute()

List all A records

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dnsdata"
)

func main() {

	apiClient := dnsdata.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.WapiV2123RecordaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.WapiV2123RecordaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WapiV2123RecordaGet`: RecordA
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.WapiV2123RecordaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIWapiV2123RecordaGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | [default to 0]
**paging** | **int32** | Select 1 if paging is required. If SET, _max_results and _return_as_object must be entered. | [default to 0]
**pageId** | **string** | Enter the page ID for fetching the next page | 
**comment** | **string** | Enter the value of the field | 
**creator** | **string** | Enter the value of the field | 
**ddnsPrincipal** | **string** | Enter the value of the field | 
**ipv4addr** | **string** | Enter the value of the field | 
**name** | **string** | Enter the value of the field | 
**reclaimable** | **string** | Enter the value of the field | 
**view** | **string** | Enter the value of the field | 
**zone** | **string** | Enter the value of the field | 

### Return type

[**RecordA**](RecordA.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WapiV2123RecordaPost

> RecordA WapiV2123RecordaPost(ctx).RecordARequest(recordARequest).ReturnFields(returnFields).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dnsdata"
)

func main() {
	recordARequest := *dnsdata.NewRecordARequest() // RecordARequest | Enter the request body here

	apiClient := dnsdata.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.WapiV2123RecordaPost(context.Background()).RecordARequest(recordARequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.WapiV2123RecordaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WapiV2123RecordaPost`: RecordA
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.WapiV2123RecordaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIWapiV2123RecordaPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordARequest** | [**RecordARequest**](RecordARequest.md) | Enter the request body here | 
**returnFields** | **string** | Enter the field names followed by comma | 

### Return type

[**RecordA**](RecordA.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WapiV2123RecordaRecordaReferenceDelete

> WapiV2123RecordARecordAReferenceDelete200Response WapiV2123RecordaRecordaReferenceDelete(ctx, recordaReference).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dnsdata"
)

func main() {
	recordaReference := "resourceID:resourceName" // string | Enter the reference for record:a

	apiClient := dnsdata.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.WapiV2123RecordaRecordaReferenceDelete(context.Background(), recordaReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.WapiV2123RecordaRecordaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WapiV2123RecordaRecordaReferenceDelete`: WapiV2123RecordARecordAReferenceDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.WapiV2123RecordaRecordaReferenceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaReference** | **string** | Enter the reference for record:a | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIWapiV2123RecordaRecordaReferenceDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**WapiV2123RecordARecordAReferenceDelete200Response**](WapiV2123RecordARecordAReferenceDelete200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WapiV2123RecordaRecordaReferenceGet

> RecordA WapiV2123RecordaRecordaReferenceGet(ctx, recordaReference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dnsdata"
)

func main() {
	recordaReference := "resourceID:resourceName" // string | Enter the reference for record:a

	apiClient := dnsdata.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.WapiV2123RecordaRecordaReferenceGet(context.Background(), recordaReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.WapiV2123RecordaRecordaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WapiV2123RecordaRecordaReferenceGet`: RecordA
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.WapiV2123RecordaRecordaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaReference** | **string** | Enter the reference for record:a | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIWapiV2123RecordaRecordaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | [default to 0]

### Return type

[**RecordA**](RecordA.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WapiV2123RecordaRecordaReferencePut

> RecordA WapiV2123RecordaRecordaReferencePut(ctx, recordaReference).RecordARequest(recordARequest).ReturnFields(returnFields).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dnsdata"
)

func main() {
	recordaReference := "resourceID:resourceName" // string | Enter the reference for record:a
	recordARequest := *dnsdata.NewRecordARequest() // RecordARequest | Enter the request body here

	apiClient := dnsdata.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.WapiV2123RecordaRecordaReferencePut(context.Background(), recordaReference).RecordARequest(recordARequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.WapiV2123RecordaRecordaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WapiV2123RecordaRecordaReferencePut`: RecordA
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.WapiV2123RecordaRecordaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaReference** | **string** | Enter the reference for record:a | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIWapiV2123RecordaRecordaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordARequest** | [**RecordARequest**](RecordARequest.md) | Enter the request body here | 
**returnFields** | **string** | Enter the field names followed by comma | 

### Return type

[**RecordA**](RecordA.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

