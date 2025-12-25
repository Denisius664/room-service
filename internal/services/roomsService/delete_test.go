package roomsservice

import (
	"context"
	"testing"

	"github.com/Denisius664/room-service/internal/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDelete_Success(t *testing.T) {
	ctx := context.Background()
	rs := &mockRoomStorage{}
	rp := &mockEventProducer{}
	cache := &mockRoomCache{}
	svc := NewRoomService(ctx, rs, rp, cache)

	rs.On("Delete", mock.Anything, "room1").Return(nil)
	rp.On("Produce", mock.Anything, mock.MatchedBy(func(e *models.RoomEvent) bool {
		return e != nil && e.Name == "room1" && e.Content == "deleted"
	})).Return(nil)
	cache.On("DeleteRoom", mock.Anything, "room1").Return(nil)

	err := svc.Delete(ctx, "room1")
	require.NoError(t, err)

	rs.AssertExpectations(t)
	rp.AssertExpectations(t)
	cache.AssertExpectations(t)
}

func TestDelete_Invalid(t *testing.T) {
	ctx := context.Background()
	rs := &mockRoomStorage{}
	rp := &mockEventProducer{}
	svc := NewRoomService(ctx, rs, rp, nil)

	err := svc.Delete(ctx, "")
	require.Error(t, err)
}
