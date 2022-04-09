package api

import (
	"CSA/model"
	"CSA/service"
	"CSA/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func registerStu(ctx *gin.Context) {
	username := ctx.PostForm("username")
	nickname := ctx.PostForm("nickname")
	password := ctx.PostForm("password")
	question := ctx.PostForm("question")
	answer := ctx.PostForm("answer")
	majorNum := ctx.PostForm("major") //前端做一个选择项
	//输入信息不能为空
	if username != "" && password != "" && question != "" && answer != "" && majorNum != "" && nickname != "" {
		l1 := len([]rune(username))
		l2 := len([]rune(password))
		l3 := len([]rune(question))
		l4 := len([]rune(answer))
		l5 := len([]rune(nickname))
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
		if l5 > 10 {
			tool.RespErrorWithData(ctx, "昵称请在10个字以内")
		}
		ImajorNum, _ := strconv.Atoi(majorNum)
		if ImajorNum >= 20 || ImajorNum < 1 {
			tool.RespErrorWithData(ctx, "专业号在1-20之间，如有疑问请查询")
			return
		}
		flag, err := service.IsRepeatStuNickName(nickname)
		if err != nil {
			fmt.Println("judge repeat username err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		if flag {
			tool.RespErrorWithData(ctx, "用户名已经存在")
			return
		}
		user := model.Student{
			Name:     username,
			NickName: nickname,
			Password: password,
			Question: question,
			Answer:   answer,
			MajorNum: ImajorNum,
		}

		err = service.RegisterStu(user)
		if err != nil {
			fmt.Println("register err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		info := "你好," + username
		tool.RespSuccessfulWithData(ctx, info)
	}
	tool.RespErrorWithData(ctx, "请将信息输入完整")
	return
}

func loginStu(ctx *gin.Context) {
	nickname := ctx.PostForm("nickname")
	password := ctx.PostForm("password")
	id := service.GetIdByTNickName(nickname)
	username := service.GetNameByStuId(id)
	flag, err := service.IsStuPasswordCorrect(id, password)
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
		ID:       id,
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

func stuSecretSecurity(ctx *gin.Context) {
	nickname := ctx.PostForm("nickname")
	id := service.GetIdByStuNickName(nickname)
	answer := ctx.PostForm("answer")
	newPassword := ctx.PostForm("new_password")
	if answer == service.GetAnswerByStuId(id) {
		l1 := len([]rune(newPassword))
		if l1 <= 16 && l1 >= 6 { //强制规定密码小于16位并大于6位
			//修改新密码
			err := service.ChangeStuPassword(id, newPassword)
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

func stuQuestion(ctx *gin.Context) {
	nickname := ctx.PostForm("nickname")
	id := service.GetIdByTNickName(nickname)
	question := service.GetQuestionByStuId(id)
	if question == "" {
		tool.RespErrorWithData(ctx, "没有此人的密保")
		return
	}
	tool.RespErrorWithData(ctx, question)
}

func changeStuPassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	Iid, _ := ctx.Get("id")
	l1 := len([]rune(newPassword))
	if l1 <= 16 && l1 >= 6 { //强制规定密码小于16位并大于6位
		id := Iid.(int)
		//检验旧密码是否正确
		flag, err := service.IsStuPasswordCorrect(id, oldPassword)
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
		err = service.ChangeStuPassword(id, newPassword)
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
