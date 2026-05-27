# AddRelatedHolidayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**CategoryId** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**StartDay** | Pointer to **int32** |  | [optional] 
**Published** | **bool** |  | 
**CreateTwin** | Pointer to **bool** | Also create the reverse relationship. | [optional] 

## Methods

### NewAddRelatedHolidayRequest

`func NewAddRelatedHolidayRequest(destinationId string, categoryId string, published bool, ) *AddRelatedHolidayRequest`

NewAddRelatedHolidayRequest instantiates a new AddRelatedHolidayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRelatedHolidayRequestWithDefaults

`func NewAddRelatedHolidayRequestWithDefaults() *AddRelatedHolidayRequest`

NewAddRelatedHolidayRequestWithDefaults instantiates a new AddRelatedHolidayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *AddRelatedHolidayRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *AddRelatedHolidayRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *AddRelatedHolidayRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetCategoryId

`func (o *AddRelatedHolidayRequest) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *AddRelatedHolidayRequest) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *AddRelatedHolidayRequest) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetStartDay

`func (o *AddRelatedHolidayRequest) GetStartDay() int32`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *AddRelatedHolidayRequest) GetStartDayOk() (*int32, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *AddRelatedHolidayRequest) SetStartDay(v int32)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *AddRelatedHolidayRequest) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### GetPublished

`func (o *AddRelatedHolidayRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AddRelatedHolidayRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AddRelatedHolidayRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetCreateTwin

`func (o *AddRelatedHolidayRequest) GetCreateTwin() bool`

GetCreateTwin returns the CreateTwin field if non-nil, zero value otherwise.

### GetCreateTwinOk

`func (o *AddRelatedHolidayRequest) GetCreateTwinOk() (*bool, bool)`

GetCreateTwinOk returns a tuple with the CreateTwin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTwin

`func (o *AddRelatedHolidayRequest) SetCreateTwin(v bool)`

SetCreateTwin sets CreateTwin field to given value.

### HasCreateTwin

`func (o *AddRelatedHolidayRequest) HasCreateTwin() bool`

HasCreateTwin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


