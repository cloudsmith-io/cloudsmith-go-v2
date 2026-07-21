# Packages

## Overview

### Available Operations

* [GetVulnerabilitiesOsvRepoSummary](#getvulnerabilitiesosvreposummary) - API to provide OSV vulnerability data for a specific repo.
* [GetVulnerabilitiesOsvNamespaceSummary](#getvulnerabilitiesosvnamespacesummary) - API to provide OSV vulnerability data for a specific namespace.
* [GetOsvVulnerabilities](#getosvvulnerabilities) - API to provide OSV vulnerability data for a specific package.
* [GetOsvCounts](#getosvcounts) - API to provide OSV severity counts for a specific package.
* [Delete](#delete) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Deletes a package.
* [UploadCran](#uploadcran) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new CRAN package
* [UploadDocker](#uploaddocker) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new Docker package
* [UploadGeneric](#uploadgeneric) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new Generic package
* [UploadHex](#uploadhex) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new Hex package
* [UploadHuggingface](#uploadhuggingface) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new HuggingFace package
* [UploadLuarocks](#uploadluarocks) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new LuaRocks package
* [UploadNpm](#uploadnpm) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new npm package
* [ValidateCargo](#validatecargo) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new Cargo package payload
* [ValidateHuggingface](#validatehuggingface) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new HuggingFace package payload
* [ValidateLuarocks](#validateluarocks) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new LuaRocks package payload
* [ValidateP2](#validatep2) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new P2 package payload

## GetVulnerabilitiesOsvRepoSummary

API to provide OSV vulnerability data for a specific repo.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_osv_repo_summary_list" method="get" path="/packages/{owner}/{repo}/vulnerabilities/osv/summary/" -->
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

    res, err := s.Packages.GetVulnerabilitiesOsvRepoSummary(ctx, operations.PackagesVulnerabilitiesOsvRepoSummaryListRequest{
        Owner: "<value>",
        Page: 495181,
        Repo: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedWebPackageVulnerabilitySummaryList != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.PackagesVulnerabilitiesOsvRepoSummaryListRequest](../../models/operations/packagesvulnerabilitiesosvreposummarylistrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.PackagesVulnerabilitiesOsvRepoSummaryListResponse](../../models/operations/packagesvulnerabilitiesosvreposummarylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVulnerabilitiesOsvNamespaceSummary

API to provide OSV vulnerability data for a specific namespace.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_osv_namespace_summary_list" method="get" path="/packages/{owner}/vulnerabilities/osv/summary/" -->
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

    res, err := s.Packages.GetVulnerabilitiesOsvNamespaceSummary(ctx, operations.PackagesVulnerabilitiesOsvNamespaceSummaryListRequest{
        Owner: "<value>",
        Page: 488651,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedWebPackageVulnerabilitySummaryList != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.PackagesVulnerabilitiesOsvNamespaceSummaryListRequest](../../models/operations/packagesvulnerabilitiesosvnamespacesummarylistrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |
| `opts`                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                             | :heavy_minus_sign:                                                                                                                                   | The options for this request.                                                                                                                        |

### Response

**[*operations.PackagesVulnerabilitiesOsvNamespaceSummaryListResponse](../../models/operations/packagesvulnerabilitiesosvnamespacesummarylistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetOsvVulnerabilities

API to provide OSV vulnerability data for a specific package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_osv_list" method="get" path="/packages/{package_slug_perm}/vulnerabilities/osv/" -->
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

    res, err := s.Packages.GetOsvVulnerabilities(ctx, operations.PackagesVulnerabilitiesOsvListRequest{
        PackageSlugPerm: "<value>",
        Page: 91700,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedWebOSVVulnerabilityList != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PackagesVulnerabilitiesOsvListRequest](../../models/operations/packagesvulnerabilitiesosvlistrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PackagesVulnerabilitiesOsvListResponse](../../models/operations/packagesvulnerabilitiesosvlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetOsvCounts

API to provide OSV severity counts for a specific package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_osv_counts_list" method="get" path="/packages/{package_slug_perm}/vulnerabilities/osv/counts/" -->
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

    res, err := s.Packages.GetOsvCounts(ctx, "<value>", cloudsmith.Pointer(true))
    if err != nil {
        log.Fatal(err)
    }
    if res.WebOSVSeverityCounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                               | Type                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                | Description                                                                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                      | The context to use for the request.                                                                                                                                                                                     |
| `packageSlugPerm`                                                                                                                                                                                                       | `string`                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                      | N/A                                                                                                                                                                                                                     |
| `primaryRecords`                                                                                                                                                                                                        | `*bool`                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                      | When true (default), only primary vulnerability records are returned, excluding duplicates identified by the primary record election process. When false, all matched records are returned including hidden duplicates. |
| `opts`                                                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                      | The options for this request.                                                                                                                                                                                           |

### Response

**[*operations.PackagesVulnerabilitiesOsvCountsListResponse](../../models/operations/packagesvulnerabilitiesosvcountslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Deletes a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_destroy" method="delete" path="/workspaces/{workspace}/repositories/{repository}/packages/{identifier}/" -->
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

    res, err := s.Packages.Delete(ctx, "<value>", "<value>", "<value>")
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
| `identifier`                                             | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` or `slug` (identifier) of the package.   |
| `repository`                                             | `string`                                                 | :heavy_check_mark:                                       | The `slug` of the repository within the workspace.       |
| `workspace`                                              | `string`                                                 | :heavy_check_mark:                                       | The `slug` of the workspace.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkspacesRepositoriesPackagesDestroyResponse](../../models/operations/workspacesrepositoriespackagesdestroyresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404              | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UploadCran

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new CRAN package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_cran_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/cran/" -->
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

    res, err := s.Packages.UploadCran(ctx, "<value>", "<value>", components.CranPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
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
| `body`                                                                           | [components.CranPackageUploadV2](../../models/components/cranpackageuploadv2.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadCranCreateResponse](../../models/operations/workspacesrepositoriespackagesuploadcrancreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UploadDocker

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new Docker package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_docker_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/docker/" -->
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

    res, err := s.Packages.UploadDocker(ctx, "<value>", "<value>", components.DockerPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `repository`                                                                         | `string`                                                                             | :heavy_check_mark:                                                                   | The `slug` of the repository within the workspace.                                   |
| `workspace`                                                                          | `string`                                                                             | :heavy_check_mark:                                                                   | The `slug` of the workspace.                                                         |
| `body`                                                                               | [components.DockerPackageUploadV2](../../models/components/dockerpackageuploadv2.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadDockerCreateResponse](../../models/operations/workspacesrepositoriespackagesuploaddockercreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UploadGeneric

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new Generic package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_generic_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/generic/" -->
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

    res, err := s.Packages.UploadGeneric(ctx, "<value>", "<value>", components.GenericPackageUploadV2{
        Filepath: "/private/tmp/slow.vst",
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `repository`                                                                           | `string`                                                                               | :heavy_check_mark:                                                                     | The `slug` of the repository within the workspace.                                     |
| `workspace`                                                                            | `string`                                                                               | :heavy_check_mark:                                                                     | The `slug` of the workspace.                                                           |
| `body`                                                                                 | [components.GenericPackageUploadV2](../../models/components/genericpackageuploadv2.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadGenericCreateResponse](../../models/operations/workspacesrepositoriespackagesuploadgenericcreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UploadHex

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new Hex package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_hex_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/hex/" -->
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

    res, err := s.Packages.UploadHex(ctx, "<value>", "<value>", components.HexPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `repository`                                                                   | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the repository within the workspace.                             |
| `workspace`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the workspace.                                                   |
| `body`                                                                         | [components.HexPackageUploadV2](../../models/components/hexpackageuploadv2.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadHexCreateResponse](../../models/operations/workspacesrepositoriespackagesuploadhexcreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UploadHuggingface

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new HuggingFace package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_huggingface_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/huggingface/" -->
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

    res, err := s.Packages.UploadHuggingface(ctx, "<value>", "<value>", components.HuggingfacePackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
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
| `body`                                                                                         | [components.HuggingfacePackageUploadV2](../../models/components/huggingfacepackageuploadv2.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadHuggingfaceCreateResponse](../../models/operations/workspacesrepositoriespackagesuploadhuggingfacecreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UploadLuarocks

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new LuaRocks package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_luarocks_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/luarocks/" -->
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

    res, err := s.Packages.UploadLuarocks(ctx, "<value>", "<value>", components.LuarocksPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `repository`                                                                             | `string`                                                                                 | :heavy_check_mark:                                                                       | The `slug` of the repository within the workspace.                                       |
| `workspace`                                                                              | `string`                                                                                 | :heavy_check_mark:                                                                       | The `slug` of the workspace.                                                             |
| `body`                                                                                   | [components.LuarocksPackageUploadV2](../../models/components/luarockspackageuploadv2.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadLuarocksCreateResponse](../../models/operations/workspacesrepositoriespackagesuploadluarockscreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UploadNpm

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new npm package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_npm_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/npm/" -->
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

    res, err := s.Packages.UploadNpm(ctx, "<value>", "<value>", components.NpmPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `repository`                                                                   | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the repository within the workspace.                             |
| `workspace`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the workspace.                                                   |
| `body`                                                                         | [components.NpmPackageUploadV2](../../models/components/npmpackageuploadv2.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadNpmCreateResponse](../../models/operations/workspacesrepositoriespackagesuploadnpmcreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateCargo

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new Cargo package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_cargo_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/cargo/" -->
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

    res, err := s.Packages.ValidateCargo(ctx, "<value>", "<value>", components.CargoPackageUploadV2{
        PackageFile: "<value>",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `repository`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | The `slug` of the repository within the workspace.                                 |
| `workspace`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | The `slug` of the workspace.                                                       |
| `body`                                                                             | [components.CargoPackageUploadV2](../../models/components/cargopackageuploadv2.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadCargoCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploadcargocreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateHuggingface

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new HuggingFace package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_huggingface_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/huggingface/" -->
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

    res, err := s.Packages.ValidateHuggingface(ctx, "<value>", "<value>", components.HuggingfacePackageUploadV2{
        PackageFile: "<value>",
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
| `body`                                                                                         | [components.HuggingfacePackageUploadV2](../../models/components/huggingfacepackageuploadv2.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadHuggingfaceCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploadhuggingfacecreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateLuarocks

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new LuaRocks package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_luarocks_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/luarocks/" -->
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

    res, err := s.Packages.ValidateLuarocks(ctx, "<value>", "<value>", components.LuarocksPackageUploadV2{
        PackageFile: "<value>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `repository`                                                                             | `string`                                                                                 | :heavy_check_mark:                                                                       | The `slug` of the repository within the workspace.                                       |
| `workspace`                                                                              | `string`                                                                                 | :heavy_check_mark:                                                                       | The `slug` of the workspace.                                                             |
| `body`                                                                                   | [components.LuarocksPackageUploadV2](../../models/components/luarockspackageuploadv2.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadLuarocksCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploadluarockscreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateP2

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new P2 package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_p2_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/p2/" -->
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

    res, err := s.Packages.ValidateP2(ctx, "<value>", "<value>", components.P2PackageUploadV2{
        PackageFile: "<value>",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `repository`                                                                 | `string`                                                                     | :heavy_check_mark:                                                           | The `slug` of the repository within the workspace.                           |
| `workspace`                                                                  | `string`                                                                     | :heavy_check_mark:                                                           | The `slug` of the workspace.                                                 |
| `body`                                                                       | [components.P2PackageUploadV2](../../models/components/p2packageuploadv2.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadP2CreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploadp2createresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |