# Card


## Supported Types

### PolicyHuggingfaceUnknownCardTyped

```go
card := components.CreateCardPolicyHuggingfaceUnknownCardTyped(components.PolicyHuggingfaceUnknownCardTyped{/* values here */})
```

### PolicyHuggingfaceModelCardTyped

```go
card := components.CreateCardPolicyHuggingfaceModelCardTyped(components.PolicyHuggingfaceModelCardTyped{/* values here */})
```

### PolicyHuggingfaceDatasetCardTyped

```go
card := components.CreateCardPolicyHuggingfaceDatasetCardTyped(components.PolicyHuggingfaceDatasetCardTyped{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch card.Type {
	case components.CardTypePolicyHuggingfaceUnknownCardTyped:
		// card.PolicyHuggingfaceUnknownCardTyped is populated
	case components.CardTypePolicyHuggingfaceModelCardTyped:
		// card.PolicyHuggingfaceModelCardTyped is populated
	case components.CardTypePolicyHuggingfaceDatasetCardTyped:
		// card.PolicyHuggingfaceDatasetCardTyped is populated
	default:
		// Unknown type - use card.GetUnknownRaw() for raw JSON
}
```
