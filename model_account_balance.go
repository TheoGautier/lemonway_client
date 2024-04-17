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

// AccountBalance struct for AccountBalance
type AccountBalance struct {
	// Payment account ID
	Id *string `json:"id,omitempty"`
	// Payment account balance
	Balance *int32 `json:"balance,omitempty"`
	// Payment account status<br/>2 = Registered, KYC incomplete.<br/>3 = Registered, rejected KYC.<br/>5 = Registered, KYC 1 (status given at registration).<br/>6 = Registered, KYC 2.<br/>7 = Registered, KYC 3.<br/>8 = Registered, expired KYC.<br/>10 = Blocked.<br/>12 = Closed.<br/>13 = Registered, status is being updated from KYC 2 to KYC 3.<br/>14 = One-time customer.<br/>15 = Special account for crowdlending.<br/>16 = Technical account.<br/>
	Status *int32 `json:"status,omitempty"`
}

// NewAccountBalance instantiates a new AccountBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBalance() *AccountBalance {
	this := AccountBalance{}
	return &this
}

// NewAccountBalanceWithDefaults instantiates a new AccountBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBalanceWithDefaults() *AccountBalance {
	this := AccountBalance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountBalance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountBalance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountBalance) SetId(v string) {
	o.Id = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *AccountBalance) GetBalance() int32 {
	if o == nil || o.Balance == nil {
		var ret int32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalance) GetBalanceOk() (*int32, bool) {
	if o == nil || o.Balance == nil {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *AccountBalance) HasBalance() bool {
	if o != nil && o.Balance != nil {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int32 and assigns it to the Balance field.
func (o *AccountBalance) SetBalance(v int32) {
	o.Balance = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountBalance) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBalance) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountBalance) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AccountBalance) SetStatus(v int32) {
	o.Status = &v
}

func (o AccountBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Balance != nil {
		toSerialize["balance"] = o.Balance
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAccountBalance struct {
	value *AccountBalance
	isSet bool
}

func (v NullableAccountBalance) Get() *AccountBalance {
	return v.value
}

func (v *NullableAccountBalance) Set(val *AccountBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBalance(val *AccountBalance) *NullableAccountBalance {
	return &NullableAccountBalance{value: val, isSet: true}
}

func (v NullableAccountBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

