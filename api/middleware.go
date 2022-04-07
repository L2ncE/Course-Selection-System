package api

import (
	"CSA/model"
	"CSA/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

//func TlsHandler(port int) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		secureMiddleware := secure.New(secure.Options{
//			SSLRedirect: true,
//			SSLHost:     ":" + strconv.Itoa(port),
//		})
//		err := secureMiddleware.Process(c.Writer, c.Request)
//
//		// If there was an error, do not continue.
//		if err != nil {
//			return
//		}
//
//		c.Next()
//	}
//}
