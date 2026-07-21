# Vulnerabilities

## Overview

### Available Operations

* [Retrieve](#retrieve) - API to provide OSV.dev vulnerability data.

## Retrieve

API to provide OSV.dev vulnerability data.

### Example Usage

<!-- UsageSnippet language="go" operationID="vulnerabilities_osv_retrieve" method="get" path="/vulnerabilities/osv/{vulnerability_id}/" -->
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

    res, err := s.Vulnerabilities.Retrieve(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PolicyOSVVulnerability != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `vulnerabilityID`                                        | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VulnerabilitiesOsvRetrieveResponse](../../models/operations/vulnerabilitiesosvretrieveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |