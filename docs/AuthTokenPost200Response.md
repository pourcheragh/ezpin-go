# AuthTokenPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | Access token for APIs | [optional] 
**Expire** | Pointer to **decimal.Decimal** | Token expire time in (ms) | [optional] 

## Methods

### NewAuthTokenPost200Response

`func NewAuthTokenPost200Response() *AuthTokenPost200Response`

NewAuthTokenPost200Response instantiates a new AuthTokenPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenPost200ResponseWithDefaults

`func NewAuthTokenPost200ResponseWithDefaults() *AuthTokenPost200Response`

NewAuthTokenPost200ResponseWithDefaults instantiates a new AuthTokenPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AuthTokenPost200Response) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AuthTokenPost200Response) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AuthTokenPost200Response) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AuthTokenPost200Response) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetExpire

`func (o *AuthTokenPost200Response) GetExpire() decimal.Decimal`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *AuthTokenPost200Response) GetExpireOk() (*decimal.Decimal, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *AuthTokenPost200Response) SetExpire(v decimal.Decimal)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *AuthTokenPost200Response) HasExpire() bool`

HasExpire returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


