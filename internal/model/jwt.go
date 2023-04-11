package model

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	Username string `json:"name"`
	jwt.RegisteredClaims
}

func (m *JWTClaims) Valid() error {
	if len(m.Username) == 0 {
		return fmt.Errorf("model(jwtValid): %w", jwt.ErrTokenInvalidClaims)
	}

	return nil
}
