package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(word string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(word), 2)

	return string(hash)
}

func CompareHash(word, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(word))
}
