# Gate 281 Registry Audit — Resolvent Branch Semantics / Projector-to-Sector Orientation Seal

## Gate statement

Gate 281 audits whether the conditional contact projectors derived after the `ResolventAdjunctionSeal` possess native physical semantics. Specifically, it asks whether the Gate-273 Morita trace multiplicities

```text
κ_C : κ_Q = 1 : 3
```

can prefer one of the six possible sealed projector-sector orientations:

```text
3 resolvent branches × 2 projector orientations = 6 candidate states.
```

## Inputs inherited from prior gates

| Source | Inherited fact | Status |
|---|---|---|
| Gate 273 | Morita finite-Hilbert trace multiplicity `κ_C:κ_Q = 1:3` | finite-derived multiplicity ledger |
| Gate 275 | scalar-Morita branches `r = (3591 ± 136√123)/3099` | finite shape constraint, branch ambiguous |
| Gate 277 | sector-level pairing `{u,d}|{e,ν}` selected by `τ_eta + B_gap` | conditional support |
| Gate 279 | no nontrivial rational contact projector over `Q` | strict Galois no-go |
| Gate 280 | resolvent adjunction constructs three conditional `2+2` projector branches | sealed conditional support |

## Trace and norm semantic audit

For each of the three conditional resolvent branches, Gate 281 audits the two projectors `P_A` and `P_B`.

The invariant result is the same for all branches:

```text
Tr(P_A) = 2
Tr(P_B) = 2
rank(P_A) ≈ 2
rank(P_B) ≈ 2
P_A + P_B = I
P_A P_B = 0
```

Therefore the contact projector split is:

```text
2 | 2
```

while the Morita multiplicity split is:

```text
1 | 3
```

These live on different carriers:

```text
contact companion module: quartic root orbit, split by resolvent projectors
finite Hilbert bimodule: lepton/quark trace multiplicity, split by Morita sectors
```

The Morita `1|3` multiplicity cannot natively orient a `2|2` contact projector pair.

Naive Frobenius norms of projector matrices are recorded only as basis-dependent diagnostics. They are not promoted to finite-core semantics because they depend on the companion power basis and are not invariant contact data.

## ProjectorSectorOrientationSeal

Because no native orientation is found, Gate 281 activates a conditional seal:

```text
ProjectorSectorOrientationSeal
```

Representative sealed witness:

```text
branch: z_high_pairing_q1q2_q3q4
z ≈ 0.793092963834819
P_(q1,q2) -> {u,d}
P_(q3,q4) -> {e,ν}
```

This is a quarantined orientation witness only. It does not rewrite the native theorem status.

## Amplitude branch audit

Gate 281 then asks whether the sealed projector orientation determines the Gate-275 amplitude branch:

```text
r_+ = (3591 + 136√123)/3099 ≈ 1.645470463011191
r_- = (3591 - 136√123)/3099 ≈ 0.672051318208557
```

Result:

```text
no derived algebraic map: resolvent branch -> r_+ or r_-
```

A contact resolvent orientation selects a `2+2` root/projector split. The Gate-275 branches are roots of the scalar-Morita amplitude-shape equation. No native functor currently maps the former to the latter.

## Status ledger

```text
CONDITIONAL_SUPPORT_PROJECTOR_TRACE_NORM_SEMANTIC_AUDIT_COMPLETED
CONDITIONAL_SUPPORT_PROJECTOR_SECTOR_ORIENTATION_SEAL_ACTIVATED
CONDITIONAL_SUPPORT_REPRESENTATIVE_PROJECTOR_SECTOR_ORIENTATION_ASSIGNED
CONDITIONAL_SUPPORT_SEELEY_DE_WITT_PREPARATION_OBLIGATIONS_DOCUMENTED
CONDITIONAL_SUPPORT_PROJECTOR_ORIENTATION_FIREWALLS_PRESERVED
FAILED_ROUTE_NO_NATIVE_PROJECTOR_ORIENTATION_SELECTOR_DERIVED
FAILED_ROUTE_1_PLUS_3_MULTIPLICITY_DOES_NOT_PREFER_2_PLUS_2_PROJECTOR_ORIENTATION
FAILED_ROUTE_PROJECTOR_ORIENTATION_DOES_NOT_DERIVE_RESOLVENT_TO_R_BRANCH_MAP
FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED
FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED
```

## Direct answer to the Gate-281 question

The `1⊕3` Morita trace multiplicities do **not** mathematically prefer a projector orientation natively. They count lepton/quark multiplicities in the finite Hilbert bimodule, whereas resolvent projectors split the contact companion module as `2⊕2`.

A sealed projector-sector orientation is therefore needed for downstream stress tests. However, even after sealing a representative 1-in-6 orientation, the amplitude branch `r_+` versus `r_-` remains unselected unless a future theorem derives a resolvent-to-scalar-Morita branch map.

## Remaining obligations before a Higgs-ratio claim

1. Derive a native map from selected contact resolvent branch to Gate-275 scalar-Morita branch.
2. Complete physical charge conjugation `J` and opposite action.
3. Complete chiral/hypercharge representation on the finite Hilbert space.
4. Derive the heat-kernel / Seeley-de Witt projection and subtraction scheme.
5. Separate scalar and gauge kinetic normalizations.

Until these are satisfied, no `a₂/a₄` Higgs mass ratio is claimed.
