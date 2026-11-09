# EventAction

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
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
)

value := operations.EventActionSync
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `EventActionSync`            | SYNC                         |
| `EventActionDelete`          | DELETE                       |
| `EventActionCopy`            | COPY                         |
| `EventActionMove`            | MOVE                         |
| `EventActionResync`          | RESYNC                       |
| `EventActionRepublish`       | REPUBLISH                    |
| `EventActionUnpublish`       | UNPUBLISH                    |
| `EventActionTagsChanged`     | TAGS_CHANGED                 |
| `EventActionQuarantine`      | QUARANTINE                   |
| `EventActionRestore`         | RESTORE                      |
| `EventActionPolicyViolation` | POLICY_VIOLATION             |
| `EventActionHardDelete`      | HARD_DELETE                  |
| `EventActionHidden`          | HIDDEN                       |