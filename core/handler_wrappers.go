package core

import (
	"net/http"

	"github.com/GandarfHSE/dentistryBackend/utils/cookie"
	img "github.com/GandarfHSE/dentistryBackend/utils/image"
	"github.com/ansel1/merry"
	"github.com/rs/zerolog/log"
)

/*
* http.Request --decoder--> Request --handler--> Response --encoder --> http.ResponseWriter
 */
func handlerWrapper[Request any, Response any](
	decoder func(*http.Request, *Request) error,
	handler func(Request, *cookie.Cookie) (*Response, merry.Error),
	encoder func(http.ResponseWriter, *Response) error,
) func(http.ResponseWriter, *http.Request) {

	wrapper := func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Processing HTTP Request...")

		var req Request
		err := decoder(r, &req)
		if err != nil {
			log.Error().Err(err).Msg("Error while decoding request!")
			w.WriteHeader(merry.HTTPCode(err))
			errorEncoder(w, err)
			return
		}

		cookie, err := cookieDecoder(r)
		respPtr, err := handler(req, cookie)
		if err != nil {
			log.Error().Err(err).Msg("Error while handling request!")
			w.WriteHeader(merry.HTTPCode(err))
			errorEncoder(w, err)
			return
		}

		err = encoder(w, respPtr)
		if err != nil {
			log.Error().Err(err).Msg("Error while encoding response!")
			w.WriteHeader(merry.HTTPCode(err))
			errorEncoder(w, err)
			return
		}

		log.Info().Msg("Request has been processed successfully!")
	}

	return wrapper
}

func jsonHandlerWrapper[Request any, Response any](
	handler func(Request, *cookie.Cookie) (*Response, merry.Error),
) func(http.ResponseWriter, *http.Request) {
	decoder := jsonDecoder[Request]
	encoder := jsonEncoder[Response]
	return handlerWrapper(decoder, handler, encoder)
}

func noBodyHandlerWrapper[Request any, Response any](
	handler func(Request, *cookie.Cookie) (*Response, merry.Error),
) func(http.ResponseWriter, *http.Request) {
	decoder := emptyDecoder[Request]
	encoder := jsonEncoder[Response]
	return handlerWrapper(decoder, handler, encoder)
}

func imageUploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Processing image upload request...")

	image, err := imageDecoder(r)
	if err != nil {
		log.Error().Err(err).Msg("Error while decoding image!")
		w.WriteHeader(500)
		errorEncoder(w, err)
		return
	}

	imgSrc, err := img.UploadImage(image)
	if err != nil {
		log.Error().Err(err).Msg("Error while uploading image!")
		w.WriteHeader(500)
		errorEncoder(w, err)
		return
	}

	imgResponse := &img.ImageUploadResponse{ImageSource: imgSrc}
	err = jsonEncoder(w, imgResponse)
	if err != nil {
		log.Error().Err(err).Msg("Error while encoding image response!")
		w.WriteHeader(500)
		errorEncoder(w, err)
		return
	}

	log.Info().Msg("Image upload request has been processed successfully!")
}
