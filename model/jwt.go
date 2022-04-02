package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
