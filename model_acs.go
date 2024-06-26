/*
Lemonway DirectKit API 2.0

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ACS struct for ACS
type ACS struct {
	ActionUrl *string `json:"actionUrl,omitempty"`
	ActionMethod *string `json:"actionMethod,omitempty"`
	PareqFieldName *string `json:"pareqFieldName,omitempty"`
	PareqFieldValue *string `json:"pareqFieldValue,omitempty"`
	TermurlFieldName *string `json:"termurlFieldName,omitempty"`
	MdFieldName *string `json:"mdFieldName,omitempty"`
	MdFieldValue *string `json:"mdFieldValue,omitempty"`
	MpiResult *string `json:"mpiResult,omitempty"`
}

// NewACS instantiates a new ACS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACS() *ACS {
	this := ACS{}
	return &this
}

// NewACSWithDefaults instantiates a new ACS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACSWithDefaults() *ACS {
	this := ACS{}
	return &this
}

// GetActionUrl returns the ActionUrl field value if set, zero value otherwise.
func (o *ACS) GetActionUrl() string {
	if o == nil || o.ActionUrl == nil {
		var ret string
		return ret
	}
	return *o.ActionUrl
}

// GetActionUrlOk returns a tuple with the ActionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACS) GetActionUrlOk() (*string, bool) {
	if o == nil || o.ActionUrl == nil {
		return nil, false
	}
	return o.ActionUrl, true
}

// HasActionUrl returns a boolean if a field has been set.
func (o *ACS) HasActionUrl() bool {
	if o != nil && o.ActionUrl != nil {
		return true
	}

	return false
}

// SetActionUrl gets a reference to the given string and assigns it to the ActionUrl field.
func (o *ACS) SetActionUrl(v string) {
	o.ActionUrl = &v
}

// GetActionMethod returns the ActionMethod field value if set, zero value otherwise.
func (o *ACS) GetActionMethod() string {
	if o == nil || o.ActionMethod == nil {
		var ret string
		return ret
	}
	return *o.ActionMethod
}

// GetActionMethodOk returns a tuple with the ActionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACS) GetActionMethodOk() (*string, bool) {
	if o == nil || o.ActionMethod == nil {
		return nil, false
	}
	return o.ActionMethod, true
}

// HasActionMethod returns a boolean if a field has been set.
func (o *ACS) HasActionMethod() bool {
	if o != nil && o.ActionMethod != nil {
		return true
	}

	return false
}

// SetActionMethod gets a reference to the given string and assigns it to the ActionMethod field.
func (o *ACS) SetActionMethod(v string) {
	o.ActionMethod = &v
}

// GetPareqFieldName returns the PareqFieldName field value if set, zero value otherwise.
func (o *ACS) GetPareqFieldName() string {
	if o == nil || o.PareqFieldName == nil {
		var ret string
		return ret
	}
	return *o.PareqFieldName
}

// GetPareqFieldNameOk returns a tuple with the PareqFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACS) GetPareqFieldNameOk() (*string, bool) {
	if o == nil || o.PareqFieldName == nil {
		return nil, false
	}
	return o.PareqFieldName, true
}

// HasPareqFieldName returns a boolean if a field has been set.
func (o *ACS) HasPareqFieldName() bool {
	if o != nil && o.PareqFieldName != nil {
		return true
	}

	return false
}

// SetPareqFieldName gets a reference to the given string and assigns it to the PareqFieldName field.
func (o *ACS) SetPareqFieldName(v string) {
	o.PareqFieldName = &v
}

// GetPareqFieldValue returns the PareqFieldValue field value if set, zero value otherwise.
func (o *ACS) GetPareqFieldValue() string {
	if o == nil || o.PareqFieldValue == nil {
		var ret string
		return ret
	}
	return *o.PareqFieldValue
}

// GetPareqFieldValueOk returns a tuple with the PareqFieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACS) GetPareqFieldValueOk() (*string, bool) {
	if o == nil || o.PareqFieldValue == nil {
		return nil, false
	}
	return o.PareqFieldValue, true
}

// HasPareqFieldValue returns a boolean if a field has been set.
func (o *ACS) HasPareqFieldValue() bool {
	if o != nil && o.PareqFieldValue != nil {
		return true
	}

	return false
}

// SetPareqFieldValue gets a reference to the given string and assigns it to the PareqFieldValue field.
func (o *ACS) SetPareqFieldValue(v string) {
	o.PareqFieldValue = &v
}

// GetTermurlFieldName returns the TermurlFieldName field value if set, zero value otherwise.
func (o *ACS) GetTermurlFieldName() string {
	if o == nil || o.TermurlFieldName == nil {
		var ret string
		return ret
	}
	return *o.TermurlFieldName
}

// GetTermurlFieldNameOk returns a tuple with the TermurlFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACS) GetTermurlFieldNameOk() (*string, bool) {
	if o == nil || o.TermurlFieldName == nil {
		return nil, false
	}
	return o.TermurlFieldName, true
}

// HasTermurlFieldName returns a boolean if a field has been set.
func (o *ACS) HasTermurlFieldName() bool {
	if o != nil && o.TermurlFieldName != nil {
		return true
	}

	return false
}

// SetTermurlFieldName gets a reference to the given string and assigns it to the TermurlFieldName field.
func (o *ACS) SetTermurlFieldName(v string) {
	o.TermurlFieldName = &v
}

// GetMdFieldName returns the MdFieldName field value if set, zero value otherwise.
func (o *ACS) GetMdFieldName() string {
	if o == nil || o.MdFieldName == nil {
		var ret string
		return ret
	}
	return *o.MdFieldName
}

// GetMdFieldNameOk returns a tuple with the MdFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACS) GetMdFieldNameOk() (*string, bool) {
	if o == nil || o.MdFieldName == nil {
		return nil, false
	}
	return o.MdFieldName, true
}

// HasMdFieldName returns a boolean if a field has been set.
func (o *ACS) HasMdFieldName() bool {
	if o != nil && o.MdFieldName != nil {
		return true
	}

	return false
}

// SetMdFieldName gets a reference to the given string and assigns it to the MdFieldName field.
func (o *ACS) SetMdFieldName(v string) {
	o.MdFieldName = &v
}

// GetMdFieldValue returns the MdFieldValue field value if set, zero value otherwise.
func (o *ACS) GetMdFieldValue() string {
	if o == nil || o.MdFieldValue == nil {
		var ret string
		return ret
	}
	return *o.MdFieldValue
}

// GetMdFieldValueOk returns a tuple with the MdFieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACS) GetMdFieldValueOk() (*string, bool) {
	if o == nil || o.MdFieldValue == nil {
		return nil, false
	}
	return o.MdFieldValue, true
}

// HasMdFieldValue returns a boolean if a field has been set.
func (o *ACS) HasMdFieldValue() bool {
	if o != nil && o.MdFieldValue != nil {
		return true
	}

	return false
}

// SetMdFieldValue gets a reference to the given string and assigns it to the MdFieldValue field.
func (o *ACS) SetMdFieldValue(v string) {
	o.MdFieldValue = &v
}

// GetMpiResult returns the MpiResult field value if set, zero value otherwise.
func (o *ACS) GetMpiResult() string {
	if o == nil || o.MpiResult == nil {
		var ret string
		return ret
	}
	return *o.MpiResult
}

// GetMpiResultOk returns a tuple with the MpiResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACS) GetMpiResultOk() (*string, bool) {
	if o == nil || o.MpiResult == nil {
		return nil, false
	}
	return o.MpiResult, true
}

// HasMpiResult returns a boolean if a field has been set.
func (o *ACS) HasMpiResult() bool {
	if o != nil && o.MpiResult != nil {
		return true
	}

	return false
}

// SetMpiResult gets a reference to the given string and assigns it to the MpiResult field.
func (o *ACS) SetMpiResult(v string) {
	o.MpiResult = &v
}

func (o ACS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionUrl != nil {
		toSerialize["actionUrl"] = o.ActionUrl
	}
	if o.ActionMethod != nil {
		toSerialize["actionMethod"] = o.ActionMethod
	}
	if o.PareqFieldName != nil {
		toSerialize["pareqFieldName"] = o.PareqFieldName
	}
	if o.PareqFieldValue != nil {
		toSerialize["pareqFieldValue"] = o.PareqFieldValue
	}
	if o.TermurlFieldName != nil {
		toSerialize["termurlFieldName"] = o.TermurlFieldName
	}
	if o.MdFieldName != nil {
		toSerialize["mdFieldName"] = o.MdFieldName
	}
	if o.MdFieldValue != nil {
		toSerialize["mdFieldValue"] = o.MdFieldValue
	}
	if o.MpiResult != nil {
		toSerialize["mpiResult"] = o.MpiResult
	}
	return json.Marshal(toSerialize)
}

type NullableACS struct {
	value *ACS
	isSet bool
}

func (v NullableACS) Get() *ACS {
	return v.value
}

func (v *NullableACS) Set(val *ACS) {
	v.value = val
	v.isSet = true
}

func (v NullableACS) IsSet() bool {
	return v.isSet
}

func (v *NullableACS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACS(val *ACS) *NullableACS {
	return &NullableACS{value: val, isSet: true}
}

func (v NullableACS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


