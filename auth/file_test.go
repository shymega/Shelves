package auth

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	err := Connect("./users.test.sqlite3")
	if err != nil {
		t.Fatal("Failed to connect to the Sqlite DB. Error message: ", err)
	}

	// Tidy up left over sqlite file
	err = os.Remove("./users.test.sqlite3")
	if err != nil {
		t.Fatal("Failed to tidy up after myself! Error message", err)
	}
}
