package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

func ParseJWT(tokenString string) (*User, error) {
	claims := JwtClaims{}
	jwtSigningKey := []byte(viper.GetString("auth.secret"))

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	user := &User{
		ID:       claims.Subject,
		Username: claims.Username,
		Name:     claims.Name,
	}

	return user, nil
}

func MakeJWT(user *User) (string, error) {
	if user == nil {
		return "", errors.New("no user provided")
	}

	claims := JwtClaims{
		Username: user.Username,
		Name:     user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "calapi",
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwtSigningKey := []byte(viper.GetString("auth.secret"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}
