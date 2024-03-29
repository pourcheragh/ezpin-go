/*
EZPin Access API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ezpin

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

// checks if the CardsCheckGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardsCheckGet200Response{}

// CardsCheckGet200Response struct for CardsCheckGet200Response
type CardsCheckGet200Response struct {
	// barcode of the physical card
	Barcode *string `json:"barcode,omitempty"`
	// name of the physical product
	Product *string `json:"product,omitempty"`
	// sku of the physical product
	Sku *int32 `json:"sku,omitempty"`
	// status of activation process
	StatusText *string `json:"status_text,omitempty"`
	// *`-1` Pending *`0` Inactive *`1` Activated
	Status *int32 `json:"status,omitempty"`
	// minimumm price available for physical product
	MinPrice *decimal.Decimal `json:"min_price,omitempty"`
	// maximumm price available for physical product
	MaxPrice *decimal.Decimal `json:"max_price,omitempty"`
	// price that the card has been activated with
	ActivatePrice *decimal.Decimal `json:"activate_price,omitempty"`
	// time of activation
	ActivatedTime *string `json:"activated_time,omitempty"`
	// currency of physical product
	Currency *string `json:"currency,omitempty"`
}

// NewCardsCheckGet200Response instantiates a new CardsCheckGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardsCheckGet200Response() *CardsCheckGet200Response {
	this := CardsCheckGet200Response{}
	return &this
}

// NewCardsCheckGet200ResponseWithDefaults instantiates a new CardsCheckGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardsCheckGet200ResponseWithDefaults() *CardsCheckGet200Response {
	this := CardsCheckGet200Response{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetBarcode() string {
	if o == nil || IsNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetBarcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *CardsCheckGet200Response) SetBarcode(v string) {
	o.Barcode = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *CardsCheckGet200Response) SetProduct(v string) {
	o.Product = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetSku() int32 {
	if o == nil || IsNil(o.Sku) {
		var ret int32
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetSkuOk() (*int32, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given int32 and assigns it to the Sku field.
func (o *CardsCheckGet200Response) SetSku(v int32) {
	o.Sku = &v
}

// GetStatusText returns the StatusText field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetStatusText() string {
	if o == nil || IsNil(o.StatusText) {
		var ret string
		return ret
	}
	return *o.StatusText
}

// GetStatusTextOk returns a tuple with the StatusText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetStatusTextOk() (*string, bool) {
	if o == nil || IsNil(o.StatusText) {
		return nil, false
	}
	return o.StatusText, true
}

// HasStatusText returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasStatusText() bool {
	if o != nil && !IsNil(o.StatusText) {
		return true
	}

	return false
}

// SetStatusText gets a reference to the given string and assigns it to the StatusText field.
func (o *CardsCheckGet200Response) SetStatusText(v string) {
	o.StatusText = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *CardsCheckGet200Response) SetStatus(v int32) {
	o.Status = &v
}

// GetMinPrice returns the MinPrice field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetMinPrice() decimal.Decimal {
	if o == nil || IsNil(o.MinPrice) {
		var ret decimal.Decimal
		return ret
	}
	return *o.MinPrice
}

// GetMinPriceOk returns a tuple with the MinPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetMinPriceOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.MinPrice) {
		return nil, false
	}
	return o.MinPrice, true
}

// HasMinPrice returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasMinPrice() bool {
	if o != nil && !IsNil(o.MinPrice) {
		return true
	}

	return false
}

// SetMinPrice gets a reference to the given decimal.Decimal and assigns it to the MinPrice field.
func (o *CardsCheckGet200Response) SetMinPrice(v decimal.Decimal) {
	o.MinPrice = &v
}

// GetMaxPrice returns the MaxPrice field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetMaxPrice() decimal.Decimal {
	if o == nil || IsNil(o.MaxPrice) {
		var ret decimal.Decimal
		return ret
	}
	return *o.MaxPrice
}

// GetMaxPriceOk returns a tuple with the MaxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetMaxPriceOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.MaxPrice) {
		return nil, false
	}
	return o.MaxPrice, true
}

// HasMaxPrice returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasMaxPrice() bool {
	if o != nil && !IsNil(o.MaxPrice) {
		return true
	}

	return false
}

// SetMaxPrice gets a reference to the given decimal.Decimal and assigns it to the MaxPrice field.
func (o *CardsCheckGet200Response) SetMaxPrice(v decimal.Decimal) {
	o.MaxPrice = &v
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetActivatePrice() decimal.Decimal {
	if o == nil || IsNil(o.ActivatePrice) {
		var ret decimal.Decimal
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetActivatePriceOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasActivatePrice() bool {
	if o != nil && !IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given decimal.Decimal and assigns it to the ActivatePrice field.
func (o *CardsCheckGet200Response) SetActivatePrice(v decimal.Decimal) {
	o.ActivatePrice = &v
}

// GetActivatedTime returns the ActivatedTime field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetActivatedTime() string {
	if o == nil || IsNil(o.ActivatedTime) {
		var ret string
		return ret
	}
	return *o.ActivatedTime
}

// GetActivatedTimeOk returns a tuple with the ActivatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetActivatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ActivatedTime) {
		return nil, false
	}
	return o.ActivatedTime, true
}

// HasActivatedTime returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasActivatedTime() bool {
	if o != nil && !IsNil(o.ActivatedTime) {
		return true
	}

	return false
}

// SetActivatedTime gets a reference to the given string and assigns it to the ActivatedTime field.
func (o *CardsCheckGet200Response) SetActivatedTime(v string) {
	o.ActivatedTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CardsCheckGet200Response) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsCheckGet200Response) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CardsCheckGet200Response) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CardsCheckGet200Response) SetCurrency(v string) {
	o.Currency = &v
}

func (o CardsCheckGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardsCheckGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.StatusText) {
		toSerialize["status_text"] = o.StatusText
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.MinPrice) {
		toSerialize["min_price"] = o.MinPrice
	}
	if !IsNil(o.MaxPrice) {
		toSerialize["max_price"] = o.MaxPrice
	}
	if !IsNil(o.ActivatePrice) {
		toSerialize["activate_price"] = o.ActivatePrice
	}
	if !IsNil(o.ActivatedTime) {
		toSerialize["activated_time"] = o.ActivatedTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableCardsCheckGet200Response struct {
	value *CardsCheckGet200Response
	isSet bool
}

func (v NullableCardsCheckGet200Response) Get() *CardsCheckGet200Response {
	return v.value
}

func (v *NullableCardsCheckGet200Response) Set(val *CardsCheckGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCardsCheckGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCardsCheckGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardsCheckGet200Response(val *CardsCheckGet200Response) *NullableCardsCheckGet200Response {
	return &NullableCardsCheckGet200Response{value: val, isSet: true}
}

func (v NullableCardsCheckGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardsCheckGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


