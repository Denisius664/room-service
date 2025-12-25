package roomsservice

import (
	"context"
	"testing"

	"github.com/Denisius664/room-service/internal/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockRoomStorage struct{ mock.Mock }

func (m *mockRoomStorage) Create(ctx context.Context, room *models.RoomInfo) error {
	args := m.Called(ctx, room)
	return args.Error(0)
}
func (m *mockRoomStorage) Get(ctx context.Context, id string) (*models.RoomInfo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.RoomInfo), args.Error(1)
}
func (m *mockRoomStorage) Update(ctx context.Context, room *models.RoomInfo) error {
	args := m.Called(ctx, room)
	return args.Error(0)
}
func (m *mockRoomStorage) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type mockEventProducer struct{ mock.Mock }

func (m *mockEventProducer) Produce(ctx context.Context, event *models.RoomEvent) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

type mockRoomCache struct{ mock.Mock }

func (m *mockRoomCache) GetRoom(ctx context.Context, name string) (*models.RoomInfo, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.RoomInfo), args.Error(1)
}
func (m *mockRoomCache) SetRoom(ctx context.Context, room *models.RoomInfo) error {
	args := m.Called(ctx, room)
	return args.Error(0)
}
func (m *mockRoomCache) DeleteRoom(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func TestCreate_Success(t *testing.T) {
	ctx := context.Background()
	rs := &mockRoomStorage{}
	rp := &mockEventProducer{}

	cache := &mockRoomCache{}
	svc := NewRoomService(ctx, rs, rp, cache)

	room := &models.RoomInfo{Name: "room1", Users: []string{"alice"}}

	rs.On("Create", mock.Anything, room).Return(nil)
	rp.On("Produce", mock.Anything, mock.MatchedBy(func(e *models.RoomEvent) bool {
		return e != nil && e.Name == room.Name && e.Content == "created"
	})).Return(nil)
	cache.On("SetRoom", mock.Anything, room).Return(nil)

	err := svc.Create(ctx, room)
	require.NoError(t, err)

	rs.AssertExpectations(t)
	rp.AssertExpectations(t)
	cache.AssertExpectations(t)
}

func TestCreate_Invalid(t *testing.T) {
	ctx := context.Background()
	rs := &mockRoomStorage{}
	rp := &mockEventProducer{}
	svc := NewRoomService(ctx, rs, rp, nil)

	err := svc.Create(ctx, nil)
	require.Error(t, err)

	err = svc.Create(ctx, &models.RoomInfo{Name: ""})
	require.Error(t, err)
}
