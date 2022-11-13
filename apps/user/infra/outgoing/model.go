package outgoing

import (
	"github.com/baransonmez/coffein/internal/user/business/domain"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func (dbPrd *User) ToUser() *domain.User {
	id, _ := uuid.Parse(dbPrd.ID)
	u := domain.User{
		ID:          id,
		Name:        dbPrd.Name,
		DateCreated: dbPrd.DateCreated,
		DateUpdated: dbPrd.DateUpdated,
	}
	return &u
}
