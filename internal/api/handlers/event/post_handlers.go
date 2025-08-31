package event

import (
	"net/http"

	"go.uber.org/zap"
)

type eventService interface{}

type PostHandler struct {
	logger       *zap.Logger
	eventService eventService
}

func NewPostHandler(l *zap.Logger, s eventService) *PostHandler {
	return &PostHandler{
		logger:       l,
		eventService: s,
	}
}

func (h *PostHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	
}