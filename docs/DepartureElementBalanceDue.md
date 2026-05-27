# DepartureElementBalanceDue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calculated** | Pointer to **bool** | Determines if the given balance due date was calculated or has been explicitly set. | [optional] 
**Date** | Pointer to **time.Time** | Date by which the full balance for the element must be paid. | [optional] 

## Methods

### NewDepartureElementBalanceDue

`func NewDepartureElementBalanceDue() *DepartureElementBalanceDue`

NewDepartureElementBalanceDue instantiates a new DepartureElementBalanceDue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartureElementBalanceDueWithDefaults

`func NewDepartureElementBalanceDueWithDefaults() *DepartureElementBalanceDue`

NewDepartureElementBalanceDueWithDefaults instantiates a new DepartureElementBalanceDue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculated

`func (o *DepartureElementBalanceDue) GetCalculated() bool`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *DepartureElementBalanceDue) GetCalculatedOk() (*bool, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculated

`func (o *DepartureElementBalanceDue) SetCalculated(v bool)`

SetCalculated sets Calculated field to given value.

### HasCalculated

`func (o *DepartureElementBalanceDue) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### GetDate

`func (o *DepartureElementBalanceDue) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DepartureElementBalanceDue) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DepartureElementBalanceDue) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *DepartureElementBalanceDue) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


