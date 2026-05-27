# ElementOptionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | 
**Name** | **string** |  | 
**Occupancy** | [**Occupancy**](Occupancy.md) |  | 
**PriceUnit** | [**PriceUnit**](PriceUnit.md) |  | 
**Published** | Pointer to **bool** |  | [optional] [default to true]
**WebBookable** | Pointer to **bool** |  | [optional] [default to true]
**IsLead** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewElementOptionCreate

`func NewElementOptionCreate(categoryId string, name string, occupancy Occupancy, priceUnit PriceUnit, ) *ElementOptionCreate`

NewElementOptionCreate instantiates a new ElementOptionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElementOptionCreateWithDefaults

`func NewElementOptionCreateWithDefaults() *ElementOptionCreate`

NewElementOptionCreateWithDefaults instantiates a new ElementOptionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *ElementOptionCreate) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ElementOptionCreate) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ElementOptionCreate) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetName

`func (o *ElementOptionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ElementOptionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ElementOptionCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOccupancy

`func (o *ElementOptionCreate) GetOccupancy() Occupancy`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *ElementOptionCreate) GetOccupancyOk() (*Occupancy, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *ElementOptionCreate) SetOccupancy(v Occupancy)`

SetOccupancy sets Occupancy field to given value.


### GetPriceUnit

`func (o *ElementOptionCreate) GetPriceUnit() PriceUnit`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ElementOptionCreate) GetPriceUnitOk() (*PriceUnit, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ElementOptionCreate) SetPriceUnit(v PriceUnit)`

SetPriceUnit sets PriceUnit field to given value.


### GetPublished

`func (o *ElementOptionCreate) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ElementOptionCreate) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ElementOptionCreate) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ElementOptionCreate) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetWebBookable

`func (o *ElementOptionCreate) GetWebBookable() bool`

GetWebBookable returns the WebBookable field if non-nil, zero value otherwise.

### GetWebBookableOk

`func (o *ElementOptionCreate) GetWebBookableOk() (*bool, bool)`

GetWebBookableOk returns a tuple with the WebBookable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBookable

`func (o *ElementOptionCreate) SetWebBookable(v bool)`

SetWebBookable sets WebBookable field to given value.

### HasWebBookable

`func (o *ElementOptionCreate) HasWebBookable() bool`

HasWebBookable returns a boolean if a field has been set.

### GetIsLead

`func (o *ElementOptionCreate) GetIsLead() bool`

GetIsLead returns the IsLead field if non-nil, zero value otherwise.

### GetIsLeadOk

`func (o *ElementOptionCreate) GetIsLeadOk() (*bool, bool)`

GetIsLeadOk returns a tuple with the IsLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLead

`func (o *ElementOptionCreate) SetIsLead(v bool)`

SetIsLead sets IsLead field to given value.

### HasIsLead

`func (o *ElementOptionCreate) HasIsLead() bool`

HasIsLead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


