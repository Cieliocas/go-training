# Quickstart: Bank Account Simulation

## Prerequisites

- Go 1.21+ installed (`go version`)
- Working directory: `Base/14-go-with-speckit/`

## Setup

```bash
cd Base/14-go-with-speckit
go mod init bank-account-sim   # only if go.mod does not already exist
```

## Run the tests

```bash
go test ./... -v
```

**Expected outcome**: all tests pass, covering the scenarios in
[spec.md](./spec.md) — deposit/withdraw success and failure paths
(insufficient funds, invalid amount), and transfer success/failure
(insufficient funds, same-account), per the contract in
[contracts/account-api.md](./contracts/account-api.md).

## Run the demo

```bash
go run main.go
```

**Expected outcome**: console output demonstrating, in order:
1. Opening an account and printing its initial balance (User Story 1).
2. A successful deposit and a rejected over-withdrawal, then a successful
   withdrawal, each with the resulting balance printed (User Story 2).
3. A transfer between two accounts, printing both balances before and
   after (User Story 3).

## Validating success criteria manually

- **SC-001/SC-002**: Balances printed after each operation in the demo
  output should change by exactly the requested amount on success, and
  stay identical on any rejected operation.
- **SC-003**: Sum the two transfer-demo balances before and after the
  transfer — the totals must match.
- **SC-004**: All of the above is verifiable from printed balances alone,
  with no need to read `account.go`'s internals.
