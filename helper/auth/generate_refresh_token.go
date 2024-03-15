package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/muhafs/go-clean-architecture/domain/entity"
)

func GenerateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	refreshToken, err = token.SignedString([]byte(secret))

	return
}
