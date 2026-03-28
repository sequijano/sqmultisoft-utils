package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID         uint   `json:"user_id"`
	Username       string `json:"username"`
	Role           string `json:"role"`
	TenantID       uint   `json:"tenant_id"`
	TenantPlan     string `json:"tenant_plan"`
	SessionVersion int    `json:"session_version"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username, role string, tenantID uint, tenantPlan string, sessionVersion int, secret string) (string, error) {
	claims := Claims{
		UserID:         userID,
		Username:       username,
		Role:           role,
		TenantID:       tenantID,
		TenantPlan:     tenantPlan,
		SessionVersion: sessionVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenStr, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
