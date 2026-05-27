<!-- Start SDK Example Usage [usage] -->
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