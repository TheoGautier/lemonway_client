# MoneyOutInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Payment Account ID to be debited | 
**IbanId** | Pointer to **int64** | IBAN ID  If no IBAN is specified, the last verified(validated) IBAN will be used. | [optional] 
**TotalAmount** | Pointer to **int32** | Total amount to debit from the Wallet  The client will receive on his bank account[totalAmount] minus[commissionAmount].  Amounts are given as integer numbers in cents | [optional] 
**CommissionAmount** | Pointer to **int32** | Your fee  Amounts are given as integer numbers in cents | [optional] 
**Comment** | Pointer to **string** | Payment Comment | [optional] 
**AutoCommission** | **bool** | This should be set to No (0) for most sites  If true:  1. [amountCom] will be ignored and will be replaced with Lemonway&#39;s fee.  2. You will not receive any fee. | 
**Reference** | Pointer to **string** | Unique ID of the call, generated by your server. This ID can be used as a search field when looking for operation details | [optional] 

## Methods

### NewMoneyOutInput

`func NewMoneyOutInput(accountId string, autoCommission bool, ) *MoneyOutInput`

NewMoneyOutInput instantiates a new MoneyOutInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyOutInputWithDefaults

`func NewMoneyOutInputWithDefaults() *MoneyOutInput`

NewMoneyOutInputWithDefaults instantiates a new MoneyOutInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MoneyOutInput) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MoneyOutInput) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MoneyOutInput) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetIbanId

`func (o *MoneyOutInput) GetIbanId() int64`

GetIbanId returns the IbanId field if non-nil, zero value otherwise.

### GetIbanIdOk

`func (o *MoneyOutInput) GetIbanIdOk() (*int64, bool)`

GetIbanIdOk returns a tuple with the IbanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanId

`func (o *MoneyOutInput) SetIbanId(v int64)`

SetIbanId sets IbanId field to given value.

### HasIbanId

`func (o *MoneyOutInput) HasIbanId() bool`

HasIbanId returns a boolean if a field has been set.

### GetTotalAmount

`func (o *MoneyOutInput) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *MoneyOutInput) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *MoneyOutInput) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *MoneyOutInput) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetCommissionAmount

`func (o *MoneyOutInput) GetCommissionAmount() int32`

GetCommissionAmount returns the CommissionAmount field if non-nil, zero value otherwise.

### GetCommissionAmountOk

`func (o *MoneyOutInput) GetCommissionAmountOk() (*int32, bool)`

GetCommissionAmountOk returns a tuple with the CommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAmount

`func (o *MoneyOutInput) SetCommissionAmount(v int32)`

SetCommissionAmount sets CommissionAmount field to given value.

### HasCommissionAmount

`func (o *MoneyOutInput) HasCommissionAmount() bool`

HasCommissionAmount returns a boolean if a field has been set.

### GetComment

`func (o *MoneyOutInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MoneyOutInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MoneyOutInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MoneyOutInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAutoCommission

`func (o *MoneyOutInput) GetAutoCommission() bool`

GetAutoCommission returns the AutoCommission field if non-nil, zero value otherwise.

### GetAutoCommissionOk

`func (o *MoneyOutInput) GetAutoCommissionOk() (*bool, bool)`

GetAutoCommissionOk returns a tuple with the AutoCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCommission

`func (o *MoneyOutInput) SetAutoCommission(v bool)`

SetAutoCommission sets AutoCommission field to given value.


### GetReference

`func (o *MoneyOutInput) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *MoneyOutInput) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *MoneyOutInput) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *MoneyOutInput) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

