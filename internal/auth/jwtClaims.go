package auth

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}
