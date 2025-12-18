package bootstrap

import (
	"fmt"

	"github.com/Denisius664/room-service/config"
	roomeventproducer "github.com/Denisius664/room-service/internal/producer/room_event_producer"
)

func InitRoomEventsProducer(cfg *config.Config) *roomeventproducer.RoomEventProducer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return roomeventproducer.NewRoomEventProducer(kafkaBrockers, cfg.Kafka.RoomEventsTopic)
}
