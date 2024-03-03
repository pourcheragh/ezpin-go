# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **float32** | id of activation transaction | [optional] 
**Barcode** | Pointer to **string** | barcode of physical card | [optional] 
**Product** | Pointer to **string** | name of product | [optional] 
**Status** | Pointer to **float32** | Pending(-1) , Activated(1) , Inactive(0) | [optional] 
**StatusText** | Pointer to **string** | status of activation order | [optional] 
**TotalFaceValue** | Pointer to **float32** | physical card price | [optional] 
**TotalFees** | Pointer to **float32** | activation fee of physical card | [optional] 
**TotalDiscounts** | Pointer to **float32** | discount of physical card | [optional] 
**TotalCustomerCost** | Pointer to **float32** | price that customer should pay | [optional] 
**TerminalId** | Pointer to **float32** | Terminal ID of the sub-user | [optional] 
**Currency** | Pointer to **string** | name of currency | [optional] 
**CreatedTime** | Pointer to **string** | time of activation | [optional] 
**ReferenceCode** | Pointer to **string** | reference code of activation | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *Transaction) GetTransactionId() float32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*float32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v float32)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Transaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetBarcode

`func (o *Transaction) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Transaction) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Transaction) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Transaction) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetProduct

`func (o *Transaction) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Transaction) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Transaction) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Transaction) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusText

`func (o *Transaction) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *Transaction) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *Transaction) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *Transaction) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetTotalFaceValue

`func (o *Transaction) GetTotalFaceValue() float32`

GetTotalFaceValue returns the TotalFaceValue field if non-nil, zero value otherwise.

### GetTotalFaceValueOk

`func (o *Transaction) GetTotalFaceValueOk() (*float32, bool)`

GetTotalFaceValueOk returns a tuple with the TotalFaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFaceValue

`func (o *Transaction) SetTotalFaceValue(v float32)`

SetTotalFaceValue sets TotalFaceValue field to given value.

### HasTotalFaceValue

`func (o *Transaction) HasTotalFaceValue() bool`

HasTotalFaceValue returns a boolean if a field has been set.

### GetTotalFees

`func (o *Transaction) GetTotalFees() float32`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *Transaction) GetTotalFeesOk() (*float32, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *Transaction) SetTotalFees(v float32)`

SetTotalFees sets TotalFees field to given value.

### HasTotalFees

`func (o *Transaction) HasTotalFees() bool`

HasTotalFees returns a boolean if a field has been set.

### GetTotalDiscounts

`func (o *Transaction) GetTotalDiscounts() float32`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *Transaction) GetTotalDiscountsOk() (*float32, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *Transaction) SetTotalDiscounts(v float32)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *Transaction) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalCustomerCost

`func (o *Transaction) GetTotalCustomerCost() float32`

GetTotalCustomerCost returns the TotalCustomerCost field if non-nil, zero value otherwise.

### GetTotalCustomerCostOk

`func (o *Transaction) GetTotalCustomerCostOk() (*float32, bool)`

GetTotalCustomerCostOk returns a tuple with the TotalCustomerCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomerCost

`func (o *Transaction) SetTotalCustomerCost(v float32)`

SetTotalCustomerCost sets TotalCustomerCost field to given value.

### HasTotalCustomerCost

`func (o *Transaction) HasTotalCustomerCost() bool`

HasTotalCustomerCost returns a boolean if a field has been set.

### GetTerminalId

`func (o *Transaction) GetTerminalId() float32`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *Transaction) GetTerminalIdOk() (*float32, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *Transaction) SetTerminalId(v float32)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *Transaction) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetCurrency

`func (o *Transaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Transaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Transaction) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Transaction) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Transaction) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Transaction) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetReferenceCode

`func (o *Transaction) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *Transaction) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *Transaction) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.

### HasReferenceCode

`func (o *Transaction) HasReferenceCode() bool`

HasReferenceCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


