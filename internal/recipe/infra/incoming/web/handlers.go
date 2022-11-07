package web

import (
	"encoding/json"
	"fmt"
	"github.com/baransonmez/coffein/internal/recipe/business/usecases"
	"github.com/google/uuid"

	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handlers struct {
	RecipeService *usecases.Service
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var nr usecases.NewRecipe
	if err := decode(r, &nr); err != nil {
		msg := fmt.Errorf("unable to decode payload: %w", err)
		Respond(w, ErrorResponse{Error: msg.Error()}, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	prod, err := h.RecipeService.CreateNewRecipe(ctx, nr)
	if err != nil {
		msg := fmt.Errorf("creating new coffee bean, nr[%+v]: %w", nr, err)
		Respond(w, ErrorResponse{Error: msg.Error()}, http.StatusBadRequest)
		return
	}

	Respond(w, prod, http.StatusCreated)
}

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	recipeUUID, err := readIDParam(r)
	if err != nil {
		msg := fmt.Errorf("unable to decode id")
		Respond(w, ErrorResponse{Error: msg.Error()}, http.StatusBadRequest)
		return
	}

	prod, err := h.RecipeService.GetRecipe(ctx, recipeUUID)
	if err != nil {
		msg := fmt.Errorf("getting recipe, recipeUUID[%+v]: %w", recipeUUID, err)
		Respond(w, ErrorResponse{Error: msg.Error()}, http.StatusBadRequest)
		return
	}

	Respond(w, prod, http.StatusCreated)
}

func decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}

func readIDParam(r *http.Request) (uuid.UUID, error) {
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

// Respond converts a Go value to JSON and sends it to the client.
func Respond(w http.ResponseWriter, data any, statusCode int) {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
	}

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		println(err)
	}

	jsonData = append(jsonData, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		println(err)
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}
