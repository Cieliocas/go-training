# Phase 1 Data Model: Bank Account Simulation

## Entity: Account

Represents a single bank account (spec Key Entities section).

| Field     | Type     | Notes                                                              |
|-----------|----------|---------------------------------------------------------------------|
| `owner`   | `string` | Non-empty display name, set at creation, immutable afterward       |
| `balance` | `int64`  | Whole currency units; must never be negative (see Validation Rules) |

Both fields are unexported; external code reads them only via `Owner()` and
`Balance()` methods, and mutates `balance` only via `Deposit`, `Withdraw`,
and `Transfer` (FR-008 — state must be observable but not directly
alterable from outside).

### Validation Rules

- **Creation** (`NewAccount`): `owner` must be non-empty; `initialBalance`
  must be `>= 0` → else `ErrNegativeInitialBalance` (spec Edge Case: negative
  initial balance).
- **Deposit**: `amount` must be `> 0` → else `ErrInvalidAmount` (FR-003).
- **Withdraw**: `amount` must be `> 0` → else `ErrInvalidAmount`; and
  `amount <= balance` → else `ErrInsufficientFunds` (FR-005).
- **Transfer**: same amount validation as Withdraw applied to the source
  account; additionally source and destination must be different accounts
  → else `ErrSameAccount` (spec Edge Case: same-account transfer).

### State Transitions

- `Deposit(amount)`: `balance = balance + amount` (only if amount valid).
- `Withdraw(amount)`: `balance = balance - amount` (only if amount valid and
  `amount <= balance`).
- `Transfer(to, amount)`: equivalent to `from.Withdraw(amount)` followed by
  `to.Deposit(amount)`, but must be all-or-nothing — if the withdrawal step
  fails, the deposit step MUST NOT occur and neither balance changes
  (SC-002, SC-003: conservation of funds).

No other entities are required — the spec's scope is a single account type
with no transaction history/audit log (per spec Assumptions).
