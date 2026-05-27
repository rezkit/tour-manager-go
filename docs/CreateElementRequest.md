# CreateElementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CategoryId** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**DefaultInventory** | Pointer to [**Inventory**](Inventory.md) |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] [default to true]
**IsPackage** | Pointer to **bool** |  | [optional] [default to true]
**BalanceDue** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateElementRequest

`func NewCreateElementRequest(name string, categoryId string, ) *CreateElementRequest`

NewCreateElementRequest instantiates a new CreateElementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateElementRequestWithDefaults

`func NewCreateElementRequestWithDefaults() *CreateElementRequest`

NewCreateElementRequestWithDefaults instantiates a new CreateElementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateElementRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateElementRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateElementRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCategoryId

`func (o *CreateElementRequest) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CreateElementRequest) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CreateElementRequest) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetDefaultInventory

`func (o *CreateElementRequest) GetDefaultInventory() Inventory`

GetDefaultInventory returns the DefaultInventory field if non-nil, zero value otherwise.

### GetDefaultInventoryOk

`func (o *CreateElementRequest) GetDefaultInventoryOk() (*Inventory, bool)`

GetDefaultInventoryOk returns a tuple with the DefaultInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInventory

`func (o *CreateElementRequest) SetDefaultInventory(v Inventory)`

SetDefaultInventory sets DefaultInventory field to given value.

### HasDefaultInventory

`func (o *CreateElementRequest) HasDefaultInventory() bool`

HasDefaultInventory returns a boolean if a field has been set.

### GetPublished

`func (o *CreateElementRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CreateElementRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CreateElementRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CreateElementRequest) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetIsPackage

`func (o *CreateElementRequest) GetIsPackage() bool`

GetIsPackage returns the IsPackage field if non-nil, zero value otherwise.

### GetIsPackageOk

`func (o *CreateElementRequest) GetIsPackageOk() (*bool, bool)`

GetIsPackageOk returns a tuple with the IsPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPackage

`func (o *CreateElementRequest) SetIsPackage(v bool)`

SetIsPackage sets IsPackage field to given value.

### HasIsPackage

`func (o *CreateElementRequest) HasIsPackage() bool`

HasIsPackage returns a boolean if a field has been set.

### GetBalanceDue

`func (o *CreateElementRequest) GetBalanceDue() int32`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *CreateElementRequest) GetBalanceDueOk() (*int32, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *CreateElementRequest) SetBalanceDue(v int32)`

SetBalanceDue sets BalanceDue field to given value.

### HasBalanceDue

`func (o *CreateElementRequest) HasBalanceDue() bool`

HasBalanceDue returns a boolean if a field has been set.

### SetBalanceDueNil

`func (o *CreateElementRequest) SetBalanceDueNil(b bool)`

 SetBalanceDueNil sets the value for BalanceDue to be an explicit nil

### UnsetBalanceDue
`func (o *CreateElementRequest) UnsetBalanceDue()`

UnsetBalanceDue ensures that no value is present for BalanceDue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


