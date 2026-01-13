package bootstrap

import (
	"fmt"
	"log"

	"github.com/Denisius664/room-service/config"
	"github.com/Denisius664/room-service/internal/storage/pgstorage"
)

func InitPGStorage(cfg *config.Config) *pgstorage.PGstorage {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode)
	storage, err := pgstorage.NewPGStorge(connectionString)
	if err != nil {
		log.Panicf("ошибка инициализации БД, %v", err)
		panic(err)
	}
	return storage
}
