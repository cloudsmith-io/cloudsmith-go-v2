# AvailabilityEnum

* `UNAVAILABLE` - Unavailable
* `PARTIAL` - Partially complete
* `COMPLETE` - Complete

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.AvailabilityEnumUnavailable

// Open enum: custom values can be created with a direct type cast
custom := components.AvailabilityEnum("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `AvailabilityEnumUnavailable` | UNAVAILABLE                   |
| `AvailabilityEnumPartial`     | PARTIAL                       |
| `AvailabilityEnumComplete`    | COMPLETE                      |