# Contract: Account Package API

This exercise is a Go package, not a web service — its "contract" is the
exported Go API surface that `main.go` and `account_test.go` depend on.

## Package

`package main` (matches this repo's single-package-per-folder convention;
no internal `src/` split).

## Type

```go
type Account struct {
    // unexported fields: owner string, balance int64
}
```

## Sentinel Errors

```go
var (
    ErrInvalidAmount         = errors.New("amount must be positive")
    ErrInsufficientFunds     = errors.New("insufficient funds")
    ErrNegativeInitialBalance = errors.New("initial balance cannot be negative")
    ErrSameAccount           = errors.New("cannot transfer to the same account")
)
```

Callers distinguish failure reasons with `errors.Is(err, ErrInsufficientFunds)`, etc.

## Functions & Methods

| Signature | Behavior | Spec Ref |
|-----------|----------|----------|
| `func NewAccount(owner string, initialBalance int64) (*Account, error)` | Creates an account. Returns `ErrNegativeInitialBalance` if `initialBalance < 0`. | FR-001 |
| `func (a *Account) Owner() string` | Returns the owner name. Never mutates state. | FR-008 |
| `func (a *Account) Balance() int64` | Returns the current balance. Never mutates state. | FR-008 |
| `func (a *Account) Deposit(amount int64) error` | Adds `amount` to balance. Returns `ErrInvalidAmount` if `amount <= 0`. | FR-002, FR-003 |
| `func (a *Account) Withdraw(amount int64) error` | Subtracts `amount` from balance. Returns `ErrInvalidAmount` if `amount <= 0`; returns `ErrInsufficientFunds` if `amount > balance`. Balance unchanged on error. | FR-004, FR-005 |
| `func (a *Account) Transfer(to *Account, amount int64) error` | Withdraws `amount` from `a`, deposits into `to`. Returns `ErrSameAccount` if `to == a`; propagates `Withdraw`'s errors. All-or-nothing: on any error, neither balance changes. | FR-006, FR-007 |

All mutating methods use pointer receivers (`*Account`) so balance changes
are visible to the caller — this is the exercise's core learning goal.

## Usage Example (illustrative, not implementation)

```go
alice, err := NewAccount("Alice", 100)
if err != nil { /* handle */ }

if err := alice.Deposit(50); err != nil { /* handle */ }
fmt.Println(alice.Balance()) // 150

bob, _ := NewAccount("Bob", 20)
if err := alice.Transfer(bob, 200); errors.Is(err, ErrInsufficientFunds) {
    fmt.Println("transfer rejected: insufficient funds")
}
```
