# Gate 663 — BoundaryWeightedDeficitClosure Stationarity and Beta-Balance Audit

## Purpose

Gate 662 showed that the active boundary-weighted deficit closure

```text
E_72(mu) = kappa_lambda + kappa_e - W_72(mu)
```

is selected by `Lambda_12` in the current v1 scale sweep.  Gate 663 asks whether the selection is caused by a true stationarity / beta-balance condition at `Lambda_12`, or by a sharp near-zero crossing of the closure function.

This is a bridge-layer stationarity and beta-balance audit only.  It does not derive Higgs mass, scalar stability, flavor, CKM/PMNS, gauge unification, boundary stress, or a native `7/72` theorem.

## Implementation

- Package: `pkg/bridge/generation2boundaryweighteddeficitclosurestationarityaudit`
- Theorem: `Generation2BoundaryWeightedDeficitClosureStationarityAndBetaBalanceAuditTheorem()`
- Runtime marker: `gate663-boundary-weighted-deficit-closure-stationarity-beta-balance-audit-20260518`

## Scale function

Gate 663 keeps the Gate662 v1 convention:

```text
W_72(mu)=|lambda(mu)|+(7/72)[G(mu)-|lambda(mu)|]
G(mu)=g3(mu)/((g1(mu)+g2(mu))/2)-1
E_72(mu)=K_sum-W_72(mu).
```

At `Lambda_12`, this reproduces:

```text
lambda(Lambda_12) = -0.0497009420776833
G(Lambda_12)      = R_3-1 = 0.0509933868964996
E_72(Lambda_12)   ≈ 8.53e-10.
```

## First derivative and beta-balance

The derivative audit computes:

```text
dE_72/dln(mu) ≈ +9.54918e-4
```

at `Lambda_12`.  This is not close to zero relative to the closure residual.  Therefore Gate 663 classifies the v1 alignment as a sharp near-zero crossing, not as a stationary beta-balance point.

In derivative form:

```text
dE_72/dt = -[(65/72)d|lambda|/dt + (7/72)dG/dt],  t=ln(mu).
```

The v1 values are approximately:

```text
d|lambda|/dt ≈ +6.41764e-4
dG/dt        ≈ -1.57813e-2
```

so the weighted gauge-residual slope dominates the weighted scalar-wound slope.  A true stationarity condition would require a different slope balance.

## Zero-scale offset

Solving the v1 closure equation gives a zero very close to the electroweak meeting point:

```text
ln(mu_zero/Lambda_12) ≈ -8.93e-7
mu_zero/Lambda_12    ≈ 0.999999107.
```

Thus the closure zero is aligned with `Lambda_12` in the exact v1 ledger, but the curve is crossing the zero rather than becoming stationary there.

## Local shape

The local curvature is small compared with the first derivative:

```text
d²E_72/dt² ≈ 7.97e-5.
```

The resulting threshold widths are controlled mainly by the slope:

```text
|E_72| < 1e-6  width ≈ 0.00209 in log scale
|E_72| < 1e-5  width ≈ 0.0209 in log scale
|E_72| < 1e-4  width ≈ 0.209 in log scale
```

This confirms a sharp crossing rather than a flat basin.

## Best weight versus scale

At `Lambda_12`, the best interpolation weight remains close to `7/72`:

```text
w_best - 7/72 ≈ 6.60e-7.
```

At local shifts `±0.1`, the best weight moves substantially away from `7/72`, so the weight alignment is also scale-sharp in the v1 ledger.

## Verdict

```text
PASS_GATE662_SCALE_SWEEP_INHERITED
PASS_E72_SCALE_FUNCTION_DEFINED
PASS_FIRST_DERIVATIVE_AUDITED
PASS_BETA_BALANCE_EQUATION_COMPUTED
CONDITIONAL_SUPPORT_LAMBDA12_IS_ZERO_CROSSING_NOT_STATIONARY
CONDITIONAL_SUPPORT_CLOSURE_ZERO_ALIGNED_WITH_ELECTROWEAK_MEETING_SCALE_IN_V1
PASS_CURVATURE_OR_LOCAL_SHAPE_AUDITED
PASS_BEST_WEIGHT_VERSUS_SCALE_AUDITED
CONDITIONAL_SUPPORT_LAMBDA12_SELECTED_BY_V1_CLOSURE
FAILED_ROUTE_NO_NATIVE_SCALE_SELECTION_THEOREM
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FAILED_ROUTE_NO_FULL_UNCERTAINTY_PROPAGATION
FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM
FIREWALL_PRESERVED_GATE663_STATIONARITY_BETA_BALANCE_BOUNDARY
```
