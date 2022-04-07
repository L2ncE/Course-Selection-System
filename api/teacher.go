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
		id := service.GetIDByTInfoName(user)
		tool.RespSuccessfulWithData(ctx, id)
		return
	} else {
		tool.RespErrorWithData(ctx, "请将信息输入完整")
		return
	}
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

func tSecretSecurity(ctx *gin.Context) {
	username := ctx.PostForm("username")
	answer := ctx.PostForm("answer")
	newPassword := ctx.PostForm("new_password")
	if answer == service.GetAnswerByTName(username) {
		l1 := len([]rune(newPassword))
		if l1 <= 16 && l1 >= 6 { //强制规定密码小于16位并大于6位
			//修改新密码
			err := service.ChangeTPassword(username, newPassword)
			if err != nil {
				fmt.Println("change password err: ", err)
				tool.RespInternalError(ctx)
				return
			}

			tool.RespSuccessfulWithData(ctx, "密码正确,密码修改成功")
			return
		} else {
			tool.RespErrorWithData(ctx, "密码请在6位到16位之内")
			return
		}
	}
	tool.RespErrorWithData(ctx, "答案错误")
	return
}

func tQuestion(ctx *gin.Context) {
	username := ctx.PostForm("username")
	question := service.GetQuestionByTName(username)
	if question == "" {
		tool.RespErrorWithData(ctx, "没有此人的密保")
		return
	}
	tool.RespErrorWithData(ctx, question)
}

func changeTPassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	iUsername, _ := ctx.Get("username")
	l1 := len([]rune(newPassword))
	if l1 <= 16 && l1 >= 6 { //强制规定密码小于16位并大于6位
		username := iUsername.(string)

		//检验旧密码是否正确
		flag, err := service.IsTPasswordCorrect(username, oldPassword)
		if err != nil {
			fmt.Println("judge password correct err: ", err)
			tool.RespInternalError(ctx)
			return
		}

		if !flag {
			tool.RespErrorWithData(ctx, "旧密码输入错误")
			return
		}

		//修改新密码
		err = service.ChangeTPassword(username, newPassword)
		if err != nil {
			fmt.Println("change password err: ", err)
			tool.RespInternalError(ctx)
			return
		}

		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithData(ctx, "密码请在6位到16位之内")
		return
	}
}
