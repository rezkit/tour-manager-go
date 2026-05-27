# \FieldsAPI

All URIs are relative to *https://tours.api.rezkit.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateField**](FieldsAPI.md#CreateField) | **Post** /fields/{type} | Create Field
[**CreateGroup**](FieldsAPI.md#CreateGroup) | **Post** /fields/{type}/groups | Create Group
[**GetFields**](FieldsAPI.md#GetFields) | **Get** /fields/{type} | List Fields
[**UpdateField**](FieldsAPI.md#UpdateField) | **Patch** /fields/{type}/{name} | Update Field



## CreateField

> FieldDefinition CreateField(ctx, type_).FieldDefinition(fieldDefinition).Execute()

Create Field



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | Item type
	fieldDefinition := openapiclient.FieldDefinition{BooleanFieldDefinition: openapiclient.NewBooleanFieldDefinition("01gpkgcy6t0m84czh8gy4kjatk", "Name_example", "GroupId_example", "Label_example", "Type_example")} // FieldDefinition | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldsAPI.CreateField(context.Background(), type_).FieldDefinition(fieldDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldsAPI.CreateField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateField`: FieldDefinition
	fmt.Fprintf(os.Stdout, "Response from `FieldsAPI.CreateField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) | Item type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldDefinition** | [**FieldDefinition**](FieldDefinition.md) |  | 

### Return type

[**FieldDefinition**](FieldDefinition.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> FieldGroup CreateGroup(ctx, type_).CreateGroupRequest(createGroupRequest).Execute()

Create Group



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | Item type
	createGroupRequest := *openapiclient.NewCreateGroupRequest("Name_example") // CreateGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldsAPI.CreateGroup(context.Background(), type_).CreateGroupRequest(createGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldsAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: FieldGroup
	fmt.Fprintf(os.Stdout, "Response from `FieldsAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) | Item type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupRequest** | [**CreateGroupRequest**](CreateGroupRequest.md) |  | 

### Return type

[**FieldGroup**](FieldGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFields

> []GetFields200ResponseInner GetFields(ctx, type_).Execute()

List Fields



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | Item type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldsAPI.GetFields(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldsAPI.GetFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFields`: []GetFields200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `FieldsAPI.GetFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) | Item type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetFields200ResponseInner**](GetFields200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateField

> FieldDefinition UpdateField(ctx, type_, name).UpdateFieldRequest(updateFieldRequest).Execute()

Update Field



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
	type_ := openapiclient.EntityType("accommodation") // EntityType | Item type
	name := "name_example" // string | Field Name
	updateFieldRequest := openapiclient.updateField_request{UpdateNumericField: openapiclient.NewUpdateNumericField()} // UpdateFieldRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldsAPI.UpdateField(context.Background(), type_, name).UpdateFieldRequest(updateFieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldsAPI.UpdateField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateField`: FieldDefinition
	fmt.Fprintf(os.Stdout, "Response from `FieldsAPI.UpdateField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EntityType**](.md) | Item type | 
**name** | **string** | Field Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFieldRequest** | [**UpdateFieldRequest**](UpdateFieldRequest.md) |  | 

### Return type

[**FieldDefinition**](FieldDefinition.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

