package roomsservice

import (
	"context"
	"testing"

	"github.com/Denisius664/room-service/internal/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGet_Success(t *testing.T) {
	ctx := context.Background()
	rs := &mockRoomStorage{}
	rp := &mockEventProducer{}
	cache := &mockRoomCache{}
	svc := NewRoomService(ctx, rs, rp, cache)

	room := &models.RoomInfo{Name: "room1", Users: []string{"alice"}}
	// Simulate cache miss -> storage hit
	cache.On("GetRoom", mock.Anything, "room1").Return(nil, nil)
	rs.On("Get", mock.Anything, "room1").Return(room, nil)
	cache.On("SetRoom", mock.Anything, room).Return(nil)

	got, err := svc.Get(ctx, "room1")
	require.NoError(t, err)
	require.Equal(t, room, got)

	cache.AssertExpectations(t)
	rs.AssertExpectations(t)
}

func TestGet_Invalid(t *testing.T) {
	ctx := context.Background()
	rs := &mockRoomStorage{}
	rp := &mockEventProducer{}
	svc := NewRoomService(ctx, rs, rp, nil)

	_, err := svc.Get(ctx, "")
	require.Error(t, err)
}
