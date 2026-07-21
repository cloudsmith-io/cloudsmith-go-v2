# MethodEnum

* `UNKNOWN` - Unknown
* `DELETE` - DELETE
* `GET` - GET
* `HEAD` - HEAD
* `OPTIONS` - OPTIONS
* `PATCH` - PATCH
* `POST` - POST
* `PUT` - PUT

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.MethodEnumUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.MethodEnum("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `MethodEnumUnknown` | UNKNOWN             |
| `MethodEnumDelete`  | DELETE              |
| `MethodEnumGet`     | GET                 |
| `MethodEnumHead`    | HEAD                |
| `MethodEnumOptions` | OPTIONS             |
| `MethodEnumPatch`   | PATCH               |
| `MethodEnumPost`    | POST                |
| `MethodEnumPut`     | PUT                 |