package recipe

import (
	"context"
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/google/uuid"
)

type Store interface {
	Get(id string) (*Recipe, error)
	Create(ctx context.Context, e Recipe) (ID error)
}
type Adapter struct {
	store Store
}

func NewRecipeAdapter(store Store) (Adapter, error) {
	return Adapter{store: store}, nil
}

func (r Adapter) Create(ctx context.Context, recipe domain.Recipe) error {
	recipeForDb := &Recipe{
		ID:          recipe.ID.String(),
		Description: recipe.Description,
		UserID:      recipe.UserID.String(),
		BeanID:      recipe.BeanID.String(),
		Steps:       StepsFromDomainModel(recipe.Steps),
		DateCreated: recipe.DateCreated,
		DateUpdated: recipe.DateUpdated,
	}
	return r.store.Create(ctx, *recipeForDb)
}

func (r Adapter) Get(id uuid.UUID) (*domain.Recipe, error) {
	recipe, err := r.store.Get(id.String())
	if err != nil {
		return nil, err
	}
	return recipe.ToRecipe(), nil
}
