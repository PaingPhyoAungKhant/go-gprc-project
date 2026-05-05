package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Payload implement jwt.Claim Interface and contain the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"jti"`
	Username  string    `json:"sub"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
}

// NewPayload creates a new token payload with specific username and duration
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

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredAt), nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedAt), nil
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (payload *Payload) GetIssuer() (string, error) {
	return "", nil
}

func (payload *Payload) GetSubject() (string, error) {
	return payload.Username, nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrTokenExpired
	}
	return nil
}
