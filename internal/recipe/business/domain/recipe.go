package domain

import (
	"github.com/google/uuid"
	"time"
)

type Recipe struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	BeanID      uuid.UUID
	Description string
	Steps       []Step
	DateCreated time.Time
	DateUpdated time.Time
}

type Step struct {
	Order             uint8
	Description       string
	DurationInSeconds int32
}
