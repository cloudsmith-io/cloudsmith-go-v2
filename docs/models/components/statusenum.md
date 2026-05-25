# StatusEnum

* `UNSAFE` - Unsafe
* `UNKNOWN` - Unknown
* `QUEUED` - Queued
* `SAFE` - Safe

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.StatusEnumUnsafe

// Open enum: custom values can be created with a direct type cast
custom := components.StatusEnum("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `StatusEnumUnsafe`  | UNSAFE              |
| `StatusEnumUnknown` | UNKNOWN             |
| `StatusEnumQueued`  | QUEUED              |
| `StatusEnumSafe`    | SAFE                |