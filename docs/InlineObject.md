# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of matching holidays | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number | [optional] 
**LastPage** | Pointer to **int32** | Total number of pages | [optional] 
**From** | Pointer to **int32** | Index of the first result on the current page. | [optional] 
**To** | Pointer to **int32** | Index of the last result on the current page. | [optional] 
**Data** | Pointer to [**[]Category**](Category.md) | Matching Categories | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineObject) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineObject) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineObject) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineObject) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCurrentPage

`func (o *InlineObject) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *InlineObject) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *InlineObject) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *InlineObject) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetLastPage

`func (o *InlineObject) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *InlineObject) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *InlineObject) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *InlineObject) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetFrom

`func (o *InlineObject) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InlineObject) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InlineObject) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *InlineObject) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *InlineObject) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InlineObject) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InlineObject) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *InlineObject) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetData

`func (o *InlineObject) GetData() []Category`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InlineObject) GetDataOk() (*[]Category, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InlineObject) SetData(v []Category)`

SetData sets Data field to given value.

### HasData

`func (o *InlineObject) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


