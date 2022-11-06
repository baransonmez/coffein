package main

import (
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/baransonmez/coffein/internal/recipe/business/usecases"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing"
	"github.com/google/uuid"
)

func main() {
	recipeAdapter, _ := outgoing.NewRecipeAdapter()
	userAdapter, _ := outgoing.NewUserAdapter()
	service := usecases.NewService(recipeAdapter, userAdapter)
	recipe, err := service.CreateNewRecipe(nil, usecases.NewRecipe{
		UserID:      uuid.New().String(),
		BeanID:      uuid.New().String(),
		Description: "30 seconds blooming",
		Steps: []domain.Step{
			{
				Description:       "blooming",
				DurationInSeconds: 24,
			},
			{
				Description:       "brewing",
				DurationInSeconds: 76,
			}},
	})
	if err != nil {
		print(err)
	}

	print(recipe.String())
}
