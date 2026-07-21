# AnalyticsRepositoryResponseRepositoryType

Repository type - `Public`, `Private` or `Open Source`.

* `PUBLIC` - Public
* `PRIVATE` - Private
* `OPEN_SOURCE` - Open-Source


## Supported Types

### AnalyticsRepositoryResponseRepositoryTypeEnum

```go
analyticsRepositoryResponseRepositoryType := components.CreateAnalyticsRepositoryResponseRepositoryTypeAnalyticsRepositoryResponseRepositoryTypeEnum(components.AnalyticsRepositoryResponseRepositoryTypeEnum{/* values here */})
```

### 

```go
analyticsRepositoryResponseRepositoryType := components.CreateAnalyticsRepositoryResponseRepositoryTypeAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch analyticsRepositoryResponseRepositoryType.Type {
	case components.AnalyticsRepositoryResponseRepositoryTypeTypeAnalyticsRepositoryResponseRepositoryTypeEnum:
		// analyticsRepositoryResponseRepositoryType.AnalyticsRepositoryResponseRepositoryTypeEnum is populated
	case components.AnalyticsRepositoryResponseRepositoryTypeTypeAny:
		// analyticsRepositoryResponseRepositoryType.Any is populated
	default:
		// Unknown type - use analyticsRepositoryResponseRepositoryType.GetUnknownRaw() for raw JSON
}
```
