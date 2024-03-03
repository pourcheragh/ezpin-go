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

// checks if the ExchangeRatesGet200ResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeRatesGet200ResponseResultsInner{}

// ExchangeRatesGet200ResponseResultsInner struct for ExchangeRatesGet200ResponseResultsInner
type ExchangeRatesGet200ResponseResultsInner struct {
	FromCurrency *ExchangeRatesGet200ResponseResultsInnerFromCurrency `json:"from_currency,omitempty"`
	ToCurrency *ExchangeRatesGet200ResponseResultsInnerToCurrency `json:"to_currency,omitempty"`
	// rate of exchange
	Rate *decimal.Decimal `json:"rate,omitempty"`
}

// NewExchangeRatesGet200ResponseResultsInner instantiates a new ExchangeRatesGet200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeRatesGet200ResponseResultsInner() *ExchangeRatesGet200ResponseResultsInner {
	this := ExchangeRatesGet200ResponseResultsInner{}
	return &this
}

// NewExchangeRatesGet200ResponseResultsInnerWithDefaults instantiates a new ExchangeRatesGet200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeRatesGet200ResponseResultsInnerWithDefaults() *ExchangeRatesGet200ResponseResultsInner {
	this := ExchangeRatesGet200ResponseResultsInner{}
	return &this
}

// GetFromCurrency returns the FromCurrency field value if set, zero value otherwise.
func (o *ExchangeRatesGet200ResponseResultsInner) GetFromCurrency() ExchangeRatesGet200ResponseResultsInnerFromCurrency {
	if o == nil || IsNil(o.FromCurrency) {
		var ret ExchangeRatesGet200ResponseResultsInnerFromCurrency
		return ret
	}
	return *o.FromCurrency
}

// GetFromCurrencyOk returns a tuple with the FromCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeRatesGet200ResponseResultsInner) GetFromCurrencyOk() (*ExchangeRatesGet200ResponseResultsInnerFromCurrency, bool) {
	if o == nil || IsNil(o.FromCurrency) {
		return nil, false
	}
	return o.FromCurrency, true
}

// HasFromCurrency returns a boolean if a field has been set.
func (o *ExchangeRatesGet200ResponseResultsInner) HasFromCurrency() bool {
	if o != nil && !IsNil(o.FromCurrency) {
		return true
	}

	return false
}

// SetFromCurrency gets a reference to the given ExchangeRatesGet200ResponseResultsInnerFromCurrency and assigns it to the FromCurrency field.
func (o *ExchangeRatesGet200ResponseResultsInner) SetFromCurrency(v ExchangeRatesGet200ResponseResultsInnerFromCurrency) {
	o.FromCurrency = &v
}

// GetToCurrency returns the ToCurrency field value if set, zero value otherwise.
func (o *ExchangeRatesGet200ResponseResultsInner) GetToCurrency() ExchangeRatesGet200ResponseResultsInnerToCurrency {
	if o == nil || IsNil(o.ToCurrency) {
		var ret ExchangeRatesGet200ResponseResultsInnerToCurrency
		return ret
	}
	return *o.ToCurrency
}

// GetToCurrencyOk returns a tuple with the ToCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeRatesGet200ResponseResultsInner) GetToCurrencyOk() (*ExchangeRatesGet200ResponseResultsInnerToCurrency, bool) {
	if o == nil || IsNil(o.ToCurrency) {
		return nil, false
	}
	return o.ToCurrency, true
}

// HasToCurrency returns a boolean if a field has been set.
func (o *ExchangeRatesGet200ResponseResultsInner) HasToCurrency() bool {
	if o != nil && !IsNil(o.ToCurrency) {
		return true
	}

	return false
}

// SetToCurrency gets a reference to the given ExchangeRatesGet200ResponseResultsInnerToCurrency and assigns it to the ToCurrency field.
func (o *ExchangeRatesGet200ResponseResultsInner) SetToCurrency(v ExchangeRatesGet200ResponseResultsInnerToCurrency) {
	o.ToCurrency = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ExchangeRatesGet200ResponseResultsInner) GetRate() decimal.Decimal {
	if o == nil || IsNil(o.Rate) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeRatesGet200ResponseResultsInner) GetRateOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ExchangeRatesGet200ResponseResultsInner) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given decimal.Decimal and assigns it to the Rate field.
func (o *ExchangeRatesGet200ResponseResultsInner) SetRate(v decimal.Decimal) {
	o.Rate = &v
}

func (o ExchangeRatesGet200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeRatesGet200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromCurrency) {
		toSerialize["from_currency"] = o.FromCurrency
	}
	if !IsNil(o.ToCurrency) {
		toSerialize["to_currency"] = o.ToCurrency
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	return toSerialize, nil
}

type NullableExchangeRatesGet200ResponseResultsInner struct {
	value *ExchangeRatesGet200ResponseResultsInner
	isSet bool
}

func (v NullableExchangeRatesGet200ResponseResultsInner) Get() *ExchangeRatesGet200ResponseResultsInner {
	return v.value
}

func (v *NullableExchangeRatesGet200ResponseResultsInner) Set(val *ExchangeRatesGet200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeRatesGet200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeRatesGet200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeRatesGet200ResponseResultsInner(val *ExchangeRatesGet200ResponseResultsInner) *NullableExchangeRatesGet200ResponseResultsInner {
	return &NullableExchangeRatesGet200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableExchangeRatesGet200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeRatesGet200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


