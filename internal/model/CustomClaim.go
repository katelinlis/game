package model

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	User     int64  `json:"user"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
