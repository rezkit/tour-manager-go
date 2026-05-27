# Cake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**ParentId** | Pointer to **NullableString** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DisplayType** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**RequiredToConfirm** | Pointer to **bool** |  | [optional] 
**RequiredBy** | Pointer to **string** |  | [optional] 
**Validation** | Pointer to **string** |  | [optional] 
**Global** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of most recent update | [optional] 
**DeletedAt** | Pointer to **NullableTime** | Timestamp of the item&#39;s deletion. Null if the item is not deleted. | [optional] 

## Methods

### NewCake

`func NewCake() *Cake`

NewCake instantiates a new Cake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCakeWithDefaults

`func NewCakeWithDefaults() *Cake`

NewCakeWithDefaults instantiates a new Cake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cake) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cake) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cake) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cake) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *Cake) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Cake) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Cake) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Cake) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *Cake) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Cake) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetName

`func (o *Cake) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cake) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cake) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cake) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Cake) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cake) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cake) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Cake) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayType

`func (o *Cake) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *Cake) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *Cake) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *Cake) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetLabel

`func (o *Cake) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Cake) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Cake) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Cake) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *Cake) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cake) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cake) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cake) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublished

`func (o *Cake) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *Cake) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *Cake) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *Cake) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRequiredToConfirm

`func (o *Cake) GetRequiredToConfirm() bool`

GetRequiredToConfirm returns the RequiredToConfirm field if non-nil, zero value otherwise.

### GetRequiredToConfirmOk

`func (o *Cake) GetRequiredToConfirmOk() (*bool, bool)`

GetRequiredToConfirmOk returns a tuple with the RequiredToConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredToConfirm

`func (o *Cake) SetRequiredToConfirm(v bool)`

SetRequiredToConfirm sets RequiredToConfirm field to given value.

### HasRequiredToConfirm

`func (o *Cake) HasRequiredToConfirm() bool`

HasRequiredToConfirm returns a boolean if a field has been set.

### GetRequiredBy

`func (o *Cake) GetRequiredBy() string`

GetRequiredBy returns the RequiredBy field if non-nil, zero value otherwise.

### GetRequiredByOk

`func (o *Cake) GetRequiredByOk() (*string, bool)`

GetRequiredByOk returns a tuple with the RequiredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBy

`func (o *Cake) SetRequiredBy(v string)`

SetRequiredBy sets RequiredBy field to given value.

### HasRequiredBy

`func (o *Cake) HasRequiredBy() bool`

HasRequiredBy returns a boolean if a field has been set.

### GetValidation

`func (o *Cake) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *Cake) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *Cake) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *Cake) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### GetGlobal

`func (o *Cake) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *Cake) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *Cake) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *Cake) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Cake) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cake) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cake) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Cake) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Cake) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cake) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cake) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Cake) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Cake) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Cake) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Cake) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Cake) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *Cake) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *Cake) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


