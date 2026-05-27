# UpdateFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**MinLength** | Pointer to **NullableInt32** |  | [optional] 
**MaxLength** | Pointer to **NullableInt32** |  | [optional] 
**MinimumValue** | Pointer to **NullableInt32** | Minimum allowed value (inclusive) | [optional] 
**MaximumValue** | Pointer to **NullableInt32** | Maximum allowed value (inclusive) | [optional] 
**Precision** | Pointer to **NullableInt32** | Decimal places | [optional] 

## Methods

### NewUpdateFieldRequest

`func NewUpdateFieldRequest() *UpdateFieldRequest`

NewUpdateFieldRequest instantiates a new UpdateFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFieldRequestWithDefaults

`func NewUpdateFieldRequestWithDefaults() *UpdateFieldRequest`

NewUpdateFieldRequestWithDefaults instantiates a new UpdateFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UpdateFieldRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateFieldRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateFieldRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UpdateFieldRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateFieldRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMinLength

`func (o *UpdateFieldRequest) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *UpdateFieldRequest) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *UpdateFieldRequest) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *UpdateFieldRequest) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *UpdateFieldRequest) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *UpdateFieldRequest) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *UpdateFieldRequest) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *UpdateFieldRequest) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *UpdateFieldRequest) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *UpdateFieldRequest) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *UpdateFieldRequest) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *UpdateFieldRequest) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetMinimumValue

`func (o *UpdateFieldRequest) GetMinimumValue() int32`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *UpdateFieldRequest) GetMinimumValueOk() (*int32, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *UpdateFieldRequest) SetMinimumValue(v int32)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *UpdateFieldRequest) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.

### SetMinimumValueNil

`func (o *UpdateFieldRequest) SetMinimumValueNil(b bool)`

 SetMinimumValueNil sets the value for MinimumValue to be an explicit nil

### UnsetMinimumValue
`func (o *UpdateFieldRequest) UnsetMinimumValue()`

UnsetMinimumValue ensures that no value is present for MinimumValue, not even an explicit nil
### GetMaximumValue

`func (o *UpdateFieldRequest) GetMaximumValue() int32`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *UpdateFieldRequest) GetMaximumValueOk() (*int32, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *UpdateFieldRequest) SetMaximumValue(v int32)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *UpdateFieldRequest) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### SetMaximumValueNil

`func (o *UpdateFieldRequest) SetMaximumValueNil(b bool)`

 SetMaximumValueNil sets the value for MaximumValue to be an explicit nil

### UnsetMaximumValue
`func (o *UpdateFieldRequest) UnsetMaximumValue()`

UnsetMaximumValue ensures that no value is present for MaximumValue, not even an explicit nil
### GetPrecision

`func (o *UpdateFieldRequest) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *UpdateFieldRequest) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *UpdateFieldRequest) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *UpdateFieldRequest) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### SetPrecisionNil

`func (o *UpdateFieldRequest) SetPrecisionNil(b bool)`

 SetPrecisionNil sets the value for Precision to be an explicit nil

### UnsetPrecision
`func (o *UpdateFieldRequest) UnsetPrecision()`

UnsetPrecision ensures that no value is present for Precision, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


