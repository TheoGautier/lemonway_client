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

// CancelMoneyInInput struct for CancelMoneyInInput
type CancelMoneyInInput struct {
	// Payment Account ID
	Account string `json:"account"`
}

// NewCancelMoneyInInput instantiates a new CancelMoneyInInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelMoneyInInput(account string) *CancelMoneyInInput {
	this := CancelMoneyInInput{}
	this.Account = account
	return &this
}

// NewCancelMoneyInInputWithDefaults instantiates a new CancelMoneyInInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelMoneyInInputWithDefaults() *CancelMoneyInInput {
	this := CancelMoneyInInput{}
	return &this
}

// GetAccount returns the Account field value
func (o *CancelMoneyInInput) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *CancelMoneyInInput) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *CancelMoneyInInput) SetAccount(v string) {
	o.Account = v
}

func (o CancelMoneyInInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableCancelMoneyInInput struct {
	value *CancelMoneyInInput
	isSet bool
}

func (v NullableCancelMoneyInInput) Get() *CancelMoneyInInput {
	return v.value
}

func (v *NullableCancelMoneyInInput) Set(val *CancelMoneyInInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelMoneyInInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelMoneyInInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelMoneyInInput(val *CancelMoneyInInput) *NullableCancelMoneyInInput {
	return &NullableCancelMoneyInInput{value: val, isSet: true}
}

func (v NullableCancelMoneyInInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelMoneyInInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

