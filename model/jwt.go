package model

import "github.com/golang-jwt/jwt"

type JWTClaim struct {
	Login string
	jwt.StandardClaims
}

type ContextType string

const ContextLogin = ContextType("user")
