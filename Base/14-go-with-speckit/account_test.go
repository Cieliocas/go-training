package main

import (
	"errors"
	"testing"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("Alice", 100)
	if err != nil {
		t.Fatalf("unexpected error creating account: %v", err)
	}
	if acc.Owner() != "Alice" {
		t.Errorf("owner = %q, want %q", acc.Owner(), "Alice")
	}
	if acc.Balance() != 100 {
		t.Errorf("balance = %d, want %d", acc.Balance(), 100)
	}

	_, err = NewAccount("Bob", -10)
	if !errors.Is(err, ErrNegativeInitialBalance) {
		t.Errorf("NewAccount with negative balance: got err %v, want ErrNegativeInitialBalance", err)
	}
}

func TestAccount_BalanceAndOwner(t *testing.T) {
	acc, err := NewAccount("Alice", 100)
	if err != nil {
		t.Fatalf("unexpected error creating account: %v", err)
	}

	if got := acc.Balance(); got != 100 {
		t.Errorf("Balance() = %d, want %d", got, 100)
	}
	if got := acc.Owner(); got != "Alice" {
		t.Errorf("Owner() = %q, want %q", got, "Alice")
	}

	// Reading balance/owner again must not change state.
	if got := acc.Balance(); got != 100 {
		t.Errorf("Balance() after repeated read = %d, want %d", got, 100)
	}
}
