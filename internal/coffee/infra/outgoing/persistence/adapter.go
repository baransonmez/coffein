package persistence

import (
	"context"
	"github.com/baransonmez/coffein/internal/coffee/business/domain"
	"github.com/google/uuid"
	"time"
)

type store struct {
}

type Bean struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Roaster     string    `db:"roaster"`
	Origin      string    `db:"origin"`
	Price       int       `db:"price"`
	RoastDate   time.Time `db:"roast_date"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func (dbPrd *Bean) ToBean() *domain.Bean {
	id, _ := uuid.Parse(dbPrd.ID)
	pu := domain.Bean{
		ID:          id,
		Name:        dbPrd.Name,
		Roaster:     dbPrd.Roaster,
		Origin:      dbPrd.Origin,
		RoastDate:   dbPrd.RoastDate,
		Price:       dbPrd.Price,
		DateCreated: dbPrd.DateCreated,
		DateUpdated: dbPrd.DateUpdated,
	}
	return &pu
}

func NewStore() (store, error) {

	return store{}, nil
}

func (s store) Create(_ context.Context, bean domain.Bean) error {

	return nil
}

func (s store) Get(id uuid.UUID) (*domain.Bean, error) {
	var bean Bean

	return bean.ToBean(), nil
}
