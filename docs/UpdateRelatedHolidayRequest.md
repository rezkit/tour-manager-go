# UpdateRelatedHolidayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**StartDay** | Pointer to **int32** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Ordering** | Pointer to [**OrderingCommand**](OrderingCommand.md) |  | [optional] 

## Methods

### NewUpdateRelatedHolidayRequest

`func NewUpdateRelatedHolidayRequest() *UpdateRelatedHolidayRequest`

NewUpdateRelatedHolidayRequest instantiates a new UpdateRelatedHolidayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRelatedHolidayRequestWithDefaults

`func NewUpdateRelatedHolidayRequestWithDefaults() *UpdateRelatedHolidayRequest`

NewUpdateRelatedHolidayRequestWithDefaults instantiates a new UpdateRelatedHolidayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *UpdateRelatedHolidayRequest) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *UpdateRelatedHolidayRequest) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *UpdateRelatedHolidayRequest) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *UpdateRelatedHolidayRequest) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetStartDay

`func (o *UpdateRelatedHolidayRequest) GetStartDay() int32`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *UpdateRelatedHolidayRequest) GetStartDayOk() (*int32, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *UpdateRelatedHolidayRequest) SetStartDay(v int32)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *UpdateRelatedHolidayRequest) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### GetPublished

`func (o *UpdateRelatedHolidayRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *UpdateRelatedHolidayRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *UpdateRelatedHolidayRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *UpdateRelatedHolidayRequest) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetOrdering

`func (o *UpdateRelatedHolidayRequest) GetOrdering() OrderingCommand`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *UpdateRelatedHolidayRequest) GetOrderingOk() (*OrderingCommand, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *UpdateRelatedHolidayRequest) SetOrdering(v OrderingCommand)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *UpdateRelatedHolidayRequest) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


