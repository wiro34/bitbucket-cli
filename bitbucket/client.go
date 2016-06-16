package bitbucket

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/spf13/viper"
	"github.com/wiro34/bitbucket-cli/config"
)

type Client struct {
	PullRequestService *PullRequestService
}

type BitbucketError struct {
	Errors []struct {
		Context       string `json:"context"`
		Message       string `json:"message"`
		ExceptionName string `json:"exceptionName"`
	} `json:"errors"`
}

func (e BitbucketError) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprintf("BitbucketError: %v", e.Errors[0].Message)
	}
	return ""
}

func NewClient(httpClient *http.Client) *Client {
	ri, err := config.LoadRepositoryInfo()
	if err != nil {
		return nil
	}

	sling := sling.New().Client(httpClient).
		Base("https://git.so-net.co.jp:8443/rest/api/1.0/").
		Path("projects/").Path(ri.Project+"/").
		Path("repos/").Path(ri.Repository+"/").
		SetBasicAuth(viper.GetString("bitbucket.username"), viper.GetString("bitbucket.password"))

	return &Client{
		PullRequestService: NewPullRequestService(sling),
	}
}
