# \ConfigurationAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateOrganizationSettings**](ConfigurationAPI.md#UpdateOrganizationSettings) | **Put** /organization/settings | Update Settings



## UpdateOrganizationSettings

> Organization UpdateOrganizationSettings(ctx).UpdateOrganizationSettingsRequest(updateOrganizationSettingsRequest).Execute()

Update Settings



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
	updateOrganizationSettingsRequest := *openapiclient.NewUpdateOrganizationSettingsRequest() // UpdateOrganizationSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.UpdateOrganizationSettings(context.Background()).UpdateOrganizationSettingsRequest(updateOrganizationSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.UpdateOrganizationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSettings`: Organization
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.UpdateOrganizationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrganizationSettingsRequest** | [**UpdateOrganizationSettingsRequest**](UpdateOrganizationSettingsRequest.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

