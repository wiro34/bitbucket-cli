package command

import (
	"strings"
)

type PrCommand struct {
	Meta
}

func (c *PrCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *PrCommand) Synopsis() string {
	return ""
}

func (c *PrCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
