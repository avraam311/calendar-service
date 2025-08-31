package models

import (
	"time"
)

type Event struct {
	UserId int       `json:"user_id" validate:"required"`
	Event  string    `json:"event" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
}
