package main

import (
	"fmt"
	"os"

	"github.com/Denisius664/room-service/config"
	"github.com/Denisius664/room-service/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig(os.Getenv("configPath"))
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфига, %v", err))
	}
	roomStorage := bootstrap.InitPGStorage(cfg)
	roomEventProducer := bootstrap.InitRoomEventsProducer(cfg)
	roomService := bootstrap.InitRoomService(roomStorage, cfg, roomEventProducer)

	playerCommandProducer := bootstrap.InitPlayerCommandProducer(cfg)
	playerService := bootstrap.InitPlayerService(playerCommandProducer)

	bootstrap.AppRun(*roomService, *playerService)
}
