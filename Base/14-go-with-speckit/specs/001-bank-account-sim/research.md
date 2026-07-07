# Phase 0 Research: Bank Account Simulation

All Technical Context items were resolvable from repo conventions and the
user's explicit direction — no `NEEDS CLARIFICATION` markers remain.

## Decision: Standard library only, no third-party dependencies

- **Rationale**: Constitution Principle II (Idiomatic, Simple Go) requires
  favoring the standard library over frameworks. `fmt`, `errors`, and
  `testing` fully cover this exercise's needs (printing, sentinel errors,
  unit tests).
- **Alternatives considered**: A third-party assertion library (e.g.
  `testify`) for tests — rejected; the standard `testing` package with
  table-driven tests is sufficient and keeps the exercise focused on Go
  fundamentals rather than tooling.

## Decision: Add a local `go.mod` for this folder

- **Rationale**: This is the first folder in the repo with an accompanying
  `_test.go` file. Modern Go tooling (`go test`) requires module context;
  prior folders only ever used `go run main.go` and never needed one.
- **Alternatives considered**: Relying on legacy GOPATH-mode compilation —
  rejected as deprecated and inconsistent with the Go 1.21+ target already
  declared in the repo's root `README.md`.

## Decision: Idiomatic error handling via sentinel errors + `errors.Is`

- **Rationale**: The user explicitly asked for idiomatic Go error handling
  for insufficient-funds withdrawals. Returning a distinct sentinel error
  per failure condition (invalid amount, insufficient funds, negative
  initial balance, same-account transfer) lets callers use `errors.Is` to
  branch on the specific cause, which is the standard Go pattern for
  expected/recoverable error conditions.
- **Alternatives considered**:
  - `panic`/`recover` — rejected; panics are idiomatically reserved for
    programmer errors and unrecoverable states, not expected business
    conditions like an overdrawn account.
  - A single boolean `ok` return — rejected; it can't distinguish *why* an
    operation failed, which matters for both callers and the tests in
    `account_test.go`.

## Decision: Balance mutation only through pointer-receiver methods

- **Rationale**: The user explicitly asked for pointer receivers, and it
  matches this exercise's learning goal — a value-receiver method would
  silently fail to persist balance changes to the caller's `Account`,
  which is exactly the pitfall this exercise is meant to teach.
- **Alternatives considered**: Value-receiver methods returning a new,
  updated `Account` — rejected; it avoids pointers entirely and does not
  exercise the intended learning goal.

## Decision: `balance` stored as `int64` (whole currency units)

- **Rationale**: All spec examples use whole numbers (100, 50, 150). Using
  an integer type avoids floating-point rounding entirely — sidestepping a
  classic (and idiomatically well-known) pitfall of using `float64` for
  money — without adding the complexity of a cents-based or
  arbitrary-precision decimal type, which is out of scope for this
  exercise.
- **Alternatives considered**: `float64` — rejected, introduces rounding
  error risk for a feature whose success criteria demand *exact* balance
  changes (SC-001, SC-003). A cents-scaled integer or `big.Rat`/decimal
  library — rejected as unnecessary complexity beyond this exercise's
  scope (no fractional-currency requirement in the spec).
