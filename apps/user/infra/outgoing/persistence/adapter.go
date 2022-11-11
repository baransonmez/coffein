package persistence

import (
	"context"
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

type store struct {
}

func NewStore() (store, error) {

	return store{}, nil
}

func (s store) Create(_ context.Context, user domain.User) error {

	return nil
}

func (s store) Get(id uuid.UUID) (*domain.User, error) {
	var user User

	return user.ToUser(), nil
}
