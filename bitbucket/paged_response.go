package bitbucket

type PagedResponse struct {
	Size       int  `json:"size"`
	Limit      int  `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Start      int  `json:"start"`
}
