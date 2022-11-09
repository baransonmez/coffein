package web

import (
	"fmt"
	"github.com/baransonmez/coffein/internal/coffee/business/usecases"
	"github.com/baransonmez/coffein/kit/web"
	"net/http"
)

type Handlers struct {
	CommandService *usecases.CommandService
	QueryService   *usecases.QueryService
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) error {
	var ncb usecases.NewCoffeeBean
	if err := web.Decode(r, &ncb); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	ctx := r.Context()

	prod, err := h.CommandService.CreateCoffeeBean(ctx, ncb)
	if err != nil {
		return fmt.Errorf("creating new coffee bean, ncb[%+v]: %w", ncb, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}

func (h Handlers) GetBean(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	id, err := web.ReadIDParam(r)
	if err != nil {
		return err
	}
	bean, err := h.QueryService.GetBean(ctx, id)
	if err != nil {
		return fmt.Errorf("getting coffee bean, id[%+v]: %w", id, err)
	}

	return web.Respond(w, bean, http.StatusOK)
}
