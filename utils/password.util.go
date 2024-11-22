package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword function is used to hash the password. It returns the hashed password and an error if any.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash function is used to compare the password and the hash. It returns true if the password and the hash are matched, otherwise it returns false.
func CheckPasswordHash(hashedPassword string, reqPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(reqPassword))
}
