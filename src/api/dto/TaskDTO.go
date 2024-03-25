package dto

import "time"

type TaskCreateDTO struct {
	Tag         string `json:"tag"`
	Sprint      int    `json:"sprint"`
	Description string `json:"description"`
}

type TaskUpdateDTO struct {
	Tag         string    `json:"tag"`
	Sprint      int       `json:"sprint"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}
