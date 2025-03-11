# RecordaAPI

All URIs are relative to *https://172.28.83.87/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](RecordaAPI.md#Get) | **Get** /record:a | 
[**Post**](RecordaAPI.md#Post) | **Post** /record:a | 
[**RecordaReferenceDelete**](RecordaAPI.md#RecordaReferenceDelete) | **Delete** /record:a/{record:a_reference} | 
[**RecordaReferenceGet**](RecordaAPI.md#RecordaReferenceGet) | **Get** /record:a/{record:a_reference} | 
[**RecordaReferencePut**](RecordaAPI.md#RecordaReferencePut) | **Put** /record:a/{record:a_reference} | 



## Get

> []RecordA Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).ProxySearch(proxySearch).Schema(schema).SchemaVersion(schemaVersion).GetDoc(getDoc).SchemaSearchable(schemaSearchable).Inheritance(inheritance).Comment(comment).Creator(creator).DdnsPrincipal(ddnsPrincipal).Ipv4addr(ipv4addr).Name(name).Reclaimable(reclaimable).View(view).Zone(zone).Execute()



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
	resp, r, err := apiClient.RecordaAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: []RecordA
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIGetRequest` struct via the builder pattern


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

[**[]RecordA**](RecordA.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> string Post(ctx).RecordARequest(recordARequest).ReturnFields(returnFields).Execute()



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
	recordARequest := *dns.NewRecordARequest() // RecordARequest | Enter the request body here

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.Post(context.Background()).RecordARequest(recordARequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: string
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordARequest** | [**RecordARequest**](RecordARequest.md) | Enter the request body here | 
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


## RecordaReferenceDelete

> string RecordaReferenceDelete(ctx, recordaReference).Execute()





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
	recordaReference := "resourceID:resourceName" // string | Enter the reference for record:a

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.RecordaReferenceDelete(context.Background(), recordaReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.RecordaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaReferenceDelete`: string
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.RecordaReferenceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaReference** | **string** | Enter the reference for record:a | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIRecordaReferenceDeleteRequest` struct via the builder pattern


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


## RecordaReferenceGet

> RecordA RecordaReferenceGet(ctx, recordaReference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()



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
	recordaReference := "resourceID:resourceName" // string | Enter the reference for record:a

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.RecordaReferenceGet(context.Background(), recordaReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.RecordaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaReferenceGet`: RecordA
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.RecordaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaReference** | **string** | Enter the reference for record:a | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIRecordaReferenceGetRequest` struct via the builder pattern


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


## RecordaReferencePut

> string RecordaReferencePut(ctx, recordaReference).RecordARequest(recordARequest).ReturnFields(returnFields).Execute()





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
	recordaReference := "resourceID:resourceName" // string | Enter the reference for record:a
	recordARequest := *dns.NewRecordARequest() // RecordARequest | Enter the request body here

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaAPI.RecordaReferencePut(context.Background(), recordaReference).RecordARequest(recordARequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaAPI.RecordaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaReferencePut`: string
	fmt.Fprintf(os.Stdout, "Response from `RecordaAPI.RecordaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordaReference** | **string** | Enter the reference for record:a | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaAPIRecordaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordARequest** | [**RecordARequest**](RecordARequest.md) | Enter the request body here | 
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

