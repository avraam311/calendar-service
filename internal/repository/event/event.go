package event

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/avraam311/calendar-service/internal/models/domain"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateEvent(ctx context.Context, event domain.Event) (int, error) {
	query := `
		INSERT INTO events (
		    user_id, event, date
		) VALUES ($1, $2, $3)
		RETURNING id;
    `

	err := r.db.QueryRow(
		ctx, query, event.UserId, event.Event, event.Date,
	).Scan(&event.Id)
	if err != nil {
		return -1, fmt.Errorf("repository/event/event.go, failed to create event - %w", err)
	}

	return event.Id, nil
}
