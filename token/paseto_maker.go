package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

// JWTMaker is a JSON Web Token Maker
type PasetoMaker struct {
	paseto        *paseto.V2
	symemetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exectly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:        paseto.NewV2(),
		symemetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// CreateToken creates a new token for specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symemetricKey, payload, nil)
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(tokenStr string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(tokenStr, maker.symemetricKey, payload, nil)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
