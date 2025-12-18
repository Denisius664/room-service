package roomserviceapi

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Mux {
	r := chi.NewRouter()
	return *r
}

func RegisterRoomHandler(router *chi.Mux, room *RoomHandler) {
	router.Post("/rooms/{name}", room.CreateRoom)
	router.Post("/rooms/{name}/join", room.JoinRoom)
	router.Get("/rooms/{name}", room.GetRoom)
	router.Delete("/rooms/{name}", room.DeleteRoom)
}

func RegisterPlayerHandler(router *chi.Mux, player *PlayerHandler) {
	router.Post("/player/{id}/play", player.Play)
	router.Post("/player/{id}/pause", player.Pause)
	router.Post("/player/{id}/seek", player.Seek)
}
