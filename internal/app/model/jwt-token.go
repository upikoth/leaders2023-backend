package model

import "github.com/golang-jwt/jwt/v4"

type JwtTokenUserData struct {
	UserID   string `json:"userId"`
	UserRole Role   `json:"userRole"`
}

type JwtTokenClaims struct {
	UserData JwtTokenUserData `json:"userData"`
	jwt.StandardClaims
}
