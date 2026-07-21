# Experimental.Files

## Overview

### Available Operations

* [CreateFileUploadUrls](#createfileuploadurls) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Requests upload URLs for a new package file.
* [AbortUpload](#abortupload) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Aborts a multipart file upload.
* [CompleteUpload](#completeupload) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Completes a multipart file upload.
* [GetFileInfo](#getfileinfo) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Returns information needed to perform a multipart file upload.
* [ValidateFileParameters](#validatefileparameters) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates parameters used to request a file upload.

## CreateFileUploadUrls

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Requests upload URLs for a new package file.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/files/" -->
```go
package main

import(
	"context"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity(components.Security{
            Apikey: cloudsmith.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Experimental.Files.CreateFileUploadUrls(ctx, "<value>", "<value>", components.PackageFileUploadV2Request{
        Filename: "example.file",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageFileUploadV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `repository`                                                                                   | `string`                                                                                       | :heavy_check_mark:                                                                             | The `slug` of the repository within the workspace.                                             |
| `workspace`                                                                                    | `string`                                                                                       | :heavy_check_mark:                                                                             | The `slug` of the workspace.                                                                   |
| `body`                                                                                         | [components.PackageFileUploadV2Request](../../models/components/packagefileuploadv2request.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.WorkspacesRepositoriesFilesCreateResponse](../../models/operations/workspacesrepositoriesfilescreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## AbortUpload

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Aborts a multipart file upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_abort_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/files/{identifier}/abort/" -->
```go
package main

import(
	"context"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity(components.Security{
            Apikey: cloudsmith.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Experimental.Files.AbortUpload(ctx, "<value>", "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.WorkspacesRepositoriesFilesAbortCreateResponse](../../models/operations/workspacesrepositoriesfilesabortcreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404              | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## CompleteUpload

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Completes a multipart file upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_complete_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/files/{identifier}/complete/" -->
```go
package main

import(
	"context"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity(components.Security{
            Apikey: cloudsmith.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Experimental.Files.CompleteUpload(ctx, "<value>", "<value>", "<value>")
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

## GetFileInfo

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Returns information needed to perform a multipart file upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_info_retrieve" method="get" path="/workspaces/{workspace}/repositories/{repository}/files/{identifier}/info/" -->
```go
package main

import(
	"context"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity(components.Security{
            Apikey: cloudsmith.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Experimental.Files.GetFileInfo(ctx, operations.WorkspacesRepositoriesFilesInfoRetrieveRequest{
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

## ValidateFileParameters

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates parameters used to request a file upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_validate_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/files/validate/" -->
```go
package main

import(
	"context"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity(components.Security{
            Apikey: cloudsmith.Pointer("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Experimental.Files.ValidateFileParameters(ctx, "<value>", "<value>", components.PackageFileUploadV2Request{
        Filename: "example.file",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `repository`                                                                                   | `string`                                                                                       | :heavy_check_mark:                                                                             | The `slug` of the repository within the workspace.                                             |
| `workspace`                                                                                    | `string`                                                                                       | :heavy_check_mark:                                                                             | The `slug` of the workspace.                                                                   |
| `body`                                                                                         | [components.PackageFileUploadV2Request](../../models/components/packagefileuploadv2request.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.WorkspacesRepositoriesFilesValidateCreateResponse](../../models/operations/workspacesrepositoriesfilesvalidatecreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |