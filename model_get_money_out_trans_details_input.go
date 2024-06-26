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

// GetMoneyOutTransDetailsInput struct for GetMoneyOutTransDetailsInput
type GetMoneyOutTransDetailsInput struct {
	// Money-Out ID
	Transactionid *int64 `json:"transactionid,omitempty"`
	// Money-Out Comment
	TransactionComment *string `json:"transactionComment,omitempty"`
	// Unique ID Generated by your Server
	Reference *string `json:"reference,omitempty"`
}

// NewGetMoneyOutTransDetailsInput instantiates a new GetMoneyOutTransDetailsInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMoneyOutTransDetailsInput() *GetMoneyOutTransDetailsInput {
	this := GetMoneyOutTransDetailsInput{}
	return &this
}

// NewGetMoneyOutTransDetailsInputWithDefaults instantiates a new GetMoneyOutTransDetailsInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMoneyOutTransDetailsInputWithDefaults() *GetMoneyOutTransDetailsInput {
	this := GetMoneyOutTransDetailsInput{}
	return &this
}

// GetTransactionid returns the Transactionid field value if set, zero value otherwise.
func (o *GetMoneyOutTransDetailsInput) GetTransactionid() int64 {
	if o == nil || o.Transactionid == nil {
		var ret int64
		return ret
	}
	return *o.Transactionid
}

// GetTransactionidOk returns a tuple with the Transactionid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyOutTransDetailsInput) GetTransactionidOk() (*int64, bool) {
	if o == nil || o.Transactionid == nil {
		return nil, false
	}
	return o.Transactionid, true
}

// HasTransactionid returns a boolean if a field has been set.
func (o *GetMoneyOutTransDetailsInput) HasTransactionid() bool {
	if o != nil && o.Transactionid != nil {
		return true
	}

	return false
}

// SetTransactionid gets a reference to the given int64 and assigns it to the Transactionid field.
func (o *GetMoneyOutTransDetailsInput) SetTransactionid(v int64) {
	o.Transactionid = &v
}

// GetTransactionComment returns the TransactionComment field value if set, zero value otherwise.
func (o *GetMoneyOutTransDetailsInput) GetTransactionComment() string {
	if o == nil || o.TransactionComment == nil {
		var ret string
		return ret
	}
	return *o.TransactionComment
}

// GetTransactionCommentOk returns a tuple with the TransactionComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyOutTransDetailsInput) GetTransactionCommentOk() (*string, bool) {
	if o == nil || o.TransactionComment == nil {
		return nil, false
	}
	return o.TransactionComment, true
}

// HasTransactionComment returns a boolean if a field has been set.
func (o *GetMoneyOutTransDetailsInput) HasTransactionComment() bool {
	if o != nil && o.TransactionComment != nil {
		return true
	}

	return false
}

// SetTransactionComment gets a reference to the given string and assigns it to the TransactionComment field.
func (o *GetMoneyOutTransDetailsInput) SetTransactionComment(v string) {
	o.TransactionComment = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GetMoneyOutTransDetailsInput) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMoneyOutTransDetailsInput) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GetMoneyOutTransDetailsInput) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GetMoneyOutTransDetailsInput) SetReference(v string) {
	o.Reference = &v
}

func (o GetMoneyOutTransDetailsInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transactionid != nil {
		toSerialize["transactionid"] = o.Transactionid
	}
	if o.TransactionComment != nil {
		toSerialize["transactionComment"] = o.TransactionComment
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	return json.Marshal(toSerialize)
}

type NullableGetMoneyOutTransDetailsInput struct {
	value *GetMoneyOutTransDetailsInput
	isSet bool
}

func (v NullableGetMoneyOutTransDetailsInput) Get() *GetMoneyOutTransDetailsInput {
	return v.value
}

func (v *NullableGetMoneyOutTransDetailsInput) Set(val *GetMoneyOutTransDetailsInput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMoneyOutTransDetailsInput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMoneyOutTransDetailsInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMoneyOutTransDetailsInput(val *GetMoneyOutTransDetailsInput) *NullableGetMoneyOutTransDetailsInput {
	return &NullableGetMoneyOutTransDetailsInput{value: val, isSet: true}
}

func (v NullableGetMoneyOutTransDetailsInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMoneyOutTransDetailsInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


