package models

import (
	"fmt"
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
	RollCost   TransactionDescription = "lose-roll"
	WinRoll    TransactionDescription = "win-roll"
)

const (
	FieldPlayerBalance            = "balance"
	FieldPlayerIsPlaying          = "is_playing"
	FieldPlayerTargetNumber       = "target_number"
	FieldPlayerHasRolledFirstDice = "has_rolled_first_dice"
)

type (
	Shared struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}
	Player struct {
		Shared
		FirstName         string `json:"first_name"`
		LastName          string `json:"last_name"`
		FullName          string `json:"full_name"`
		IsPlaying         bool   `json:"is_playing" gorm:"column:is_playing;type:boolean;default:false"`
		WalletBalance     int    `json:"wallet_balance"`
		TargetNumber      int    `json:"target_number"`
		DiceSum           int    `json:"dice_sum"`
		HasRolledFirstDie bool   `json:"has_rolled"`
	}
	Game struct {
		Shared
	}
	WalletTransaction struct {
		Shared
		PlayerId        uuid.UUID              `json:"player_id"`
		Amount          int                    `json:"amount"`
		Description     TransactionDescription `json:"description"`
		TransactionType TransactionType        `json:"transaction_type"`
	}
)

func (s Shared) Init() {
	s.ID = uuid.New()
	fmt.Println(s)
	s.CreatedAt = time.Now().UTC()
}

func (u Player) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (Player) TableName() string {
	return "Players"
}

func (WalletTransaction) TableName() string {
	return "wallet_transaction"
}
