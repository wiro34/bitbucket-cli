package command

import (
	"github.com/mitchellh/cli"
	"github.com/wiro34/bitbucket-cli/config"
)

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	Ui             cli.Ui
	RepositoryInfo *config.RepositoryInfo
}
