package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoTokenGenerator struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoTokenGenerator(symmetricKey string) (TokenGenerator, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	gen := &PasetoTokenGenerator{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return gen, nil
}

// GenerateToken creates a new token for a specific username and duration
func (gen *PasetoTokenGenerator) GenerateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := gen.paseto.Encrypt(gen.symmetricKey, payload, nil)
	return token, payload, err
}

// ValidateToken checks if the token is valid or not
func (gen *PasetoTokenGenerator) ValidateToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := gen.paseto.Decrypt(token, gen.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
