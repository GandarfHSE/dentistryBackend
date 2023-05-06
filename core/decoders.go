package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/GandarfHSE/dentistryBackend/core/auth"
	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	img "github.com/GandarfHSE/dentistryBackend/utils/image"
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

func emptyDecoder[Request any](r *http.Request, v *Request) error {
	return nil
}

func imageDecoder(r *http.Request) (*img.Image, error) {
	err := r.ParseMultipartForm(1 << 25)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing multipart form!")
		return nil, err
	}

	file, _, err := r.FormFile("image_data")
	if err == http.ErrMissingFile {
		log.Info().Msg("No image data was provided!")
		return nil, err
	}
	if err != nil {
		log.Error().Err(err).Msg("Error occured while parsing image data!")
		return nil, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while copying image data!")
		return nil, err
	}
	data := buf.Bytes()

	if len(data) == 0 {
		err = errors.New("Image body is empty!")
		log.Error().Err(err).Msg("Image body is empty!")
		return nil, err
	}

	ext := r.FormValue("image_extension")
	if ext == "" {
		err = errors.New("File extension is empty!")
		log.Error().Err(err).Msg("File extension is empty!")
		return nil, err
	}

	log.Info().Msg("Successfully parsed an image!")
	return &img.Image{Ext: ext, Data: data}, nil
}
