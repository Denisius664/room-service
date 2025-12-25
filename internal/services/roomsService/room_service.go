package roomsservice

import (
	"context"

	"github.com/Denisius664/room-service/internal/models"
)

type RoomStorage interface {
	Create(ctx context.Context, room *models.RoomInfo) error
	Get(ctx context.Context, id string) (*models.RoomInfo, error)
	Update(ctx context.Context, room *models.RoomInfo) error
	Delete(ctx context.Context, id string) error
}

type RoomEventProducer interface {
	Produce(ctx context.Context, event *models.RoomEvent) error
}

type RoomService struct {
	roomStorage       RoomStorage
	roomEventProducer RoomEventProducer
	roomCache         RoomCache
}

// RoomCache is a minimal cache interface used by RoomService (e.g., Redis adapter).
type RoomCache interface {
	GetRoom(ctx context.Context, name string) (*models.RoomInfo, error)
	SetRoom(ctx context.Context, room *models.RoomInfo) error
	DeleteRoom(ctx context.Context, name string) error
}

// NewRoomService constructs a RoomService. Pass nil for roomCache to disable caching.
func NewRoomService(ctx context.Context, roomStorage RoomStorage, roomEventProducer RoomEventProducer, roomCache RoomCache) *RoomService {
	return &RoomService{
		roomStorage:       roomStorage,
		roomEventProducer: roomEventProducer,
		roomCache:         roomCache,
	}
}
