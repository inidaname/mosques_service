package types

import (
	"log/slog"

	// cache "github.com/inidaname/mosque/mosques-service/cache"
	"github.com/inidaname/mosque/mosques-service/internal/cache"
	db "github.com/inidaname/mosque/mosques-service/internal/db/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Config Config
	Logger *slog.Logger
	Store  *db.Queries
	Db     *pgxpool.Pool
	// Mailer        mailer.Mailer
	Cache               cache.CacheService
	HealthAuthenticator HealthCheckableAuthenticator
}
