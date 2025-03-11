# RecordaaaaAPI

All URIs are relative to *https://172.28.83.87/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](RecordaaaaAPI.md#Get) | **Get** /record:aaaa | 
[**Post**](RecordaaaaAPI.md#Post) | **Post** /record:aaaa | 
[**RecordaaaaReferenceDelete**](RecordaaaaAPI.md#RecordaaaaReferenceDelete) | **Delete** /record:aaaa/{record:aaaa_reference} | 
[**RecordaaaaReferenceGet**](RecordaaaaAPI.md#RecordaaaaReferenceGet) | **Get** /record:aaaa/{record:aaaa_reference} | 
[**RecordaaaaReferencePut**](RecordaaaaAPI.md#RecordaaaaReferencePut) | **Put** /record:aaaa/{record:aaaa_reference} | 



## Get

> []RecordAAAA Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).ProxySearch(proxySearch).Schema(schema).SchemaVersion(schemaVersion).GetDoc(getDoc).SchemaSearchable(schemaSearchable).Inheritance(inheritance).Comment(comment).Creator(creator).DdnsPrincipal(ddnsPrincipal).Ipv4addr(ipv4addr).Name(name).Reclaimable(reclaimable).View(view).Zone(zone).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: []RecordAAAA
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | [default to 0]
**paging** | **int32** | Select 1 if paging is required. If SET, _max_results and _return_as_object must be entered. | [default to 0]
**pageId** | **string** | Enter the page ID for fetching the next page | 
**proxySearch** | **string** | If set to GM, the request is redirected to Grid master for processing. If set to LOCAL, the request is processed locally. This option is applicable only on vConnector grid members. The default is LOCAL. | 
**schema** | **string** | If this option is specified, a WAPI schema will be returned | 
**schemaVersion** | **int32** | If this option is specified, a WAPI schema of particular version will be returned. If options is omitted, schema version is assumed to be 1 | 
**getDoc** | **int32** | When set to 1, it returns the documentation of the object.Applicable only when _schema_version is 2 | 
**schemaSearchable** | **int32** | If this option is specified, search only fields will also be returned. Applicable only when _schema_version is 2 | 
**inheritance** | **bool** | If this option is set to True, fields which support inheritance, will display data properly. | 
**comment** | **string** | Enter the value of the field | 
**creator** | **string** | Enter the value of the field | 
**ddnsPrincipal** | **string** | Enter the value of the field | 
**ipv4addr** | **string** | Enter the value of the field | 
**name** | **string** | Enter the value of the field | 
**reclaimable** | **string** | Enter the value of the field | 
**view** | **string** | Enter the value of the field | 
**zone** | **string** | Enter the value of the field | 

### Return type

[**[]RecordAAAA**](RecordAAAA.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> string Post(ctx).RecordAAAARequest(recordAAAARequest).ReturnFields(returnFields).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dns"
)

func main() {
	recordAAAARequest := *dns.NewRecordAAAARequest() // RecordAAAARequest | Enter the request body here

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.Post(context.Background()).RecordAAAARequest(recordAAAARequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: string
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAAAARequest** | [**RecordAAAARequest**](RecordAAAARequest.md) | Enter the request body here | 
**returnFields** | **string** | Enter the field names followed by comma | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaaaaReferenceDelete

> string RecordaaaaReferenceDelete(ctx, recordaaaaReference).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dns"
)

func main() {
	recordaaaaReference := "resourceID:resourceName" // string | Enter the reference for record:aaaa

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.RecordaaaaReferenceDelete(context.Background(), recordaaaaReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.RecordaaaaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaaaaReferenceDelete`: string
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.RecordaaaaReferenceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaaaaReference** | **string** | Enter the reference for record:aaaa | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIRecordaaaaReferenceDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaaaaReferenceGet

> RecordAAAA RecordaaaaReferenceGet(ctx, recordaaaaReference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dns"
)

func main() {
	recordaaaaReference := "resourceID:resourceName" // string | Enter the reference for record:aaaa

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.RecordaaaaReferenceGet(context.Background(), recordaaaaReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.RecordaaaaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaaaaReferenceGet`: RecordAAAA
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.RecordaaaaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaaaaReference** | **string** | Enter the reference for record:aaaa | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIRecordaaaaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | [default to 0]

### Return type

[**RecordAAAA**](RecordAAAA.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaaaaReferencePut

> string RecordaaaaReferencePut(ctx, recordaaaaReference).RecordAAAARequest(recordAAAARequest).ReturnFields(returnFields).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/unasra/nios-go-client/dns"
)

func main() {
	recordaaaaReference := "resourceID:resourceName" // string | Enter the reference for record:aaaa
	recordAAAARequest := *dns.NewRecordAAAARequest() // RecordAAAARequest | Enter the request body here

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.RecordaaaaReferencePut(context.Background(), recordaaaaReference).RecordAAAARequest(recordAAAARequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.RecordaaaaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaaaaReferencePut`: string
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.RecordaaaaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaaaaReference** | **string** | Enter the reference for record:aaaa | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIRecordaaaaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAAAARequest** | [**RecordAAAARequest**](RecordAAAARequest.md) | Enter the request body here | 
**returnFields** | **string** | Enter the field names followed by comma | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

