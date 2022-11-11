package usecases

import (
	"context"
	"fmt"
	"github.com/baransonmez/coffein/internal/user/business/domain"
	"github.com/google/uuid"
)

type QueryService struct {
	repository Repository
}

func NewQueryService(r Repository) *QueryService {
	return &QueryService{
		repository: r,
	}
}

func (s QueryService) GetUser(_ context.Context, id uuid.UUID) (*domain.User, error) {
	bean, err := s.repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}
	return bean, nil
}
