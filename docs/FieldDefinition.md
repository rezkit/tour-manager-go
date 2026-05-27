# FieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**Name** | **string** | Field Name | 
**GroupId** | **string** | Field group title | 
**Rank** | Pointer to **int32** |  | [optional] 
**Label** | **string** |  | 
**Required** | **bool** | Determines if this field is required | 
**Type** | **string** | ## Field Type  Determines the specific type of field. All the field types for a text field handle string values but the specific &#x60;type&#x60; value determines the semantic input type used  * &#x60;text&#x60; - Plain text field * &#x60;rich_text&#x60; - Rich text editor * &#x60;url&#x60; - URL field, value must be a valid URL * &#x60;color&#x60; - Color field, value must a valid hex colour  | 
**MinLength** | Pointer to **int32** | Minimum value length | [optional] 
**MaxLength** | Pointer to **int32** | Minimum value length | [optional] 
**Format** | Pointer to **string** | RegEx for field value validation. Applicable only when &#x60;type&#x60; is &#x60;text&#x60; | [optional] 
**MinimumValue** | Pointer to **int32** | Minimum allowed value (inclusive) | [optional] 
**MaximumValue** | Pointer to **int32** | Maximum allowed value (inclusive) | [optional] 
**Precision** | **int32** | Decimal places | 
**Group** | **string** | Field group title | 
**Values** | **[]string** | Allowed values | 

## Methods

### NewFieldDefinition

`func NewFieldDefinition(id string, name string, groupId string, label string, required bool, type_ string, precision int32, group string, values []string, ) *FieldDefinition`

NewFieldDefinition instantiates a new FieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDefinitionWithDefaults

`func NewFieldDefinitionWithDefaults() *FieldDefinition`

NewFieldDefinitionWithDefaults instantiates a new FieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetGroupId

`func (o *FieldDefinition) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FieldDefinition) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FieldDefinition) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetRank

`func (o *FieldDefinition) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *FieldDefinition) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *FieldDefinition) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *FieldDefinition) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLabel

`func (o *FieldDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FieldDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FieldDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequired

`func (o *FieldDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetType

`func (o *FieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetMinLength

`func (o *FieldDefinition) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *FieldDefinition) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *FieldDefinition) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *FieldDefinition) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *FieldDefinition) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *FieldDefinition) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *FieldDefinition) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *FieldDefinition) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetFormat

`func (o *FieldDefinition) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *FieldDefinition) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *FieldDefinition) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *FieldDefinition) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMinimumValue

`func (o *FieldDefinition) GetMinimumValue() int32`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *FieldDefinition) GetMinimumValueOk() (*int32, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *FieldDefinition) SetMinimumValue(v int32)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *FieldDefinition) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.

### GetMaximumValue

`func (o *FieldDefinition) GetMaximumValue() int32`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *FieldDefinition) GetMaximumValueOk() (*int32, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *FieldDefinition) SetMaximumValue(v int32)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *FieldDefinition) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetPrecision

`func (o *FieldDefinition) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *FieldDefinition) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *FieldDefinition) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.


### GetGroup

`func (o *FieldDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FieldDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FieldDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetValues

`func (o *FieldDefinition) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FieldDefinition) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FieldDefinition) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


