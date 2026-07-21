# PolicyDecisionLogV1FacetsResponse


## Supported Types

### PolicyDecisionLogV1FacetsPackageResponse

```go
policyDecisionLogV1FacetsResponse := components.CreatePolicyDecisionLogV1FacetsResponsePolicyDecisionLogV1FacetsPackageResponse(components.PolicyDecisionLogV1FacetsPackageResponse{/* values here */})
```

### PolicyDecisionLogV1FacetsRepositoryResponse

```go
policyDecisionLogV1FacetsResponse := components.CreatePolicyDecisionLogV1FacetsResponsePolicyDecisionLogV1FacetsRepositoryResponse(components.PolicyDecisionLogV1FacetsRepositoryResponse{/* values here */})
```

### PolicyDecisionLogV1FacetsPackageFormatResponse

```go
policyDecisionLogV1FacetsResponse := components.CreatePolicyDecisionLogV1FacetsResponsePolicyDecisionLogV1FacetsPackageFormatResponse(components.PolicyDecisionLogV1FacetsPackageFormatResponse{/* values here */})
```

### PolicyDecisionLogV1FacetsPolicyResponse

```go
policyDecisionLogV1FacetsResponse := components.CreatePolicyDecisionLogV1FacetsResponsePolicyDecisionLogV1FacetsPolicyResponse(components.PolicyDecisionLogV1FacetsPolicyResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch policyDecisionLogV1FacetsResponse.Type {
	case components.PolicyDecisionLogV1FacetsResponseTypePolicyDecisionLogV1FacetsPackageResponse:
		// policyDecisionLogV1FacetsResponse.PolicyDecisionLogV1FacetsPackageResponse is populated
	case components.PolicyDecisionLogV1FacetsResponseTypePolicyDecisionLogV1FacetsRepositoryResponse:
		// policyDecisionLogV1FacetsResponse.PolicyDecisionLogV1FacetsRepositoryResponse is populated
	case components.PolicyDecisionLogV1FacetsResponseTypePolicyDecisionLogV1FacetsPackageFormatResponse:
		// policyDecisionLogV1FacetsResponse.PolicyDecisionLogV1FacetsPackageFormatResponse is populated
	case components.PolicyDecisionLogV1FacetsResponseTypePolicyDecisionLogV1FacetsPolicyResponse:
		// policyDecisionLogV1FacetsResponse.PolicyDecisionLogV1FacetsPolicyResponse is populated
	default:
		// Unknown type - use policyDecisionLogV1FacetsResponse.GetUnknownRaw() for raw JSON
}
```
