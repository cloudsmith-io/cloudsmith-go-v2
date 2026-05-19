# SourceKindEnum

* `UNKNOWN` - Unknown
* `SYSTEM` - System
* `UPSTREAM` - Upstream
* `CUSTOM` - Custom
* `THIRD_PARTY` - Third Party

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.SourceKindEnumUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.SourceKindEnum("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `SourceKindEnumUnknown`    | UNKNOWN                    |
| `SourceKindEnumSystem`     | SYSTEM                     |
| `SourceKindEnumUpstream`   | UPSTREAM                   |
| `SourceKindEnumCustom`     | CUSTOM                     |
| `SourceKindEnumThirdParty` | THIRD_PARTY                |