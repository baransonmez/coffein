package web

import (
	"fmt"
	"github.com/baransonmez/coffein/internal/recipe/business/usecases"
	"github.com/baransonmez/coffein/kit/web"
	"net/http"
)

type Handlers struct {
	CommandService *usecases.CommandService
	QueryService   *usecases.QueryService
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) error {
	var nr usecases.NewRecipe
	if err := web.Decode(r, &nr); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	ctx := r.Context()

	prod, err := h.CommandService.CreateNewRecipe(ctx, nr)
	if err != nil {
		return fmt.Errorf("creating new coffee bean, nr[%+v]: %w", nr, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	recipeUUID, err := web.ReadIDParam(r)
	if err != nil {
		return fmt.Errorf("unable to decode id")
	}

	prod, err := h.QueryService.GetRecipe(ctx, recipeUUID)
	if err != nil {
		return fmt.Errorf("getting recipe, recipeUUID[%+v]: %w", recipeUUID, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}
