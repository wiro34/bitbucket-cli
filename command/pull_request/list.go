package pullRequest

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/wiro34/bitbucket-cli/bitbucket"
)

type ListCommand struct {
	Ui cli.Ui
}

func (c *ListCommand) Run(args []string) int {

	client := bitbucket.NewClient(nil)
	list, err := client.PullRequestService.List()
	if err != nil {
		c.Ui.Error(err.Error())
		return -1
	}

	for _, v := range list {
		c.Ui.Output(fmt.Sprintf(" #%-4d %-24s %s -> %s",
			v.ID, v.Title, v.From.DisplayID, v.To.DisplayID))
	}

	return 0
}

func (c *ListCommand) Synopsis() string {
	return "Print opened pull requests"
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
