package config

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func LoadConfig() {
	dir, err := homedir.Dir()
	if err != nil {
		panic(fmt.Errorf("Cannot get user home directory."))
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir + "/.bitbucket-cli")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load config file: %s \n", err))
	}

	if username := viper.GetString("bitbucket.username"); username == "" {
		panic(fmt.Errorf("username is not set"))
	}

	if password := viper.GetString("bitbucket.password"); password == "" {
		panic(fmt.Errorf("password is not set"))
	}

	if baseURL := viper.GetString("bitbucket.base_url"); baseURL == "" {
		panic(fmt.Errorf("base_url is not set"))
	}
}
