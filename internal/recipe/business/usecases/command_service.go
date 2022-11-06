package usecases

import (
	"context"
	"fmt"
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/google/uuid"
)

type Service struct {
	recipeRepository RecipeRepository
	userRepository   UserRepository
}

func NewService(r RecipeRepository, u UserRepository) *Service {
	return &Service{
		recipeRepository: r,
		userRepository:   u,
	}
}

func (c Service) CreateNewRecipe(ctx context.Context, np NewRecipe) (uuid.UUID, error) {
	err := np.Validate()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}
	membershipType, err := c.userRepository.GetMembershipType(ctx, np.UserID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}
	if membershipType == "" {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}
	recipe := np.toDomainModel()
	if err := c.recipeRepository.Create(ctx, recipe); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return recipe.ID, nil
}

func (c Service) GetRecipe(_ context.Context, id uuid.UUID) (*domain.Recipe, error) {
	recipe, err := c.recipeRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return recipe, nil
}
