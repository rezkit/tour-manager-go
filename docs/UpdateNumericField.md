# UpdateNumericField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**MinimumValue** | Pointer to **NullableInt32** | Minimum allowed value (inclusive) | [optional] 
**MaximumValue** | Pointer to **NullableInt32** | Maximum allowed value (inclusive) | [optional] 
**Precision** | Pointer to **NullableInt32** | Decimal places | [optional] 

## Methods

### NewUpdateNumericField

`func NewUpdateNumericField() *UpdateNumericField`

NewUpdateNumericField instantiates a new UpdateNumericField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNumericFieldWithDefaults

`func NewUpdateNumericFieldWithDefaults() *UpdateNumericField`

NewUpdateNumericFieldWithDefaults instantiates a new UpdateNumericField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UpdateNumericField) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateNumericField) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateNumericField) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UpdateNumericField) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateNumericField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateNumericField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateNumericField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateNumericField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMinimumValue

`func (o *UpdateNumericField) GetMinimumValue() int32`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *UpdateNumericField) GetMinimumValueOk() (*int32, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *UpdateNumericField) SetMinimumValue(v int32)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *UpdateNumericField) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.

### SetMinimumValueNil

`func (o *UpdateNumericField) SetMinimumValueNil(b bool)`

 SetMinimumValueNil sets the value for MinimumValue to be an explicit nil

### UnsetMinimumValue
`func (o *UpdateNumericField) UnsetMinimumValue()`

UnsetMinimumValue ensures that no value is present for MinimumValue, not even an explicit nil
### GetMaximumValue

`func (o *UpdateNumericField) GetMaximumValue() int32`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *UpdateNumericField) GetMaximumValueOk() (*int32, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *UpdateNumericField) SetMaximumValue(v int32)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *UpdateNumericField) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### SetMaximumValueNil

`func (o *UpdateNumericField) SetMaximumValueNil(b bool)`

 SetMaximumValueNil sets the value for MaximumValue to be an explicit nil

### UnsetMaximumValue
`func (o *UpdateNumericField) UnsetMaximumValue()`

UnsetMaximumValue ensures that no value is present for MaximumValue, not even an explicit nil
### GetPrecision

`func (o *UpdateNumericField) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *UpdateNumericField) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *UpdateNumericField) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *UpdateNumericField) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### SetPrecisionNil

`func (o *UpdateNumericField) SetPrecisionNil(b bool)`

 SetPrecisionNil sets the value for Precision to be an explicit nil

### UnsetPrecision
`func (o *UpdateNumericField) UnsetPrecision()`

UnsetPrecision ensures that no value is present for Precision, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


