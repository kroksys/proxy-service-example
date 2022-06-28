package utils

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Proxy ProxyCongif `mapstructure:"proxy"`
}

type ProxyCongif struct {
	Listen string `mapstructure:"listen"`
	Target string `mapstructure:"target"`
	Log    string `mapstructure:"log"`
}

// Reads the configuration file and returns pointer to Config
func ReadConfig(name string) (*Config, error) {
	name = prepareConfigFileName(name)
	viper.SetConfigName(name)
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// move config data into struct
	c := Config{}
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Removes .json, .yaml or .toml file extensions from config file name
func prepareConfigFileName(conf string) string {
	conf = strings.TrimSuffix(conf, ".yaml")
	conf = strings.TrimSuffix(conf, ".toml")
	conf = strings.TrimSuffix(conf, ".json")
	return conf
}
