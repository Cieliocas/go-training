---
description: "Task list for Bank Account Simulation"
---

# Tasks: Bank Account Simulation

**Input**: Design documents from `/specs/001-bank-account-sim/`

**Prerequisites**: plan.md, spec.md, data-model.md, contracts/account-api.md, research.md, quickstart.md

**Tests**: Included. Constitution Principle III (Test-First for Non-Trivial
Logic) requires tests for this exercise (struct with behavior + error
handling), so test tasks are mandatory here, not optional.

**Organization**: Tasks are grouped by user story (from spec.md: US1 = P1
Open Account & Check Balance, US2 = P2 Deposit/Withdraw, US3 = P3 Transfer).

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

This exercise is a single flat Go package at the repository root for this
folder (`Base/14-go-with-speckit/`) — no `src/`/`tests/` split, matching
every other numbered folder in `go-training`. All paths below are relative
to `Base/14-go-with-speckit/`.

**Note on parallelism**: This exercise touches only 4 files total
(`go.mod`, `account.go`, `account_test.go`, `main.go`), and every user
story adds to the *same* `account.go`/`account_test.go`/`main.go` files
rather than new ones. Per the "avoid same-file conflicts" rule, almost no
task here is safely parallelizable — tasks are executed sequentially in
the order listed.

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization

- [X] T001 Initialize a local Go module by running `go mod init bank-account-sim` in `Base/14-go-with-speckit/`, creating `Base/14-go-with-speckit/go.mod` (per research.md: first folder in the repo needing `go test`)

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core `Account` type and errors that every user story depends on

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T002 Define the `Account` struct (unexported `owner string`, `balance int64` fields) and the sentinel errors `ErrInvalidAmount`, `ErrInsufficientFunds`, `ErrNegativeInitialBalance`, `ErrSameAccount` (via `errors.New`, per contracts/account-api.md) in `Base/14-go-with-speckit/account.go`
- [X] T003 Implement `func NewAccount(owner string, initialBalance int64) (*Account, error)` in `Base/14-go-with-speckit/account.go`, returning `ErrNegativeInitialBalance` when `initialBalance < 0` (depends on T002; per data-model.md Validation Rules)

**Checkpoint**: `Account` type and constructor exist — user story implementation can now begin (sequentially, per the parallelism note above)

---

## Phase 3: User Story 1 - Open an Account and Check Balance (Priority: P1) 🎯 MVP

**Goal**: A learner can open an account with an owner and initial balance, then read back the balance and owner without altering state.

**Independent Test**: Call `NewAccount("Alice", 100)`, then call `.Balance()` and `.Owner()` and confirm they return `100` and `"Alice"`.

### Tests for User Story 1

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation (T003's constructor already exists from Foundational, but Balance()/Owner() do not yet)**

- [X] T004 [US1] Write `TestNewAccount` (valid creation, and rejection of a negative initial balance via `errors.Is(err, ErrNegativeInitialBalance)`) in `Base/14-go-with-speckit/account_test.go`
- [X] T005 [US1] Write `TestAccount_BalanceAndOwner` (confirms getters return the values passed to `NewAccount` and that calling them does not change state) in `Base/14-go-with-speckit/account_test.go` (depends on T004 — same file)

### Implementation for User Story 1

- [X] T006 [US1] Implement `func (a *Account) Owner() string` and `func (a *Account) Balance() int64` getter methods in `Base/14-go-with-speckit/account.go` (depends on T003; makes T004/T005 pass)
- [X] T007 [US1] Add a demo block to `func main()` in `Base/14-go-with-speckit/main.go` that opens an account and prints its owner and balance (depends on T006; per quickstart.md Run the demo, step 1)

**Checkpoint**: User Story 1 is fully functional and testable independently — `go test ./... -v` and `go run main.go` both demonstrate opening an account and reading its balance.

---

## Phase 4: User Story 2 - Deposit and Withdraw Funds (Priority: P2)

**Goal**: A learner can deposit into and withdraw from an account, with withdrawals blocked when they would overdraw the account.

**Independent Test**: From a known balance, deposit and confirm the exact increase; attempt an over-withdrawal and confirm it's rejected with `ErrInsufficientFunds` and the balance unchanged; then withdraw a valid amount and confirm the exact decrease.

### Tests for User Story 2

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T008 [US2] Write `TestAccount_Deposit` (successful deposit increases balance by exact amount; zero/negative amount rejected via `errors.Is(err, ErrInvalidAmount)` with balance unchanged) in `Base/14-go-with-speckit/account_test.go` (depends on T005 — same file)
- [ ] T009 [US2] Write `TestAccount_Withdraw` (successful withdrawal decreases balance by exact amount; over-withdrawal rejected via `errors.Is(err, ErrInsufficientFunds)` with balance unchanged; zero/negative amount rejected via `ErrInvalidAmount`) in `Base/14-go-with-speckit/account_test.go` (depends on T008 — same file)

### Implementation for User Story 2

- [ ] T010 [US2] Implement `func (a *Account) Deposit(amount int64) error` in `Base/14-go-with-speckit/account.go` (depends on T006; makes T008 pass; per data-model.md State Transitions)
- [ ] T011 [US2] Implement `func (a *Account) Withdraw(amount int64) error` in `Base/14-go-with-speckit/account.go` (depends on T010 — same file; makes T009 pass)
- [ ] T012 [US2] Extend `func main()` in `Base/14-go-with-speckit/main.go` with a deposit, a rejected over-withdrawal, and a successful withdrawal, printing the balance after each step (depends on T011 and T007; per quickstart.md Run the demo, step 2)

**Checkpoint**: User Stories 1 AND 2 both work independently — deposits/withdrawals are validated and idiomatic-error-checked in tests and demonstrated in the demo.

---

## Phase 5: User Story 3 - Transfer Between Accounts (Priority: P3)

**Goal**: A learner can move funds from one account to another as a single all-or-nothing operation.

**Independent Test**: From two accounts with known balances, transfer a valid amount and confirm the source decreased and destination increased by exactly that amount; attempt a transfer exceeding the source balance and confirm both balances stay unchanged.

### Tests for User Story 3

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T013 [US3] Write `TestAccount_Transfer` (successful transfer moves the exact amount between two accounts; over-transfer rejected via `errors.Is(err, ErrInsufficientFunds)` with both balances unchanged; transfer to self rejected via `errors.Is(err, ErrSameAccount)`) in `Base/14-go-with-speckit/account_test.go` (depends on T009 — same file)

### Implementation for User Story 3

- [ ] T014 [US3] Implement `func (a *Account) Transfer(to *Account, amount int64) error` in `Base/14-go-with-speckit/account.go`, returning `ErrSameAccount` when `to == a`, otherwise withdrawing from `a` and depositing into `to` as an all-or-nothing operation (depends on T011; makes T013 pass; per data-model.md State Transitions)
- [ ] T015 [US3] Extend `func main()` in `Base/14-go-with-speckit/main.go` with a two-account transfer demo, printing both balances before and after (depends on T014 and T012; per quickstart.md Run the demo, step 3)

**Checkpoint**: All user stories are independently functional — the full simulation (open, deposit, withdraw, transfer) works end to end.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Final quality pass across the whole feature

- [ ] T016 Add a one-line doc comment above the `Account` type or in a top-of-file comment in `Base/14-go-with-speckit/account.go` recording the key takeaway (pointer receivers for mutation; sentinel errors for expected failure conditions), per constitution Principle V (Documented Learnings)
- [ ] T017 Run `gofmt -l .` and `go vet ./...` in `Base/14-go-with-speckit/` and fix any issues found
- [ ] T018 Run `go test ./... -v` in `Base/14-go-with-speckit/` and manually walk through quickstart.md's "Validating success criteria manually" section against the `go run main.go` output

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies — start immediately.
- **Foundational (Phase 2)**: Depends on Setup (T001) — BLOCKS all user stories.
- **User Story 1 (Phase 3)**: Depends on Foundational (T002, T003) only.
- **User Story 2 (Phase 4)**: Depends on User Story 1's `Balance()` getter (T006) to observe results, and reuses `account_test.go` from US1 (T005) — implement after US1.
- **User Story 3 (Phase 5)**: Depends on User Story 2's `Withdraw`/`Deposit` (T010, T011), since `Transfer` composes them — implement after US2.
- **Polish (Phase 6)**: Depends on all user stories being complete.

### Within Each User Story

- Tests MUST be written and FAIL before implementation (constitution Principle III).
- Getter/constructor tasks before mutating-method tasks.
- `account.go` implementation before the corresponding `main.go` demo block.

### Parallel Opportunities

None of substance: this is an intentionally small, single-package exercise
(4 files total) where every task after T001 touches one of `account.go`,
`account_test.go`, or `main.go`, and each story builds on the previous
story's methods (`Transfer` needs `Withdraw`/`Deposit`; `Withdraw`/`Deposit`
need the constructor and getters). Execute T001 → T018 in order.

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (T001).
2. Complete Phase 2: Foundational (T002–T003).
3. Complete Phase 3: User Story 1 (T004–T007).
4. **STOP and VALIDATE**: `go test ./... -v` and `go run main.go` — confirm an account can be opened and its balance read back correctly.

### Incremental Delivery

1. Setup + Foundational → foundation ready (T001–T003).
2. Add User Story 1 → test independently (T004–T007) → MVP.
3. Add User Story 2 → test independently (T008–T012).
4. Add User Story 3 → test independently (T013–T015).
5. Polish (T016–T018).

## Notes

- [Story] label maps each task to its user story for traceability.
- Commit after each task or logical group.
- Stop at any checkpoint to validate that story's behavior independently via `go test` and `go run main.go`.
- Avoid: vague tasks, same-file conflicts, cross-story dependencies that break independent testability of the *behavior* (even though the same files are shared, each story's tests/behavior remain independently verifiable).
