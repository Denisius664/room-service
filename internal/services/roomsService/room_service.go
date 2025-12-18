package roomsservice

import (
	"context"

	"github.com/Denisius664/room-service/internal/models"
)

type roomStorage interface {
	Create(ctx context.Context, room *models.RoomInfo) error
	Get(ctx context.Context, id string) (*models.RoomInfo, error)
	Update(ctx context.Context, room *models.RoomInfo) error
	Delete(ctx context.Context, id string) error
}

type roomEventProducer interface {
	Produce(ctx context.Context, event *models.RoomEvent) error
}

type RoomService struct {
	roomStorage       roomStorage
	roomEventProducer roomEventProducer
}

func NewRoomService(ctx context.Context, roomStorage roomStorage, roomEventProducer roomEventProducer) *RoomService {
	return &RoomService{
		roomStorage:       roomStorage,
		roomEventProducer: roomEventProducer,
	}
}
