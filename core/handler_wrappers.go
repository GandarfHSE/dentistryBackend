package core

import (
	"net/http"

	"github.com/ansel1/merry/v2"
	"github.com/rs/zerolog/log"
)

/*
* http.Request --decoder--> Request --handler--> Response --encoder --> http.ResponseWriter
 */
func handlerWrapper[Request any, Response any](
	decoder func(*http.Request, *Request) error,
	handler func(Request) (*Response, error),
	encoder func(http.ResponseWriter, *Response) error,
) func(http.ResponseWriter, *http.Request) {

	wrapper := func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Processing HTTP Request...")

		var req Request
		err := decoder(r, &req)
		if err != nil {
			log.Error().Err(err).Msg("Error while decoding request!")
			errorEncoder(w, err)
			w.WriteHeader(merry.HTTPCode(err))
			return
		}

		respPtr, err := handler(req)
		if err != nil {
			log.Error().Err(err).Msg("Error while handling request!")
			errorEncoder(w, err)
			w.WriteHeader(merry.HTTPCode(err))
			return
		}

		err = encoder(w, respPtr)
		if err != nil {
			log.Error().Err(err).Msg("Error while encoding response!")
			errorEncoder(w, err)
			w.WriteHeader(merry.HTTPCode(err))
			return
		}

		w.WriteHeader(200)
		log.Info().Msg("Request has been processed successfully!")
	}

	return wrapper
}

func jsonHandlerWrapper[Request any, Response any](
	handler func(Request) (*Response, error),
) func(http.ResponseWriter, *http.Request) {
	decoder := jsonDecoder[Request]
	encoder := jsonEncoder[Response]
	return handlerWrapper(decoder, handler, encoder)
}
