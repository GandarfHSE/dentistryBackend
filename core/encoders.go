package core

import (
	"encoding/json"
	"net/http"

	"github.com/ansel1/merry/v2"
)

func jsonEncoder[Response any](w http.ResponseWriter, r *Response) error {
	err := json.NewEncoder(w).Encode(*r)
	if err != nil {
		return err
	}
	return nil
}

func errorEncoder(w http.ResponseWriter, err error) error {
	resp := make(map[string]string)
	resp["err"] = merry.Details(err)

	e := json.NewEncoder(w).Encode(resp)
	if e != nil {
		return e
	}
	return nil
}
