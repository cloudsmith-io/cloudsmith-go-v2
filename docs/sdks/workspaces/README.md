# Workspaces

## Overview

### Available Operations

* [WorkspacesPoliciesList](#workspacespolicieslist) - List policies.
* [WorkspacesPoliciesCreate](#workspacespoliciescreate) - Create a policy.
* [WorkspacesPoliciesRetrieve](#workspacespoliciesretrieve) - Retrieve a policy.
* [WorkspacesPoliciesUpdate](#workspacespoliciesupdate) - Update a policy.
* [WorkspacesPoliciesPartialUpdate](#workspacespoliciespartialupdate) - Partially update a policy.
* [WorkspacesPoliciesDestroy](#workspacespoliciesdestroy) - Delete a policy.
* [WorkspacesPoliciesActionsList](#workspacespoliciesactionslist) - List all actions for the policy.
* [WorkspacesPoliciesActionsCreate](#workspacespoliciesactionscreate) - Create an action for a policy.
* [WorkspacesPoliciesActionsRetrieve](#workspacespoliciesactionsretrieve) - Retrieve an action for a policy.
* [WorkspacesPoliciesActionsUpdate](#workspacespoliciesactionsupdate) - Update an action for a policy.
* [WorkspacesPoliciesActionsPartialUpdate](#workspacespoliciesactionspartialupdate) - Partially update an action for a policy.
* [WorkspacesPoliciesActionsDestroy](#workspacespoliciesactionsdestroy) - Delete an action for a policy.
* [WorkspacesPoliciesSimulateList](#workspacespoliciessimulatelist) - Simulates the execution of a Policy to verify its behavior without taking any actions.

Before using this endpoint, the Policy to simulate must be created with the creation endpoint. To experiment
with the policy, it is recommended to create the policy with disabled=True, so that it does not take effect
outside the simulation endpoint.

This endpoint evaluates all packages in the workspace, generating a "decision log" for each evaluation.
Each log contains:
- Package metadata provided to the policy engine at runtime.
- Output from the user-defined Rego policy.

No actions associated with the policy are executed. Instead, the endpoint reports what would happen to each package if the policy were active.
* [WorkspacesPoliciesDecisionLogsV1List](#workspacespoliciesdecisionlogsv1list) - List policy evaluation decision logs.
* [WorkspacesPoliciesDecisionLogsV1Retrieve](#workspacespoliciesdecisionlogsv1retrieve) - Retrieve the full decision log JSON from S3.

## WorkspacesPoliciesList

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

    res, err := s.Workspaces.WorkspacesPoliciesList(ctx, operations.WorkspacesPoliciesListRequest{
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

    res, err := s.Workspaces.WorkspacesPoliciesList(ctx, operations.WorkspacesPoliciesListRequest{
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

## WorkspacesPoliciesCreate

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

    res, err := s.Workspaces.WorkspacesPoliciesCreate(ctx, "<value>", components.PolicyRequest{
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

## WorkspacesPoliciesRetrieve

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

    res, err := s.Workspaces.WorkspacesPoliciesRetrieve(ctx, "<value>", "<value>")
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

## WorkspacesPoliciesUpdate

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

    res, err := s.Workspaces.WorkspacesPoliciesUpdate(ctx, "<value>", "<value>", components.PolicyRequest{
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

## WorkspacesPoliciesPartialUpdate

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

    res, err := s.Workspaces.WorkspacesPoliciesPartialUpdate(ctx, "<value>", "<value>", nil)
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

## WorkspacesPoliciesDestroy

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

    res, err := s.Workspaces.WorkspacesPoliciesDestroy(ctx, "<value>", "<value>")
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

## WorkspacesPoliciesActionsList

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

    res, err := s.Workspaces.WorkspacesPoliciesActionsList(ctx, "<value>", "<value>", nil, nil)
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

## WorkspacesPoliciesActionsCreate

Create an action for a policy.

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_policies_actions_create" method="post" path="/workspaces/{workspace}/policies/{policy}/actions/" -->
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

    res, err := s.Workspaces.WorkspacesPoliciesActionsCreate(ctx, "<value>", "<value>", nil)
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
| `policy`                                                                          | `string`                                                                          | :heavy_check_mark:                                                                | The `slug_perm` (identifier) of the Policy.                                       |
| `workspace`                                                                       | `string`                                                                          | :heavy_check_mark:                                                                | The `name` of the Workspace.                                                      |
| `body`                                                                            | [*components.PolicyActionRequest](../../models/components/policyactionrequest.md) | :heavy_minus_sign:                                                                | N/A                                                                               |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.WorkspacesPoliciesActionsCreateResponse](../../models/operations/workspacespoliciesactionscreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 422                   | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## WorkspacesPoliciesActionsRetrieve

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

    res, err := s.Workspaces.WorkspacesPoliciesActionsRetrieve(ctx, "<value>", "<value>", "<value>")
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

## WorkspacesPoliciesActionsUpdate

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

    res, err := s.Workspaces.WorkspacesPoliciesActionsUpdate(ctx, "<value>", "<value>", "<value>", nil)
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

## WorkspacesPoliciesActionsPartialUpdate

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

    res, err := s.Workspaces.WorkspacesPoliciesActionsPartialUpdate(ctx, "<value>", "<value>", "<value>", nil)
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

## WorkspacesPoliciesActionsDestroy

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

    res, err := s.Workspaces.WorkspacesPoliciesActionsDestroy(ctx, "<value>", "<value>", "<value>")
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

## WorkspacesPoliciesSimulateList

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

    res, err := s.Workspaces.WorkspacesPoliciesSimulateList(ctx, operations.WorkspacesPoliciesSimulateListRequest{
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

## WorkspacesPoliciesDecisionLogsV1List

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

    res, err := s.Workspaces.WorkspacesPoliciesDecisionLogsV1List(ctx, operations.WorkspacesPoliciesDecisionLogsV1ListRequest{
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

    res, err := s.Workspaces.WorkspacesPoliciesDecisionLogsV1List(ctx, operations.WorkspacesPoliciesDecisionLogsV1ListRequest{
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

## WorkspacesPoliciesDecisionLogsV1Retrieve

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

    res, err := s.Workspaces.WorkspacesPoliciesDecisionLogsV1Retrieve(ctx, "<id>", "<value>", nil)
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