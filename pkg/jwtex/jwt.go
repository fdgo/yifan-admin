package jwtex

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"yifan/pkg/stringx"
)

type JwtToken struct {
	SigningKey []byte
}

func (j *JwtToken) CreateToken(claims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
func (j *JwtToken) ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("toke错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("toke过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("toke错误")
			} else {
				return nil, errors.New("toke错误")
			}
		}
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("toke错误")
}

//Decode 解码
func (j *JwtToken) Decode(r *http.Request, w http.ResponseWriter) (*jwt.StandardClaims, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errors.New("缺少token")
	}
	if !strings.HasPrefix(token, "Bearer ") {
		return nil, errors.New("缺少token")
	}
	if stringx.Length(token) < 128 {
		return nil, errors.New("缺少token")
	}
	token = stringx.SubString(token, stringx.Length("Bearer "), stringx.Length(token)-stringx.Length("Bearer "))
	sub, err := j.ParseToken(token)
	if sub == nil {
		return nil, err
	}
	return sub, nil
}
