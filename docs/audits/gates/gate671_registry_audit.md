# Gate 671 — HistoryWallBalance Normal-Vector Source and Minimality Audit

## Purpose

Gate 670 defined the active bridge as the oriented wall-distance hyperplane:

```text
W_72 =
kappa_lambda
+ kappa_e
+ (65/72)lambda(Lambda_12)
- (7/72)(R_3-1)
≈ 0.
```

Gate 671 audits the source type, typed minimality, coordinate normalization, and scale-local status of the normal vector:

```text
n_72 = (1, 1, 65/72, -7/72)
```

on:

```text
(kappa_lambda, kappa_e, lambda(Lambda_12), R_3-1).
```

This is a bridge-layer normal-vector audit only.  It does not derive boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS, or a native `7/72` theorem.

## Implemented package

```text
pkg/bridge/generation2historywallbalancenormalvectorsourceaudit
```

Registered theorem:

```text
generation2historywallbalancenormalvectorsourceaudit.Generation2HistoryWallBalanceNormalVectorSourceAndMinimalityAuditTheorem()
```

## Normal vector

Gate 671 decomposes:

```text
n_72 = (1, 1, 65/72, -7/72)
```

as:

```text
history side  = (1,1),
boundary side = (65/72,-7/72).
```

The history side gives unit weight to the scalar matching deficit and flavor wall deficit.  The boundary side gives a scalar-dominant signed wall interpolation with:

```text
65/72 + 7/72 = 1.
```

The gauge wall enters as the signed pull:

```text
-(7/72)(R_3-1).
```

## Typed minimality audit

Gate 671 compares only typed alternatives:

```text
(1,1,1,0)
(1,1,1,-1)
(1,1,7/8,-1/8)
(1,1,9/10,-1/10)
(1,1,65/72,-7/72)
(1,1,63/70,-7/70)
```

using the signed wall form:

```text
kappa_lambda+kappa_e+a lambda(Lambda_12)-b(R_3-1).
```

For exact `kappa_e`, the best typed candidate among this ledger is:

```text
(1,1,65/72,-7/72)
```

with residual:

```text
8.52583441346e-10.
```

Nearby typed alternatives are much weaker in the exact ledger, for example:

```text
(1,1,9/10,-1/10) residual ≈ -3.58927191327e-6.
```

## Exact versus OrientationBalance kappa_e

Replacing exact `kappa_e` by:

```text
kappa_e_orient = sin²(theta13)/4 - J_CKM
               = 0.00550633006471245
```

raises the `n_72` residual to:

```text
2.77672572133e-6.
```

In the OrientationBalance-substituted ledger, nearby typed weights become competitive.  Gate 671 therefore keeps `n_72` as strongest in the exact wall ledger, not as an independently derived flavor theorem.

## Coordinate-normalization audit

Gate 671 classifies the normal vector as coordinate-sealed to the Gate669 canonical wall-distance normalization:

```text
kappa_lambda,
kappa_e,
lambda(Lambda_12),
R_3-1.
```

Arbitrary rescaling of `lambda` or `R_3-1` destroys the weight interpretation.  Thus the normal vector is meaningful as a bridge-layer wall-coordinate object, not a coordinate-free native law.

## Scale-local audit

Gate 671 inherits the Gate662 result that the `n_72` closure is selected at `Lambda_12` in the v1 scale sweep and local perturbation grid.  This supports the wall-normal as a v1 Lambda12-local bridge object, while preserving the no-native-scale-selection theorem firewall.

## Source-type candidates

Gate 671 classifies possible sources:

```text
1. augmented chamber trace:
   7/72 from 7 over 70+2;

2. boundary interpolation:
   7/72 as scalar/gauge wall split weight;

3. history-deficit conservation:
   kappa_lambda+kappa_e balanced against boundary wall projection;

4. coordinate artifact risk:
   the normal is meaningful only in canonical wall-distance coordinates.
```

All remain candidates.  No source theorem is certified.

## Verdict

```text
PASS_GATE670_HISTORY_WALL_BALANCE_INHERITED
PASS_NORMAL_VECTOR_DEFINED
PASS_NORMAL_VECTOR_DECOMPOSITION_AUDITED
PASS_TYPED_ALTERNATIVE_NORMALS_COMPARED
PASS_COORDINATE_NORMALIZATION_AUDITED
PASS_EXACT_VERSUS_ORIENTATION_KAPPA_AUDITED
PASS_SCALE_LOCAL_AUDIT_COMPUTED
PASS_SOURCE_TYPE_CANDIDATES_AUDITED
CONDITIONAL_SUPPORT_N72_IS_BEST_TYPED_WALL_BALANCE_NORMAL_IN_V1
CONDITIONAL_SUPPORT_HISTORY_WALL_BALANCE_NORMAL_IS_COORDINATE_SEALED
CONDITIONAL_SUPPORT_AUGMENTED_CHAMBER_TRACE_SOURCE_CANDIDATE
CONDITIONAL_SUPPORT_BOUNDARY_INTERPOLATION_SOURCE_CANDIDATE
FAILED_ROUTE_NO_NATIVE_NORMAL_VECTOR_SOURCE_THEOREM
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FAILED_ROUTE_NO_NATIVE_WALL_DISTANCE_AIRLOCK_THEOREM
FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION
FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM
FIREWALL_PRESERVED_GATE671_NORMAL_VECTOR_BOUNDARY
```

## Interpretation

Gate 671 strengthens the `HistoryWallBalanceSeal` by showing that its normal vector is the best typed wall-balance normal in the exact v1 ledger among the tested candidates.  It also sharply classifies the normal as coordinate-sealed to the current wall-distance normalization.  The next missing theorem is not another numerical fit; it is a native source theorem for the normal vector or a wall-distance airlock theorem explaining why these are the correct bridge coordinates.
