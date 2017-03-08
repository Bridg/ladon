package ladon

// ContextKeyValuesEqualCondition is a condition which is fulfilled if the given keys,
// Key1 and Key2 have the same lookup values in the provided context map value
type ContextKeyValuesEqualCondition struct {
	Key1 string `json:"key1"`

	Key2 string `json:"key2"`
}

// Fulfills returns true if the given value is a map of strings to strings and
// the maps values for ContextKeyValuesEqualCondition.Key1 and
// ContextKeyValuesEqualCondition.Key2 are equal
func (c *ContextKeyValuesEqualCondition) Fulfills(value interface{}, _ *Request) bool {
	m, ok := value.(map[string]string)

	return ok && m[c.Key1] == m[c.Key2]
}

// GetName returns the condition's name.
func (c *ContextKeyValuesEqualCondition) GetName() string {
	return "ContextKeyValuesEqualCondition"
}
