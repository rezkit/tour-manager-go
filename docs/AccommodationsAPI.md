# \AccommodationsAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccommodation**](AccommodationsAPI.md#CreateAccommodation) | **Post** /accommodations | Create Accommodation
[**CreateRoomType**](AccommodationsAPI.md#CreateRoomType) | **Post** /accommodations/{id}/roomTypes | Create Room Type
[**CreateRoomTypePrice**](AccommodationsAPI.md#CreateRoomTypePrice) | **Post** /accommodations/{id}/roomTypes/{room_type_id}/prices | Create Room Type Price
[**DeleteAccommodation**](AccommodationsAPI.md#DeleteAccommodation) | **Delete** /accommodations/{id} | Delete Accommodation
[**DeleteRoomType**](AccommodationsAPI.md#DeleteRoomType) | **Delete** /accommodations/{id}/roomTypes/{room_type_id} | Delete Room Type
[**DeleteRoomTypePrice**](AccommodationsAPI.md#DeleteRoomTypePrice) | **Delete** /accommodations/{id}/roomTypes/{room_type_id}/prices/{price_id} | Delete Room Type Price
[**GetAccommodation**](AccommodationsAPI.md#GetAccommodation) | **Get** /accommodations/{id} | Get Accommodation
[**GetRoomType**](AccommodationsAPI.md#GetRoomType) | **Get** /accommodations/{id}/roomTypes/{room_type_id} | Get Room Type
[**GetRoomTypePrice**](AccommodationsAPI.md#GetRoomTypePrice) | **Get** /accommodations/{id}/roomTypes/{room_type_id}/prices/{price_id} | Get Room Type Price
[**ListAccommodations**](AccommodationsAPI.md#ListAccommodations) | **Get** /accommodations | List Accommodations
[**ListRoomTypePrices**](AccommodationsAPI.md#ListRoomTypePrices) | **Get** /accommodations/{id}/roomTypes/{room_type_id}/prices | List Room Type Prices
[**ListRoomTypes**](AccommodationsAPI.md#ListRoomTypes) | **Get** /accommodations/{id}/roomTypes | List Room Types
[**RestoreAccommodation**](AccommodationsAPI.md#RestoreAccommodation) | **Put** /accommodations/{id}/restore | Restore Accommodation
[**RestoreRoomType**](AccommodationsAPI.md#RestoreRoomType) | **Put** /accommodations/{id}/roomTypes/{room_type_id}/restore | Restore Room Type
[**RestoreRoomTypePrice**](AccommodationsAPI.md#RestoreRoomTypePrice) | **Put** /accommodations/{id}/roomTypes/{room_type_id}/prices/{price_id}/restore | Restore Room Type Price
[**UpdateAccommodation**](AccommodationsAPI.md#UpdateAccommodation) | **Patch** /accommodations/{id} | Update Accommodation
[**UpdateRoomType**](AccommodationsAPI.md#UpdateRoomType) | **Patch** /accommodations/{id}/roomTypes/{room_type_id} | Update Room Type
[**UpdateRoomTypePrice**](AccommodationsAPI.md#UpdateRoomTypePrice) | **Patch** /accommodations/{id}/roomTypes/{room_type_id}/prices/{price_id} | Update Room Type Price



## CreateAccommodation

> CreateAccommodation(ctx).Accommodation(accommodation).Execute()

Create Accommodation

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
	accommodation := *openapiclient.NewAccommodation() // Accommodation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.CreateAccommodation(context.Background()).Accommodation(accommodation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.CreateAccommodation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccommodationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accommodation** | [**Accommodation**](Accommodation.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoomType

> CreateRoomType(ctx, id).Execute()

Create Room Type

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
	r, err := apiClient.AccommodationsAPI.CreateRoomType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.CreateRoomType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomTypeRequest struct via the builder pattern


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


## CreateRoomTypePrice

> CreateRoomTypePrice(ctx, id, roomTypeId).Execute()

Create Room Type Price

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
	roomTypeId := "roomTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.CreateRoomTypePrice(context.Background(), id, roomTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.CreateRoomTypePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomTypePriceRequest struct via the builder pattern


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


## DeleteAccommodation

> DeleteAccommodation(ctx, id).Execute()

Delete Accommodation

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
	r, err := apiClient.AccommodationsAPI.DeleteAccommodation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.DeleteAccommodation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccommodationRequest struct via the builder pattern


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


## DeleteRoomType

> DeleteRoomType(ctx, id, roomTypeId).Execute()

Delete Room Type

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
	roomTypeId := "roomTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.DeleteRoomType(context.Background(), id, roomTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.DeleteRoomType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomTypeRequest struct via the builder pattern


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


## DeleteRoomTypePrice

> DeleteRoomTypePrice(ctx, id, roomTypeId, priceId).Execute()

Delete Room Type Price

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
	roomTypeId := "roomTypeId_example" // string | 
	priceId := "priceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.DeleteRoomTypePrice(context.Background(), id, roomTypeId, priceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.DeleteRoomTypePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomTypePriceRequest struct via the builder pattern


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


## GetAccommodation

> GetAccommodation(ctx, id).Execute()

Get Accommodation

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
	r, err := apiClient.AccommodationsAPI.GetAccommodation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.GetAccommodation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccommodationRequest struct via the builder pattern


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


## GetRoomType

> GetRoomType(ctx, id, roomTypeId).Execute()

Get Room Type

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
	roomTypeId := "roomTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.GetRoomType(context.Background(), id, roomTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.GetRoomType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomTypeRequest struct via the builder pattern


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


## GetRoomTypePrice

> GetRoomTypePrice(ctx, id, roomTypeId, priceId).Execute()

Get Room Type Price

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
	roomTypeId := "roomTypeId_example" // string | 
	priceId := "priceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.GetRoomTypePrice(context.Background(), id, roomTypeId, priceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.GetRoomTypePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomTypePriceRequest struct via the builder pattern


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


## ListAccommodations

> ListAccommodations200Response ListAccommodations(ctx).Execute()

List Accommodations

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
	resp, r, err := apiClient.AccommodationsAPI.ListAccommodations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.ListAccommodations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccommodations`: ListAccommodations200Response
	fmt.Fprintf(os.Stdout, "Response from `AccommodationsAPI.ListAccommodations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAccommodationsRequest struct via the builder pattern


### Return type

[**ListAccommodations200Response**](ListAccommodations200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomTypePrices

> ListRoomTypePrices(ctx, id, roomTypeId).Execute()

List Room Type Prices

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
	roomTypeId := "roomTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.ListRoomTypePrices(context.Background(), id, roomTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.ListRoomTypePrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoomTypePricesRequest struct via the builder pattern


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


## ListRoomTypes

> ListRoomTypes(ctx, id).Execute()

List Room Types

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
	r, err := apiClient.AccommodationsAPI.ListRoomTypes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.ListRoomTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoomTypesRequest struct via the builder pattern


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


## RestoreAccommodation

> RestoreAccommodation(ctx, id).Execute()

Restore Accommodation

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
	r, err := apiClient.AccommodationsAPI.RestoreAccommodation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.RestoreAccommodation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreAccommodationRequest struct via the builder pattern


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


## RestoreRoomType

> RestoreRoomType(ctx, id, roomTypeId).Execute()

Restore Room Type

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
	roomTypeId := "roomTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.RestoreRoomType(context.Background(), id, roomTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.RestoreRoomType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreRoomTypeRequest struct via the builder pattern


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


## RestoreRoomTypePrice

> RestoreRoomTypePrice(ctx, id, roomTypeId, priceId).Execute()

Restore Room Type Price

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
	roomTypeId := "roomTypeId_example" // string | 
	priceId := "priceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.RestoreRoomTypePrice(context.Background(), id, roomTypeId, priceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.RestoreRoomTypePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreRoomTypePriceRequest struct via the builder pattern


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


## UpdateAccommodation

> UpdateAccommodation(ctx, id).Execute()

Update Accommodation

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
	r, err := apiClient.AccommodationsAPI.UpdateAccommodation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.UpdateAccommodation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccommodationRequest struct via the builder pattern


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


## UpdateRoomType

> UpdateRoomType(ctx, id, roomTypeId).Execute()

Update Room Type

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
	roomTypeId := "roomTypeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.UpdateRoomType(context.Background(), id, roomTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.UpdateRoomType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomTypeRequest struct via the builder pattern


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


## UpdateRoomTypePrice

> UpdateRoomTypePrice(ctx, id, roomTypeId, priceId).Execute()

Update Room Type Price

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
	roomTypeId := "roomTypeId_example" // string | 
	priceId := "priceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccommodationsAPI.UpdateRoomTypePrice(context.Background(), id, roomTypeId, priceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccommodationsAPI.UpdateRoomTypePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**roomTypeId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomTypePriceRequest struct via the builder pattern


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

