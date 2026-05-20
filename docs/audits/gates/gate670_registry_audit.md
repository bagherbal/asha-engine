# Gate 670 — Oriented Wall-Distance Hyperplane Audit

## Purpose

Gate 669 showed that the active scalar/flavor/gauge bridge uses wall-distance coordinates:

```text
epsilon_e       = charged-lepton wall offset,
|lambda|        = scalar zero-wall depth,
R_3 - 1         = gauge meeting-wall excess.
```

Gate 670 asks whether the active closure is more cleanly written as one signed oriented wall-distance hyperplane on:

```text
kappa_lambda,
kappa_e,
lambda(Lambda_12),
R_3 - 1.
```

This is a bridge-layer wall-balance audit only.  It does not derive boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS, or a native `7/72` theorem.

## Implemented package

```text
pkg/bridge/generation2orientedwalldistancehyperplaneaudit
```

Registered theorem:

```text
generation2orientedwalldistancehyperplaneaudit.Generation2OrientedWallDistanceHyperplaneAuditTheorem()
```

## Signed form

Gate 670 confirms that the positive-distance form:

```text
kappa_lambda + kappa_e
-
[(65/72)|lambda(Lambda_12)| + (7/72)(R_3-1)]
≈ 0
```

is equivalent, because `lambda(Lambda_12)<0`, to the signed wall form:

```text
kappa_lambda + kappa_e
+ (65/72)lambda(Lambda_12)
- (7/72)(R_3-1)
≈ 0.
```

Numerically:

```text
W_72_wall =
kappa_lambda + kappa_e
+ (65/72)lambda(Lambda_12)
- (7/72)(R_3-1)

W_72_wall ≈ 8.52583441346e-10.
```

## HistoryWallBalanceSeal

Gate 670 defines the bridge-layer seal:

```text
HistoryWallBalanceSeal:
  W_72_wall =
  kappa_lambda
  + kappa_e
  + (65/72)lambda(Lambda_12)
  - (7/72)(R_3-1)
  ≈ 0.
```

The wall-coordinate roles are:

```text
kappa_lambda       = scalar low-scale matching wall-deficit coordinate,
kappa_e            = flavor loop-wall deficit coordinate,
lambda(Lambda_12)  = signed scalar zero-wall coordinate,
R_3-1              = signed gauge meeting-wall excess.
```

## Hyperplane normal

The oriented wall hyperplane has normal vector:

```text
(1, 1, 65/72, -7/72)
```

on the ordered coordinate list:

```text
(kappa_lambda, kappa_e, lambda(Lambda_12), R_3-1).
```

The `65/72` and `7/72` split remains the active typed weight from the scalar/flavor/boundary closure:

```text
65/72 + 7/72 = 1.
```

The empirical best interpolation weight remains:

```text
w_best = 0.0972228818894104,
7/72   = 0.0972222222222222,
w_best - 7/72 ≈ 6.5967e-7.
```

## Orientation approximation

Replacing exact `kappa_e` by the OrientationBalance approximation:

```text
kappa_e_orient = sin²(theta13)/4 - J_CKM
               = 0.00550633006471245
```

raises the wall-hyperplane residual to:

```text
≈ 2.77672572133e-6.
```

This is still a structured bridge diagnostic, but the exact `kappa_e` closure is much sharper.

## Hessian-layer firewall

Gate 670 preserves the Gate668/Gate669 layer separation:

```text
|lambda|   = active quartic zero-wall coordinate,
2|lambda|  = scalar Hessian / squared-mass layer.
```

The oriented wall hyperplane uses signed `lambda`, not the Hessian coordinate `2lambda` or `2|lambda|`.

## Verdict

```text
PASS_GATE669_WALL_COORDINATE_AUDIT_INHERITED
PASS_SIGNED_WALL_FORM_WRITTEN
PASS_HISTORY_WALL_BALANCE_FUNCTIONAL_DEFINED
PASS_WALL_COORDINATE_ROLES_CLASSIFIED
PASS_NORMAL_VECTOR_AND_7_OVER_72_WEIGHT_AUDITED
PASS_ORIENTATION_APPROXIMATION_AUDITED
PASS_HESSIAN_LAYER_FIREWALL_PRESERVED
CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_IS_ORIENTED_WALL_DISTANCE_HYPERPLANE
CONDITIONAL_SUPPORT_HISTORY_WALL_BALANCE_SEAL_DEFINED
CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_TYPED_AS_HYPERPLANE_NORMAL_WEIGHT
FAILED_ROUTE_NO_NATIVE_WALL_DISTANCE_AIRLOCK_THEOREM
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION
FAILED_ROUTE_NO_NATIVE_SCALAR_ZERO_BOUNDARY_THEOREM
FAILED_ROUTE_NO_HIGGS_STABILITY_GAUGE_UNIFICATION_FLAVOR_OR_CKM_PMNS_CLAIM
FIREWALL_PRESERVED_GATE670_ORIENTED_WALL_HYPERPLANE_BOUNDARY
```

## Interpretation

Gate 670 upgrades the active bridge from pairwise wall-distance comparison to one signed affine wall-balance equation:

```text
History deficits + oriented boundary-wall coordinates ≈ 0.
```

This is the cleanest current bridge-layer form of the scalar/flavor/gauge closure.  It remains a bridge seal, not a native wall-distance theorem or boundary-stress derivation.
