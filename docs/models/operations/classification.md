# Classification

Filter results by metadata classification.

* `UNKNOWN` - Unknown
* `INTRINSIC` - Intrinsic
* `SECURITY` - Security
* `PROVENANCE` - Provenance
* `SBOM` - SBOM
* `GENERIC` - Generic

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
)

value := operations.ClassificationUnknown
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ClassificationUnknown`    | UNKNOWN                    |
| `ClassificationIntrinsic`  | INTRINSIC                  |
| `ClassificationSecurity`   | SECURITY                   |
| `ClassificationProvenance` | PROVENANCE                 |
| `ClassificationSbom`       | SBOM                       |
| `ClassificationGeneric`    | GENERIC                    |