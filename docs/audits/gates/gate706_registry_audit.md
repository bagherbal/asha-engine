# Gate 706 — Central Scalar Baseline and Uplift-Only Response Isolation Audit

## Purpose

Gate 705 rewrote the positive-distance boundary wound observable as:

```text
W_boundary = |lambda|I_H72 + S_split P_K7.
```

Gate 706 audits the structural separation between the universal scalar baseline
and the support-selected K7 uplift.  It asks whether the scalar baseline is a
central identity shift seen by any normalized observer state, while the
nontrivial history response is entirely carried by:

```text
R_uplift = S_split P_K7.
```

This is a bridge-layer baseline/uplift isolation audit only.  It does not derive
boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
CKM/PMNS, a native response theorem, a native state-selection theorem, or a
native `7/72` theorem.

## Implementation

- Package: `pkg/bridge/generation2centralscalarbaselineandupliftonlyresponseisolationaudit`
- Registered theorem: `generation2centralscalarbaselineandupliftonlyresponseisolationaudit.Generation2CentralScalarBaselineAndUpliftOnlyResponseIsolationAuditTheorem()`

## Central baseline

Define:

```text
B_scalar = |lambda|I_H72.
```

Since it is proportional to the identity:

```text
[B_scalar,P_K7]  = 0
[B_scalar,P_B]   = 0
[B_scalar,P_G]   = 0
[B_scalar,P_perp]= 0.
```

For any normalized density state `rho`:

```text
Tr(rho B_scalar)=|lambda|Tr(rho)=|lambda|.
```

Therefore the scalar baseline is observer-independent and projector-blind.

## Uplift isolation

Subtract the scalar baseline from Gate705's observable:

```text
W_boundary - |lambda|I_H72
= S_split P_K7
= R_uplift.
```

The baseline-subtracted history defect is:

```text
D_base = K_sum-|lambda|
       = kappa_lambda+kappa_e+lambda.
```

Under the full augmented no-bias state:

```text
Tr(rho_72 R_uplift)
= (7/72)S_split.
```

Thus Gate706 isolates the Gate700 response law as the uplift-only part of the
Gate705 positive-distance observable.

## Numerical audit

```text
|lambda|                 ≈ 0.0497009420776833
S_split                  ≈ 0.00129244481881630
(7/72)S_split            ≈ 0.0001256543573849
D_base                   ≈ 0.0001256552099684
residual                 ≈ 8.5258e-10
```

Equivalently:

```text
K_sum ≈ |lambda| + Tr(rho_72 R_uplift).
```

## Observer dependence

The baseline expectation is observer-independent:

```text
Tr(rho B_scalar)=|lambda|.
```

The uplift expectation is observer-dependent:

```text
Tr(rho R_uplift)=S_split Tr(rho P_K7).
```

Typed state comparisons:

```text
rho_72     -> (7/72)S_split
rho_finite -> (7/70)S_split
rho_kernel -> (7/71)S_split
rho_K7     -> S_split
```

Therefore only the uplift sector requires the full no-bias observer-state
selection.

## Support dependence

The baseline does not select `P_K7`; it is a central identity shift.  The uplift
operator is projector-sensitive:

```text
R_uplift=S_split P_K7.
```

The K7 identity still requires the Boolean-octonionic support selector:

```text
rank(P)=7,
P_B P=P,
P_G P=P
=> P=P_K7.
```

## Relation to previous gates

Gate705 positive-distance form:

```text
K_sum ≈ Tr(rho_72 W_boundary).
```

Gate706 baseline-subtracted form:

```text
D_base=K_sum-|lambda|
≈ Tr(rho_72[W_boundary-|lambda|I_H72])
= Tr(rho_72 S_split P_K7).
```

Gate706 introduces no new numerical relation.  It isolates the active response
by subtracting the central scalar baseline.

## Verdict

```text
PASS_GATE705_SCALAR_BASELINE_K7_UPLIFT_INHERITED
PASS_CENTRAL_BASELINE_OPERATOR_DEFINED
PASS_BASELINE_COMMUTES_WITH_PROJECTOR_ALGEBRA
PASS_BASELINE_EXPECTATION_OBSERVER_INDEPENDENT
PASS_UPLIFT_OPERATOR_ISOLATED
PASS_BASELINE_SUBTRACTED_RESPONSE_RECONSTRUCTED
PASS_OBSERVER_DEPENDENCE_LOCALIZED_TO_UPLIFT
PASS_SUPPORT_DEPENDENCE_LOCALIZED_TO_UPLIFT
PASS_RELATION_TO_PREVIOUS_GATES_AUDITED
CONDITIONAL_SUPPORT_SCALAR_BASELINE_IS_CENTRAL_IDENTITY_SHIFT
CONDITIONAL_SUPPORT_NONTRIVIAL_BRIDGE_CONTENT_IS_K7_UPLIFT_RESPONSE
CONDITIONAL_SUPPORT_DBASE_IS_BASELINE_SUBTRACTED_HISTORY_DEFECT
FAILED_ROUTE_BASELINE_DOES_NOT_SELECT_K7_OR_RHO72
FAILED_ROUTE_NO_NATIVE_REASON_SCALAR_WOUND_IS_FULL_CHAMBER_BASELINE
FAILED_ROUTE_NO_NATIVE_REASON_K7_RECEIVES_SPLIT_UPLIFT
FAILED_ROUTE_NO_NATIVE_BOUNDARY_WOUND_UPLIFT_THEOREM
FAILED_ROUTE_NO_NATIVE_HISTORY_RESPONSE_THEOREM
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FIREWALL_PRESERVED_GATE706_CENTRAL_BASELINE_UPLIFT_ISOLATION_BOUNDARY
```
