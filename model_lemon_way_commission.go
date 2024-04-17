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

// LemonWayCommission struct for LemonWayCommission
type LemonWayCommission struct {
	Idp2p *string `json:"idp2p,omitempty"`
	// Amounts are represented as integer in cents (Euro)  Represented as an integer in cents (Euro)
	Amount *int32 `json:"amount,omitempty"`
}

// NewLemonWayCommission instantiates a new LemonWayCommission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLemonWayCommission() *LemonWayCommission {
	this := LemonWayCommission{}
	return &this
}

// NewLemonWayCommissionWithDefaults instantiates a new LemonWayCommission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLemonWayCommissionWithDefaults() *LemonWayCommission {
	this := LemonWayCommission{}
	return &this
}

// GetIdp2p returns the Idp2p field value if set, zero value otherwise.
func (o *LemonWayCommission) GetIdp2p() string {
	if o == nil || o.Idp2p == nil {
		var ret string
		return ret
	}
	return *o.Idp2p
}

// GetIdp2pOk returns a tuple with the Idp2p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LemonWayCommission) GetIdp2pOk() (*string, bool) {
	if o == nil || o.Idp2p == nil {
		return nil, false
	}
	return o.Idp2p, true
}

// HasIdp2p returns a boolean if a field has been set.
func (o *LemonWayCommission) HasIdp2p() bool {
	if o != nil && o.Idp2p != nil {
		return true
	}

	return false
}

// SetIdp2p gets a reference to the given string and assigns it to the Idp2p field.
func (o *LemonWayCommission) SetIdp2p(v string) {
	o.Idp2p = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *LemonWayCommission) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LemonWayCommission) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *LemonWayCommission) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *LemonWayCommission) SetAmount(v int32) {
	o.Amount = &v
}

func (o LemonWayCommission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Idp2p != nil {
		toSerialize["idp2p"] = o.Idp2p
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableLemonWayCommission struct {
	value *LemonWayCommission
	isSet bool
}

func (v NullableLemonWayCommission) Get() *LemonWayCommission {
	return v.value
}

func (v *NullableLemonWayCommission) Set(val *LemonWayCommission) {
	v.value = val
	v.isSet = true
}

func (v NullableLemonWayCommission) IsSet() bool {
	return v.isSet
}

func (v *NullableLemonWayCommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLemonWayCommission(val *LemonWayCommission) *NullableLemonWayCommission {
	return &NullableLemonWayCommission{value: val, isSet: true}
}

func (v NullableLemonWayCommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLemonWayCommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

