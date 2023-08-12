

package utils

// StringInSlice returns true if given string in slice
func StringInSlice(s string, l []string) bool {
	for _, objStr := range l {
		if s == objStr {
			return true
		}
	}
	return false
}

// IntInSlice xxx
func IntInSlice(i int, l []int) bool {
	for _, obj := range l {
		if i == obj {
			return true
		}
	}
	return false
}
