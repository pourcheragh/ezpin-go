# OrdersPost400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | detail of returned response | [optional] 
**Code** | Pointer to **int32** | error code | [optional] 

## Methods

### NewOrdersPost400Response

`func NewOrdersPost400Response() *OrdersPost400Response`

NewOrdersPost400Response instantiates a new OrdersPost400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersPost400ResponseWithDefaults

`func NewOrdersPost400ResponseWithDefaults() *OrdersPost400Response`

NewOrdersPost400ResponseWithDefaults instantiates a new OrdersPost400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *OrdersPost400Response) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *OrdersPost400Response) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *OrdersPost400Response) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *OrdersPost400Response) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetCode

`func (o *OrdersPost400Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OrdersPost400Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OrdersPost400Response) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *OrdersPost400Response) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


