# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**CreatedAt** | **time.Time** | Creation timestamp | 
**UpdatedAt** | **time.Time** | Timestamp of most recent update | 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Width** | **int32** |  | 
**Height** | **int32** |  | 
**FileName** | **string** |  | 
**FileSize** | **int32** |  | 
**Focus** | Pointer to [**FocusPoint**](FocusPoint.md) |  | [optional] 
**Thumbnail** | **string** |  | 

## Methods

### NewImage

`func NewImage(id string, createdAt time.Time, updatedAt time.Time, title string, width int32, height int32, fileName string, fileSize int32, thumbnail string, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Image) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Image) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Image) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Image) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Image) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Image) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTitle

`func (o *Image) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Image) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Image) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Image) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Image) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Image) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Image) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWidth

`func (o *Image) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Image) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Image) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *Image) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Image) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Image) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetFileName

`func (o *Image) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Image) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Image) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *Image) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *Image) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *Image) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetFocus

`func (o *Image) GetFocus() FocusPoint`

GetFocus returns the Focus field if non-nil, zero value otherwise.

### GetFocusOk

`func (o *Image) GetFocusOk() (*FocusPoint, bool)`

GetFocusOk returns a tuple with the Focus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocus

`func (o *Image) SetFocus(v FocusPoint)`

SetFocus sets Focus field to given value.

### HasFocus

`func (o *Image) HasFocus() bool`

HasFocus returns a boolean if a field has been set.

### GetThumbnail

`func (o *Image) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Image) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Image) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


