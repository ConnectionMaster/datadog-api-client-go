# datadog-api-client-go

This repository contains a Go API client for the [Datadog API](https://docs.datadoghq.com/api/).
The code is generated using [openapi-generator](https://github.com/OpenAPITools/openapi-generator)
and [apigentools](https://github.com/DataDog/apigentools).

## Requirements

- Go 1.14+

## Layout

This repository contains per-major-version API client packages. Right
now, Datadog has two API versions, `v1` and `v2`.

### The API v1 Client

The client library for Datadog API v1 is located in the `api/v1/datadog` directory. Import it with

```go
import "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
```

All the documentation for this package is available [here](api/v1/datadog/README.md).

### The API v2 Client

The client library for Datadog API v2 is located in the `api/v2/datadog` directory. Import it with

```go
import "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
```

All the documentation for this package is available [here](api/v2/datadog/README.md).

## Getting Started

Here's an example creating a user:

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    body := *datadog.NewUserCreateRequest(*datadog.NewUserCreateData(*datadog.NewUserCreateAttributes("jane.doe@example.com"), datadog.UsersType("users")))

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUser(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating user: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    responseData := resp.GetData()
    fmt.Fprintf(os.Stdout, "User ID: %s", responseData.GetId())
}
```

Save it to `example.go`, then run `go get github.com/DataDog/datadog-api-client-go/api/v2/datadog`.
Set the `DD_CLIENT_API_KEY` and `DD_CLIENT_APP_KEY` to your Datadog
credentials, and then run `go run example.go`.

### Unstable Endpoints

This client includes access to Datadog API endpoints while they are in an unstable state and may undergo breaking changes. An extra configuration step is required to enable these endpoints:

```go
    configuration.SetUnstableOperationEnabled("<OperationName>", true)
```

where `<OperationName>` is the name of the method used to interact with that endpoint. For example: `GetLogsIndex`, or `UpdateLogsIndex`

### Changing Server

When talking to a different server, like the `eu` instance, change the `ContextServerVariables`:

```go
    ctx = context.WithValue(ctx,
        datadog.ContextServerVariables,
        map[string]string{
            "site": "datadoghq.eu",
    })
```

## Documentation

Documentation for API endpoints and models can be found under the docs subdirectories, in [v1](/api/v1/datadog#documentation-for-api-endpoints)
and [v2](/api/v2/datadog#documentation-for-api-endpoints).

It's also available on [pkg.go.dev](https://pkg.go.dev/github.com/DataDog/datadog-api-client-go).

## Contributing

As most of the code in this repository is generated, we will only accept PRs for files
that are not modified by our code-generation machinery (changes to the generated files
would get overwritten). We happily accept contributions to files that are not autogenerated,
such as tests and development tooling.

## Author

support@datadoghq.com
