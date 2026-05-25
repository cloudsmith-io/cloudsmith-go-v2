# PackageTagTypeEnum

* `INFO` - Info
* `SOME_VERSION` - Some Version
* `LATEST_VERSION` - Latest Version
* `ANNOTATION` - Annotation

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.PackageTagTypeEnumInfo

// Open enum: custom values can be created with a direct type cast
custom := components.PackageTagTypeEnum("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PackageTagTypeEnumInfo`          | INFO                              |
| `PackageTagTypeEnumSomeVersion`   | SOME_VERSION                      |
| `PackageTagTypeEnumLatestVersion` | LATEST_VERSION                    |
| `PackageTagTypeEnumAnnotation`    | ANNOTATION                        |