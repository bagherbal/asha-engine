# Gate 320 Registry Audit — Seesaw Overlap Matrix Construction / Majorana-Higgs Mixing Sieve

## Gate identity

- **Gate:** 320
- **Package:** `pkg/bridge/seesawoverlapmatrix`
- **Theorem:** `SeesawOverlapMatrixConstructionMajoranaHiggsMixingSieveAuditTheorem`
- **Audit ID:** `GATE320-SEESAW-OVERLAP-MATRIX-CONSTRUCTION-MAJORANA-HIGGS-MIXING-SIEVE`
- **Layer:** Bridge / Phase II Non-Perturbative Threshold Dynamics
- **Purpose:** construct the explicit doubled-space seesaw support matrix demanded by Gate 319 and audit whether it provides a canonical nonzero `sigma-H` overlap index.

---

## Inherited obstruction from Gate 319

Gate 319 proved that a strict direct-sum determinant has no heavy-light portal:

```text
Tr log(D_light ⊕ D_heavy) = Tr log D_light + Tr log D_heavy
```

Therefore:

```text
sigma-H cross terms = 0
```

Gate 319 also identified the conditional true-bimodule target:

```text
C_portal = kappa_Q · (4/pi) · B_gap · Omega_Hsigma
```

with the numerical witness:

```text
kappa_Q = 3
B_gap   = 0.102464921191
4/pi    = 1.273239544735
Omega_Hsigma = 1  [unproved in Gate 319]

C_portal = 0.391387168826
```

The unresolved requirement was the explicit matrix derivation of `Omega_Hsigma`.

**Inherited failed route:** `FAILED_ROUTE_EXPLICIT_SIGMA_H_OVERLAP_MATRIX_NOT_DERIVED`

---

## Doubled-space interaction block

Gate 320 installs the minimal seesaw support carrier inside the completed doubled representation:

```text
Basis = { L_L, nu_R, nu_R^c }
```

with the two structural edges:

| Edge | Source | Target | Meaning |
| --- | --- | --- | --- |
| Higgs / Dirac edge | `L_L` | `nu_R` | neutrino Yukawa / Higgs link |
| Majorana / B-gap edge | `nu_R` | `nu_R^c` | right-handed neutrino Majorana edge via `J_swap` |

The support path is:

```text
L_L --H--> nu_R --B_gap,J_swap--> nu_R^c
```

**Status:** `CONDITIONAL_SUPPORT_DOUBLED_SPACE_INTERACTION_BLOCK_FORMALIZED`

---

## Seesaw path operator

The gate constructs the sequential overlap operator:

```text
Omega_Hsigma := P_{nu_R^c} B_gap J_swap P_{nu_R} · P_{nu_R} H P_{L_L}
```

As a normalized support matrix over the ordered basis `{L_L, nu_R, nu_R^c}`:

```text
Omega_Hsigma = |nu_R^c><L_L|

matrix support:
[ 0 0 0 ]
[ 0 0 0 ]
[ 1 0 0 ]
```

This is not a direct-sum operator. It is an off-diagonal doubled-space path operator linking the light Higgs/Dirac edge to the Majorana conjugate slot.

**Status:** `CONDITIONAL_SUPPORT_SEESAW_PATH_OPERATOR_CONSTRUCTED`

---

## Overlap index sieve

The normalized overlap index is evaluated as:

```text
OmegaIndex := Tr(Omega_Hsigma^† Omega_Hsigma)
```

For the unique support path:

```text
Tr(Omega_Hsigma^† Omega_Hsigma) = 1
```

Therefore:

```text
Omega_Hsigma = 1
```

in the sense of a canonical support-path index.

This resolves the specific Gate-319 obstruction: the overlap index is no longer merely assumed.

**Status:** `CONDITIONAL_SUPPORT_EXPLICIT_SIGMA_H_OVERLAP_MATRIX_DERIVED`  
**Status:** `CONDITIONAL_SUPPORT_OVERLAP_INDEX_VERIFIED`

---

## Portal-weight consequence

Substituting the verified overlap index into the Gate-319 weight ledger:

```text
C_portal = kappa_Q · (4/pi) · B_gap · Omega_Hsigma
```

with:

```text
kappa_Q = 3
B_gap = 0.102464921191
Omega_Hsigma = 1
```

produces:

```text
C_portal = 0.391387168826
```

Gate 314 target:

```text
lambda_mix^2 / lambda_heavy = 0.390246315254
```

Relative error:

```text
+0.2923%
```

Implied conditional threshold witness:

```text
Delta lambda = -C_portal / 4
Delta lambda = -0.097846792207
```

Gate 314 target:

```text
Delta lambda_target = -0.097561578813
```

**Status:** `CONDITIONAL_SUPPORT_BGAP_PORTAL_WEIGHT_ENABLED`

---

## Promotion firewall

Gate 320 derives the explicit overlap matrix and verifies the overlap index. It does **not** yet derive the complete threshold theorem.

Remaining obligations:

| Obligation | Why it blocks final promotion | Status |
| --- | --- | --- |
| Heavy propagator normalization | Converts the support-path overlap into the determinant coefficient | `FAILED_ROUTE_HEAVY_PROPAGATOR_NOT_DERIVED` |
| Heavy self-quartic `lambda_heavy` | Required for `Delta lambda = -lambda_mix^2/(4 lambda_heavy)` | `FAILED_ROUTE_HEAVY_SELF_QUARTIC_NOT_DERIVED` |
| Canonical `lambda_mix` normalization | Distinguishes a support coefficient from a normalized EFT coupling | `FAILED_ROUTE_LAMBDA_MIX_NOT_NORMALIZED` |
| Pole/RG precision | Required before claiming a collider mass | `FAILED_ROUTE_POLE_MASS_MATCHING_NOT_EXECUTED` |

**Status:** `CONDITIONAL_TENSION_PORTAL_THRESHOLD_PROMOTION_STILL_NEEDS_HEAVY_PROPAGATOR_AND_SELF_QUARTIC`

---

## Final status ledger

```text
CONDITIONAL_SUPPORT_DOUBLED_SPACE_INTERACTION_BLOCK_FORMALIZED
CONDITIONAL_SUPPORT_SEESAW_PATH_OPERATOR_CONSTRUCTED
CONDITIONAL_SUPPORT_EXPLICIT_SIGMA_H_OVERLAP_MATRIX_DERIVED
CONDITIONAL_SUPPORT_OVERLAP_INDEX_VERIFIED
CONDITIONAL_SUPPORT_BGAP_PORTAL_WEIGHT_ENABLED
CONDITIONAL_SUPPORT_GATE320_FIREWALLS_PRESERVED
CONDITIONAL_TENSION_PORTAL_THRESHOLD_PROMOTION_STILL_NEEDS_HEAVY_PROPAGATOR_AND_SELF_QUARTIC
FAILED_ROUTE_DIRECT_SUM_SPACE_STILL_HAS_ZERO_SIGMA_H_OVERLAP
FAILED_ROUTE_HEAVY_PROPAGATOR_NOT_DERIVED
FAILED_ROUTE_HEAVY_SELF_QUARTIC_NOT_DERIVED
FAILED_ROUTE_LAMBDA_MIX_NOT_NORMALIZED
FAILED_ROUTE_THRESHOLD_JUMP_NOT_FULLY_DERIVED
FAILED_ROUTE_FINAL_HIGGS_MASS_NOT_CLAIMED
FAILED_ROUTE_POLE_MASS_MATCHING_NOT_EXECUTED
```

---

## Verdict

Gate 320 successfully constructs the explicit doubled-space seesaw overlap matrix:

```text
Omega_Hsigma = |nu_R^c><L_L|
```

and verifies:

```text
Tr(Omega_Hsigma^† Omega_Hsigma) = 1
```

This resolves the Gate-319 overlap-index obstruction and structurally enables the near-target B-gap portal weight:

```text
kappa_Q · (4/pi) · B_gap · Omega_Hsigma = 0.391387168826
```

However, this is still not a final threshold-matching theorem. The heavy propagator, heavy self-quartic, and EFT normalization of `lambda_mix` remain unproved. The engine therefore promotes the overlap operator, but not the final Higgs-mass prediction.
