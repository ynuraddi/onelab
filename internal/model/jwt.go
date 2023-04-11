package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	Username       string
	StandartClaims jwt.RegisteredClaims
}

func (m *JWTClaims) Valid() error {
	// if m.Username == "" {
	// 	return jwt.ErrTokenInvalidClaims
	// }

	// if m.StandartClaims.ExpiresAt.Before(time.Now()) {
	// 	return jwt.ErrTokenExpired
	// }

	return nil
}
