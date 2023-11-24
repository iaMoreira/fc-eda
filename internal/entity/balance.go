package entity

import "time"

type Balance struct {
	AccountIDFrom      string
	AccountIDTo        string
	BalanceAccountFrom float64
	BalanceAccountTo   float64
	CreatedAt          time.Time
}

func NewBalance(client *Balance) *Balance {
	if client == nil {
		return nil
	}
	Balance := &Balance{
		CreatedAt: time.Now(),
	}
	return Balance
}
