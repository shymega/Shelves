package database

import (
	"reflect"
	"testing"
)

func TestMongoConnect(t *testing.T) {
	session, err := MongoConnect("localhost")
	result1 := reflect.TypeOf(session).String()
	if result1 != "*mgo.Session" {
		t.Fatal("Expected '*mgo.Session', got: ", result1)
	}

	if err != nil {
		t.Fatal("Encountered an error, failing. Error message: ", err)
	}
}
