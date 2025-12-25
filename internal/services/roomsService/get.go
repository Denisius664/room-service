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
	// try cache first
	if s.roomCache != nil {
		if v, err := s.roomCache.GetRoom(ctx, id); err == nil && v != nil {
			return v, nil
		}
		// on error we fallthrough to storage
	}

	room, err := s.roomStorage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if s.roomCache != nil && room != nil {
		if err := s.roomCache.SetRoom(ctx, room); err != nil {
			// best-effort
		}
	}

	return room, nil
}
