package hash

import "golang.org/x/crypto/bcrypt"

//Untuk Hash password dari bcrypt
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//Untuk melakukan validasi password pada saat login
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
