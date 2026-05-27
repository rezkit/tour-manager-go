# DepartureElementOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of most recent update | [optional] 
**DeletedAt** | Pointer to **NullableTime** | Timestamp of the item&#39;s deletion. Null if the item is not deleted. | [optional] 
**Category** | Pointer to [**Category**](Category.md) |  | [optional] 
**Published** | Pointer to **bool** | Element option is published and viewable | [optional] 
**WebBookable** | Pointer to **bool** | Determines if the element option can be booked online. If false then the element can only be booked directly.  | [optional] 
**IsLead** | Pointer to **bool** | Determines if the prices for this option are valid candidates for a leading price. | [optional] 
**Name** | Pointer to **string** | Option Name | [optional] 
**PriceUnit** | Pointer to [**PriceUnit**](PriceUnit.md) |  | [optional] 
**Occupancy** | Pointer to [**Occupancy**](Occupancy.md) |  | [optional] 
**Constraints** | Pointer to [**ElementOptionPropertiesConstraints**](ElementOptionPropertiesConstraints.md) |  | [optional] 
**Prices** | Pointer to [**[]Price**](Price.md) | Prices for this element-option on this departure, one per operator currency. Exactly one item will always exist in this list for every operator currency.  | [optional] 

## Methods

### NewDepartureElementOptionsInner

`func NewDepartureElementOptionsInner() *DepartureElementOptionsInner`

NewDepartureElementOptionsInner instantiates a new DepartureElementOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartureElementOptionsInnerWithDefaults

`func NewDepartureElementOptionsInnerWithDefaults() *DepartureElementOptionsInner`

NewDepartureElementOptionsInnerWithDefaults instantiates a new DepartureElementOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DepartureElementOptionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepartureElementOptionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepartureElementOptionsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DepartureElementOptionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DepartureElementOptionsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DepartureElementOptionsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DepartureElementOptionsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DepartureElementOptionsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DepartureElementOptionsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DepartureElementOptionsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DepartureElementOptionsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DepartureElementOptionsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DepartureElementOptionsInner) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DepartureElementOptionsInner) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DepartureElementOptionsInner) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DepartureElementOptionsInner) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DepartureElementOptionsInner) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DepartureElementOptionsInner) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetCategory

`func (o *DepartureElementOptionsInner) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DepartureElementOptionsInner) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DepartureElementOptionsInner) SetCategory(v Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DepartureElementOptionsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPublished

`func (o *DepartureElementOptionsInner) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *DepartureElementOptionsInner) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *DepartureElementOptionsInner) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *DepartureElementOptionsInner) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetWebBookable

`func (o *DepartureElementOptionsInner) GetWebBookable() bool`

GetWebBookable returns the WebBookable field if non-nil, zero value otherwise.

### GetWebBookableOk

`func (o *DepartureElementOptionsInner) GetWebBookableOk() (*bool, bool)`

GetWebBookableOk returns a tuple with the WebBookable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebBookable

`func (o *DepartureElementOptionsInner) SetWebBookable(v bool)`

SetWebBookable sets WebBookable field to given value.

### HasWebBookable

`func (o *DepartureElementOptionsInner) HasWebBookable() bool`

HasWebBookable returns a boolean if a field has been set.

### GetIsLead

`func (o *DepartureElementOptionsInner) GetIsLead() bool`

GetIsLead returns the IsLead field if non-nil, zero value otherwise.

### GetIsLeadOk

`func (o *DepartureElementOptionsInner) GetIsLeadOk() (*bool, bool)`

GetIsLeadOk returns a tuple with the IsLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLead

`func (o *DepartureElementOptionsInner) SetIsLead(v bool)`

SetIsLead sets IsLead field to given value.

### HasIsLead

`func (o *DepartureElementOptionsInner) HasIsLead() bool`

HasIsLead returns a boolean if a field has been set.

### GetName

`func (o *DepartureElementOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepartureElementOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepartureElementOptionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DepartureElementOptionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceUnit

`func (o *DepartureElementOptionsInner) GetPriceUnit() PriceUnit`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *DepartureElementOptionsInner) GetPriceUnitOk() (*PriceUnit, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *DepartureElementOptionsInner) SetPriceUnit(v PriceUnit)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *DepartureElementOptionsInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetOccupancy

`func (o *DepartureElementOptionsInner) GetOccupancy() Occupancy`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *DepartureElementOptionsInner) GetOccupancyOk() (*Occupancy, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *DepartureElementOptionsInner) SetOccupancy(v Occupancy)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *DepartureElementOptionsInner) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.

### GetConstraints

`func (o *DepartureElementOptionsInner) GetConstraints() ElementOptionPropertiesConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *DepartureElementOptionsInner) GetConstraintsOk() (*ElementOptionPropertiesConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *DepartureElementOptionsInner) SetConstraints(v ElementOptionPropertiesConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *DepartureElementOptionsInner) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetPrices

`func (o *DepartureElementOptionsInner) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *DepartureElementOptionsInner) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *DepartureElementOptionsInner) SetPrices(v []Price)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *DepartureElementOptionsInner) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


