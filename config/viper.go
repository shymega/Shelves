package config

import (
	"github.com/spf13/viper"
)

// GetViper takes the configPath and configName from the main Shelves file, and returns
// a Viper instance, and a error message, *should* anything go wrong.
func GetViper(configPath string, configName string) (*viper.Viper, error) {

	config := viper.New() // Create a new Viper instance.

	config.AddConfigPath(configPath) //	Add the configPath passed in the function arguments.
	config.SetConfigName(configName) // Add the configName passed in the function arguments.

	err := config.ReadInConfig()
	if err != nil {
		return config, err // If a error is found when getting the config to read the configuration file,
		// return the config as it is, and the error associated with it.
	}

	return config, nil // Return *viper.Viper for configuration, and a nil error value.
}
