package models

import "github.com/google/uuid"

type User struct {
	Id            uuid.UUID `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	FullName      string    `json:"full_name"`
	IsPlaying     bool      `json:"is_playing"`
	WalletBalance int       `json:"wallet_balance"`
}

type Game struct {
}
