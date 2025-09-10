package event

import (
	"context"
	"errors"
	"fmt"

	"github.com/avraam311/calendar-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrEventNotFound = errors.New("event not found")
	ErrNoNewData     = errors.New("nothing to update")
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateEvent(ctx context.Context, event *models.EventCreate) (uint, error) {
	query := `
		INSERT INTO events (
		    user_id, event, date
		) VALUES ($1, $2, $3)
		RETURNING id;
    `
	var ID uint
	err := r.db.QueryRow(ctx, query, event.UserID, event.Event, event.Date).Scan(&ID)
	if err != nil {
		return 0, fmt.Errorf("repository/CreateEvent - %w", err)
	}

	return ID, nil
}

func (r *Repository) UpdateEvent(ctx context.Context, event *models.EventUpdate) (uint, error) {
	query := `
		UPDATE events
		SET
			user_id = $1,
			event = $2,
		    date = $3
		WHERE id = $4;
	`

	cmdTag, err := r.db.Exec(ctx, query, event.UserID, event.Event, event.Date, event.ID)
	if cmdTag.RowsAffected() == 0 {
		return 0, ErrNoNewData
	}

	if err != nil {
		return 0, fmt.Errorf("repository/UpdateEvent - %w", err)
	}

	return event.ID, nil
}

func (r *Repository) DeleteEvent(ctx context.Context, ID uint) (uint, error) {
	query := `
   		DELETE FROM events
   		WHERE id = $1;
    `

	cmdTag, err := r.db.Exec(ctx, query, ID)
	if cmdTag.RowsAffected() == 0 {
		fmt.Println("on right way")
		return 0, ErrEventNotFound
	}

	if err != nil {
		return 0, fmt.Errorf("repository/DeleteEvent - %w", err)
	}

	return ID, nil
}
