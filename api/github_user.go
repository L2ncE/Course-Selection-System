package api

import (
	"CSA/model"
	"CSA/service"
	"CSA/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func registerGitHub(ctx *gin.Context, name string) {
	majorNum := ctx.PostForm("major") //前端做一个选择项
	username := ctx.PostForm("username")
	//输入信息不能为空
	if majorNum != "" {
		ImajorNum, _ := strconv.Atoi(majorNum)
		if ImajorNum >= 20 || ImajorNum < 1 {
			tool.RespErrorWithData(ctx, "专业号在1-20之间，如有疑问请查询")
			return
		}

		user := model.GitHubUser{
			Name:     name,
			NickName: username,
			MajorNum: ImajorNum,
		}

		err := service.RegisterGitHub(user)
		if err != nil {
			fmt.Println("register err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		info := "你好," + username
		tool.RespSuccessfulWithData(ctx, info)
		return
	}
	tool.RespErrorWithData(ctx, "请将信息输入完整")
	return
}

func Oauth(ctx *gin.Context) {
	var err error
	var code = ctx.Query("code")

	var tokenAuthUrl = service.GetTokenAuthUrl(code)
	var token *model.Token
	if token, err = service.GetToken(tokenAuthUrl); err != nil {
		tool.RespInternalError(ctx)
		return
	}

	// 通过token，获取用户信息
	var userInfo map[string]interface{}
	userInfo, err = service.GetUserInfo(token)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}
	var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	str := ""
	length := len(chars)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		str += chars[rand.Intn(length)]
	}
	user := userInfo["login"].(string) + str
	flag, _ := service.IsRepeatGitHubName(user)

	if flag == false {
		registerGitHub(ctx, user)
	}

	c := model.MyClaims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 6000000,
			Issuer:    "YuanXinHao",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithTwoData(ctx, user, s)
	return

}
