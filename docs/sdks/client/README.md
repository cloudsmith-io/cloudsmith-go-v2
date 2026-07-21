# Analytics.Metrics.Client

## Overview

### Available Operations

* [List](#list) - API to provide usage data derived from client logs as single aggregated values. Available for _Member_, _Manager_ and _Owner_ level Users. All times in UTC.

## List

API to provide usage data derived from client logs as single aggregated values. Available for _Member_, _Manager_ and _Owner_ level Users. All times in UTC.

### Example Usage

<!-- UsageSnippet language="go" operationID="analytics_metrics_client_list" method="get" path="/analytics/metrics/client/{workspace}/" -->
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

    res, err := s.Analytics.Metrics.Client.List(ctx, operations.AnalyticsMetricsClientListRequest{
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsSingleValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.AnalyticsMetricsClientListRequest](../../models/operations/analyticsmetricsclientlistrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AnalyticsMetricsClientListResponse](../../models/operations/analyticsmetricsclientlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |