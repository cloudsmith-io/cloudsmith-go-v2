# SourceKind

Filter results by metadata source kind.

* `UNKNOWN` - Unknown
* `SYSTEM` - System
* `UPSTREAM` - Upstream
* `CUSTOM` - Custom
* `THIRD_PARTY` - Third Party

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
)

value := operations.SourceKindUnknown
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `SourceKindUnknown`    | UNKNOWN                |
| `SourceKindSystem`     | SYSTEM                 |
| `SourceKindUpstream`   | UPSTREAM               |
| `SourceKindCustom`     | CUSTOM                 |
| `SourceKindThirdParty` | THIRD_PARTY            |