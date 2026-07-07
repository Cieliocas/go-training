# Feature Specification: Bank Account Simulation

**Feature Branch**: `001-bank-account-sim`

**Created**: 2026-07-04

**Status**: Draft

**Input**: User description: "build a simple bank account simulation using pointers and methods."

**Learning Goal**: Practice defining a struct that represents real-world state
(an account balance) and mutating it safely through methods with pointer
receivers, instead of passing values around by copy.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Open an Account and Check Balance (Priority: P1)

A learner opens a new account for an owner with a starting balance, then
checks the current balance to confirm it was recorded correctly.

**Why this priority**: This is the foundation every other operation depends
on — without a working account and balance check, deposits, withdrawals, and
transfers can't be verified.

**Independent Test**: Create an account with an owner name and an initial
balance, then read back the balance — it must match what was provided at
creation.

**Acceptance Scenarios**:

1. **Given** no existing account, **When** a new account is opened for
   "Alice" with an initial balance of 100, **Then** the account's balance is
   100 and its owner is "Alice".
2. **Given** an open account, **When** its balance is checked, **Then** the
   current balance is returned without changing it.

---

### User Story 2 - Deposit and Withdraw Funds (Priority: P2)

A learner deposits money into an account and withdraws money from it,
seeing the balance update correctly, and sees withdrawals blocked when funds
are insufficient.

**Why this priority**: Deposits and withdrawals are the core balance-mutating
operations and the main place pointer receivers matter (the change must be
visible to the caller after the method returns).

**Independent Test**: Starting from a known balance, deposit an amount and
verify the balance increased by exactly that amount; then attempt to
withdraw more than the balance and verify it is rejected and the balance is
unchanged; then withdraw a valid amount and verify the balance decreased by
exactly that amount.

**Acceptance Scenarios**:

1. **Given** an account with balance 100, **When** 50 is deposited,
   **Then** the balance becomes 150.
2. **Given** an account with balance 100, **When** a withdrawal of 150 is
   requested, **Then** the withdrawal is rejected and the balance remains
   100.
3. **Given** an account with balance 100, **When** a withdrawal of 40 is
   requested, **Then** the withdrawal succeeds and the balance becomes 60.

---

### User Story 3 - Transfer Between Accounts (Priority: P3)

A learner moves funds from one account to another in a single operation,
seeing both balances update consistently.

**Why this priority**: Transfers combine withdraw-then-deposit logic across
two separate struct instances, which is good practice for working with
multiple pointers at once, but the simulation is still useful without it.

**Independent Test**: Starting from two accounts with known balances,
transfer an amount from one to the other and verify the source balance
decreased and the destination balance increased by exactly that amount, and
that a transfer exceeding the source balance is rejected and leaves both
balances unchanged.

**Acceptance Scenarios**:

1. **Given** account A with balance 100 and account B with balance 20,
   **When** 30 is transferred from A to B, **Then** A's balance is 70 and
   B's balance is 50.
2. **Given** account A with balance 100 and account B with balance 20,
   **When** 150 is transferred from A to B, **Then** the transfer is
   rejected and both balances remain unchanged.

---

### Edge Cases

- What happens when a deposit or withdrawal amount is zero or negative? The
  operation MUST be rejected and the balance MUST remain unchanged.
- What happens when a transfer's source and destination account are the
  same account? The operation MUST be rejected (or treated as a no-op) with
  the balance unchanged either way.
- What happens when an account is opened with a negative initial balance?
  The operation MUST be rejected.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST allow opening a new account with an owner name and
  a non-negative initial balance.
- **FR-002**: System MUST allow depositing a positive amount into an
  account, increasing its balance by exactly that amount.
- **FR-003**: System MUST reject deposits of zero or negative amounts,
  leaving the balance unchanged.
- **FR-004**: System MUST allow withdrawing an amount up to and including
  the current balance, decreasing the balance by exactly that amount.
- **FR-005**: System MUST reject withdrawals that exceed the current
  balance or that are zero/negative, leaving the balance unchanged in both
  cases.
- **FR-006**: System MUST allow transferring a valid amount from one
  account to a different account, decreasing the source balance and
  increasing the destination balance by exactly that amount.
- **FR-007**: System MUST reject a transfer if the source account has
  insufficient balance, leaving both accounts' balances unchanged.
- **FR-008**: System MUST report an account's current balance and owner on
  request without altering its state.

### Key Entities

- **Account**: Represents a single bank account. Key attributes: owner name,
  current balance. Balance must never go negative as a result of any
  operation.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 100% of valid deposit and withdrawal operations change the
  account balance by exactly the requested amount.
- **SC-002**: 100% of withdrawal or transfer attempts that exceed the
  source account's balance are rejected without any change to either
  account's balance.
- **SC-003**: For any sequence of valid transfers between two accounts, the
  combined total balance of both accounts stays exactly the same before and
  after (funds are neither created nor destroyed).
- **SC-004**: A learner can verify every scenario above by only observing
  reported balances — no inspection of internal state is required.

## Assumptions

- Single in-memory simulation for one program run; no persistence to disk
  or a database is required.
- Single currency, whole-number or standard decimal amounts; no
  multi-currency conversion is in scope.
- No concurrent/simultaneous access to the same account is required for
  this exercise (single-threaded use).
- No authentication, authorization, or multi-user access control is
  required — this is a local practice simulation, not a real banking
  system.
- Transaction history/audit log is out of scope; only current balance needs
  to be observable.
