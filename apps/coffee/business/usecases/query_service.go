package usecases

import (
	"context"
	"fmt"
	"github.com/baransonmez/coffein/internal/coffee/business/domain"
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

func (q QueryService) GetBean(_ context.Context, id uuid.UUID) (*domain.Bean, error) {
	bean, err := q.repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return bean, nil
}
