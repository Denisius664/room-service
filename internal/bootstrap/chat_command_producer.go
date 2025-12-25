package bootstrap

import (
	"fmt"

	"github.com/Denisius664/room-service/config"
	chatcommandproducer "github.com/Denisius664/room-service/internal/producer/chat_command_producer"
)

func InitChatCommandProducer(cfg *config.Config) *chatcommandproducer.ChatCommandProducer {
	kafkaBrokers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return chatcommandproducer.NewChatCommandProducer(kafkaBrokers, cfg.Kafka.ChatCommandTopic)
}
