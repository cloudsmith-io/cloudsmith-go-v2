# ExperimentalFiles

## Overview

### Available Operations

* [Create](#create) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Requests upload URLs for a new package file.
* [Complete](#complete) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Completes a multipart file upload.
* [GetInfo](#getinfo) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Returns information needed to perform a multipart file upload.

## Create

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Requests upload URLs for a new package file.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/files/" -->
```go
package main

import(
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.ExperimentalFiles.Create(ctx, "<value>", "<value>", components.PackageFileUploadV2{
        Filename: "example.file",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageFileUploadInfoV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `repository`                                                                     | `string`                                                                         | :heavy_check_mark:                                                               | The `slug` of the repository within the workspace.                               |
| `workspace`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | The `slug` of the workspace.                                                     |
| `body`                                                                           | [components.PackageFileUploadV2](../../models/components/packagefileuploadv2.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.WorkspacesRepositoriesFilesCreateResponse](../../models/operations/workspacesrepositoriesfilescreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## Complete

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Completes a multipart file upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_complete_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/files/{identifier}/complete/" -->
```go
package main

import(
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.ExperimentalFiles.Complete(ctx, "<value>", "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageFileUploadResultV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` (identifier) of the package file.        |
| `repository`                                             | `string`                                                 | :heavy_check_mark:                                       | The `slug` of the repository within the workspace.       |
| `workspace`                                              | `string`                                                 | :heavy_check_mark:                                       | The `slug` of the workspace.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkspacesRepositoriesFilesCompleteCreateResponse](../../models/operations/workspacesrepositoriesfilescompletecreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404              | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## GetInfo

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Returns information needed to perform a multipart file upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_info_retrieve" method="get" path="/workspaces/{workspace}/repositories/{repository}/files/{identifier}/info/" -->
```go
package main

import(
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.ExperimentalFiles.GetInfo(ctx, operations.WorkspacesRepositoriesFilesInfoRetrieveRequest{
        Filename: "example.file",
        Identifier: "<value>",
        Repository: "<value>",
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageFilePartsUpload != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.WorkspacesRepositoriesFilesInfoRetrieveRequest](../../models/operations/workspacesrepositoriesfilesinforetrieverequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.WorkspacesRepositoriesFilesInfoRetrieveResponse](../../models/operations/workspacesrepositoriesfilesinforetrieveresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404              | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |