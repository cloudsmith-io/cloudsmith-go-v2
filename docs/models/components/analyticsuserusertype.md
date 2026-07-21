# AnalyticsUserUserType


## Supported Types

### AnalyticsUserUserTypeEnum

```go
analyticsUserUserType := components.CreateAnalyticsUserUserTypeAnalyticsUserUserTypeEnum(components.AnalyticsUserUserTypeEnum{/* values here */})
```

### 

```go
analyticsUserUserType := components.CreateAnalyticsUserUserTypeAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch analyticsUserUserType.Type {
	case components.AnalyticsUserUserTypeTypeAnalyticsUserUserTypeEnum:
		// analyticsUserUserType.AnalyticsUserUserTypeEnum is populated
	case components.AnalyticsUserUserTypeTypeAny:
		// analyticsUserUserType.Any is populated
	default:
		// Unknown type - use analyticsUserUserType.GetUnknownRaw() for raw JSON
}
```
