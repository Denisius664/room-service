package roomsservice

import (
	"context"
	"errors"
	"log"

	"github.com/Denisius664/room-service/internal/models"
)

func (s *RoomService) Update(ctx context.Context, room *models.RoomInfo) error {
	if room == nil {
		return errors.New("room is required")
	}
	if room.Name == "" {
		return errors.New("room name is required")
	}

	if err := s.roomStorage.Update(ctx, room); err != nil {
		return err
	}

	// update cache
	if s.roomCache != nil {
		if err := s.roomCache.SetRoom(ctx, room); err != nil {
			log.Printf("failed to set cache for room %s: %v", room.Name, err)
		}
	}

	if err := s.roomEventProducer.Produce(ctx, &models.RoomEvent{Name: room.Name, Content: "updated"}); err != nil {
		log.Printf("failed to produce room updated event for %s: %v", room.Name, err)
	}
	return nil
}
