# \HolidayVersionsAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHolidayVersion**](HolidayVersionsAPI.md#CreateHolidayVersion) | **Post** /holidays/{holiday_id}/versions | Create Holiday Version
[**DeleteHolidayVersion**](HolidayVersionsAPI.md#DeleteHolidayVersion) | **Delete** /holidays/{holiday_id}/versions/{version_id} | Delete Holiday Version
[**GetHolidayVersion**](HolidayVersionsAPI.md#GetHolidayVersion) | **Get** /holidays/{holiday_id}/versions/{version_id} | Get Holiday version
[**ListHolidayVersions**](HolidayVersionsAPI.md#ListHolidayVersions) | **Get** /holidays/{holiday_id}/versions | List Holiday Versions
[**RestoreHolidayVersion**](HolidayVersionsAPI.md#RestoreHolidayVersion) | **Put** /holidays/{holiday_id}/versions/{version_id}/restore | Restore Holiday Version
[**UpdateHolidayVersion**](HolidayVersionsAPI.md#UpdateHolidayVersion) | **Patch** /holidays/{holiday_id}/versions/{version_id} | Update Holiday Version



## CreateHolidayVersion

> HolidayVersion CreateHolidayVersion(ctx, holidayId).CreateHolidayRequest(createHolidayRequest).Execute()

Create Holiday Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager"
)

func main() {
	holidayId := "holidayId_example" // string | Holiday ID
	createHolidayRequest := *openapiclient.NewCreateHolidayRequest("Name_example", "H001") // CreateHolidayRequest | New Holiday Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidayVersionsAPI.CreateHolidayVersion(context.Background(), holidayId).CreateHolidayRequest(createHolidayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidayVersionsAPI.CreateHolidayVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHolidayVersion`: HolidayVersion
	fmt.Fprintf(os.Stdout, "Response from `HolidayVersionsAPI.CreateHolidayVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holidayId** | **string** | Holiday ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHolidayVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createHolidayRequest** | [**CreateHolidayRequest**](CreateHolidayRequest.md) | New Holiday Payload | 

### Return type

[**HolidayVersion**](HolidayVersion.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHolidayVersion

> DeleteHolidayVersion(ctx, holidayId, versionId).Execute()

Delete Holiday Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager"
)

func main() {
	holidayId := "holidayId_example" // string | Holiday ID
	versionId := "versionId_example" // string | Holiday Version ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HolidayVersionsAPI.DeleteHolidayVersion(context.Background(), holidayId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidayVersionsAPI.DeleteHolidayVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holidayId** | **string** | Holiday ID | 
**versionId** | **string** | Holiday Version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHolidayVersionRequest struct via the builder pattern


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


## GetHolidayVersion

> HolidayVersion GetHolidayVersion(ctx, holidayId, versionId).Execute()

Get Holiday version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager"
)

func main() {
	holidayId := "holidayId_example" // string | Holiday ID
	versionId := "versionId_example" // string | Holiday Version ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidayVersionsAPI.GetHolidayVersion(context.Background(), holidayId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidayVersionsAPI.GetHolidayVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHolidayVersion`: HolidayVersion
	fmt.Fprintf(os.Stdout, "Response from `HolidayVersionsAPI.GetHolidayVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holidayId** | **string** | Holiday ID | 
**versionId** | **string** | Holiday Version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHolidayVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HolidayVersion**](HolidayVersion.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHolidayVersions

> ListHolidayVersions200Response ListHolidayVersions(ctx, holidayId).Search(search).Name(name).Code(code).Page(page).Limit(limit).Order(order).Sort(sort).Trash(trash).Published(published).Execute()

List Holiday Versions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager"
)

func main() {
	holidayId := "holidayId_example" // string | Holiday ID
	search := "search_example" // string | (UI) Holiday search query (optional)
	name := "name_example" // string | Filter holidays by name (contains) (optional)
	code := "code_example" // string | Filter holidays by code (prefix) (optional)
	page := int32(56) // int32 | Page number (optional) (default to 1)
	limit := int32(56) // int32 | Maximum number of results to return (optional) (default to 20)
	order := openapiclient.SortOrder("asc") // SortOrder | Sort order (optional)
	sort := "sort_example" // string | Sort field (optional)
	trash := int32(56) // int32 | View trash instead of active records if true (optional)
	published := int32(56) // int32 | Filter holidays by their publish state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidayVersionsAPI.ListHolidayVersions(context.Background(), holidayId).Search(search).Name(name).Code(code).Page(page).Limit(limit).Order(order).Sort(sort).Trash(trash).Published(published).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidayVersionsAPI.ListHolidayVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHolidayVersions`: ListHolidayVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `HolidayVersionsAPI.ListHolidayVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holidayId** | **string** | Holiday ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHolidayVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | (UI) Holiday search query | 
 **name** | **string** | Filter holidays by name (contains) | 
 **code** | **string** | Filter holidays by code (prefix) | 
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Maximum number of results to return | [default to 20]
 **order** | [**SortOrder**](SortOrder.md) | Sort order | 
 **sort** | **string** | Sort field | 
 **trash** | **int32** | View trash instead of active records if true | 
 **published** | **int32** | Filter holidays by their publish state | 

### Return type

[**ListHolidayVersions200Response**](ListHolidayVersions200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreHolidayVersion

> HolidayVersion RestoreHolidayVersion(ctx, holidayId, versionId).Execute()

Restore Holiday Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager"
)

func main() {
	holidayId := "holidayId_example" // string | Holiday ID
	versionId := "versionId_example" // string | Holiday Version ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidayVersionsAPI.RestoreHolidayVersion(context.Background(), holidayId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidayVersionsAPI.RestoreHolidayVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreHolidayVersion`: HolidayVersion
	fmt.Fprintf(os.Stdout, "Response from `HolidayVersionsAPI.RestoreHolidayVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holidayId** | **string** | Holiday ID | 
**versionId** | **string** | Holiday Version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreHolidayVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HolidayVersion**](HolidayVersion.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHolidayVersion

> HolidayVersion UpdateHolidayVersion(ctx, holidayId, versionId).UpdateHolidayRequest(updateHolidayRequest).Execute()

Update Holiday Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rezkit/tour-manager"
)

func main() {
	holidayId := "holidayId_example" // string | Holiday ID
	versionId := "versionId_example" // string | Holiday Version ID
	updateHolidayRequest := *openapiclient.NewUpdateHolidayRequest() // UpdateHolidayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidayVersionsAPI.UpdateHolidayVersion(context.Background(), holidayId, versionId).UpdateHolidayRequest(updateHolidayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidayVersionsAPI.UpdateHolidayVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHolidayVersion`: HolidayVersion
	fmt.Fprintf(os.Stdout, "Response from `HolidayVersionsAPI.UpdateHolidayVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holidayId** | **string** | Holiday ID | 
**versionId** | **string** | Holiday Version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHolidayVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateHolidayRequest** | [**UpdateHolidayRequest**](UpdateHolidayRequest.md) |  | 

### Return type

[**HolidayVersion**](HolidayVersion.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

