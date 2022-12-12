package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (hashBytes []byte, err error) {
	hashBytes, err = bcrypt.GenerateFromPassword([]byte(password), 4)
	return
}

func ComparePasswordHash(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
