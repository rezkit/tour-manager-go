# ElementOptionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Occupancy** | Pointer to [**Occupancy**](Occupancy.md) |  | [optional] 
**PriceUnit** | Pointer to [**PriceUnit**](PriceUnit.md) |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] [default to true]
**WebBookable** | Pointer to **bool** |  | [optional] [default to true]
**IsLead** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewElementOptionParams

`func NewElementOptionParams() *ElementOptionParams`

NewElementOptionParams instantiates a new ElementOptionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElementOptionParamsWithDefaults

`func NewElementOptionParamsWithDefaults() *ElementOptionParams`

NewElementOptionParamsWithDefaults instantiates a new ElementOptionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *ElementOptionParams) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ElementOptionParams) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ElementOptionParams) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ElementOptionParams) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetName

`func (o *ElementOptionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ElementOptionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ElementOptionParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ElementOptionParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOccupancy

`func (o *ElementOptionParams) GetOccupancy() Occupancy`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *ElementOptionParams) GetOccupancyOk() (*Occupancy, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *ElementOptionParams) SetOccupancy(v Occupancy)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *ElementOptionParams) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ElementOptionParams) GetPriceUnit() PriceUnit`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ElementOptionParams) GetPriceUnitOk() (*PriceUnit, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ElementOptionParams) SetPriceUnit(v PriceUnit)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ElementOptionParams) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetPublished

`func (o *ElementOptionParams) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ElementOptionParams) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ElementOptionParams) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ElementOptionParams) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetWebBookable

`func (o *ElementOptionParams) GetWebBookable() bool`

GetWebBookable returns the WebBookable field if non-nil, zero value otherwise.

### GetWebBookableOk

`func (o *ElementOptionParams) GetWebBookableOk() (*bool, bool)`

GetWebBookableOk returns a tuple with the WebBookable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBookable

`func (o *ElementOptionParams) SetWebBookable(v bool)`

SetWebBookable sets WebBookable field to given value.

### HasWebBookable

`func (o *ElementOptionParams) HasWebBookable() bool`

HasWebBookable returns a boolean if a field has been set.

### GetIsLead

`func (o *ElementOptionParams) GetIsLead() bool`

GetIsLead returns the IsLead field if non-nil, zero value otherwise.

### GetIsLeadOk

`func (o *ElementOptionParams) GetIsLeadOk() (*bool, bool)`

GetIsLeadOk returns a tuple with the IsLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLead

`func (o *ElementOptionParams) SetIsLead(v bool)`

SetIsLead sets IsLead field to given value.

### HasIsLead

`func (o *ElementOptionParams) HasIsLead() bool`

HasIsLead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


