package main

import (
	"github.com/baransonmez/coffein/internal/user/business/usecases"
	"github.com/baransonmez/coffein/internal/user/infra/outgoing/persistence"
)

func main() {
	store, _ := persistence.NewStore()
	commandService := usecases.NewCommandService(store)
	_ = usecases.NewQueryService(store)
	user, err := commandService.CreateNewUser(nil, usecases.NewUser{Name: "Baran"})
	if err != nil {
		print(err)
	}
	print(user.String())

}
