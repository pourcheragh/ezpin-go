# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int32** | id of order | [optional] 
**DeliveryType** | Pointer to **int32** | * &#x60;0&#x60; None *&#x60;1&#x60; email *&#x60;2&#x60; SMS *&#x60;3&#x60; WhatsApp | [optional] 
**Destination** | Pointer to **string** | E-mail or Phone/Whatsapp number you want the order to be delivered to. | [optional] 
**Status** | Pointer to **int32** | *&#x60;-1&#x60; Reject *&#x60;0&#x60; Pending *&#x60;1&#x60; Accept | [optional] 
**StatusText** | Pointer to **string** | status of order | [optional] 
**CreatedTime** | Pointer to **string** | time of order creation | [optional] 
**TerminalId** | Pointer to **int32** | terminal id of sub user | [optional] 
**ReferenceCode** | Pointer to **string** | reference code of order | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**Count** | Pointer to **int32** | order quantity | [optional] 
**TotalFaceValue** | Pointer to **decimal.Decimal** | total of product original prices | [optional] 
**TotalFees** | Pointer to **decimal.Decimal** | sum of activation fees | [optional] 
**TotalDiscounts** | Pointer to **string** | sum of discounts | [optional] 
**TotalCustomerCost** | Pointer to **decimal.Decimal** | total price customer should pay | [optional] 
**IsCompleted** | Pointer to **bool** | is order process completed or not | [optional] 
**ShareLink** | Pointer to **string** | link of order cards. | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *Order) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Order) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetDeliveryType

`func (o *Order) GetDeliveryType() int32`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *Order) GetDeliveryTypeOk() (*int32, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *Order) SetDeliveryType(v int32)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *Order) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDestination

`func (o *Order) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Order) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Order) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Order) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusText

`func (o *Order) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *Order) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *Order) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *Order) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Order) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Order) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Order) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Order) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetTerminalId

`func (o *Order) GetTerminalId() int32`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *Order) GetTerminalIdOk() (*int32, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *Order) SetTerminalId(v int32)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *Order) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetReferenceCode

`func (o *Order) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *Order) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *Order) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.

### HasReferenceCode

`func (o *Order) HasReferenceCode() bool`

HasReferenceCode returns a boolean if a field has been set.

### GetProduct

`func (o *Order) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Order) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Order) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Order) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCount

`func (o *Order) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Order) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Order) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Order) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalFaceValue

`func (o *Order) GetTotalFaceValue() decimal.Decimal`

GetTotalFaceValue returns the TotalFaceValue field if non-nil, zero value otherwise.

### GetTotalFaceValueOk

`func (o *Order) GetTotalFaceValueOk() (*decimal.Decimal, bool)`

GetTotalFaceValueOk returns a tuple with the TotalFaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFaceValue

`func (o *Order) SetTotalFaceValue(v decimal.Decimal)`

SetTotalFaceValue sets TotalFaceValue field to given value.

### HasTotalFaceValue

`func (o *Order) HasTotalFaceValue() bool`

HasTotalFaceValue returns a boolean if a field has been set.

### GetTotalFees

`func (o *Order) GetTotalFees() decimal.Decimal`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *Order) GetTotalFeesOk() (*decimal.Decimal, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *Order) SetTotalFees(v decimal.Decimal)`

SetTotalFees sets TotalFees field to given value.

### HasTotalFees

`func (o *Order) HasTotalFees() bool`

HasTotalFees returns a boolean if a field has been set.

### GetTotalDiscounts

`func (o *Order) GetTotalDiscounts() string`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *Order) GetTotalDiscountsOk() (*string, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *Order) SetTotalDiscounts(v string)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *Order) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalCustomerCost

`func (o *Order) GetTotalCustomerCost() decimal.Decimal`

GetTotalCustomerCost returns the TotalCustomerCost field if non-nil, zero value otherwise.

### GetTotalCustomerCostOk

`func (o *Order) GetTotalCustomerCostOk() (*decimal.Decimal, bool)`

GetTotalCustomerCostOk returns a tuple with the TotalCustomerCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomerCost

`func (o *Order) SetTotalCustomerCost(v decimal.Decimal)`

SetTotalCustomerCost sets TotalCustomerCost field to given value.

### HasTotalCustomerCost

`func (o *Order) HasTotalCustomerCost() bool`

HasTotalCustomerCost returns a boolean if a field has been set.

### GetIsCompleted

`func (o *Order) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *Order) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *Order) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *Order) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### GetShareLink

`func (o *Order) GetShareLink() string`

GetShareLink returns the ShareLink field if non-nil, zero value otherwise.

### GetShareLinkOk

`func (o *Order) GetShareLinkOk() (*string, bool)`

GetShareLinkOk returns a tuple with the ShareLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLink

`func (o *Order) SetShareLink(v string)`

SetShareLink sets ShareLink field to given value.

### HasShareLink

`func (o *Order) HasShareLink() bool`

HasShareLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


