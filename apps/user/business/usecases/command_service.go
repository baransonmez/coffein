package usecases

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type CommandService struct {
	repository Repository
}

func NewCommandService(r Repository) *CommandService {
	return &CommandService{
		repository: r,
	}
}

func (s CommandService) CreateNewUser(ctx context.Context, np NewUser) (uuid.UUID, error) {
	err := np.Validate()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}
	user := np.toDomainModel()
	if err := s.repository.Create(ctx, user); err != nil {
		return uuid.UUID{}, fmt.Errorf("create: %w", err)
	}

	return user.ID, nil
}
