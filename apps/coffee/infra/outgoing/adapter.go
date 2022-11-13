package outgoing

import (
	"context"
	"github.com/baransonmez/coffein/internal/coffee/business/domain"
	"github.com/google/uuid"
)

type Store interface {
	Get(id string) (*Bean, error)
	Create(ctx context.Context, e *Bean) (ID error)
}
type Adapter struct {
	store Store
}

func NewBeanAdapter(store Store) (Adapter, error) {
	return Adapter{store: store}, nil
}

func (s Adapter) Create(ctx context.Context, bean domain.Bean) error {
	beanForDB := &Bean{
		ID:          bean.ID.String(),
		Name:        bean.Name,
		Roaster:     bean.Roaster,
		Origin:      bean.Origin,
		Price:       bean.Price,
		RoastDate:   bean.RoastDate,
		DateCreated: bean.DateCreated,
		DateUpdated: bean.DateUpdated,
	}

	return s.store.Create(ctx, beanForDB)
}

func (s Adapter) Get(id uuid.UUID) (*domain.Bean, error) {
	bean, err := s.store.Get(id.String())
	if err != nil {
		return nil, err
	}
	return bean.ToBean(), nil
}
