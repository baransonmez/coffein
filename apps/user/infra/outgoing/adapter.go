package outgoing

import (
	"context"
	"github.com/baransonmez/coffein/internal/user/business/domain"
	"github.com/google/uuid"
)

type Store interface {
	Get(id string) (*User, error)
	Create(ctx context.Context, e *User) (ID error)
}
type Adapter struct {
	store Store
}

func NewUserAdapter(store Store) Adapter {
	return Adapter{store: store}
}

func (a Adapter) Create(ctx context.Context, user domain.User) error {
	userForDB := &User{
		ID:          user.ID.String(),
		Name:        user.Name,
		DateCreated: user.DateCreated,
		DateUpdated: user.DateUpdated,
	}

	return a.store.Create(ctx, userForDB)
}

func (a Adapter) Get(id uuid.UUID) (*domain.User, error) {
	user, err := a.store.Get(id.String())
	if err != nil {
		return nil, err
	}

	return user.ToUser(), nil
}
