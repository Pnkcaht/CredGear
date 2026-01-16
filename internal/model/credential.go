package model

type Credential struct {
	URL      string `json:"url"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
