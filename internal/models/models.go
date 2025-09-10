package models

import "time"

type EventDelete struct {
	ID uint `json:"id" validate:"required"`
}

type EventCreate struct {
	UserID int       `json:"user_id" validate:"required"`
	Event  string    `json:"event" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
}

type EventUpdate struct {
	ID     uint      `json:"id" validate:"required"`
	UserID int       `json:"user_id" validate:"required"`
	Event  string    `json:"event" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
}
