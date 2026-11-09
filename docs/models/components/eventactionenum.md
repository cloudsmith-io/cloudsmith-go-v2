# EventActionEnum

* `SYNC` - sync
* `DELETE` - delete
* `COPY` - copy
* `MOVE` - move
* `RESYNC` - resync
* `REPUBLISH` - republish
* `UNPUBLISH` - unpublish
* `TAGS_CHANGED` - tags changed
* `QUARANTINE` - quarantined
* `RESTORE` - restored
* `POLICY_VIOLATION` - policy violation
* `HARD_DELETE` - hard delete
* `HIDDEN` - hidden

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.EventActionEnumSync

// Open enum: custom values can be created with a direct type cast
custom := components.EventActionEnum("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `EventActionEnumSync`            | SYNC                             |
| `EventActionEnumDelete`          | DELETE                           |
| `EventActionEnumCopy`            | COPY                             |
| `EventActionEnumMove`            | MOVE                             |
| `EventActionEnumResync`          | RESYNC                           |
| `EventActionEnumRepublish`       | REPUBLISH                        |
| `EventActionEnumUnpublish`       | UNPUBLISH                        |
| `EventActionEnumTagsChanged`     | TAGS_CHANGED                     |
| `EventActionEnumQuarantine`      | QUARANTINE                       |
| `EventActionEnumRestore`         | RESTORE                          |
| `EventActionEnumPolicyViolation` | POLICY_VIOLATION                 |
| `EventActionEnumHardDelete`      | HARD_DELETE                      |
| `EventActionEnumHidden`          | HIDDEN                           |