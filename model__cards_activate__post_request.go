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

// checks if the CardsActivatePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardsActivatePostRequest{}

// CardsActivatePostRequest struct for CardsActivatePostRequest
type CardsActivatePostRequest struct {
	// Barcode Number Of Physical Card
	Barcode *string `json:"barcode,omitempty"`
	// SKU Of Product
	Sku *int32 `json:"sku,omitempty"`
	// Item price
	Price *decimal.Decimal `json:"price,omitempty"`
	// Terminal ID of the sub-users that can be defined in setting section in user panel.
	TerminalId *int32 `json:"terminal_id,omitempty"`
	// Pin defined for sub-user
	TerminalPin *int32 `json:"terminal_pin,omitempty"`
	// A unique UUID v4 referece code must be included in request
	ReferenceCode *string `json:"reference_code,omitempty"`
}

// NewCardsActivatePostRequest instantiates a new CardsActivatePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardsActivatePostRequest() *CardsActivatePostRequest {
	this := CardsActivatePostRequest{}
	return &this
}

// NewCardsActivatePostRequestWithDefaults instantiates a new CardsActivatePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardsActivatePostRequestWithDefaults() *CardsActivatePostRequest {
	this := CardsActivatePostRequest{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *CardsActivatePostRequest) GetBarcode() string {
	if o == nil || IsNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsActivatePostRequest) GetBarcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *CardsActivatePostRequest) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *CardsActivatePostRequest) SetBarcode(v string) {
	o.Barcode = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *CardsActivatePostRequest) GetSku() int32 {
	if o == nil || IsNil(o.Sku) {
		var ret int32
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsActivatePostRequest) GetSkuOk() (*int32, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *CardsActivatePostRequest) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given int32 and assigns it to the Sku field.
func (o *CardsActivatePostRequest) SetSku(v int32) {
	o.Sku = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CardsActivatePostRequest) GetPrice() decimal.Decimal {
	if o == nil || IsNil(o.Price) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsActivatePostRequest) GetPriceOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CardsActivatePostRequest) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given decimal.Decimal and assigns it to the Price field.
func (o *CardsActivatePostRequest) SetPrice(v decimal.Decimal) {
	o.Price = &v
}

// GetTerminalId returns the TerminalId field value if set, zero value otherwise.
func (o *CardsActivatePostRequest) GetTerminalId() int32 {
	if o == nil || IsNil(o.TerminalId) {
		var ret int32
		return ret
	}
	return *o.TerminalId
}

// GetTerminalIdOk returns a tuple with the TerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsActivatePostRequest) GetTerminalIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TerminalId) {
		return nil, false
	}
	return o.TerminalId, true
}

// HasTerminalId returns a boolean if a field has been set.
func (o *CardsActivatePostRequest) HasTerminalId() bool {
	if o != nil && !IsNil(o.TerminalId) {
		return true
	}

	return false
}

// SetTerminalId gets a reference to the given int32 and assigns it to the TerminalId field.
func (o *CardsActivatePostRequest) SetTerminalId(v int32) {
	o.TerminalId = &v
}

// GetTerminalPin returns the TerminalPin field value if set, zero value otherwise.
func (o *CardsActivatePostRequest) GetTerminalPin() int32 {
	if o == nil || IsNil(o.TerminalPin) {
		var ret int32
		return ret
	}
	return *o.TerminalPin
}

// GetTerminalPinOk returns a tuple with the TerminalPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsActivatePostRequest) GetTerminalPinOk() (*int32, bool) {
	if o == nil || IsNil(o.TerminalPin) {
		return nil, false
	}
	return o.TerminalPin, true
}

// HasTerminalPin returns a boolean if a field has been set.
func (o *CardsActivatePostRequest) HasTerminalPin() bool {
	if o != nil && !IsNil(o.TerminalPin) {
		return true
	}

	return false
}

// SetTerminalPin gets a reference to the given int32 and assigns it to the TerminalPin field.
func (o *CardsActivatePostRequest) SetTerminalPin(v int32) {
	o.TerminalPin = &v
}

// GetReferenceCode returns the ReferenceCode field value if set, zero value otherwise.
func (o *CardsActivatePostRequest) GetReferenceCode() string {
	if o == nil || IsNil(o.ReferenceCode) {
		var ret string
		return ret
	}
	return *o.ReferenceCode
}

// GetReferenceCodeOk returns a tuple with the ReferenceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardsActivatePostRequest) GetReferenceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceCode) {
		return nil, false
	}
	return o.ReferenceCode, true
}

// HasReferenceCode returns a boolean if a field has been set.
func (o *CardsActivatePostRequest) HasReferenceCode() bool {
	if o != nil && !IsNil(o.ReferenceCode) {
		return true
	}

	return false
}

// SetReferenceCode gets a reference to the given string and assigns it to the ReferenceCode field.
func (o *CardsActivatePostRequest) SetReferenceCode(v string) {
	o.ReferenceCode = &v
}

func (o CardsActivatePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardsActivatePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.TerminalId) {
		toSerialize["terminal_id"] = o.TerminalId
	}
	if !IsNil(o.TerminalPin) {
		toSerialize["terminal_pin"] = o.TerminalPin
	}
	if !IsNil(o.ReferenceCode) {
		toSerialize["reference_code"] = o.ReferenceCode
	}
	return toSerialize, nil
}

type NullableCardsActivatePostRequest struct {
	value *CardsActivatePostRequest
	isSet bool
}

func (v NullableCardsActivatePostRequest) Get() *CardsActivatePostRequest {
	return v.value
}

func (v *NullableCardsActivatePostRequest) Set(val *CardsActivatePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardsActivatePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardsActivatePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardsActivatePostRequest(val *CardsActivatePostRequest) *NullableCardsActivatePostRequest {
	return &NullableCardsActivatePostRequest{value: val, isSet: true}
}

func (v NullableCardsActivatePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardsActivatePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


