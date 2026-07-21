# Experimental.Files

## Overview

### Available Operations

* [Abort](#abort) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Aborts a multipart file upload.
* [Validate](#validate) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates parameters used to request a file upload.

## Abort

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Aborts a multipart file upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_abort_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/files/{identifier}/abort/" -->
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

    res, err := s.Experimental.Files.Abort(ctx, "<value>", "<value>", "<value>")
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

## Validate

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates parameters used to request a file upload.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_files_validate_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/files/validate/" -->
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

    res, err := s.Experimental.Files.Validate(ctx, "<value>", "<value>", components.PackageFileUploadV2{
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `repository`                                                                     | `string`                                                                         | :heavy_check_mark:                                                               | The `slug` of the repository within the workspace.                               |
| `workspace`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | The `slug` of the workspace.                                                     |
| `body`                                                                           | [components.PackageFileUploadV2](../../models/components/packagefileuploadv2.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.WorkspacesRepositoriesFilesValidateCreateResponse](../../models/operations/workspacesrepositoriesfilesvalidatecreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |