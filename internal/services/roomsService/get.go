package roomsservice

import (
	"context"
	"errors"

	"github.com/Denisius664/room-service/internal/models"
)

func (s *RoomService) Get(ctx context.Context, id string) (*models.RoomInfo, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.roomStorage.Get(ctx, id)
}
