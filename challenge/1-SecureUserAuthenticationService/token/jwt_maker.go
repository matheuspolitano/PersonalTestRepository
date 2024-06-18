package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var MethodInvalid = errors.New("token method is wrong")

type JWTMaker struct {
	secretKey string
}

func (jwtMaker *JWTMaker) CreateToken(payload *Payload) (string, *Payload, error) {
	tokenClain := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := tokenClain.SignedString([]byte(jwtMaker.secretKey))
	if err != nil {
		return "", nil, err
	}
	return token, payload, nil
}

func (jwtMaker *JWTMaker) CheckToken(token string) (*Payload, error) {

	parseFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, MethodInvalid
		}
		return []byte(jwtMaker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, parseFunc)

	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("erro to get claims")
	}

	if err = payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
