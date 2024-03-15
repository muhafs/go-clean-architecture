package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ExtractID(requestToken string, secret string) (userID string, err error) {
	kf := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}

	token, err := jwt.Parse(requestToken, kf)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		err = fmt.Errorf("invalid Token")
		return
	}

	userID = claims["id"].(string)
	return
}
