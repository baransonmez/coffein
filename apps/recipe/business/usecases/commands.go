package usecases

import (
	"errors"
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/google/uuid"
	"time"
)

type NewRecipe struct {
	UserID      string        `json:"user_id"`
	BeanID      string        `json:"bean_id"`
	Description string        `json:"description"`
	Steps       []domain.Step `json:"steps"`
}

func (r *NewRecipe) toDomainModel() domain.Recipe {
	userId, _ := uuid.Parse(r.UserID)
	beanId, _ := uuid.Parse(r.BeanID)
	recipe := domain.Recipe{
		ID:          uuid.New(),
		Description: r.Description,
		BeanID:      beanId,
		UserID:      userId,
		Steps:       r.Steps,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	return recipe
}

func (r *NewRecipe) Validate() error {
	if r.UserID == "" {
		return errors.New("userId cannot be empty")
	}
	if r.BeanID == "" {
		return errors.New("coffeeID cannot be empty")
	}
	if r.Description == "" {
		return errors.New("description cannot be empty")
	}
	if len(r.Steps) < 1 {
		return errors.New("steps length must be greater than 1")
	}

	return nil
}
