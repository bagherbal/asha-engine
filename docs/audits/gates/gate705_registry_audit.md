# Gate 705 — Scalar Baseline and K7 Boundary-Split Uplift Observable Audit

## Purpose

Gate 704 typed the positive-distance closure as the no-bias expectation of the
K7/complement two-payoff boundary wound observable:

```text
W_boundary = (R_3-1)P_K7 + |lambda(Lambda_12)|P_perp.
```

Gate 705 audits the equivalent decomposition:

```text
W_boundary = |lambda|I_H72 + S_split P_K7,
S_split = (R_3-1)-|lambda| = lambda+(R_3-1).
```

This is a bridge-layer observable-decomposition audit only.  It does not derive
boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
CKM/PMNS, a native response theorem, a native state-selection theorem, or a
native `7/72` theorem.

## Implementation

- Package: `pkg/bridge/generation2scalarbaselineandk7boundarysplitupliftobservableaudit`
- Registered theorem: `generation2scalarbaselineandk7boundarysplitupliftobservableaudit.Generation2ScalarBaselineAndK7BoundarySplitUpliftObservableAuditTheorem()`

## Inherited objects

```text
rho_72 = I_H72/72
P_perp = I_H72-P_K7
lambda = lambda(Lambda_12)<0
|lambda| = -lambda
R = R_3-1
S_split = lambda+R = R-|lambda|
K_sum = kappa_lambda+kappa_e
```

## Gate704 observable

```text
W_boundary = R P_K7 + |lambda|P_perp.
```

Under `rho_72`:

```text
Tr(rho_72 W_boundary)
= (7/72)R + (65/72)|lambda|.
```

This matches `K_sum` up to the inherited residual.

## Baseline-plus-uplift decomposition

Using `P_perp=I_H72-P_K7`:

```text
W_boundary
= R P_K7 + |lambda|(I_H72-P_K7)
= |lambda|I_H72 + (R-|lambda|)P_K7
= |lambda|I_H72 + S_split P_K7.
```

Thus the sharper source typing is:

```text
full H72 chamber -> scalar zero-wall baseline |lambda|
K7 support       -> boundary split uplift S_split
K7 payoff        -> |lambda|+S_split = R_3-1
```

So `K7` does not need to be described as primitively receiving the gauge wound.
It receives the boundary split uplift over the scalar baseline.

## Expectation audit

```text
Tr(rho_72 W_boundary)
= Tr(rho_72 |lambda|I_H72) + Tr(rho_72 S_split P_K7)
= |lambda| + (7/72)S_split.
```

Numerically:

```text
|lambda|            ≈ 0.0497009420776833
(7/72)S_split       ≈ 0.0001256543573849
sum                 ≈ 0.0498265964350682
K_sum               ≈ 0.0498265972876517
residual            ≈ 8.5258e-10
```

## Relation to Gate700

Gate700 response law:

```text
D_base ≈ (7/72)S_split.
```

Because `lambda<0`:

```text
K_sum = |lambda| + D_base.
```

Therefore Gate705 positive-distance law is not a new numerical relation:

```text
K_sum ≈ |lambda| + (7/72)S_split.
```

It is Gate700 plus the scalar baseline identity.

## Alternative decompositions

- `R I_H72 - S_split P_perp`: algebraically equivalent, but less natural because
  Gate703 identifies the signed scalar wall as the active airlock/baseline.
- `xi_boundary I_H72` plus signed corrections: less minimal and not the active
  quotient form.
- `|lambda|I_H72 + S_split(P_+-P_-)`: gives the inactive `1/72` signed-polarity
  split response.
- `|lambda|I_H72 + S_split P_K7`: active scalar-baseline K7-uplift form.

## Verdict

```text
PASS_GATE704_BOUNDARY_WOUND_MIXTURE_INHERITED
PASS_TWO_PAYOFF_OBSERVABLE_REWRITTEN
PASS_SCALAR_BASELINE_PLUS_K7_UPLIFT_DECOMPOSITION_COMPUTED
PASS_EXPECTATION_REPRODUCES_KSUM_CLOSURE
PASS_RELATION_TO_GATE700_RESPONSE_LAW_AUDITED
PASS_SOURCE_TYPE_UPGRADE_AUDITED
PASS_ALTERNATIVE_BASELINE_DECOMPOSITIONS_AUDITED
CONDITIONAL_SUPPORT_KSUM_IS_SCALAR_BASELINE_PLUS_EXPECTED_K7_SPLIT_UPLIFT
CONDITIONAL_SUPPORT_K7_RECEIVES_BOUNDARY_SPLIT_UPLIFT_NOT_PRIMITIVE_GAUGE_WOUND
CONDITIONAL_SUPPORT_SCALAR_WALL_AIRLOCK_SUPPORTS_SCALAR_BASELINE_READING
FAILED_ROUTE_NO_NATIVE_REASON_SCALAR_WOUND_IS_FULL_CHAMBER_BASELINE
FAILED_ROUTE_NO_NATIVE_REASON_K7_RECEIVES_SPLIT_UPLIFT
FAILED_ROUTE_NO_NATIVE_BOUNDARY_WOUND_UPLIFT_THEOREM
FAILED_ROUTE_NO_NATIVE_HISTORY_RESPONSE_THEOREM
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FIREWALL_PRESERVED_GATE705_SCALAR_BASELINE_K7_UPLIFT_BOUNDARY
```
