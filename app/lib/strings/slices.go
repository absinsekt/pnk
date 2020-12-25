package strings

// Contains is a helper function to check if a string exists in a slice
func Contains(vals []string, s string) bool {
	for _, v := range vals {
		if v == s {
			return true
		}
	}

	return false
}

// Filter filters values in a slice
func Filter(vals []string, s string) []string {
	result := []string{}

	for _, v := range vals {
		if v != s {
			result = append(result, v)
		}
	}

	return result
}
