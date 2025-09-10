package event

import (
	"context"
	"fmt"

	"github.com/avraam311/calendar-service/internal/models"
)

type eventRepo interface {
	CreateEvent(ctx context.Context, event *models.EventCreate) (uint, error)
	UpdateEvent(ctx context.Context, event *models.EventUpdate) (uint, error)
	DeleteEvent(ctx context.Context, ID uint) (uint, error)
}

type Service struct {
	eventRepo eventRepo
}

func New(r eventRepo) *Service {
	return &Service{
		eventRepo: r,
	}
}

func (s *Service) CreateEvent(ctx context.Context, event *models.EventCreate) (uint, error) {
	ID, err := s.eventRepo.CreateEvent(ctx, event)
	if err != nil {
		return 0, fmt.Errorf("service/CreateEvent - %w", err)
	}

	return ID, nil
}

func (s *Service) UpdateEvent(ctx context.Context, event *models.EventUpdate) (uint, error) {
	ID, err := s.eventRepo.UpdateEvent(ctx, event)
	if err != nil {
		return 0, fmt.Errorf("service/UpdateEvent - %w", err)
	}

	return ID, nil
}

func (s *Service) DeleteEvent(ctx context.Context, ID uint) (uint, error) {
	ID, err := s.eventRepo.DeleteEvent(ctx, ID)
	if err != nil {
		return 0, fmt.Errorf("service/DeleteEvent - %w", err)
	}

	return ID, nil
}
