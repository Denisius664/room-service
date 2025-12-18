package bootstrap

import (
	"log"
	"net/http"

	roomserviceapi "github.com/Denisius664/room-service/internal/api/room_service_api"
	playerservice "github.com/Denisius664/room-service/internal/services/playerService"
	roomsservice "github.com/Denisius664/room-service/internal/services/roomsService"
)

func AppRun(roomService roomsservice.RoomService, playerService playerservice.PlayerService) {
	router := roomserviceapi.NewRouter()
	roomHandler := roomserviceapi.NewRoomHandler(&roomService)
	roomserviceapi.RegisterRoomHandler(&router, roomHandler)
	playerHandler := roomserviceapi.NewPlayerHandler(&playerService)
	roomserviceapi.RegisterPlayerHandler(&router, playerHandler)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", &router)
}
