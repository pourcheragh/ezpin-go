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

// checks if the AuthTokenPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenPost200Response{}

// AuthTokenPost200Response struct for AuthTokenPost200Response
type AuthTokenPost200Response struct {
	// Access token for APIs
	Access *string `json:"access,omitempty"`
	// Token expire time in (ms)
	Expire *float64 `json:"expire,omitempty"`
}

// NewAuthTokenPost200Response instantiates a new AuthTokenPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenPost200Response() *AuthTokenPost200Response {
	this := AuthTokenPost200Response{}
	return &this
}

// NewAuthTokenPost200ResponseWithDefaults instantiates a new AuthTokenPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenPost200ResponseWithDefaults() *AuthTokenPost200Response {
	this := AuthTokenPost200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AuthTokenPost200Response) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenPost200Response) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AuthTokenPost200Response) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *AuthTokenPost200Response) SetAccess(v string) {
	o.Access = &v
}

// GetExpire returns the Expire field value if set, zero value otherwise.
func (o *AuthTokenPost200Response) GetExpire() float64 {
	if o == nil || IsNil(o.Expire) {
		var ret float64
		return ret
	}
	return *o.Expire
}

// GetExpireOk returns a tuple with the Expire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenPost200Response) GetExpireOk() (*float64, bool) {
	if o == nil || IsNil(o.Expire) {
		return nil, false
	}
	return o.Expire, true
}

// HasExpire returns a boolean if a field has been set.
func (o *AuthTokenPost200Response) HasExpire() bool {
	if o != nil && !IsNil(o.Expire) {
		return true
	}

	return false
}

// SetExpire gets a reference to the given float64 and assigns it to the Expire field.
func (o *AuthTokenPost200Response) SetExpire(v float64) {
	o.Expire = &v
}

func (o AuthTokenPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.Expire) {
		toSerialize["expire"] = o.Expire
	}
	return toSerialize, nil
}

type NullableAuthTokenPost200Response struct {
	value *AuthTokenPost200Response
	isSet bool
}

func (v NullableAuthTokenPost200Response) Get() *AuthTokenPost200Response {
	return v.value
}

func (v *NullableAuthTokenPost200Response) Set(val *AuthTokenPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenPost200Response(val *AuthTokenPost200Response) *NullableAuthTokenPost200Response {
	return &NullableAuthTokenPost200Response{value: val, isSet: true}
}

func (v NullableAuthTokenPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


