# Workspaces.DecisionLogs

## Overview

### Available Operations

* [List](#list) - List policy evaluation decision logs.

## List

List policy evaluation decision logs.

### Example Usage: ChronologicalOrder

<!-- UsageSnippet language="go" operationID="workspaces_policies_decision_logs_v1_list" method="get" path="/workspaces/{workspace}/policies/decision-logs-v1/" example="ChronologicalOrder" -->
```go
package main

import(
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/types"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Workspaces.DecisionLogs.List(ctx, operations.WorkspacesPoliciesDecisionLogsV1ListRequest{
        CreatedAfter: types.MustTimeFromString("2024-11-12T13:20:12.290Z"),
        Sort: cloudsmith.Pointer("id"),
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyDecisionLogV1SummaryList != nil {
        // handle response
    }
}
```
### Example Usage: ReverseChronologicalOrder(theDefault)

<!-- UsageSnippet language="go" operationID="workspaces_policies_decision_logs_v1_list" method="get" path="/workspaces/{workspace}/policies/decision-logs-v1/" example="ReverseChronologicalOrder(theDefault)" -->
```go
package main

import(
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/types"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Workspaces.DecisionLogs.List(ctx, operations.WorkspacesPoliciesDecisionLogsV1ListRequest{
        CreatedAfter: types.MustTimeFromString("2025-02-17T14:08:57.543Z"),
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyDecisionLogV1SummaryList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.WorkspacesPoliciesDecisionLogsV1ListRequest](../../models/operations/workspacespoliciesdecisionlogsv1listrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.WorkspacesPoliciesDecisionLogsV1ListResponse](../../models/operations/workspacespoliciesdecisionlogsv1listresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorDetail   | 400, 401, 403, 404, 422 | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |