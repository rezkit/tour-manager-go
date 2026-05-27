# \ItineraryAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItineraryEntry**](ItineraryAPI.md#CreateItineraryEntry) | **Post** /holidays/versions/{version_id}/itinerary | Create Itinerary Entry
[**DeleteItineraryEntry**](ItineraryAPI.md#DeleteItineraryEntry) | **Delete** /holidays/versions/{version_id}/itinerary/{entry_id} | Delete an Itinerary Entry
[**ListItineraryEntries**](ItineraryAPI.md#ListItineraryEntries) | **Get** /holidays/versions/{version_id}/itinerary | List Itinerary Entries
[**UpdateItineraryEntry**](ItineraryAPI.md#UpdateItineraryEntry) | **Patch** /holidays/versions/{version_id}/itinerary/{entry_id} | Update Itinerary Entry



## CreateItineraryEntry

> ItineraryEntry CreateItineraryEntry(ctx, versionId).ItineraryEntryProperties(itineraryEntryProperties).Execute()

Create Itinerary Entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager-go"
)

func main() {
	versionId := "versionId_example" // string | Holiday Version ID
	itineraryEntryProperties := *openapiclient.NewItineraryEntryProperties(int32(123), int32(123), "Title_example") // ItineraryEntryProperties | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItineraryAPI.CreateItineraryEntry(context.Background(), versionId).ItineraryEntryProperties(itineraryEntryProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItineraryAPI.CreateItineraryEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItineraryEntry`: ItineraryEntry
	fmt.Fprintf(os.Stdout, "Response from `ItineraryAPI.CreateItineraryEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Holiday Version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateItineraryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itineraryEntryProperties** | [**ItineraryEntryProperties**](ItineraryEntryProperties.md) |  | 

### Return type

[**ItineraryEntry**](ItineraryEntry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItineraryEntry

> DeleteItineraryEntry(ctx, versionId, entryId).Execute()

Delete an Itinerary Entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager-go"
)

func main() {
	versionId := "versionId_example" // string | Holiday Version ID
	entryId := "entryId_example" // string | Itinerary Entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItineraryAPI.DeleteItineraryEntry(context.Background(), versionId, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItineraryAPI.DeleteItineraryEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Holiday Version ID | 
**entryId** | **string** | Itinerary Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItineraryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItineraryEntries

> ListItineraryEntries200Response ListItineraryEntries(ctx, versionId).Page(page).Limit(limit).Order(order).Sort(sort).Trash(trash).Published(published).Execute()

List Itinerary Entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager-go"
)

func main() {
	versionId := "versionId_example" // string | Holiday Version ID
	page := int32(56) // int32 | Page number (optional) (default to 1)
	limit := int32(56) // int32 | Maximum number of results to return (optional) (default to 20)
	order := openapiclient.SortOrder("asc") // SortOrder | Sort order (optional)
	sort := "sort_example" // string | Sort field (optional) (default to "start_day")
	trash := int32(56) // int32 | View trash instead of active records if true (optional)
	published := int32(56) // int32 | Filter holidays by their publish state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItineraryAPI.ListItineraryEntries(context.Background(), versionId).Page(page).Limit(limit).Order(order).Sort(sort).Trash(trash).Published(published).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItineraryAPI.ListItineraryEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListItineraryEntries`: ListItineraryEntries200Response
	fmt.Fprintf(os.Stdout, "Response from `ItineraryAPI.ListItineraryEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Holiday Version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListItineraryEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Maximum number of results to return | [default to 20]
 **order** | [**SortOrder**](SortOrder.md) | Sort order | 
 **sort** | **string** | Sort field | [default to &quot;start_day&quot;]
 **trash** | **int32** | View trash instead of active records if true | 
 **published** | **int32** | Filter holidays by their publish state | 

### Return type

[**ListItineraryEntries200Response**](ListItineraryEntries200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItineraryEntry

> ItineraryEntry UpdateItineraryEntry(ctx, versionId, entryId).ItineraryEntryParams(itineraryEntryParams).Execute()

Update Itinerary Entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager-go"
)

func main() {
	versionId := "versionId_example" // string | Holiday Version ID
	entryId := "entryId_example" // string | Itinerary Entry ID
	itineraryEntryParams := *openapiclient.NewItineraryEntryParams() // ItineraryEntryParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItineraryAPI.UpdateItineraryEntry(context.Background(), versionId, entryId).ItineraryEntryParams(itineraryEntryParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItineraryAPI.UpdateItineraryEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItineraryEntry`: ItineraryEntry
	fmt.Fprintf(os.Stdout, "Response from `ItineraryAPI.UpdateItineraryEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | Holiday Version ID | 
**entryId** | **string** | Itinerary Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItineraryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **itineraryEntryParams** | [**ItineraryEntryParams**](ItineraryEntryParams.md) |  | 

### Return type

[**ItineraryEntry**](ItineraryEntry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

