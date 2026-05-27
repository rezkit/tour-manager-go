# NumberFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**Name** | **string** | Field Name | 
**GroupId** | **string** | Field group ID | 
**Rank** | Pointer to **int32** |  | [optional] 
**Label** | **string** |  | 
**Required** | **bool** | Determines if this field is required | 
**Type** | **string** |  | 
**MinimumValue** | Pointer to **int32** | Minimum allowed value (inclusive) | [optional] 
**MaximumValue** | Pointer to **int32** | Maximum allowed value (inclusive) | [optional] 
**Precision** | **int32** | Decimal places | 

## Methods

### NewNumberFieldDefinition

`func NewNumberFieldDefinition(id string, name string, groupId string, label string, required bool, type_ string, precision int32, ) *NumberFieldDefinition`

NewNumberFieldDefinition instantiates a new NumberFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberFieldDefinitionWithDefaults

`func NewNumberFieldDefinitionWithDefaults() *NumberFieldDefinition`

NewNumberFieldDefinitionWithDefaults instantiates a new NumberFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NumberFieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumberFieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumberFieldDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NumberFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NumberFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NumberFieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetGroupId

`func (o *NumberFieldDefinition) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NumberFieldDefinition) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NumberFieldDefinition) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetRank

`func (o *NumberFieldDefinition) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *NumberFieldDefinition) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *NumberFieldDefinition) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *NumberFieldDefinition) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLabel

`func (o *NumberFieldDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NumberFieldDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NumberFieldDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequired

`func (o *NumberFieldDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NumberFieldDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NumberFieldDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetType

`func (o *NumberFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumberFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumberFieldDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetMinimumValue

`func (o *NumberFieldDefinition) GetMinimumValue() int32`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *NumberFieldDefinition) GetMinimumValueOk() (*int32, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *NumberFieldDefinition) SetMinimumValue(v int32)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *NumberFieldDefinition) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.

### GetMaximumValue

`func (o *NumberFieldDefinition) GetMaximumValue() int32`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *NumberFieldDefinition) GetMaximumValueOk() (*int32, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *NumberFieldDefinition) SetMaximumValue(v int32)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *NumberFieldDefinition) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetPrecision

`func (o *NumberFieldDefinition) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *NumberFieldDefinition) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *NumberFieldDefinition) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


