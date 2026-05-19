# RepositoryTypeEnum

* `PUBLIC` - Public
* `PRIVATE` - Private
* `OPEN_SOURCE` - Open-Source

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.RepositoryTypeEnumPublic

// Open enum: custom values can be created with a direct type cast
custom := components.RepositoryTypeEnum("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `RepositoryTypeEnumPublic`     | PUBLIC                         |
| `RepositoryTypeEnumPrivate`    | PRIVATE                        |
| `RepositoryTypeEnumOpenSource` | OPEN_SOURCE                    |