# UpdateOrganizationSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currencies** | Pointer to **[]string** | List of currencies to provide pricing in. If a new value is added to the list then uninitialized prices will be generated for all values. If a value is removed, then existing prices will be retained but no new prices will be generated.  | [optional] 
**DepositDefaults** | Pointer to [**UpdateOrganizationSettingsRequestDepositDefaults**](UpdateOrganizationSettingsRequestDepositDefaults.md) |  | [optional] 

## Methods

### NewUpdateOrganizationSettingsRequest

`func NewUpdateOrganizationSettingsRequest() *UpdateOrganizationSettingsRequest`

NewUpdateOrganizationSettingsRequest instantiates a new UpdateOrganizationSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationSettingsRequestWithDefaults

`func NewUpdateOrganizationSettingsRequestWithDefaults() *UpdateOrganizationSettingsRequest`

NewUpdateOrganizationSettingsRequestWithDefaults instantiates a new UpdateOrganizationSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencies

`func (o *UpdateOrganizationSettingsRequest) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *UpdateOrganizationSettingsRequest) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *UpdateOrganizationSettingsRequest) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *UpdateOrganizationSettingsRequest) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetDepositDefaults

`func (o *UpdateOrganizationSettingsRequest) GetDepositDefaults() UpdateOrganizationSettingsRequestDepositDefaults`

GetDepositDefaults returns the DepositDefaults field if non-nil, zero value otherwise.

### GetDepositDefaultsOk

`func (o *UpdateOrganizationSettingsRequest) GetDepositDefaultsOk() (*UpdateOrganizationSettingsRequestDepositDefaults, bool)`

GetDepositDefaultsOk returns a tuple with the DepositDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositDefaults

`func (o *UpdateOrganizationSettingsRequest) SetDepositDefaults(v UpdateOrganizationSettingsRequestDepositDefaults)`

SetDepositDefaults sets DepositDefaults field to given value.

### HasDepositDefaults

`func (o *UpdateOrganizationSettingsRequest) HasDepositDefaults() bool`

HasDepositDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


