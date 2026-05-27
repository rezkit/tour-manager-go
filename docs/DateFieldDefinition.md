# DateFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**Name** | **string** | Field Name | 
**GroupId** | **string** | Field group ID | 
**Rank** | Pointer to **int32** |  | [optional] 
**Label** | **string** |  | 
**Required** | **bool** | Determines if this field is required | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewDateFieldDefinition

`func NewDateFieldDefinition(id string, name string, groupId string, label string, required bool, ) *DateFieldDefinition`

NewDateFieldDefinition instantiates a new DateFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateFieldDefinitionWithDefaults

`func NewDateFieldDefinitionWithDefaults() *DateFieldDefinition`

NewDateFieldDefinitionWithDefaults instantiates a new DateFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DateFieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DateFieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DateFieldDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DateFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DateFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DateFieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetGroupId

`func (o *DateFieldDefinition) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DateFieldDefinition) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DateFieldDefinition) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetRank

`func (o *DateFieldDefinition) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *DateFieldDefinition) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *DateFieldDefinition) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *DateFieldDefinition) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLabel

`func (o *DateFieldDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DateFieldDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DateFieldDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequired

`func (o *DateFieldDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DateFieldDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DateFieldDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetType

`func (o *DateFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DateFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DateFieldDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DateFieldDefinition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


