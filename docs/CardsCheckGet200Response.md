# CardsCheckGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **string** | barcode of the physical card | [optional] 
**Product** | Pointer to **string** | name of the physical product | [optional] 
**Sku** | Pointer to **int32** | sku of the physical product | [optional] 
**StatusText** | Pointer to **string** | status of activation process | [optional] 
**Status** | Pointer to **int32** | *&#x60;-1&#x60; Pending *&#x60;0&#x60; Inactive *&#x60;1&#x60; Activated | [optional] 
**MinPrice** | Pointer to **decimal.Decimal** | minimumm price available for physical product | [optional] 
**MaxPrice** | Pointer to **decimal.Decimal** | maximumm price available for physical product | [optional] 
**ActivatePrice** | Pointer to **decimal.Decimal** | price that the card has been activated with | [optional] 
**ActivatedTime** | Pointer to **string** | time of activation | [optional] 
**Currency** | Pointer to **string** | currency of physical product | [optional] 

## Methods

### NewCardsCheckGet200Response

`func NewCardsCheckGet200Response() *CardsCheckGet200Response`

NewCardsCheckGet200Response instantiates a new CardsCheckGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardsCheckGet200ResponseWithDefaults

`func NewCardsCheckGet200ResponseWithDefaults() *CardsCheckGet200Response`

NewCardsCheckGet200ResponseWithDefaults instantiates a new CardsCheckGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *CardsCheckGet200Response) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *CardsCheckGet200Response) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *CardsCheckGet200Response) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *CardsCheckGet200Response) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetProduct

`func (o *CardsCheckGet200Response) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CardsCheckGet200Response) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CardsCheckGet200Response) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CardsCheckGet200Response) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSku

`func (o *CardsCheckGet200Response) GetSku() int32`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CardsCheckGet200Response) GetSkuOk() (*int32, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CardsCheckGet200Response) SetSku(v int32)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CardsCheckGet200Response) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetStatusText

`func (o *CardsCheckGet200Response) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *CardsCheckGet200Response) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *CardsCheckGet200Response) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *CardsCheckGet200Response) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetStatus

`func (o *CardsCheckGet200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardsCheckGet200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardsCheckGet200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardsCheckGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMinPrice

`func (o *CardsCheckGet200Response) GetMinPrice() decimal.Decimal`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *CardsCheckGet200Response) GetMinPriceOk() (*decimal.Decimal, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *CardsCheckGet200Response) SetMinPrice(v decimal.Decimal)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *CardsCheckGet200Response) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *CardsCheckGet200Response) GetMaxPrice() decimal.Decimal`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *CardsCheckGet200Response) GetMaxPriceOk() (*decimal.Decimal, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *CardsCheckGet200Response) SetMaxPrice(v decimal.Decimal)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *CardsCheckGet200Response) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetActivatePrice

`func (o *CardsCheckGet200Response) GetActivatePrice() decimal.Decimal`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *CardsCheckGet200Response) GetActivatePriceOk() (*decimal.Decimal, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *CardsCheckGet200Response) SetActivatePrice(v decimal.Decimal)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *CardsCheckGet200Response) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetActivatedTime

`func (o *CardsCheckGet200Response) GetActivatedTime() string`

GetActivatedTime returns the ActivatedTime field if non-nil, zero value otherwise.

### GetActivatedTimeOk

`func (o *CardsCheckGet200Response) GetActivatedTimeOk() (*string, bool)`

GetActivatedTimeOk returns a tuple with the ActivatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedTime

`func (o *CardsCheckGet200Response) SetActivatedTime(v string)`

SetActivatedTime sets ActivatedTime field to given value.

### HasActivatedTime

`func (o *CardsCheckGet200Response) HasActivatedTime() bool`

HasActivatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *CardsCheckGet200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CardsCheckGet200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CardsCheckGet200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CardsCheckGet200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


