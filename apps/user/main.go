package main

import (
	"github.com/baransonmez/coffein/internal/user/business/usecases"
	"github.com/baransonmez/coffein/internal/user/infra/incoming/web"
	"github.com/baransonmez/coffein/internal/user/infra/outgoing"
	"github.com/baransonmez/coffein/internal/user/infra/outgoing/persistence"
	kitweb "github.com/baransonmez/coffein/kit/web"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	inmem := persistence.NewInMem()
	adapter := outgoing.NewUserAdapter(inmem)
	commandService := usecases.NewCommandService(adapter)
	queryService := usecases.NewQueryService(adapter)
	user, err := commandService.CreateNewUser(nil, usecases.NewUser{Name: "Baran"})
	if err != nil {
		print(err)
	}
	print(user.String())

	userAPI := web.Handlers{CommandService: commandService, QueryService: queryService}
	handler := routes(userAPI)

	servPort := ":8083"
	log.Printf("starting server on %s\n", servPort)

	srv := &http.Server{
		Addr:         servPort,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 45 * time.Second,
	}
	err = srv.ListenAndServe()

}

func routes(userAPI web.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/user", kitweb.Handle(userAPI.Create))
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", kitweb.Handle(userAPI.Get))
	return router
}
