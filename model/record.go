package model

import (
	"github.com/jackc/pgx/pgtype"
)

type Record struct {
	ID          int         `json:"id,omitempty"`
	ServiceName string      `json:"service_name,omitempty"`
	Price       int         `json:"price,omitempty"`
	UserID      pgtype.UUID `json:"user_id,omitempty"`
	StartDate   string      `json:"start_date,omitempty"`
}
