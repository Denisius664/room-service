package roomserviceapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type playerService interface {
	Play(ctx context.Context, playerID string) error
	Pause(ctx context.Context, playerID string) error
	Seek(ctx context.Context, playerID string, pos int) error
}

type PlayerHandler struct {
	service playerService
}

func NewPlayerHandler(s playerService) *PlayerHandler {
	return &PlayerHandler{service: s}
}

func (h *PlayerHandler) Play(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.service.Play(r.Context(), id)

	w.WriteHeader(http.StatusOK)
}

func (h *PlayerHandler) Pause(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	h.service.Pause(r.Context(), id)

	w.WriteHeader(http.StatusOK)
}

func (h *PlayerHandler) Seek(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body struct{ Position int }
	json.NewDecoder(r.Body).Decode(&body)
	h.service.Seek(r.Context(), id, body.Position)

	w.WriteHeader(http.StatusOK)
}
