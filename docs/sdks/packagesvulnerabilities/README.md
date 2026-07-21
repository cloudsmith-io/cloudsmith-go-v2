# Packages.Vulnerabilities

## Overview

### Available Operations

* [ListOsv](#listosv) - API to provide OSV vulnerability data for a specific package.

## ListOsv

API to provide OSV vulnerability data for a specific package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_vulnerabilities_osv_list" method="get" path="/packages/{package_slug_perm}/vulnerabilities/osv/" -->
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

    res, err := s.Packages.Vulnerabilities.ListOsv(ctx, operations.PackagesVulnerabilitiesOsvListRequest{
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