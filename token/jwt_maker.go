package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token Maker
type JWTMaker struct {
	secretKey string
}

var (
	ErrTokenExpired = errors.New("token expired")
	ErrTokenInvalid = errors.New("invalid token")
)

func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid sevret key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

// CreateToken creates a new token for specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(tokenStr string) (*Payload, error) {
	fn := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", token.Header["alg"])
		}
		return []byte(maker.secretKey), nil
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, fn)
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrTokenExpired
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenInvalidClaims):
			return nil, ErrTokenExpired
		default:
			return nil, ErrTokenInvalid
		}
	}

	payload, ok := token.Claims.(*Payload)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return payload, nil
}
