# Experimental.Packages

## Overview

### Available Operations

* [UploadMaven](#uploadmaven) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new Maven package
* [UploadRpm](#uploadrpm) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new RedHat package
* [ValidateUploadConda](#validateuploadconda) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new Conda package payload
* [ValidateUploadCran](#validateuploadcran) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new CRAN package payload
* [ValidateUploadDeb](#validateuploaddeb) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new Debian package payload
* [ValidateUploadDocker](#validateuploaddocker) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new Docker package payload
* [ValidateUploadNpm](#validateuploadnpm) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new npm package payload
* [ValidateUploadNuget](#validateuploadnuget) - **Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new NuGet package payload

## UploadMaven

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new Maven package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_maven_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/maven/" -->
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

    res, err := s.Experimental.Packages.UploadMaven(ctx, "<value>", "<value>", components.MavenPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `repository`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | The `slug` of the repository within the workspace.                                 |
| `workspace`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | The `slug` of the workspace.                                                       |
| `body`                                                                             | [components.MavenPackageUploadV2](../../models/components/mavenpackageuploadv2.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadMavenCreateResponse](../../models/operations/workspacesrepositoriespackagesuploadmavencreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## UploadRpm

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Creates a new RedHat package

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_upload_rpm_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/upload/rpm/" -->
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

    res, err := s.Experimental.Packages.UploadRpm(ctx, "<value>", "<value>", components.RpmPackageUploadV2{
        Distribution: "<value>",
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PackageV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `repository`                                                                   | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the repository within the workspace.                             |
| `workspace`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the workspace.                                                   |
| `body`                                                                         | [components.RpmPackageUploadV2](../../models/components/rpmpackageuploadv2.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.WorkspacesRepositoriesPackagesUploadRpmCreateResponse](../../models/operations/workspacesrepositoriespackagesuploadrpmcreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateUploadConda

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new Conda package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_conda_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/conda/" -->
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

    res, err := s.Experimental.Packages.ValidateUploadConda(ctx, "<value>", "<value>", components.CondaPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `repository`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | The `slug` of the repository within the workspace.                                 |
| `workspace`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | The `slug` of the workspace.                                                       |
| `body`                                                                             | [components.CondaPackageUploadV2](../../models/components/condapackageuploadv2.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadCondaCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploadcondacreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateUploadCran

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new CRAN package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_cran_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/cran/" -->
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

    res, err := s.Experimental.Packages.ValidateUploadCran(ctx, "<value>", "<value>", components.CranPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `repository`                                                                     | `string`                                                                         | :heavy_check_mark:                                                               | The `slug` of the repository within the workspace.                               |
| `workspace`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | The `slug` of the workspace.                                                     |
| `body`                                                                           | [components.CranPackageUploadV2](../../models/components/cranpackageuploadv2.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadCranCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploadcrancreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateUploadDeb

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new Debian package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_deb_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/deb/" -->
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

    res, err := s.Experimental.Packages.ValidateUploadDeb(ctx, "<value>", "<value>", components.DebPackageUploadV2{
        Distribution: "<value>",
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `repository`                                                                   | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the repository within the workspace.                             |
| `workspace`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the workspace.                                                   |
| `body`                                                                         | [components.DebPackageUploadV2](../../models/components/debpackageuploadv2.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadDebCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploaddebcreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateUploadDocker

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new Docker package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_docker_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/docker/" -->
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

    res, err := s.Experimental.Packages.ValidateUploadDocker(ctx, "<value>", "<value>", components.DockerPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `repository`                                                                         | `string`                                                                             | :heavy_check_mark:                                                                   | The `slug` of the repository within the workspace.                                   |
| `workspace`                                                                          | `string`                                                                             | :heavy_check_mark:                                                                   | The `slug` of the workspace.                                                         |
| `body`                                                                               | [components.DockerPackageUploadV2](../../models/components/dockerpackageuploadv2.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadDockerCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploaddockercreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateUploadNpm

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new npm package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_npm_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/npm/" -->
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

    res, err := s.Experimental.Packages.ValidateUploadNpm(ctx, "<value>", "<value>", components.NpmPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `repository`                                                                   | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the repository within the workspace.                             |
| `workspace`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | The `slug` of the workspace.                                                   |
| `body`                                                                         | [components.NpmPackageUploadV2](../../models/components/npmpackageuploadv2.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadNpmCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploadnpmcreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## ValidateUploadNuget

**Experimental**: this endpoint is under active development and may change in backwards-incompatible ways without notice.

Validates a new NuGet package payload

### Example Usage

<!-- UsageSnippet language="go" operationID="workspaces_repositories_packages_validate_upload_nuget_create" method="post" path="/workspaces/{workspace}/repositories/{repository}/packages/validate-upload/nuget/" -->
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

    res, err := s.Experimental.Packages.ValidateUploadNuget(ctx, "<value>", "<value>", components.NugetPackageUploadV2{
        PackageFile: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `repository`                                                                       | `string`                                                                           | :heavy_check_mark:                                                                 | The `slug` of the repository within the workspace.                                 |
| `workspace`                                                                        | `string`                                                                           | :heavy_check_mark:                                                                 | The `slug` of the workspace.                                                       |
| `body`                                                                             | [components.NugetPackageUploadV2](../../models/components/nugetpackageuploadv2.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.WorkspacesRepositoriesPackagesValidateUploadNugetCreateResponse](../../models/operations/workspacesrepositoriespackagesvalidateuploadnugetcreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 400, 404, 422         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |