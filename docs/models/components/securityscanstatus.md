# SecurityScanStatus

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

value := components.SecurityScanStatusAwaitingScan

// Open enum: custom values can be created with a direct type cast
custom := components.SecurityScanStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `SecurityScanStatusAwaitingScan`       | AWAITING_SCAN                          |
| `SecurityScanStatusScanning`           | SCANNING                               |
| `SecurityScanStatusVulnerable`         | VULNERABLE                             |
| `SecurityScanStatusNotVulnerable`      | NOT_VULNERABLE                         |
| `SecurityScanStatusDisabled`           | DISABLED                               |
| `SecurityScanStatusFailed`             | FAILED                                 |
| `SecurityScanStatusSkipped`            | SKIPPED                                |
| `SecurityScanStatusUnsupportedPackage` | UNSUPPORTED_PACKAGE                    |
| `SecurityScanStatusUnsupportedFormat`  | UNSUPPORTED_FORMAT                     |