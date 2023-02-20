package core

import (
	"crypto/md5"

	"github.com/ansel1/merry"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

func (h *AuthHandlers) GenerateEncodedPassword(password string) [16]byte {
	hh := md5.New()
	var res [16]byte
	copy(res[:], hh.Sum([]byte(password + "kek" + password))[0:16])
	return res
}

func (h *AuthHandlers) CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": username,
	})
	return token.SignedString(h.jwtPrivate)
}

// TODO: return struct instead of string
func (h *AuthHandlers) ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Error().Msg("Wrong token signing method!")
			return nil, merry.New("Wrong token signing method!").WithHTTPCode(405)
		}
		return h.jwtPublic, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"]
		if !ok {
			return "", merry.New("JWT claims do not have username field").WithHTTPCode(401)
		}
		usernameString, ok := username.(string)
		if !ok {
			return "", merry.New("Can't cast username to string").WithHTTPCode(400)
		}
		return usernameString, nil
	} else {
		return "", merry.New("JWT invalid").WithHTTPCode(401)
	}
}
