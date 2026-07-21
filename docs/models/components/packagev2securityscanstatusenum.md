# PackageV2SecurityScanStatusEnum

* `AWAITING_SCAN` - Awaiting Security Scan
* `SCANNING` - Security Scanning in Progress
* `VULNERABLE` - Scan Detected Vulnerabilities
* `NOT_VULNERABLE` - Scan Detected No Vulnerabilities
* `DISABLED` - Security Scanning Disabled
* `FAILED` - Security Scanning Failed
* `SKIPPED` - Security Scanning Skipped
* `UNSUPPORTED_PACKAGE` - Security Scanning Not Supported
* `UNSUPPORTED_FORMAT` - Security Scanning Not Supported

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.PackageV2SecurityScanStatusEnumAwaitingScan

// Open enum: custom values can be created with a direct type cast
custom := components.PackageV2SecurityScanStatusEnum("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `PackageV2SecurityScanStatusEnumAwaitingScan`       | AWAITING_SCAN                                       |
| `PackageV2SecurityScanStatusEnumScanning`           | SCANNING                                            |
| `PackageV2SecurityScanStatusEnumVulnerable`         | VULNERABLE                                          |
| `PackageV2SecurityScanStatusEnumNotVulnerable`      | NOT_VULNERABLE                                      |
| `PackageV2SecurityScanStatusEnumDisabled`           | DISABLED                                            |
| `PackageV2SecurityScanStatusEnumFailed`             | FAILED                                              |
| `PackageV2SecurityScanStatusEnumSkipped`            | SKIPPED                                             |
| `PackageV2SecurityScanStatusEnumUnsupportedPackage` | UNSUPPORTED_PACKAGE                                 |
| `PackageV2SecurityScanStatusEnumUnsupportedFormat`  | UNSUPPORTED_FORMAT                                  |