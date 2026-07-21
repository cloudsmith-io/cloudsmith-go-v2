# ThreatLevel

Filter by threat level. Matches the single best classified severity. Accepts normalized threat levels (critical, high, medium, low, unknown, malware) and vendor aliases (important, moderate, negligible, none, unimportant).

* `critical` - critical
* `high` - high
* `important` - important
* `low` - low
* `malware` - malware
* `medium` - medium
* `moderate` - moderate
* `negligible` - negligible
* `none` - none
* `unimportant` - unimportant
* `unknown` - unknown

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
)

value := operations.ThreatLevelCritical
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ThreatLevelCritical`    | critical                 |
| `ThreatLevelHigh`        | high                     |
| `ThreatLevelImportant`   | important                |
| `ThreatLevelLow`         | low                      |
| `ThreatLevelMalware`     | malware                  |
| `ThreatLevelMedium`      | medium                   |
| `ThreatLevelModerate`    | moderate                 |
| `ThreatLevelNegligible`  | negligible               |
| `ThreatLevelNone`        | none                     |
| `ThreatLevelUnimportant` | unimportant              |
| `ThreatLevelUnknown`     | unknown                  |