# ScanSummaryEnum

* `UNSAFE` - Unsafe
* `UNKNOWN` - Unknown
* `QUEUED` - Queued
* `SAFE` - Safe

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.ScanSummaryEnumUnsafe

// Open enum: custom values can be created with a direct type cast
custom := components.ScanSummaryEnum("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ScanSummaryEnumUnsafe`  | UNSAFE                   |
| `ScanSummaryEnumUnknown` | UNKNOWN                  |
| `ScanSummaryEnumQueued`  | QUEUED                   |
| `ScanSummaryEnumSafe`    | SAFE                     |