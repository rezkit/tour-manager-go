# ItineraryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**VersionId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of most recent update | [optional] 
**StartDay** | Pointer to **int32** | Day index when this entry starts | [optional] 
**EndDay** | Pointer to **int32** | Day index when this entry ends | [optional] 
**Title** | Pointer to **string** | Itinerary Entry Title | [optional] 
**Introduction** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IncludesBreakfast** | Pointer to **bool** | Determines if a complementary breakfast meal is included | [optional] 
**IncludesLunch** | Pointer to **bool** | Determines if a complementary lunch meal is included | [optional] 
**IncludesDinner** | Pointer to **bool** | Determines if a complementary dinner meal is included | [optional] 
**Published** | Pointer to **bool** |  | [optional] 

## Methods

### NewItineraryEntry

`func NewItineraryEntry() *ItineraryEntry`

NewItineraryEntry instantiates a new ItineraryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItineraryEntryWithDefaults

`func NewItineraryEntryWithDefaults() *ItineraryEntry`

NewItineraryEntryWithDefaults instantiates a new ItineraryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItineraryEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItineraryEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItineraryEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItineraryEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionId

`func (o *ItineraryEntry) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ItineraryEntry) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ItineraryEntry) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ItineraryEntry) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ItineraryEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ItineraryEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ItineraryEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ItineraryEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ItineraryEntry) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ItineraryEntry) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ItineraryEntry) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ItineraryEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStartDay

`func (o *ItineraryEntry) GetStartDay() int32`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *ItineraryEntry) GetStartDayOk() (*int32, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *ItineraryEntry) SetStartDay(v int32)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *ItineraryEntry) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### GetEndDay

`func (o *ItineraryEntry) GetEndDay() int32`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *ItineraryEntry) GetEndDayOk() (*int32, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *ItineraryEntry) SetEndDay(v int32)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *ItineraryEntry) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### GetTitle

`func (o *ItineraryEntry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItineraryEntry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItineraryEntry) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItineraryEntry) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIntroduction

`func (o *ItineraryEntry) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *ItineraryEntry) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *ItineraryEntry) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *ItineraryEntry) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetDescription

`func (o *ItineraryEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItineraryEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItineraryEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItineraryEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludesBreakfast

`func (o *ItineraryEntry) GetIncludesBreakfast() bool`

GetIncludesBreakfast returns the IncludesBreakfast field if non-nil, zero value otherwise.

### GetIncludesBreakfastOk

`func (o *ItineraryEntry) GetIncludesBreakfastOk() (*bool, bool)`

GetIncludesBreakfastOk returns a tuple with the IncludesBreakfast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesBreakfast

`func (o *ItineraryEntry) SetIncludesBreakfast(v bool)`

SetIncludesBreakfast sets IncludesBreakfast field to given value.

### HasIncludesBreakfast

`func (o *ItineraryEntry) HasIncludesBreakfast() bool`

HasIncludesBreakfast returns a boolean if a field has been set.

### GetIncludesLunch

`func (o *ItineraryEntry) GetIncludesLunch() bool`

GetIncludesLunch returns the IncludesLunch field if non-nil, zero value otherwise.

### GetIncludesLunchOk

`func (o *ItineraryEntry) GetIncludesLunchOk() (*bool, bool)`

GetIncludesLunchOk returns a tuple with the IncludesLunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesLunch

`func (o *ItineraryEntry) SetIncludesLunch(v bool)`

SetIncludesLunch sets IncludesLunch field to given value.

### HasIncludesLunch

`func (o *ItineraryEntry) HasIncludesLunch() bool`

HasIncludesLunch returns a boolean if a field has been set.

### GetIncludesDinner

`func (o *ItineraryEntry) GetIncludesDinner() bool`

GetIncludesDinner returns the IncludesDinner field if non-nil, zero value otherwise.

### GetIncludesDinnerOk

`func (o *ItineraryEntry) GetIncludesDinnerOk() (*bool, bool)`

GetIncludesDinnerOk returns a tuple with the IncludesDinner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesDinner

`func (o *ItineraryEntry) SetIncludesDinner(v bool)`

SetIncludesDinner sets IncludesDinner field to given value.

### HasIncludesDinner

`func (o *ItineraryEntry) HasIncludesDinner() bool`

HasIncludesDinner returns a boolean if a field has been set.

### GetPublished

`func (o *ItineraryEntry) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ItineraryEntry) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ItineraryEntry) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ItineraryEntry) HasPublished() bool`

HasPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


