# \CategoriesAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachCategories**](CategoriesAPI.md#AttachCategories) | **Patch** /{type}/{item}/categories | Attach Categories
[**CreateCategory**](CategoriesAPI.md#CreateCategory) | **Post** /{type}/categories | Create Category
[**DeleteCategory**](CategoriesAPI.md#DeleteCategory) | **Delete** /{type}/categories/{category} | Delete Category
[**DetachCategories**](CategoriesAPI.md#DetachCategories) | **Delete** /{type}/{item}/categories | Detach Categories
[**ListAttachedCategories**](CategoriesAPI.md#ListAttachedCategories) | **Get** /{type}/{item}/categories | List Attached Categories
[**ListCategories**](CategoriesAPI.md#ListCategories) | **Get** /{type}/categories | List Categories
[**ReplaceCategories**](CategoriesAPI.md#ReplaceCategories) | **Put** /{type}/{item}/categories | Replace Categories
[**RestoreCategory**](CategoriesAPI.md#RestoreCategory) | **Put** /{type}/categories/{category}/restore | Restore Category
[**UpdateCategory**](CategoriesAPI.md#UpdateCategory) | **Patch** /{type}/categories/{category} | Update Category



## AttachCategories

> []Category AttachCategories(ctx, type_, item).ReplaceCategoriesRequest(replaceCategoriesRequest).Execute()

Attach Categories



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | Item Type
	item := "item_example" // string | Item ID
	replaceCategoriesRequest := *openapiclient.NewReplaceCategoriesRequest([]string{"01gpkgcy6t0m84czh8gy4kjatk"}) // ReplaceCategoriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.AttachCategories(context.Background(), type_, item).ReplaceCategoriesRequest(replaceCategoriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.AttachCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachCategories`: []Category
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.AttachCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) | Item Type | 
**item** | **string** | Item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceCategoriesRequest** | [**ReplaceCategoriesRequest**](ReplaceCategoriesRequest.md) |  | 

### Return type

[**[]Category**](Category.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCategory

> Category CreateCategory(ctx, type_).CreateCategoryRequest(createCategoryRequest).Execute()

Create Category



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | 
	createCategoryRequest := *openapiclient.NewCreateCategoryRequest("Name_example") // CreateCategoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.CreateCategory(context.Background(), type_).CreateCategoryRequest(createCategoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CreateCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCategory`: Category
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.CreateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCategoryRequest** | [**CreateCategoryRequest**](CreateCategoryRequest.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCategory

> DeleteCategory(ctx, type_, category).Execute()

Delete Category



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | 
	category := "category_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CategoriesAPI.DeleteCategory(context.Background(), type_, category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.DeleteCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) |  | 
**category** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


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


## DetachCategories

> DetachCategories(ctx, type_, item).Ids(ids).Execute()

Detach Categories



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | Item Type
	item := "item_example" // string | Item ID
	ids := []string{"01gpkgcy6t0m84czh8gy4kjatk"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CategoriesAPI.DetachCategories(context.Background(), type_, item).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.DetachCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) | Item Type | 
**item** | **string** | Item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ids** | **[]string** |  | 

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


## ListAttachedCategories

> InlineObject ListAttachedCategories(ctx, type_, item).Children(children).ParentId(parentId).Name(name).Search(search).Sort(sort).Order(order).Published(published).Searchable(searchable).Page(page).Limit(limit).Execute()

List Attached Categories

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
	type_ := openapiclient.EntityType("accommodation") // EntityType | Item Type
	item := "item_example" // string | Item ID
	children := true // bool | Include child categories on returned results (optional)
	parentId := "parentId_example" // string | Filter by parent ID (optional)
	name := "name_example" // string | Filter by name (optional)
	search := "search_example" // string | Free-text search (optional)
	sort := "sort_example" // string | Sort, defaults to category order (optional)
	order := openapiclient.SortOrder("asc") // SortOrder |  (optional)
	published := int32(56) // int32 | Filter Categories by their publish state (optional)
	searchable := int32(56) // int32 | Filter Categories by their searchability (optional)
	page := int32(56) // int32 | Page number (optional) (default to 1)
	limit := int32(56) // int32 | Maximum number of results to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.ListAttachedCategories(context.Background(), type_, item).Children(children).ParentId(parentId).Name(name).Search(search).Sort(sort).Order(order).Published(published).Searchable(searchable).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.ListAttachedCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAttachedCategories`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.ListAttachedCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) | Item Type | 
**item** | **string** | Item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAttachedCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **children** | **bool** | Include child categories on returned results | 
 **parentId** | **string** | Filter by parent ID | 
 **name** | **string** | Filter by name | 
 **search** | **string** | Free-text search | 
 **sort** | **string** | Sort, defaults to category order | 
 **order** | [**SortOrder**](SortOrder.md) |  | 
 **published** | **int32** | Filter Categories by their publish state | 
 **searchable** | **int32** | Filter Categories by their searchability | 
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Maximum number of results to return | [default to 20]

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategories

> InlineObject ListCategories(ctx, type_).Children(children).ParentId(parentId).Name(name).Search(search).Sort(sort).Order(order).Published(published).Searchable(searchable).Page(page).Limit(limit).Execute()

List Categories



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | 
	children := true // bool | Include child categories on returned results (optional)
	parentId := "parentId_example" // string | Filter by parent ID (optional)
	name := "name_example" // string | Filter by name (optional)
	search := "search_example" // string | Free-text search (optional)
	sort := "sort_example" // string | Sort, defaults to category order (optional)
	order := openapiclient.SortOrder("asc") // SortOrder |  (optional)
	published := int32(56) // int32 | Filter Categories by their publish state (optional)
	searchable := int32(56) // int32 | Filter Categories by their searchability (optional)
	page := int32(56) // int32 | Page number (optional) (default to 1)
	limit := int32(56) // int32 | Maximum number of results to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.ListCategories(context.Background(), type_).Children(children).ParentId(parentId).Name(name).Search(search).Sort(sort).Order(order).Published(published).Searchable(searchable).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.ListCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCategories`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.ListCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **children** | **bool** | Include child categories on returned results | 
 **parentId** | **string** | Filter by parent ID | 
 **name** | **string** | Filter by name | 
 **search** | **string** | Free-text search | 
 **sort** | **string** | Sort, defaults to category order | 
 **order** | [**SortOrder**](SortOrder.md) |  | 
 **published** | **int32** | Filter Categories by their publish state | 
 **searchable** | **int32** | Filter Categories by their searchability | 
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Maximum number of results to return | [default to 20]

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCategories

> []Category ReplaceCategories(ctx, type_, item).ReplaceCategoriesRequest(replaceCategoriesRequest).Execute()

Replace Categories



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | Item Type
	item := "item_example" // string | Item ID
	replaceCategoriesRequest := *openapiclient.NewReplaceCategoriesRequest([]string{"01gpkgcy6t0m84czh8gy4kjatk"}) // ReplaceCategoriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.ReplaceCategories(context.Background(), type_, item).ReplaceCategoriesRequest(replaceCategoriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.ReplaceCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCategories`: []Category
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.ReplaceCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) | Item Type | 
**item** | **string** | Item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceCategoriesRequest** | [**ReplaceCategoriesRequest**](ReplaceCategoriesRequest.md) |  | 

### Return type

[**[]Category**](Category.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreCategory

> RestoreCategory(ctx, type_, category).Execute()

Restore Category

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
	type_ := openapiclient.EntityType("accommodation") // EntityType | 
	category := "category_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CategoriesAPI.RestoreCategory(context.Background(), type_, category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.RestoreCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) |  | 
**category** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreCategoryRequest struct via the builder pattern


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


## UpdateCategory

> Category UpdateCategory(ctx, type_, category).UpdateCategoryRequest(updateCategoryRequest).Execute()

Update Category



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | 
	category := "category_example" // string | 
	updateCategoryRequest := *openapiclient.NewUpdateCategoryRequest() // UpdateCategoryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.UpdateCategory(context.Background(), type_, category).UpdateCategoryRequest(updateCategoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.UpdateCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCategory`: Category
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.UpdateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) |  | 
**category** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCategoryRequest** | [**UpdateCategoryRequest**](UpdateCategoryRequest.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

