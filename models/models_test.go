package models

import (
	"reflect"
	"testing"
)

func TestReturnDBLoggers(t *testing.T) {
	LogMongoDRV, LogMongoTX := ReturnDBLoggers()

	result1 := reflect.TypeOf(LogMongoDRV).String()
	result2 := reflect.TypeOf(LogMongoTX).String()

	if result1 != "*log15.logger" {
		t.Fatal("Expected a type of *log15.logger, but got: ", result1)
	}

	if result2 != "*log15.logger" {
		t.Fatal("Expected a type of *log15.logger, but got: ", result2)
	}
}

func TestReturnDGUILoggers(t *testing.T) {
	LogDGUICore := ReturnDGUILoggers()

	result1 := reflect.TypeOf(LogDGUICore).String()

	if result1 != "*log15.logger" {
		t.Fatal("Expected a type of *log15.logger, but got: ", result1)
	}
}

func TestReturnAuthLoggers(t *testing.T) {
	LogAuthFileConnect, LogAuthCryptGen, LogAuthCryptCompare := ReturnAuthLoggers()

	result1 := reflect.TypeOf(LogAuthFileConnect).String()
	result2 := reflect.TypeOf(LogAuthCryptGen).String()
	result3 := reflect.TypeOf(LogAuthCryptCompare).String()

	if result1 != "*log15.logger" {
		t.Fatal("Expected a type of *log15.logger, but got: ", result1)
	}

	if result2 != "*log15.logger" {
		t.Fatal("Expected a type of *log15.logger, but got: ", result2)
	}

	if result3 != "*log15.logger" {
		t.Fatal("Expected a type of *log15.logger, but got: ", result3)
	}

}
