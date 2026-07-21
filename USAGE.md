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

	res, err := s.Analytics.AnalyticsLogsClientList(ctx, operations.AnalyticsLogsClientListRequest{
		Page:      408040,
		Workspace: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PaginatedLogsResponseList != nil {
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
<!-- End SDK Example Usage [usage] -->