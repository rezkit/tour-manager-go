# UpdateElementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CategoryId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**DefaultInventory** | Pointer to [**Inventory**](Inventory.md) |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**IsPackage** | Pointer to **bool** |  | [optional] 
**BalanceDue** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUpdateElementRequest

`func NewUpdateElementRequest() *UpdateElementRequest`

NewUpdateElementRequest instantiates a new UpdateElementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateElementRequestWithDefaults

`func NewUpdateElementRequestWithDefaults() *UpdateElementRequest`

NewUpdateElementRequestWithDefaults instantiates a new UpdateElementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateElementRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateElementRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateElementRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateElementRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategoryId

`func (o *UpdateElementRequest) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *UpdateElementRequest) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *UpdateElementRequest) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *UpdateElementRequest) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDefaultInventory

`func (o *UpdateElementRequest) GetDefaultInventory() Inventory`

GetDefaultInventory returns the DefaultInventory field if non-nil, zero value otherwise.

### GetDefaultInventoryOk

`func (o *UpdateElementRequest) GetDefaultInventoryOk() (*Inventory, bool)`

GetDefaultInventoryOk returns a tuple with the DefaultInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInventory

`func (o *UpdateElementRequest) SetDefaultInventory(v Inventory)`

SetDefaultInventory sets DefaultInventory field to given value.

### HasDefaultInventory

`func (o *UpdateElementRequest) HasDefaultInventory() bool`

HasDefaultInventory returns a boolean if a field has been set.

### GetPublished

`func (o *UpdateElementRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *UpdateElementRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *UpdateElementRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *UpdateElementRequest) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetIsPackage

`func (o *UpdateElementRequest) GetIsPackage() bool`

GetIsPackage returns the IsPackage field if non-nil, zero value otherwise.

### GetIsPackageOk

`func (o *UpdateElementRequest) GetIsPackageOk() (*bool, bool)`

GetIsPackageOk returns a tuple with the IsPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPackage

`func (o *UpdateElementRequest) SetIsPackage(v bool)`

SetIsPackage sets IsPackage field to given value.

### HasIsPackage

`func (o *UpdateElementRequest) HasIsPackage() bool`

HasIsPackage returns a boolean if a field has been set.

### GetBalanceDue

`func (o *UpdateElementRequest) GetBalanceDue() int32`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *UpdateElementRequest) GetBalanceDueOk() (*int32, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *UpdateElementRequest) SetBalanceDue(v int32)`

SetBalanceDue sets BalanceDue field to given value.

### HasBalanceDue

`func (o *UpdateElementRequest) HasBalanceDue() bool`

HasBalanceDue returns a boolean if a field has been set.

### SetBalanceDueNil

`func (o *UpdateElementRequest) SetBalanceDueNil(b bool)`

 SetBalanceDueNil sets the value for BalanceDue to be an explicit nil

### UnsetBalanceDue
`func (o *UpdateElementRequest) UnsetBalanceDue()`

UnsetBalanceDue ensures that no value is present for BalanceDue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


