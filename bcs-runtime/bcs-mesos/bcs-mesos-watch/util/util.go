

// Package util xxx
package util

// GetHashId get ID for dispatch channel
func GetHashId(s string, maxInt int) int {
	if maxInt <= 1 {
		return 0
	}

	seed := 131
	hash := 0
	char := []byte(s)

	for _, c := range char {
		hash = hash*seed + int(c)
	}

	return (hash & 0x7FFFFFFF) % maxInt
}
