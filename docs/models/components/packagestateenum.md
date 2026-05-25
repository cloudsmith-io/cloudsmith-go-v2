# PackageStateEnum

* `AVAILABLE` - The package is available for download.
* `DELETED` - The package is deleted.
* `QUARANTINED` - The package is quarantined.
* `HIDDEN` - The package is hidden.

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.PackageStateEnumAvailable

// Open enum: custom values can be created with a direct type cast
custom := components.PackageStateEnum("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PackageStateEnumAvailable`   | AVAILABLE                     |
| `PackageStateEnumDeleted`     | DELETED                       |
| `PackageStateEnumQuarantined` | QUARANTINED                   |
| `PackageStateEnumHidden`      | HIDDEN                        |