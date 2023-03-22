package algo

import (
	"crypto/md5"
)

func GenerateEncodedPassword(password string) string {
	hh := md5.New()
	return string(hh.Sum([]byte(password + "kek" + password)))
}
