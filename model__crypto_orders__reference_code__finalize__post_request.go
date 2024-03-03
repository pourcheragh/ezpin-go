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

// checks if the CryptoOrdersReferenceCodeFinalizePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptoOrdersReferenceCodeFinalizePostRequest{}

// CryptoOrdersReferenceCodeFinalizePostRequest struct for CryptoOrdersReferenceCodeFinalizePostRequest
type CryptoOrdersReferenceCodeFinalizePostRequest struct {
	// Cryptocurrency wallet, which will receive the order amount
	CryptoWallet *string `json:"crypto_wallet,omitempty"`
}

// NewCryptoOrdersReferenceCodeFinalizePostRequest instantiates a new CryptoOrdersReferenceCodeFinalizePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoOrdersReferenceCodeFinalizePostRequest() *CryptoOrdersReferenceCodeFinalizePostRequest {
	this := CryptoOrdersReferenceCodeFinalizePostRequest{}
	return &this
}

// NewCryptoOrdersReferenceCodeFinalizePostRequestWithDefaults instantiates a new CryptoOrdersReferenceCodeFinalizePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoOrdersReferenceCodeFinalizePostRequestWithDefaults() *CryptoOrdersReferenceCodeFinalizePostRequest {
	this := CryptoOrdersReferenceCodeFinalizePostRequest{}
	return &this
}

// GetCryptoWallet returns the CryptoWallet field value if set, zero value otherwise.
func (o *CryptoOrdersReferenceCodeFinalizePostRequest) GetCryptoWallet() string {
	if o == nil || IsNil(o.CryptoWallet) {
		var ret string
		return ret
	}
	return *o.CryptoWallet
}

// GetCryptoWalletOk returns a tuple with the CryptoWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoOrdersReferenceCodeFinalizePostRequest) GetCryptoWalletOk() (*string, bool) {
	if o == nil || IsNil(o.CryptoWallet) {
		return nil, false
	}
	return o.CryptoWallet, true
}

// HasCryptoWallet returns a boolean if a field has been set.
func (o *CryptoOrdersReferenceCodeFinalizePostRequest) HasCryptoWallet() bool {
	if o != nil && !IsNil(o.CryptoWallet) {
		return true
	}

	return false
}

// SetCryptoWallet gets a reference to the given string and assigns it to the CryptoWallet field.
func (o *CryptoOrdersReferenceCodeFinalizePostRequest) SetCryptoWallet(v string) {
	o.CryptoWallet = &v
}

func (o CryptoOrdersReferenceCodeFinalizePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoOrdersReferenceCodeFinalizePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CryptoWallet) {
		toSerialize["crypto_wallet"] = o.CryptoWallet
	}
	return toSerialize, nil
}

type NullableCryptoOrdersReferenceCodeFinalizePostRequest struct {
	value *CryptoOrdersReferenceCodeFinalizePostRequest
	isSet bool
}

func (v NullableCryptoOrdersReferenceCodeFinalizePostRequest) Get() *CryptoOrdersReferenceCodeFinalizePostRequest {
	return v.value
}

func (v *NullableCryptoOrdersReferenceCodeFinalizePostRequest) Set(val *CryptoOrdersReferenceCodeFinalizePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoOrdersReferenceCodeFinalizePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoOrdersReferenceCodeFinalizePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoOrdersReferenceCodeFinalizePostRequest(val *CryptoOrdersReferenceCodeFinalizePostRequest) *NullableCryptoOrdersReferenceCodeFinalizePostRequest {
	return &NullableCryptoOrdersReferenceCodeFinalizePostRequest{value: val, isSet: true}
}

func (v NullableCryptoOrdersReferenceCodeFinalizePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoOrdersReferenceCodeFinalizePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


