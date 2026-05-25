# Packages

## Overview

### Available Operations

* [PackagesMetadataList](#packagesmetadatalist) - List metadata for a package
* [PackagesMetadataCreate](#packagesmetadatacreate) - Attach metadata to a package
* [PackagesMetadataRetrieve](#packagesmetadataretrieve) - Retrieve a package metadata entry
* [PackagesMetadataPartialUpdate](#packagesmetadatapartialupdate) - Update a package metadata entry
* [PackagesMetadataDestroy](#packagesmetadatadestroy) - Delete a package metadata entry

## PackagesMetadataList

List all metadata entries for a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_metadata_list" method="get" path="/packages/{package_slug_perm}/metadata/" -->
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

    res, err := s.Packages.PackagesMetadataList(ctx, operations.PackagesMetadataListRequest{
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

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PackagesMetadataListRequest](../../models/operations/packagesmetadatalistrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PackagesMetadataListResponse](../../models/operations/packagesmetadatalistresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 403, 404, 422    | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## PackagesMetadataCreate

Attach arbitrary metadata to a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_metadata_create" method="post" path="/packages/{package_slug_perm}/metadata/" example="Unauthenticated" -->
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

    res, err := s.Packages.PackagesMetadataCreate(ctx, "<value>", components.ArtifactMetadataWriteInput{
        Content: "<value>",
        ContentType: "<value>",
        SourceIdentity: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ArtifactMetadataWrite != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `packageSlugPerm`                                                                              | `string`                                                                                       | :heavy_check_mark:                                                                             | The `slug_perm` of the package.                                                                |
| `body`                                                                                         | [components.ArtifactMetadataWriteInput](../../models/components/artifactmetadatawriteinput.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PackagesMetadataCreateResponse](../../models/operations/packagesmetadatacreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 403, 404, 422    | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## PackagesMetadataRetrieve

Retrieve a specific metadata entry for a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_metadata_retrieve" method="get" path="/packages/{package_slug_perm}/metadata/{slug_perm}/" -->
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

    res, err := s.Packages.PackagesMetadataRetrieve(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ArtifactMetadataRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `packageSlugPerm`                                        | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` of the package.                          |
| `slugPerm`                                               | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` of the metadata entry.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PackagesMetadataRetrieveResponse](../../models/operations/packagesmetadataretrieveresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 403, 404         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## PackagesMetadataPartialUpdate

Edit a specific metadata entry for a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_metadata_partial_update" method="patch" path="/packages/{package_slug_perm}/metadata/{slug_perm}/" -->
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

    res, err := s.Packages.PackagesMetadataPartialUpdate(ctx, "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ArtifactMetadataWrite != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |
| `packageSlugPerm`                                                                                   | `string`                                                                                            | :heavy_check_mark:                                                                                  | The `slug_perm` of the package.                                                                     |
| `slugPerm`                                                                                          | `string`                                                                                            | :heavy_check_mark:                                                                                  | The `slug_perm` of the metadata entry.                                                              |
| `body`                                                                                              | [*components.PatchedArtifactMetadataWrite](../../models/components/patchedartifactmetadatawrite.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |

### Response

**[*operations.PackagesMetadataPartialUpdateResponse](../../models/operations/packagesmetadatapartialupdateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 403, 404, 422    | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## PackagesMetadataDestroy

Remove a specific metadata entry from a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="packages_metadata_destroy" method="delete" path="/packages/{package_slug_perm}/metadata/{slug_perm}/" -->
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

    res, err := s.Packages.PackagesMetadataDestroy(ctx, "<value>", "<value>")
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
| `packageSlugPerm`                                        | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` of the package.                          |
| `slugPerm`                                               | `string`                                                 | :heavy_check_mark:                                       | The `slug_perm` of the metadata entry.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PackagesMetadataDestroyResponse](../../models/operations/packagesmetadatadestroyresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 403, 404         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |