# cloudsmith

Developer-friendly & type-safe Go SDK specifically catered to leverage *cloudsmith* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=cloudsmith&utm_campaign=go)
[![License: Apache 2.0](https://img.shields.io/badge/LICENSE_//_Apache_2.0-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://www.apache.org/licenses/LICENSE-2.0)


<!-- Start Summary [summary] -->
## Summary

Cloudsmith API: The API to the Cloudsmith Service

For more information about the API: [Cloudsmith API Reference](https://help.cloudsmith.io/reference/introduction)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [cloudsmith](#cloudsmith)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/cloudsmith-io/cloudsmith-go-v2
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
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

	res, err := s.Metadata.MetadataPackagesList(ctx, operations.MetadataPackagesListRequest{
		PackageSlugPerm: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PaginatedArtifactMetadataReadList != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  |
| -------- | ------ | ------- |
| `Apikey` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
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

	res, err := s.Metadata.MetadataPackagesList(ctx, operations.MetadataPackagesListRequest{
		PackageSlugPerm: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PaginatedArtifactMetadataReadList != nil {
		// handle response
	}
}

```

### Environment Variables

When no explicit option is provided to `New()`, the SDK falls back to the following environment variables:

| Variable               | Purpose                                                                                              |
| ---------------------- | ---------------------------------------------------------------------------------------------------- |
| `CLOUDSMITH_API_KEY`   | API key used for authentication when `WithSecurity(...)` is not passed.                              |
| `CLOUDSMITH_API_HOST`  | Base API host used when `WithServerURL(...)` / `WithServerIndex(...)` is not passed. The value is normalized and `/v2/` is appended if not already present. |

Explicit options always take precedence over environment variables; the variables are only consulted as a fallback. With both variables set, the client can be constructed with no options:

```go
package main

import (
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	// Reads CLOUDSMITH_API_KEY and CLOUDSMITH_API_HOST from the environment.
	s := cloudsmith.New()

	res, err := s.Metadata.MetadataPackagesList(ctx, operations.MetadataPackagesListRequest{
		PackageSlugPerm: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PaginatedArtifactMetadataReadList != nil {
		// handle response
	}
}
```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Metadata](docs/sdks/metadata/README.md)

* [MetadataPackagesList](docs/sdks/metadata/README.md#metadatapackageslist) - List metadata for a package
* [MetadataPackagesCreate](docs/sdks/metadata/README.md#metadatapackagescreate) - Attach metadata to a package
* [MetadataPackagesRetrieve](docs/sdks/metadata/README.md#metadatapackagesretrieve) - Retrieve a package metadata entry
* [MetadataPackagesPartialUpdate](docs/sdks/metadata/README.md#metadatapackagespartialupdate) - Update a package metadata entry
* [MetadataPackagesDestroy](docs/sdks/metadata/README.md#metadatapackagesdestroy) - Delete a package metadata entry
* [MetadataValidateCreate](docs/sdks/metadata/README.md#metadatavalidatecreate) - Validate a metadata payload

### [Workspaces](docs/sdks/workspaces/README.md)

* [WorkspacesPoliciesList](docs/sdks/workspaces/README.md#workspacespolicieslist) - List policies.
* [WorkspacesPoliciesCreate](docs/sdks/workspaces/README.md#workspacespoliciescreate) - Create a policy.
* [WorkspacesPoliciesRetrieve](docs/sdks/workspaces/README.md#workspacespoliciesretrieve) - Retrieve a policy.
* [WorkspacesPoliciesUpdate](docs/sdks/workspaces/README.md#workspacespoliciesupdate) - Update a policy.
* [WorkspacesPoliciesPartialUpdate](docs/sdks/workspaces/README.md#workspacespoliciespartialupdate) - Partially update a policy.
* [WorkspacesPoliciesDestroy](docs/sdks/workspaces/README.md#workspacespoliciesdestroy) - Delete a policy.
* [WorkspacesPoliciesActionsList](docs/sdks/workspaces/README.md#workspacespoliciesactionslist) - List all actions for the policy.
* [WorkspacesPoliciesActionsCreate](docs/sdks/workspaces/README.md#workspacespoliciesactionscreate) - Create an action for a policy.
* [WorkspacesPoliciesActionsRetrieve](docs/sdks/workspaces/README.md#workspacespoliciesactionsretrieve) - Retrieve an action for a policy.
* [WorkspacesPoliciesActionsUpdate](docs/sdks/workspaces/README.md#workspacespoliciesactionsupdate) - Update an action for a policy.
* [WorkspacesPoliciesActionsPartialUpdate](docs/sdks/workspaces/README.md#workspacespoliciesactionspartialupdate) - Partially update an action for a policy.
* [WorkspacesPoliciesActionsDestroy](docs/sdks/workspaces/README.md#workspacespoliciesactionsdestroy) - Delete an action for a policy.
* [WorkspacesPoliciesSimulateList](docs/sdks/workspaces/README.md#workspacespoliciessimulatelist) - Simulates the execution of a Policy to verify its behavior without taking any actions.

Before using this endpoint, the Policy to simulate must be created with the creation endpoint. To experiment
with the policy, it is recommended to create the policy with disabled=True, so that it does not take effect
outside the simulation endpoint.

This endpoint evaluates all packages in the workspace, generating a "decision log" for each evaluation.
Each log contains:
- Package metadata provided to the policy engine at runtime.
- Output from the user-defined Rego policy.

No actions associated with the policy are executed. Instead, the endpoint reports what would happen to each package if the policy were active.
* [WorkspacesPoliciesDecisionLogsV1List](docs/sdks/workspaces/README.md#workspacespoliciesdecisionlogsv1list) - List policy evaluation decision logs.
* [WorkspacesPoliciesDecisionLogsV1Retrieve](docs/sdks/workspaces/README.md#workspacespoliciesdecisionlogsv1retrieve) - Retrieve the full decision log JSON from S3.

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
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
		Page:      694333,
		Query:     cloudsmith.Pointer("name: foo OR name: bar"),
		Sort:      cloudsmith.Pointer("-version,updated_at"),
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
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"github.com/cloudsmith-io/cloudsmith-go-v2/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := cloudsmith.New(
		cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Metadata.MetadataPackagesList(ctx, operations.MetadataPackagesListRequest{
		PackageSlugPerm: "<value>",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.PaginatedArtifactMetadataReadList != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"github.com/cloudsmith-io/cloudsmith-go-v2/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := cloudsmith.New(
		cloudsmith.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Metadata.MetadataPackagesList(ctx, operations.MetadataPackagesListRequest{
		PackageSlugPerm: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PaginatedArtifactMetadataReadList != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `MetadataPackagesList` function may return the following errors:

| Error Type            | Status Code        | Content Type     |
| --------------------- | ------------------ | ---------------- |
| apierrors.ErrorDetail | 401, 403, 404, 422 | application/json |
| apierrors.APIError    | 4XX, 5XX           | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/apierrors"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := cloudsmith.New(
		cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Metadata.MetadataPackagesList(ctx, operations.MetadataPackagesListRequest{
		PackageSlugPerm: "<value>",
	})
	if err != nil {

		var e *apierrors.ErrorDetail
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	cloudsmith "github.com/cloudsmith-io/cloudsmith-go-v2"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := cloudsmith.New(
		cloudsmith.WithServerURL("https://api.cloudsmith.io/v2/"),
		cloudsmith.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Metadata.MetadataPackagesList(ctx, operations.MetadataPackagesListRequest{
		PackageSlugPerm: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PaginatedArtifactMetadataReadList != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/cloudsmith-io/cloudsmith-go-v2"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = cloudsmith.New(cloudsmith.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is generally available (GA) and ready for production use. We follow [semantic versioning](https://semver.org/), so you can expect backwards-compatible changes within a major version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=cloudsmith&utm_campaign=go)
