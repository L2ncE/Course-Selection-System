package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}
