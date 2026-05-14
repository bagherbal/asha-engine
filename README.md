# ASHA Engine

A typed Go research engine for finite algebra, Clifford phase-space geometry, Boolean–Octonionic contact structure, theorem gates, and bridge-layer physics.

The first scaffold intentionally starts with the **exact algebraic foundation** only:

1. exterior-grade structure of `ΛR⁸`;
2. Clifford algebra dimension/signature verification for `Cℓ(1,7)`;
3. covariant phase-space bookkeeping;
4. theorem-gated reporting.

Physical constants are **not** hard-coded in the engine. Observed values belong in external validation datasets, not in the finite core.

## Run

```bash
go test ./...
go run ./cmd/asha
```

## Design law

The engine separates:

- `EXACT_FINITE`: directly computed algebraic facts;
- `VARIATIONAL`: finite action/eigenproblem results;
- `OPEN_TEST`: sharply stated frontier tests;
- `BRIDGE_REQUIRED`: physical interpretation needing a continuum/normalization bridge;
- `INVALID_COMPARISON`: quantities that must not be compared yet.

## First theorem gates

```text
Gate 0 — Exterior grade structure
Gate 1 — Clifford phase-space algebra
Gate 2 — Covariant phase-space bookkeeping
```

Next implementation target:

```text
Gate 3 — Boolean incidence support P_B
Gate 4 — Octonionic G₂ calibration P_G
Gate 5 — Contact vacuum K₇
```
