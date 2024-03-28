package dto

import "time"

type TaskCreate struct {
	Tag         string `json:"tag"`
	Sprint      int    `json:"sprint"`
	Description string `json:"description"`
}

type TaskUpdate struct {
	Tag         string    `json:"tag"`
	Sprint      int       `json:"sprint"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}
