# LogsResponseUserType

User type - `Service` or `User`.

* `USER` - User
* `SERVICE` - Service


## Supported Types

### LogsResponseUserTypeEnum

```go
logsResponseUserType := components.CreateLogsResponseUserTypeLogsResponseUserTypeEnum(components.LogsResponseUserTypeEnum{/* values here */})
```

### 

```go
logsResponseUserType := components.CreateLogsResponseUserTypeAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch logsResponseUserType.Type {
	case components.LogsResponseUserTypeTypeLogsResponseUserTypeEnum:
		// logsResponseUserType.LogsResponseUserTypeEnum is populated
	case components.LogsResponseUserTypeTypeAny:
		// logsResponseUserType.Any is populated
	default:
		// Unknown type - use logsResponseUserType.GetUnknownRaw() for raw JSON
}
```
