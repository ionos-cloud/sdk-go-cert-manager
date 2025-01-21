# \AutoCertificateApi

All URIs are relative to *https://certificate-manager.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**AutoCertificatesDelete**](AutoCertificateApi.md#AutoCertificatesDelete) | **Delete** /auto-certificates/{autoCertificateId} | Delete AutoCertificate|
|[**AutoCertificatesFindById**](AutoCertificateApi.md#AutoCertificatesFindById) | **Get** /auto-certificates/{autoCertificateId} | Retrieve AutoCertificate|
|[**AutoCertificatesGet**](AutoCertificateApi.md#AutoCertificatesGet) | **Get** /auto-certificates | Retrieve all AutoCertificate|
|[**AutoCertificatesPatch**](AutoCertificateApi.md#AutoCertificatesPatch) | **Patch** /auto-certificates/{autoCertificateId} | Updates AutoCertificate|
|[**AutoCertificatesPost**](AutoCertificateApi.md#AutoCertificatesPost) | **Post** /auto-certificates | Create AutoCertificate|



## AutoCertificatesDelete

```go
var result  = AutoCertificatesDelete(ctx, autoCertificateId)
                      .Execute()
```

Delete AutoCertificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-cert-manager"
)

func main() {
    autoCertificateId := "f88467f8-a2d6-5871-83b9-e10f23d0a48a" // string | The ID (UUID) of the AutoCertificate.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.AutoCertificateApi.AutoCertificatesDelete(context.Background(), autoCertificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoCertificateApi.AutoCertificatesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**autoCertificateId** | **string** | The ID (UUID) of the AutoCertificate. | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoCertificatesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AutoCertificateApiService.AutoCertificatesDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AutoCertificateApiService.AutoCertificatesDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AutoCertificateApiService.AutoCertificatesDelete": {
    "port": "8443",
},
})
```


## AutoCertificatesFindById

```go
var result AutoCertificateRead = AutoCertificatesFindById(ctx, autoCertificateId)
                      .Execute()
```

Retrieve AutoCertificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-cert-manager"
)

func main() {
    autoCertificateId := "f88467f8-a2d6-5871-83b9-e10f23d0a48a" // string | The ID (UUID) of the AutoCertificate.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AutoCertificateApi.AutoCertificatesFindById(context.Background(), autoCertificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoCertificateApi.AutoCertificatesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoCertificatesFindById`: AutoCertificateRead
    fmt.Fprintf(os.Stdout, "Response from `AutoCertificateApi.AutoCertificatesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**autoCertificateId** | **string** | The ID (UUID) of the AutoCertificate. | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoCertificatesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**AutoCertificateRead**](../models/AutoCertificateRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AutoCertificateApiService.AutoCertificatesFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AutoCertificateApiService.AutoCertificatesFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AutoCertificateApiService.AutoCertificatesFindById": {
    "port": "8443",
},
})
```


## AutoCertificatesGet

```go
var result AutoCertificateReadList = AutoCertificatesGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterCommonName(filterCommonName)
                      .Execute()
```

Retrieve all AutoCertificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-cert-manager"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    filterCommonName := "www.example.com" // string | Filter by the common name (DNS).  (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AutoCertificateApi.AutoCertificatesGet(context.Background()).Offset(offset).Limit(limit).FilterCommonName(filterCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoCertificateApi.AutoCertificatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoCertificatesGet`: AutoCertificateReadList
    fmt.Fprintf(os.Stdout, "Response from `AutoCertificateApi.AutoCertificatesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiAutoCertificatesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **filterCommonName** | **string** | Filter by the common name (DNS).  | |

### Return type

[**AutoCertificateReadList**](../models/AutoCertificateReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AutoCertificateApiService.AutoCertificatesGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AutoCertificateApiService.AutoCertificatesGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AutoCertificateApiService.AutoCertificatesGet": {
    "port": "8443",
},
})
```


## AutoCertificatesPatch

```go
var result AutoCertificateRead = AutoCertificatesPatch(ctx, autoCertificateId)
                      .AutoCertificatePatch(autoCertificatePatch)
                      .Execute()
```

Updates AutoCertificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-cert-manager"
)

func main() {
    autoCertificateId := "f88467f8-a2d6-5871-83b9-e10f23d0a48a" // string | The ID (UUID) of the AutoCertificate.
    autoCertificatePatch := *openapiclient.NewAutoCertificatePatch(*openapiclient.NewPatchName("My name")) // AutoCertificatePatch | patch AutoCertificate

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AutoCertificateApi.AutoCertificatesPatch(context.Background(), autoCertificateId).AutoCertificatePatch(autoCertificatePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoCertificateApi.AutoCertificatesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoCertificatesPatch`: AutoCertificateRead
    fmt.Fprintf(os.Stdout, "Response from `AutoCertificateApi.AutoCertificatesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**autoCertificateId** | **string** | The ID (UUID) of the AutoCertificate. | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoCertificatesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **autoCertificatePatch** | [**AutoCertificatePatch**](../models/AutoCertificatePatch.md) | patch AutoCertificate | |

### Return type

[**AutoCertificateRead**](../models/AutoCertificateRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AutoCertificateApiService.AutoCertificatesPatch"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AutoCertificateApiService.AutoCertificatesPatch": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AutoCertificateApiService.AutoCertificatesPatch": {
    "port": "8443",
},
})
```


## AutoCertificatesPost

```go
var result AutoCertificateRead = AutoCertificatesPost(ctx)
                      .AutoCertificateCreate(autoCertificateCreate)
                      .Execute()
```

Create AutoCertificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-cert-manager"
)

func main() {
    autoCertificateCreate := *openapiclient.NewAutoCertificateCreate(*openapiclient.NewAutoCertificate("b471cd03-ef51-52c5-91a5-49195b0a04d4", "www.example.com", "rsa4096", "My Auto renewed certificate")) // AutoCertificateCreate | AutoCertificate to create.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.AutoCertificateApi.AutoCertificatesPost(context.Background()).AutoCertificateCreate(autoCertificateCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoCertificateApi.AutoCertificatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoCertificatesPost`: AutoCertificateRead
    fmt.Fprintf(os.Stdout, "Response from `AutoCertificateApi.AutoCertificatesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiAutoCertificatesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **autoCertificateCreate** | [**AutoCertificateCreate**](../models/AutoCertificateCreate.md) | AutoCertificate to create. | |

### Return type

[**AutoCertificateRead**](../models/AutoCertificateRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"AutoCertificateApiService.AutoCertificatesPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "AutoCertificateApiService.AutoCertificatesPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "AutoCertificateApiService.AutoCertificatesPost": {
    "port": "8443",
},
})
```

