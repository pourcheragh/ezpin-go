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

// checks if the CatalogsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogsGet200Response{}

// CatalogsGet200Response struct for CatalogsGet200Response
type CatalogsGet200Response struct {
	// Number of total catalogs
	Count *float32 `json:"count,omitempty"`
	Results []Catalog `json:"results,omitempty"`
}

// NewCatalogsGet200Response instantiates a new CatalogsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogsGet200Response() *CatalogsGet200Response {
	this := CatalogsGet200Response{}
	return &this
}

// NewCatalogsGet200ResponseWithDefaults instantiates a new CatalogsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogsGet200ResponseWithDefaults() *CatalogsGet200Response {
	this := CatalogsGet200Response{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CatalogsGet200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogsGet200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CatalogsGet200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *CatalogsGet200Response) SetCount(v float32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CatalogsGet200Response) GetResults() []Catalog {
	if o == nil || IsNil(o.Results) {
		var ret []Catalog
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogsGet200Response) GetResultsOk() ([]Catalog, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CatalogsGet200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Catalog and assigns it to the Results field.
func (o *CatalogsGet200Response) SetResults(v []Catalog) {
	o.Results = v
}

func (o CatalogsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableCatalogsGet200Response struct {
	value *CatalogsGet200Response
	isSet bool
}

func (v NullableCatalogsGet200Response) Get() *CatalogsGet200Response {
	return v.value
}

func (v *NullableCatalogsGet200Response) Set(val *CatalogsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogsGet200Response(val *CatalogsGet200Response) *NullableCatalogsGet200Response {
	return &NullableCatalogsGet200Response{value: val, isSet: true}
}

func (v NullableCatalogsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


