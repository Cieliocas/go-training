<!--
Sync Impact Report
==================
Version change: (template) → 1.0.0
Modified principles: N/A (initial ratification, no prior version)
Added sections:
  - Core Principles: I. Spec-First Learning, II. Idiomatic, Simple Go,
    III. Test-First for Non-Trivial Logic, IV. Incremental Complexity,
    V. Documented Learnings
  - Spec-Driven Workflow (Section 2)
  - Scope & Constraints (Section 3)
  - Governance
Removed sections: none (all template placeholders replaced)
Templates requiring updates:
  - .specify/templates/plan-template.md ✅ compatible as-is (generic Constitution
    Check gate references this file; no principle-specific edits needed)
  - .specify/templates/spec-template.md ✅ compatible as-is
  - .specify/templates/tasks-template.md ✅ compatible as-is (Tests remain
    OPTIONAL per template; Principle III narrows when they're expected)
  - .specify/templates/commands/*.md — no agent-specific references found requiring update
Follow-up TODOs: none
-->

# 14-go-with-speckit Constitution

<!-- This is folder 14 of the personal `go-training` repository: a Go
fundamentals learning project, practiced using the Spec-Driven Development
(speckit) workflow itself as the learning method. -->

## Core Principles

### I. Spec-First Learning

Every exercise or feature in this folder MUST start with a spec (via
`/speckit-specify`) before any code is written. The spec MUST state the
learning goal, not just the desired output (e.g., "practice interface
composition" rather than only "build a shape calculator"). Skipping the spec
step is only acceptable for throwaway one-line syntax checks that are deleted
immediately after.

**Rationale**: The point of this folder is deliberate practice, not just
working code. Writing the spec first forces clarity about what concept is
being learned before diving into syntax.

### II. Idiomatic, Simple Go

Code MUST favor plain Go idioms and the standard library over cleverness,
external frameworks, or premature abstraction. Generics, advanced patterns,
or third-party packages MUST NOT be introduced until the plain/idiomatic
approach has been demonstrated first. No organizational-only packages or
speculative interfaces.

**Rationale**: As a learner, the priority is internalizing Go fundamentals
(types, errors, structs, pointers) as taught idiomatically — not showcasing
advanced techniques before the basics are solid.

### III. Test-First for Non-Trivial Logic

Any exercise involving logic beyond a trivial syntax demo (structs with
behavior, error handling, concurrency, algorithms) MUST include at least one
`_test.go` file written before or alongside the implementation, following
Go's standard `testing` package. Pure syntax demonstrations (e.g., variable
declarations, basic loops) are exempt.

**Rationale**: Builds real testing habits early without overburdening
trivial exercises with test ceremony that adds no learning value.

### IV. Incremental Complexity

New exercises in this folder MUST build on concepts already introduced in
earlier numbered folders (01–13) and MUST follow the roadmap order declared
in the repository root `README.md` (Fundamentos → Estruturas de Dados →
Ponteiros/Métodos → Concorrência → Interfaces/Erros → DevOps Tooling). A
concept from a later roadmap stage MUST NOT be used as a silent dependency
of an earlier-stage exercise.

**Rationale**: Keeps the learning progression coherent and makes it possible
to look back at any folder and know exactly what was already known at that
point.

### V. Documented Learnings

Each exercise/feature MUST record, in its spec or a short top-of-file
comment, which concept is being practiced and one key takeaway once
completed. Full doc comments or verbose explanations are not required —
brevity is preferred.

**Rationale**: Converts one-off practice into retained knowledge and gives a
quick way to review progress across the roadmap later.

## Spec-Driven Workflow

This folder uses the speckit command sequence for each exercise/feature:
`/speckit-constitution` (this file, rarely re-run) → `/speckit-specify` →
`/speckit-clarify` (only if the spec has ambiguity worth resolving) →
`/speckit-plan` → `/speckit-tasks` → `/speckit-implement`. Given this is a
single-person learning repository, the workflow MAY be run lightly (e.g.,
skipping `/speckit-clarify` for small exercises), but the spec → plan →
tasks → implement order MUST NOT be skipped entirely for any feature tracked
under `specs/`.

## Scope & Constraints

- **Audience**: Single learner (repository owner); no external contributors
  or production users. Process overhead should stay proportional to that.
- **Language/Runtime**: Go 1.21+, matching the root `README.md` badge.
- **Out of scope for this folder**: Production deployment, CI/CD pipelines,
  multi-user concerns, and performance tuning — those belong to later
  DevOps-tooling stages of the roadmap, not to fundamentals practice.
- **Language of artifacts**: Spec/task prose may be written in Portuguese or
  English, matching the root README's bilingual style; code, identifiers,
  and comments should stay in English for consistency with idiomatic Go
  conventions.

## Governance

This constitution supersedes ad-hoc habits for any feature developed under
this folder's `specs/` directory. Amendments are made by re-running
`/speckit-constitution` when the learning focus or workflow materially
changes, and MUST update this file's Sync Impact Report and version per the
rules below:

- **MAJOR**: Removing or redefining a principle in a backward-incompatible
  way (e.g., dropping test-first expectations entirely).
- **MINOR**: Adding a new principle or materially expanding guidance.
- **PATCH**: Wording clarifications, typo fixes, non-semantic edits.

Because this is a personal learning repository, there is no formal PR
review gate — but before starting `/speckit-implement` on any feature, the
spec and plan SHOULD be reread against these principles as a self-check.

**Version**: 1.0.0 | **Ratified**: 2026-07-04 | **Last Amended**: 2026-07-04
