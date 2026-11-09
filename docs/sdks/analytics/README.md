# Analytics

## Overview

### Available Operations

* [AnalyticsLogsPackageList](#analyticslogspackagelist) - Returns package log events, including uploads, copies, deletes, etc.

## AnalyticsLogsPackageList

Returns package log events, including uploads, copies, deletes, etc.

### Example Usage

<!-- UsageSnippet language="go" operationID="analytics_logs_package_list" method="get" path="/analytics/logs/package/{workspace}/" example="PackageNamedSemverOrAnyRuby-formatPackage" -->
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

    res, err := s.Analytics.AnalyticsLogsPackageList(ctx, operations.AnalyticsLogsPackageListRequest{
        Query: cloudsmith.Pointer("package:semver OR package_format:ruby"),
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPackageLogsResponseList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AnalyticsLogsPackageListRequest](../../models/operations/analyticslogspackagelistrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AnalyticsLogsPackageListResponse](../../models/operations/analyticslogspackagelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |