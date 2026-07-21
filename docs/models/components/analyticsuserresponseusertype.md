# AnalyticsUserResponseUserType

User type - `Service` or `User`.

* `USER` - User
* `SERVICE` - Service


## Supported Types

### AnalyticsUserResponseUserTypeEnum

```go
analyticsUserResponseUserType := components.CreateAnalyticsUserResponseUserTypeAnalyticsUserResponseUserTypeEnum(components.AnalyticsUserResponseUserTypeEnum{/* values here */})
```

### 

```go
analyticsUserResponseUserType := components.CreateAnalyticsUserResponseUserTypeAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch analyticsUserResponseUserType.Type {
	case components.AnalyticsUserResponseUserTypeTypeAnalyticsUserResponseUserTypeEnum:
		// analyticsUserResponseUserType.AnalyticsUserResponseUserTypeEnum is populated
	case components.AnalyticsUserResponseUserTypeTypeAny:
		// analyticsUserResponseUserType.Any is populated
	default:
		// Unknown type - use analyticsUserResponseUserType.GetUnknownRaw() for raw JSON
}
```
