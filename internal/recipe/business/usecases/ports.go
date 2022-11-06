package usecases

import (
	"context"
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/google/uuid"
)

type RecipeRepository interface {
	Get(id uuid.UUID) (*domain.Recipe, error)
	Create(ctx context.Context, e domain.Recipe) (ID error)
}

type UserRepository interface {
	GetMembershipType(ctx context.Context, userId string) (string, error)
}
