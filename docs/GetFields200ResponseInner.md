# GetFields200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**Label** | Pointer to **string** | Field group label | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Fields** | Pointer to [**[]FieldDefinition**](FieldDefinition.md) |  | [optional] 

## Methods

### NewGetFields200ResponseInner

`func NewGetFields200ResponseInner() *GetFields200ResponseInner`

NewGetFields200ResponseInner instantiates a new GetFields200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFields200ResponseInnerWithDefaults

`func NewGetFields200ResponseInnerWithDefaults() *GetFields200ResponseInner`

NewGetFields200ResponseInnerWithDefaults instantiates a new GetFields200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetFields200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetFields200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetFields200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetFields200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetFields200ResponseInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetFields200ResponseInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetFields200ResponseInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetFields200ResponseInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRank

`func (o *GetFields200ResponseInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *GetFields200ResponseInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *GetFields200ResponseInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *GetFields200ResponseInner) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetFields

`func (o *GetFields200ResponseInner) GetFields() []FieldDefinition`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetFields200ResponseInner) GetFieldsOk() (*[]FieldDefinition, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetFields200ResponseInner) SetFields(v []FieldDefinition)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GetFields200ResponseInner) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


