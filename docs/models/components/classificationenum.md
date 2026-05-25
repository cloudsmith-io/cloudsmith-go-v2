# ClassificationEnum

* `UNKNOWN` - Unknown
* `INTRINSIC` - Intrinsic
* `SECURITY` - Security
* `PROVENANCE` - Provenance
* `SBOM` - SBOM
* `GENERIC` - Generic

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.ClassificationEnumUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.ClassificationEnum("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ClassificationEnumUnknown`    | UNKNOWN                        |
| `ClassificationEnumIntrinsic`  | INTRINSIC                      |
| `ClassificationEnumSecurity`   | SECURITY                       |
| `ClassificationEnumProvenance` | PROVENANCE                     |
| `ClassificationEnumSbom`       | SBOM                           |
| `ClassificationEnumGeneric`    | GENERIC                        |