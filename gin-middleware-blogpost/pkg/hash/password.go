package hash

import (
	"crypto/sha1"
	"fmt"
)

func Hash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := "PASSWORD_SALT"
	hashed := fmt.Sprintf("%x", hash.Sum([]byte(salt)))
	fmt.Println(`hashed=`, hashed)
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
