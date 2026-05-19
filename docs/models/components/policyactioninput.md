# PolicyActionInput


## Supported Types

### SetPackageStatePolicyActionTypedInput

```go
policyActionInput := components.CreatePolicyActionInputSetPackageState(components.SetPackageStatePolicyActionTypedInput{/* values here */})
```

### AddPackageTagsPolicyActionTypedInput

```go
policyActionInput := components.CreatePolicyActionInputAddPackageTags(components.AddPackageTagsPolicyActionTypedInput{/* values here */})
```

### RemovePackageTagsPolicyActionTypedInput

```go
policyActionInput := components.CreatePolicyActionInputRemovePackageTags(components.RemovePackageTagsPolicyActionTypedInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch policyActionInput.Type {
	case components.PolicyActionInputTypeSetPackageState:
		// policyActionInput.SetPackageStatePolicyActionTypedInput is populated
	case components.PolicyActionInputTypeAddPackageTags:
		// policyActionInput.AddPackageTagsPolicyActionTypedInput is populated
	case components.PolicyActionInputTypeRemovePackageTags:
		// policyActionInput.RemovePackageTagsPolicyActionTypedInput is populated
}
```
