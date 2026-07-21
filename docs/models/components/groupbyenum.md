# GroupByEnum

* `REPOSITORY` - Repository
* `USER` - User
* `ENTITLEMENT_TOKEN` - Entitlement Token
* `PACKAGE_FORMAT` - Package Format
* `IP_ADDRESS` - IP Address
* `PACKAGE` - Package
* `COUNTRY` - Country
* `HTTP_STATUS` - HTTP Status
* `EDGE_RESPONSE` - Edge Response
* `HOST` - Host

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.GroupByEnumRepository

// Open enum: custom values can be created with a direct type cast
custom := components.GroupByEnum("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `GroupByEnumRepository`       | REPOSITORY                    |
| `GroupByEnumUser`             | USER                          |
| `GroupByEnumEntitlementToken` | ENTITLEMENT_TOKEN             |
| `GroupByEnumPackageFormat`    | PACKAGE_FORMAT                |
| `GroupByEnumIPAddress`        | IP_ADDRESS                    |
| `GroupByEnumPackage`          | PACKAGE                       |
| `GroupByEnumCountry`          | COUNTRY                       |
| `GroupByEnumHTTPStatus`       | HTTP_STATUS                   |
| `GroupByEnumEdgeResponse`     | EDGE_RESPONSE                 |
| `GroupByEnumHost`             | HOST                          |