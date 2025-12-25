package bootstrap

import (
	"log"
	"net/http"

	roomserviceapi "github.com/Denisius664/room-service/internal/api/room_service_api"
	chatservice "github.com/Denisius664/room-service/internal/services/chatService"
	playerservice "github.com/Denisius664/room-service/internal/services/playerService"
	roomsservice "github.com/Denisius664/room-service/internal/services/roomsService"
)

func AppRun(roomService roomsservice.RoomService, playerService playerservice.PlayerService, chatService chatservice.ChatService) {
	router := roomserviceapi.NewRouter()
	roomHandler := roomserviceapi.NewRoomHandler(&roomService)
	roomserviceapi.RegisterRoomHandler(&router, roomHandler)
	playerHandler := roomserviceapi.NewPlayerHandler(&playerService)
	roomserviceapi.RegisterPlayerHandler(&router, playerHandler)
	chatHandler := roomserviceapi.NewChatHandler(&chatService)
	roomserviceapi.RegisterChatHandler(&router, chatHandler)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", &router)
}
