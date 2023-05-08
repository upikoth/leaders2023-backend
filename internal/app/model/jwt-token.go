package model

import "github.com/golang-jwt/jwt/v4"

type JwtTokenUserData struct {
	UserId int `json:"userId"`
}

type JwtTokenClaims struct {
	UserData JwtTokenUserData `json:"userData"`
	jwt.StandardClaims
}
