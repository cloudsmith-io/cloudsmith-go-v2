# PatchedPolicyAction


## Supported Types

### PatchedSetPackageStatePolicyActionTyped

```go
patchedPolicyAction := components.CreatePatchedPolicyActionSetPackageState(components.PatchedSetPackageStatePolicyActionTyped{/* values here */})
```

### PatchedAddPackageTagsPolicyActionTyped

```go
patchedPolicyAction := components.CreatePatchedPolicyActionAddPackageTags(components.PatchedAddPackageTagsPolicyActionTyped{/* values here */})
```

### PatchedRemovePackageTagsPolicyActionTyped

```go
patchedPolicyAction := components.CreatePatchedPolicyActionRemovePackageTags(components.PatchedRemovePackageTagsPolicyActionTyped{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch patchedPolicyAction.Type {
	case components.PatchedPolicyActionTypeSetPackageState:
		// patchedPolicyAction.PatchedSetPackageStatePolicyActionTyped is populated
	case components.PatchedPolicyActionTypeAddPackageTags:
		// patchedPolicyAction.PatchedAddPackageTagsPolicyActionTyped is populated
	case components.PatchedPolicyActionTypeRemovePackageTags:
		// patchedPolicyAction.PatchedRemovePackageTagsPolicyActionTyped is populated
}
```
