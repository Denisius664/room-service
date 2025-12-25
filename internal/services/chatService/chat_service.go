package chatservice

import (
	"context"

	"github.com/Denisius664/room-service/internal/models"
)

type chatCommandProducer interface {
	Produce(ctx context.Context, cmd *models.SendMessageCommand) error
}

type ChatService struct {
	chatCommandProducer chatCommandProducer
}

func NewChatService(chatCommandProducer chatCommandProducer) *ChatService {
	return &ChatService{chatCommandProducer: chatCommandProducer}
}

func (s *ChatService) Send(ctx context.Context, toRoom string, sender string, content string) error {
	return s.chatCommandProducer.Produce(ctx, &models.SendMessageCommand{ToRoom: toRoom, Sender: sender, Content: content})
}
