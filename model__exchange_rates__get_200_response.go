/*
EZPin Access API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ezpin

import (
	"encoding/json"
)

// checks if the ExchangeRatesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeRatesGet200Response{}

// ExchangeRatesGet200Response struct for ExchangeRatesGet200Response
type ExchangeRatesGet200Response struct {
	Results []ExchangeRatesGet200ResponseResultsInner `json:"results,omitempty"`
}

// NewExchangeRatesGet200Response instantiates a new ExchangeRatesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeRatesGet200Response() *ExchangeRatesGet200Response {
	this := ExchangeRatesGet200Response{}
	return &this
}

// NewExchangeRatesGet200ResponseWithDefaults instantiates a new ExchangeRatesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeRatesGet200ResponseWithDefaults() *ExchangeRatesGet200Response {
	this := ExchangeRatesGet200Response{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ExchangeRatesGet200Response) GetResults() []ExchangeRatesGet200ResponseResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []ExchangeRatesGet200ResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeRatesGet200Response) GetResultsOk() ([]ExchangeRatesGet200ResponseResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ExchangeRatesGet200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ExchangeRatesGet200ResponseResultsInner and assigns it to the Results field.
func (o *ExchangeRatesGet200Response) SetResults(v []ExchangeRatesGet200ResponseResultsInner) {
	o.Results = v
}

func (o ExchangeRatesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeRatesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableExchangeRatesGet200Response struct {
	value *ExchangeRatesGet200Response
	isSet bool
}

func (v NullableExchangeRatesGet200Response) Get() *ExchangeRatesGet200Response {
	return v.value
}

func (v *NullableExchangeRatesGet200Response) Set(val *ExchangeRatesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeRatesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeRatesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeRatesGet200Response(val *ExchangeRatesGet200Response) *NullableExchangeRatesGet200Response {
	return &NullableExchangeRatesGet200Response{value: val, isSet: true}
}

func (v NullableExchangeRatesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeRatesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


