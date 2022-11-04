package usecases

import (
	"context"
	"github.com/baransonmez/coffein/internal/coffee/business/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Get(id uuid.UUID) (*domain.Bean, error)
	Create(ctx context.Context, b domain.Bean) (ID error)
}
