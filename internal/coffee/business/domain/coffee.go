package domain

import (
	"github.com/google/uuid"
	"time"
)

type Bean struct {
	ID          uuid.UUID
	Name        string
	Roaster     string
	Origin      string
	Price       int
	RoastDate   time.Time
	DateCreated time.Time
	DateUpdated time.Time
}
