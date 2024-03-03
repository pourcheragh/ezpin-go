# OrdersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **int32** | Catalog SKU | [optional] 
**Quantity** | Pointer to **int32** | Item count | [optional] 
**PreOder** | Pointer to **bool** | To determine whether you want to pre order your request or not. The product must support pre order mood. if you pass true your cards will get ready in the future. | [optional] 
**Price** | Pointer to **decimal.Decimal** | Item price | [optional] 
**DeliveryType** | Pointer to **int32** | * &#x60;0&#x60; None *&#x60;1&#x60; email *&#x60;2&#x60; SMS *&#x60;3&#x60; WhatsApp | [optional] 
**Destination** | Pointer to **string** | E-mail or Phone/Whatsapp number you want the order to be delivered to. | [optional] 
**TerminalId** | Pointer to **int32** | Terminal ID of the sub-users that can be defined in setting section in user panel | [optional] 
**TerminalPin** | Pointer to **int32** | Pin defined for sub-user | [optional] 
**ReferenceCode** | Pointer to **string** | Refernce code. A unique UUID V4 referece code must be included in request | [optional] 

## Methods

### NewOrdersPostRequest

`func NewOrdersPostRequest() *OrdersPostRequest`

NewOrdersPostRequest instantiates a new OrdersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersPostRequestWithDefaults

`func NewOrdersPostRequestWithDefaults() *OrdersPostRequest`

NewOrdersPostRequestWithDefaults instantiates a new OrdersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *OrdersPostRequest) GetSku() int32`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *OrdersPostRequest) GetSkuOk() (*int32, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *OrdersPostRequest) SetSku(v int32)`

SetSku sets Sku field to given value.

### HasSku

`func (o *OrdersPostRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetQuantity

`func (o *OrdersPostRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrdersPostRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrdersPostRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrdersPostRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPreOder

`func (o *OrdersPostRequest) GetPreOder() bool`

GetPreOder returns the PreOder field if non-nil, zero value otherwise.

### GetPreOderOk

`func (o *OrdersPostRequest) GetPreOderOk() (*bool, bool)`

GetPreOderOk returns a tuple with the PreOder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOder

`func (o *OrdersPostRequest) SetPreOder(v bool)`

SetPreOder sets PreOder field to given value.

### HasPreOder

`func (o *OrdersPostRequest) HasPreOder() bool`

HasPreOder returns a boolean if a field has been set.

### GetPrice

`func (o *OrdersPostRequest) GetPrice() decimal.Decimal`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrdersPostRequest) GetPriceOk() (*decimal.Decimal, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrdersPostRequest) SetPrice(v decimal.Decimal)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrdersPostRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDeliveryType

`func (o *OrdersPostRequest) GetDeliveryType() int32`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *OrdersPostRequest) GetDeliveryTypeOk() (*int32, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *OrdersPostRequest) SetDeliveryType(v int32)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *OrdersPostRequest) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDestination

`func (o *OrdersPostRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *OrdersPostRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *OrdersPostRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *OrdersPostRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetTerminalId

`func (o *OrdersPostRequest) GetTerminalId() int32`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *OrdersPostRequest) GetTerminalIdOk() (*int32, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *OrdersPostRequest) SetTerminalId(v int32)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *OrdersPostRequest) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetTerminalPin

`func (o *OrdersPostRequest) GetTerminalPin() int32`

GetTerminalPin returns the TerminalPin field if non-nil, zero value otherwise.

### GetTerminalPinOk

`func (o *OrdersPostRequest) GetTerminalPinOk() (*int32, bool)`

GetTerminalPinOk returns a tuple with the TerminalPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPin

`func (o *OrdersPostRequest) SetTerminalPin(v int32)`

SetTerminalPin sets TerminalPin field to given value.

### HasTerminalPin

`func (o *OrdersPostRequest) HasTerminalPin() bool`

HasTerminalPin returns a boolean if a field has been set.

### GetReferenceCode

`func (o *OrdersPostRequest) GetReferenceCode() string`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *OrdersPostRequest) GetReferenceCodeOk() (*string, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *OrdersPostRequest) SetReferenceCode(v string)`

SetReferenceCode sets ReferenceCode field to given value.

### HasReferenceCode

`func (o *OrdersPostRequest) HasReferenceCode() bool`

HasReferenceCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


