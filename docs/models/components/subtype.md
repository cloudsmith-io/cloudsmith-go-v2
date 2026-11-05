# Subtype

Distinguish skills from plugins in shared/cross-format package listings and policies.

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.SubtypeSkill

// Open enum: custom values can be created with a direct type cast
custom := components.Subtype("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `SubtypeSkill`  | skill           |
| `SubtypePlugin` | plugin          |