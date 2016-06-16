package pullRequest

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/toqueteos/webbrowser"
	"github.com/wiro34/bitbucket-cli/bitbucket"
	"github.com/wiro34/bitbucket-cli/config"
)

type CreateCommand struct {
	Ui cli.Ui

	Name           string
	CommandName    string
	RepositoryInfo *config.RepositoryInfo
}

func (c *CreateCommand) Run(args []string) int {
	if len(args) < 5 {
		c.Ui.Warn(c.Help())
		return 1
	}

	client := bitbucket.NewClient(nil)
	pr, err := client.PullRequestService.Create(args[0], args[1], args[2], args[3], args[4:])
	if err != nil {
		c.Ui.Error(err.Error())
		return -1
	}

	webbrowser.Open(c.RepositoryInfo.BaseURL + pr.Link.URL)

	return 0
}

func (c *CreateCommand) Synopsis() string {
	return "Create a new pull request"
}

func (c *CreateCommand) Help() string {
	helpText := fmt.Sprintf(`
usage: %s %s create <title> <description> <from> <to> <reviewers ...>
`, c.Name, c.CommandName)
	return strings.TrimSpace(helpText)
}
