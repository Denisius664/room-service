package playerservice

import (
	"context"
	"fmt"

	"github.com/Denisius664/room-service/internal/models"
)

type playerCommandProducer interface {
	Produce(ctx context.Context, event *models.PlayerCommand) error
}

type PlayerService struct {
	playerCommandProducer playerCommandProducer
}

func NewPlayerService(playerCommandProducer playerCommandProducer) *PlayerService {
	return &PlayerService{playerCommandProducer: playerCommandProducer}
}

func (s *PlayerService) Play(ctx context.Context, playerID string) error {
	return s.playerCommandProducer.Produce(ctx, &models.PlayerCommand{Name: playerID, Content: "Play"})
}

func (s *PlayerService) Pause(ctx context.Context, playerID string) error {
	return s.playerCommandProducer.Produce(ctx, &models.PlayerCommand{Name: playerID, Content: "Pause"})
}

func (s *PlayerService) Seek(ctx context.Context, playerID string, pos int) error {
	return s.playerCommandProducer.Produce(ctx, &models.PlayerCommand{Name: playerID, Content: fmt.Sprintf("Seek %d", pos)})
}
