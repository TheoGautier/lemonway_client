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

// BnplPurchaseItem struct for BnplPurchaseItem
type BnplPurchaseItem struct {
	Type string `json:"type"`
	Description string `json:"description"`
	Reference string `json:"reference"`
	Quantity int32 `json:"quantity"`
	UnitAmount int32 `json:"unitAmount"`
	MerchantAccountId string `json:"merchantAccountId"`
}

// NewBnplPurchaseItem instantiates a new BnplPurchaseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnplPurchaseItem(type_ string, description string, reference string, quantity int32, unitAmount int32, merchantAccountId string) *BnplPurchaseItem {
	this := BnplPurchaseItem{}
	this.Type = type_
	this.Description = description
	this.Reference = reference
	this.Quantity = quantity
	this.UnitAmount = unitAmount
	this.MerchantAccountId = merchantAccountId
	return &this
}

// NewBnplPurchaseItemWithDefaults instantiates a new BnplPurchaseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnplPurchaseItemWithDefaults() *BnplPurchaseItem {
	this := BnplPurchaseItem{}
	return &this
}

// GetType returns the Type field value
func (o *BnplPurchaseItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BnplPurchaseItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BnplPurchaseItem) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *BnplPurchaseItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BnplPurchaseItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BnplPurchaseItem) SetDescription(v string) {
	o.Description = v
}

// GetReference returns the Reference field value
func (o *BnplPurchaseItem) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *BnplPurchaseItem) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *BnplPurchaseItem) SetReference(v string) {
	o.Reference = v
}

// GetQuantity returns the Quantity field value
func (o *BnplPurchaseItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *BnplPurchaseItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *BnplPurchaseItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetUnitAmount returns the UnitAmount field value
func (o *BnplPurchaseItem) GetUnitAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value
// and a boolean to check if the value has been set.
func (o *BnplPurchaseItem) GetUnitAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitAmount, true
}

// SetUnitAmount sets field value
func (o *BnplPurchaseItem) SetUnitAmount(v int32) {
	o.UnitAmount = v
}

// GetMerchantAccountId returns the MerchantAccountId field value
func (o *BnplPurchaseItem) GetMerchantAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccountId
}

// GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field value
// and a boolean to check if the value has been set.
func (o *BnplPurchaseItem) GetMerchantAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccountId, true
}

// SetMerchantAccountId sets field value
func (o *BnplPurchaseItem) SetMerchantAccountId(v string) {
	o.MerchantAccountId = v
}

func (o BnplPurchaseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["unitAmount"] = o.UnitAmount
	}
	if true {
		toSerialize["merchantAccountId"] = o.MerchantAccountId
	}
	return json.Marshal(toSerialize)
}

type NullableBnplPurchaseItem struct {
	value *BnplPurchaseItem
	isSet bool
}

func (v NullableBnplPurchaseItem) Get() *BnplPurchaseItem {
	return v.value
}

func (v *NullableBnplPurchaseItem) Set(val *BnplPurchaseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBnplPurchaseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBnplPurchaseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnplPurchaseItem(val *BnplPurchaseItem) *NullableBnplPurchaseItem {
	return &NullableBnplPurchaseItem{value: val, isSet: true}
}

func (v NullableBnplPurchaseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnplPurchaseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


