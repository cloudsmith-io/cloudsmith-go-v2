# MetadataPackagesListClassification

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

value := operations.MetadataPackagesListClassificationUnknown
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `MetadataPackagesListClassificationUnknown`    | UNKNOWN                                        |
| `MetadataPackagesListClassificationIntrinsic`  | INTRINSIC                                      |
| `MetadataPackagesListClassificationSecurity`   | SECURITY                                       |
| `MetadataPackagesListClassificationProvenance` | PROVENANCE                                     |
| `MetadataPackagesListClassificationSbom`       | SBOM                                           |
| `MetadataPackagesListClassificationGeneric`    | GENERIC                                        |