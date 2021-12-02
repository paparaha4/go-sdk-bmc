# Go API client for tagapi

The Tags Manager API. </br></br>**All URLs are relative to (https://api.phoenixnap.com/tag-manager/v1/)**

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build date: 2021-12-02T15:04:17.392906Z[Etc/UTC]
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://phoenixnap.com/](https://phoenixnap.com/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./tagapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.phoenixnap.com/tag-manager/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*TagsApi* | [**TagsGet**](docs/TagsApi.md#tagsget) | **Get** /tags | List tags.
*TagsApi* | [**TagsPost**](docs/TagsApi.md#tagspost) | **Post** /tags | Create a Tag.
*TagsApi* | [**TagsTagIdDelete**](docs/TagsApi.md#tagstagiddelete) | **Delete** /tags/{tagId} | Delete a Tag.
*TagsApi* | [**TagsTagIdGet**](docs/TagsApi.md#tagstagidget) | **Get** /tags/{tagId} | Get a Tag.
*TagsApi* | [**TagsTagIdPatch**](docs/TagsApi.md#tagstagidpatch) | **Patch** /tags/{tagId} | Modify a Tag.


## Documentation For Models

 - [DeleteResult](docs/DeleteResult.md)
 - [Error](docs/Error.md)
 - [ResourceAssignment](docs/ResourceAssignment.md)
 - [Tag](docs/Tag.md)
 - [TagCreate](docs/TagCreate.md)
 - [TagUpdate](docs/TagUpdate.md)


## Documentation For Authorization



### OAuth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **tags**: Grants full access to tags-api.
 - **tags.read**: Grants read only access to tags-api.

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@phoenixnap.com
