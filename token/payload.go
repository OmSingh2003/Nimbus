package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different types of error returned by verifying token function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// GetExpirationTime returns the expiration time for JWT Claims interface
func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}

// GetIssuedAt returns the issued at time for JWT Claims interface
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}

// GetNotBefore returns nil for JWT Claims interface
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

// GetIssuer returns empty string for JWT Claims interface
func (payload *Payload) GetIssuer() (string, error) {
	return "", nil
}

// GetSubject returns the username for JWT Claims interface
func (payload *Payload) GetSubject() (string, error) {
	return payload.Username, nil
}

// GetAudience returns nil for JWT Claims interface
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}
