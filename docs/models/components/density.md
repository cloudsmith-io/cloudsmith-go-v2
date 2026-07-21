# Density

Row density for the table. Null means use the default.

* `COMFORTABLE` - Comfortable
* `COMPACT` - Compact


## Supported Types

### DensityEnum

```go
density := components.CreateDensityDensityEnum(components.DensityEnum{/* values here */})
```

### 

```go
density := components.CreateDensityAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch density.Type {
	case components.DensityTypeDensityEnum:
		// density.DensityEnum is populated
	case components.DensityTypeAny:
		// density.Any is populated
	default:
		// Unknown type - use density.GetUnknownRaw() for raw JSON
}
```
