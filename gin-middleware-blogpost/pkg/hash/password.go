package hash

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func Hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := os.Getenv("PASSWORD_SALT")
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
