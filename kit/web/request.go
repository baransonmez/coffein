package web

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}

func ReadIDParam(r *http.Request) (uuid.UUID, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := stringToID(params.ByName("id"))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func stringToID(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	return id, err
}
