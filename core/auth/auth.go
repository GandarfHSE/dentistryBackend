package auth

import (
	"crypto/rsa"
	"os"

	"github.com/GandarfHSE/dentistryBackend/utils/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

type AuthHandlers struct {
	// TODO: use postresql
	jwtPrivate *rsa.PrivateKey
	jwtPublic  *rsa.PublicKey
}

var authHandlers *AuthHandlers

func GetAuthHandlers() (*AuthHandlers, error) {
	if authHandlers == nil {
		log.Warn().Msg("GetAuthHandlers from nil: trying to load authHandlers...")
		err := LoadAuthHandlers()
		if err != nil {
			log.Error().Err(err).Msg("GetAuthHandlers from nil: failed to load authHandlers!")
			return nil, err
		}
	}
	return authHandlers, nil
}

func LoadAuthHandlers() error {
	authHandlers = nil

	jwtPrivatePath, err := config.GetAbsPrivatePath()
	if err != nil {
		log.Error().Err(err).Msg("Can't get private path from config...")
		return err
	}
	jwtPublicPath, err := config.GetAbsPublicPath()
	if err != nil {
		log.Error().Err(err).Msg("Can't get public path from config...")
		return err
	}

	private, err := os.ReadFile(jwtPrivatePath)
	if err != nil {
		log.Error().Err(err).Msg("Can't read jwtprivateFile!")
		return err
	}
	public, err := os.ReadFile(jwtPublicPath)
	if err != nil {
		log.Error().Err(err).Msg("Can't read jwtPublicFile!")
		return err
	}

	jwtPrivate, err := jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		log.Error().Err(err).Msg("Can't parse jwtprivateFile!")
		return err
	}
	jwtPublic, err := jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		log.Error().Err(err).Msg("Can't parse jwtPublicFile!")
		return err
	}

	authHandlers = &AuthHandlers{
		jwtPrivate: jwtPrivate,
		jwtPublic:  jwtPublic,
	}
	log.Info().Msg("AuthHandlers have been set successfully!")
	return nil
}
