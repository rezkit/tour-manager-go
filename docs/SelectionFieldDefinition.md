# SelectionFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**Name** | **string** | Field Name | 
**Group** | **string** | Field group title | 
**Rank** | Pointer to **int32** |  | [optional] 
**Label** | **string** |  | 
**Required** | **bool** | Determines if this field is required | 
**Type** | **string** |  | 
**Values** | **[]string** | Allowed values | 

## Methods

### NewSelectionFieldDefinition

`func NewSelectionFieldDefinition(id string, name string, group string, label string, required bool, type_ string, values []string, ) *SelectionFieldDefinition`

NewSelectionFieldDefinition instantiates a new SelectionFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectionFieldDefinitionWithDefaults

`func NewSelectionFieldDefinitionWithDefaults() *SelectionFieldDefinition`

NewSelectionFieldDefinitionWithDefaults instantiates a new SelectionFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelectionFieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectionFieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectionFieldDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SelectionFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectionFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectionFieldDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetGroup

`func (o *SelectionFieldDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SelectionFieldDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SelectionFieldDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetRank

`func (o *SelectionFieldDefinition) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SelectionFieldDefinition) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SelectionFieldDefinition) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SelectionFieldDefinition) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLabel

`func (o *SelectionFieldDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SelectionFieldDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SelectionFieldDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequired

`func (o *SelectionFieldDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SelectionFieldDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SelectionFieldDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetType

`func (o *SelectionFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectionFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectionFieldDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetValues

`func (o *SelectionFieldDefinition) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SelectionFieldDefinition) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SelectionFieldDefinition) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


