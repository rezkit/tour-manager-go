# DepartureProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** |  | 
**End** | **time.Time** |  | 
**RangeType** | Pointer to [**DepartureRangeType**](DepartureRangeType.md) |  | [optional] 
**Inventory** | [**Inventory**](Inventory.md) |  | 

## Methods

### NewDepartureProperties

`func NewDepartureProperties(start time.Time, end time.Time, inventory Inventory, ) *DepartureProperties`

NewDepartureProperties instantiates a new DepartureProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeparturePropertiesWithDefaults

`func NewDeparturePropertiesWithDefaults() *DepartureProperties`

NewDeparturePropertiesWithDefaults instantiates a new DepartureProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *DepartureProperties) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DepartureProperties) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DepartureProperties) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *DepartureProperties) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DepartureProperties) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DepartureProperties) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetRangeType

`func (o *DepartureProperties) GetRangeType() DepartureRangeType`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *DepartureProperties) GetRangeTypeOk() (*DepartureRangeType, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *DepartureProperties) SetRangeType(v DepartureRangeType)`

SetRangeType sets RangeType field to given value.

### HasRangeType

`func (o *DepartureProperties) HasRangeType() bool`

HasRangeType returns a boolean if a field has been set.

### GetInventory

`func (o *DepartureProperties) GetInventory() Inventory`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *DepartureProperties) GetInventoryOk() (*Inventory, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *DepartureProperties) SetInventory(v Inventory)`

SetInventory sets Inventory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


