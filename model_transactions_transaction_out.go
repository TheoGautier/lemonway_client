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

// TransactionsTransactionOut struct for TransactionsTransactionOut
type TransactionsTransactionOut struct {
	Value []TransactionOut `json:"value,omitempty"`
}

// NewTransactionsTransactionOut instantiates a new TransactionsTransactionOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsTransactionOut() *TransactionsTransactionOut {
	this := TransactionsTransactionOut{}
	return &this
}

// NewTransactionsTransactionOutWithDefaults instantiates a new TransactionsTransactionOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsTransactionOutWithDefaults() *TransactionsTransactionOut {
	this := TransactionsTransactionOut{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TransactionsTransactionOut) GetValue() []TransactionOut {
	if o == nil || o.Value == nil {
		var ret []TransactionOut
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsTransactionOut) GetValueOk() ([]TransactionOut, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TransactionsTransactionOut) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []TransactionOut and assigns it to the Value field.
func (o *TransactionsTransactionOut) SetValue(v []TransactionOut) {
	o.Value = v
}

func (o TransactionsTransactionOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsTransactionOut struct {
	value *TransactionsTransactionOut
	isSet bool
}

func (v NullableTransactionsTransactionOut) Get() *TransactionsTransactionOut {
	return v.value
}

func (v *NullableTransactionsTransactionOut) Set(val *TransactionsTransactionOut) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsTransactionOut) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsTransactionOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsTransactionOut(val *TransactionsTransactionOut) *NullableTransactionsTransactionOut {
	return &NullableTransactionsTransactionOut{value: val, isSet: true}
}

func (v NullableTransactionsTransactionOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsTransactionOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

