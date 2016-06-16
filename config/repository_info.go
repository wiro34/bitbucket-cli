package config

import (
	"errors"
	"regexp"

	"github.com/spf13/viper"
	"github.com/tcnksm/go-gitconfig"
)

var rexp = regexp.MustCompile(`.+?://(.+@)?([^/]+?):?(\d+)?/([^/]+?)/([^/]+)\.git`)

type RepositoryInfo struct {
	BaseURL    string
	Project    string
	Repository string
}

func LoadRepositoryInfo() (*RepositoryInfo, error) {
	url, err := gitconfig.OriginURL()
	if err != nil {
		return nil, err
	}

	url = "ssh://git@git.so-net.co.jp:7999/sandbox/pullrequestflowtest.git"

	result := rexp.FindAllStringSubmatch(url, -1)
	if len(result) != 1 || len(result[0]) != 6 {
		return nil, errors.New("This repository may not be the one of Bitbucket.")
	}

	return &RepositoryInfo{
		BaseURL:    viper.GetString("bitbucket.base_url"),
		Project:    result[0][4],
		Repository: result[0][5],
	}, nil
}
