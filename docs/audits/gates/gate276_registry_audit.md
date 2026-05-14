# Gate 276 Registry Audit — Scalar-Morita Spectral Shape Bridge / Branch Selector and Heat-Kernel Normalization

## Gate ID

`GATE276-SCALAR-MORITA-SPECTRAL-SHAPE-BRIDGE-BRANCH-SELECTOR-HEAT-KERNEL-NORMALIZATION-AUDIT`

## Purpose

Gate 276 audits whether the Gate 275 scalar-Morita two-branch amplitude-shape constraint can be promoted into a physical Seeley-de Witt `a₂/a₄` Higgs-ratio prediction.

It preserves the central firewall:

```text
scale-free raw finite spectral shape ≠ physical heat-kernel coefficient ratio
```

## Inputs inherited from Gate 275

Gate 275 connected two independent finite ledgers:

```text
Gate 169: λ_contact = 1197/4624
Gate 273: κ_C : κ_Q = 1 : 3
```

with the Morita trace-shape formulas:

```text
Tr(D_F²) = |x|² + 3|y|²
Tr(D_F⁴) = |x|⁴ + 3|y|⁴
```

The scalar-Morita bridge equation is:

```text
(|x|⁴ + 3|y|⁴) / (|x|² + 3|y|²)² = 1197/4624.
```

For

```text
r = |y/x|²
```

this gives:

```text
3099r² - 7182r + 3427 = 0
r = (3591 ± 136√123)/3099.
```

## Branch diagnostics

### Upper branch

```text
r_+ = (3591 + 136√123)/3099
r_+ ≈ 1.645470463011191
|y/x|_+ ≈ 1.282758926303454
Tr(D_F²)|x=1 ≈ 5.936411389033573
Tr(D_F⁴)|x=1 ≈ 9.122719133926790
Tr(D_F⁴)/Tr(D_F²)² ≈ 1197/4624
Tr(D_F⁴)/Tr(D_F²)|x=1 ≈ 1.536739712948354
```

### Lower branch

```text
r_- = (3591 - 136√123)/3099
r_- ≈ 0.672051318208557
|y/x|_- ≈ 0.819787361581378
Tr(D_F²)|x=1 ≈ 3.016153954625672
Tr(D_F⁴)|x=1 ≈ 2.354958922917579
Tr(D_F⁴)/Tr(D_F²)² ≈ 1197/4624
Tr(D_F⁴)/Tr(D_F²)|x=1 ≈ 0.780782068271395
```

Both branches reproduce the scale-free shape. Neither branch is selected by the finite data currently available.

## Branch selector audit

The following native selector candidates were audited:

| Selector | Result | Verdict |
|---|---:|---|
| Positivity / raw stability | both branches pass | cannot select |
| Charge/anomaly ledgers | insensitive to amplitude branch | cannot select |
| Physical `J` / parity orientation | physical `J` not derived | blocked |
| Energy/action minimization | no finite potential/action functional derived | blocked |

The branch ambiguity therefore remains exact.

```text
FAILED_ROUTE_TWO_BRANCH_VACUUM_AMBIGUITY_REMAINS
```

## Heat-kernel / Seeley-de Witt audit

Gate 276 records the formal expansion obligation:

```text
Tr(f(D/Λ)) ~ f₄Λ⁴ a₀ + f₂Λ² a₂ + f₀ a₄ + ...
```

But the project still lacks the following required maps:

```text
cutoff moments f₀,f₂,f₄
subtraction / renormalization scheme
scalar fluctuation map
Higgs mass / Higgs quartic projection
gauge kinetic projection
field normalization convention
physical anti-linear J
full chiral hypercharge representation
```

Therefore the gate refuses to identify raw finite traces with physical Seeley-de Witt coefficients.

```text
FAILED_ROUTE_HEAT_KERNEL_PROJECTION_NOT_DERIVED
FAILED_ROUTE_SCALAR_GAUGE_FIELD_NORMALIZATION_MISSING
FAILED_ROUTE_SEELEY_DE_WITT_A2_A4_NOT_DERIVED
FAILED_ROUTE_HIGGS_MASS_RATIO_NOT_CLAIMED
```

## Final status ledger

```text
CONDITIONAL_SUPPORT_GATE275_TWO_BRANCH_SHAPE_CONSTRAINT_INHERITED
CONDITIONAL_SUPPORT_SCALAR_MORITA_BRIDGE_FORMALIZED
CONDITIONAL_SUPPORT_BRANCH_SELECTOR_AUDIT_COMPLETED
CONDITIONAL_SUPPORT_BRANCH_CANDIDATE_MOMENTS_RECOMPUTED
CONDITIONAL_SUPPORT_FORMAL_HEAT_KERNEL_PROJECTION_REQUIREMENTS_DEFINED
CONDITIONAL_SUPPORT_SPECTRAL_FIREWALLS_PRESERVED
FAILED_ROUTE_TWO_BRANCH_VACUUM_AMBIGUITY_REMAINS
FAILED_ROUTE_ABSOLUTE_DF_SCALE_NOT_DERIVED
FAILED_ROUTE_HEAT_KERNEL_PROJECTION_NOT_DERIVED
FAILED_ROUTE_SCALAR_GAUGE_FIELD_NORMALIZATION_MISSING
FAILED_ROUTE_SEELEY_DE_WITT_A2_A4_NOT_DERIVED
FAILED_ROUTE_HIGGS_MASS_RATIO_NOT_CLAIMED
```

## Interpretation

Gate 276 confirms the scalar-Morita bridge as a real finite geometric shape constraint. It does not yet produce a Higgs mass prediction.

The exact boundary is now:

```text
Gate 169 scalar shape + Gate 273 Morita multiplicity
  -> two exact r branches
  -> raw candidate finite spectral shapes
  -> branch selector + heat-kernel projection still required
  -> no a₂/a₄ theorem yet
```

## Future theorem obligations

A future gate must derive at least the following before promoting the result to a Higgs-ratio theorem:

1. A finite-core selector for `r_+` or `r_-`.
2. A physical anti-linear `J` on the completed finite Hilbert space.
3. Full chiral hypercharge representation for `C ⊕ H ⊕ M3(C)`.
4. Heat-kernel cutoff moment normalization.
5. Scalar/gauge projection from raw finite moments to Lagrangian coefficients.
6. Field normalization and subtraction scheme.
7. A statement of what dimensionless quantity is actually predicted.

## Tests

Focused tests passed:

```bash
go test -p=1 ./pkg/bridge/scalarmoritaspectralbridge -count=1 -timeout=120s -v

go test -p=1 ./pkg/bridge/scalarmoritaspectralbridge ./pkg/bridge/physicalfinitehilbertcompletion ./pkg/bridge/nativeweakquaternionicalgebra -count=1 -timeout=120s -v

go list ./internal/app

go list ./cmd/asha
```

No full internal tests, full package tests, or `go test ./...` were run.
