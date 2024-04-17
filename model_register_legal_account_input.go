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

// RegisterLegalAccountInput struct for RegisterLegalAccountInput
type RegisterLegalAccountInput struct {
	Company Company `json:"company"`
	// Indicates if the legal representative is also an Ultimate Beneficial owner (i.e. shareholder with &gt;25% of capital or voting rights).
	IsUltimateBeneficialOwner *bool `json:"isUltimateBeneficialOwner,omitempty"`
	// Payment account ID that you use to identify the customer.Choose your unique number.<br /><b>Note:</b> If you plan to credit payments accounts  by fund transfer, please use short alphanumeric payment account identifiers (max 20 char.).   Your customers will have to write their payment account identifier in the transfer label/comment, a label of more that 20 characters could be cut when passing the the banking system.
	AccountId string `json:"accountId"`
	// Unique Email
	Email string `json:"email"`
	// Client title
	Title *string `json:"title,omitempty"`
	// Client first name
	FirstName string `json:"firstName"`
	// Client last name
	LastName string `json:"lastName"`
	Adresse *Address `json:"adresse,omitempty"`
	Birth *Birth `json:"birth,omitempty"`
	// Client nationality, using ISO 3166-1 alpha-3 format  Please separate multiple nationalities with a comma
	Nationality string `json:"nationality"`
	// Phone number with MSISDN format: international number with country code without \"00\" neither \"+\".
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Mobile phone number with MSISDN format: international number with country code without \"00\" neither \"+\".   This will be used by default when eletronically signing documents.
	MobileNumber *string `json:"mobileNumber,omitempty"`
	// For crowdfunding/loan platforms, indicates if the wallet is created for a debtor.
	IsDebtor *bool `json:"isDebtor,omitempty"`
	// Indicates if the payment account is created for a Payer or a Beneficiary.  1: Payer (default)  2: Beneficiary
	PayerOrBeneficiary int32 `json:"payerOrBeneficiary"`
	// Indicates if the payment account is for a one-time customer.   If yes, the payment account will be created with status 14, allowing only one payment.   The maximum amount will be defined with Lemonway.
	IsOneTimeCustomerAccount *bool `json:"isOneTimeCustomerAccount,omitempty"`
	// **Note:** This option is available depending on your contract  True, in case this option is enabled in your contract.  Otherwise it will be considered a Client Wallet.
	IsTechnicalAccount *bool `json:"isTechnicalAccount,omitempty"`
}

// NewRegisterLegalAccountInput instantiates a new RegisterLegalAccountInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterLegalAccountInput(company Company, accountId string, email string, firstName string, lastName string, nationality string, payerOrBeneficiary int32) *RegisterLegalAccountInput {
	this := RegisterLegalAccountInput{}
	this.Company = company
	this.AccountId = accountId
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.Nationality = nationality
	this.PayerOrBeneficiary = payerOrBeneficiary
	return &this
}

// NewRegisterLegalAccountInputWithDefaults instantiates a new RegisterLegalAccountInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterLegalAccountInputWithDefaults() *RegisterLegalAccountInput {
	this := RegisterLegalAccountInput{}
	return &this
}

// GetCompany returns the Company field value
func (o *RegisterLegalAccountInput) GetCompany() Company {
	if o == nil {
		var ret Company
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetCompanyOk() (*Company, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *RegisterLegalAccountInput) SetCompany(v Company) {
	o.Company = v
}

// GetIsUltimateBeneficialOwner returns the IsUltimateBeneficialOwner field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetIsUltimateBeneficialOwner() bool {
	if o == nil || o.IsUltimateBeneficialOwner == nil {
		var ret bool
		return ret
	}
	return *o.IsUltimateBeneficialOwner
}

// GetIsUltimateBeneficialOwnerOk returns a tuple with the IsUltimateBeneficialOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetIsUltimateBeneficialOwnerOk() (*bool, bool) {
	if o == nil || o.IsUltimateBeneficialOwner == nil {
		return nil, false
	}
	return o.IsUltimateBeneficialOwner, true
}

// HasIsUltimateBeneficialOwner returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasIsUltimateBeneficialOwner() bool {
	if o != nil && o.IsUltimateBeneficialOwner != nil {
		return true
	}

	return false
}

// SetIsUltimateBeneficialOwner gets a reference to the given bool and assigns it to the IsUltimateBeneficialOwner field.
func (o *RegisterLegalAccountInput) SetIsUltimateBeneficialOwner(v bool) {
	o.IsUltimateBeneficialOwner = &v
}

// GetAccountId returns the AccountId field value
func (o *RegisterLegalAccountInput) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *RegisterLegalAccountInput) SetAccountId(v string) {
	o.AccountId = v
}

// GetEmail returns the Email field value
func (o *RegisterLegalAccountInput) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RegisterLegalAccountInput) SetEmail(v string) {
	o.Email = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RegisterLegalAccountInput) SetTitle(v string) {
	o.Title = &v
}

// GetFirstName returns the FirstName field value
func (o *RegisterLegalAccountInput) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *RegisterLegalAccountInput) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *RegisterLegalAccountInput) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *RegisterLegalAccountInput) SetLastName(v string) {
	o.LastName = v
}

// GetAdresse returns the Adresse field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetAdresse() Address {
	if o == nil || o.Adresse == nil {
		var ret Address
		return ret
	}
	return *o.Adresse
}

// GetAdresseOk returns a tuple with the Adresse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetAdresseOk() (*Address, bool) {
	if o == nil || o.Adresse == nil {
		return nil, false
	}
	return o.Adresse, true
}

// HasAdresse returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasAdresse() bool {
	if o != nil && o.Adresse != nil {
		return true
	}

	return false
}

// SetAdresse gets a reference to the given Address and assigns it to the Adresse field.
func (o *RegisterLegalAccountInput) SetAdresse(v Address) {
	o.Adresse = &v
}

// GetBirth returns the Birth field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetBirth() Birth {
	if o == nil || o.Birth == nil {
		var ret Birth
		return ret
	}
	return *o.Birth
}

// GetBirthOk returns a tuple with the Birth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetBirthOk() (*Birth, bool) {
	if o == nil || o.Birth == nil {
		return nil, false
	}
	return o.Birth, true
}

// HasBirth returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasBirth() bool {
	if o != nil && o.Birth != nil {
		return true
	}

	return false
}

// SetBirth gets a reference to the given Birth and assigns it to the Birth field.
func (o *RegisterLegalAccountInput) SetBirth(v Birth) {
	o.Birth = &v
}

// GetNationality returns the Nationality field value
func (o *RegisterLegalAccountInput) GetNationality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetNationalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nationality, true
}

// SetNationality sets field value
func (o *RegisterLegalAccountInput) SetNationality(v string) {
	o.Nationality = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *RegisterLegalAccountInput) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetMobileNumber() string {
	if o == nil || o.MobileNumber == nil {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetMobileNumberOk() (*string, bool) {
	if o == nil || o.MobileNumber == nil {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasMobileNumber() bool {
	if o != nil && o.MobileNumber != nil {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *RegisterLegalAccountInput) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetIsDebtor returns the IsDebtor field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetIsDebtor() bool {
	if o == nil || o.IsDebtor == nil {
		var ret bool
		return ret
	}
	return *o.IsDebtor
}

// GetIsDebtorOk returns a tuple with the IsDebtor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetIsDebtorOk() (*bool, bool) {
	if o == nil || o.IsDebtor == nil {
		return nil, false
	}
	return o.IsDebtor, true
}

// HasIsDebtor returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasIsDebtor() bool {
	if o != nil && o.IsDebtor != nil {
		return true
	}

	return false
}

// SetIsDebtor gets a reference to the given bool and assigns it to the IsDebtor field.
func (o *RegisterLegalAccountInput) SetIsDebtor(v bool) {
	o.IsDebtor = &v
}

// GetPayerOrBeneficiary returns the PayerOrBeneficiary field value
func (o *RegisterLegalAccountInput) GetPayerOrBeneficiary() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PayerOrBeneficiary
}

// GetPayerOrBeneficiaryOk returns a tuple with the PayerOrBeneficiary field value
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetPayerOrBeneficiaryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerOrBeneficiary, true
}

// SetPayerOrBeneficiary sets field value
func (o *RegisterLegalAccountInput) SetPayerOrBeneficiary(v int32) {
	o.PayerOrBeneficiary = v
}

// GetIsOneTimeCustomerAccount returns the IsOneTimeCustomerAccount field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetIsOneTimeCustomerAccount() bool {
	if o == nil || o.IsOneTimeCustomerAccount == nil {
		var ret bool
		return ret
	}
	return *o.IsOneTimeCustomerAccount
}

// GetIsOneTimeCustomerAccountOk returns a tuple with the IsOneTimeCustomerAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetIsOneTimeCustomerAccountOk() (*bool, bool) {
	if o == nil || o.IsOneTimeCustomerAccount == nil {
		return nil, false
	}
	return o.IsOneTimeCustomerAccount, true
}

// HasIsOneTimeCustomerAccount returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasIsOneTimeCustomerAccount() bool {
	if o != nil && o.IsOneTimeCustomerAccount != nil {
		return true
	}

	return false
}

// SetIsOneTimeCustomerAccount gets a reference to the given bool and assigns it to the IsOneTimeCustomerAccount field.
func (o *RegisterLegalAccountInput) SetIsOneTimeCustomerAccount(v bool) {
	o.IsOneTimeCustomerAccount = &v
}

// GetIsTechnicalAccount returns the IsTechnicalAccount field value if set, zero value otherwise.
func (o *RegisterLegalAccountInput) GetIsTechnicalAccount() bool {
	if o == nil || o.IsTechnicalAccount == nil {
		var ret bool
		return ret
	}
	return *o.IsTechnicalAccount
}

// GetIsTechnicalAccountOk returns a tuple with the IsTechnicalAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterLegalAccountInput) GetIsTechnicalAccountOk() (*bool, bool) {
	if o == nil || o.IsTechnicalAccount == nil {
		return nil, false
	}
	return o.IsTechnicalAccount, true
}

// HasIsTechnicalAccount returns a boolean if a field has been set.
func (o *RegisterLegalAccountInput) HasIsTechnicalAccount() bool {
	if o != nil && o.IsTechnicalAccount != nil {
		return true
	}

	return false
}

// SetIsTechnicalAccount gets a reference to the given bool and assigns it to the IsTechnicalAccount field.
func (o *RegisterLegalAccountInput) SetIsTechnicalAccount(v bool) {
	o.IsTechnicalAccount = &v
}

func (o RegisterLegalAccountInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["company"] = o.Company
	}
	if o.IsUltimateBeneficialOwner != nil {
		toSerialize["isUltimateBeneficialOwner"] = o.IsUltimateBeneficialOwner
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if o.Adresse != nil {
		toSerialize["adresse"] = o.Adresse
	}
	if o.Birth != nil {
		toSerialize["birth"] = o.Birth
	}
	if true {
		toSerialize["nationality"] = o.Nationality
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.MobileNumber != nil {
		toSerialize["mobileNumber"] = o.MobileNumber
	}
	if o.IsDebtor != nil {
		toSerialize["isDebtor"] = o.IsDebtor
	}
	if true {
		toSerialize["payerOrBeneficiary"] = o.PayerOrBeneficiary
	}
	if o.IsOneTimeCustomerAccount != nil {
		toSerialize["isOneTimeCustomerAccount"] = o.IsOneTimeCustomerAccount
	}
	if o.IsTechnicalAccount != nil {
		toSerialize["isTechnicalAccount"] = o.IsTechnicalAccount
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterLegalAccountInput struct {
	value *RegisterLegalAccountInput
	isSet bool
}

func (v NullableRegisterLegalAccountInput) Get() *RegisterLegalAccountInput {
	return v.value
}

func (v *NullableRegisterLegalAccountInput) Set(val *RegisterLegalAccountInput) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterLegalAccountInput) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterLegalAccountInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterLegalAccountInput(val *RegisterLegalAccountInput) *NullableRegisterLegalAccountInput {
	return &NullableRegisterLegalAccountInput{value: val, isSet: true}
}

func (v NullableRegisterLegalAccountInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterLegalAccountInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

