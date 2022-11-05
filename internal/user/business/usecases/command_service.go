package usecases

import (
	"context"
	"fmt"
	"github.com/baransonmez/coffein/internal/user/business/domain"
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (c Service) CreateNewUser(ctx context.Context, np NewUser) (uuid.UUID, error) {
	err := np.Validate()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}
	user := np.toDomainModel()
	if err := c.repository.Create(ctx, user); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return user.ID, nil
}

func (c Service) GetUser(_ context.Context, id uuid.UUID) (*domain.User, error) {
	bean, err := c.repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return bean, nil
}
