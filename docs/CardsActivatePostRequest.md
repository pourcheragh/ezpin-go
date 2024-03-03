# CardsActivatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **string** | Barcode Number Of Physical Card | [optional] 
**Sku** | Pointer to **int32** | SKU Of Product | [optional] 
**Price** | Pointer to **decimal.Decimal** | Item price | [optional] 
**TerminalId** | Pointer to **int32** | Terminal ID of the sub-users that can be defined in setting section in user panel. | [optional] 
**TerminalPin** | Pointer to **int32** | Pin defined for sub-user | [optional] 
**ReferenceCode** | Pointer to **string** | A unique UUID v4 referece code must be included in request | [optional] 

## Methods

### NewCardsActivatePostRequest

`func NewCardsActivatePostRequest() *CardsActivatePostRequest`

NewCardsActivatePostRequest instantiates a new CardsActivatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardsActivatePostRequestWithDefaults

`func NewCardsActivatePostRequestWithDefaults() *CardsActivatePostRequest`

NewCardsActivatePostRequestWithDefaults instantiates a new CardsActivatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *CardsActivatePostRequest) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *CardsActivatePostRequest) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *CardsActivatePostRequest) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *CardsActivatePostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetSku

`func (o *CardsActivatePostRequest) GetSku() int32`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CardsActivatePostRequest) GetSkuOk() (*int32, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CardsActivatePostRequest) SetSku(v int32)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CardsActivatePostRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPrice

`func (o *CardsActivatePostRequest) GetPrice() decimal.Decimal`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CardsActivatePostRequest) GetPriceOk() (*decimal.Decimal, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CardsActivatePostRequest) SetPrice(v decimal.Decimal)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CardsActivatePostRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTerminalId

`func (o *CardsActivatePostRequest) GetTerminalId() int32`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *CardsActivatePostRequest) GetTerminalIdOk() (*int32, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *CardsActivatePostRequest) SetTerminalId(v int32)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *CardsActivatePostRequest) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetTerminalPin

`func (o *CardsActivatePostRequest) GetTerminalPin() int32`

GetTerminalPin returns the TerminalPin field if non-nil, zero value otherwise.

### GetTerminalPinOk

`func (o *CardsActivatePostRequest) GetTerminalPinOk() (*int32, bool)`

GetTerminalPinOk returns a tuple with the TerminalPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPin

`func (o *CardsActivatePostRequest) SetTerminalPin(v int32)`

SetTerminalPin sets TerminalPin field to given value.

### HasTerminalPin

`func (o *CardsActivatePostRequest) HasTerminalPin() bool`

HasTerminalPin returns a boolean if a field has been set.

### GetReferenceCode

`func (o *CardsActivatePostRequest) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *CardsActivatePostRequest) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *CardsActivatePostRequest) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.

### HasReferenceCode

`func (o *CardsActivatePostRequest) HasReferenceCode() bool`

HasReferenceCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


