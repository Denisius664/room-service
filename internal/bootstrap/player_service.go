package bootstrap

import (
	playercommandproducer "github.com/Denisius664/room-service/internal/producer/player_command_producer"
	playerservice "github.com/Denisius664/room-service/internal/services/playerService"
)

func InitPlayerService(playerCommandProducer *playercommandproducer.PlayerCommandProducer) *playerservice.PlayerService {
	return playerservice.NewPlayerService(playerCommandProducer)
}
