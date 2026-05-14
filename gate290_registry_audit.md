# Gate 290 Registry Audit — Bimodule Trace Capacity Sieve / Sector Hierarchy Audit

## Gate status

Gate 290 audits whether the Morita `1⊕3` trace multiplicity ledger can become a native branch selector for the two scalar-Morita amplitude branches left open by Gates 288–289.

The proposed intuition is:

```text
κ_C = 1
κ_Q = 3
```

so perhaps the quark sector should carry more finite spectral action than the lepton sector. Gate 290 separates this into two possible statements:

1. **Weak total capacity:** the total quark-sector trace should exceed the total lepton-sector trace.
2. **Strong per-slot monotonicity:** the quark trace per Morita slot should exceed the lepton trace per Morita slot.

The first statement is compatible with the derived multiplicities but too weak to select a branch. The second statement would select `r_+`, but it is not derived from multiplicity alone and would be an extra amplitude-ordering axiom.

## Inputs inherited

| Input | Value | Provenance |
|---|---:|---|
| `κ_C` | `1` | Gate 273 Morita trace multiplicity |
| `κ_Q` | `3` | Gate 273 Morita trace multiplicity |
| `r_+` | `(3591 + 136√123)/3099 ≈ 1.645470463011191` | Gate 275 / Gate 288 |
| `r_-` | `(3591 - 136√123)/3099 ≈ 0.672051318208557` | Gate 275 / Gate 288 |
| `X_+` | `≈ 0.9680658202595966` | Gate 288 contact-spectral cutoff |
| `X_-` | `≈ 1.905352660102002` | Gate 288 contact-spectral cutoff |

The sector moment proxies are:

```text
Tr(P_C D_F²) = X
Tr(P_Q D_F²) = 3Xr
Tr(P_C D_F⁴) = X²
Tr(P_Q D_F⁴) = 3X²r²
```

## Branch stress test

| Branch | `Tr(P_C D_F²)` | `Tr(P_Q D_F²)` | total D² bound | `Tr(P_C D_F⁴)` | `Tr(P_Q D_F⁴)` | total D⁴ bound |
|---|---:|---:|---|---:|---:|---|
| `r_+` | `0.968065820260` | `4.778771140464` | pass | `0.937151432355` | `7.612217870976` | pass |
| `r_-` | `1.905352660102` | `3.841484300621` | pass | `3.630368759358` | `4.919000543973` | pass |

The weak total-capacity bound

```text
Tr(P_Q D_F^{2n}) >= Tr(P_C D_F^{2n})
```

is therefore too weak: both branches pass.

## Per-slot diagnostic

Per Morita slot, the branches behave differently:

```text
per-slot C D² = X
per-slot Q D² = Xr
per-slot C D⁴ = X²
per-slot Q D⁴ = X²r²
```

Since `r_+ > 1` and `r_- < 1`, a strong per-slot monotonicity law

```text
Tr(P_Q D_F^{2n})/κ_Q >= Tr(P_C D_F^{2n})/κ_C
```

would select `r_+` and veto `r_-`.

But this is precisely the missing dynamical assumption. The derived Morita theorem counts how many quark slots exist relative to lepton slots. It does not prove that each quark slot must carry greater edge-map norm than the lepton slot. Promoting per-slot monotonicity into a theorem would convert multiplicity into amplitude, which the previous gates repeatedly firewalled.

## Geometric veto audit

Gate 290 therefore distinguishes:

| Proposed rule | Branch effect | Native status |
|---|---|---|
| Total quark trace ≥ total lepton trace | both branches pass | diagnostic only |
| Per-slot quark trace ≥ per-slot lepton trace | selects `r_+` | not derived; extra axiom/seal required |

So the lower branch is not mathematically falsified by current finite geometry.

## Status ledger

```text
CONDITIONAL_SUPPORT_GATE289_SECTOR_TRACE_DIAGNOSTIC_INHERITED
CONDITIONAL_SUPPORT_TRACE_CAPACITY_CANDIDATES_FORMALIZED
CONDITIONAL_SUPPORT_BRANCH_STRESS_TEST_COMPLETED
CONDITIONAL_SUPPORT_TOTAL_CAPACITY_BOUND_AUDITED
CONDITIONAL_SUPPORT_PER_SLOT_MONOTONIC_BOUND_DIAGNOSTIC_EXPOSED
CONDITIONAL_SUPPORT_TRACE_CAPACITY_FIREWALLS_PRESERVED
FAILED_ROUTE_NO_NATIVE_TRACE_CAPACITY_BOUND_DERIVED
FAILED_ROUTE_TOTAL_CAPACITY_BOUND_DOES_NOT_SELECT_BRANCH
FAILED_ROUTE_PER_SLOT_MONOTONIC_BOUND_IS_EXTRA_SELECTION_AXIOM
FAILED_ROUTE_GEOMETRY_CANNOT_STRICTLY_VETO_R_MINUS_DISTRIBUTION
FAILED_ROUTE_BRANCH_NOT_SELECTED_BY_TRACE_CAPACITY_BOUND
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

## Verdict

Gate 290 does **not** select the amplitude branch.

The Morita `1⊕3` multiplicity ledger provides a valid sector-capacity diagnostic, but not a native law that quark edge amplitudes must exceed lepton edge amplitudes per slot. The `r_+` branch is compatible with a heavy-quark/per-slot monotonic intuition, but that intuition remains unproved. Both `r_+` and `r_-` survive the finite-core audit.

A future gate must derive a genuine branch-sensitive sector functional, complete the physical `J`/hypercharge representation, or explicitly introduce a sealed amplitude-ordering principle before discarding `r_-`.
