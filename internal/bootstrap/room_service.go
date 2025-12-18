package bootstrap

import (
	"context"

	"github.com/Denisius664/room-service/config"
	roomeventproducer "github.com/Denisius664/room-service/internal/producer/room_event_producer"
	roomsservice "github.com/Denisius664/room-service/internal/services/roomsService"
	"github.com/Denisius664/room-service/internal/storage/pgstorage"
)

func InitRoomService(storage *pgstorage.PGstorage, cfg *config.Config, producer *roomeventproducer.RoomEventProducer) *roomsservice.RoomService {

	return roomsservice.NewRoomService(context.Background(), storage, producer)
}
