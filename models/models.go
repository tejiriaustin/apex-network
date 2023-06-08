package models

import (
	"github.com/google/uuid"
	"time"
)

type TransactionType string

type TransactionDescription string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

const (
	FundWallet TransactionDescription = "fund-wallet"
	StartGame  TransactionDescription = "start-game"
	LoseRoll   TransactionDescription = "lose-roll"
	WinRoll    TransactionDescription = "win-roll"
)

type (
	Shared struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}
	User struct {
		Shared
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		FullName      string `json:"full_name"`
		IsPlaying     bool   `json:"is_playing"`
		WalletBalance int    `json:"wallet_balance"`
	}
	Game struct {
		Shared
		UserId         string `json:"user_id"`
		TargetNumber   int    `json:"target_number"`
		FirstDiceRoll  int    `json:"first_dice_roll"`
		SecondDiceRoll int    `json:"second_dice_roll"`
		IsComplete     bool   `json:"is_complete"`
	}
	WalletTransaction struct {
		Shared
		Description     string `json:"description"`
		TransactionType string `json:"transaction_type"`
	}
)

func (s Shared) Init() {
	id := uuid.New()
	s.ID = id
	s.CreatedAt = time.Now().UTC()
}
