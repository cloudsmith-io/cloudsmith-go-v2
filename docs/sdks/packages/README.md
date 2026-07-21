# Packages

## Overview

### Available Operations

* [OsvSummary](#osvsummary) - API to provide OSV vulnerability data for a specific repo.
* [GetOsvNamespaceSummary](#getosvnamespacesummary) - API to provide OSV vulnerability data for a specific namespace.
* [GetOsvCounts](#getosvcounts) - API to provide OSV severity counts for a specific package.

## OsvSummary

API to provide OSV vulnerability data for a specific repo.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_osv_repo_summary_list" method="get" path="/packages/{owner}/{repo}/vulnerabilities/osv/summary/" -->
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

    res, err := s.Packages.OsvSummary(ctx, operations.PackagesVulnerabilitiesOsvRepoSummaryListRequest{
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

## GetOsvNamespaceSummary

API to provide OSV vulnerability data for a specific namespace.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_osv_namespace_summary_list" method="get" path="/packages/{owner}/vulnerabilities/osv/summary/" -->
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

    res, err := s.Packages.GetOsvNamespaceSummary(ctx, operations.PackagesVulnerabilitiesOsvNamespaceSummaryListRequest{
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

## GetOsvCounts

API to provide OSV severity counts for a specific package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_osv_counts_list" method="get" path="/packages/{package_slug_perm}/vulnerabilities/osv/counts/" -->
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