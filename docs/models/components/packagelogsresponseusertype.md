# PackageLogsResponseUserType

Kind of actor that performed the event.

* `USER` - User
* `SERVICE` - Service


## Supported Types

### PackageLogsResponseUserTypeEnum

```go
packageLogsResponseUserType := components.CreatePackageLogsResponseUserTypePackageLogsResponseUserTypeEnum(components.PackageLogsResponseUserTypeEnum{/* values here */})
```

### 

```go
packageLogsResponseUserType := components.CreatePackageLogsResponseUserTypeAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch packageLogsResponseUserType.Type {
	case components.PackageLogsResponseUserTypeTypePackageLogsResponseUserTypeEnum:
		// packageLogsResponseUserType.PackageLogsResponseUserTypeEnum is populated
	case components.PackageLogsResponseUserTypeTypeAny:
		// packageLogsResponseUserType.Any is populated
	default:
		// Unknown type - use packageLogsResponseUserType.GetUnknownRaw() for raw JSON
}
```
