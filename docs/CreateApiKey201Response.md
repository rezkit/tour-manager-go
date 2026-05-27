# CreateApiKey201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlainTextToken** | Pointer to **string** | The API Key itself.  **Warning: This is the _only_ time the key will be shown. If the key is lost it _cannot_ be retrieved.**  | [optional] 
**AccessToken** | Pointer to [**ApiKey**](ApiKey.md) |  | [optional] 

## Methods

### NewCreateApiKey201Response

`func NewCreateApiKey201Response() *CreateApiKey201Response`

NewCreateApiKey201Response instantiates a new CreateApiKey201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKey201ResponseWithDefaults

`func NewCreateApiKey201ResponseWithDefaults() *CreateApiKey201Response`

NewCreateApiKey201ResponseWithDefaults instantiates a new CreateApiKey201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlainTextToken

`func (o *CreateApiKey201Response) GetPlainTextToken() string`

GetPlainTextToken returns the PlainTextToken field if non-nil, zero value otherwise.

### GetPlainTextTokenOk

`func (o *CreateApiKey201Response) GetPlainTextTokenOk() (*string, bool)`

GetPlainTextTokenOk returns a tuple with the PlainTextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainTextToken

`func (o *CreateApiKey201Response) SetPlainTextToken(v string)`

SetPlainTextToken sets PlainTextToken field to given value.

### HasPlainTextToken

`func (o *CreateApiKey201Response) HasPlainTextToken() bool`

HasPlainTextToken returns a boolean if a field has been set.

### GetAccessToken

`func (o *CreateApiKey201Response) GetAccessToken() ApiKey`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *CreateApiKey201Response) GetAccessTokenOk() (*ApiKey, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *CreateApiKey201Response) SetAccessToken(v ApiKey)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *CreateApiKey201Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


