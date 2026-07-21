# LogsResponseRepositoryType

Repository type - `Public`, `Private` or `Open Source`.

* `PUBLIC` - Public
* `PRIVATE` - Private
* `OPEN_SOURCE` - Open-Source


## Supported Types

### LogsResponseRepositoryTypeEnum

```go
logsResponseRepositoryType := components.CreateLogsResponseRepositoryTypeLogsResponseRepositoryTypeEnum(components.LogsResponseRepositoryTypeEnum{/* values here */})
```

### 

```go
logsResponseRepositoryType := components.CreateLogsResponseRepositoryTypeAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch logsResponseRepositoryType.Type {
	case components.LogsResponseRepositoryTypeTypeLogsResponseRepositoryTypeEnum:
		// logsResponseRepositoryType.LogsResponseRepositoryTypeEnum is populated
	case components.LogsResponseRepositoryTypeTypeAny:
		// logsResponseRepositoryType.Any is populated
	default:
		// Unknown type - use logsResponseRepositoryType.GetUnknownRaw() for raw JSON
}
```
