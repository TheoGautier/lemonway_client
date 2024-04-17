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

// Iban struct for Iban
type Iban struct {
	// IBAN ID
	Id *int64 `json:"id,omitempty"`
	// Only status 5 allows the use of the IBAN for a Money-Out.<br/>1 = None.<br/>2 = Internal.<br/>3 = Not used.<br/>4 = Waiting to be verified by Lemon Way.<br/>5 = Activated.<br/>6 = Rejected by the bank.<br/>7 = Rejected, no owner.<br/>8 = Deactivated.<br/>9 = Rejected.<br/>
	Status *int32 `json:"status,omitempty"`
	// IBAN saved
	Iban *string `json:"iban,omitempty"`
	// SWIFT BIC code associated with the IBAN
	Swift *string `json:"swift,omitempty"`
	// IBAN Holder
	Holder *string `json:"holder,omitempty"`
	// Indicates if a merchant IBAN or a virtual client IBAN  1: Merchant IBAN  2: IBAN virtual client
	Type *int32 `json:"type,omitempty"`
}

// NewIban instantiates a new Iban object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIban() *Iban {
	this := Iban{}
	return &this
}

// NewIbanWithDefaults instantiates a new Iban object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIbanWithDefaults() *Iban {
	this := Iban{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Iban) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iban) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Iban) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Iban) SetId(v int64) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Iban) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iban) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Iban) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *Iban) SetStatus(v int32) {
	o.Status = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *Iban) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iban) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *Iban) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *Iban) SetIban(v string) {
	o.Iban = &v
}

// GetSwift returns the Swift field value if set, zero value otherwise.
func (o *Iban) GetSwift() string {
	if o == nil || o.Swift == nil {
		var ret string
		return ret
	}
	return *o.Swift
}

// GetSwiftOk returns a tuple with the Swift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iban) GetSwiftOk() (*string, bool) {
	if o == nil || o.Swift == nil {
		return nil, false
	}
	return o.Swift, true
}

// HasSwift returns a boolean if a field has been set.
func (o *Iban) HasSwift() bool {
	if o != nil && o.Swift != nil {
		return true
	}

	return false
}

// SetSwift gets a reference to the given string and assigns it to the Swift field.
func (o *Iban) SetSwift(v string) {
	o.Swift = &v
}

// GetHolder returns the Holder field value if set, zero value otherwise.
func (o *Iban) GetHolder() string {
	if o == nil || o.Holder == nil {
		var ret string
		return ret
	}
	return *o.Holder
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iban) GetHolderOk() (*string, bool) {
	if o == nil || o.Holder == nil {
		return nil, false
	}
	return o.Holder, true
}

// HasHolder returns a boolean if a field has been set.
func (o *Iban) HasHolder() bool {
	if o != nil && o.Holder != nil {
		return true
	}

	return false
}

// SetHolder gets a reference to the given string and assigns it to the Holder field.
func (o *Iban) SetHolder(v string) {
	o.Holder = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Iban) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iban) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Iban) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Iban) SetType(v int32) {
	o.Type = &v
}

func (o Iban) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if o.Swift != nil {
		toSerialize["swift"] = o.Swift
	}
	if o.Holder != nil {
		toSerialize["holder"] = o.Holder
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIban struct {
	value *Iban
	isSet bool
}

func (v NullableIban) Get() *Iban {
	return v.value
}

func (v *NullableIban) Set(val *Iban) {
	v.value = val
	v.isSet = true
}

func (v NullableIban) IsSet() bool {
	return v.isSet
}

func (v *NullableIban) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIban(val *Iban) *NullableIban {
	return &NullableIban{value: val, isSet: true}
}

func (v NullableIban) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIban) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

