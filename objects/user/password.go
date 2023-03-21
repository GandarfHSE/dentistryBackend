package user

import (
	"crypto/md5"
)

func generateEncodedPassword(password string) string {
	hh := md5.New()
	return string(hh.Sum([]byte(password + "kek" + password)))
}
