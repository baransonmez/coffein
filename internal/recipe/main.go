package main

import (
	"encoding/json"
	"fmt"
	"github.com/baransonmez/coffein/internal/recipe/business/domain"
	"github.com/baransonmez/coffein/internal/recipe/business/usecases"
	"github.com/baransonmez/coffein/internal/recipe/infra/incoming/web"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing/persistence"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
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
	recipeAPI := web.Handlers{RecipeService: service}
	handler := routes(recipeAPI)

	servPort := ":8089"
	log.Printf("starting server on %s\n", servPort)

	srv := &http.Server{
		Addr:         servPort,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 45 * time.Second,
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func routes(recipeAPI web.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/recipe", recipeAPI.Create)
	router.HandlerFunc(http.MethodGet, "/v1/recipe/:id", recipeAPI.Get)
	return router
}
