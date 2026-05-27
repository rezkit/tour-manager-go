# OrganizationDepositDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | Pointer to **float32** | Default percentage rate to apply as a deposit to all prices with no deposit value set | [optional] 
**BalanceDue** | Pointer to **int32** | Number of days before departure that a balance is due by default unless a specific date is set. | [optional] 

## Methods

### NewOrganizationDepositDefaults

`func NewOrganizationDepositDefaults() *OrganizationDepositDefaults`

NewOrganizationDepositDefaults instantiates a new OrganizationDepositDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDepositDefaultsWithDefaults

`func NewOrganizationDepositDefaultsWithDefaults() *OrganizationDepositDefaults`

NewOrganizationDepositDefaultsWithDefaults instantiates a new OrganizationDepositDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *OrganizationDepositDefaults) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *OrganizationDepositDefaults) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *OrganizationDepositDefaults) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *OrganizationDepositDefaults) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetBalanceDue

`func (o *OrganizationDepositDefaults) GetBalanceDue() int32`

GetBalanceDue returns the BalanceDue field if non-nil, zero value otherwise.

### GetBalanceDueOk

`func (o *OrganizationDepositDefaults) GetBalanceDueOk() (*int32, bool)`

GetBalanceDueOk returns a tuple with the BalanceDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDue

`func (o *OrganizationDepositDefaults) SetBalanceDue(v int32)`

SetBalanceDue sets BalanceDue field to given value.

### HasBalanceDue

`func (o *OrganizationDepositDefaults) HasBalanceDue() bool`

HasBalanceDue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


