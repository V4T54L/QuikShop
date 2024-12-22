package utils

import "crypto"

func HashPassword(password string) string {
	// Generate SHA256 hash
	hash := crypto.SHA256.New()
	hash.Write([]byte(password))
	return string(hash.Sum(nil))
}

func ComparePassword(password, hashedPassword string) bool {
	// Compare the password with the hashed password
	return HashPassword(password) == hashedPassword
}
