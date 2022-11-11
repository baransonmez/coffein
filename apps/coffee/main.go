package main

import (
	"github.com/baransonmez/coffein/internal/coffee/business/usecases"
	"github.com/baransonmez/coffein/internal/coffee/infra/incoming/web"
	"github.com/baransonmez/coffein/internal/coffee/infra/outgoing"
	kitweb "github.com/baransonmez/coffein/kit/web"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	store, _ := outgoing.NewStore()
	commandService := usecases.NewCommandService(store)
	queryService := usecases.NewQueryService(store)
	beanId, err := commandService.CreateCoffeeBean(nil, usecases.NewCoffeeBean{
		Name:      "Yirgaciffe",
		Roaster:   "Montag",
		Origin:    "Etiopia",
		Price:     23,
		RoastDate: time.Now().AddDate(2, 3, 4),
	})

	if err != nil {
		print(err)
	}
	print(beanId.String())

	recipeAPI := web.Handlers{CommandService: commandService, QueryService: queryService}
	handler := routes(recipeAPI)

	servPort := ":8086"
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

func routes(recipeAPI web.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/bean", kitweb.Handle(recipeAPI.Create))
	router.HandlerFunc(http.MethodGet, "/v1/bean/:id", kitweb.Handle(recipeAPI.GetBean))
	return router
}
