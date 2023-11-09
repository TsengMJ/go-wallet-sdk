package core

import (
	"math/big"
)

const (
	TransactionTypeWithdraw   TransactionType = "Withdraw"
	TransactionTypeForcedExit TransactionType = "ForcedExit"
)

type Withdraw struct {
	Type      string     `json:"type"`
	AccountId uint32     `json:"accountId"`
	From      string     `json:"from"`
	To        string     `json:"to"`
	TokenId   uint32     `json:"token"`
	Amount    *big.Int   `json:"amount"`
	Fee       string     `json:"fee"`
	Nonce     uint32     `json:"nonce"`
	Signature *Signature `json:"signature"`
	*TimeRange
}

func (t *Withdraw) getType() string {
	return "Withdraw"
}

type ForcedExit struct {
	Type      string     `json:"type"`
	AccountId uint32     `json:"initiatorAccountId"`
	Target    string     `json:"target"`
	TokenId   uint32     `json:"token"`
	Amount    *big.Int   `json:"amount"`
	Fee       string     `json:"fee"`
	Nonce     uint32     `json:"nonce"`
	Signature *Signature `json:"signature"`
	*TimeRange
}

func (t *ForcedExit) getType() string {
	return "ForcedExit"
}