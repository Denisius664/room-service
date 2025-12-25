package roomsservice

import (
	"context"
	"testing"

	"github.com/Denisius664/room-service/internal/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdate_Success(t *testing.T) {
	ctx := context.Background()
	rs := &mockRoomStorage{}
	rp := &mockEventProducer{}
	cache := &mockRoomCache{}
	svc := NewRoomService(ctx, rs, rp, cache)

	room := &models.RoomInfo{Name: "room1", Users: []string{"alice"}}
	rs.On("Update", mock.Anything, room).Return(nil)
	rp.On("Produce", mock.Anything, mock.MatchedBy(func(e *models.RoomEvent) bool {
		return e != nil && e.Name == room.Name && e.Content == "updated"
	})).Return(nil)
	cache.On("SetRoom", mock.Anything, room).Return(nil)

	err := svc.Update(ctx, room)
	require.NoError(t, err)

	rs.AssertExpectations(t)
	rp.AssertExpectations(t)
	cache.AssertExpectations(t)
}

func TestUpdate_Invalid(t *testing.T) {
	ctx := context.Background()
	rs := &mockRoomStorage{}
	rp := &mockEventProducer{}
	svc := NewRoomService(ctx, rs, rp, nil)

	err := svc.Update(ctx, nil)
	require.Error(t, err)

	err = svc.Update(ctx, &models.RoomInfo{Name: ""})
	require.Error(t, err)
}
