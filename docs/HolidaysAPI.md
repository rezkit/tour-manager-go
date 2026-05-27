# \HolidaysAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRelatedHoliday**](HolidaysAPI.md#AddRelatedHoliday) | **Post** /holidays/{id}/relations | Add Related Holiday
[**CopyHoliday**](HolidaysAPI.md#CopyHoliday) | **Put** /holidays/{id}/copy | Copy Holiday
[**CreateHoliday**](HolidaysAPI.md#CreateHoliday) | **Post** /holidays | Create Holiday
[**DeleteHoliday**](HolidaysAPI.md#DeleteHoliday) | **Delete** /holidays/{id} | Delete holiday
[**DeleteRelatedHoliday**](HolidaysAPI.md#DeleteRelatedHoliday) | **Delete** /holidays/{id}/relations/{relation} | Delete Relation
[**GetHoliday**](HolidaysAPI.md#GetHoliday) | **Get** /holidays/{id} | Get Holiday
[**GetRelatedHoliday**](HolidaysAPI.md#GetRelatedHoliday) | **Get** /holidays/{id}/relations/{relation} | Get Relation
[**ListHolidays**](HolidaysAPI.md#ListHolidays) | **Get** /holidays | List Holidays
[**ListRelatedHolidays**](HolidaysAPI.md#ListRelatedHolidays) | **Get** /holidays/{id}/relations | List Related Holidays
[**ReindexHolidays**](HolidaysAPI.md#ReindexHolidays) | **Post** /holidays/reindex | Trigger Search Reindex
[**RestoreAHoliday**](HolidaysAPI.md#RestoreAHoliday) | **Put** /holidays/{holiday_id}/restore | Restore a Holiday
[**SearchHolidays**](HolidaysAPI.md#SearchHolidays) | **Get** /holidays/search | Holiday Search
[**SuggestHolidays**](HolidaysAPI.md#SuggestHolidays) | **Get** /holidays/suggest | Holiday Suggestions
[**UpdateHoliday**](HolidaysAPI.md#UpdateHoliday) | **Patch** /holidays/{id} | Update Holiday
[**UpdateRelatedHoliday**](HolidaysAPI.md#UpdateRelatedHoliday) | **Patch** /holidays/{id}/relations/{relation} | Update Relation



## AddRelatedHoliday

> HolidayEdge AddRelatedHoliday(ctx, id).AddRelatedHolidayRequest(addRelatedHolidayRequest).Execute()

Add Related Holiday



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
	id := "id_example" // string | 
	addRelatedHolidayRequest := *openapiclient.NewAddRelatedHolidayRequest("01gpkgcy6t0m84czh8gy4kjatk", "01gpkgcy6t0m84czh8gy4kjatk", false) // AddRelatedHolidayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.AddRelatedHoliday(context.Background(), id).AddRelatedHolidayRequest(addRelatedHolidayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.AddRelatedHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRelatedHoliday`: HolidayEdge
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.AddRelatedHoliday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRelatedHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addRelatedHolidayRequest** | [**AddRelatedHolidayRequest**](AddRelatedHolidayRequest.md) |  | 

### Return type

[**HolidayEdge**](HolidayEdge.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyHoliday

> Holiday CopyHoliday(ctx, id).Execute()

Copy Holiday



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
	id := "id_example" // string | Holiday ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.CopyHoliday(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.CopyHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyHoliday`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.CopyHoliday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Holiday ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Holiday**](Holiday.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHoliday

> Holiday CreateHoliday(ctx).CreateHolidayRequest(createHolidayRequest).Execute()

Create Holiday



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
	createHolidayRequest := *openapiclient.NewCreateHolidayRequest("Name_example", "H001") // CreateHolidayRequest | New Holiday Payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.CreateHoliday(context.Background()).CreateHolidayRequest(createHolidayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.CreateHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHoliday`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.CreateHoliday`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createHolidayRequest** | [**CreateHolidayRequest**](CreateHolidayRequest.md) | New Holiday Payload | 

### Return type

[**Holiday**](Holiday.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHoliday

> DeleteHoliday(ctx, id).Execute()

Delete holiday



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
	id := "id_example" // string | Holiday ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HolidaysAPI.DeleteHoliday(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.DeleteHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Holiday ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHolidayRequest struct via the builder pattern


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


## DeleteRelatedHoliday

> DeleteRelatedHoliday(ctx, id, relation).Execute()

Delete Relation

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
	id := "id_example" // string | 
	relation := "relation_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HolidaysAPI.DeleteRelatedHoliday(context.Background(), id, relation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.DeleteRelatedHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**relation** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRelatedHolidayRequest struct via the builder pattern


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


## GetHoliday

> Holiday GetHoliday(ctx, id).Execute()

Get Holiday



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
	id := "id_example" // string | Holiday ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.GetHoliday(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.GetHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHoliday`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.GetHoliday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Holiday ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Holiday**](Holiday.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedHoliday

> HolidayEdge GetRelatedHoliday(ctx, id, relation).Execute()

Get Relation

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
	id := "id_example" // string | 
	relation := "relation_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.GetRelatedHoliday(context.Background(), id, relation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.GetRelatedHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRelatedHoliday`: HolidayEdge
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.GetRelatedHoliday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**relation** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelatedHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HolidayEdge**](HolidayEdge.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHolidays

> ListHolidays200Response ListHolidays(ctx).Search(search).Name(name).Code(code).Page(page).Limit(limit).Order(order).Sort(sort).Trash(trash).Published(published).Execute()

List Holidays



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
	resp, r, err := apiClient.HolidaysAPI.ListHolidays(context.Background()).Search(search).Name(name).Code(code).Page(page).Limit(limit).Order(order).Sort(sort).Trash(trash).Published(published).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.ListHolidays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHolidays`: ListHolidays200Response
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.ListHolidays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHolidaysRequest struct via the builder pattern


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

[**ListHolidays200Response**](ListHolidays200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRelatedHolidays

> []HolidayEdge ListRelatedHolidays(ctx, id).Execute()

List Related Holidays



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.ListRelatedHolidays(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.ListRelatedHolidays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRelatedHolidays`: []HolidayEdge
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.ListRelatedHolidays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRelatedHolidaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HolidayEdge**](HolidayEdge.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReindexHolidays

> ReindexHolidays200Response ReindexHolidays(ctx).Execute()

Trigger Search Reindex



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.ReindexHolidays(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.ReindexHolidays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReindexHolidays`: ReindexHolidays200Response
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.ReindexHolidays`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReindexHolidaysRequest struct via the builder pattern


### Return type

[**ReindexHolidays200Response**](ReindexHolidays200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreAHoliday

> Holiday RestoreAHoliday(ctx, holidayId).Execute()

Restore a Holiday



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
	holidayId := "holidayId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.RestoreAHoliday(context.Background(), holidayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.RestoreAHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreAHoliday`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.RestoreAHoliday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**holidayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreAHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Holiday**](Holiday.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchHolidays

> SearchHolidays(ctx).K(k).Ccy(ccy).Q(q).O(o).I(i).J(j).Pf(pf).Pt(pt).Df(df).Dt(dt).L(l).Fh(fh).Fd(fd).Pi(pi).C0(c0).C1(c1).C2(c2).C3(c3).C4(c4).C5(c5).C6(c6).C7(c7).C8(c8).C9(c9).Execute()

Holiday Search

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
	k := "rezkit" // string | Public search key
	ccy := "GBP" // string | Active search currency
	q := "walking holidays" // string | Search query (optional)
	o := int32(0) // int32 | Pagination offset (optional)
	i := int32(25) // int32 | Pagination page size (optional)
	j := "d" // string | Search join/filtering mode (optional)
	pf := float32(123.45) // float32 | Minimum Price (inclusive) (optional)
	pt := float32(5000) // float32 | Maximum Price (inclusive) (optional)
	df := "now+30d" // string | Date range lower bound (inclusive) Can be either a specific timestamp value or a relative date calculation e.g. `now+7d`  (optional)
	dt := "now+1y" // string | Date range upper bound (inclusive) Can be either a specific timestamp value or a relative date calculation e.g. `now+7d`  (optional)
	l := map[string][]string{"key": []string{"Inner_example"}} // map[string][]string | Filter by location (optional)
	fh := map[string]SearchHolidaysFhParameterValue{"key": openapiclient.searchHolidays_fh_parameter_value{Bool: new(bool)}} // map[string]SearchHolidaysFhParameterValue | Filter based on custom field properties, using an exact match only.  (optional)
	fd := map[string]SearchHolidaysFhParameterValue{"key": openapiclient.searchHolidays_fh_parameter_value{Bool: new(bool)}} // map[string]SearchHolidaysFhParameterValue | Filter departures based on custom field properties, using an exact match only (optional)
	pi := int32(56) // int32 | Price histogram interval (optional)
	c0 := []string{"Inner_example"} // []string | Filter by Categories (Level 0) (optional)
	c1 := []string{"Inner_example"} // []string | Filter by Categories (Level 1) (optional)
	c2 := []string{"Inner_example"} // []string | Filter by Categories (Level 2) (optional)
	c3 := []string{"Inner_example"} // []string | Filter by Categories (Level 3) (optional)
	c4 := []string{"Inner_example"} // []string | Filter by Categories (Level 4) (optional)
	c5 := []string{"Inner_example"} // []string | Filter by Categories (Level 5) (optional)
	c6 := []string{"Inner_example"} // []string | Filter by Categories (Level 6) (optional)
	c7 := []string{"Inner_example"} // []string | Filter by Categories (Level 7) (optional)
	c8 := []string{"Inner_example"} // []string | Filter by Categories (Level 8) (optional)
	c9 := []string{"Inner_example"} // []string | Filter by Categories (Level 9) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HolidaysAPI.SearchHolidays(context.Background()).K(k).Ccy(ccy).Q(q).O(o).I(i).J(j).Pf(pf).Pt(pt).Df(df).Dt(dt).L(l).Fh(fh).Fd(fd).Pi(pi).C0(c0).C1(c1).C2(c2).C3(c3).C4(c4).C5(c5).C6(c6).C7(c7).C8(c8).C9(c9).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.SearchHolidays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchHolidaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **k** | **string** | Public search key | 
 **ccy** | **string** | Active search currency | 
 **q** | **string** | Search query | 
 **o** | **int32** | Pagination offset | 
 **i** | **int32** | Pagination page size | 
 **j** | **string** | Search join/filtering mode | 
 **pf** | **float32** | Minimum Price (inclusive) | 
 **pt** | **float32** | Maximum Price (inclusive) | 
 **df** | **string** | Date range lower bound (inclusive) Can be either a specific timestamp value or a relative date calculation e.g. &#x60;now+7d&#x60;  | 
 **dt** | **string** | Date range upper bound (inclusive) Can be either a specific timestamp value or a relative date calculation e.g. &#x60;now+7d&#x60;  | 
 **l** | **map[string][]string** | Filter by location | 
 **fh** | [**map[string]SearchHolidaysFhParameterValue**](SearchHolidaysFhParameterValue.md) | Filter based on custom field properties, using an exact match only.  | 
 **fd** | [**map[string]SearchHolidaysFhParameterValue**](SearchHolidaysFhParameterValue.md) | Filter departures based on custom field properties, using an exact match only | 
 **pi** | **int32** | Price histogram interval | 
 **c0** | **[]string** | Filter by Categories (Level 0) | 
 **c1** | **[]string** | Filter by Categories (Level 1) | 
 **c2** | **[]string** | Filter by Categories (Level 2) | 
 **c3** | **[]string** | Filter by Categories (Level 3) | 
 **c4** | **[]string** | Filter by Categories (Level 4) | 
 **c5** | **[]string** | Filter by Categories (Level 5) | 
 **c6** | **[]string** | Filter by Categories (Level 6) | 
 **c7** | **[]string** | Filter by Categories (Level 7) | 
 **c8** | **[]string** | Filter by Categories (Level 8) | 
 **c9** | **[]string** | Filter by Categories (Level 9) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuggestHolidays

> []string SuggestHolidays(ctx).Q(q).Execute()

Holiday Suggestions



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
	q := "q_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.SuggestHolidays(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.SuggestHolidays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuggestHolidays`: []string
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.SuggestHolidays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuggestHolidaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHoliday

> Holiday UpdateHoliday(ctx, id).UpdateHolidayRequest(updateHolidayRequest).Execute()

Update Holiday



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
	id := "id_example" // string | Holiday ID
	updateHolidayRequest := *openapiclient.NewUpdateHolidayRequest() // UpdateHolidayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.UpdateHoliday(context.Background(), id).UpdateHolidayRequest(updateHolidayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.UpdateHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHoliday`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.UpdateHoliday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Holiday ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHolidayRequest** | [**UpdateHolidayRequest**](UpdateHolidayRequest.md) |  | 

### Return type

[**Holiday**](Holiday.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRelatedHoliday

> HolidayEdge UpdateRelatedHoliday(ctx, id, relation).UpdateRelatedHolidayRequest(updateRelatedHolidayRequest).Execute()

Update Relation

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
	id := "id_example" // string | 
	relation := "relation_example" // string | 
	updateRelatedHolidayRequest := *openapiclient.NewUpdateRelatedHolidayRequest() // UpdateRelatedHolidayRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.UpdateRelatedHoliday(context.Background(), id, relation).UpdateRelatedHolidayRequest(updateRelatedHolidayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.UpdateRelatedHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRelatedHoliday`: HolidayEdge
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.UpdateRelatedHoliday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**relation** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRelatedHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRelatedHolidayRequest** | [**UpdateRelatedHolidayRequest**](UpdateRelatedHolidayRequest.md) |  | 

### Return type

[**HolidayEdge**](HolidayEdge.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

