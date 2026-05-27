# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**RezkitId** | Pointer to **string** | Organization&#39;s RezKit ID | [optional] 
**Name** | Pointer to **string** | Organization Name | [optional] 
**Currencies** | Pointer to **[]string** | List of presently available currencies.  | [optional] 
**DepositDefaults** | Pointer to [**OrganizationDepositDefaults**](OrganizationDepositDefaults.md) |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRezkitId

`func (o *Organization) GetRezkitId() string`

GetRezkitId returns the RezkitId field if non-nil, zero value otherwise.

### GetRezkitIdOk

`func (o *Organization) GetRezkitIdOk() (*string, bool)`

GetRezkitIdOk returns a tuple with the RezkitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRezkitId

`func (o *Organization) SetRezkitId(v string)`

SetRezkitId sets RezkitId field to given value.

### HasRezkitId

`func (o *Organization) HasRezkitId() bool`

HasRezkitId returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrencies

`func (o *Organization) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *Organization) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *Organization) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *Organization) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetDepositDefaults

`func (o *Organization) GetDepositDefaults() OrganizationDepositDefaults`

GetDepositDefaults returns the DepositDefaults field if non-nil, zero value otherwise.

### GetDepositDefaultsOk

`func (o *Organization) GetDepositDefaultsOk() (*OrganizationDepositDefaults, bool)`

GetDepositDefaultsOk returns a tuple with the DepositDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositDefaults

`func (o *Organization) SetDepositDefaults(v OrganizationDepositDefaults)`

SetDepositDefaults sets DepositDefaults field to given value.

### HasDepositDefaults

`func (o *Organization) HasDepositDefaults() bool`

HasDepositDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


