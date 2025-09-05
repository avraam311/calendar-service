package event

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/avraam311/calendar-service/internal/models/domain"
	"github.com/avraam311/calendar-service/internal/pkg/validator"
)

type eventService interface {
	CreateEvent(ctx context.Context, event domain.Event) (int, error)
}

type PostHandler struct {
	logger       *zap.Logger
	validator    *validator.GoValidator
	eventService eventService
}

func NewPostHandler(l *zap.Logger, v *validator.GoValidator, s eventService) *PostHandler {
	return &PostHandler{
		logger:       l,
		eventService: s,
		validator:    v,
	}
}

func (h *PostHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.logger.Warn("not allowed methods")
		h.handleError(w, http.StatusBadRequest, "only method POST allowed")
		return
	}

	var event domain.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		h.logger.Warn("failed to decode JSON", zap.Error(err))
		h.handleError(w, http.StatusBadRequest, "invalid json")
		return
	}

	err = h.validator.Validate(event)
	if err != nil {
		h.logger.Warn("validation error", zap.Error(err))
		h.handleError(w, http.StatusBadRequest, "validation error")
		return
	}

	userId, err := h.eventService.CreateEvent(r.Context(), event)
	if err != nil {
		h.logger.Error("failed to create event", zap.Error(err))
		h.handleError(w, http.StatusInternalServerError, "internal error")
		return
	}

	h.logger.Info("event created", zap.Any("event", event))

	response := map[string]string{
		"result": strconv.Itoa(userId),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		h.logger.Error("failed to encode error response", zap.Error(err))
		http.Error(w, "error response encoding error", http.StatusInternalServerError)
	}
}

func (h *PostHandler) handleError(w http.ResponseWriter, code int, msg string) {
	errorResponse := map[string]string{
		"error": msg,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		h.logger.Error("failed to encode error response", zap.Error(err))
		http.Error(w, "error response encoding error", http.StatusInternalServerError)
	}
}
