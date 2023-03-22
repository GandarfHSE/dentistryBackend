package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/GandarfHSE/dentistryBackend/core/auth"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	"github.com/rs/zerolog/log"
)

func cookieDecoder(r *http.Request) (*cookie.Cookie, error) {
	raw_cookie, err := r.Cookie("jwt")
	if err != nil {
		log.Warn().Msg("No JWT in cookie!")
		return nil, err
	}

	authHandlers, err := auth.GetAuthHandlers()
	cookie, err := authHandlers.ParseToken(raw_cookie.Value)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get cookie from JWT!")
		return nil, err
	}

	return cookie, nil
}

func jsonDecoder[Request any](r *http.Request, v *Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error().Err(err).Msg("Error occured while reading request body")
		return err
	}

	err = json.Unmarshal(reqBody, v)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing json")
		return err
	}

	log.Info().Msg("Successfylly got JSON request")
	return nil
}
