# PackagesMetadataListClassification

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

value := operations.PackagesMetadataListClassificationUnknown
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `PackagesMetadataListClassificationUnknown`    | UNKNOWN                                        |
| `PackagesMetadataListClassificationIntrinsic`  | INTRINSIC                                      |
| `PackagesMetadataListClassificationSecurity`   | SECURITY                                       |
| `PackagesMetadataListClassificationProvenance` | PROVENANCE                                     |
| `PackagesMetadataListClassificationSbom`       | SBOM                                           |
| `PackagesMetadataListClassificationGeneric`    | GENERIC                                        |