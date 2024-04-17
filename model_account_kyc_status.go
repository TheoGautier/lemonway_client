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

// AccountKycStatus struct for AccountKycStatus
type AccountKycStatus struct {
	// Payment account ID
	Id *string `json:"id,omitempty"`
	// Payment account status<br/>2 = Registered, KYC incomplete.<br/>3 = Registered, rejected KYC.<br/>5 = Registered, KYC 1 (status given at registration).<br/>6 = Registered, KYC 2.<br/>7 = Registered, KYC 3.<br/>8 = Registered, expired KYC.<br/>10 = Blocked.<br/>12 = Closed.<br/>13 = Registered, status is being updated from KYC 2 to KYC 3.<br/>14 = One-time customer.<br/>15 = Special account for crowdlending.<br/>16 = Technical account.<br/>
	Status *int32 `json:"status,omitempty"`
	// Modification date of the payment account status in Second UTC
	Date *string `json:"date,omitempty"`
	Documents []Document `json:"documents,omitempty"`
	Ibans []Iban `json:"ibans,omitempty"`
	SddMandates []SddMandate `json:"sddMandates,omitempty"`
}

// NewAccountKycStatus instantiates a new AccountKycStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountKycStatus() *AccountKycStatus {
	this := AccountKycStatus{}
	return &this
}

// NewAccountKycStatusWithDefaults instantiates a new AccountKycStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountKycStatusWithDefaults() *AccountKycStatus {
	this := AccountKycStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountKycStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatus) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountKycStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountKycStatus) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountKycStatus) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatus) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountKycStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AccountKycStatus) SetStatus(v int32) {
	o.Status = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AccountKycStatus) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatus) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AccountKycStatus) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *AccountKycStatus) SetDate(v string) {
	o.Date = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *AccountKycStatus) GetDocuments() []Document {
	if o == nil || o.Documents == nil {
		var ret []Document
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatus) GetDocumentsOk() ([]Document, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *AccountKycStatus) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []Document and assigns it to the Documents field.
func (o *AccountKycStatus) SetDocuments(v []Document) {
	o.Documents = v
}

// GetIbans returns the Ibans field value if set, zero value otherwise.
func (o *AccountKycStatus) GetIbans() []Iban {
	if o == nil || o.Ibans == nil {
		var ret []Iban
		return ret
	}
	return o.Ibans
}

// GetIbansOk returns a tuple with the Ibans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatus) GetIbansOk() ([]Iban, bool) {
	if o == nil || o.Ibans == nil {
		return nil, false
	}
	return o.Ibans, true
}

// HasIbans returns a boolean if a field has been set.
func (o *AccountKycStatus) HasIbans() bool {
	if o != nil && o.Ibans != nil {
		return true
	}

	return false
}

// SetIbans gets a reference to the given []Iban and assigns it to the Ibans field.
func (o *AccountKycStatus) SetIbans(v []Iban) {
	o.Ibans = v
}

// GetSddMandates returns the SddMandates field value if set, zero value otherwise.
func (o *AccountKycStatus) GetSddMandates() []SddMandate {
	if o == nil || o.SddMandates == nil {
		var ret []SddMandate
		return ret
	}
	return o.SddMandates
}

// GetSddMandatesOk returns a tuple with the SddMandates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountKycStatus) GetSddMandatesOk() ([]SddMandate, bool) {
	if o == nil || o.SddMandates == nil {
		return nil, false
	}
	return o.SddMandates, true
}

// HasSddMandates returns a boolean if a field has been set.
func (o *AccountKycStatus) HasSddMandates() bool {
	if o != nil && o.SddMandates != nil {
		return true
	}

	return false
}

// SetSddMandates gets a reference to the given []SddMandate and assigns it to the SddMandates field.
func (o *AccountKycStatus) SetSddMandates(v []SddMandate) {
	o.SddMandates = v
}

func (o AccountKycStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.Ibans != nil {
		toSerialize["ibans"] = o.Ibans
	}
	if o.SddMandates != nil {
		toSerialize["sddMandates"] = o.SddMandates
	}
	return json.Marshal(toSerialize)
}

type NullableAccountKycStatus struct {
	value *AccountKycStatus
	isSet bool
}

func (v NullableAccountKycStatus) Get() *AccountKycStatus {
	return v.value
}

func (v *NullableAccountKycStatus) Set(val *AccountKycStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountKycStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountKycStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountKycStatus(val *AccountKycStatus) *NullableAccountKycStatus {
	return &NullableAccountKycStatus{value: val, isSet: true}
}

func (v NullableAccountKycStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountKycStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

