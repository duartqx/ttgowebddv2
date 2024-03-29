package dto

import "time"

type Login struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	Status    bool      `json:"status"`
}
