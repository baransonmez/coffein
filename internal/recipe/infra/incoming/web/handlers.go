package web

import (
	"fmt"
	"github.com/baransonmez/coffein/internal/recipe/business/usecases"
	"github.com/baransonmez/coffein/kit/web"
	"net/http"
)

type Handlers struct {
	RecipeService *usecases.Service
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var nr usecases.NewRecipe
	if err := web.Decode(r, &nr); err != nil {
		msg := fmt.Errorf("unable to decode payload: %w", err)
		web.Respond(w, web.ErrorResponse{Error: msg.Error()}, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	prod, err := h.RecipeService.CreateNewRecipe(ctx, nr)
	if err != nil {
		msg := fmt.Errorf("creating new coffee bean, nr[%+v]: %w", nr, err)
		web.Respond(w, web.ErrorResponse{Error: msg.Error()}, http.StatusBadRequest)
		return
	}

	web.Respond(w, prod, http.StatusCreated)
}

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	recipeUUID, err := web.ReadIDParam(r)
	if err != nil {
		msg := fmt.Errorf("unable to decode id")
		web.Respond(w, web.ErrorResponse{Error: msg.Error()}, http.StatusBadRequest)
		return
	}

	prod, err := h.RecipeService.GetRecipe(ctx, recipeUUID)
	if err != nil {
		msg := fmt.Errorf("getting recipe, recipeUUID[%+v]: %w", recipeUUID, err)
		web.Respond(w, web.ErrorResponse{Error: msg.Error()}, http.StatusBadRequest)
		return
	}

	web.Respond(w, prod, http.StatusCreated)
}
