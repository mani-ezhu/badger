package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// SlackConfig ...
type SlackConfig struct {
	UserName string `mapstructure:"username"`
	HookUrl  string `mapstructure:"url"`
	Channel  string `mapstructure:"channel_name"`
}

// Global variable for config
var Config SlackConfig

type configReader func() error

// ConfigReaderFunc is the viper config reader implementations
func ConfigReaderFunc() configReader {
	return func() error {
		return viper.ReadInConfig()
	}
}

// LoadConfigFromFile ...
// function will load config from yml file
func LoadConfigFromFile(filePath string, cr configReader) (err error) {
	for {
		// setting config file on viper
		viper.SetConfigFile(filePath)
		// reading config and thorwing panic if error occured
		if err := cr(); err != nil {
			break
		}
		// creating slack config
		mapstructure.Decode(viper.GetStringMap("slack"), &Config)
		break
	}
	return

}
