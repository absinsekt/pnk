package strings

func MapValues(m map[string]interface{}) []interface{} {
	result := make([]interface{}, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
