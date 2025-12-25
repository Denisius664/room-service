package roomserviceapi

import (
	"context"
	"encoding/json"
	"net/http"
)

type chatService interface {
	Send(ctx context.Context, toRoom string, sender string, content string) error
}

type ChatHandler struct {
	service chatService
}

func NewChatHandler(s chatService) *ChatHandler {
	return &ChatHandler{service: s}
}

func (h *ChatHandler) Send(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ToRoom  string `json:"toRoom"`
		Sender  string `json:"sender"`
		Content string `json:"content"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	h.service.Send(r.Context(), body.ToRoom, body.Sender, body.Content)

	w.WriteHeader(http.StatusOK)
}
