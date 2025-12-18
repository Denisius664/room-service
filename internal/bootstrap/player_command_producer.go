package bootstrap

import (
	"fmt"

	"github.com/Denisius664/room-service/config"
	playercommandproducer "github.com/Denisius664/room-service/internal/producer/player_command_producer"
)

func InitPlayerCommandProducer(cfg *config.Config) *playercommandproducer.PlayerCommandProducer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return playercommandproducer.NewPlayerCommandProducer(kafkaBrockers, cfg.Kafka.PlayerCommandTopic)
}
