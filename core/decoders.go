package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

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
