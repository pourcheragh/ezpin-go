# CryptoCurrencyCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **float32** | id of the fiat currency | [optional] 
**CurrencyTitle** | Pointer to **string** | Title of the fiat currency | [optional] 
**Symbol** | Pointer to **string** | symbol of this fiat currency | [optional] 
**Code** | Pointer to **string** | code of this fiat currency | [optional] 
**Price** | Pointer to **float32** | Price of crypto that you want to buy | [optional] 
**MinAmount** | Pointer to **float32** | Minumumm amount that you can buy more than this price. | [optional] 
**MaxAmount** | Pointer to **float32** | Maximumm amount that you can buy this crpto currency less than this amount | [optional] 
**CryptoCurrencyId** | Pointer to **string** | id of this crypto currency | [optional] 
**CryptoTitle** | Pointer to **string** | name of this crypto currency | [optional] 
**CryptoSymbol** | Pointer to **string** | symbol of this crypto currency | [optional] 
**CryptoCode** | Pointer to **string** | code of this crypto currency | [optional] 

## Methods

### NewCryptoCurrencyCatalog

`func NewCryptoCurrencyCatalog() *CryptoCurrencyCatalog`

NewCryptoCurrencyCatalog instantiates a new CryptoCurrencyCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoCurrencyCatalogWithDefaults

`func NewCryptoCurrencyCatalogWithDefaults() *CryptoCurrencyCatalog`

NewCryptoCurrencyCatalogWithDefaults instantiates a new CryptoCurrencyCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *CryptoCurrencyCatalog) GetCurrencyId() float32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *CryptoCurrencyCatalog) GetCurrencyIdOk() (*float32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *CryptoCurrencyCatalog) SetCurrencyId(v float32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *CryptoCurrencyCatalog) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyTitle

`func (o *CryptoCurrencyCatalog) GetCurrencyTitle() string`

GetCurrencyTitle returns the CurrencyTitle field if non-nil, zero value otherwise.

### GetCurrencyTitleOk

`func (o *CryptoCurrencyCatalog) GetCurrencyTitleOk() (*string, bool)`

GetCurrencyTitleOk returns a tuple with the CurrencyTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyTitle

`func (o *CryptoCurrencyCatalog) SetCurrencyTitle(v string)`

SetCurrencyTitle sets CurrencyTitle field to given value.

### HasCurrencyTitle

`func (o *CryptoCurrencyCatalog) HasCurrencyTitle() bool`

HasCurrencyTitle returns a boolean if a field has been set.

### GetSymbol

`func (o *CryptoCurrencyCatalog) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CryptoCurrencyCatalog) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CryptoCurrencyCatalog) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CryptoCurrencyCatalog) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCode

`func (o *CryptoCurrencyCatalog) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CryptoCurrencyCatalog) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CryptoCurrencyCatalog) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CryptoCurrencyCatalog) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPrice

`func (o *CryptoCurrencyCatalog) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CryptoCurrencyCatalog) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CryptoCurrencyCatalog) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CryptoCurrencyCatalog) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMinAmount

`func (o *CryptoCurrencyCatalog) GetMinAmount() float32`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *CryptoCurrencyCatalog) GetMinAmountOk() (*float32, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *CryptoCurrencyCatalog) SetMinAmount(v float32)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *CryptoCurrencyCatalog) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetMaxAmount

`func (o *CryptoCurrencyCatalog) GetMaxAmount() float32`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *CryptoCurrencyCatalog) GetMaxAmountOk() (*float32, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *CryptoCurrencyCatalog) SetMaxAmount(v float32)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *CryptoCurrencyCatalog) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetCryptoCurrencyId

`func (o *CryptoCurrencyCatalog) GetCryptoCurrencyId() string`

GetCryptoCurrencyId returns the CryptoCurrencyId field if non-nil, zero value otherwise.

### GetCryptoCurrencyIdOk

`func (o *CryptoCurrencyCatalog) GetCryptoCurrencyIdOk() (*string, bool)`

GetCryptoCurrencyIdOk returns a tuple with the CryptoCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrencyId

`func (o *CryptoCurrencyCatalog) SetCryptoCurrencyId(v string)`

SetCryptoCurrencyId sets CryptoCurrencyId field to given value.

### HasCryptoCurrencyId

`func (o *CryptoCurrencyCatalog) HasCryptoCurrencyId() bool`

HasCryptoCurrencyId returns a boolean if a field has been set.

### GetCryptoTitle

`func (o *CryptoCurrencyCatalog) GetCryptoTitle() string`

GetCryptoTitle returns the CryptoTitle field if non-nil, zero value otherwise.

### GetCryptoTitleOk

`func (o *CryptoCurrencyCatalog) GetCryptoTitleOk() (*string, bool)`

GetCryptoTitleOk returns a tuple with the CryptoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoTitle

`func (o *CryptoCurrencyCatalog) SetCryptoTitle(v string)`

SetCryptoTitle sets CryptoTitle field to given value.

### HasCryptoTitle

`func (o *CryptoCurrencyCatalog) HasCryptoTitle() bool`

HasCryptoTitle returns a boolean if a field has been set.

### GetCryptoSymbol

`func (o *CryptoCurrencyCatalog) GetCryptoSymbol() string`

GetCryptoSymbol returns the CryptoSymbol field if non-nil, zero value otherwise.

### GetCryptoSymbolOk

`func (o *CryptoCurrencyCatalog) GetCryptoSymbolOk() (*string, bool)`

GetCryptoSymbolOk returns a tuple with the CryptoSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoSymbol

`func (o *CryptoCurrencyCatalog) SetCryptoSymbol(v string)`

SetCryptoSymbol sets CryptoSymbol field to given value.

### HasCryptoSymbol

`func (o *CryptoCurrencyCatalog) HasCryptoSymbol() bool`

HasCryptoSymbol returns a boolean if a field has been set.

### GetCryptoCode

`func (o *CryptoCurrencyCatalog) GetCryptoCode() string`

GetCryptoCode returns the CryptoCode field if non-nil, zero value otherwise.

### GetCryptoCodeOk

`func (o *CryptoCurrencyCatalog) GetCryptoCodeOk() (*string, bool)`

GetCryptoCodeOk returns a tuple with the CryptoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCode

`func (o *CryptoCurrencyCatalog) SetCryptoCode(v string)`

SetCryptoCode sets CryptoCode field to given value.

### HasCryptoCode

`func (o *CryptoCurrencyCatalog) HasCryptoCode() bool`

HasCryptoCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


