package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func Verify(requestToken string, secret string) (isAuthorized bool, err error) {
	kf := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	}

	if _, err = jwt.Parse(requestToken, kf); err != nil {
		return
	}

	isAuthorized = true
	return
}
