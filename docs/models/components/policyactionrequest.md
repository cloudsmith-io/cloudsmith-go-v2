# PolicyActionRequest


## Supported Types

### SetPackageStatePolicyActionTypedRequest

```go
policyActionRequest := components.CreatePolicyActionRequestSetPackageState(components.SetPackageStatePolicyActionTypedRequest{/* values here */})
```

### AddPackageTagsPolicyActionTypedRequest

```go
policyActionRequest := components.CreatePolicyActionRequestAddPackageTags(components.AddPackageTagsPolicyActionTypedRequest{/* values here */})
```

### RemovePackageTagsPolicyActionTypedRequest

```go
policyActionRequest := components.CreatePolicyActionRequestRemovePackageTags(components.RemovePackageTagsPolicyActionTypedRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch policyActionRequest.Type {
	case components.PolicyActionRequestTypeSetPackageState:
		// policyActionRequest.SetPackageStatePolicyActionTypedRequest is populated
	case components.PolicyActionRequestTypeAddPackageTags:
		// policyActionRequest.AddPackageTagsPolicyActionTypedRequest is populated
	case components.PolicyActionRequestTypeRemovePackageTags:
		// policyActionRequest.RemovePackageTagsPolicyActionTypedRequest is populated
}
```
