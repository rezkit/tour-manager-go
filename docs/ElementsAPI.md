# \ElementsAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateElement**](ElementsAPI.md#CreateElement) | **Post** /holidays/versions/{version}/elements | Create Element
[**CreateOption**](ElementsAPI.md#CreateOption) | **Post** /holidays/elements/{element}/options | Create Option
[**DeleteElement**](ElementsAPI.md#DeleteElement) | **Delete** /holidays/versions/{version}/elements/{element} | Delete Element
[**DeleteOption**](ElementsAPI.md#DeleteOption) | **Delete** /holidays/elements/{element}/options/{option} | Delete Option
[**GetElement**](ElementsAPI.md#GetElement) | **Get** /holidays/versions/{version}/elements/{element} | Get Element
[**GetOption**](ElementsAPI.md#GetOption) | **Get** /holidays/elements/{element}/options/{option} | Get Option
[**ListElements**](ElementsAPI.md#ListElements) | **Get** /holidays/versions/{version}/elements | List Elements
[**RestoreAnOption**](ElementsAPI.md#RestoreAnOption) | **Put** /holidays/elements/{element}/options/{option}/restore | Restore an Option
[**UpdateElement**](ElementsAPI.md#UpdateElement) | **Patch** /holidays/versions/{version}/elements/{element} | Update Element
[**UpdateOption**](ElementsAPI.md#UpdateOption) | **Patch** /holidays/elements/{element}/options/{option} | Update Option



## CreateElement

> Element CreateElement(ctx, version).CreateElementRequest(createElementRequest).Execute()

Create Element



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
	version := "version_example" // string | Version ID
	createElementRequest := *openapiclient.NewCreateElementRequest("Name_example", "01gpkgcy6t0m84czh8gy4kjatk") // CreateElementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElementsAPI.CreateElement(context.Background(), version).CreateElementRequest(createElementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.CreateElement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateElement`: Element
	fmt.Fprintf(os.Stdout, "Response from `ElementsAPI.CreateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | Version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createElementRequest** | [**CreateElementRequest**](CreateElementRequest.md) |  | 

### Return type

[**Element**](Element.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOption

> ElementOptionProperties CreateOption(ctx, element).ElementOptionCreate(elementOptionCreate).Execute()

Create Option



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
	element := "element_example" // string | Element ID
	elementOptionCreate := *openapiclient.NewElementOptionCreate("01gpkgcy6t0m84czh8gy4kjatk", "Name_example", *openapiclient.NewOccupancy(int32(1), int32(2)), openapiclient.PriceUnit("unit")) // ElementOptionCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElementsAPI.CreateOption(context.Background(), element).ElementOptionCreate(elementOptionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.CreateOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOption`: ElementOptionProperties
	fmt.Fprintf(os.Stdout, "Response from `ElementsAPI.CreateOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**element** | **string** | Element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **elementOptionCreate** | [**ElementOptionCreate**](ElementOptionCreate.md) |  | 

### Return type

[**ElementOptionProperties**](ElementOptionProperties.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteElement

> DeleteElement(ctx, version, element).Execute()

Delete Element

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
	version := "version_example" // string | Version ID
	element := "element_example" // string | Element ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ElementsAPI.DeleteElement(context.Background(), version, element).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.DeleteElement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | Version ID | 
**element** | **string** | Element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteElementRequest struct via the builder pattern


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


## DeleteOption

> DeleteOption(ctx, element, option).Execute()

Delete Option



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
	element := "element_example" // string | Element ID
	option := "option_example" // string | Option ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ElementsAPI.DeleteOption(context.Background(), element, option).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.DeleteOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**element** | **string** | Element ID | 
**option** | **string** | Option ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOptionRequest struct via the builder pattern


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


## GetElement

> Element GetElement(ctx, version, element).Execute()

Get Element

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
	version := "version_example" // string | Version ID
	element := "element_example" // string | Element ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElementsAPI.GetElement(context.Background(), version, element).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.GetElement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetElement`: Element
	fmt.Fprintf(os.Stdout, "Response from `ElementsAPI.GetElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | Version ID | 
**element** | **string** | Element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Element**](Element.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOption

> ElementOptionProperties GetOption(ctx, element, option).Execute()

Get Option



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
	element := "element_example" // string | Element ID
	option := "option_example" // string | Option ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElementsAPI.GetOption(context.Background(), element, option).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.GetOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOption`: ElementOptionProperties
	fmt.Fprintf(os.Stdout, "Response from `ElementsAPI.GetOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**element** | **string** | Element ID | 
**option** | **string** | Option ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ElementOptionProperties**](ElementOptionProperties.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, applciation/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListElements

> ListElements200Response ListElements(ctx, version).Name(name).Page(page).Limit(limit).Order(order).Sort(sort).Trash(trash).Published(published).Execute()

List Elements



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
	version := "version_example" // string | Version ID
	name := "name_example" // string | Filter elements by name (contains) (optional)
	page := int32(56) // int32 | Page number (optional) (default to 1)
	limit := int32(56) // int32 | Maximum number of results to return (optional) (default to 20)
	order := openapiclient.SortOrder("asc") // SortOrder | Sort order (optional)
	sort := "sort_example" // string | Sort field (optional)
	trash := int32(56) // int32 | View trash instead of active records if true (optional)
	published := int32(56) // int32 | Filter holidays by their publish state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElementsAPI.ListElements(context.Background(), version).Name(name).Page(page).Limit(limit).Order(order).Sort(sort).Trash(trash).Published(published).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.ListElements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListElements`: ListElements200Response
	fmt.Fprintf(os.Stdout, "Response from `ElementsAPI.ListElements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | Version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListElementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter elements by name (contains) | 
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Maximum number of results to return | [default to 20]
 **order** | [**SortOrder**](SortOrder.md) | Sort order | 
 **sort** | **string** | Sort field | 
 **trash** | **int32** | View trash instead of active records if true | 
 **published** | **int32** | Filter holidays by their publish state | 

### Return type

[**ListElements200Response**](ListElements200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreAnOption

> ElementOptionProperties RestoreAnOption(ctx, element, option).Execute()

Restore an Option

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
	element := "element_example" // string | Element ID
	option := "option_example" // string | Option ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElementsAPI.RestoreAnOption(context.Background(), element, option).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.RestoreAnOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreAnOption`: ElementOptionProperties
	fmt.Fprintf(os.Stdout, "Response from `ElementsAPI.RestoreAnOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**element** | **string** | Element ID | 
**option** | **string** | Option ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreAnOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ElementOptionProperties**](ElementOptionProperties.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateElement

> Element UpdateElement(ctx, version, element).UpdateElementRequest(updateElementRequest).Execute()

Update Element

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
	version := "version_example" // string | Version ID
	element := "element_example" // string | Element ID
	updateElementRequest := *openapiclient.NewUpdateElementRequest() // UpdateElementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElementsAPI.UpdateElement(context.Background(), version, element).UpdateElementRequest(updateElementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.UpdateElement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateElement`: Element
	fmt.Fprintf(os.Stdout, "Response from `ElementsAPI.UpdateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** | Version ID | 
**element** | **string** | Element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateElementRequest** | [**UpdateElementRequest**](UpdateElementRequest.md) |  | 

### Return type

[**Element**](Element.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOption

> ElementOptionProperties UpdateOption(ctx, element, option).ElementOptionParams(elementOptionParams).Execute()

Update Option

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
	element := "element_example" // string | Element ID
	option := "option_example" // string | Option ID
	elementOptionParams := *openapiclient.NewElementOptionParams() // ElementOptionParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElementsAPI.UpdateOption(context.Background(), element, option).ElementOptionParams(elementOptionParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElementsAPI.UpdateOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOption`: ElementOptionProperties
	fmt.Fprintf(os.Stdout, "Response from `ElementsAPI.UpdateOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**element** | **string** | Element ID | 
**option** | **string** | Option ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **elementOptionParams** | [**ElementOptionParams**](ElementOptionParams.md) |  | 

### Return type

[**ElementOptionProperties**](ElementOptionProperties.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

