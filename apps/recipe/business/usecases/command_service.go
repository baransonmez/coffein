package usecases

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type CommandService struct {
	recipeRepository RecipeRepository
	userRepository   UserRepository
}

func NewCommandService(r RecipeRepository, u UserRepository) *CommandService {
	return &CommandService{
		recipeRepository: r,
		userRepository:   u,
	}
}

func (c CommandService) CreateNewRecipe(ctx context.Context, np NewRecipe) (uuid.UUID, error) {
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
