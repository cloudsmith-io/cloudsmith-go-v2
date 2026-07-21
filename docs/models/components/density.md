# Density

Row density for the table. Null means use the default.

* `COMFORTABLE` - Comfortable
* `COMPACT` - Compact

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.DensityComfortable

// Open enum: custom values can be created with a direct type cast
custom := components.Density("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `DensityComfortable` | COMFORTABLE          |
| `DensityCompact`     | COMPACT              |