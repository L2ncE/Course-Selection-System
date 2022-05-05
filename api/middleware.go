package api

import (
	"CSA/model"
	"CSA/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var mySigningKey = []byte("CSA")

// JWTAuth JWT登录
func JWTAuth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token == "" {
		tool.RespErrorWithData(ctx, "游客你好！没有您的信息,请先登录!")
		ctx.Abort()
		return
	}
	ctx.Set("id", ParseToken(token))
	ctx.Next()
}

func ParseToken(s string) int {
	//解析传过来的token
	tokenClaims, err := jwt.ParseWithClaims(s, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return tokenClaims.Claims.(*model.MyClaims).ID
}

//CORS 跨域中间件
func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token, x-access-token")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}

func LoggerToFile() gin.HandlerFunc {
	logFilePath := "./"
	logFileName := "csa"
	fileName := path.Join(logFilePath, logFileName)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = f
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增hook
	logger.AddHook(hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUrl := c.Request.URL
		statuCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %s | %s",
			statuCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUrl,
		)
	}
}
