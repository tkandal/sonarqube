package model

type SQError struct {
	Errors []struct {
		Msg string `json:"msg"`
	} `json:"errors"`
}

type SQPaging struct {
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
	Total     int64 `json:"total"`
}

type SQGroups struct {
	Paging SQPaging `json:"paging"`
	Groups []struct {
		ID           int64  `json:"id"`
		Name         string `json:"name,omitempty"`
		Description  string `json:"description,omitempty"`
		MembersCount int64  `json:"membersCount"`
		Default      bool   `json:"default"`
	} `json:"groups"`
}

type SQUsers struct {
	Paging SQPaging `json:"paging"`
	Users  []struct {
		Login            string   `json:"login"`
		Name             string   `json:"name,omitempty"`
		Active           bool     `json:"active"`
		Email            string   `json:"email,omitempty"`
		Groups           []string `json:"groups"`
		TokenCounts      int64    `json:"tokensCount"`
		Local            bool     `json:"local"`
		ExternalIdentity string   `json:"externalIdentity,omitempty"`
		ExternalProvider string   `json:"externalProvider,omitempty"`
		Avatar           string   `json:"avatar,omitempty"`
	} `json:"users"`
}
