# ItineraryEntryProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDay** | **int32** |  | 
**EndDay** | **int32** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Introduction** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**IncludesBreakfast** | Pointer to **bool** |  | [optional] 
**IncludesLunch** | Pointer to **bool** |  | [optional] 
**IncludesDinner** | Pointer to **bool** |  | [optional] 

## Methods

### NewItineraryEntryProperties

`func NewItineraryEntryProperties(startDay int32, endDay int32, title string, ) *ItineraryEntryProperties`

NewItineraryEntryProperties instantiates a new ItineraryEntryProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItineraryEntryPropertiesWithDefaults

`func NewItineraryEntryPropertiesWithDefaults() *ItineraryEntryProperties`

NewItineraryEntryPropertiesWithDefaults instantiates a new ItineraryEntryProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDay

`func (o *ItineraryEntryProperties) GetStartDay() int32`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *ItineraryEntryProperties) GetStartDayOk() (*int32, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *ItineraryEntryProperties) SetStartDay(v int32)`

SetStartDay sets StartDay field to given value.


### GetEndDay

`func (o *ItineraryEntryProperties) GetEndDay() int32`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *ItineraryEntryProperties) GetEndDayOk() (*int32, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *ItineraryEntryProperties) SetEndDay(v int32)`

SetEndDay sets EndDay field to given value.


### GetTitle

`func (o *ItineraryEntryProperties) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItineraryEntryProperties) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItineraryEntryProperties) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ItineraryEntryProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItineraryEntryProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItineraryEntryProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItineraryEntryProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIntroduction

`func (o *ItineraryEntryProperties) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *ItineraryEntryProperties) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *ItineraryEntryProperties) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *ItineraryEntryProperties) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetPublished

`func (o *ItineraryEntryProperties) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ItineraryEntryProperties) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ItineraryEntryProperties) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ItineraryEntryProperties) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetIncludesBreakfast

`func (o *ItineraryEntryProperties) GetIncludesBreakfast() bool`

GetIncludesBreakfast returns the IncludesBreakfast field if non-nil, zero value otherwise.

### GetIncludesBreakfastOk

`func (o *ItineraryEntryProperties) GetIncludesBreakfastOk() (*bool, bool)`

GetIncludesBreakfastOk returns a tuple with the IncludesBreakfast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesBreakfast

`func (o *ItineraryEntryProperties) SetIncludesBreakfast(v bool)`

SetIncludesBreakfast sets IncludesBreakfast field to given value.

### HasIncludesBreakfast

`func (o *ItineraryEntryProperties) HasIncludesBreakfast() bool`

HasIncludesBreakfast returns a boolean if a field has been set.

### GetIncludesLunch

`func (o *ItineraryEntryProperties) GetIncludesLunch() bool`

GetIncludesLunch returns the IncludesLunch field if non-nil, zero value otherwise.

### GetIncludesLunchOk

`func (o *ItineraryEntryProperties) GetIncludesLunchOk() (*bool, bool)`

GetIncludesLunchOk returns a tuple with the IncludesLunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesLunch

`func (o *ItineraryEntryProperties) SetIncludesLunch(v bool)`

SetIncludesLunch sets IncludesLunch field to given value.

### HasIncludesLunch

`func (o *ItineraryEntryProperties) HasIncludesLunch() bool`

HasIncludesLunch returns a boolean if a field has been set.

### GetIncludesDinner

`func (o *ItineraryEntryProperties) GetIncludesDinner() bool`

GetIncludesDinner returns the IncludesDinner field if non-nil, zero value otherwise.

### GetIncludesDinnerOk

`func (o *ItineraryEntryProperties) GetIncludesDinnerOk() (*bool, bool)`

GetIncludesDinnerOk returns a tuple with the IncludesDinner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesDinner

`func (o *ItineraryEntryProperties) SetIncludesDinner(v bool)`

SetIncludesDinner sets IncludesDinner field to given value.

### HasIncludesDinner

`func (o *ItineraryEntryProperties) HasIncludesDinner() bool`

HasIncludesDinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


