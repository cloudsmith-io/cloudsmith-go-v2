# PackageTypeEnum

* `BINARY` - Binary
* `SOURCE` - Source
* `COMBINED` - Combined
* `OTHER` - Other

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.PackageTypeEnumBinary

// Open enum: custom values can be created with a direct type cast
custom := components.PackageTypeEnum("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `PackageTypeEnumBinary`   | BINARY                    |
| `PackageTypeEnumSource`   | SOURCE                    |
| `PackageTypeEnumCombined` | COMBINED                  |
| `PackageTypeEnumOther`    | OTHER                     |