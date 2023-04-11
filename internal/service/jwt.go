package service

import (
	"time"

	"app/config"
	"app/internal/model"

	"github.com/golang-jwt/jwt/v4"
)

const expireTime = 72 * time.Hour

type JWT struct {
	jwtKey []byte
	User   IUserService
}

func NewJWT(conf *config.Config, user IUserService) *JWT {
	return &JWT{
		jwtKey: []byte(conf.HTTP.JWTKey),
		User:   user,
	}
}

func (m *JWT) GenerateToken(username string) (tokenStr string, err error) {
	claims := &model.JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.jwtKey)
}

// func (m *JWTAuth) ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		token := m.extractToken(c.Request())
// 		if token != "test" {
// 			claims, err := m.ValidateToken(token)
// 			if err != nil {
// 				return echo.NewHTTPError(http.StatusForbidden, err.Error())
// 			}

// 			ctx := context.WithValue(c.Request().Context(), model.ContextUsername, claims.Username)
// 			c.SetRequest(c.Request().WithContext(ctx))
// 		}
// 		return next(c)
// 	}
// }

// func (m *JWTAuth) ValidateToken(signedToken string) (*model.JWTClaims, error) {
// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&model.JWTClaims{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return m.jwtKey, nil
// 		},
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	claims, ok := token.Claims.(*model.JWTClaims)
// 	if !ok {
// 		return nil, errors.New("couldn't parse claims")
// 	}
// 	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
// 		return nil, errors.New("token expired")
// 	}
// 	return claims, nil
// }

// func (m *JWTAuth) extractToken(r *http.Request) string {
// 	bearToken := r.Header.Get("Authorization")
// 	strArr := strings.Split(bearToken, " ")
// 	if len(strArr) == 2 {
// 		return strArr[1]
// 	}
// 	return ""
// }
