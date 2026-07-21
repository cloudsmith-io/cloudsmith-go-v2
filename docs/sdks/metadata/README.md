# Metadata

## Overview

### Available Operations

* [ListPackageMetadata](#listpackagemetadata) - List metadata for a package
* [AttachPackageMetadata](#attachpackagemetadata) - Attach metadata to a package
* [GetPackageMetadata](#getpackagemetadata) - Retrieve a package metadata entry
* [UpdatePackageMetadata](#updatepackagemetadata) - Update a package metadata entry
* [DeletePackageMetadata](#deletepackagemetadata) - Delete a package metadata entry
* [Validate](#validate) - Validate a metadata payload

## ListPackageMetadata

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

    res, err := s.Metadata.ListPackageMetadata(ctx, operations.MetadataPackagesListRequest{
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

## AttachPackageMetadata

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

    res, err := s.Metadata.AttachPackageMetadata(ctx, "<value>", components.ArtifactMetadataWriteRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `packageSlugPerm`                                                                                  | `string`                                                                                           | :heavy_check_mark:                                                                                 | The `slug_perm` of the package.                                                                    |
| `body`                                                                                             | [components.ArtifactMetadataWriteRequest](../../models/components/artifactmetadatawriterequest.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.MetadataPackagesCreateResponse](../../models/operations/metadatapackagescreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorDetail   | 401, 402, 403, 404, 422 | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetPackageMetadata

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

    res, err := s.Metadata.GetPackageMetadata(ctx, "<value>", "<value>")
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

## UpdatePackageMetadata

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

    res, err := s.Metadata.UpdatePackageMetadata(ctx, "<value>", "<value>", nil)
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

## DeletePackageMetadata

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

    res, err := s.Metadata.DeletePackageMetadata(ctx, "<value>", "<value>")
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

## Validate

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

    res, err := s.Metadata.Validate(ctx, components.ArtifactMetadataValidate{
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