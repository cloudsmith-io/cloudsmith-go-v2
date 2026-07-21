# Workspaces.Actions

## Overview

### Available Operations

* [Create](#create) - Create an action for a policy.
* [PartialUpdate](#partialupdate) - Partially update an action for a policy.

## Create

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

    res, err := s.Workspaces.Actions.Create(ctx, "<value>", "<value>", nil)
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

## PartialUpdate

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

    res, err := s.Workspaces.Actions.PartialUpdate(ctx, "<value>", "<value>", "<value>", nil)
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