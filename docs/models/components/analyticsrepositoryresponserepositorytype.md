# AnalyticsRepositoryResponseRepositoryType

Repository type - `Public`, `Private` or `Open Source`.

* `PUBLIC` - Public
* `PRIVATE` - Private
* `OPEN_SOURCE` - Open-Source

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.AnalyticsRepositoryResponseRepositoryTypePublic

// Open enum: custom values can be created with a direct type cast
custom := components.AnalyticsRepositoryResponseRepositoryType("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `AnalyticsRepositoryResponseRepositoryTypePublic`     | PUBLIC                                                |
| `AnalyticsRepositoryResponseRepositoryTypePrivate`    | PRIVATE                                               |
| `AnalyticsRepositoryResponseRepositoryTypeOpenSource` | OPEN_SOURCE                                           |