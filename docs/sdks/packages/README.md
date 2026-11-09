# Packages

## Overview

### Available Operations

* [PackagesVulnerabilitiesList](#packagesvulnerabilitieslist) - OSV advisory data for a specific package.
* [PackagesVulnerabilitiesCountsList](#packagesvulnerabilitiescountslist) - OSV advisory counts for a specific package.

## PackagesVulnerabilitiesList

OSV advisory data for a specific package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_list" method="get" path="/packages/{package_slug_perm}/vulnerabilities/" -->
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

    res, err := s.Packages.PackagesVulnerabilitiesList(ctx, operations.PackagesVulnerabilitiesListRequest{
        PackageSlugPerm: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedWebOSVVulnerabilityList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PackagesVulnerabilitiesListRequest](../../models/operations/packagesvulnerabilitieslistrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PackagesVulnerabilitiesListResponse](../../models/operations/packagesvulnerabilitieslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PackagesVulnerabilitiesCountsList

OSV advisory counts for a specific package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_counts_list" method="get" path="/packages/{package_slug_perm}/vulnerabilities/counts/" -->
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

    res, err := s.Packages.PackagesVulnerabilitiesCountsList(ctx, "<value>", cloudsmith.Pointer(true))
    if err != nil {
        log.Fatal(err)
    }
    if res.WebOSVSeverityCounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                               | Type                                                                                                                                                                                    | Required                                                                                                                                                                                | Description                                                                                                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                      | The context to use for the request.                                                                                                                                                     |
| `packageSlugPerm`                                                                                                                                                                       | `string`                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                      | N/A                                                                                                                                                                                     |
| `primaryRecords`                                                                                                                                                                        | `*bool`                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                      | When true (default), only primary vulnerability records are returned, excluding records that describe the same flaw. When false, all matched records are returned including duplicates. |
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |

### Response

**[*operations.PackagesVulnerabilitiesCountsListResponse](../../models/operations/packagesvulnerabilitiescountslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |