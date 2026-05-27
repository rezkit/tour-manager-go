# \DeparturesAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeparture**](DeparturesAPI.md#CreateDeparture) | **Post** /holidays/departures | Create Departure
[**DeleteDeparture**](DeparturesAPI.md#DeleteDeparture) | **Delete** /holidays/departures/{departure} | Delete Departure
[**GetDeparture**](DeparturesAPI.md#GetDeparture) | **Get** /holidays/departures/{departure} | Get Departure
[**ListDepartures**](DeparturesAPI.md#ListDepartures) | **Get** /holidays/departures | List Departures
[**UpdateDeparture**](DeparturesAPI.md#UpdateDeparture) | **Patch** /holidays/departures/{departure} | Update Departure



## CreateDeparture

> Departure CreateDeparture(ctx).DepartureProperties(departureProperties).Execute()

Create Departure



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/rezkit/tour-manager-go"
)

func main() {
	departureProperties := *openapiclient.NewDepartureProperties(time.Now(), time.Now(), *openapiclient.NewInventory("Type_example")) // DepartureProperties | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeparturesAPI.CreateDeparture(context.Background()).DepartureProperties(departureProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeparturesAPI.CreateDeparture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeparture`: Departure
	fmt.Fprintf(os.Stdout, "Response from `DeparturesAPI.CreateDeparture`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDepartureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **departureProperties** | [**DepartureProperties**](DepartureProperties.md) |  | 

### Return type

[**Departure**](Departure.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeparture

> DeleteDeparture(ctx, departure).Execute()

Delete Departure



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
	departure := "departure_example" // string | Departure ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeparturesAPI.DeleteDeparture(context.Background(), departure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeparturesAPI.DeleteDeparture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departure** | **string** | Departure ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDepartureRequest struct via the builder pattern


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


## GetDeparture

> Departure GetDeparture(ctx, departure).Execute()

Get Departure



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
	departure := "departure_example" // string | Departure ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeparturesAPI.GetDeparture(context.Background(), departure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeparturesAPI.GetDeparture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeparture`: Departure
	fmt.Fprintf(os.Stdout, "Response from `DeparturesAPI.GetDeparture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departure** | **string** | Departure ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDepartureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Departure**](Departure.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDepartures

> ListDepartures200Response ListDepartures(ctx).Version(version).Holiday(holiday).Before(before).After(after).Published(published).Execute()

List Departures



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/rezkit/tour-manager-go"
)

func main() {
	version := "version_example" // string | Filter by Version (optional)
	holiday := "holiday_example" // string | Filter by holiday (optional)
	before := time.Now() // time.Time | Return only departures which *may return before* the given date. (optional)
	after := time.Now() // time.Time | Return only departures which *may depart after* the given date. (optional)
	published := int32(56) // int32 | Return only published/non-published departures (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeparturesAPI.ListDepartures(context.Background()).Version(version).Holiday(holiday).Before(before).After(after).Published(published).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeparturesAPI.ListDepartures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDepartures`: ListDepartures200Response
	fmt.Fprintf(os.Stdout, "Response from `DeparturesAPI.ListDepartures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeparturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | Filter by Version | 
 **holiday** | **string** | Filter by holiday | 
 **before** | **time.Time** | Return only departures which *may return before* the given date. | 
 **after** | **time.Time** | Return only departures which *may depart after* the given date. | 
 **published** | **int32** | Return only published/non-published departures | 

### Return type

[**ListDepartures200Response**](ListDepartures200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeparture

> Departure UpdateDeparture(ctx, departure).DepartureParams(departureParams).Execute()

Update Departure



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
	departure := "departure_example" // string | Departure ID
	departureParams := *openapiclient.NewDepartureParams() // DepartureParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeparturesAPI.UpdateDeparture(context.Background(), departure).DepartureParams(departureParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeparturesAPI.UpdateDeparture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeparture`: Departure
	fmt.Fprintf(os.Stdout, "Response from `DeparturesAPI.UpdateDeparture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departure** | **string** | Departure ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDepartureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **departureParams** | [**DepartureParams**](DepartureParams.md) |  | 

### Return type

[**Departure**](Departure.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

