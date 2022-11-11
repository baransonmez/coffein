package usecases

import (
	"errors"
	"github.com/baransonmez/coffein/internal/user/business/domain"
	"github.com/google/uuid"
	"time"
)

type NewUser struct {
	Name string `json:"name"`
}

func (u *NewUser) toDomainModel() domain.User {
	user := domain.User{
		ID:          uuid.New(),
		Name:        u.Name,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	return user
}

func (u *NewUser) Validate() error {
	if u.Name == "" {
		//return &common.CannotBeSmallerError{Field: "name", Limit: 2}
		return errors.New("name cannot be empty")
	}
	return nil
}
