package main

import (
	"encoding/json"
	"fmt"
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/baransonmez/coffein/internal/recipe/business/usecases"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing/persistence"
	"github.com/google/uuid"
)

func main() {
	mem := persistence.NewInMem()
	recipeAdapter, _ := outgoing.NewRecipeAdapter(mem)
	userAdapter, _ := outgoing.NewUserAdapter()
	service := usecases.NewService(recipeAdapter, userAdapter)
	recipeID, err := service.CreateNewRecipe(nil, usecases.NewRecipe{
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

	print(recipeID.String())

	recipeFromDb, err := mem.Get(recipeID.String())
	if err != nil {
		print(err)
	}
	fmt.Println(prettyPrint(recipeFromDb))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
