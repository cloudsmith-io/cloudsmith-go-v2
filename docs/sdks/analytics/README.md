# Analytics

## Overview

### Available Operations

* [ListClientLogs](#listclientlogs) - Returns individual log-level records of client downloads, including bandwidth in bytes. Available for _Manager_ or _Owner_ level users only. All times in UTC.
* [ListPackageLogs](#listpackagelogs) - Returns package log events, including uploads, copies, deletes, etc.
* [GetClientTimeSeriesMetrics](#getclienttimeseriesmetrics) - API to provide usage data derived from client logs in time series format. Available for _Member_, _Manager_ and _Owner_ level Users. All times in UTC.

## ListClientLogs

Returns individual log-level records of client downloads, including bandwidth in bytes. Available for _Manager_ or _Owner_ level users only. All times in UTC.

### Example Usage

<!-- UsageSnippet language="go" operationID="analytics_logs_client_list" method="get" path="/analytics/logs/client/{workspace}/" -->
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

    res, err := s.Analytics.ListClientLogs(ctx, operations.AnalyticsLogsClientListRequest{
        Page: 408040,
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedLogsResponseList != nil {
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.AnalyticsLogsClientListRequest](../../models/operations/analyticslogsclientlistrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.AnalyticsLogsClientListResponse](../../models/operations/analyticslogsclientlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListPackageLogs

Returns package log events, including uploads, copies, deletes, etc.

### Example Usage

<!-- UsageSnippet language="go" operationID="analytics_logs_package_list" method="get" path="/analytics/logs/package/{workspace}/" example="PackageNamedSemverOrAnyRuby-formatPackage" -->
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

    res, err := s.Analytics.ListPackageLogs(ctx, operations.AnalyticsLogsPackageListRequest{
        Page: 557482,
        Query: cloudsmith.Pointer("package:semver OR package_format:ruby"),
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPackageLogsResponseList != nil {
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

## GetClientTimeSeriesMetrics

API to provide usage data derived from client logs in time series format. Available for _Member_, _Manager_ and _Owner_ level Users. All times in UTC.

### Example Usage

<!-- UsageSnippet language="go" operationID="analytics_metrics_client_time_series_list" method="get" path="/analytics/metrics/client/time-series/{workspace}/" -->
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

    res, err := s.Analytics.GetClientTimeSeriesMetrics(ctx, operations.AnalyticsMetricsClientTimeSeriesListRequest{
        Interval: operations.IntervalFifteenMinutes,
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsTimeSeriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.AnalyticsMetricsClientTimeSeriesListRequest](../../models/operations/analyticsmetricsclienttimeserieslistrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.AnalyticsMetricsClientTimeSeriesListResponse](../../models/operations/analyticsmetricsclienttimeserieslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |