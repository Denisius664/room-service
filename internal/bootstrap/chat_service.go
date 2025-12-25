package bootstrap

import (
	chatcommandproducer "github.com/Denisius664/room-service/internal/producer/chat_command_producer"
	chatservice "github.com/Denisius664/room-service/internal/services/chatService"
)

func InitChatService(chatCommandProducer *chatcommandproducer.ChatCommandProducer) *chatservice.ChatService {
	return chatservice.NewChatService(chatCommandProducer)
}
