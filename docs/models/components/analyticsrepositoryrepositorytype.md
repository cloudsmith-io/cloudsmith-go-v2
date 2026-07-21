# AnalyticsRepositoryRepositoryType


## Supported Types

### AnalyticsRepositoryRepositoryTypeEnum

```go
analyticsRepositoryRepositoryType := components.CreateAnalyticsRepositoryRepositoryTypeAnalyticsRepositoryRepositoryTypeEnum(components.AnalyticsRepositoryRepositoryTypeEnum{/* values here */})
```

### 

```go
analyticsRepositoryRepositoryType := components.CreateAnalyticsRepositoryRepositoryTypeAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch analyticsRepositoryRepositoryType.Type {
	case components.AnalyticsRepositoryRepositoryTypeTypeAnalyticsRepositoryRepositoryTypeEnum:
		// analyticsRepositoryRepositoryType.AnalyticsRepositoryRepositoryTypeEnum is populated
	case components.AnalyticsRepositoryRepositoryTypeTypeAny:
		// analyticsRepositoryRepositoryType.Any is populated
	default:
		// Unknown type - use analyticsRepositoryRepositoryType.GetUnknownRaw() for raw JSON
}
```
