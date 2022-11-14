package main

import (
	"github.com/baransonmez/coffein/internal/recipe/business/usecases"
	"github.com/baransonmez/coffein/internal/recipe/infra/incoming/web"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing/recipe"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing/recipe/persistence"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing/user"
	kitweb "github.com/baransonmez/coffein/kit/web"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	mem := persistence.NewInMem()
	recipeAdapter, _ := recipe.NewRecipeAdapter(mem)
	userAdapter, _ := user.NewUserAdapter()
	commandService := usecases.NewCommandService(recipeAdapter, userAdapter)
	queryService := usecases.NewQueryService(recipeAdapter)

	recipeAPI := web.Handlers{CommandService: commandService, QueryService: queryService}
	handler := routes(recipeAPI)

	servPort := ":8089"
	log.Printf("starting server on %s\n", servPort)

	srv := &http.Server{
		Addr:         servPort,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 45 * time.Second,
	}
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func routes(recipeAPI web.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/recipe", kitweb.Handle(recipeAPI.Create))
	router.HandlerFunc(http.MethodGet, "/v1/recipe/:id", kitweb.Handle(recipeAPI.Get))
	return router
}
