package config

import (
	"reflect"
	"testing"
)

func TestGetViper(t *testing.T) {
	viper, err := GetViper("./", "config")
	if err != nil {
		t.Fatal("There was a error getting a viper.Viper type returned from the Shelves Config package. The error message was: ", err)
	}

	vipertype := reflect.TypeOf(viper).String()

	if vipertype != "*viper.Viper" {
		t.Fatal("Expected '*viper.Viper', not: ", vipertype)
	}
}
