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

// MoneyInTransferInitOutput struct for MoneyInTransferInitOutput
type MoneyInTransferInitOutput struct {
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewMoneyInTransferInitOutput instantiates a new MoneyInTransferInitOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInTransferInitOutput() *MoneyInTransferInitOutput {
	this := MoneyInTransferInitOutput{}
	return &this
}

// NewMoneyInTransferInitOutputWithDefaults instantiates a new MoneyInTransferInitOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInTransferInitOutputWithDefaults() *MoneyInTransferInitOutput {
	this := MoneyInTransferInitOutput{}
	return &this
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *MoneyInTransferInitOutput) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInTransferInitOutput) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *MoneyInTransferInitOutput) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *MoneyInTransferInitOutput) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MoneyInTransferInitOutput) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInTransferInitOutput) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MoneyInTransferInitOutput) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MoneyInTransferInitOutput) SetId(v int64) {
	o.Id = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MoneyInTransferInitOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInTransferInitOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MoneyInTransferInitOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *MoneyInTransferInitOutput) SetError(v Error) {
	o.Error = &v
}

func (o MoneyInTransferInitOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyInTransferInitOutput struct {
	value *MoneyInTransferInitOutput
	isSet bool
}

func (v NullableMoneyInTransferInitOutput) Get() *MoneyInTransferInitOutput {
	return v.value
}

func (v *NullableMoneyInTransferInitOutput) Set(val *MoneyInTransferInitOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInTransferInitOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInTransferInitOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInTransferInitOutput(val *MoneyInTransferInitOutput) *NullableMoneyInTransferInitOutput {
	return &NullableMoneyInTransferInitOutput{value: val, isSet: true}
}

func (v NullableMoneyInTransferInitOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInTransferInitOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

