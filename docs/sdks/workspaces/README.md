# Workspaces

## Overview

### Available Operations

* [CreatePolicy](#createpolicy) - Create a policy.
* [RetrievePolicy](#retrievepolicy) - Retrieve a policy.
* [UpdatePolicy](#updatepolicy) - Update a policy.
* [DeletePolicy](#deletepolicy) - Delete a policy.
* [GetPolicyActions](#getpolicyactions) - List all actions for the policy.
* [UpdatePolicyAction](#updatepolicyaction) - Update an action for a policy.
* [PatchPolicyAction](#patchpolicyaction) - Partially update an action for a policy.
* [DeletePolicyAction](#deletepolicyaction) - Delete an action for a policy.
* [SimulatePolicy](#simulatepolicy) - Simulates the execution of a Policy to verify its behavior without taking any actions.

Before using this endpoint, the Policy to simulate must be created with the creation endpoint. To experiment
with the policy, it is recommended to create the policy with disabled=True, so that it does not take effect
outside the simulation endpoint.

This endpoint evaluates all packages in the workspace, generating a "decision log" for each evaluation.
Each log contains:
- Package metadata provided to the policy engine at runtime.
- Output from the user-defined Rego policy.

No actions associated with the policy are executed. Instead, the endpoint reports what would happen to each package if the policy were active.
* [GetPoliciesDecisionLogsV1](#getpoliciesdecisionlogsv1) - List policy evaluation decision logs.
* [GetDecisionLog](#getdecisionlog) - Retrieve the full decision log JSON from S3.
* [GetDecisionLogFacets](#getdecisionlogfacets) - Internal endpoint that returns counts of policy decision logs grouped by a
single dimension. Used to populate the Policy Decision Logs filter dropdowns in
the web app with the full set of entities that had activity in the timeframe.

Excluded from the public OpenAPI spec in prd; visible in dev/stg.

## CreatePolicy

Create a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_create" method="post" path="/workspaces/{workspace}/policies/" example="MaxPoliciesReached" -->
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

    res, err := s.Workspaces.CreatePolicy(ctx, "<value>", components.PolicyRequest{
        Name: "<value>",
        Rego: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Policy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `workspace`                                                          | `string`                                                             | :heavy_check_mark:                                                   | The `name` of the Workspace.                                         |
| `body`                                                               | [components.PolicyRequest](../../models/components/policyrequest.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.WorkspacesPoliciesCreateResponse](../../models/operations/workspacespoliciescreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 422                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## RetrievePolicy

Retrieve a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_retrieve" method="get" path="/workspaces/{workspace}/policies/{policy}/" -->
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

    res, err := s.Workspaces.RetrievePolicy(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Policy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policy`                                                 | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` (identifier) of the Policy.              |
| `workspace`                                              | `string`                                                 | :heavy_check_mark:                                       | The `name` of the Workspace.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkspacesPoliciesRetrieveResponse](../../models/operations/workspacespoliciesretrieveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdatePolicy

Update a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_update" method="put" path="/workspaces/{workspace}/policies/{policy}/" -->
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

    res, err := s.Workspaces.UpdatePolicy(ctx, "<value>", "<value>", components.PolicyRequest{
        Name: "<value>",
        Rego: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Policy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `policy`                                                             | `string`                                                             | :heavy_check_mark:                                                   | The `slug_perm` (identifier) of the Policy.                          |
| `workspace`                                                          | `string`                                                             | :heavy_check_mark:                                                   | The `name` of the Workspace.                                         |
| `body`                                                               | [components.PolicyRequest](../../models/components/policyrequest.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.WorkspacesPoliciesUpdateResponse](../../models/operations/workspacespoliciesupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeletePolicy

Delete a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_destroy" method="delete" path="/workspaces/{workspace}/policies/{policy}/" -->
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

    res, err := s.Workspaces.DeletePolicy(ctx, "<value>", "<value>")
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
| `policy`                                                 | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` (identifier) of the Policy.              |
| `workspace`                                              | `string`                                                 | :heavy_check_mark:                                       | The `name` of the Workspace.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkspacesPoliciesDestroyResponse](../../models/operations/workspacespoliciesdestroyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetPolicyActions

List all actions for the policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_actions_list" method="get" path="/workspaces/{workspace}/policies/{policy}/actions/" -->
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

    res, err := s.Workspaces.GetPolicyActions(ctx, 452996, "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyActionList != nil {
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `page`                                                   | `int64`                                                  | :heavy_check_mark:                                       | A page number within the paginated result set.           |
| `policy`                                                 | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` (identifier) of the Policy.              |
| `workspace`                                              | `string`                                                 | :heavy_check_mark:                                       | The `name` of the Workspace.                             |
| `pageSize`                                               | `*int64`                                                 | :heavy_minus_sign:                                       | Number of results to return per page.                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkspacesPoliciesActionsListResponse](../../models/operations/workspacespoliciesactionslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdatePolicyAction

Update an action for a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_actions_update" method="put" path="/workspaces/{workspace}/policies/{policy}/actions/{action}/" -->
```go
package main

import(
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"log"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Workspaces.UpdatePolicyAction(ctx, "<value>", "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyAction != nil {
        switch res.PolicyAction.Type {
            case components.PolicyActionTypeSetPackageState:
                // res.PolicyAction.SetPackageStatePolicyActionTyped is populated
            case components.PolicyActionTypeAddPackageTags:
                // res.PolicyAction.AddPackageTagsPolicyActionTyped is populated
            case components.PolicyActionTypeRemovePackageTags:
                // res.PolicyAction.RemovePackageTagsPolicyActionTyped is populated
            default:
                // Unknown type - use res.PolicyAction.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `action`                                                                          | `string`                                                                          | :heavy_check_mark:                                                                | The `slug_perm` (identifier) of the Action.                                       |
| `policy`                                                                          | `string`                                                                          | :heavy_check_mark:                                                                | The `slug_perm` (identifier) of the Policy.                                       |
| `workspace`                                                                       | `string`                                                                          | :heavy_check_mark:                                                                | The `name` of the Workspace.                                                      |
| `body`                                                                            | [*components.PolicyActionRequest](../../models/components/policyactionrequest.md) | :heavy_minus_sign:                                                                | N/A                                                                               |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.WorkspacesPoliciesActionsUpdateResponse](../../models/operations/workspacespoliciesactionsupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PatchPolicyAction

Partially update an action for a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_actions_partial_update" method="patch" path="/workspaces/{workspace}/policies/{policy}/actions/{action}/" -->
```go
package main

import(
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"log"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Workspaces.PatchPolicyAction(ctx, "<value>", "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyAction != nil {
        switch res.PolicyAction.Type {
            case components.PolicyActionTypeSetPackageState:
                // res.PolicyAction.SetPackageStatePolicyActionTyped is populated
            case components.PolicyActionTypeAddPackageTags:
                // res.PolicyAction.AddPackageTagsPolicyActionTyped is populated
            case components.PolicyActionTypeRemovePackageTags:
                // res.PolicyAction.RemovePackageTagsPolicyActionTyped is populated
            default:
                // Unknown type - use res.PolicyAction.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `action`                                                                          | `string`                                                                          | :heavy_check_mark:                                                                | The `slug_perm` (identifier) of the Action.                                       |
| `policy`                                                                          | `string`                                                                          | :heavy_check_mark:                                                                | The `slug_perm` (identifier) of the Policy.                                       |
| `workspace`                                                                       | `string`                                                                          | :heavy_check_mark:                                                                | The `name` of the Workspace.                                                      |
| `body`                                                                            | [*components.PatchedPolicyAction](../../models/components/patchedpolicyaction.md) | :heavy_minus_sign:                                                                | N/A                                                                               |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.WorkspacesPoliciesActionsPartialUpdateResponse](../../models/operations/workspacespoliciesactionspartialupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeletePolicyAction

Delete an action for a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_actions_destroy" method="delete" path="/workspaces/{workspace}/policies/{policy}/actions/{action}/" -->
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

    res, err := s.Workspaces.DeletePolicyAction(ctx, "<value>", "<value>", "<value>")
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
| `action`                                                 | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` (identifier) of the Action.              |
| `policy`                                                 | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` (identifier) of the Policy.              |
| `workspace`                                              | `string`                                                 | :heavy_check_mark:                                       | The `name` of the Workspace.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkspacesPoliciesActionsDestroyResponse](../../models/operations/workspacespoliciesactionsdestroyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SimulatePolicy

Simulates the execution of a Policy to verify its behavior without taking any actions.

Before using this endpoint, the Policy to simulate must be created with the creation endpoint. To experiment
with the policy, it is recommended to create the policy with disabled=True, so that it does not take effect
outside the simulation endpoint.

This endpoint evaluates all packages in the workspace, generating a "decision log" for each evaluation.
Each log contains:
- Package metadata provided to the policy engine at runtime.
- Output from the user-defined Rego policy.

No actions associated with the policy are executed. Instead, the endpoint reports what would happen to each package if the policy were active.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_simulate_list" method="get" path="/workspaces/{workspace}/policies/{policy}/simulate/" -->
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

    res, err := s.Workspaces.SimulatePolicy(ctx, operations.WorkspacesPoliciesSimulateListRequest{
        Page: 439618,
        Policy: "<value>",
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyDecisionLogList != nil {
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
| `request`                                                                                                            | [operations.WorkspacesPoliciesSimulateListRequest](../../models/operations/workspacespoliciessimulatelistrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.WorkspacesPoliciesSimulateListResponse](../../models/operations/workspacespoliciessimulatelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetPoliciesDecisionLogsV1

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

    res, err := s.Workspaces.GetPoliciesDecisionLogsV1(ctx, operations.WorkspacesPoliciesDecisionLogsV1ListRequest{
        CreatedAfter: types.MustTimeFromString("2024-11-12T13:20:12.290Z"),
        Page: 583225,
        Sort: cloudsmith.Pointer("id"),
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyDecisionLogV1SummaryList != nil {
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

    res, err := s.Workspaces.GetPoliciesDecisionLogsV1(ctx, operations.WorkspacesPoliciesDecisionLogsV1ListRequest{
        CreatedAfter: types.MustTimeFromString("2025-02-17T14:08:57.543Z"),
        Page: 377363,
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyDecisionLogV1SummaryList != nil {
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

## GetDecisionLog

Retrieve the full decision log JSON from S3.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_decision_logs_v1_retrieve" method="get" path="/workspaces/{workspace}/policies/decision-logs-v1/{decision_log_id}/" -->
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

    res, err := s.Workspaces.GetDecisionLog(ctx, "<id>", "<value>", nil)
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

## GetDecisionLogFacets

Internal endpoint that returns counts of policy decision logs grouped by a
single dimension. Used to populate the Policy Decision Logs filter dropdowns in
the web app with the full set of entities that had activity in the timeframe.

Excluded from the public OpenAPI spec in prd; visible in dev/stg.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_decision_logs_v1_facets_list" method="get" path="/workspaces/{workspace}/policies/decision-logs-v1/facets/" -->
```go
package main

import(
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/types"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"log"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

func main() {
    ctx := context.Background()

    s := cloudsmith.New(
        cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Workspaces.GetDecisionLogFacets(ctx, operations.WorkspacesPoliciesDecisionLogsV1FacetsListRequest{
        CreatedAfter: types.MustTimeFromString("2026-07-12T18:06:56.305Z"),
        GroupBy: operations.GroupByRepository,
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyDecisionLogV1FacetsResponse != nil {
        switch res.PolicyDecisionLogV1FacetsResponse.Type {
            case components.PolicyDecisionLogV1FacetsResponseTypePolicyDecisionLogV1FacetsPackageResponse:
                // res.PolicyDecisionLogV1FacetsResponse.PolicyDecisionLogV1FacetsPackageResponse is populated
            case components.PolicyDecisionLogV1FacetsResponseTypePolicyDecisionLogV1FacetsRepositoryResponse:
                // res.PolicyDecisionLogV1FacetsResponse.PolicyDecisionLogV1FacetsRepositoryResponse is populated
            case components.PolicyDecisionLogV1FacetsResponseTypePolicyDecisionLogV1FacetsPackageFormatResponse:
                // res.PolicyDecisionLogV1FacetsResponse.PolicyDecisionLogV1FacetsPackageFormatResponse is populated
            case components.PolicyDecisionLogV1FacetsResponseTypePolicyDecisionLogV1FacetsPolicyResponse:
                // res.PolicyDecisionLogV1FacetsResponse.PolicyDecisionLogV1FacetsPolicyResponse is populated
            default:
                // Unknown type - use res.PolicyDecisionLogV1FacetsResponse.GetUnknownRaw() for raw JSON
        }

    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.WorkspacesPoliciesDecisionLogsV1FacetsListRequest](../../models/operations/workspacespoliciesdecisionlogsv1facetslistrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.WorkspacesPoliciesDecisionLogsV1FacetsListResponse](../../models/operations/workspacespoliciesdecisionlogsv1facetslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |