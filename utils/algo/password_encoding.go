package algo

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateEncodedPassword(password string) string {
	hash := md5.Sum([]byte(password + "kek" + password))
	return hex.EncodeToString(hash[:])
}
