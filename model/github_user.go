package model

type Conf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type GitHubUser struct {
	Id       int
	Name     string
	NickName string `gorm:"column:NickName"`
	Identity int
	MajorNum int `gorm:"column:MajorNum"`
}

func (GitHubUser) TableName() string {
	return "github_user"
}
