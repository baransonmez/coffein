package main

import (
	"github.com/baransonmez/coffein/internal/coffee/business/usecases"
	"github.com/baransonmez/coffein/internal/coffee/infra/incoming/web"
	"github.com/baransonmez/coffein/internal/coffee/infra/outgoing"
	"github.com/baransonmez/coffein/internal/coffee/infra/outgoing/persistence"
	kitweb "github.com/baransonmez/coffein/kit/web"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	inmemStore := persistence.NewInMem()
	outgoingAdapter, _ := outgoing.NewBeanAdapter(inmemStore)
	commandService := usecases.NewCommandService(outgoingAdapter)
	queryService := usecases.NewQueryService(outgoingAdapter)

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
	err := srv.ListenAndServe()
	log.Fatal(err)

}

func routes(beanAPI web.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/bean", kitweb.Handle(beanAPI.Create))
	router.HandlerFunc(http.MethodGet, "/v1/bean/:id", kitweb.Handle(beanAPI.GetBean))
	return router
}
