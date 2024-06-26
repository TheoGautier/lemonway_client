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

// MoneyIn3DInitInput struct for MoneyIn3DInitInput
type MoneyIn3DInitInput struct {
	Card *CardInfo `json:"card,omitempty"`
	// Return URL of the 3-D Secure website. Your client will be redirected on this URL, which should be your website URL.
	ReturnUrl *string `json:"returnUrl,omitempty"`
	CardId *int64 `json:"cardId,omitempty"`
	RegisterCard *bool `json:"registerCard,omitempty"`
	ThreeDs *ThreeDs `json:"threeDS,omitempty"`
	RiskAnalysis *RiskAnalysis `json:"riskAnalysis,omitempty"`
	// Average amount of future recurring payments to cover recurring payments with variable amounts. **Note:** Only if the amount is higher than the amount of the first transaction. cover recurring payments with variable amounts. **Note:** Only if the amount is higher than the amount of the first transaction.
	RecurringAvgAmount *int32 `json:"recurringAvgAmount,omitempty"`
	// **Recurring**: If a payment follows a recurring amount or **One_shot**: when there is no future recurring payments.
	PaymentPattern *string `json:"paymentPattern,omitempty"`
	// Unique ID of the call, generated by your server. This ID can be used as a search field when looking for operation details
	Reference *string `json:"reference,omitempty"`
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

// NewMoneyIn3DInitInput instantiates a new MoneyIn3DInitInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoneyIn3DInitInput(accountId string) *MoneyIn3DInitInput {
	this := MoneyIn3DInitInput{}
	this.AccountId = accountId
	return &this
}

// NewMoneyIn3DInitInputWithDefaults instantiates a new MoneyIn3DInitInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoneyIn3DInitInputWithDefaults() *MoneyIn3DInitInput {
	this := MoneyIn3DInitInput{}
	return &this
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetCard() CardInfo {
	if o == nil || o.Card == nil {
		var ret CardInfo
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetCardOk() (*CardInfo, bool) {
	if o == nil || o.Card == nil {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasCard() bool {
	if o != nil && o.Card != nil {
		return true
	}

	return false
}

// SetCard gets a reference to the given CardInfo and assigns it to the Card field.
func (o *MoneyIn3DInitInput) SetCard(v CardInfo) {
	o.Card = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetReturnUrl() string {
	if o == nil || o.ReturnUrl == nil {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetReturnUrlOk() (*string, bool) {
	if o == nil || o.ReturnUrl == nil {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasReturnUrl() bool {
	if o != nil && o.ReturnUrl != nil {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *MoneyIn3DInitInput) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetCardId() int64 {
	if o == nil || o.CardId == nil {
		var ret int64
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetCardIdOk() (*int64, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given int64 and assigns it to the CardId field.
func (o *MoneyIn3DInitInput) SetCardId(v int64) {
	o.CardId = &v
}

// GetRegisterCard returns the RegisterCard field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetRegisterCard() bool {
	if o == nil || o.RegisterCard == nil {
		var ret bool
		return ret
	}
	return *o.RegisterCard
}

// GetRegisterCardOk returns a tuple with the RegisterCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetRegisterCardOk() (*bool, bool) {
	if o == nil || o.RegisterCard == nil {
		return nil, false
	}
	return o.RegisterCard, true
}

// HasRegisterCard returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasRegisterCard() bool {
	if o != nil && o.RegisterCard != nil {
		return true
	}

	return false
}

// SetRegisterCard gets a reference to the given bool and assigns it to the RegisterCard field.
func (o *MoneyIn3DInitInput) SetRegisterCard(v bool) {
	o.RegisterCard = &v
}

// GetThreeDS returns the ThreeDs field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetThreeDS() ThreeDs {
	if o == nil || o.ThreeDs == nil {
		var ret ThreeDs
		return ret
	}
	return *o.ThreeDs
}

// GetThreeDSOk returns a tuple with the ThreeDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetThreeDSOk() (*ThreeDs, bool) {
	if o == nil || o.ThreeDs == nil {
		return nil, false
	}
	return o.ThreeDs, true
}

// HasThreeDS returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasThreeDS() bool {
	if o != nil && o.ThreeDs != nil {
		return true
	}

	return false
}

// SetThreeDS gets a reference to the given ThreeDs and assigns it to the ThreeDs field.
func (o *MoneyIn3DInitInput) SetThreeDS(v ThreeDs) {
	o.ThreeDs = &v
}

// GetRiskAnalysis returns the RiskAnalysis field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetRiskAnalysis() RiskAnalysis {
	if o == nil || o.RiskAnalysis == nil {
		var ret RiskAnalysis
		return ret
	}
	return *o.RiskAnalysis
}

// GetRiskAnalysisOk returns a tuple with the RiskAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetRiskAnalysisOk() (*RiskAnalysis, bool) {
	if o == nil || o.RiskAnalysis == nil {
		return nil, false
	}
	return o.RiskAnalysis, true
}

// HasRiskAnalysis returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasRiskAnalysis() bool {
	if o != nil && o.RiskAnalysis != nil {
		return true
	}

	return false
}

// SetRiskAnalysis gets a reference to the given RiskAnalysis and assigns it to the RiskAnalysis field.
func (o *MoneyIn3DInitInput) SetRiskAnalysis(v RiskAnalysis) {
	o.RiskAnalysis = &v
}

// GetRecurringAvgAmount returns the RecurringAvgAmount field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetRecurringAvgAmount() int32 {
	if o == nil || o.RecurringAvgAmount == nil {
		var ret int32
		return ret
	}
	return *o.RecurringAvgAmount
}

// GetRecurringAvgAmountOk returns a tuple with the RecurringAvgAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetRecurringAvgAmountOk() (*int32, bool) {
	if o == nil || o.RecurringAvgAmount == nil {
		return nil, false
	}
	return o.RecurringAvgAmount, true
}

// HasRecurringAvgAmount returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasRecurringAvgAmount() bool {
	if o != nil && o.RecurringAvgAmount != nil {
		return true
	}

	return false
}

// SetRecurringAvgAmount gets a reference to the given int32 and assigns it to the RecurringAvgAmount field.
func (o *MoneyIn3DInitInput) SetRecurringAvgAmount(v int32) {
	o.RecurringAvgAmount = &v
}

// GetPaymentPattern returns the PaymentPattern field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetPaymentPattern() string {
	if o == nil || o.PaymentPattern == nil {
		var ret string
		return ret
	}
	return *o.PaymentPattern
}

// GetPaymentPatternOk returns a tuple with the PaymentPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetPaymentPatternOk() (*string, bool) {
	if o == nil || o.PaymentPattern == nil {
		return nil, false
	}
	return o.PaymentPattern, true
}

// HasPaymentPattern returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasPaymentPattern() bool {
	if o != nil && o.PaymentPattern != nil {
		return true
	}

	return false
}

// SetPaymentPattern gets a reference to the given string and assigns it to the PaymentPattern field.
func (o *MoneyIn3DInitInput) SetPaymentPattern(v string) {
	o.PaymentPattern = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MoneyIn3DInitInput) SetReference(v string) {
	o.Reference = &v
}

// GetAccountId returns the AccountId field value
func (o *MoneyIn3DInitInput) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MoneyIn3DInitInput) SetAccountId(v string) {
	o.AccountId = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetTotalAmount() int32 {
	if o == nil || o.TotalAmount == nil {
		var ret int32
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetTotalAmountOk() (*int32, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasTotalAmount() bool {
	if o != nil && o.TotalAmount != nil {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int32 and assigns it to the TotalAmount field.
func (o *MoneyIn3DInitInput) SetTotalAmount(v int32) {
	o.TotalAmount = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetCommissionAmount() int32 {
	if o == nil || o.CommissionAmount == nil {
		var ret int32
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetCommissionAmountOk() (*int32, bool) {
	if o == nil || o.CommissionAmount == nil {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasCommissionAmount() bool {
	if o != nil && o.CommissionAmount != nil {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given int32 and assigns it to the CommissionAmount field.
func (o *MoneyIn3DInitInput) SetCommissionAmount(v int32) {
	o.CommissionAmount = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MoneyIn3DInitInput) SetComment(v string) {
	o.Comment = &v
}

// GetAutoCommission returns the AutoCommission field value if set, zero value otherwise.
func (o *MoneyIn3DInitInput) GetAutoCommission() bool {
	if o == nil || o.AutoCommission == nil {
		var ret bool
		return ret
	}
	return *o.AutoCommission
}

// GetAutoCommissionOk returns a tuple with the AutoCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoneyIn3DInitInput) GetAutoCommissionOk() (*bool, bool) {
	if o == nil || o.AutoCommission == nil {
		return nil, false
	}
	return o.AutoCommission, true
}

// HasAutoCommission returns a boolean if a field has been set.
func (o *MoneyIn3DInitInput) HasAutoCommission() bool {
	if o != nil && o.AutoCommission != nil {
		return true
	}

	return false
}

// SetAutoCommission gets a reference to the given bool and assigns it to the AutoCommission field.
func (o *MoneyIn3DInitInput) SetAutoCommission(v bool) {
	o.AutoCommission = &v
}

func (o MoneyIn3DInitInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Card != nil {
		toSerialize["card"] = o.Card
	}
	if o.ReturnUrl != nil {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if o.CardId != nil {
		toSerialize["cardId"] = o.CardId
	}
	if o.RegisterCard != nil {
		toSerialize["registerCard"] = o.RegisterCard
	}
	if o.ThreeDs != nil {
		toSerialize["threeDS"] = o.ThreeDs
	}
	if o.RiskAnalysis != nil {
		toSerialize["riskAnalysis"] = o.RiskAnalysis
	}
	if o.RecurringAvgAmount != nil {
		toSerialize["recurringAvgAmount"] = o.RecurringAvgAmount
	}
	if o.PaymentPattern != nil {
		toSerialize["paymentPattern"] = o.PaymentPattern
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
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

type NullableMoneyIn3DInitInput struct {
	value *MoneyIn3DInitInput
	isSet bool
}

func (v NullableMoneyIn3DInitInput) Get() *MoneyIn3DInitInput {
	return v.value
}

func (v *NullableMoneyIn3DInitInput) Set(val *MoneyIn3DInitInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMoneyIn3DInitInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMoneyIn3DInitInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoneyIn3DInitInput(val *MoneyIn3DInitInput) *NullableMoneyIn3DInitInput {
	return &NullableMoneyIn3DInitInput{value: val, isSet: true}
}

func (v NullableMoneyIn3DInitInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoneyIn3DInitInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


