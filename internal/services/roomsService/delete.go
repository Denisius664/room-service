package roomsservice

import (
	"context"
	"errors"
	"log"

	"github.com/Denisius664/room-service/internal/models"
)

func (s *RoomService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	if err := s.roomStorage.Delete(ctx, id); err != nil {
		return err
	}

	if s.roomCache != nil {
		if err := s.roomCache.DeleteRoom(ctx, id); err != nil {
			log.Printf("failed to delete cache for room %s: %v", id, err)
		}
	}

	if err := s.roomEventProducer.Produce(ctx, &models.RoomEvent{Name: id, Content: "deleted"}); err != nil {
		log.Printf("failed to produce room deleted event for %s: %v", id, err)
	}
	return nil
}
