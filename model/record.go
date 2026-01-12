package model

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	ServiceName string    `json:"service_name,omitempty"`
	Price       int       `json:"price,omitempty"`
	UserID      uuid.UUID `json:"user_id,omitempty"`
	StartDate   time.Time `json:"start_date,omitempty"`
}
