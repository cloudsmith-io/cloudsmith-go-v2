# PolicyPolymorphicPackage


## Supported Types

### PolicyPackageTyped

```go
policyPolymorphicPackage := components.CreatePolicyPolymorphicPackagePackage(components.PolicyPackageTyped{/* values here */})
```

### PolicyHuggingfacePackageTyped

```go
policyPolymorphicPackage := components.CreatePolicyPolymorphicPackageHuggingfacePackage(components.PolicyHuggingfacePackageTyped{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch policyPolymorphicPackage.Type {
	case components.PolicyPolymorphicPackageTypePackage:
		// policyPolymorphicPackage.PolicyPackageTyped is populated
	case components.PolicyPolymorphicPackageTypeHuggingfacePackage:
		// policyPolymorphicPackage.PolicyHuggingfacePackageTyped is populated
	default:
		// Unknown type - use policyPolymorphicPackage.GetUnknownRaw() for raw JSON
}
```
