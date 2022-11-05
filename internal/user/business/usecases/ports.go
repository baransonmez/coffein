package usecases

import (
	"context"
	"github.com/baransonmez/coffein/internal/user/business/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Get(id uuid.UUID) (*domain.User, error)
	Create(ctx context.Context, b domain.User) (ID error)
}
