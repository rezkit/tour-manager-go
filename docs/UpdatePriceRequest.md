# UpdatePriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float32** | Set the price value | [optional] 
**Deposit** | Pointer to **NullableFloat32** | Set the deposit value. Remove the deposit and revert to default calculation by setting to null. | [optional] 
**OnSale** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdatePriceRequest

`func NewUpdatePriceRequest() *UpdatePriceRequest`

NewUpdatePriceRequest instantiates a new UpdatePriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceRequestWithDefaults

`func NewUpdatePriceRequestWithDefaults() *UpdatePriceRequest`

NewUpdatePriceRequestWithDefaults instantiates a new UpdatePriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdatePriceRequest) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdatePriceRequest) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdatePriceRequest) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdatePriceRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDeposit

`func (o *UpdatePriceRequest) GetDeposit() float32`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *UpdatePriceRequest) GetDepositOk() (*float32, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *UpdatePriceRequest) SetDeposit(v float32)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *UpdatePriceRequest) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### SetDepositNil

`func (o *UpdatePriceRequest) SetDepositNil(b bool)`

 SetDepositNil sets the value for Deposit to be an explicit nil

### UnsetDeposit
`func (o *UpdatePriceRequest) UnsetDeposit()`

UnsetDeposit ensures that no value is present for Deposit, not even an explicit nil
### GetOnSale

`func (o *UpdatePriceRequest) GetOnSale() bool`

GetOnSale returns the OnSale field if non-nil, zero value otherwise.

### GetOnSaleOk

`func (o *UpdatePriceRequest) GetOnSaleOk() (*bool, bool)`

GetOnSaleOk returns a tuple with the OnSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSale

`func (o *UpdatePriceRequest) SetOnSale(v bool)`

SetOnSale sets OnSale field to given value.

### HasOnSale

`func (o *UpdatePriceRequest) HasOnSale() bool`

HasOnSale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


