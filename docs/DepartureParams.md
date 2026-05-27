# DepartureParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** |  | [optional] 
**End** | Pointer to **time.Time** |  | [optional] 
**RangeType** | Pointer to [**DepartureRangeType**](DepartureRangeType.md) |  | [optional] 
**Inventory** | Pointer to [**Inventory**](Inventory.md) |  | [optional] 

## Methods

### NewDepartureParams

`func NewDepartureParams() *DepartureParams`

NewDepartureParams instantiates a new DepartureParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartureParamsWithDefaults

`func NewDepartureParamsWithDefaults() *DepartureParams`

NewDepartureParamsWithDefaults instantiates a new DepartureParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *DepartureParams) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DepartureParams) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DepartureParams) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *DepartureParams) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *DepartureParams) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DepartureParams) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DepartureParams) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *DepartureParams) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetRangeType

`func (o *DepartureParams) GetRangeType() DepartureRangeType`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *DepartureParams) GetRangeTypeOk() (*DepartureRangeType, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *DepartureParams) SetRangeType(v DepartureRangeType)`

SetRangeType sets RangeType field to given value.

### HasRangeType

`func (o *DepartureParams) HasRangeType() bool`

HasRangeType returns a boolean if a field has been set.

### GetInventory

`func (o *DepartureParams) GetInventory() Inventory`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *DepartureParams) GetInventoryOk() (*Inventory, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *DepartureParams) SetInventory(v Inventory)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *DepartureParams) HasInventory() bool`

HasInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


