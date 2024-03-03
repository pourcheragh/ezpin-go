# OrdersPost406Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | detail of returned response | [optional] 
**Code** | Pointer to **int32** | error code | [optional] 

## Methods

### NewOrdersPost406Response

`func NewOrdersPost406Response() *OrdersPost406Response`

NewOrdersPost406Response instantiates a new OrdersPost406Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersPost406ResponseWithDefaults

`func NewOrdersPost406ResponseWithDefaults() *OrdersPost406Response`

NewOrdersPost406ResponseWithDefaults instantiates a new OrdersPost406Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *OrdersPost406Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *OrdersPost406Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *OrdersPost406Response) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *OrdersPost406Response) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetCode

`func (o *OrdersPost406Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OrdersPost406Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OrdersPost406Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *OrdersPost406Response) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


