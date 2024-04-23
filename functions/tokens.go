package functions

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Polilo-User/buildings/functions/errors"

	"github.com/dgrijalva/jwt-go"
)

// Параметры jwt-токена
const (
	SECRET_KEY        = "33446a9dcf9ea060a0a6532b166da32f303af0de"
	ACCESS_TOKEN_TTL  = "24h"
	REFRESH_TOKEN_TTL = "72h"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type TokenManager interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.Unauthorized.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func (m *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

// ф-я дешифровки токена
func BearerAuth(r *http.Request) (userID string, err error) {
	header := r.Header.Get("Authorization")
	if len(header) == 0 {
		return "", errors.Unauthorized.New("empty auth header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.Unauthorized.New("invalid auth header")
	}
	if len(headerParts[1]) == 0 {
		return "", errors.Unauthorized.New("token is empty")
	}
	manag, err := NewManager(SECRET_KEY)
	if err != nil {
		return "", errors.Unauthorized.New(err.Error())
	}
	id, err := manag.Parse(headerParts[1])
	if err != nil {
		return "", errors.Unauthorized.New("invalid auth token")
	}

	return id, nil
}
