package session

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

type Session struct {
	ID          string
	UserID      string
	CsrfToken   string
	Expires     time.Time
	MaxDuration time.Duration
}

func NewSession(userID string) (*Session, error) {
	sessionID := generateSessionID(userID)
	csrfToken := generateCsrfToken(sessionID)
	return &Session{
		ID:          sessionID,
		UserID:      userID,
		CsrfToken:   csrfToken,
		Expires:     time.Now().Add(24 * time.Hour),
		MaxDuration: 24 * time.Hour,
	}, nil
}

func generateSessionID(userID string) string {
	hash := sha256.Sum256([]byte(userID))
	sessionID := base64.StdEncoding.EncodeToString(hash[:])
	return sessionID
}

func generateCsrfToken(sessionID string) string {
	hash := sha256.Sum256([]byte(sessionID))
	token := base64.URLEncoding.EncodeToString(hash[:])
	return token
}
