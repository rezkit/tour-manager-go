# AllocationInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Capacity** | **int32** | Total number of units in inventory. | 

## Methods

### NewAllocationInventory

`func NewAllocationInventory(type_ string, capacity int32, ) *AllocationInventory`

NewAllocationInventory instantiates a new AllocationInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationInventoryWithDefaults

`func NewAllocationInventoryWithDefaults() *AllocationInventory`

NewAllocationInventoryWithDefaults instantiates a new AllocationInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AllocationInventory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AllocationInventory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AllocationInventory) SetType(v string)`

SetType sets Type field to given value.


### GetCapacity

`func (o *AllocationInventory) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *AllocationInventory) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *AllocationInventory) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


