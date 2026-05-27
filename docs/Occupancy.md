# Occupancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **int32** | Minimum number of passengers required to book. | 
**To** | **int32** | Maximum number of passengers bookable per reservation. | 

## Methods

### NewOccupancy

`func NewOccupancy(from int32, to int32, ) *Occupancy`

NewOccupancy instantiates a new Occupancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOccupancyWithDefaults

`func NewOccupancyWithDefaults() *Occupancy`

NewOccupancyWithDefaults instantiates a new Occupancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *Occupancy) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Occupancy) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Occupancy) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *Occupancy) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Occupancy) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Occupancy) SetTo(v int32)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


