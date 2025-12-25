package bootstrap

import (
	"log"

	"github.com/Denisius664/room-service/config"
	rediscache "github.com/Denisius664/room-service/internal/cache/redis_cache"
)

func InitRedis(cfg *config.Config) *rediscache.RedisCache {
	if cfg == nil {
		return nil
	}
	r, err := rediscache.NewRedisCache(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB, cfg.Redis.TTLSec)
	if err != nil {
		log.Printf("failed to init redis cache: %v", err)
		return nil
	}
	return r
}
