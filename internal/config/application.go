package config

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/inidaname/mosque/mosques-service/internal/cache"
	db "github.com/inidaname/mosque/mosques-service/internal/db"
	"github.com/inidaname/mosque/mosques-service/internal/types"
)

var (
	instance *types.Application
	once     sync.Once
)

func CreateApplication() *types.Application {
	once.Do(func() {
		cfg, err := LoadConfig("internal/config/config.yaml")

		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		store, dbPool, err := db.ConnectDB(ctx, logger, cfg)
		if err != nil {
			logger.Error("failed to connect to database", "error", err)
			os.Exit(1)
		}

		// Initialize thread-safe cache
		cacheService := cache.NewCacheService(5*time.Minute, 10*time.Minute)

		instance = &types.Application{
			Config: *cfg,
			Logger: logger,
			Store:  store,
			Db:     dbPool,
			Cache:  *cacheService,
		}
	})

	return instance
}
