# PackageLogsResponseUserType

Kind of actor that performed the event.

* `USER` - User
* `SERVICE` - Service

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.PackageLogsResponseUserTypeUser

// Open enum: custom values can be created with a direct type cast
custom := components.PackageLogsResponseUserType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `PackageLogsResponseUserTypeUser`    | USER                                 |
| `PackageLogsResponseUserTypeService` | SERVICE                              |