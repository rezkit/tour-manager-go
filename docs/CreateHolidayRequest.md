# CreateHolidayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Holiday Name | 
**Code** | **string** | Holiday Code | 
**Introduction** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Holiday description (HTML Content) | [optional] 
**Published** | Pointer to **bool** | Determines if the holiday is publicly viewable. | [optional] [default to false]

## Methods

### NewCreateHolidayRequest

`func NewCreateHolidayRequest(name string, code string, ) *CreateHolidayRequest`

NewCreateHolidayRequest instantiates a new CreateHolidayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHolidayRequestWithDefaults

`func NewCreateHolidayRequestWithDefaults() *CreateHolidayRequest`

NewCreateHolidayRequestWithDefaults instantiates a new CreateHolidayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateHolidayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHolidayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHolidayRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *CreateHolidayRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateHolidayRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateHolidayRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetIntroduction

`func (o *CreateHolidayRequest) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *CreateHolidayRequest) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *CreateHolidayRequest) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *CreateHolidayRequest) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetDescription

`func (o *CreateHolidayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateHolidayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateHolidayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateHolidayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublished

`func (o *CreateHolidayRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CreateHolidayRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CreateHolidayRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CreateHolidayRequest) HasPublished() bool`

HasPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


