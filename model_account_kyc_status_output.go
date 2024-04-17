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

// AccountKycStatusOutput struct for AccountKycStatusOutput
type AccountKycStatusOutput struct {
	Transactions *TransactionsTransactionAccount `json:"transactions,omitempty"`
	Links *Links `json:"_links,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewAccountKycStatusOutput instantiates a new AccountKycStatusOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountKycStatusOutput() *AccountKycStatusOutput {
	this := AccountKycStatusOutput{}
	return &this
}

// NewAccountKycStatusOutputWithDefaults instantiates a new AccountKycStatusOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountKycStatusOutputWithDefaults() *AccountKycStatusOutput {
	this := AccountKycStatusOutput{}
	return &this
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *AccountKycStatusOutput) GetTransactions() TransactionsTransactionAccount {
	if o == nil || o.Transactions == nil {
		var ret TransactionsTransactionAccount
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatusOutput) GetTransactionsOk() (*TransactionsTransactionAccount, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *AccountKycStatusOutput) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given TransactionsTransactionAccount and assigns it to the Transactions field.
func (o *AccountKycStatusOutput) SetTransactions(v TransactionsTransactionAccount) {
	o.Transactions = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AccountKycStatusOutput) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatusOutput) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AccountKycStatusOutput) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *AccountKycStatusOutput) SetLinks(v Links) {
	o.Links = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AccountKycStatusOutput) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatusOutput) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AccountKycStatusOutput) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *AccountKycStatusOutput) SetError(v Error) {
	o.Error = &v
}

func (o AccountKycStatusOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableAccountKycStatusOutput struct {
	value *AccountKycStatusOutput
	isSet bool
}

func (v NullableAccountKycStatusOutput) Get() *AccountKycStatusOutput {
	return v.value
}

func (v *NullableAccountKycStatusOutput) Set(val *AccountKycStatusOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountKycStatusOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountKycStatusOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountKycStatusOutput(val *AccountKycStatusOutput) *NullableAccountKycStatusOutput {
	return &NullableAccountKycStatusOutput{value: val, isSet: true}
}

func (v NullableAccountKycStatusOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountKycStatusOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

