package model

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
