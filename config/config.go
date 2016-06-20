package config

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

type Config struct {
	Username string
	Password string
	BaseURL  string
}

const separator = string(filepath.Separator)
const configDirName = ".bitbucket-cli"

var config Config

func LoadConfig() (*Config, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return nil, errors.New("Cannot get user home directory.")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir + separator + configDirName)
	err = viper.ReadInConfig()
	if err != nil {
		// it's ok.
		return nil, nil
	}

	var username, password, baseURL string
	if username = viper.GetString("bitbucket.username"); username == "" {
		return nil, errors.New("username is not set")
	}

	if password = viper.GetString("bitbucket.password"); password == "" {
		return nil, errors.New("password is not set")
	}

	if baseURL = viper.GetString("bitbucket.base_url"); baseURL == "" {
		return nil, errors.New("base_url is not set")
	}

	return &Config{
		Username: username,
		Password: password,
		BaseURL:  baseURL,
	}, nil
}

func InputConfig() (*Config, error) {
	username, err := inputWithPrompt("Bitbucket username: ")
	if err != nil {
		return nil, err
	}

	password, err := inputPasswordWithPrompt("Password for " + username + " (this stores $HOME/.bitbucket-cli/config.yml): ")
	if err != nil {
		return nil, err
	}

	baseURL, err := inputWithPrompt("Bitbucket base URL (i.e. https://your.domain.com/): ")
	if err != nil {
		return nil, err
	}

	return &Config{
		Username: username,
		Password: password,
		BaseURL:  baseURL,
	}, nil
}

func SaveConfig(config *Config) error {
	dir, err := homedir.Dir()
	if err != nil {
		return errors.New("Cannot get user home directory.")
	}

	configFilePath := dir + separator + configDirName + separator + "config.yml"
	os.MkdirAll(filepath.Dir(configFilePath), 0755)

	configStr := strings.Trim(fmt.Sprintf(`
bitbucket:
  username: %s
  password: %s
  base_url: %s
`, config.Username, config.Password, config.BaseURL), "\n")
	return ioutil.WriteFile(configFilePath, []byte(configStr), os.ModePerm)
}

func ResetConfig() (*Config, error) {
	conf, err := InputConfig()
	if err != nil {
		return nil, err
	}
	if err = SaveConfig(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func inputWithPrompt(prompt string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		if input := scanner.Text(); len(input) > 0 {
			return input, nil
		}
	}
	return "", errors.New("Aborted")
}

func inputPasswordWithPrompt(prompt string) (string, error) {
	fmt.Print(prompt)
	password, err := terminal.ReadPassword(0)
	if err != nil {
		return "", errors.New("Failed to input password")
	}
	fmt.Println("")
	return string(password), nil
}
