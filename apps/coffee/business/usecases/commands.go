package usecases

import (
	"errors"
	"github.com/baransonmez/coffein/internal/coffee/business/domain"
	"github.com/google/uuid"
	"time"
)

var NameCannotBeEmptyError = errors.New("name cannot be empty")
var OriginCannotBeEmptyError = errors.New("origin cannot be empty")
var RoasterCannotBeEmptyError = errors.New("roaster cannot be empty")
var PriceCannotBeLowerThan1Error = errors.New("price cannot be empty")

type NewCoffeeBean struct {
	Name      string    `json:"name"`
	Roaster   string    `json:"roaster"`
	Origin    string    `json:"origin"`
	Price     int       `json:"price"`
	RoastDate time.Time `json:"roast_created"`
}

func (c *NewCoffeeBean) toDomainModel() *domain.Bean {
	coffeeBean := domain.Bean{
		ID:          uuid.New(),
		Name:        c.Name,
		Roaster:     c.Roaster,
		Origin:      c.Origin,
		RoastDate:   c.RoastDate,
		Price:       c.Price,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	return &coffeeBean
}

func (c *NewCoffeeBean) Validate() error {
	if c.Name == "" {
		return NameCannotBeEmptyError
	}
	if c.Origin == "" {
		return OriginCannotBeEmptyError
	}
	if c.Roaster == "" {
		return RoasterCannotBeEmptyError
	}
	if c.Price < 1 {
		return PriceCannotBeLowerThan1Error
	}

	return nil
}
