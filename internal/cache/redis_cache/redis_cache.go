package rediscache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"github.com/Denisius664/room-service/internal/models"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(host string, port int, password string, db int, ttlSeconds int) (*RedisCache, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	opt := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
	client := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis")
	}

	return &RedisCache{client: client, ttl: time.Duration(ttlSeconds) * time.Second}, nil
}

func (r *RedisCache) key(name string) string {
	return "room:" + name
}

func (r *RedisCache) GetRoom(ctx context.Context, name string) (*models.RoomInfo, error) {
	b, err := r.client.Get(ctx, r.key(name)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, errors.Wrap(err, "redis get")
	}
	var room models.RoomInfo
	if err := json.Unmarshal(b, &room); err != nil {
		return nil, errors.Wrap(err, "unmarshal room")
	}
	return &room, nil
}

func (r *RedisCache) SetRoom(ctx context.Context, room *models.RoomInfo) error {
	b, err := json.Marshal(room)
	if err != nil {
		return errors.Wrap(err, "marshal room")
	}
	if err := r.client.Set(ctx, r.key(room.Name), b, r.ttl).Err(); err != nil {
		return errors.Wrap(err, "redis set")
	}
	return nil
}

func (r *RedisCache) DeleteRoom(ctx context.Context, name string) error {
	if err := r.client.Del(ctx, r.key(name)).Err(); err != nil {
		return errors.Wrap(err, "redis del")
	}
	return nil
}
