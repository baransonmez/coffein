package main

import (
	"github.com/baransonmez/coffein/internal/coffee/business/usecases"
	"github.com/baransonmez/coffein/internal/coffee/infra/outgoing/persistence"
	"time"
)

func main() {
	store, _ := persistence.NewStore()
	service := usecases.NewService(store)
	beanId, err := service.CreateCoffeeBean(nil, usecases.NewCoffeeBean{
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
}
