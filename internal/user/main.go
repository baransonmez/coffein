package main

import (
	"github.com/baransonmez/coffein/internal/user/business/usecases"
	"github.com/baransonmez/coffein/internal/user/infra/outgoing/persistence"
)

func main() {
	store, _ := persistence.NewStore()
	service := usecases.NewService(store)
	user, err := service.CreateNewUser(nil, usecases.NewUser{Name: "Baran"})
	if err != nil {
		print(err)
	}
	print(user.String())

}
