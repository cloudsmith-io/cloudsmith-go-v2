# Workspaces

## Overview

### Available Operations

* [ListPolicies](#listpolicies) - List policies.
* [CreatePolicy](#createpolicy) - Create a policy.
* [GetPolicy](#getpolicy) - Retrieve a policy.
* [ReplacePolicy](#replacepolicy) - Update a policy.
* [UpdatePolicy](#updatepolicy) - Partially update a policy.
* [DeletePolicy](#deletepolicy) - Delete a policy.
* [GetPolicyActions](#getpolicyactions) - List all actions for the policy.
* [GetPolicyAction](#getpolicyaction) - Retrieve an action for a policy.
* [UpdatePolicyAction](#updatepolicyaction) - Update an action for a policy.
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
* [GetPolicyDecisionLogFacets](#getpolicydecisionlogfacets) - Internal endpoint that returns counts of policy decision logs grouped by a
single dimension. Used to populate the Policy Decision Logs filter dropdowns in
the web app with the full set of entities that had activity in the timeframe.

Excluded from the public OpenAPI spec in prd; visible in dev/stg.

## ListPolicies

List policies.

### Example Usage: PoliciesNamedFooOrBar

<!-- UsageSnippet language="go" operationID="workspaces_policies_list" method="get" path="/workspaces/{workspace}/policies/" example="PoliciesNamedFooOrBar" -->
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

    res, err := s.Workspaces.ListPolicies(ctx, operations.WorkspacesPoliciesListRequest{
        Page: 694333,
        Query: cloudsmith.Pointer("name: foo OR name: bar"),
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyList != nil {
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
### Example Usage: `version`Descending,`updatedAt`Ascending

<!-- UsageSnippet language="go" operationID="workspaces_policies_list" method="get" path="/workspaces/{workspace}/policies/" example="`version`Descending,`updatedAt`Ascending" -->
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

    res, err := s.Workspaces.ListPolicies(ctx, operations.WorkspacesPoliciesListRequest{
        Page: 475392,
        Sort: cloudsmith.Pointer("-version,updated_at"),
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyList != nil {
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.WorkspacesPoliciesListRequest](../../models/operations/workspacespolicieslistrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.WorkspacesPoliciesListResponse](../../models/operations/workspacespolicieslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

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

## GetPolicy

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

    res, err := s.Workspaces.GetPolicy(ctx, "<value>", "<value>")
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

## ReplacePolicy

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

    res, err := s.Workspaces.ReplacePolicy(ctx, "<value>", "<value>", components.PolicyRequest{
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

## UpdatePolicy

Partially update a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_partial_update" method="patch" path="/workspaces/{workspace}/policies/{policy}/" -->
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

    res, err := s.Workspaces.UpdatePolicy(ctx, "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Policy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `policy`                                                              | `string`                                                              | :heavy_check_mark:                                                    | The `slug_perm` (identifier) of the Policy.                           |
| `workspace`                                                           | `string`                                                              | :heavy_check_mark:                                                    | The `name` of the Workspace.                                          |
| `body`                                                                | [*components.PatchedPolicy](../../models/components/patchedpolicy.md) | :heavy_minus_sign:                                                    | N/A                                                                   |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.WorkspacesPoliciesPartialUpdateResponse](../../models/operations/workspacespoliciespartialupdateresponse.md), error**

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

    res, err := s.Workspaces.GetPolicyActions(ctx, "<value>", "<value>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyActionList != nil {
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
| `page`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | A page number within the paginated result set.           |
| `pageSize`                                               | `*int64`                                                 | :heavy_minus_sign:                                       | Number of results to return per page.                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkspacesPoliciesActionsListResponse](../../models/operations/workspacespoliciesactionslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetPolicyAction

Retrieve an action for a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_actions_retrieve" method="get" path="/workspaces/{workspace}/policies/{policy}/actions/{action}/" -->
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

    res, err := s.Workspaces.GetPolicyAction(ctx, "<value>", "<value>", "<value>")
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `action`                                                 | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` (identifier) of the Action.              |
| `policy`                                                 | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` (identifier) of the Policy.              |
| `workspace`                                              | `string`                                                 | :heavy_check_mark:                                       | The `name` of the Workspace.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WorkspacesPoliciesActionsRetrieveResponse](../../models/operations/workspacespoliciesactionsretrieveresponse.md), error**

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
        Policy: "<value>",
        Workspace: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedPolicyDecisionLogList != nil {
        // handle response
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

## GetPolicyDecisionLogFacets

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

    res, err := s.Workspaces.GetPolicyDecisionLogFacets(ctx, operations.WorkspacesPoliciesDecisionLogsV1FacetsListRequest{
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