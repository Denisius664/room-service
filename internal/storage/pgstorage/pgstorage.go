package pgstorage

import (
	"context"
	"fmt"

	"github.com/Denisius664/room-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type PGstorage struct {
	db *pgxpool.Pool
}

func NewPGStorge(connString string) (*PGstorage, error) {

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка парсинга конфига")
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка подключения")
	}
	storage := &PGstorage{
		db: db,
	}
	err = storage.initTables()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *PGstorage) initTables() error {
	sql := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %v (
        %v TEXT PRIMARY KEY,
        %v TEXT[]
    )`, tableName, NameColumnName, UsersColumnName)
	_, err := s.db.Exec(context.Background(), sql)
	if err != nil {
		return errors.Wrap(err, "initition tables")
	}
	return nil
}

// Create inserts a new room record.
func (s *PGstorage) Create(ctx context.Context, room *models.RoomInfo) error {
	sql := fmt.Sprintf("INSERT INTO %v (%v, %v) VALUES ($1, $2)", tableName, NameColumnName, UsersColumnName)
	_, err := s.db.Exec(ctx, sql, room.Name, room.Users)
	if err != nil {
		return errors.Wrap(err, "create room")
	}
	return nil
}

// Get retrieves a room by id (name).
func (s *PGstorage) Get(ctx context.Context, id string) (*models.RoomInfo, error) {
	sql := fmt.Sprintf("SELECT %v, %v FROM %v WHERE %v = $1", NameColumnName, UsersColumnName, tableName, NameColumnName)
	row := s.db.QueryRow(ctx, sql, id)
	var name string
	var users []string
	if err := row.Scan(&name, &users); err != nil {
		return nil, errors.Wrap(err, "get room")
	}
	return &models.RoomInfo{Name: name, Users: users}, nil
}

// Update updates the users list for an existing room.
func (s *PGstorage) Update(ctx context.Context, room *models.RoomInfo) error {
	sql := fmt.Sprintf("UPDATE %v SET %v = $1 WHERE %v = $2", tableName, UsersColumnName, NameColumnName)
	_, err := s.db.Exec(ctx, sql, room.Users, room.Name)
	if err != nil {
		return errors.Wrap(err, "update room")
	}
	return nil
}

// Delete removes a room by id (name).
func (s *PGstorage) Delete(ctx context.Context, id string) error {
	sql := fmt.Sprintf("DELETE FROM %v WHERE %v = $1", tableName, NameColumnName)
	_, err := s.db.Exec(ctx, sql, id)
	if err != nil {
		return errors.Wrap(err, "delete room")
	}
	return nil
}
