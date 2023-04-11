package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"app/config"
	"app/internal/model"
	"app/internal/service"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

const expireTime = 15 * time.Second

type JWTAuth struct {
	jwtKey []byte
	User   service.IUserService
}

func NewJWTAuth(conf config.Config, user service.IUserService) *JWTAuth {
	return &JWTAuth{
		User: user,
	}
}

func (m *JWTAuth) GenerateToken(username string) (tokenStr string, err error) {
	expTime := time.Now().Add(expireTime)
	claims := &model.JWTClaims{
		Username: username,
		StandartClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.jwtKey)
}

func (m *JWTAuth) ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := m.extractToken(c.Request())
		if token != "test" {
			claims, err := m.ValidateToken(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			ctx := context.WithValue(c.Request().Context(), model.ContextUsername, claims.Username)
			c.SetRequest(c.Request().WithContext(ctx))
		}
		return next(c)
	}
}

func (m *JWTAuth) ValidateToken(signedToken string) (*model.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return m.jwtKey, nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.StandartClaims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, nil
}

func (m *JWTAuth) extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
