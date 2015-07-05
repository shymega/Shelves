package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// PWDECrypt takes two arguments, the password to	be encrypted, and the cost.
// It then returns and the hashed Password generated in a string form.
// It also returns an error value, which normally would be nil, but
// if anything went wrong it would return the error.
func PWDECrypt(password string, cost int) (string, error) {
	bytePassword := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, cost)
	if err != nil {
		return string(hashedPassword), err
	}
	return string(hashedPassword), nil
}

// PWDCryptComp takes two arguments, the hashed Password from generation
// earlier, and the password sent to Shelves from the Web Service, for
// example. It then returns an error type, which is nil if the password
// matches, or an error if it doesn't match.
func PWDCryptComp(hashedPassword string, password string) error {
	bytehashedPassword := []byte(hashedPassword)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(bytehashedPassword, bytePassword)
	return err
}
