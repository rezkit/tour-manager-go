# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier. Unique system-wide to a single entity. Consists of 26 alphanumeric characters.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp of most recent update | [optional] 
**Currency** | Pointer to **string** | ISO Currency Code | [optional] 
**OnSale** | Pointer to **bool** | Determines if this price is available for sale. Configurations that are not possible or disallowed are disabled.  | [optional] 
**Initialized** | Pointer to **interface{}** | Determines if the price has been initialized with a value.  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Deposit** | Pointer to [**PriceDeposit**](PriceDeposit.md) |  | [optional] 

## Methods

### NewPrice

`func NewPrice() *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Price) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Price) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Price) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Price) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Price) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Price) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Price) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Price) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Price) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Price) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Price) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Price) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Price) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOnSale

`func (o *Price) GetOnSale() bool`

GetOnSale returns the OnSale field if non-nil, zero value otherwise.

### GetOnSaleOk

`func (o *Price) GetOnSaleOk() (*bool, bool)`

GetOnSaleOk returns a tuple with the OnSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSale

`func (o *Price) SetOnSale(v bool)`

SetOnSale sets OnSale field to given value.

### HasOnSale

`func (o *Price) HasOnSale() bool`

HasOnSale returns a boolean if a field has been set.

### GetInitialized

`func (o *Price) GetInitialized() interface{}`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *Price) GetInitializedOk() (*interface{}, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *Price) SetInitialized(v interface{})`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *Price) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### SetInitializedNil

`func (o *Price) SetInitializedNil(b bool)`

 SetInitializedNil sets the value for Initialized to be an explicit nil

### UnsetInitialized
`func (o *Price) UnsetInitialized()`

UnsetInitialized ensures that no value is present for Initialized, not even an explicit nil
### GetValue

`func (o *Price) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Price) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Price) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Price) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDeposit

`func (o *Price) GetDeposit() PriceDeposit`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *Price) GetDepositOk() (*PriceDeposit, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *Price) SetDeposit(v PriceDeposit)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *Price) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


