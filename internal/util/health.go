package util

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/inidaname/mosque/mosques-service/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

// HealthAuthenticator wraps an Authenticator with health check capabilities
type HealthAuthenticator struct {
	types.Authenticator
	logger *slog.Logger
	mu     sync.RWMutex

	// Health state
	lastHealthy   bool
	lastCheck     time.Time
	lastError     error
	checkInterval time.Duration

	// Metrics
	checkCount   int
	failureCount int
	totalLatency time.Duration
}

// NewHealthAuthenticator creates a new health-wrapped authenticator
func NewHealthAuthenticator(
	auth types.Authenticator,
	logger *slog.Logger,
	checkInterval time.Duration,
) *HealthAuthenticator {
	if checkInterval == 0 {
		checkInterval = 30 * time.Second
	}

	return &HealthAuthenticator{
		Authenticator: auth,
		logger:        logger,
		checkInterval: checkInterval,
		lastCheck:     time.Now().Add(-checkInterval), // Force immediate check
	}
}

// HealthCheck performs a comprehensive health verification
func (h *HealthAuthenticator) HealthCheck(ctx context.Context) error {
	h.mu.RLock()

	// Return cached result if within interval
	if time.Since(h.lastCheck) < h.checkInterval {
		healthy := h.lastHealthy
		err := h.lastError
		h.mu.RUnlock()

		if !healthy {
			return fmt.Errorf("cached unhealthy state: %w", err)
		}
		return nil
	}
	h.mu.RUnlock()

	// Perform fresh check
	h.mu.Lock()
	defer h.mu.Unlock()

	start := time.Now()
	defer func() {
		h.totalLatency += time.Since(start)
		h.lastCheck = time.Now()
		h.checkCount++
	}()

	// Test token generation
	testClaims := types.Claims{
		"healthcheck": true,
		"exp":         time.Now().Add(5 * time.Minute).Unix(),
		"iat":         time.Now().Unix(),
	}

	token, err := h.GenerateToken(testClaims)
	if err != nil {
		h.recordFailure("token generation", err)
		return fmt.Errorf("token generation failed: %w", err)
	}

	// Test token validation
	if _, err := h.ValidateToken(token); err != nil {
		h.recordFailure("token validation", err)
		return fmt.Errorf("token validation failed: %w", err)
	}

	// Test invalid token rejection
	if _, err := h.ValidateToken("invalid.token.here"); err == nil {
		err := fmt.Errorf("invalid token was accepted")
		h.recordFailure("invalid token check", err)
		return err
	}

	h.lastHealthy = true
	h.lastError = nil
	return nil
}

// Stats returns health metrics and statistics
func (h *HealthAuthenticator) Stats() map[string]interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()

	avgLatency := time.Duration(0)
	if h.checkCount > 0 {
		avgLatency = h.totalLatency / time.Duration(h.checkCount)
	}

	return map[string]interface{}{
		"healthy":        h.lastHealthy,
		"last_check":     h.lastCheck.Format(time.RFC3339),
		"last_error":     fmt.Sprintf("%v", h.lastError),
		"check_count":    h.checkCount,
		"failure_count":  h.failureCount,
		"check_interval": h.checkInterval.String(),
		"avg_latency":    avgLatency.String(),
		"uptime":         time.Since(h.lastCheck).String(),
	}
}

func (h *HealthAuthenticator) recordFailure(operation string, err error) {
	h.lastHealthy = false
	h.lastError = err
	h.failureCount++

	h.logger.Error("Authentication health check failed",
		"operation", operation,
		"error", err,
		"next_check", time.Now().Add(h.checkInterval).Format(time.RFC3339),
	)
}

type HealthDB struct {
	*pgxpool.Pool
}

func (db *HealthDB) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Test both connection and basic query
	var now time.Time
	err := db.Pool.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}

func (db *HealthDB) Stats() *types.DBStats {
	stats := db.Pool.Stat()
	return &types.DBStats{
		MaxConns:     stats.MaxConns(),
		UsedConns:    stats.TotalConns(),
		IdleConns:    stats.IdleConns(),
		WaitCount:    stats.EmptyAcquireCount(),
		WaitDuration: stats.AcquireDuration().Milliseconds(),
		// MaxIdleTime:  stats.MaxIdleTime().Milliseconds(),
	}
}
