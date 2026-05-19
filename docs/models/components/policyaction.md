# PolicyAction


## Supported Types

### SetPackageStatePolicyActionTyped

```go
policyAction := components.CreatePolicyActionSetPackageState(components.SetPackageStatePolicyActionTyped{/* values here */})
```

### AddPackageTagsPolicyActionTyped

```go
policyAction := components.CreatePolicyActionAddPackageTags(components.AddPackageTagsPolicyActionTyped{/* values here */})
```

### RemovePackageTagsPolicyActionTyped

```go
policyAction := components.CreatePolicyActionRemovePackageTags(components.RemovePackageTagsPolicyActionTyped{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch policyAction.Type {
	case components.PolicyActionTypeSetPackageState:
		// policyAction.SetPackageStatePolicyActionTyped is populated
	case components.PolicyActionTypeAddPackageTags:
		// policyAction.AddPackageTagsPolicyActionTyped is populated
	case components.PolicyActionTypeRemovePackageTags:
		// policyAction.RemovePackageTagsPolicyActionTyped is populated
	default:
		// Unknown type - use policyAction.GetUnknownRaw() for raw JSON
}
```
