# Gate 639 — CompactSplitTwistResidual Invariant Audit

## Purpose

Gate 638 proved that the compact octonionic pullback `Omega_0`, the inherited compact metric `g_K`, and the Hodge split polarization `B_K=g_K S_K` are lawfully related but unfused:

```text
g_Omega ≈ c g_K,
B_K ≈ c^{-1} g_Omega S_K,
```

while the admissible `S_K` twists fail to produce a `B_K`-compatible split-G2 3-form.  Gate 639 audits the repeated residual itself:

```text
rho_twist ≈ 0.470317081001772.
```

This is an internal finite-geometry obstruction audit only.  It does not derive boundary stress, scalar RG matching, physical spacetime, flavor, Higgs mass, CKM/PMNS, gauge unification, split-G2, or a native `7/72` theorem.

## Inherited Gate638 data

```text
g_Omega ≈ c g_K,
c ≈ 8.63167457503e-05,
relative residual ≈ 8.37e-15.

B_K ≈ c^{-1} g_Omega S_K,
scaled residual ≈ 8.36e-15.
```

Gate638 already certified that the compact calibration and Hodge split bilinear are native but not fused into a compatible split 3-form.

## Repeated residual cluster

The residual appears in three independent Gate638 routes:

| Route | Inertia | Relative residual to `B_K` | Cluster role |
|---|---:|---:|---|
| `omega_1_alt = Alt[Omega_0(S_Kx,y,z)]` | `(4,3,0)` | `0.470317081001771` | admissible split twist |
| `omega_2_alt = Alt[Omega_0(S_Kx,S_Ky,z)]` | `(3,4,0)` | `0.470317081001770` | admissible split twist |
| `omega_B_alt = Alt[B_K(x ×_{Omega_0} y,z)]` | `(4,3,0)` | `0.470317081001773` | antisymmetrized cross-product route |

The cluster mean is:

```text
rho_twist = 0.470317081001772,
spread ≈ 2.44e-15.
```

The compact untwisted routes remain far away from `B_K`:

```text
omega_0 residual ≈ 0.989743318610787,
omega_3 residual ≈ 0.989743318610787.
```

## Invariance audit

Gate 639 treats `rho_twist` as a projective metric residual:

```text
rho = min_c ||g_candidate - c B_K||_F / ||B_K||_F.
```

The audit records that this projective residual is not removed by the typed normalizations that should be gauge-like for this comparison:

| Probe | Result |
|---|---|
| Orthogonal basis change | invariant |
| `Omega` / candidate metric rescaling | invariant |
| Target sign / orientation flip `B_K -> -B_K` | invariant |
| `S_K` orientation flip | invariant |
| determinant-volume normalization | invariant |
| trace-free projective comparison | invariant |

The maximum recorded drift across these projective probes is zero in the normalized audit ledger.

## Source sweep audit

Switching among the Gate637 compact `P_G` pullback sources does not remove the obstruction:

```text
best compact source = omega_t,
best compact source residual ≈ 0.989743318610787,
best split twist residual ≈ 0.470317081001772.
```

Therefore the repeated `rho_twist` is not explained away by choosing a different compact `P_G` pullback source.

## Classification

Gate 639 conditionally classifies `rho_twist` as an internal compact/split obstruction witness:

```text
rho_twist ≈ 0.470317081001772,
rho_twist^2 ≈ 0.221198156552,
arcsin(rho_twist) ≈ 0.489650043230328 rad ≈ 28.0548809155 degrees.
```

The angle diagnostic is only a magnitude diagnostic.  It is not promoted to a native angle theorem or physical observable.

## Native ASHA status

Certified:

```text
compact octonionic calibration aligned with g_K,
Hodge split polarization B_K=g_K S_K,
repeated projective twist residual rho_twist,
projective invariance under typed normalization probes.
```

Not certified:

```text
B_K-compatible native split-G2 3-form,
physical spacetime metric,
boundary-stress assignment,
scalar/flavor transport theorem,
gauge unification,
native 7/72 trace theorem.
```

## Verdict

```text
PASS_GATE638_TWO_NATIVE_STRUCTURES_REMAIN_UNFUSED
PASS_TWIST_RESIDUAL_REPEATED_ACROSS_ROUTES
PASS_RESIDUAL_INVARIANCE_TESTS_COMPUTED
PASS_RHO_TWIST_NOT_REMOVED_BY_SCALE_OR_ORIENTATION_NORMALIZATION
PASS_PROJECTIVE_METRIC_RESIDUAL_AUDITED
CONDITIONAL_SUPPORT_RHO_TWIST_IS_COMPACT_SPLIT_OBSTRUCTION_INVARIANT
FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FIREWALL_PRESERVED_GATE639_RHO_TWIST_IS_INTERNAL_OBSTRUCTION_ONLY
```
