# DepartureElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of most recent update | [optional] 
**DeletedAt** | Pointer to **NullableTime** | Timestamp of the item&#39;s deletion. Null if the item is not deleted. | [optional] 
**Inventory** | Pointer to [**Inventory**](Inventory.md) |  | [optional] 
**Element** | Pointer to [**DepartureElementElement**](DepartureElementElement.md) |  | [optional] 
**BalanceDue** | Pointer to [**DepartureElementBalanceDue**](DepartureElementBalanceDue.md) |  | [optional] 
**Options** | Pointer to [**[]DepartureElementOptionsInner**](DepartureElementOptionsInner.md) |  | [optional] 

## Methods

### NewDepartureElement

`func NewDepartureElement() *DepartureElement`

NewDepartureElement instantiates a new DepartureElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartureElementWithDefaults

`func NewDepartureElementWithDefaults() *DepartureElement`

NewDepartureElementWithDefaults instantiates a new DepartureElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DepartureElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepartureElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepartureElement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DepartureElement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DepartureElement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DepartureElement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DepartureElement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DepartureElement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DepartureElement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DepartureElement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DepartureElement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DepartureElement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DepartureElement) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DepartureElement) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DepartureElement) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DepartureElement) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DepartureElement) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DepartureElement) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetInventory

`func (o *DepartureElement) GetInventory() Inventory`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *DepartureElement) GetInventoryOk() (*Inventory, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *DepartureElement) SetInventory(v Inventory)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *DepartureElement) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetElement

`func (o *DepartureElement) GetElement() DepartureElementElement`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *DepartureElement) GetElementOk() (*DepartureElementElement, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *DepartureElement) SetElement(v DepartureElementElement)`

SetElement sets Element field to given value.

### HasElement

`func (o *DepartureElement) HasElement() bool`

HasElement returns a boolean if a field has been set.

### GetBalanceDue

`func (o *DepartureElement) GetBalanceDue() DepartureElementBalanceDue`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *DepartureElement) GetBalanceDueOk() (*DepartureElementBalanceDue, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *DepartureElement) SetBalanceDue(v DepartureElementBalanceDue)`

SetBalanceDue sets BalanceDue field to given value.

### HasBalanceDue

`func (o *DepartureElement) HasBalanceDue() bool`

HasBalanceDue returns a boolean if a field has been set.

### GetOptions

`func (o *DepartureElement) GetOptions() []DepartureElementOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DepartureElement) GetOptionsOk() (*[]DepartureElementOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DepartureElement) SetOptions(v []DepartureElementOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DepartureElement) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


