package psw

import "golang.org/x/crypto/bcrypt"

func Verify(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	return err == nil
}
