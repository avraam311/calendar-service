package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/avraam311/calendar-service/internal/api/handlers/event"
)

func NewRouter(eventPostHandler *event.PostHandler, eventGetHandler *event.GetHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*"},
		AllowedMethods:   []string{"POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Post("/create_event", eventPostHandler.CreateEvent)
		r.Put("/update_event", eventPostHandler.UpdateEvent)
		r.Delete("/delete_event", eventPostHandler.DeleteEvent)
		r.Get("/events_for_day", eventGetHandler.GetEventsForDay)
		r.Get("/events_for_week", eventGetHandler.GetEventsForWeek)
		r.Get("/events_for_month", eventGetHandler.GetEventsForMonth)
	})

	return r
}

func NewServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
