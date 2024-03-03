# \DefaultAPI

All URIs are relative to *https://api.ezpaypin.com/vendors/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenPost**](DefaultAPI.md#AuthTokenPost) | **Post** /auth/token/ | 
[**BalanceGet**](DefaultAPI.md#BalanceGet) | **Get** /balance/ | 
[**CardsActivatePost**](DefaultAPI.md#CardsActivatePost) | **Post** /cards/activate/ | 
[**CardsCheckGet**](DefaultAPI.md#CardsCheckGet) | **Get** /cards/check/ | 
[**CardsReferenceCodeGet**](DefaultAPI.md#CardsReferenceCodeGet) | **Get** /cards/{reference_code}/ | 
[**CatalogsGet**](DefaultAPI.md#CatalogsGet) | **Get** /catalogs/ | 
[**CatalogsProductSkuAvailabilityGet**](DefaultAPI.md#CatalogsProductSkuAvailabilityGet) | **Get** /catalogs/{product_sku}/availability/ | 
[**CryptoCatalogGet**](DefaultAPI.md#CryptoCatalogGet) | **Get** /crypto/catalog/ | 
[**CryptoOrdersGet**](DefaultAPI.md#CryptoOrdersGet) | **Get** /crypto/orders/ | 
[**CryptoOrdersPost**](DefaultAPI.md#CryptoOrdersPost) | **Post** /crypto/orders/ | 
[**CryptoOrdersReferenceCodeFinalizePost**](DefaultAPI.md#CryptoOrdersReferenceCodeFinalizePost) | **Post** /crypto/orders/{reference_code}/finalize/ | 
[**CryptoOrdersReferenceCodeGet**](DefaultAPI.md#CryptoOrdersReferenceCodeGet) | **Get** /crypto/orders/{reference_code}/ | 
[**ExchangeRatesGet**](DefaultAPI.md#ExchangeRatesGet) | **Get** /exchange_rates/ | 
[**NotificationConfigGet**](DefaultAPI.md#NotificationConfigGet) | **Get** /notification_config/ | 
[**NotificationConfigPost**](DefaultAPI.md#NotificationConfigPost) | **Post** /notification_config/ | 
[**OrdersGet**](DefaultAPI.md#OrdersGet) | **Get** /orders/ | 
[**OrdersPost**](DefaultAPI.md#OrdersPost) | **Post** /orders/ | 
[**OrdersReferenceCodeCardsGet**](DefaultAPI.md#OrdersReferenceCodeCardsGet) | **Get** /orders/{reference_code}/cards/ | 
[**OrdersReferenceCodeGet**](DefaultAPI.md#OrdersReferenceCodeGet) | **Get** /orders/{reference_code}/ | 



## AuthTokenPost

> AuthTokenPost200Response AuthTokenPost(ctx).AuthTokenPostRequest(authTokenPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	authTokenPostRequest := *openapiclient.NewAuthTokenPostRequest() // AuthTokenPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AuthTokenPost(context.Background()).AuthTokenPostRequest(authTokenPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AuthTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenPost`: AuthTokenPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AuthTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authTokenPostRequest** | [**AuthTokenPostRequest**](AuthTokenPostRequest.md) |  | 

### Return type

[**AuthTokenPost200Response**](AuthTokenPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BalanceGet

> []Balance BalanceGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BalanceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BalanceGet`: []Balance
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BalanceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBalanceGetRequest struct via the builder pattern


### Return type

[**[]Balance**](Balance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CardsActivatePost

> Transaction CardsActivatePost(ctx).CardsActivatePostRequest(cardsActivatePostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	cardsActivatePostRequest := *openapiclient.NewCardsActivatePostRequest() // CardsActivatePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CardsActivatePost(context.Background()).CardsActivatePostRequest(cardsActivatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CardsActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CardsActivatePost`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CardsActivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCardsActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardsActivatePostRequest** | [**CardsActivatePostRequest**](CardsActivatePostRequest.md) |  | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CardsCheckGet

> CardsCheckGet200Response CardsCheckGet(ctx).Barcode(barcode).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	barcode := "888462551206058120006239903648" // string | Barcode Number Of Physical Card

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CardsCheckGet(context.Background()).Barcode(barcode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CardsCheckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CardsCheckGet`: CardsCheckGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CardsCheckGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCardsCheckGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **barcode** | **string** | Barcode Number Of Physical Card | 

### Return type

[**CardsCheckGet200Response**](CardsCheckGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CardsReferenceCodeGet

> Transaction CardsReferenceCodeGet(ctx, referenceCode).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	referenceCode := "046ed58d-53b1-4ec5-be46-693096a0bb98" // string | Reference Code of the activation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CardsReferenceCodeGet(context.Background(), referenceCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CardsReferenceCodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CardsReferenceCodeGet`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CardsReferenceCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceCode** | **string** | Reference Code of the activation | 

### Other Parameters

Other parameters are passed through a pointer to a apiCardsReferenceCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogsGet

> CatalogsGet200Response CatalogsGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CatalogsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CatalogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogsGet`: CatalogsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CatalogsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogsGetRequest struct via the builder pattern


### Return type

[**CatalogsGet200Response**](CatalogsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatalogsProductSkuAvailabilityGet

> CatalogsProductSkuAvailabilityGet200Response CatalogsProductSkuAvailabilityGet(ctx, productSku).ItemCount(itemCount).Price(price).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	productSku := int32(2785) // int32 | SKU of desired product
	itemCount := int32(10) // int32 | Number of items
	price := "10" // decimal.Decimal | Product price

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CatalogsProductSkuAvailabilityGet(context.Background(), productSku).ItemCount(itemCount).Price(price).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CatalogsProductSkuAvailabilityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatalogsProductSkuAvailabilityGet`: CatalogsProductSkuAvailabilityGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CatalogsProductSkuAvailabilityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productSku** | **int32** | SKU of desired product | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatalogsProductSkuAvailabilityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemCount** | **int32** | Number of items | 
 **price** | **decimal.Decimal** | Product price | 

### Return type

[**CatalogsProductSkuAvailabilityGet200Response**](CatalogsProductSkuAvailabilityGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCatalogGet

> CryptoCatalogGet200Response CryptoCatalogGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CryptoCatalogGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CryptoCatalogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CryptoCatalogGet`: CryptoCatalogGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CryptoCatalogGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCatalogGetRequest struct via the builder pattern


### Return type

[**CryptoCatalogGet200Response**](CryptoCatalogGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoOrdersGet

> CryptoOrdersGet200Response CryptoOrdersGet(ctx).Limit(limit).Offset(offset).StartDate(startDate).EndDate(endDate).Status(status).CanPay(canPay).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	limit := int32(10) // int32 | Maximum number of results (optional)
	offset := int32(1) // int32 | Offset number in results (optional)
	startDate := "2020-01-03T04:00:00.000Z" // string | Report Start Date (optional)
	endDate := "2020-01-07T11:00:00.000Z" // string | Report End Date (optional)
	status := int32(1) // int32 | Status of the order (optional)
	canPay := true // bool | Shows if order can be finalized and has not been rejected or expired (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CryptoOrdersGet(context.Background()).Limit(limit).Offset(offset).StartDate(startDate).EndDate(endDate).Status(status).CanPay(canPay).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CryptoOrdersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CryptoOrdersGet`: CryptoOrdersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CryptoOrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCryptoOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of results | 
 **offset** | **int32** | Offset number in results | 
 **startDate** | **string** | Report Start Date | 
 **endDate** | **string** | Report End Date | 
 **status** | **int32** | Status of the order | 
 **canPay** | **bool** | Shows if order can be finalized and has not been rejected or expired | 

### Return type

[**CryptoOrdersGet200Response**](CryptoOrdersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoOrdersPost

> CryptoCurrencyOrder CryptoOrdersPost(ctx).CryptoOrdersPostRequest(cryptoOrdersPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	cryptoOrdersPostRequest := *openapiclient.NewCryptoOrdersPostRequest() // CryptoOrdersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CryptoOrdersPost(context.Background()).CryptoOrdersPostRequest(cryptoOrdersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CryptoOrdersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CryptoOrdersPost`: CryptoCurrencyOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CryptoOrdersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCryptoOrdersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cryptoOrdersPostRequest** | [**CryptoOrdersPostRequest**](CryptoOrdersPostRequest.md) |  | 

### Return type

[**CryptoCurrencyOrder**](CryptoCurrencyOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoOrdersReferenceCodeFinalizePost

> CryptoOrdersReferenceCodeFinalizePost200Response CryptoOrdersReferenceCodeFinalizePost(ctx, referenceCode).CryptoOrdersReferenceCodeFinalizePostRequest(cryptoOrdersReferenceCodeFinalizePostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	referenceCode := "046ed58d-53b1-4ec5-be46-693096a0bb98" // string | Reference code of cryptocurrency order
	cryptoOrdersReferenceCodeFinalizePostRequest := *openapiclient.NewCryptoOrdersReferenceCodeFinalizePostRequest() // CryptoOrdersReferenceCodeFinalizePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CryptoOrdersReferenceCodeFinalizePost(context.Background(), referenceCode).CryptoOrdersReferenceCodeFinalizePostRequest(cryptoOrdersReferenceCodeFinalizePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CryptoOrdersReferenceCodeFinalizePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CryptoOrdersReferenceCodeFinalizePost`: CryptoOrdersReferenceCodeFinalizePost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CryptoOrdersReferenceCodeFinalizePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceCode** | **string** | Reference code of cryptocurrency order | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoOrdersReferenceCodeFinalizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cryptoOrdersReferenceCodeFinalizePostRequest** | [**CryptoOrdersReferenceCodeFinalizePostRequest**](CryptoOrdersReferenceCodeFinalizePostRequest.md) |  | 

### Return type

[**CryptoOrdersReferenceCodeFinalizePost200Response**](CryptoOrdersReferenceCodeFinalizePost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoOrdersReferenceCodeGet

> CryptoCurrencyOrder CryptoOrdersReferenceCodeGet(ctx, referenceCode).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	referenceCode := "046ed58d-53b1-4ec5-be46-693096a0bb98" // string | reference code of cryptocurrency order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CryptoOrdersReferenceCodeGet(context.Background(), referenceCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CryptoOrdersReferenceCodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CryptoOrdersReferenceCodeGet`: CryptoCurrencyOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CryptoOrdersReferenceCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceCode** | **string** | reference code of cryptocurrency order | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoOrdersReferenceCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CryptoCurrencyOrder**](CryptoCurrencyOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExchangeRatesGet

> ExchangeRatesGet200Response ExchangeRatesGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ExchangeRatesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExchangeRatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeRatesGet`: ExchangeRatesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ExchangeRatesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExchangeRatesGetRequest struct via the builder pattern


### Return type

[**ExchangeRatesGet200Response**](ExchangeRatesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationConfigGet

> Notification NotificationConfigGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.NotificationConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.NotificationConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationConfigGet`: Notification
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.NotificationConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationConfigGetRequest struct via the builder pattern


### Return type

[**Notification**](Notification.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationConfigPost

> NotificationConfigPost200Response NotificationConfigPost(ctx).Notification(notification).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	notification := *openapiclient.NewNotification() // Notification |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.NotificationConfigPost(context.Background()).Notification(notification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.NotificationConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationConfigPost`: NotificationConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.NotificationConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notification** | [**Notification**](Notification.md) |  | 

### Return type

[**NotificationConfigPost200Response**](NotificationConfigPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersGet

> OrdersGet200Response OrdersGet(ctx).Limit(limit).Offset(offset).StartDate(startDate).EndDate(endDate).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	limit := int32(10) // int32 | Number of items in list (optional)
	offset := int32(1) // int32 | Offset number of list (optional)
	startDate := "2020-01-03T04:00:00.000Z" // string | Report Start Date (optional)
	endDate := "2020-01-07T11:00:00.000Z" // string | Report End Date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrdersGet(context.Background()).Limit(limit).Offset(offset).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrdersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrdersGet`: OrdersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of items in list | 
 **offset** | **int32** | Offset number of list | 
 **startDate** | **string** | Report Start Date | 
 **endDate** | **string** | Report End Date | 

### Return type

[**OrdersGet200Response**](OrdersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersPost

> Order OrdersPost(ctx).OrdersPostRequest(ordersPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	ordersPostRequest := *openapiclient.NewOrdersPostRequest() // OrdersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrdersPost(context.Background()).OrdersPostRequest(ordersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrdersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrdersPost`: Order
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrdersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrdersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordersPostRequest** | [**OrdersPostRequest**](OrdersPostRequest.md) |  | 

### Return type

[**Order**](Order.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersReferenceCodeCardsGet

> OrdersReferenceCodeCardsGet200Response OrdersReferenceCodeCardsGet(ctx, referenceCode).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	referenceCode := "046ed58d-53b1-4ec5-be46-693096a0bb98" // string | Reference Code of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrdersReferenceCodeCardsGet(context.Background(), referenceCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrdersReferenceCodeCardsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrdersReferenceCodeCardsGet`: OrdersReferenceCodeCardsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrdersReferenceCodeCardsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceCode** | **string** | Reference Code of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersReferenceCodeCardsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrdersReferenceCodeCardsGet200Response**](OrdersReferenceCodeCardsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersReferenceCodeGet

> Order OrdersReferenceCodeGet(ctx, referenceCode).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/ezpin-go"
)

func main() {
	referenceCode := "046ed58d-53b1-4ec5-be46-693096a0bb98" // string | Reference Code of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.OrdersReferenceCodeGet(context.Background(), referenceCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OrdersReferenceCodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrdersReferenceCodeGet`: Order
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OrdersReferenceCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceCode** | **string** | Reference Code of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersReferenceCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Order**](Order.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

