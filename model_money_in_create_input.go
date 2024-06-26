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

// MoneyInCreateInput struct for MoneyInCreateInput
type MoneyInCreateInput struct {
	Card CardInfoExtended `json:"card"`
	Transaction TransactionInfo `json:"transaction"`
	// Service Provider
	ServiceProvider *string `json:"ServiceProvider,omitempty"`
	// Payment Account ID to Credit
	AccountId string `json:"accountId"`
	// Amount to Debit  Amounts are given as integer numbers in cents
	TotalAmount *int32 `json:"totalAmount,omitempty"`
	// Your Fee  Amounts are given as integer numbers in cents
	CommissionAmount *int32 `json:"commissionAmount,omitempty"`
	// Comment Regarding the Transaction
	Comment *string `json:"comment,omitempty"`
	// If true:  1. [amountCom] will be ignored and will be replaced with Lemonway's fee  2. You will not receive any fee
	AutoCommission *bool `json:"autoCommission,omitempty"`
}

// NewMoneyInCreateInput instantiates a new MoneyInCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyInCreateInput(card CardInfoExtended, transaction TransactionInfo, accountId string) *MoneyInCreateInput {
	this := MoneyInCreateInput{}
	this.Card = card
	this.Transaction = transaction
	this.AccountId = accountId
	return &this
}

// NewMoneyInCreateInputWithDefaults instantiates a new MoneyInCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyInCreateInputWithDefaults() *MoneyInCreateInput {
	this := MoneyInCreateInput{}
	return &this
}

// GetCard returns the Card field value
func (o *MoneyInCreateInput) GetCard() CardInfoExtended {
	if o == nil {
		var ret CardInfoExtended
		return ret
	}

	return o.Card
}

// GetCardOk returns a tuple with the Card field value
// and a boolean to check if the value has been set.
func (o *MoneyInCreateInput) GetCardOk() (*CardInfoExtended, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Card, true
}

// SetCard sets field value
func (o *MoneyInCreateInput) SetCard(v CardInfoExtended) {
	o.Card = v
}

// GetTransaction returns the Transaction field value
func (o *MoneyInCreateInput) GetTransaction() TransactionInfo {
	if o == nil {
		var ret TransactionInfo
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *MoneyInCreateInput) GetTransactionOk() (*TransactionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *MoneyInCreateInput) SetTransaction(v TransactionInfo) {
	o.Transaction = v
}

// GetServiceProvider returns the ServiceProvider field value if set, zero value otherwise.
func (o *MoneyInCreateInput) GetServiceProvider() string {
	if o == nil || o.ServiceProvider == nil {
		var ret string
		return ret
	}
	return *o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInCreateInput) GetServiceProviderOk() (*string, bool) {
	if o == nil || o.ServiceProvider == nil {
		return nil, false
	}
	return o.ServiceProvider, true
}

// HasServiceProvider returns a boolean if a field has been set.
func (o *MoneyInCreateInput) HasServiceProvider() bool {
	if o != nil && o.ServiceProvider != nil {
		return true
	}

	return false
}

// SetServiceProvider gets a reference to the given string and assigns it to the ServiceProvider field.
func (o *MoneyInCreateInput) SetServiceProvider(v string) {
	o.ServiceProvider = &v
}

// GetAccountId returns the AccountId field value
func (o *MoneyInCreateInput) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MoneyInCreateInput) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MoneyInCreateInput) SetAccountId(v string) {
	o.AccountId = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MoneyInCreateInput) GetTotalAmount() int32 {
	if o == nil || o.TotalAmount == nil {
		var ret int32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInCreateInput) GetTotalAmountOk() (*int32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MoneyInCreateInput) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int32 and assigns it to the TotalAmount field.
func (o *MoneyInCreateInput) SetTotalAmount(v int32) {
	o.TotalAmount = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *MoneyInCreateInput) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInCreateInput) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *MoneyInCreateInput) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *MoneyInCreateInput) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MoneyInCreateInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInCreateInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MoneyInCreateInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MoneyInCreateInput) SetComment(v string) {
	o.Comment = &v
}

// GetAutoCommission returns the AutoCommission field value if set, zero value otherwise.
func (o *MoneyInCreateInput) GetAutoCommission() bool {
	if o == nil || o.AutoCommission == nil {
		var ret bool
		return ret
	}
	return *o.AutoCommission
}

// GetAutoCommissionOk returns a tuple with the AutoCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyInCreateInput) GetAutoCommissionOk() (*bool, bool) {
	if o == nil || o.AutoCommission == nil {
		return nil, false
	}
	return o.AutoCommission, true
}

// HasAutoCommission returns a boolean if a field has been set.
func (o *MoneyInCreateInput) HasAutoCommission() bool {
	if o != nil && o.AutoCommission != nil {
		return true
	}

	return false
}

// SetAutoCommission gets a reference to the given bool and assigns it to the AutoCommission field.
func (o *MoneyInCreateInput) SetAutoCommission(v bool) {
	o.AutoCommission = &v
}

func (o MoneyInCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card"] = o.Card
	}
	if true {
		toSerialize["transaction"] = o.Transaction
	}
	if o.ServiceProvider != nil {
		toSerialize["ServiceProvider"] = o.ServiceProvider
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if o.TotalAmount != nil {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if o.CommissionAmount != nil {
		toSerialize["commissionAmount"] = o.CommissionAmount
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.AutoCommission != nil {
		toSerialize["autoCommission"] = o.AutoCommission
	}
	return json.Marshal(toSerialize)
}

type NullableMoneyInCreateInput struct {
	value *MoneyInCreateInput
	isSet bool
}

func (v NullableMoneyInCreateInput) Get() *MoneyInCreateInput {
	return v.value
}

func (v *NullableMoneyInCreateInput) Set(val *MoneyInCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyInCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyInCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyInCreateInput(val *MoneyInCreateInput) *NullableMoneyInCreateInput {
	return &NullableMoneyInCreateInput{value: val, isSet: true}
}

func (v NullableMoneyInCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyInCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


