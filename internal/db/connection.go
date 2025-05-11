package db

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	db "github.com/inidaname/mosque/mosques-service/internal/db/models"
	"github.com/inidaname/mosque/mosques-service/internal/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context, logger *slog.Logger, cfg *types.Config) (*db.Queries, *pgxpool.Pool, error) {
	const op = "db.Connection"
	logger = logger.With("operation", op)

	config, err := pgxpool.ParseConfig(cfg.Database.Url)
	if err != nil {
		logger.Error("failed to parse DB config", "error", err)
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	// Recommended production settings
	config.MaxConns = 50
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 30 * time.Second

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.Error("failed to create connection pool", "error", err)
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := conn.Ping(ctx); err != nil {
		conn.Close()
		logger.Error("failed to ping database", "error", err)
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("database connection established",
		"max_conns", config.MaxConns,
		"min_conns", config.MinConns)

	return db.New(conn), conn, nil
}

type DBTracer struct {
	logger *slog.Logger
}

func (t *DBTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	start := time.Now()
	t.logger.Debug("executing query",
		"sql", data.SQL,
		"args", maskSensitiveArgs(data.Args))
	return context.WithValue(ctx, "query_start", start)
}

type QueryTracer struct {
	logger *slog.Logger
}

func (t *QueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	// Get start time from context
	start, ok := ctx.Value("query_start").(time.Time)
	if !ok {
		t.logger.Warn("missing query start time in context")
		return
	}

	duration := time.Since(start)

	if data.Err != nil {
		t.logger.Warn("query failed",
			"error", data.Err,
			"duration_ms", duration.Milliseconds())
	} else {
		t.logger.Debug("query completed",
			"duration_ms", duration.Milliseconds(),
			"rows_affected", data.CommandTag.RowsAffected())
	}
}

// maskSensitiveArgs hides sensitive data in query arguments
func maskSensitiveArgs(args []interface{}) []interface{} {
	masked := make([]interface{}, len(args))
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			if len(v) > 8 && (i == 0 || isSensitiveParam(i)) {
				masked[i] = "*****"
			} else {
				masked[i] = v
			}
		case []byte:
			masked[i] = "*****"
		default:
			masked[i] = v
		}
	}
	return masked
}

func isSensitiveParam(index int) bool {
	// Add logic to identify sensitive parameter positions
	// Example: return index == 1 // If second parameter is always sensitive
	return false
}
