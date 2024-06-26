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

// UpdateAccountStatus struct for UpdateAccountStatus
type UpdateAccountStatus struct {
	// Updated payment account ID
	Id *string `json:"id,omitempty"`
	// Status of the payment account
	Status *int32 `json:"status,omitempty"`
}

// NewUpdateAccountStatus instantiates a new UpdateAccountStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountStatus() *UpdateAccountStatus {
	this := UpdateAccountStatus{}
	return &this
}

// NewUpdateAccountStatusWithDefaults instantiates a new UpdateAccountStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountStatusWithDefaults() *UpdateAccountStatus {
	this := UpdateAccountStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateAccountStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountStatus) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateAccountStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateAccountStatus) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateAccountStatus) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountStatus) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateAccountStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UpdateAccountStatus) SetStatus(v int32) {
	o.Status = &v
}

func (o UpdateAccountStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAccountStatus struct {
	value *UpdateAccountStatus
	isSet bool
}

func (v NullableUpdateAccountStatus) Get() *UpdateAccountStatus {
	return v.value
}

func (v *NullableUpdateAccountStatus) Set(val *UpdateAccountStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountStatus(val *UpdateAccountStatus) *NullableUpdateAccountStatus {
	return &NullableUpdateAccountStatus{value: val, isSet: true}
}

func (v NullableUpdateAccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


