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
