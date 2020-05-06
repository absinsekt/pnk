package strings

import (
	"crypto/rand"
)

// GenerateRandomString returns random string with a given length
func GenerateRandomString(length int64) string {
	const dict = "ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890abcdefghijklmnopqrstuvwxyz"
	rnd := make([]byte, length)

	if _, err := rand.Read(rnd); err != nil {
		return ""
	}

	for idx, bt := range rnd {
		rnd[idx] = dict[bt%byte(len(dict))]
	}

	return string(rnd)
}
