package token

import "time"

type TokenGenerator interface {
	GenerateToken(string, time.Duration) (string, *Payload, error)
	ValidateToken(string) (*Payload, error)
}
