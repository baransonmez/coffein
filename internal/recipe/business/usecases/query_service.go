package usecases

import (
	"context"
	"fmt"
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/google/uuid"
)

type QueryService struct {
	recipeRepository RecipeRepository
}

func NewQueryService(r RecipeRepository) *QueryService {
	return &QueryService{
		recipeRepository: r,
	}
}

func (q QueryService) GetRecipe(_ context.Context, id uuid.UUID) (*domain.Recipe, error) {
	recipe, err := q.recipeRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return recipe, nil
}
