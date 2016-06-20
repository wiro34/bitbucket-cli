package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/wiro34/bitbucket-cli/command"
	"github.com/wiro34/bitbucket-cli/config"
)

func Run(args []string) int {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if conf == nil {
		conf, err = config.ResetConfig()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}

	ri, err := config.LoadRepositoryInfo()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	// Meta-option for executables.
	// It defines output color and its stdout/stderr stream.
	meta := &command.Meta{
		Ui: &cli.ColoredUi{
			InfoColor:  cli.UiColorBlue,
			ErrorColor: cli.UiColorRed,
			Ui: &cli.BasicUi{
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
				Reader:      os.Stdin,
			},
		},
		RepositoryInfo: ri,
		Config:         conf,
	}

	return RunCustom(args, Commands(meta))
}

func RunCustom(args []string, commands map[string]cli.CommandFactory) int {

	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	// Convert short-form command to full-form
	if len(args) > 0 {
		switch args[0] {
		case "pr":
			args[0] = "pull-request"
		}
	}

	cli := &cli.CLI{
		Args:       args,
		Commands:   commands,
		Version:    Version,
		HelpFunc:   cli.BasicHelpFunc(Name),
		HelpWriter: os.Stdout,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute: %s\n", err.Error())
	}

	return exitCode
}
