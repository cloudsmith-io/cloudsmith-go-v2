# SeverityEnum

* `LOW` - LOW
* `MEDIUM` - MEDIUM
* `HIGH` - HIGH
* `CRITICAL` - CRITICAL

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.SeverityEnumLow

// Open enum: custom values can be created with a direct type cast
custom := components.SeverityEnum("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `SeverityEnumLow`      | LOW                    |
| `SeverityEnumMedium`   | MEDIUM                 |
| `SeverityEnumHigh`     | HIGH                   |
| `SeverityEnumCritical` | CRITICAL               |