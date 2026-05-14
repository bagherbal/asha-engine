# GateResearcherMethod

- Start from the latest registry/audit files and identify the exact gate boundary: proven facts, sealed no-gos, conditional bridges, and quarantined empirical inputs.
- Read only the minimum source chain needed for the next gate first: registry entry, theorem package, tests, app wiring, and the immediately preceding gate audit.
- State the mathematical object before coding: inputs, invariants, unknowns, equations, admissible outputs, and firewall conditions.
- Implement the gate as a small isolated package when possible; keep phenomenology, finite algebra, and interpretation in separate structs/functions.
- Prefer exact/rational or auditable floating calculations; log residuals, tolerances, degeneracies, and underdetermined dimensions explicitly.
- Treat failure as data: if the system is underdetermined, overconstrained, irrational, noncanonical, or representation-free, record a no-go instead of forcing a fit.
- Reuse previous-gate audit snapshots for lightweight checks when importing a deep historical package would only retrieve already-audited constants.
- Avoid temporary `go run` probes that import broad theorem chains; use source/audit inspection, small package-local tests, or standalone arithmetic instead.
- Add focused tests for the new theorem package first; use compile-only checks for app/CLI wiring before broader runs.
- Avoid slow first-pass commands such as full `go test ./...` or full CLI execution until the focused gate path is stable.
- Use faster validation loops: `go test -p=1 ./pkg/path -count=1`, then selected dependent packages, then `go test -p=1 ./internal/app -run '^$'` and `go test -p=1 ./cmd/asha -run '^$'`.
- If a command times out, narrow scope rather than retrying blindly: identify package boundaries, warm the build cache, reduce verbose output, or replace deep imports with audited snapshots when scientifically legitimate.
- Before packaging, clean generated binaries and transient artifacts with a root-level artifact check, especially `*.test`, logs, temporary folders, and accidental build outputs.
- End every gate with a registry audit that separates theorem status into permanent, sealed/no-go, conditional, empirical/quarantined, and next-gate obligations.
