package api

import (
	"CSA/model"
	"CSA/service"
	"CSA/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func registerT(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	question := ctx.PostForm("question")
	answer := ctx.PostForm("answer")
	//输入信息不能为空
	if username != "" && password != "" && question != "" && answer != "" {
		l1 := len([]rune(username))
		l2 := len([]rune(password))
		l3 := len([]rune(question))
		l4 := len([]rune(answer))
		if l1 > 8 || l1 < 1 {
			tool.RespErrorWithData(ctx, "姓名请在1到8位之间")
			return
		}
		if l2 > 16 || l2 < 6 {
			tool.RespErrorWithData(ctx, "密码请在6到16位之间")
			return
		}
		if l3 > 16 {
			tool.RespErrorWithData(ctx, "密保问题请在16个字以内")
			return
		}
		if l4 > 16 {
			tool.RespErrorWithData(ctx, "密保答案请在16个字以内")
			return
		}
		user := model.Teacher{
			Name:     username,
			Password: password,
			Question: question,
			Answer:   answer,
		}

		err := service.RegisterT(user)
		if err != nil {
			fmt.Println("register err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	}
	tool.RespErrorWithData(ctx, "请将信息输入完整")
	return
}

func loginT(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	flag, err := service.IsTPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithData(ctx, "密码错误")
		return
	}
	//jwt
	c := model.MyClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 2592000, //30天，仅做测试
			Issuer:    "YuanXinHao",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessfulWithData(ctx, s)
}
