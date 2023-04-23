package auth

import (
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/ansel1/merry"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

// TODO: username -> login
func (h *AuthHandlers) CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": username,
	})
	return token.SignedString(h.jwtPrivate)
}

func (h *AuthHandlers) ParseToken(tokenString string) (*cookie.Cookie, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Error().Msg("Wrong token signing method!")
			return nil, merry.New("Wrong token signing method!").WithHTTPCode(405)
		}
		return h.jwtPublic, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"]
		if !ok {
			return nil, merry.New("JWT claims do not have username field").WithHTTPCode(401)
		}
		usernameString, ok := username.(string)
		if !ok {
			return nil, merry.New("Can't cast username to string").WithHTTPCode(400)
		}

		cookie := cookie.Cookie{Username: usernameString}
		return &cookie, nil
	} else {
		return nil, merry.New("JWT invalid").WithHTTPCode(401)
	}
}
