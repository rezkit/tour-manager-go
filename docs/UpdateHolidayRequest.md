# UpdateHolidayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Holiday Name | [optional] 
**Code** | Pointer to **string** | Holiday Code | [optional] 
**Introduction** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** | Holiday description (HTML Content) | [optional] 
**Fields** | Pointer to [**map[string]CustomFieldsDataValue**](CustomFieldsDataValue.md) | User-defined Fields Data | [optional] 
**Published** | Pointer to **bool** | Determines if the holiday is publicly viewable. | [optional] 
**Ordering** | Pointer to [**OrderingCommand**](OrderingCommand.md) |  | [optional] 

## Methods

### NewUpdateHolidayRequest

`func NewUpdateHolidayRequest() *UpdateHolidayRequest`

NewUpdateHolidayRequest instantiates a new UpdateHolidayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHolidayRequestWithDefaults

`func NewUpdateHolidayRequestWithDefaults() *UpdateHolidayRequest`

NewUpdateHolidayRequestWithDefaults instantiates a new UpdateHolidayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateHolidayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateHolidayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateHolidayRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateHolidayRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateHolidayRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateHolidayRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateHolidayRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateHolidayRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIntroduction

`func (o *UpdateHolidayRequest) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *UpdateHolidayRequest) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *UpdateHolidayRequest) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *UpdateHolidayRequest) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### SetIntroductionNil

`func (o *UpdateHolidayRequest) SetIntroductionNil(b bool)`

 SetIntroductionNil sets the value for Introduction to be an explicit nil

### UnsetIntroduction
`func (o *UpdateHolidayRequest) UnsetIntroduction()`

UnsetIntroduction ensures that no value is present for Introduction, not even an explicit nil
### GetDescription

`func (o *UpdateHolidayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateHolidayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateHolidayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateHolidayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateHolidayRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateHolidayRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFields

`func (o *UpdateHolidayRequest) GetFields() map[string]CustomFieldsDataValue`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UpdateHolidayRequest) GetFieldsOk() (*map[string]CustomFieldsDataValue, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UpdateHolidayRequest) SetFields(v map[string]CustomFieldsDataValue)`

SetFields sets Fields field to given value.

### HasFields

`func (o *UpdateHolidayRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetPublished

`func (o *UpdateHolidayRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *UpdateHolidayRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *UpdateHolidayRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *UpdateHolidayRequest) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetOrdering

`func (o *UpdateHolidayRequest) GetOrdering() OrderingCommand`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *UpdateHolidayRequest) GetOrderingOk() (*OrderingCommand, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *UpdateHolidayRequest) SetOrdering(v OrderingCommand)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *UpdateHolidayRequest) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


