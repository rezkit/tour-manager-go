# ListHolidays200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of matching holidays | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number | [optional] 
**LastPage** | Pointer to **int32** | Total number of pages | [optional] 
**From** | Pointer to **int32** | Index of the first result on the current page. | [optional] 
**To** | Pointer to **int32** | Index of the last result on the current page. | [optional] 
**Data** | Pointer to [**[]Holiday**](Holiday.md) | Matching Holidays | [optional] 

## Methods

### NewListHolidays200Response

`func NewListHolidays200Response() *ListHolidays200Response`

NewListHolidays200Response instantiates a new ListHolidays200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHolidays200ResponseWithDefaults

`func NewListHolidays200ResponseWithDefaults() *ListHolidays200Response`

NewListHolidays200ResponseWithDefaults instantiates a new ListHolidays200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListHolidays200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListHolidays200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListHolidays200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListHolidays200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ListHolidays200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListHolidays200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListHolidays200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ListHolidays200Response) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetLastPage

`func (o *ListHolidays200Response) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ListHolidays200Response) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ListHolidays200Response) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *ListHolidays200Response) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetFrom

`func (o *ListHolidays200Response) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListHolidays200Response) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListHolidays200Response) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ListHolidays200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *ListHolidays200Response) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListHolidays200Response) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListHolidays200Response) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *ListHolidays200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetData

`func (o *ListHolidays200Response) GetData() []Holiday`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListHolidays200Response) GetDataOk() (*[]Holiday, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListHolidays200Response) SetData(v []Holiday)`

SetData sets Data field to given value.

### HasData

`func (o *ListHolidays200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


