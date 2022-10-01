package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates and returns a signed token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken returns payload stored inside body of the token if it is valid
	VerifyToken(token string) (*Payload, error)
}
