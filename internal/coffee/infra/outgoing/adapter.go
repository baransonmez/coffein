package outgoing

import (
	"context"
	"github.com/baransonmez/coffein/internal/coffee/business/domain"
	"github.com/google/uuid"
)

type Store struct {
}

func NewStore() (Store, error) {

	return Store{}, nil
}

func (s Store) Create(_ context.Context, bean domain.Bean) error {

	return nil
}

func (s Store) Get(id uuid.UUID) (*domain.Bean, error) {
	var bean Bean

	return bean.ToBean(), nil
}
