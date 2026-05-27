# HolidayEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**SourceId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**DestinationId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**CategoryId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**StartDay** | Pointer to **int32** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Ordering** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of most recent update | [optional] 

## Methods

### NewHolidayEdge

`func NewHolidayEdge() *HolidayEdge`

NewHolidayEdge instantiates a new HolidayEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHolidayEdgeWithDefaults

`func NewHolidayEdgeWithDefaults() *HolidayEdge`

NewHolidayEdgeWithDefaults instantiates a new HolidayEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HolidayEdge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HolidayEdge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HolidayEdge) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HolidayEdge) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceId

`func (o *HolidayEdge) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HolidayEdge) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HolidayEdge) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *HolidayEdge) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDestinationId

`func (o *HolidayEdge) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *HolidayEdge) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *HolidayEdge) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *HolidayEdge) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetCategoryId

`func (o *HolidayEdge) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *HolidayEdge) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *HolidayEdge) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *HolidayEdge) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetStartDay

`func (o *HolidayEdge) GetStartDay() int32`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *HolidayEdge) GetStartDayOk() (*int32, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *HolidayEdge) SetStartDay(v int32)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *HolidayEdge) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### GetPublished

`func (o *HolidayEdge) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *HolidayEdge) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *HolidayEdge) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *HolidayEdge) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetOrdering

`func (o *HolidayEdge) GetOrdering() int32`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *HolidayEdge) GetOrderingOk() (*int32, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *HolidayEdge) SetOrdering(v int32)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *HolidayEdge) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HolidayEdge) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HolidayEdge) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HolidayEdge) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HolidayEdge) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HolidayEdge) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HolidayEdge) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HolidayEdge) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HolidayEdge) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


