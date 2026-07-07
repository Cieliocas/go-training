# Implementation Plan: Bank Account Simulation

**Branch**: `001-bank-account-sim` | **Date**: 2026-07-04 | **Spec**: [spec.md](./spec.md)

**Input**: Feature specification from `/specs/001-bank-account-sim/spec.md`

## Summary

Build a small in-memory bank account simulation in Go that lets a learner
open accounts, deposit, withdraw, and transfer funds, with all balance
mutation done through pointer-receiver methods on an `Account` struct and
insufficient-funds/invalid-amount conditions surfaced as idiomatic Go
errors (sentinel errors checked with `errors.Is`), per the user's explicit
direction for this plan.

## Technical Context

**Language/Version**: Go 1.21+ (matches repo root `README.md`)

**Primary Dependencies**: Go standard library only (`fmt`, `errors`,
`testing`) — no third-party packages, per constitution Principle II
(Idiomatic, Simple Go)

**Storage**: N/A — single in-memory run, no persistence (per spec
Assumptions)

**Testing**: `go test`, standard library `testing` package

**Target Platform**: Local CLI, run via `go run main.go` on any OS with Go
installed

**Project Type**: Single-file learning exercise (CLI demo), not a
library/service

**Performance Goals**: N/A — not a performance-sensitive exercise; correct
balance behavior matters, not throughput

**Constraints**: No external dependencies; no concurrency (single-threaded
use only, per spec Assumptions); no persistence layer

**Scale/Scope**: A handful of `Account` values and a few dozen operations
in a demo `main()`; scope is deliberately small (learning exercise, not
production sizing)

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Check | Status |
|-----------|-------|--------|
| I. Spec-First Learning | `spec.md` written and validated before this plan | PASS |
| II. Idiomatic, Simple Go | Plain struct + pointer-receiver methods + stdlib `errors`; no frameworks, no generics, no third-party deps | PASS |
| III. Test-First for Non-Trivial Logic | Struct with behavior + error handling is non-trivial → `account_test.go` is a required deliverable, written alongside implementation, covering deposit/withdraw/transfer/error paths | PASS |
| IV. Incremental Complexity | Builds on prior folders (07-ponteiros, 09-struct, 10-errors, 12-arrays/13-slice); does not require concurrency or interfaces (later roadmap stages) | PASS |
| V. Documented Learnings | `main.go` and this plan record the learning goal (pointer receivers + idiomatic error handling); a one-line takeaway will be added once implemented | PASS |

No violations — Complexity Tracking table is not needed.

## Project Structure

### Documentation (this feature)

```text
specs/001-bank-account-sim/
├── plan.md              # This file (/speckit-plan command output)
├── research.md          # Phase 0 output (/speckit-plan command)
├── data-model.md         # Phase 1 output (/speckit-plan command)
├── quickstart.md         # Phase 1 output (/speckit-plan command)
├── contracts/            # Phase 1 output (/speckit-plan command)
│   └── account-api.md
└── tasks.md              # Phase 2 output (/speckit-tasks command - NOT created by /speckit-plan)
```

### Source Code (repository root)

This folder (`Base/14-go-with-speckit/`) is itself the exercise's root —
consistent with every other numbered folder in `go-training` (single
`main.go`, no `src/`/`tests/` split). This feature adds a local Go module
so `go test` can run, plus two new files alongside `main.go`:

```text
Base/14-go-with-speckit/
├── go.mod              # NEW: local module (required for `go test`; first
│                        #      folder in the repo that needs it)
├── main.go             # UPDATED: demo entry point exercising the simulation
├── account.go           # NEW: Account struct + pointer-receiver methods
└── account_test.go      # NEW: unit tests (constitution Principle III)
```

**Structure Decision**: Single project, flat layout, matching this repo's
existing per-folder convention (no `src/`/`tests/` subdirectories). The only
structural addition versus prior folders is `go.mod`, needed because this is
the first exercise with an accompanying test file.

## Complexity Tracking

*No entries — Constitution Check reported no violations.*
