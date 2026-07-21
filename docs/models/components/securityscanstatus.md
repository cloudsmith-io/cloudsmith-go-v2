# SecurityScanStatus


## Supported Types

### PackageV2SecurityScanStatusEnum

```go
securityScanStatus := components.CreateSecurityScanStatusPackageV2SecurityScanStatusEnum(components.PackageV2SecurityScanStatusEnum{/* values here */})
```

### 

```go
securityScanStatus := components.CreateSecurityScanStatusAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch securityScanStatus.Type {
	case components.SecurityScanStatusTypePackageV2SecurityScanStatusEnum:
		// securityScanStatus.PackageV2SecurityScanStatusEnum is populated
	case components.SecurityScanStatusTypeAny:
		// securityScanStatus.Any is populated
	default:
		// Unknown type - use securityScanStatus.GetUnknownRaw() for raw JSON
}
```
