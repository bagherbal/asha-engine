# Gate 662 — BoundaryWeightedDeficitClosure Scale-Sweep and Sensitivity Audit

## Purpose

Gate 661 isolated the noncircular bridge diagnostic:

```text
E_72 = kappa_lambda + kappa_e - W_72,
W_72 = (65/72)|lambda(Lambda_12)| + (7/72)(R_3-1).
```

Gate 662 audits whether this closure is selected by `Lambda_12` in the current v1 transport ledger, or whether it is merely evaluated there.

This is a bridge-layer scale-sweep and sensitivity audit only. It does not derive Higgs mass, scalar stability, flavor, CKM/PMNS, gauge unification, boundary stress, or a native `7/72` theorem.

## Implementation

- Package: `pkg/bridge/generation2boundaryweighteddeficitclosurescalesweepaudit`
- Theorem: `Generation2BoundaryWeightedDeficitClosureScaleSweepAndSensitivityAuditTheorem()`
- Runtime marker: `gate662-boundary-weighted-deficit-closure-scale-sweep-sensitivity-audit-20260518`

## Inherited exact-ledger closure

```text
kappa_lambda = 0.0443230430960771
kappa_e      = 0.00550355419157456
K_sum        = 0.0498265972876517

W_72 = 0.0498265964350682
E_72 = K_sum - W_72 = 8.52583441346e-10.
```

Gate 662 keeps the scalar runtime formula lift demoted as independent evidence because `kappa_lambda` was defined from the low-scale scalar runtime gap.

## Scale sweep

Using the project’s v1 one-loop scalar/gauge transport equations as a diagnostic ledger, Gate 662 evaluates the closure at:

```text
Lambda_12,
Lambda_13,
Lambda_23,
Lambda_geom.
```

Two gauge-residual conventions are recorded:

1. `EW-mean strong residual`: `g3 / ((g1+g2)/2) - 1`, which reproduces `R_3-1` at `Lambda_12`.
2. Pair-residual diagnostics at each pairwise meeting scale, where the nonmeeting gauge coupling is compared to the meeting-pair mean.

In both conventions, the current v1 grid selects `Lambda_12` as the unique smallest residual among the tested scales.

## Local perturbation sweep

Gate 662 also sweeps:

```text
ln(mu/Lambda_12) ∈ {-2,-1,-0.5,-0.1,0,+0.1,+0.5,+1,+2}.
```

The local grid minimum occurs at the exact `Lambda_12` ledger point. Nearby shifts at `±0.1` in log-scale already enlarge `|E_72|` to roughly the `1e-4` scale, so the closure is scale-sensitive in the current v1 transport ledger.

## Weight sensitivity

The empirical best boundary interpolation weight is:

```text
w_best = [K_sum-|lambda(Lambda_12)|] / [(R_3-1)-|lambda(Lambda_12)|]
       = 0.0972228818894104.
```

The typed candidate is:

```text
7/72 = 0.0972222222222222.
```

Difference:

```text
w_best - 7/72 = 6.59667188138e-7.
```

Using the OrientationBalance approximation

```text
kappa_e_orient = sin²(theta13)/4 - J_CKM
```

shifts the best weight to approximately:

```text
0.0993706510610444,
```

so the exact-ledger closure is much sharper than the orientation-substituted one.

## Sensitivity Jacobian

The gate records the direct closure sensitivities:

```text
∂E_72/∂kappa_e       = +1
∂E_72/∂|lambda|      = -65/72
∂E_72/∂(R_3-1)       = -7/72
```

and the scalar-deficit sensitivities from

```text
kappa_lambda = 1 - [(lambda_runtime-lambda_proxy)/lambda_proxy]/L.
```

This exposes that full significance requires covariance/uncertainty propagation through low-scale scalar inputs and v1 RG transport.

## Verdict

```text
PASS_GATE661_NONCIRCULAR_CLOSURE_INHERITED
PASS_SCALE_SWEEP_COMPUTED_WITH_V1_TRANSPORT
CONDITIONAL_SUPPORT_CLOSURE_IS_LAMBDA12_SELECTED_IN_V1
PASS_LOCAL_LAMBDA12_PERTURBATION_SWEEP_COMPUTED
CONDITIONAL_SUPPORT_LOCAL_E72_MINIMUM_AT_LAMBDA12_IN_V1_GRID
PASS_WEIGHT_SENSITIVITY_COMPUTED
CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_WEIGHT_ROBUST_IN_V1_EXACT_LEDGER
CONDITIONAL_SUPPORT_ORIENTATION_APPROXIMATION_PERTURBS_BEST_WEIGHT_BUT_REMAINS_BRIDGE_SMALL
PASS_INPUT_SENSITIVITY_JACOBIAN_COMPUTED
FAILED_ROUTE_NO_NATIVE_SCALE_SELECTION_THEOREM
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FAILED_ROUTE_NO_FULL_UNCERTAINTY_PROPAGATION
FAILED_ROUTE_NO_NATIVE_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION
FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM
FIREWALL_PRESERVED_GATE662_SCALE_SWEEP_SENSITIVITY_BOUNDARY
```
