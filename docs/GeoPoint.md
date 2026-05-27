# GeoPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lat** | **float32** |  | 
**Lng** | **float32** |  | 
**Zoom** | Pointer to **int32** |  | [optional] 
**Precision** | Pointer to **float32** |  | [optional] 

## Methods

### NewGeoPoint

`func NewGeoPoint(lat float32, lng float32, ) *GeoPoint`

NewGeoPoint instantiates a new GeoPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoPointWithDefaults

`func NewGeoPointWithDefaults() *GeoPoint`

NewGeoPointWithDefaults instantiates a new GeoPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLat

`func (o *GeoPoint) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *GeoPoint) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *GeoPoint) SetLat(v float32)`

SetLat sets Lat field to given value.


### GetLng

`func (o *GeoPoint) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *GeoPoint) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *GeoPoint) SetLng(v float32)`

SetLng sets Lng field to given value.


### GetZoom

`func (o *GeoPoint) GetZoom() int32`

GetZoom returns the Zoom field if non-nil, zero value otherwise.

### GetZoomOk

`func (o *GeoPoint) GetZoomOk() (*int32, bool)`

GetZoomOk returns a tuple with the Zoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoom

`func (o *GeoPoint) SetZoom(v int32)`

SetZoom sets Zoom field to given value.

### HasZoom

`func (o *GeoPoint) HasZoom() bool`

HasZoom returns a boolean if a field has been set.

### GetPrecision

`func (o *GeoPoint) GetPrecision() float32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *GeoPoint) GetPrecisionOk() (*float32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *GeoPoint) SetPrecision(v float32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *GeoPoint) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


