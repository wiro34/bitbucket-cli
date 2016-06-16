package main

import (
	"github.com/mitchellh/cli"
	"github.com/wiro34/bitbucket-cli/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"pull-request": func() (cli.Command, error) {
			return &command.PullRequestCommand{
				Meta: *meta,
				Name: Name,
			}, nil
		},
		"browse": func() (cli.Command, error) {
			return &command.BrowseCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
