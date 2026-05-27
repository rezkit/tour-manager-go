# \PricesAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrice**](PricesAPI.md#GetPrice) | **Get** /holidays/prices/{price} | Get Price
[**UpdatePrice**](PricesAPI.md#UpdatePrice) | **Patch** /holidays/prices/{price} | Update Price



## GetPrice

> Price GetPrice(ctx, price).Execute()

Get Price



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
	price := "price_example" // string | Price ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.GetPrice(context.Background(), price).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.GetPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrice`: Price
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.GetPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**price** | **string** | Price ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Price**](Price.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrice

> Price UpdatePrice(ctx, price).UpdatePriceRequest(updatePriceRequest).Execute()

Update Price



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
	price := "price_example" // string | Price ID
	updatePriceRequest := *openapiclient.NewUpdatePriceRequest() // UpdatePriceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.UpdatePrice(context.Background(), price).UpdatePriceRequest(updatePriceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.UpdatePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrice`: Price
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.UpdatePrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**price** | **string** | Price ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePriceRequest** | [**UpdatePriceRequest**](UpdatePriceRequest.md) |  | 

### Return type

[**Price**](Price.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

