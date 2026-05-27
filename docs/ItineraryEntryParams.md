# ItineraryEntryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDay** | Pointer to **int32** |  | [optional] 
**EndDay** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Introduction** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**IncludesBreakfast** | Pointer to **bool** |  | [optional] 
**IncludesLunch** | Pointer to **bool** |  | [optional] 
**IncludesDinner** | Pointer to **bool** |  | [optional] 

## Methods

### NewItineraryEntryParams

`func NewItineraryEntryParams() *ItineraryEntryParams`

NewItineraryEntryParams instantiates a new ItineraryEntryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItineraryEntryParamsWithDefaults

`func NewItineraryEntryParamsWithDefaults() *ItineraryEntryParams`

NewItineraryEntryParamsWithDefaults instantiates a new ItineraryEntryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDay

`func (o *ItineraryEntryParams) GetStartDay() int32`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *ItineraryEntryParams) GetStartDayOk() (*int32, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *ItineraryEntryParams) SetStartDay(v int32)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *ItineraryEntryParams) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### GetEndDay

`func (o *ItineraryEntryParams) GetEndDay() int32`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *ItineraryEntryParams) GetEndDayOk() (*int32, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *ItineraryEntryParams) SetEndDay(v int32)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *ItineraryEntryParams) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### GetTitle

`func (o *ItineraryEntryParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItineraryEntryParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItineraryEntryParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItineraryEntryParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ItineraryEntryParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItineraryEntryParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItineraryEntryParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItineraryEntryParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIntroduction

`func (o *ItineraryEntryParams) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *ItineraryEntryParams) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *ItineraryEntryParams) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *ItineraryEntryParams) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetPublished

`func (o *ItineraryEntryParams) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ItineraryEntryParams) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ItineraryEntryParams) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ItineraryEntryParams) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetIncludesBreakfast

`func (o *ItineraryEntryParams) GetIncludesBreakfast() bool`

GetIncludesBreakfast returns the IncludesBreakfast field if non-nil, zero value otherwise.

### GetIncludesBreakfastOk

`func (o *ItineraryEntryParams) GetIncludesBreakfastOk() (*bool, bool)`

GetIncludesBreakfastOk returns a tuple with the IncludesBreakfast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesBreakfast

`func (o *ItineraryEntryParams) SetIncludesBreakfast(v bool)`

SetIncludesBreakfast sets IncludesBreakfast field to given value.

### HasIncludesBreakfast

`func (o *ItineraryEntryParams) HasIncludesBreakfast() bool`

HasIncludesBreakfast returns a boolean if a field has been set.

### GetIncludesLunch

`func (o *ItineraryEntryParams) GetIncludesLunch() bool`

GetIncludesLunch returns the IncludesLunch field if non-nil, zero value otherwise.

### GetIncludesLunchOk

`func (o *ItineraryEntryParams) GetIncludesLunchOk() (*bool, bool)`

GetIncludesLunchOk returns a tuple with the IncludesLunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesLunch

`func (o *ItineraryEntryParams) SetIncludesLunch(v bool)`

SetIncludesLunch sets IncludesLunch field to given value.

### HasIncludesLunch

`func (o *ItineraryEntryParams) HasIncludesLunch() bool`

HasIncludesLunch returns a boolean if a field has been set.

### GetIncludesDinner

`func (o *ItineraryEntryParams) GetIncludesDinner() bool`

GetIncludesDinner returns the IncludesDinner field if non-nil, zero value otherwise.

### GetIncludesDinnerOk

`func (o *ItineraryEntryParams) GetIncludesDinnerOk() (*bool, bool)`

GetIncludesDinnerOk returns a tuple with the IncludesDinner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesDinner

`func (o *ItineraryEntryParams) SetIncludesDinner(v bool)`

SetIncludesDinner sets IncludesDinner field to given value.

### HasIncludesDinner

`func (o *ItineraryEntryParams) HasIncludesDinner() bool`

HasIncludesDinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


