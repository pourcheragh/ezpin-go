# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** | The URL that status change will be reported to | [optional] 
**ConfidentialKey** | Pointer to **string** | The confidential key that you need for validation in your system | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *Notification) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Notification) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Notification) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Notification) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetConfidentialKey

`func (o *Notification) GetConfidentialKey() string`

GetConfidentialKey returns the ConfidentialKey field if non-nil, zero value otherwise.

### GetConfidentialKeyOk

`func (o *Notification) GetConfidentialKeyOk() (*string, bool)`

GetConfidentialKeyOk returns a tuple with the ConfidentialKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialKey

`func (o *Notification) SetConfidentialKey(v string)`

SetConfidentialKey sets ConfidentialKey field to given value.

### HasConfidentialKey

`func (o *Notification) HasConfidentialKey() bool`

HasConfidentialKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


