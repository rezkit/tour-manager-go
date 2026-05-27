# UpdateOrganizationSettingsRequestDepositDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **float32** | Default percentage rate to apply as a deposit to all prices with no deposit value set | 
**BalanceDue** | **float32** | Number of days before departure that a balance is due by default unless a specific date is set. | 

## Methods

### NewUpdateOrganizationSettingsRequestDepositDefaults

`func NewUpdateOrganizationSettingsRequestDepositDefaults(percentage float32, balanceDue float32, ) *UpdateOrganizationSettingsRequestDepositDefaults`

NewUpdateOrganizationSettingsRequestDepositDefaults instantiates a new UpdateOrganizationSettingsRequestDepositDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationSettingsRequestDepositDefaultsWithDefaults

`func NewUpdateOrganizationSettingsRequestDepositDefaultsWithDefaults() *UpdateOrganizationSettingsRequestDepositDefaults`

NewUpdateOrganizationSettingsRequestDepositDefaultsWithDefaults instantiates a new UpdateOrganizationSettingsRequestDepositDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *UpdateOrganizationSettingsRequestDepositDefaults) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *UpdateOrganizationSettingsRequestDepositDefaults) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *UpdateOrganizationSettingsRequestDepositDefaults) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetBalanceDue

`func (o *UpdateOrganizationSettingsRequestDepositDefaults) GetBalanceDue() float32`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *UpdateOrganizationSettingsRequestDepositDefaults) GetBalanceDueOk() (*float32, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *UpdateOrganizationSettingsRequestDepositDefaults) SetBalanceDue(v float32)`

SetBalanceDue sets BalanceDue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


