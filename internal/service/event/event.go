package event

import (
	"context"
	"errors"
	"fmt"

	"github.com/avraam311/calendar-service/internal/models/domain"
)

var (
	ErrInvalidDate = errors.New("invalid event date")
)

type eventRepo interface {
	CreateEvent(ctx context.Context, event domain.Event) (int, error)
}

type Service struct {
	eventRepo eventRepo
}

func New(r eventRepo) *Service {
	return &Service{
		eventRepo: r,
	}
}

func (s *Service) CreateEvent(ctx context.Context, event domain.Event) (int, error) {
	id, err := s.eventRepo.CreateEvent(ctx, event)
	if err != nil {
		return -1, fmt.Errorf("create event - %w", err)
	}

	return id, nil
}
