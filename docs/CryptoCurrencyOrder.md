# CryptoCurrencyOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | crypto order id | [optional] 
**ReferenceCode** | Pointer to **string** | refence code of this crypto order | [optional] 
**RequestedPrice** | Pointer to **int32** | The amount of crypto you ordered. | [optional] 
**Price** | Pointer to **int32** | price of this order | [optional] 
**CryptoCurrencyId** | Pointer to **int32** | id of this crypto currency | [optional] 
**CryptoCurrencyData** | Pointer to [**CryptoCurrencyOrderCryptoCurrencyData**](CryptoCurrencyOrderCryptoCurrencyData.md) |  | [optional] 
**CurrencyId** | Pointer to **int32** | id of fiat currency | [optional] 
**CurrencyDaya** | Pointer to [**CryptoCurrencyOrderCurrencyDaya**](CryptoCurrencyOrderCurrencyDaya.md) |  | [optional] 
**Status** | Pointer to **int32** | accept(1) , waiting(0) , reject(-1) , expired(-5) | [optional] 
**CryptoWallet** | Pointer to **string** | The destination wallet. | [optional] 
**StatusText** | Pointer to **string** | text of status | [optional] 
**CommissionPercent** | Pointer to **float32** | percentage of commission defined for this user | [optional] 
**CommissionPrice** | Pointer to **int32** | price of commission defined for this user. | [optional] 
**NetworkFee** | Pointer to **float32** | commission of network. | [optional] 
**CryptoAmount** | Pointer to **int32** | amount of crypto. | [optional] 
**CryptoPrice** | Pointer to **int32** | last price of that crypto currency that you want to buy. | [optional] 
**CanPay** | Pointer to **bool** | can this order be payed or not. | [optional] 
**IsUsed** | Pointer to **bool** | is order redeemed or not. | [optional] 
**UsedTime** | Pointer to **string** | the time of redemption of order. | [optional] 
**CreatedTime** | Pointer to **string** | order creation time. | [optional] 
**TerminalId** | Pointer to **int32** | sub user id | [optional] 

## Methods

### NewCryptoCurrencyOrder

`func NewCryptoCurrencyOrder() *CryptoCurrencyOrder`

NewCryptoCurrencyOrder instantiates a new CryptoCurrencyOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoCurrencyOrderWithDefaults

`func NewCryptoCurrencyOrderWithDefaults() *CryptoCurrencyOrder`

NewCryptoCurrencyOrderWithDefaults instantiates a new CryptoCurrencyOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CryptoCurrencyOrder) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CryptoCurrencyOrder) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CryptoCurrencyOrder) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *CryptoCurrencyOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferenceCode

`func (o *CryptoCurrencyOrder) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *CryptoCurrencyOrder) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *CryptoCurrencyOrder) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.

### HasReferenceCode

`func (o *CryptoCurrencyOrder) HasReferenceCode() bool`

HasReferenceCode returns a boolean if a field has been set.

### GetRequestedPrice

`func (o *CryptoCurrencyOrder) GetRequestedPrice() int32`

GetRequestedPrice returns the RequestedPrice field if non-nil, zero value otherwise.

### GetRequestedPriceOk

`func (o *CryptoCurrencyOrder) GetRequestedPriceOk() (*int32, bool)`

GetRequestedPriceOk returns a tuple with the RequestedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPrice

`func (o *CryptoCurrencyOrder) SetRequestedPrice(v int32)`

SetRequestedPrice sets RequestedPrice field to given value.

### HasRequestedPrice

`func (o *CryptoCurrencyOrder) HasRequestedPrice() bool`

HasRequestedPrice returns a boolean if a field has been set.

### GetPrice

`func (o *CryptoCurrencyOrder) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CryptoCurrencyOrder) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CryptoCurrencyOrder) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CryptoCurrencyOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCryptoCurrencyId

`func (o *CryptoCurrencyOrder) GetCryptoCurrencyId() int32`

GetCryptoCurrencyId returns the CryptoCurrencyId field if non-nil, zero value otherwise.

### GetCryptoCurrencyIdOk

`func (o *CryptoCurrencyOrder) GetCryptoCurrencyIdOk() (*int32, bool)`

GetCryptoCurrencyIdOk returns a tuple with the CryptoCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrencyId

`func (o *CryptoCurrencyOrder) SetCryptoCurrencyId(v int32)`

SetCryptoCurrencyId sets CryptoCurrencyId field to given value.

### HasCryptoCurrencyId

`func (o *CryptoCurrencyOrder) HasCryptoCurrencyId() bool`

HasCryptoCurrencyId returns a boolean if a field has been set.

### GetCryptoCurrencyData

`func (o *CryptoCurrencyOrder) GetCryptoCurrencyData() CryptoCurrencyOrderCryptoCurrencyData`

GetCryptoCurrencyData returns the CryptoCurrencyData field if non-nil, zero value otherwise.

### GetCryptoCurrencyDataOk

`func (o *CryptoCurrencyOrder) GetCryptoCurrencyDataOk() (*CryptoCurrencyOrderCryptoCurrencyData, bool)`

GetCryptoCurrencyDataOk returns a tuple with the CryptoCurrencyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrencyData

`func (o *CryptoCurrencyOrder) SetCryptoCurrencyData(v CryptoCurrencyOrderCryptoCurrencyData)`

SetCryptoCurrencyData sets CryptoCurrencyData field to given value.

### HasCryptoCurrencyData

`func (o *CryptoCurrencyOrder) HasCryptoCurrencyData() bool`

HasCryptoCurrencyData returns a boolean if a field has been set.

### GetCurrencyId

`func (o *CryptoCurrencyOrder) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *CryptoCurrencyOrder) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *CryptoCurrencyOrder) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *CryptoCurrencyOrder) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyDaya

`func (o *CryptoCurrencyOrder) GetCurrencyDaya() CryptoCurrencyOrderCurrencyDaya`

GetCurrencyDaya returns the CurrencyDaya field if non-nil, zero value otherwise.

### GetCurrencyDayaOk

`func (o *CryptoCurrencyOrder) GetCurrencyDayaOk() (*CryptoCurrencyOrderCurrencyDaya, bool)`

GetCurrencyDayaOk returns a tuple with the CurrencyDaya field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDaya

`func (o *CryptoCurrencyOrder) SetCurrencyDaya(v CryptoCurrencyOrderCurrencyDaya)`

SetCurrencyDaya sets CurrencyDaya field to given value.

### HasCurrencyDaya

`func (o *CryptoCurrencyOrder) HasCurrencyDaya() bool`

HasCurrencyDaya returns a boolean if a field has been set.

### GetStatus

`func (o *CryptoCurrencyOrder) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CryptoCurrencyOrder) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CryptoCurrencyOrder) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CryptoCurrencyOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCryptoWallet

`func (o *CryptoCurrencyOrder) GetCryptoWallet() string`

GetCryptoWallet returns the CryptoWallet field if non-nil, zero value otherwise.

### GetCryptoWalletOk

`func (o *CryptoCurrencyOrder) GetCryptoWalletOk() (*string, bool)`

GetCryptoWalletOk returns a tuple with the CryptoWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoWallet

`func (o *CryptoCurrencyOrder) SetCryptoWallet(v string)`

SetCryptoWallet sets CryptoWallet field to given value.

### HasCryptoWallet

`func (o *CryptoCurrencyOrder) HasCryptoWallet() bool`

HasCryptoWallet returns a boolean if a field has been set.

### GetStatusText

`func (o *CryptoCurrencyOrder) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *CryptoCurrencyOrder) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *CryptoCurrencyOrder) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *CryptoCurrencyOrder) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetCommissionPercent

`func (o *CryptoCurrencyOrder) GetCommissionPercent() float32`

GetCommissionPercent returns the CommissionPercent field if non-nil, zero value otherwise.

### GetCommissionPercentOk

`func (o *CryptoCurrencyOrder) GetCommissionPercentOk() (*float32, bool)`

GetCommissionPercentOk returns a tuple with the CommissionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionPercent

`func (o *CryptoCurrencyOrder) SetCommissionPercent(v float32)`

SetCommissionPercent sets CommissionPercent field to given value.

### HasCommissionPercent

`func (o *CryptoCurrencyOrder) HasCommissionPercent() bool`

HasCommissionPercent returns a boolean if a field has been set.

### GetCommissionPrice

`func (o *CryptoCurrencyOrder) GetCommissionPrice() int32`

GetCommissionPrice returns the CommissionPrice field if non-nil, zero value otherwise.

### GetCommissionPriceOk

`func (o *CryptoCurrencyOrder) GetCommissionPriceOk() (*int32, bool)`

GetCommissionPriceOk returns a tuple with the CommissionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionPrice

`func (o *CryptoCurrencyOrder) SetCommissionPrice(v int32)`

SetCommissionPrice sets CommissionPrice field to given value.

### HasCommissionPrice

`func (o *CryptoCurrencyOrder) HasCommissionPrice() bool`

HasCommissionPrice returns a boolean if a field has been set.

### GetNetworkFee

`func (o *CryptoCurrencyOrder) GetNetworkFee() float32`

GetNetworkFee returns the NetworkFee field if non-nil, zero value otherwise.

### GetNetworkFeeOk

`func (o *CryptoCurrencyOrder) GetNetworkFeeOk() (*float32, bool)`

GetNetworkFeeOk returns a tuple with the NetworkFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFee

`func (o *CryptoCurrencyOrder) SetNetworkFee(v float32)`

SetNetworkFee sets NetworkFee field to given value.

### HasNetworkFee

`func (o *CryptoCurrencyOrder) HasNetworkFee() bool`

HasNetworkFee returns a boolean if a field has been set.

### GetCryptoAmount

`func (o *CryptoCurrencyOrder) GetCryptoAmount() int32`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *CryptoCurrencyOrder) GetCryptoAmountOk() (*int32, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *CryptoCurrencyOrder) SetCryptoAmount(v int32)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *CryptoCurrencyOrder) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetCryptoPrice

`func (o *CryptoCurrencyOrder) GetCryptoPrice() int32`

GetCryptoPrice returns the CryptoPrice field if non-nil, zero value otherwise.

### GetCryptoPriceOk

`func (o *CryptoCurrencyOrder) GetCryptoPriceOk() (*int32, bool)`

GetCryptoPriceOk returns a tuple with the CryptoPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPrice

`func (o *CryptoCurrencyOrder) SetCryptoPrice(v int32)`

SetCryptoPrice sets CryptoPrice field to given value.

### HasCryptoPrice

`func (o *CryptoCurrencyOrder) HasCryptoPrice() bool`

HasCryptoPrice returns a boolean if a field has been set.

### GetCanPay

`func (o *CryptoCurrencyOrder) GetCanPay() bool`

GetCanPay returns the CanPay field if non-nil, zero value otherwise.

### GetCanPayOk

`func (o *CryptoCurrencyOrder) GetCanPayOk() (*bool, bool)`

GetCanPayOk returns a tuple with the CanPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPay

`func (o *CryptoCurrencyOrder) SetCanPay(v bool)`

SetCanPay sets CanPay field to given value.

### HasCanPay

`func (o *CryptoCurrencyOrder) HasCanPay() bool`

HasCanPay returns a boolean if a field has been set.

### GetIsUsed

`func (o *CryptoCurrencyOrder) GetIsUsed() bool`

GetIsUsed returns the IsUsed field if non-nil, zero value otherwise.

### GetIsUsedOk

`func (o *CryptoCurrencyOrder) GetIsUsedOk() (*bool, bool)`

GetIsUsedOk returns a tuple with the IsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsed

`func (o *CryptoCurrencyOrder) SetIsUsed(v bool)`

SetIsUsed sets IsUsed field to given value.

### HasIsUsed

`func (o *CryptoCurrencyOrder) HasIsUsed() bool`

HasIsUsed returns a boolean if a field has been set.

### GetUsedTime

`func (o *CryptoCurrencyOrder) GetUsedTime() string`

GetUsedTime returns the UsedTime field if non-nil, zero value otherwise.

### GetUsedTimeOk

`func (o *CryptoCurrencyOrder) GetUsedTimeOk() (*string, bool)`

GetUsedTimeOk returns a tuple with the UsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTime

`func (o *CryptoCurrencyOrder) SetUsedTime(v string)`

SetUsedTime sets UsedTime field to given value.

### HasUsedTime

`func (o *CryptoCurrencyOrder) HasUsedTime() bool`

HasUsedTime returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CryptoCurrencyOrder) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CryptoCurrencyOrder) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CryptoCurrencyOrder) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CryptoCurrencyOrder) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetTerminalId

`func (o *CryptoCurrencyOrder) GetTerminalId() int32`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *CryptoCurrencyOrder) GetTerminalIdOk() (*int32, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *CryptoCurrencyOrder) SetTerminalId(v int32)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *CryptoCurrencyOrder) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


