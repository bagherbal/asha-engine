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


## v0.3 — Octonionic / G₂ Calibration Gate

Adds the rank-14 G₂ matter-calibration support inside the middle chamber `Λ⁴R⁸`:

- standard Fano convention for the associative 3-form `φ`;
- coassociative Hodge-dual 4-form `*φ`;
- construction of `M₁₄ᴳ = 7_t ⊕ 7_s`;
- orthonormal calibration frame `Q_G`;
- support projector `P_G = Q_G Q_Gᵀ`;
- exact theorem checks for rank, trace, idempotence, and symmetry.

The next gate is `K₇ = Im(P_B) ∩ Im(P_G)`.


## v0.4 — Contact gate

The engine now constructs the Boolean--Octonionic contact space:

```text
K = Im(P_B) ∩ Im(P_G)
```

It verifies the contact projector `P_K`, containment inside both `P_B` and `P_G`, the finite contact index `I_BG`, and reports the bare contact leakage invariant `L_BG = ||P_B P_G - P_K||_F` without identifying it with the cosmological constant.

## v0.5 — B-sector dynamical vacuum gate

This version adds the first variational theorem gate:

```text
O_B = Wᵀ(I − P_G)W
S_B[b] = ||(I − P_G)Wb||² = bᵀO_Bb
```

The engine verifies that the kernel of `O_B` has dimension 7 and equals the Boolean-coordinate image of the contact space `K = Im(P_B) ∩ Im(P_G)`. This upgrades the contact space from a static intersection to a finite zero-energy sector. No physical constants are inferred from this gate.
