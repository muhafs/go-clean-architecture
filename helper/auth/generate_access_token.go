package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/muhafs/go-clean-architecture/domain/entity"
)

func GenerateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(secret))

	return
}
