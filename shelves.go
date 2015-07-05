package main

import (
	"flag"
	"fmt"
	"github.com/shymega/shelves/config"
	"github.com/shymega/shelves/controllers"
	"github.com/spf13/viper"
	"os"
)

// Initialize flags
var (
	configLocation string
	debug          bool
)

var Config *viper.Viper

func init() {
	// Populate flags.
	flag.StringVar(&configLocation, "configLocation", "./", "Where the configuration file is located, this should point to the directory NOT to the file itself")
	flag.BoolVar(&debug, "debug", false, "Run Shelves in debugging mode.")
	flag.Parse()
	// Flags initialized.

	// Check if the flag configLocation is not the default value, end-user should specify a full path to the directory, NOT the file.
	if configLocation == "./" {
		fmt.Println("You should be specifing a path to the config file, but NOT the file itself, only the directory.")
		fmt.Println("The configuration file CANNOT be in this directory.")
		fmt.Println("Try appending -help to your command.")
		os.Exit(1)
	}

	var err error // Define the err variable
	Config, err = config.GetViper(configLocation, "config")
	if err != nil {
		panic(err) // TODO: This should be handled better!
	}

}

func main() {
	// Start the REST/Web Server.
	// TODO: Look into channels, and goroutines.
	controllers.StartREST(":8080", "./public")
}
