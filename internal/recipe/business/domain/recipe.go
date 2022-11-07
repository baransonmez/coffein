package domain

import (
	"github.com/google/uuid"
	"time"
)

type Recipe struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	BeanID      uuid.UUID `json:"bean_id"`
	Description string    `json:"description"`
	Steps       []Step    `json:"steps"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type Step struct {
	Order             uint8  `json:"order"`
	Description       string `json:"description"`
	DurationInSeconds int32  `json:"duration"`
}
