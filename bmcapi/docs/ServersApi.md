# \ServersApi

All URIs are relative to *https://api.phoenixnap.com/bmc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePrivateNetwork**](ServersApi.md#DeletePrivateNetwork) | **Delete** /servers/{serverId}/network-configuration/private-network-configuration/private-networks/{privateNetworkId} | Removes the server from private network.
[**ServersGet**](ServersApi.md#ServersGet) | **Get** /servers | List servers.
[**ServersPost**](ServersApi.md#ServersPost) | **Post** /servers | Create new server.
[**ServersServerIdActionsDeprovisionPost**](ServersApi.md#ServersServerIdActionsDeprovisionPost) | **Post** /servers/{serverId}/actions/deprovision | Deprovision a server.
[**ServersServerIdActionsPowerOffPost**](ServersApi.md#ServersServerIdActionsPowerOffPost) | **Post** /servers/{serverId}/actions/power-off | Power off server.
[**ServersServerIdActionsPowerOnPost**](ServersApi.md#ServersServerIdActionsPowerOnPost) | **Post** /servers/{serverId}/actions/power-on | Power on server.
[**ServersServerIdActionsRebootPost**](ServersApi.md#ServersServerIdActionsRebootPost) | **Post** /servers/{serverId}/actions/reboot | Reboot server.
[**ServersServerIdActionsReservePost**](ServersApi.md#ServersServerIdActionsReservePost) | **Post** /servers/{serverId}/actions/reserve | Reserve server.
[**ServersServerIdActionsResetPost**](ServersApi.md#ServersServerIdActionsResetPost) | **Post** /servers/{serverId}/actions/reset | Reset server.
[**ServersServerIdActionsShutdownPost**](ServersApi.md#ServersServerIdActionsShutdownPost) | **Post** /servers/{serverId}/actions/shutdown | Shutdown server.
[**ServersServerIdDelete**](ServersApi.md#ServersServerIdDelete) | **Delete** /servers/{serverId} | Delete server.
[**ServersServerIdGet**](ServersApi.md#ServersServerIdGet) | **Get** /servers/{serverId} | Get server.
[**ServersServerIdIpBlocksIpBlockIdDelete**](ServersApi.md#ServersServerIdIpBlocksIpBlockIdDelete) | **Delete** /servers/{serverId}/network-configuration/ip-block-configurations/ip-blocks/{ipBlockId} | Unassign IP Block from Server.
[**ServersServerIdIpBlocksPost**](ServersApi.md#ServersServerIdIpBlocksPost) | **Post** /servers/{serverId}/network-configuration/ip-block-configurations/ip-blocks | Assign IP Block to Server.
[**ServersServerIdPatch**](ServersApi.md#ServersServerIdPatch) | **Patch** /servers/{serverId} | Patch a Server.
[**ServersServerIdPrivateNetworksPatch**](ServersApi.md#ServersServerIdPrivateNetworksPatch) | **Patch** /servers/{serverId}/network-configuration/private-network-configuration/private-networks/{privateNetworkId} | Updates the server&#39;s private network&#39;s IP addresses
[**ServersServerIdPrivateNetworksPost**](ServersApi.md#ServersServerIdPrivateNetworksPost) | **Post** /servers/{serverId}/network-configuration/private-network-configuration/private-networks | Adds the server to a private network.
[**ServersServerIdPublicNetworksDelete**](ServersApi.md#ServersServerIdPublicNetworksDelete) | **Delete** /servers/{serverId}/network-configuration/public-network-configuration/public-networks/{publicNetworkId} | Removes the server from the Public Network
[**ServersServerIdPublicNetworksPatch**](ServersApi.md#ServersServerIdPublicNetworksPatch) | **Patch** /servers/{serverId}/network-configuration/public-network-configuration/public-networks/{publicNetworkId} | Updates the server&#39;s public network&#39;s IP addresses.
[**ServersServerIdPublicNetworksPost**](ServersApi.md#ServersServerIdPublicNetworksPost) | **Post** /servers/{serverId}/network-configuration/public-network-configuration/public-networks | Adds the server to a Public Network.
[**ServersServerIdTagsPut**](ServersApi.md#ServersServerIdTagsPut) | **Put** /servers/{serverId}/tags | Overwrite tags assigned for Server.



## DeletePrivateNetwork

> string DeletePrivateNetwork(ctx, serverId, privateNetworkId).Execute()

Removes the server from private network.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.DeletePrivateNetwork(context.Background(), serverId, privateNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.DeletePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePrivateNetwork`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.DeletePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersGet

> []Server ServersGet(ctx).Tag(tag).Execute()

List servers.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tag := []string{"Inner_example"} // []string | A list of query parameters related to tags in the form of tagName.tagValue (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersGet(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersGet`: []Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **[]string** | A list of query parameters related to tags in the form of tagName.tagValue | 

### Return type

[**[]Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersPost

> Server ServersPost(ctx).Force(force).ServerCreate(serverCreate).Execute()

Create new server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)
    serverCreate := *openapiclient.NewServerCreate("my-server-1", "ubuntu/bionic", "s1.c1.small", "PHX") // ServerCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersPost(context.Background()).Force(force).ServerCreate(serverCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersPost`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]
 **serverCreate** | [**ServerCreate**](ServerCreate.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsDeprovisionPost

> string ServersServerIdActionsDeprovisionPost(ctx, serverId).RelinquishIpBlock(relinquishIpBlock).Execute()

Deprovision a server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    relinquishIpBlock := *openapiclient.NewRelinquishIpBlock() // RelinquishIpBlock |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdActionsDeprovisionPost(context.Background(), serverId).RelinquishIpBlock(relinquishIpBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsDeprovisionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsDeprovisionPost`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsDeprovisionPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsDeprovisionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relinquishIpBlock** | [**RelinquishIpBlock**](RelinquishIpBlock.md) |  | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsPowerOffPost

> ActionResult ServersServerIdActionsPowerOffPost(ctx, serverId).Execute()

Power off server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdActionsPowerOffPost(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsPowerOffPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsPowerOffPost`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsPowerOffPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsPowerOffPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsPowerOnPost

> ActionResult ServersServerIdActionsPowerOnPost(ctx, serverId).Execute()

Power on server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdActionsPowerOnPost(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsPowerOnPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsPowerOnPost`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsPowerOnPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsPowerOnPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsRebootPost

> ActionResult ServersServerIdActionsRebootPost(ctx, serverId).Execute()

Reboot server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdActionsRebootPost(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsRebootPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsRebootPost`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsRebootPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsRebootPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsReservePost

> Server ServersServerIdActionsReservePost(ctx, serverId).ServerReserve(serverReserve).Execute()

Reserve server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    serverReserve := *openapiclient.NewServerReserve("ONE_MONTH_RESERVATION") // ServerReserve |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdActionsReservePost(context.Background(), serverId).ServerReserve(serverReserve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsReservePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsReservePost`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsReservePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsReservePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverReserve** | [**ServerReserve**](ServerReserve.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsResetPost

> ResetResult ServersServerIdActionsResetPost(ctx, serverId).ServerReset(serverReset).Execute()

Reset server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    serverReset := *openapiclient.NewServerReset() // ServerReset |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdActionsResetPost(context.Background(), serverId).ServerReset(serverReset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsResetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsResetPost`: ResetResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsResetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverReset** | [**ServerReset**](ServerReset.md) |  | 

### Return type

[**ResetResult**](ResetResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsShutdownPost

> ActionResult ServersServerIdActionsShutdownPost(ctx, serverId).Execute()

Shutdown server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdActionsShutdownPost(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsShutdownPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsShutdownPost`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsShutdownPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsShutdownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdDelete

> DeleteResult ServersServerIdDelete(ctx, serverId).Execute()

Delete server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdDelete(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdDelete`: DeleteResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResult**](DeleteResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdGet

> Server ServersServerIdGet(ctx, serverId).Execute()

Get server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdGet(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdGet`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdIpBlocksIpBlockIdDelete

> string ServersServerIdIpBlocksIpBlockIdDelete(ctx, serverId, ipBlockId).RelinquishIpBlock(relinquishIpBlock).Execute()

Unassign IP Block from Server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    ipBlockId := "6047127fed34ecc3ba8402d2" // string | The IP Block identifier.
    relinquishIpBlock := *openapiclient.NewRelinquishIpBlock() // RelinquishIpBlock |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdIpBlocksIpBlockIdDelete(context.Background(), serverId, ipBlockId).RelinquishIpBlock(relinquishIpBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdIpBlocksIpBlockIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdIpBlocksIpBlockIdDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdIpBlocksIpBlockIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdIpBlocksIpBlockIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **relinquishIpBlock** | [**RelinquishIpBlock**](RelinquishIpBlock.md) |  | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdIpBlocksPost

> ServerIpBlock ServersServerIdIpBlocksPost(ctx, serverId).ServerIpBlock(serverIpBlock).Execute()

Assign IP Block to Server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    serverIpBlock := *openapiclient.NewServerIpBlock("60473a6115e34466c9f8f083") // ServerIpBlock |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdIpBlocksPost(context.Background(), serverId).ServerIpBlock(serverIpBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdIpBlocksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdIpBlocksPost`: ServerIpBlock
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdIpBlocksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdIpBlocksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverIpBlock** | [**ServerIpBlock**](ServerIpBlock.md) |  | 

### Return type

[**ServerIpBlock**](ServerIpBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPatch

> Server ServersServerIdPatch(ctx, serverId).ServerPatch(serverPatch).Execute()

Patch a Server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    serverPatch := *openapiclient.NewServerPatch() // ServerPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdPatch(context.Background(), serverId).ServerPatch(serverPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdPatch`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverPatch** | [**ServerPatch**](ServerPatch.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPrivateNetworksPatch

> ServerPrivateNetwork ServersServerIdPrivateNetworksPatch(ctx, serverId, privateNetworkId).Force(force).ServerNetworkUpdate(serverNetworkUpdate).Execute()

Updates the server's private network's IP addresses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.
    force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)
    serverNetworkUpdate := *openapiclient.NewServerNetworkUpdate() // ServerNetworkUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdPrivateNetworksPatch(context.Background(), serverId, privateNetworkId).Force(force).ServerNetworkUpdate(serverNetworkUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdPrivateNetworksPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdPrivateNetworksPatch`: ServerPrivateNetwork
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdPrivateNetworksPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPrivateNetworksPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]
 **serverNetworkUpdate** | [**ServerNetworkUpdate**](ServerNetworkUpdate.md) |  | 

### Return type

[**ServerPrivateNetwork**](ServerPrivateNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPrivateNetworksPost

> ServerPrivateNetwork ServersServerIdPrivateNetworksPost(ctx, serverId).Force(force).ServerPrivateNetwork(serverPrivateNetwork).Execute()

Adds the server to a private network.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)
    serverPrivateNetwork := *openapiclient.NewServerPrivateNetwork("603f3b2cfcaf050643b89a4b") // ServerPrivateNetwork |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdPrivateNetworksPost(context.Background(), serverId).Force(force).ServerPrivateNetwork(serverPrivateNetwork).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdPrivateNetworksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdPrivateNetworksPost`: ServerPrivateNetwork
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdPrivateNetworksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPrivateNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]
 **serverPrivateNetwork** | [**ServerPrivateNetwork**](ServerPrivateNetwork.md) |  | 

### Return type

[**ServerPrivateNetwork**](ServerPrivateNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPublicNetworksDelete

> string ServersServerIdPublicNetworksDelete(ctx, serverId, publicNetworkId).Execute()

Removes the server from the Public Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdPublicNetworksDelete(context.Background(), serverId, publicNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdPublicNetworksDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdPublicNetworksDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdPublicNetworksDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPublicNetworksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPublicNetworksPatch

> ServerPublicNetwork ServersServerIdPublicNetworksPatch(ctx, serverId, publicNetworkId).Force(force).ServerNetworkUpdate(serverNetworkUpdate).Execute()

Updates the server's public network's IP addresses.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.
    force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)
    serverNetworkUpdate := *openapiclient.NewServerNetworkUpdate() // ServerNetworkUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdPublicNetworksPatch(context.Background(), serverId, publicNetworkId).Force(force).ServerNetworkUpdate(serverNetworkUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdPublicNetworksPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdPublicNetworksPatch`: ServerPublicNetwork
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdPublicNetworksPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPublicNetworksPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]
 **serverNetworkUpdate** | [**ServerNetworkUpdate**](ServerNetworkUpdate.md) |  | 

### Return type

[**ServerPublicNetwork**](ServerPublicNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPublicNetworksPost

> ServerPublicNetwork ServersServerIdPublicNetworksPost(ctx, serverId).Force(force).ServerPublicNetwork(serverPublicNetwork).Execute()

Adds the server to a Public Network.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)
    serverPublicNetwork := *openapiclient.NewServerPublicNetwork("60473c2509268bc77fd06d29") // ServerPublicNetwork |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdPublicNetworksPost(context.Background(), serverId).Force(force).ServerPublicNetwork(serverPublicNetwork).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdPublicNetworksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdPublicNetworksPost`: ServerPublicNetwork
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdPublicNetworksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPublicNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]
 **serverPublicNetwork** | [**ServerPublicNetwork**](ServerPublicNetwork.md) |  | 

### Return type

[**ServerPublicNetwork**](ServerPublicNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdTagsPut

> Server ServersServerIdTagsPut(ctx, serverId).TagAssignmentRequest(tagAssignmentRequest).Execute()

Overwrite tags assigned for Server.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    tagAssignmentRequest := []openapiclient.TagAssignmentRequest{*openapiclient.NewTagAssignmentRequest("Environment")} // []TagAssignmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersApi.ServersServerIdTagsPut(context.Background(), serverId).TagAssignmentRequest(tagAssignmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdTagsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdTagsPut`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdTagsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdTagsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagAssignmentRequest** | [**[]TagAssignmentRequest**](TagAssignmentRequest.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

