# Element

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of most recent update | [optional] 
**DeletedAt** | Pointer to **NullableTime** | Timestamp of the item&#39;s deletion. Null if the item is not deleted. | [optional] 
**VersionId** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**Category** | Pointer to [**Category**](Category.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsPackage** | Pointer to **bool** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**DefaultInventory** | Pointer to [**Inventory**](Inventory.md) |  | [optional] 
**BalanceDue** | Pointer to **NullableInt32** | Rule for balance payment due relative to departure date. No balance rule specified, global default applies. | [optional] 
**Options** | Pointer to [**[]ElementOptionProperties**](ElementOptionProperties.md) | Available options for the element. All options for an element share the same inventory.  | [optional] 

## Methods

### NewElement

`func NewElement() *Element`

NewElement instantiates a new Element object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElementWithDefaults

`func NewElementWithDefaults() *Element`

NewElementWithDefaults instantiates a new Element object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Element) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Element) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Element) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Element) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Element) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Element) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Element) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Element) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Element) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Element) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Element) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Element) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Element) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Element) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Element) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Element) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *Element) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *Element) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetVersionId

`func (o *Element) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Element) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Element) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *Element) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetCategory

`func (o *Element) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Element) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Element) SetCategory(v Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Element) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetName

`func (o *Element) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Element) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Element) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Element) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsPackage

`func (o *Element) GetIsPackage() bool`

GetIsPackage returns the IsPackage field if non-nil, zero value otherwise.

### GetIsPackageOk

`func (o *Element) GetIsPackageOk() (*bool, bool)`

GetIsPackageOk returns a tuple with the IsPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPackage

`func (o *Element) SetIsPackage(v bool)`

SetIsPackage sets IsPackage field to given value.

### HasIsPackage

`func (o *Element) HasIsPackage() bool`

HasIsPackage returns a boolean if a field has been set.

### GetPublished

`func (o *Element) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *Element) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *Element) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *Element) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetDefaultInventory

`func (o *Element) GetDefaultInventory() Inventory`

GetDefaultInventory returns the DefaultInventory field if non-nil, zero value otherwise.

### GetDefaultInventoryOk

`func (o *Element) GetDefaultInventoryOk() (*Inventory, bool)`

GetDefaultInventoryOk returns a tuple with the DefaultInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInventory

`func (o *Element) SetDefaultInventory(v Inventory)`

SetDefaultInventory sets DefaultInventory field to given value.

### HasDefaultInventory

`func (o *Element) HasDefaultInventory() bool`

HasDefaultInventory returns a boolean if a field has been set.

### GetBalanceDue

`func (o *Element) GetBalanceDue() int32`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *Element) GetBalanceDueOk() (*int32, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *Element) SetBalanceDue(v int32)`

SetBalanceDue sets BalanceDue field to given value.

### HasBalanceDue

`func (o *Element) HasBalanceDue() bool`

HasBalanceDue returns a boolean if a field has been set.

### SetBalanceDueNil

`func (o *Element) SetBalanceDueNil(b bool)`

 SetBalanceDueNil sets the value for BalanceDue to be an explicit nil

### UnsetBalanceDue
`func (o *Element) UnsetBalanceDue()`

UnsetBalanceDue ensures that no value is present for BalanceDue, not even an explicit nil
### GetOptions

`func (o *Element) GetOptions() []ElementOptionProperties`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Element) GetOptionsOk() (*[]ElementOptionProperties, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Element) SetOptions(v []ElementOptionProperties)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Element) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


