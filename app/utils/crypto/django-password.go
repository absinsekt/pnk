package crypto

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"crypto/sha256"
	"crypto/subtle"

	"golang.org/x/crypto/pbkdf2"
)

// GetDjangoPasswordHash returns password hash string for django users
func GetDjangoPasswordHash(password, salt string, iterations int) string {
	pass := pbkdf2.Key([]byte(password), []byte(salt), iterations, 32, sha256.New)
	base := base64.StdEncoding.EncodeToString(pass)

	return fmt.Sprintf("pbkdf2_sha256$%d$%s$%s", iterations, salt, base)
}

// DjangoPasswordEquals checks if given password hash equals one from django db
func DjangoPasswordEquals(input, eq string) bool {
	split := strings.Split(eq, "$")

	_iterations := split[1]
	salt := split[2]

	iterations, _ := strconv.Atoi(_iterations)
	hashed := GetDjangoPasswordHash(input, salt, iterations)

	return subtle.ConstantTimeCompare([]byte(hashed), []byte(eq)) == 1
}
