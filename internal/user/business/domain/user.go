package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID
	Name        string
	DateCreated time.Time
	DateUpdated time.Time
}
