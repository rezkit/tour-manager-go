# TextFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**Name** | **string** | Field Name | 
**GroupId** | **string** | Field group ID | 
**Rank** | Pointer to **int32** |  | [optional] 
**Label** | **string** |  | 
**Required** | **bool** | Determines if this field is required | 
**Type** | **string** | ## Field Type  Determines the specific type of field. All the field types for a text field handle string values but the specific &#x60;type&#x60; value determines the semantic input type used  * &#x60;text&#x60; - Plain text field * &#x60;rich_text&#x60; - Rich text editor * &#x60;url&#x60; - URL field, value must be a valid URL * &#x60;color&#x60; - Color field, value must a valid hex colour  | 
**MinLength** | Pointer to **int32** | Minimum value length | [optional] 
**MaxLength** | Pointer to **int32** | Minimum value length | [optional] 
**Format** | Pointer to **string** | RegEx for field value validation. Applicable only when &#x60;type&#x60; is &#x60;text&#x60; | [optional] 

## Methods

### NewTextFieldDefinition

`func NewTextFieldDefinition(id string, name string, groupId string, label string, required bool, type_ string, ) *TextFieldDefinition`

NewTextFieldDefinition instantiates a new TextFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextFieldDefinitionWithDefaults

`func NewTextFieldDefinitionWithDefaults() *TextFieldDefinition`

NewTextFieldDefinitionWithDefaults instantiates a new TextFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TextFieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextFieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextFieldDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TextFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TextFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TextFieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetGroupId

`func (o *TextFieldDefinition) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TextFieldDefinition) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TextFieldDefinition) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetRank

`func (o *TextFieldDefinition) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TextFieldDefinition) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TextFieldDefinition) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *TextFieldDefinition) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLabel

`func (o *TextFieldDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TextFieldDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TextFieldDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequired

`func (o *TextFieldDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *TextFieldDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *TextFieldDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetType

`func (o *TextFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextFieldDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetMinLength

`func (o *TextFieldDefinition) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *TextFieldDefinition) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *TextFieldDefinition) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *TextFieldDefinition) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *TextFieldDefinition) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *TextFieldDefinition) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *TextFieldDefinition) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *TextFieldDefinition) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetFormat

`func (o *TextFieldDefinition) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TextFieldDefinition) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TextFieldDefinition) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *TextFieldDefinition) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


