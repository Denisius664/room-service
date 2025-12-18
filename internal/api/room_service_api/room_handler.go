package roomserviceapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Denisius664/room-service/internal/models"

	"github.com/go-chi/chi/v5"
)

type roomService interface {
	Create(ctx context.Context, room *models.RoomInfo) error
	Get(ctx context.Context, id string) (*models.RoomInfo, error)
	Update(ctx context.Context, room *models.RoomInfo) error
	Delete(ctx context.Context, id string) error
}

type RoomHandler struct {
	roomService roomService
}

func NewRoomHandler(s roomService) *RoomHandler {
	return &RoomHandler{roomService: s}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	room := models.RoomInfo{Name: chi.URLParam(r, "name")}
	if err := h.roomService.Create(r.Context(), &room); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"name": room.Name})
}

func (h *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "name")
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username required", http.StatusBadRequest)
		return
	}
	room, err := h.roomService.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	room.Join(username)
	h.roomService.Update(r.Context(), room)

	w.WriteHeader(http.StatusOK)
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "name")
	room, err := h.roomService.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "name")
	h.roomService.Delete(r.Context(), id)

	w.WriteHeader(http.StatusOK)
}
