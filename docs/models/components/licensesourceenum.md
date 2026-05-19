# LicenseSourceEnum

* `AUTOMATIC` - Automatically sourced
* `MANUAL` - Manually sourced

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.LicenseSourceEnumAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.LicenseSourceEnum("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `LicenseSourceEnumAutomatic` | AUTOMATIC                    |
| `LicenseSourceEnumManual`    | MANUAL                       |