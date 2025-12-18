package pgstorage

import (
	"context"
	"os"
	"testing"

	"github.com/Denisius664/room-service/internal/models"
	"github.com/stretchr/testify/require"
)

// These tests are integration tests against a real Postgres instance.
// They are skipped unless TEST_PG_CONN environment variable is set to a valid
// pgx connection string (for example, "postgres://user:pass@localhost:5432/dbname").
func TestPGstorage_CRUD_Integration(t *testing.T) {
	conn := os.Getenv("TEST_PG_CONN")
	if conn == "" {
		t.Skip("integration test skipped; set TEST_PG_CONN to run")
	}

	s, err := NewPGStorge(conn)
	require.NoError(t, err)

	ctx := context.Background()
	room := &models.RoomInfo{Name: "itest_room_1", Users: []string{"alice", "bob"}}

	// Create
	err = s.Create(ctx, room)
	require.NoError(t, err)

	// Get
	got, err := s.Get(ctx, room.Name)
	require.NoError(t, err)
	require.Equal(t, room.Name, got.Name)
	require.Equal(t, room.Users, got.Users)

	// Update
	room.Users = append(room.Users, "carol")
	err = s.Update(ctx, room)
	require.NoError(t, err)

	got2, err := s.Get(ctx, room.Name)
	require.NoError(t, err)
	require.Equal(t, room.Users, got2.Users)

	// Delete
	err = s.Delete(ctx, room.Name)
	require.NoError(t, err)

	// After delete, Get should return an error
	_, err = s.Get(ctx, room.Name)
	require.Error(t, err)
}
