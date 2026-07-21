# IntervalEnum

* `MINUTE` - Minute
* `HOUR` - Hour
* `DAY` - Day
* `WEEK` - Week
* `MONTH` - Month
* `YEAR` - Year
* `FIVE_MINUTES` - 5 Minutes
* `FIFTEEN_MINUTES` - 15 Minutes

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.IntervalEnumMinute

// Open enum: custom values can be created with a direct type cast
custom := components.IntervalEnum("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `IntervalEnumMinute`         | MINUTE                       |
| `IntervalEnumHour`           | HOUR                         |
| `IntervalEnumDay`            | DAY                          |
| `IntervalEnumWeek`           | WEEK                         |
| `IntervalEnumMonth`          | MONTH                        |
| `IntervalEnumYear`           | YEAR                         |
| `IntervalEnumFiveMinutes`    | FIVE_MINUTES                 |
| `IntervalEnumFifteenMinutes` | FIFTEEN_MINUTES              |