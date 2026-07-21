# Workspaces.Policies.DecisionLogsV1

## Overview

### Available Operations

* [Get](#get) - Retrieve the full decision log JSON from S3.

## Get

Retrieve the full decision log JSON from S3.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_decision_logs_v1_retrieve" method="get" path="/workspaces/{workspace}/policies/decision-logs-v1/{decision_log_id}/" -->
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

    res, err := s.Workspaces.Policies.DecisionLogsV1.Get(ctx, "<id>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyDecisionLog != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `decisionLogID`                                                                             | `string`                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `workspace`                                                                                 | `string`                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `fields`                                                                                    | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Comma-separated fields to include in the response (e.g. `id,correlation_id,policy_output`). |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.WorkspacesPoliciesDecisionLogsV1RetrieveResponse](../../models/operations/workspacespoliciesdecisionlogsv1retrieveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |