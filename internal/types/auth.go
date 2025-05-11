package types

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type Claims = jwt.MapClaims
type Token = jwt.Token

type Authenticator interface {
	GenerateToken(claims Claims) (string, error)
	ValidateToken(token string) (*Token, error)
}

// HealthCheckableAuthenticator extends the base interface with health checks
type HealthCheckableAuthenticator interface {
	Authenticator
	HealthCheck(ctx context.Context) error
	Stats() map[string]interface{}
}
