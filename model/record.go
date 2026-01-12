package model

import (
	"time"

	"github.com/jackc/pgx/pgtype"
)

type Record struct {
	ID          int         `json:"id,omitempty"`
	ServiceName string      `json:"service_name,omitempty"`
	Price       int         `json:"price,omitempty"`
	UserID      pgtype.UUID `json:"user_id,omitempty"`
	StartDate   time.Time   `json:"start_date,omitempty"`
}
