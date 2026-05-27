# ElementOptionPropertiesConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinAge** | Pointer to **int32** | Minimum age of all passengers at time of travel. (inclusive) | [optional] 
**MaxAge** | Pointer to **int32** | Maximum age of all passengers at time of travel (inclusive) | [optional] 
**PassengerSex** | Pointer to **[]string** | Permitted sex of all passengers. | [optional] 

## Methods

### NewElementOptionPropertiesConstraints

`func NewElementOptionPropertiesConstraints() *ElementOptionPropertiesConstraints`

NewElementOptionPropertiesConstraints instantiates a new ElementOptionPropertiesConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElementOptionPropertiesConstraintsWithDefaults

`func NewElementOptionPropertiesConstraintsWithDefaults() *ElementOptionPropertiesConstraints`

NewElementOptionPropertiesConstraintsWithDefaults instantiates a new ElementOptionPropertiesConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinAge

`func (o *ElementOptionPropertiesConstraints) GetMinAge() int32`

GetMinAge returns the MinAge field if non-nil, zero value otherwise.

### GetMinAgeOk

`func (o *ElementOptionPropertiesConstraints) GetMinAgeOk() (*int32, bool)`

GetMinAgeOk returns a tuple with the MinAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAge

`func (o *ElementOptionPropertiesConstraints) SetMinAge(v int32)`

SetMinAge sets MinAge field to given value.

### HasMinAge

`func (o *ElementOptionPropertiesConstraints) HasMinAge() bool`

HasMinAge returns a boolean if a field has been set.

### GetMaxAge

`func (o *ElementOptionPropertiesConstraints) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *ElementOptionPropertiesConstraints) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *ElementOptionPropertiesConstraints) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *ElementOptionPropertiesConstraints) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetPassengerSex

`func (o *ElementOptionPropertiesConstraints) GetPassengerSex() []string`

GetPassengerSex returns the PassengerSex field if non-nil, zero value otherwise.

### GetPassengerSexOk

`func (o *ElementOptionPropertiesConstraints) GetPassengerSexOk() (*[]string, bool)`

GetPassengerSexOk returns a tuple with the PassengerSex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengerSex

`func (o *ElementOptionPropertiesConstraints) SetPassengerSex(v []string)`

SetPassengerSex sets PassengerSex field to given value.

### HasPassengerSex

`func (o *ElementOptionPropertiesConstraints) HasPassengerSex() bool`

HasPassengerSex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


