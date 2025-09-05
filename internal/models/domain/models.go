package domain

import (
	"time"
)

type Event struct {
	Id     int       `json:"id"`
	UserId int       `json:"user_id" validate:"required"`
	Event  string    `json:"event" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
}
