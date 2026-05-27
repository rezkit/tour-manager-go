# BooleanFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**Name** | **string** | Field Name | 
**GroupId** | **string** | Field group title | 
**Rank** | Pointer to **int32** |  | [optional] 
**Label** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewBooleanFieldDefinition

`func NewBooleanFieldDefinition(id string, name string, groupId string, label string, type_ string, ) *BooleanFieldDefinition`

NewBooleanFieldDefinition instantiates a new BooleanFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanFieldDefinitionWithDefaults

`func NewBooleanFieldDefinitionWithDefaults() *BooleanFieldDefinition`

NewBooleanFieldDefinitionWithDefaults instantiates a new BooleanFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BooleanFieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BooleanFieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BooleanFieldDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BooleanFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BooleanFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BooleanFieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetGroupId

`func (o *BooleanFieldDefinition) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BooleanFieldDefinition) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BooleanFieldDefinition) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetRank

`func (o *BooleanFieldDefinition) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *BooleanFieldDefinition) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *BooleanFieldDefinition) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *BooleanFieldDefinition) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLabel

`func (o *BooleanFieldDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BooleanFieldDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BooleanFieldDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *BooleanFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BooleanFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BooleanFieldDefinition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


