package psw

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(result), err
}
