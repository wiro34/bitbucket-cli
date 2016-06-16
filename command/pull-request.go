package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/wiro34/bitbucket-cli/command/pull_request"
)

const CommandName = "pull-request"

type PullRequestCommand struct {
	Meta

	Name string
}

func (c *PullRequestCommand) Run(args []string) int {
	return c.runSubCommands(args)
}

func (c *PullRequestCommand) Synopsis() string {
	return ""
}

func (c *PullRequestCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

func (c *PullRequestCommand) runSubCommands(args []string) int {
	cli := &cli.CLI{
		Args:       args,
		Commands:   c.SubCommands(),
		HelpFunc:   SubCommandHelpFunc(c.Name, CommandName),
		HelpWriter: os.Stdout,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute: %s\n", err.Error())
	}

	return exitCode
}

func (c *PullRequestCommand) SubCommands() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &pullRequest.ListCommand{
				Ui: c.Ui,
			}, nil
		},
		"create": func() (cli.Command, error) {
			return &pullRequest.CreateCommand{
				Ui:             c.Ui,
				Name:           c.Name,
				CommandName:    CommandName,
				RepositoryInfo: c.RepositoryInfo,
			}, nil
		},
	}
}
