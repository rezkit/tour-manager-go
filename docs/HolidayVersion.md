# HolidayVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**HolidayId** | Pointer to **string** | Holiday ID the version belongs to. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of most recent update | [optional] 
**Name** | Pointer to **string** | Holiday Version Name | [optional] 
**Code** | Pointer to **string** | Holiday Version Code | [optional] 
**Introduction** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Holiday version description (HTML Content) | [optional] 
**Published** | Pointer to **bool** |  | [optional] 

## Methods

### NewHolidayVersion

`func NewHolidayVersion() *HolidayVersion`

NewHolidayVersion instantiates a new HolidayVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHolidayVersionWithDefaults

`func NewHolidayVersionWithDefaults() *HolidayVersion`

NewHolidayVersionWithDefaults instantiates a new HolidayVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HolidayVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HolidayVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HolidayVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HolidayVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHolidayId

`func (o *HolidayVersion) GetHolidayId() string`

GetHolidayId returns the HolidayId field if non-nil, zero value otherwise.

### GetHolidayIdOk

`func (o *HolidayVersion) GetHolidayIdOk() (*string, bool)`

GetHolidayIdOk returns a tuple with the HolidayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayId

`func (o *HolidayVersion) SetHolidayId(v string)`

SetHolidayId sets HolidayId field to given value.

### HasHolidayId

`func (o *HolidayVersion) HasHolidayId() bool`

HasHolidayId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HolidayVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HolidayVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HolidayVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HolidayVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HolidayVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HolidayVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HolidayVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HolidayVersion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *HolidayVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HolidayVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HolidayVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HolidayVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *HolidayVersion) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HolidayVersion) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HolidayVersion) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *HolidayVersion) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIntroduction

`func (o *HolidayVersion) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *HolidayVersion) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *HolidayVersion) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *HolidayVersion) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetDescription

`func (o *HolidayVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HolidayVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HolidayVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HolidayVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublished

`func (o *HolidayVersion) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *HolidayVersion) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *HolidayVersion) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *HolidayVersion) HasPublished() bool`

HasPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


