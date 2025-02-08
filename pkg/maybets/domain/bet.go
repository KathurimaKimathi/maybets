package domain

import (
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/application/enums"
)

type Bet struct {
	BetID     string        `json:"bet_id"`
	UserID    string        `json:"user_id"`
	Amount    float64       `json:"amount"`
	Odds      float64       `json:"odds"`
	Outcome   enums.Outcome `json:"outcome"`
	Timestamp time.Time     `json:"timestamp"`
}

type User struct {
	ID            string  `json:"id"`
	TotalBets     int64   `json:"total_bets,omitempty"`
	TotalWinnings float64 `json:"winnings,omitempty"`
}
