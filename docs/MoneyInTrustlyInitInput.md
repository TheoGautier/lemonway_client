# MoneyInTrustlyInitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinalCustomerEmail** | **string** | End user email | 
**FinalCustomerFirstName** | **string** | End user first name | 
**FinalCustomerLastName** | **string** | End user last name | 
**ReturnUrl** | **string** | Your site returns the URL, called to terminate the operation and on which the callback will be sent, with data in POST parameters.  This URL must contain a unique ID so you know which operation is related to the return. | 
**Reference** | Pointer to **string** | Unique ID of the call, generated by your server. This ID can be used as a search field when looking for operation details | [optional] 
**AccountId** | **string** | Payment Account ID to Credit | 
**TotalAmount** | Pointer to **int32** | Amount to Debit  Amounts are given as integer numbers in cents | [optional] 
**CommissionAmount** | Pointer to **int32** | Your Fee  Amounts are given as integer numbers in cents | [optional] 
**Comment** | Pointer to **string** | Comment Regarding the Transaction | [optional] 
**AutoCommission** | Pointer to **bool** | If true:  1. [amountCom] will be ignored and will be replaced with Lemonway&#39;s fee  2. You will not receive any fee | [optional] 

## Methods

### NewMoneyInTrustlyInitInput

`func NewMoneyInTrustlyInitInput(finalCustomerEmail string, finalCustomerFirstName string, finalCustomerLastName string, returnUrl string, accountId string, ) *MoneyInTrustlyInitInput`

NewMoneyInTrustlyInitInput instantiates a new MoneyInTrustlyInitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInTrustlyInitInputWithDefaults

`func NewMoneyInTrustlyInitInputWithDefaults() *MoneyInTrustlyInitInput`

NewMoneyInTrustlyInitInputWithDefaults instantiates a new MoneyInTrustlyInitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalCustomerEmail

`func (o *MoneyInTrustlyInitInput) GetFinalCustomerEmail() string`

GetFinalCustomerEmail returns the FinalCustomerEmail field if non-nil, zero value otherwise.

### GetFinalCustomerEmailOk

`func (o *MoneyInTrustlyInitInput) GetFinalCustomerEmailOk() (*string, bool)`

GetFinalCustomerEmailOk returns a tuple with the FinalCustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCustomerEmail

`func (o *MoneyInTrustlyInitInput) SetFinalCustomerEmail(v string)`

SetFinalCustomerEmail sets FinalCustomerEmail field to given value.


### GetFinalCustomerFirstName

`func (o *MoneyInTrustlyInitInput) GetFinalCustomerFirstName() string`

GetFinalCustomerFirstName returns the FinalCustomerFirstName field if non-nil, zero value otherwise.

### GetFinalCustomerFirstNameOk

`func (o *MoneyInTrustlyInitInput) GetFinalCustomerFirstNameOk() (*string, bool)`

GetFinalCustomerFirstNameOk returns a tuple with the FinalCustomerFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCustomerFirstName

`func (o *MoneyInTrustlyInitInput) SetFinalCustomerFirstName(v string)`

SetFinalCustomerFirstName sets FinalCustomerFirstName field to given value.


### GetFinalCustomerLastName

`func (o *MoneyInTrustlyInitInput) GetFinalCustomerLastName() string`

GetFinalCustomerLastName returns the FinalCustomerLastName field if non-nil, zero value otherwise.

### GetFinalCustomerLastNameOk

`func (o *MoneyInTrustlyInitInput) GetFinalCustomerLastNameOk() (*string, bool)`

GetFinalCustomerLastNameOk returns a tuple with the FinalCustomerLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCustomerLastName

`func (o *MoneyInTrustlyInitInput) SetFinalCustomerLastName(v string)`

SetFinalCustomerLastName sets FinalCustomerLastName field to given value.


### GetReturnUrl

`func (o *MoneyInTrustlyInitInput) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *MoneyInTrustlyInitInput) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *MoneyInTrustlyInitInput) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetReference

`func (o *MoneyInTrustlyInitInput) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *MoneyInTrustlyInitInput) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *MoneyInTrustlyInitInput) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *MoneyInTrustlyInitInput) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetAccountId

`func (o *MoneyInTrustlyInitInput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MoneyInTrustlyInitInput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MoneyInTrustlyInitInput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTotalAmount

`func (o *MoneyInTrustlyInitInput) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *MoneyInTrustlyInitInput) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *MoneyInTrustlyInitInput) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *MoneyInTrustlyInitInput) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetCommissionAmount

`func (o *MoneyInTrustlyInitInput) GetCommissionAmount() int32`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *MoneyInTrustlyInitInput) GetCommissionAmountOk() (*int32, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *MoneyInTrustlyInitInput) SetCommissionAmount(v int32)`

SetCommissionAmount sets CommissionAmount field to given value.

### HasCommissionAmount

`func (o *MoneyInTrustlyInitInput) HasCommissionAmount() bool`

HasCommissionAmount returns a boolean if a field has been set.

### GetComment

`func (o *MoneyInTrustlyInitInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MoneyInTrustlyInitInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MoneyInTrustlyInitInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MoneyInTrustlyInitInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAutoCommission

`func (o *MoneyInTrustlyInitInput) GetAutoCommission() bool`

GetAutoCommission returns the AutoCommission field if non-nil, zero value otherwise.

### GetAutoCommissionOk

`func (o *MoneyInTrustlyInitInput) GetAutoCommissionOk() (*bool, bool)`

GetAutoCommissionOk returns a tuple with the AutoCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCommission

`func (o *MoneyInTrustlyInitInput) SetAutoCommission(v bool)`

SetAutoCommission sets AutoCommission field to given value.

### HasAutoCommission

`func (o *MoneyInTrustlyInitInput) HasAutoCommission() bool`

HasAutoCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


