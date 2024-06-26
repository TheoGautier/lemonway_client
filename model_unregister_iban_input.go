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

// UnregisterIBANInput struct for UnregisterIBANInput
type UnregisterIBANInput struct {
	// Payment Account ID
	Wallet string `json:"wallet"`
}

// NewUnregisterIBANInput instantiates a new UnregisterIBANInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnregisterIBANInput(wallet string) *UnregisterIBANInput {
	this := UnregisterIBANInput{}
	this.Wallet = wallet
	return &this
}

// NewUnregisterIBANInputWithDefaults instantiates a new UnregisterIBANInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnregisterIBANInputWithDefaults() *UnregisterIBANInput {
	this := UnregisterIBANInput{}
	return &this
}

// GetWallet returns the Wallet field value
func (o *UnregisterIBANInput) GetWallet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value
// and a boolean to check if the value has been set.
func (o *UnregisterIBANInput) GetWalletOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallet, true
}

// SetWallet sets field value
func (o *UnregisterIBANInput) SetWallet(v string) {
	o.Wallet = v
}

func (o UnregisterIBANInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wallet"] = o.Wallet
	}
	return json.Marshal(toSerialize)
}

type NullableUnregisterIBANInput struct {
	value *UnregisterIBANInput
	isSet bool
}

func (v NullableUnregisterIBANInput) Get() *UnregisterIBANInput {
	return v.value
}

func (v *NullableUnregisterIBANInput) Set(val *UnregisterIBANInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUnregisterIBANInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUnregisterIBANInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnregisterIBANInput(val *UnregisterIBANInput) *NullableUnregisterIBANInput {
	return &NullableUnregisterIBANInput{value: val, isSet: true}
}

func (v NullableUnregisterIBANInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnregisterIBANInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


