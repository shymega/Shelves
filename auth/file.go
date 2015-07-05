package auth

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	// log "gopkg.in/inconshreveable/log15.v2"
	"errors"
	"os"
)

// Users defines the structure of the sqlite3 Users database.
// The database contains the User's ID, Username and HashedPassword.
type Users struct {
	ID            int    `db:"ID"`
	Username      string `db:"Username"`
	CryptPassword string `db:"CryptPassword"`
}

// schema populates the sqlite3 Users database
// and sets the ID column as the Primary Key which
// auto increments on new entries.
var schema = `
CREATE TABLE Users (
ID INTEGER NOT NULL,
Username VARCHAR(25),
CryptPassword VARCHAR(275),
PRIMARY KEY (ID)
);`

// db is a global variable of *sqlx.DB type.
var db *sqlx.DB

// Connect connects to the sqlite DB passed onto the function
// through the global variable 'db'.
func Connect(sqlitefile string) error {
	var err error

	db, err = sqlx.Connect("sqlite3", sqlitefile)
	if err != nil {
		return err
	}

	return nil
}

// PopulateDB removes the sqlite DB, and recreates and populates it.
func PopulateDB(sqlitefile string) error {
	// If the sqlite DB exists, remove it and recreate it.
	if _, err := os.Stat(sqlitefile); os.IsExist(err) {
		err := os.Remove(sqlitefile)
		if err != nil {
			return err
		}
	}

	// After removing the database, we should recreate and connect to the DB.
	err := Connect(sqlitefile)
	if err != nil {
		return err
	}

	// After we have recreated and connected to the DB, we should populate it!
	result, err := db.Exec(schema)
	if err != nil {
		return err
	}
	result.RowsAffected()

	return nil
}

// AddUser takes a string variable called username and a string variable called pwd.
// It then produces a hashed password from the crypt.go file in package auth, and
// inserts it into the SQL database provided.
func AddUser(username string, pwd string) error {
	/** Preliminary check if function arguments are empty **/
	if len(username) == 0 {
		ErrorUsernameNotSpecified := "No Username specified to this function. Ignore and Exit."
		return errors.New(ErrorUsernameNotSpecified)
	}

	if len(pwd) == 0 {
		ErrorPasswordNotSpecifed := "No 'new' password specified to this function. Ignore and Exit."
		return errors.New(ErrorPasswordNotSpecifed)
	}

	cryptedPassword, err := PWDECrypt(pwd, 10)
	if err != nil {
		return err
	}

	// Add a user to the DB.
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO Users (Username, CryptPassword) VALUES ($1, $2)", username, cryptedPassword)
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// UserPwdChange takes two arguments, a username and a new password, both a string type.
// It then checks if the username exists in the DB. If it does, it generates a new hashed password
// and updates the record of the User with it.
func UserPwdChange(username string, newpwd string) error {
	/** Preliminary check if function arguments are empty **/

	if len(username) == 0 {
		ErrorUsernameNotSpecified := "No Username specified to this function. Ignore and Exit."
		return errors.New(ErrorUsernameNotSpecified)
	}

	if len(newpwd) == 0 {
		ErrorNewPasswordNotSpecifed := "No 'new' password specified to this function. Ignore and Exit."
		return errors.New(ErrorNewPasswordNotSpecifed)
	}

	// By this point, a username and a new password must have been specified!

	// Time to check if the username specified to the function exists in the DB.
	user := Users{}
	err := db.Get(&user, "SELECT ID, Username, CryptPassword FROM Users WHERE Username=$1", username)
	if err != nil {
		return err
	}

	if len(user.Username) == 0 {
		ErrorUsernameExistsInDB := "The user specified doesn't exist in the DB. You might need to create the user first! Ignore and Exit."
		return errors.New(ErrorUsernameExistsInDB)
	}

	// Now to generate a encrypted password for the password replacement.
	cryptedPassword := []byte(newpwd)

	// The new encrypted password should have been generated now. Replacing in DB.
	_, err = db.Exec("UPDATE User SET CryptPassword = $1 WHERE Username=$2", username, cryptedPassword)
	if err != nil {
		return err
	}

	// The password should be replaced, returning nil error value to show that the change was successful.
	return nil
}

// UserAuth takes a string variable called 'username' and a string variable called webpwd.
// It grabs the encrypted password from the sqlite database and then compares it using the crypt.go
// file from the auth package with the password passed to the function.
func UserAuth(username string, webpwd string) (error, error) {
	/** Preliminary check if function arguments are empty **/
	if len(username) == 0 {
		ErrorUsernameNotSpecified := "No Username specified to this function. Ignore and Exit."
		return errors.New(ErrorUsernameNotSpecified), nil
	}

	if len(webpwd) == 0 {
		ErrorWebPasswordNotSpecifed := "No web password specifed to this function. Ignore and Exit."
		return errors.New(ErrorWebPasswordNotSpecifed), nil
	}

	// Create a local representation of the Users struct here.
	user := Users{}

	// Retrieve the CryptedPassword from the database.
	err := db.Get(&user, "SELECT CryptPassword FROM Users WHERE Username == $1;", username)
	if err != nil {
		return err, nil
	}

	// Compare the two passwords, DB and function argument, if they match return nil, if not return error.
	result := PWDCryptComp(user.CryptPassword, webpwd)
	if result != nil {
		return nil, result
	}

	return nil, nil
}
