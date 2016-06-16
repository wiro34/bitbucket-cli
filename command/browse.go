package command

import (
	"strings"

	"github.com/toqueteos/webbrowser"
)

type BrowseCommand struct {
	Meta
}

func (c *BrowseCommand) Run(args []string) int {
	// Write your code here

	url, err := c.getRepoBrowseUrl()
	if err != nil {
		c.Ui.Error(err.Error())
		return -1
	}
	webbrowser.Open(url)

	return 0
}

func (c *BrowseCommand) Synopsis() string {
	return "Open URL of a repository by browser"
}

func (c *BrowseCommand) Help() string {
	helpText := `
`
	return strings.TrimSpace(helpText)
}

func (c *BrowseCommand) getRepoBrowseUrl() (string, error) {
	browseURL := c.RepositoryInfo.BaseURL +
		"/projects/" + c.RepositoryInfo.Project +
		"/repos/" + c.RepositoryInfo.Repository
	return browseURL, nil
}
