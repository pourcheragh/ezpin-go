# AuthTokenPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Client ID provided by EZpin | [optional] 
**SecretKey** | Pointer to **string** | Secret key provided by EZpin | [optional] 

## Methods

### NewAuthTokenPostRequest

`func NewAuthTokenPostRequest() *AuthTokenPostRequest`

NewAuthTokenPostRequest instantiates a new AuthTokenPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenPostRequestWithDefaults

`func NewAuthTokenPostRequestWithDefaults() *AuthTokenPostRequest`

NewAuthTokenPostRequestWithDefaults instantiates a new AuthTokenPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AuthTokenPostRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthTokenPostRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthTokenPostRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AuthTokenPostRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecretKey

`func (o *AuthTokenPostRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AuthTokenPostRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AuthTokenPostRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AuthTokenPostRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


