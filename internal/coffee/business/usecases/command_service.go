package usecases

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type CommandService struct {
	repository Repository
}

func NewService(r Repository) *CommandService {
	return &CommandService{
		repository: r,
	}
}

func (c CommandService) CreateCoffeeBean(ctx context.Context, np NewCoffeeBean) (uuid.UUID, error) {
	err := np.Validate()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}
	coffeeBean := np.toDomainModel()

	if err := c.repository.Create(ctx, *coffeeBean); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return coffeeBean.ID, nil
}
