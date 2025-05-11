package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/inidaname/mosque/mosques-service/internal/types"
)

type JWTAuthenticator struct {
	secretKey     []byte
	signingMethod jwt.SigningMethod
	issuer        string
	audience      string
}

func NewJWTAuthenticator(secret string, issuer string, audience string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secretKey:     []byte(secret),
		signingMethod: jwt.SigningMethodHS256,
		issuer:        issuer,
		audience:      audience,
	}
}

func (j *JWTAuthenticator) GenerateToken(claims types.Claims) (string, error) {
	// Convert claims to jwt.MapClaims
	jwtClaims := jwt.MapClaims{}
	for k, v := range claims {
		jwtClaims[k] = v
	}

	// Set standard claims
	jwtClaims["iss"] = j.issuer
	jwtClaims["aud"] = j.audience

	token := jwt.NewWithClaims(j.signingMethod, jwtClaims)
	return token.SignedString(j.secretKey)
}

func (j *JWTAuthenticator) ValidateToken(tokenString string) (*types.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Convert back to our Claims type
		resultClaims := make(types.Claims)
		for k, v := range claims {
			resultClaims[k] = v
		}

		return &types.Token{
			Raw:    tokenString,
			Claims: resultClaims,
		}, nil
	}

	return nil, fmt.Errorf("invalid token")
}
