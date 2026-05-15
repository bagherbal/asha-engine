# Gate 485 Registry Audit — Koide Constraint Provenance & Topological Baseline

## Verdict

```text
CONDITIONAL_SUPPORT_C3_SHADOW_BASIS_ORTHOGONALITY_PROVED
CONDITIONAL_SUPPORT_NULL_BOUNDARY_FORCES_R_OVER_S_SQRT2
CONDITIONAL_SUPPORT_KOIDE_Q_TWO_THIRDS_DERIVED_FROM_NULL_C3_SHADOW
CONDITIONAL_SUPPORT_CHARGED_LEPTON_BRIDGE_SHADOW_COMPATIBLE_WITH_NULL_BASELINE
CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_KOIDE_PROVENANCE
```

Gate 485 closes the provenance gap left by Gate 484: `R/S = sqrt(2)` is not inserted from the observed charged-lepton masses. It follows from the C3 square-root shadow **only after** the shadow is placed on the `Cℓ(1,7)` null boundary with the democratic leg timelike and the phase-plane leg spacelike.

## Inherited boundary

| inherited object | status |
|---|---|
| Gate 480 null cone | `true` |
| Gate 481 common baseline cancellation | `true` |
| Gate 483 quark/lepton topology separation only | `true` |
| Gate 484 C3 basis validated | `true` |
| Gate 484 charged-lepton Koide shadow found | `true` |
| observed masses remain bridge data | `true` |

## C3 square-root shadow theorem

Define the normalized C3 mass-shadow coordinates

```text
x_i = sqrt(m_i) = S + R cos(θ_i - ψ)
θ_i ∈ {0, 2π/3, 4π/3}
```

The democratic component is `D_i=S`; the phase-plane component is `P_i=R cos(θ_i-ψ)`. The C3 identities are:

```text
Σ_i cos(θ_i-ψ)    = 0
Σ_i cos²(θ_i-ψ)   = 3/2
D·P                 = S R Σ_i cos(θ_i-ψ) = 0
||D||²              = 3S²
||P||²              = (3/2)R²
```

These are C3 identities, not mass fits. Gate 485 verifies them across phase samples:

| ψ | Σcos | Σcos² | D·P | ||D||² at S=1 | ||P||² at R=1 | pass |
|---:|---:|---:|---:|---:|---:|---|
| 0 | 2.22044604925e-16 | 1.5 | 2.22044604925e-16 | 3 | 1.5 | `true` |
| 0.349065850399 | 1.11022302463e-16 | 1.5 | 1.11022302463e-16 | 3 | 1.5 | `true` |
| 1.0471975512 | -1.11022302463e-16 | 1.5 | -1.11022302463e-16 | 3 | 1.5 | `true` |
| -0.628318530718 | 2.35922392733e-16 | 1.5 | 2.35922392733e-16 | 3 | 1.5 | `true` |

## Null boundary derivation

The native boundary audit uses the `Cℓ(1,7)` lightlike form with democratic/hierarchy direction timelike and the C3 phase plane spacelike:

```text
q_C3(S,R) = ||D||² - ||P||²
          = 3S² - (3/2)R²
q_C3 = 0  ⇒  3S² = (3/2)R²
          ⇒  R² = 2S²
S>0,R>0   ⇒  R/S = sqrt(2)
```

Numerical exactness check on a representative positive null branch:

| S | R | ψ | 3S² | (3/2)R² | q | R/S | residual from sqrt(2) |
|---:|---:|---:|---:|---:|---:|---:|---:|
| 1 | 1.41421356237 | 0.448798950513 | 3 | 3 | -4.4408920985e-16 | 1.41421356237 | 0 |

## Koide equivalence

For the same C3 shadow,

```text
Q = (Σ_i x_i²)/(Σ_i x_i)²
  = (||D||² + ||P||²)/(3S)²
  = (3S² + (3/2)R²)/(9S²).
```

Substitute the null result `R²=2S²`:

```text
Q = (3S² + 3S²)/(9S²) = 2/3.
```

Gate 485 therefore derives `R/S = 1.41421356237` and `Q = 2/3 = 0.666666666667` with no observed masses in the proof.

## Boundary-collapse ledger

| space | before boundary | after null boundary | collapsed | still free |
|---|---:|---:|---:|---|
| C3 square-root shadow `(S,R,ψ)` | `3` | `2` | `1` | `S scale, ψ phase` |

The boundary collapses the radial shape coordinate `R/S`, not the mass spectrum. A null C3 baseline has the form

```text
x_i = S [1 + sqrt(2) cos(θ_i - ψ)].
```

So the Koide scalar is fixed, while the absolute scale `S`, the C3 sheet/phase `ψ`, and physical sector perturbations remain unselected.

## Sector firewall

```text
charged-lepton Koide shadow = bridge-compatible with the colorless null-C3 baseline
quark Koide promotion       = rejected; quarks are color/QCD-dressed, not bare null baselines
absolute masses             = not derived
ψ phase / C3 sheet          = not selected
CKM/PMNS                    = not constructed
native flavor dimension     = 13 remains sealed
K/X/Y charged coefficients  = 9 remain sealed
```

Rejected promotions:

```text
FAILED_ROUTE_EMPIRICAL_KOIDE_SHADOW_AS_NATIVE_MASS_FIT_REJECTED
FAILED_ROUTE_NULL_KOIDE_DOES_NOT_DERIVE_ABSOLUTE_MASSES
FAILED_ROUTE_NULL_KOIDE_DOES_NOT_SELECT_C3_PHASE_PSI
FAILED_ROUTE_QUARK_SECTORS_COLOR_DRESSED_NOT_NULL_BASELINE
FAILED_ROUTE_KOIDE_BASELINE_AS_CKM_PMNS_PREDICTION_REJECTED
FAILED_ROUTE_KOIDE_BASELINE_DOES_NOT_COLLAPSE_13_MODULI
```

## Truth statement

Gate485 result: the C3 shadow null condition 3S²-(3/2)R²=0 natively forces R/S=1.41421356237 and Q=0.666666666667 for a bare colorless null baseline, collapsing one C3 shape coordinate while leaving scale, phase, quark dressing, CKM, PMNS, and physical masses sealed.

## Next step

Gate 486 — Color-dressing deformation firewall or lepton-baseline airlock. audit whether color/winding topology can define a bridge-only dressing operator for quark deviations from the null-C3 baseline without importing masses or mixing matrices as native data
