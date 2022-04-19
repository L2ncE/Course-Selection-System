package model

type Conf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

type Token struct {
	AccessToken string `json:"access_token"`
}
