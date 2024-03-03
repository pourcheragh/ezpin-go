# ExchangeRatesGet200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromCurrency** | Pointer to [**ExchangeRatesGet200ResponseResultsInnerFromCurrency**](ExchangeRatesGet200ResponseResultsInnerFromCurrency.md) |  | [optional] 
**ToCurrency** | Pointer to [**ExchangeRatesGet200ResponseResultsInnerToCurrency**](ExchangeRatesGet200ResponseResultsInnerToCurrency.md) |  | [optional] 
**Rate** | Pointer to **decimal.Decimal** | rate of exchange | [optional] 

## Methods

### NewExchangeRatesGet200ResponseResultsInner

`func NewExchangeRatesGet200ResponseResultsInner() *ExchangeRatesGet200ResponseResultsInner`

NewExchangeRatesGet200ResponseResultsInner instantiates a new ExchangeRatesGet200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRatesGet200ResponseResultsInnerWithDefaults

`func NewExchangeRatesGet200ResponseResultsInnerWithDefaults() *ExchangeRatesGet200ResponseResultsInner`

NewExchangeRatesGet200ResponseResultsInnerWithDefaults instantiates a new ExchangeRatesGet200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromCurrency

`func (o *ExchangeRatesGet200ResponseResultsInner) GetFromCurrency() ExchangeRatesGet200ResponseResultsInnerFromCurrency`

GetFromCurrency returns the FromCurrency field if non-nil, zero value otherwise.

### GetFromCurrencyOk

`func (o *ExchangeRatesGet200ResponseResultsInner) GetFromCurrencyOk() (*ExchangeRatesGet200ResponseResultsInnerFromCurrency, bool)`

GetFromCurrencyOk returns a tuple with the FromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrency

`func (o *ExchangeRatesGet200ResponseResultsInner) SetFromCurrency(v ExchangeRatesGet200ResponseResultsInnerFromCurrency)`

SetFromCurrency sets FromCurrency field to given value.

### HasFromCurrency

`func (o *ExchangeRatesGet200ResponseResultsInner) HasFromCurrency() bool`

HasFromCurrency returns a boolean if a field has been set.

### GetToCurrency

`func (o *ExchangeRatesGet200ResponseResultsInner) GetToCurrency() ExchangeRatesGet200ResponseResultsInnerToCurrency`

GetToCurrency returns the ToCurrency field if non-nil, zero value otherwise.

### GetToCurrencyOk

`func (o *ExchangeRatesGet200ResponseResultsInner) GetToCurrencyOk() (*ExchangeRatesGet200ResponseResultsInnerToCurrency, bool)`

GetToCurrencyOk returns a tuple with the ToCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrency

`func (o *ExchangeRatesGet200ResponseResultsInner) SetToCurrency(v ExchangeRatesGet200ResponseResultsInnerToCurrency)`

SetToCurrency sets ToCurrency field to given value.

### HasToCurrency

`func (o *ExchangeRatesGet200ResponseResultsInner) HasToCurrency() bool`

HasToCurrency returns a boolean if a field has been set.

### GetRate

`func (o *ExchangeRatesGet200ResponseResultsInner) GetRate() decimal.Decimal`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ExchangeRatesGet200ResponseResultsInner) GetRateOk() (*decimal.Decimal, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ExchangeRatesGet200ResponseResultsInner) SetRate(v decimal.Decimal)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ExchangeRatesGet200ResponseResultsInner) HasRate() bool`

HasRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


