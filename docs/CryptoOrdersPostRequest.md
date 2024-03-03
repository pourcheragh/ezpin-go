# CryptoOrdersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **int32** | ID of Fiat currency | [optional] 
**CryptoCurrencyId** | Pointer to **int32** | ID of Crypto Currency | [optional] 
**RequestedPrice** | Pointer to **decimal.Decimal** | The amount you need to exchange from fiat currency to Cryptocurrency | [optional] 
**ReferenceCode** | Pointer to **string** | A unique UUID v4 reference code must be included in request | [optional] 
**TerminalId** | Pointer to **int32** | Terminal ID of the sub-users that can be defined in setting section in user panel | [optional] 
**TerminalPin** | Pointer to **int32** | Pin defined for sub-user | [optional] 

## Methods

### NewCryptoOrdersPostRequest

`func NewCryptoOrdersPostRequest() *CryptoOrdersPostRequest`

NewCryptoOrdersPostRequest instantiates a new CryptoOrdersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoOrdersPostRequestWithDefaults

`func NewCryptoOrdersPostRequestWithDefaults() *CryptoOrdersPostRequest`

NewCryptoOrdersPostRequestWithDefaults instantiates a new CryptoOrdersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *CryptoOrdersPostRequest) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *CryptoOrdersPostRequest) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *CryptoOrdersPostRequest) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *CryptoOrdersPostRequest) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCryptoCurrencyId

`func (o *CryptoOrdersPostRequest) GetCryptoCurrencyId() int32`

GetCryptoCurrencyId returns the CryptoCurrencyId field if non-nil, zero value otherwise.

### GetCryptoCurrencyIdOk

`func (o *CryptoOrdersPostRequest) GetCryptoCurrencyIdOk() (*int32, bool)`

GetCryptoCurrencyIdOk returns a tuple with the CryptoCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrencyId

`func (o *CryptoOrdersPostRequest) SetCryptoCurrencyId(v int32)`

SetCryptoCurrencyId sets CryptoCurrencyId field to given value.

### HasCryptoCurrencyId

`func (o *CryptoOrdersPostRequest) HasCryptoCurrencyId() bool`

HasCryptoCurrencyId returns a boolean if a field has been set.

### GetRequestedPrice

`func (o *CryptoOrdersPostRequest) GetRequestedPrice() decimal.Decimal`

GetRequestedPrice returns the RequestedPrice field if non-nil, zero value otherwise.

### GetRequestedPriceOk

`func (o *CryptoOrdersPostRequest) GetRequestedPriceOk() (*decimal.Decimal, bool)`

GetRequestedPriceOk returns a tuple with the RequestedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPrice

`func (o *CryptoOrdersPostRequest) SetRequestedPrice(v decimal.Decimal)`

SetRequestedPrice sets RequestedPrice field to given value.

### HasRequestedPrice

`func (o *CryptoOrdersPostRequest) HasRequestedPrice() bool`

HasRequestedPrice returns a boolean if a field has been set.

### GetReferenceCode

`func (o *CryptoOrdersPostRequest) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *CryptoOrdersPostRequest) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *CryptoOrdersPostRequest) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.

### HasReferenceCode

`func (o *CryptoOrdersPostRequest) HasReferenceCode() bool`

HasReferenceCode returns a boolean if a field has been set.

### GetTerminalId

`func (o *CryptoOrdersPostRequest) GetTerminalId() int32`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *CryptoOrdersPostRequest) GetTerminalIdOk() (*int32, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *CryptoOrdersPostRequest) SetTerminalId(v int32)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *CryptoOrdersPostRequest) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetTerminalPin

`func (o *CryptoOrdersPostRequest) GetTerminalPin() int32`

GetTerminalPin returns the TerminalPin field if non-nil, zero value otherwise.

### GetTerminalPinOk

`func (o *CryptoOrdersPostRequest) GetTerminalPinOk() (*int32, bool)`

GetTerminalPinOk returns a tuple with the TerminalPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPin

`func (o *CryptoOrdersPostRequest) SetTerminalPin(v int32)`

SetTerminalPin sets TerminalPin field to given value.

### HasTerminalPin

`func (o *CryptoOrdersPostRequest) HasTerminalPin() bool`

HasTerminalPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


