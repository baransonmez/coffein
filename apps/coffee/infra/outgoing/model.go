package outgoing

import (
	"github.com/baransonmez/coffein/internal/coffee/business/domain"
	"github.com/google/uuid"
	"time"
)

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
