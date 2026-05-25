# SecurityScanStatusEnum

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

value := components.SecurityScanStatusEnumAwaitingScan

// Open enum: custom values can be created with a direct type cast
custom := components.SecurityScanStatusEnum("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `SecurityScanStatusEnumAwaitingScan`       | AWAITING_SCAN                              |
| `SecurityScanStatusEnumScanning`           | SCANNING                                   |
| `SecurityScanStatusEnumVulnerable`         | VULNERABLE                                 |
| `SecurityScanStatusEnumNotVulnerable`      | NOT_VULNERABLE                             |
| `SecurityScanStatusEnumDisabled`           | DISABLED                                   |
| `SecurityScanStatusEnumFailed`             | FAILED                                     |
| `SecurityScanStatusEnumSkipped`            | SKIPPED                                    |
| `SecurityScanStatusEnumUnsupportedPackage` | UNSUPPORTED_PACKAGE                        |
| `SecurityScanStatusEnumUnsupportedFormat`  | UNSUPPORTED_FORMAT                         |