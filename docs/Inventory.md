# Inventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Capacity** | Pointer to **int32** |  | [optional] 
**Available** | Pointer to **int32** |  | [optional] 
**Errata** | Pointer to **string** |  | [optional] 

## Methods

### NewInventory

`func NewInventory(type_ string, ) *Inventory`

NewInventory instantiates a new Inventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryWithDefaults

`func NewInventoryWithDefaults() *Inventory`

NewInventoryWithDefaults instantiates a new Inventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Inventory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Inventory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Inventory) SetType(v string)`

SetType sets Type field to given value.


### GetCapacity

`func (o *Inventory) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *Inventory) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *Inventory) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *Inventory) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetAvailable

`func (o *Inventory) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Inventory) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Inventory) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Inventory) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetErrata

`func (o *Inventory) GetErrata() string`

GetErrata returns the Errata field if non-nil, zero value otherwise.

### GetErrataOk

`func (o *Inventory) GetErrataOk() (*string, bool)`

GetErrataOk returns a tuple with the Errata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrata

`func (o *Inventory) SetErrata(v string)`

SetErrata sets Errata field to given value.

### HasErrata

`func (o *Inventory) HasErrata() bool`

HasErrata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


