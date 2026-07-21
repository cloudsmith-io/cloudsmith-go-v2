# Workspaces.Policies

## Overview

### Available Operations

* [List](#list) - List policies.
* [Update](#update) - Partially update a policy.

## List

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

    res, err := s.Workspaces.Policies.List(ctx, operations.WorkspacesPoliciesListRequest{
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

    res, err := s.Workspaces.Policies.List(ctx, operations.WorkspacesPoliciesListRequest{
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

## Update

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

    res, err := s.Workspaces.Policies.Update(ctx, "<value>", "<value>", nil)
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