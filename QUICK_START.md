# QUICK START

This project is a theorem-gated Go research engine. The normal workflow is to inspect the theorem package for a gate, run only targeted tests, and read the generated audit under `docs/audits/gates/`.

## Requirements

- Go 1.22+ or the version declared in `go.mod`.
- A shell environment capable of running `go test`.

## Basic commands

```bash
# Check that the command entrypoint builds/go-lists
go list ./cmd/asha

# Check registry wiring without running expensive app tests
go list ./internal/app

# Run a specific gate package only
go test -p=1 ./pkg/bridge/reviewerobjectionmatrix -count=1

# Run selected related packages only, not the full suite
go test -p=1   ./pkg/bridge/reviewerobjectionmatrix   ./pkg/bridge/executiveabstractclaimaudit   ./pkg/bridge/manuscriptskeletonexport   -count=1
```

## What not to run by default

```bash
# Avoid unless you intentionally want a full project validation.
go test ./...

# Avoid internal/app tests when timeout risk matters.
go test ./internal/app
```

## Where to find things

```text
README.md                         # high-level project status and boundaries
docs/architecture.md              # detailed architecture ledger
docs/INDEX.md                     # documentation/artifact structure
docs/ARTIFACT_INDEX.md            # canonical artifact index
docs/REPRODUCIBILITY_CHECKLIST.md # targeted validation policy
docs/audits/gates/INDEX.md        # gate audit index
docs/summaries/                   # summary docs and logical tower maps
docs/paper/                       # manuscript workspace
docs/visuals/                     # figures and diagrams
```

## Adding a new gate

1. Add the package under the relevant `pkg/...` namespace.
2. Add `analysis.go`, `theorem.go`, `format.go`, and `analysis_test.go` following the local gate style.
3. Wire the gate into `internal/app/app.go` if it is part of the registry.
4. Generate the audit into `docs/audits/gates/gateNNN_registry_audit.md`.
5. Run only targeted tests for the new package and selected dependencies.
6. Update `README.md` and `docs/architecture.md` with a short, bounded addendum only.

## Current artifact policy

- Root should stay clean.
- Gate audits live in `docs/audits/gates/`.
- The canonical artifact index lives in `docs/ARTIFACT_INDEX.md`.
- The reproducibility checklist lives in `docs/REPRODUCIBILITY_CHECKLIST.md`.
- Phenomenology audits live in `docs/audits/phenomenology/`.
- Summary/tower documents live in `docs/summaries/`.
- Paper drafts and final paper files live in `docs/paper/`.
- Visuals live in `docs/visuals/`.
