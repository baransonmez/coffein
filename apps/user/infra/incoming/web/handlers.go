package web

import (
	"fmt"
	"github.com/baransonmez/coffein/internal/user/business/usecases"
	"github.com/baransonmez/coffein/kit/web"
	"net/http"
)

type Handlers struct {
	CommandService *usecases.CommandService
	QueryService   *usecases.QueryService
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) error {
	var nu usecases.NewUser
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	ctx := r.Context()

	prod, err := h.CommandService.CreateNewUser(ctx, nu)
	if err != nil {
		return fmt.Errorf("creating new user, nu[%+v]: %w", nu, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	userUUID, err := web.ReadIDParam(r)
	if err != nil {
		return err
	}

	prod, err := h.QueryService.GetUser(ctx, userUUID)
	if err != nil {
		return fmt.Errorf("getting recipe, userUUID[%+v]: %w", userUUID, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}
