package main

import "errors"

// Account balances are mutated exclusively through pointer-receiver methods,
// and expected failure conditions (overdraft, invalid amounts) are reported
// as sentinel errors rather than panics — the two idioms this exercise
// practices.
type Account struct {
	owner   string
	balance int64
}

var (
	ErrInvalidAmount          = errors.New("amount must be positive")
	ErrInsufficientFunds      = errors.New("insufficient funds")
	ErrNegativeInitialBalance = errors.New("initial balance cannot be negative")
	ErrSameAccount            = errors.New("cannot transfer to the same account")
)

// NewAccount creates an account with the given owner and initial balance.
func NewAccount(owner string, initialBalance int64) (*Account, error) {
	if initialBalance < 0 {
		return nil, ErrNegativeInitialBalance
	}
	return &Account{owner: owner, balance: initialBalance}, nil
}

// Owner returns the account holder's name.
func (a *Account) Owner() string {
	return a.owner
}

// Balance returns the current balance.
func (a *Account) Balance() int64 {
	return a.balance
}
