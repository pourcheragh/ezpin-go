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

// checks if the CatalogShowingPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogShowingPrice{}

// CatalogShowingPrice struct for CatalogShowingPrice
type CatalogShowingPrice struct {
	// showing price
	Price *decimal.Decimal `json:"price,omitempty"`
	ShowingCurrency []ShowingCurrency `json:"showing_currency,omitempty"`
}

// NewCatalogShowingPrice instantiates a new CatalogShowingPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogShowingPrice() *CatalogShowingPrice {
	this := CatalogShowingPrice{}
	return &this
}

// NewCatalogShowingPriceWithDefaults instantiates a new CatalogShowingPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogShowingPriceWithDefaults() *CatalogShowingPrice {
	this := CatalogShowingPrice{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CatalogShowingPrice) GetPrice() decimal.Decimal {
	if o == nil || IsNil(o.Price) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogShowingPrice) GetPriceOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CatalogShowingPrice) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given decimal.Decimal and assigns it to the Price field.
func (o *CatalogShowingPrice) SetPrice(v decimal.Decimal) {
	o.Price = &v
}

// GetShowingCurrency returns the ShowingCurrency field value if set, zero value otherwise.
func (o *CatalogShowingPrice) GetShowingCurrency() []ShowingCurrency {
	if o == nil || IsNil(o.ShowingCurrency) {
		var ret []ShowingCurrency
		return ret
	}
	return o.ShowingCurrency
}

// GetShowingCurrencyOk returns a tuple with the ShowingCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogShowingPrice) GetShowingCurrencyOk() ([]ShowingCurrency, bool) {
	if o == nil || IsNil(o.ShowingCurrency) {
		return nil, false
	}
	return o.ShowingCurrency, true
}

// HasShowingCurrency returns a boolean if a field has been set.
func (o *CatalogShowingPrice) HasShowingCurrency() bool {
	if o != nil && !IsNil(o.ShowingCurrency) {
		return true
	}

	return false
}

// SetShowingCurrency gets a reference to the given []ShowingCurrency and assigns it to the ShowingCurrency field.
func (o *CatalogShowingPrice) SetShowingCurrency(v []ShowingCurrency) {
	o.ShowingCurrency = v
}

func (o CatalogShowingPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogShowingPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.ShowingCurrency) {
		toSerialize["showing_currency"] = o.ShowingCurrency
	}
	return toSerialize, nil
}

type NullableCatalogShowingPrice struct {
	value *CatalogShowingPrice
	isSet bool
}

func (v NullableCatalogShowingPrice) Get() *CatalogShowingPrice {
	return v.value
}

func (v *NullableCatalogShowingPrice) Set(val *CatalogShowingPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogShowingPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogShowingPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogShowingPrice(val *CatalogShowingPrice) *NullableCatalogShowingPrice {
	return &NullableCatalogShowingPrice{value: val, isSet: true}
}

func (v NullableCatalogShowingPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogShowingPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


