# PriceDeposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calculated** | Pointer to **bool** | Determines if the given deposit amount has been calculated via a global rule | [optional] 
**Value** | Pointer to **float32** | Deposit payment required to confirm a reservation. | [optional] 

## Methods

### NewPriceDeposit

`func NewPriceDeposit() *PriceDeposit`

NewPriceDeposit instantiates a new PriceDeposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDepositWithDefaults

`func NewPriceDepositWithDefaults() *PriceDeposit`

NewPriceDepositWithDefaults instantiates a new PriceDeposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculated

`func (o *PriceDeposit) GetCalculated() bool`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *PriceDeposit) GetCalculatedOk() (*bool, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *PriceDeposit) SetCalculated(v bool)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *PriceDeposit) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### GetValue

`func (o *PriceDeposit) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PriceDeposit) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PriceDeposit) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *PriceDeposit) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


