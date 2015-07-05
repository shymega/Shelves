package auth

import (
	"testing"
)

var password string = "testingpassw0rd!"
var hashedPassword string

func TestPWDECrypt(t *testing.T) {
	var err error
	// Encrypt a password with the Shelves crypto library
	hashedPassword, err = PWDECrypt(password, 10)
	if err != nil {
		t.Fatal("The function, PWDECrypt, failed to hash the password passed to it. The error message was: ", err)
	}
}

func TestPWDCryptComp(t *testing.T) {
	// Compare the unhashed password and the hashedPassword from earlier.
	err := PWDCryptComp(hashedPassword, password)
	if err != nil {
		t.Fatal("The password didn't match. The error message was: ", err)
	}
}
