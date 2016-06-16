package bitbucket

import (
	"net/http"

	"github.com/dghubble/sling"
)

type PullRequestResponse struct {
	PagedResponse
	Values []PullRequest `json:"values"`
}

type PullRequest struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	State  string `json:"state"`
	Author struct {
		User PullRequestUser `json:"user"`
	} `json:"author"`
	From PullRequestRef `json:"fromRef"`
	To   PullRequestRef `json:"toRef"`
	Link struct {
		URL string `json:"url"`
	} `json:"link"`
}

type PullRequestParam struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	State       string                `json:"state"`
	Open        bool                  `json:"open"`
	Closed      bool                  `json:"closed"`
	From        PullRequestRef        `json:"fromRef"`
	To          PullRequestRef        `json:"toRef"`
	Reviewers   []PullRequestReviewer `json:"reviewers"`
	Locked      bool                  `json:"locked"`
}

type PullRequestRef struct {
	ID        string `json:"id"`
	DisplayID string `json:"displayId"`
}

type PullRequestUser struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type PullRequestReviewer struct {
	User PullRequestUser `json:"user"`
}

type PullRequestService struct {
	sling *sling.Sling
}

func (s *PullRequestService) List() ([]PullRequest, error) {
	responseError := new(BitbucketError)
	responseBody := new(PullRequestResponse)

	res, err := s.sling.Receive(responseBody, responseError)
	if isFailure(res) {
		return nil, responseError
	} else if err != nil {
		return nil, err
	}

	return responseBody.Values, nil
}

func (s *PullRequestService) Create(title, description, from, to string, reviewers []string) (*PullRequest, error) {
	responseError := new(BitbucketError)
	responseBody := new(PullRequest)

	req := PullRequestParam{
		Title:       title,
		Description: description,
		State:       "OPEN",
		Open:        true,
		Closed:      false,
		From: PullRequestRef{
			ID: "refs/heads/" + from,
		},
		To: PullRequestRef{
			ID: "refs/heads/" + to,
		},
		Reviewers: make([]PullRequestReviewer, len(reviewers)),
		Locked:    false,
	}
	for i, name := range reviewers {
		req.Reviewers[i] = PullRequestReviewer{PullRequestUser{Name: name}}
	}
	res, err := s.sling.Post("").BodyJSON(req).Receive(responseBody, responseError)
	if isFailure(res) {
		return nil, responseError
	} else if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func NewPullRequestService(sling *sling.Sling) *PullRequestService {
	return &PullRequestService{
		sling: sling.Path("pull-requests"),
	}
}

func isFailure(res *http.Response) bool {
	return res.StatusCode < 200 || res.StatusCode >= 300
}
