# Metadata

## Overview

### Available Operations

* [MetadataPackagesList](#metadatapackageslist) - List metadata for a package
* [MetadataPackagesCreate](#metadatapackagescreate) - Attach metadata to a package
* [MetadataPackagesRetrieve](#metadatapackagesretrieve) - Retrieve a package metadata entry
* [MetadataPackagesPartialUpdate](#metadatapackagespartialupdate) - Update a package metadata entry
* [MetadataPackagesDestroy](#metadatapackagesdestroy) - Delete a package metadata entry
* [MetadataValidateCreate](#metadatavalidatecreate) - Validate a metadata payload

## MetadataPackagesList

List all metadata entries for a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="metadata_packages_list" method="get" path="/metadata/packages/{package_slug_perm}/" -->
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

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.MetadataPackagesListRequest](../../models/operations/metadatapackageslistrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.MetadataPackagesListResponse](../../models/operations/metadatapackageslistresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 403, 404, 422    | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## MetadataPackagesCreate

Attach arbitrary metadata to a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="metadata_packages_create" method="post" path="/metadata/packages/{package_slug_perm}/" example="Unauthenticated" -->
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

    res, err := s.Metadata.MetadataPackagesCreate(ctx, "<value>", components.ArtifactMetadataWriteInput{
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

**[*operations.MetadataPackagesCreateResponse](../../models/operations/metadatapackagescreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorDetail   | 401, 402, 403, 404, 422 | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MetadataPackagesRetrieve

Retrieve a specific metadata entry for a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="metadata_packages_retrieve" method="get" path="/metadata/packages/{package_slug_perm}/{slug_perm}/" -->
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

    res, err := s.Metadata.MetadataPackagesRetrieve(ctx, "<value>", "<value>")
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

**[*operations.MetadataPackagesRetrieveResponse](../../models/operations/metadatapackagesretrieveresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 403, 404         | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## MetadataPackagesPartialUpdate

Edit a specific metadata entry for a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="metadata_packages_partial_update" method="patch" path="/metadata/packages/{package_slug_perm}/{slug_perm}/" -->
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

    res, err := s.Metadata.MetadataPackagesPartialUpdate(ctx, "<value>", "<value>", nil)
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

**[*operations.MetadataPackagesPartialUpdateResponse](../../models/operations/metadatapackagespartialupdateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorDetail   | 401, 402, 403, 404, 422 | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MetadataPackagesDestroy

Remove a specific metadata entry from a package.

### Example Usage

<!-- UsageSnippet language="go" operationID="metadata_packages_destroy" method="delete" path="/metadata/packages/{package_slug_perm}/{slug_perm}/" -->
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

    res, err := s.Metadata.MetadataPackagesDestroy(ctx, "<value>", "<value>")
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

**[*operations.MetadataPackagesDestroyResponse](../../models/operations/metadatapackagesdestroyresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 402, 403, 404    | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |

## MetadataValidateCreate

Validate a metadata payload against the JSON schema for its ``content_type`` without persisting it. Returns 200 on success and 422 on validation failure, with the same field-level error shape as the create endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="metadata_validate_create" method="post" path="/metadata/validate/" example="Unauthenticated" -->
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

    res, err := s.Metadata.MetadataValidateCreate(ctx, components.ArtifactMetadataValidate{
        Content: "<value>",
        ContentType: "<value>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.ArtifactMetadataValidate](../../models/components/artifactmetadatavalidate.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.MetadataValidateCreateResponse](../../models/operations/metadatavalidatecreateresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| apierrors.ErrorDetail | 401, 422              | application/json      |
| apierrors.APIError    | 4XX, 5XX              | \*/\*                 |