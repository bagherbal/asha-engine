# Gate 732 — Boundary Raw-Moment Response Coordinate-Naturality Audit

## Purpose

Gate 731 source-typed the cubic coefficient in the residual expansion as:

```text
7/36 = 2p_K7.
```

The active residual-compression expansion is:

```text
D_base
≈
M1_wall
+
kappa_e M2_wall
-
2p_K7 M3_wall.
```

Gate 732 audits whether this expansion is naturally expressed in raw powers of the boundary response operator `R_wall`, or whether variance, central moments, cumulants, or normalized moments provide a better typed coordinate.  This is a bridge-layer moment-coordinate audit only.  It does not derive a native boundary moment expansion theorem, scalar runtime lambda, Higgs mass, Yukawa operators, CKM/PMNS, flavor hierarchy, or a native HistoryLoopUnit theorem.

## Registered theorem

```text
pkg/bridge/generation2boundaryrawmomentresponsecoordinatenaturalityaudit
```

```text
generation2boundaryrawmomentresponsecoordinatenaturalityaudit.Generation2BoundaryRawMomentResponseCoordinateNaturalityAuditTheorem()
```

## Raw response moments

For:

```text
R_wall = S_split P_K7,
p_K7   = Tr(rho_72 P_K7)=7/72,
```

projector idempotence gives:

```text
R_wall^n = S_split^n P_K7.
```

Therefore:

```text
M1 = p_K7 S_split
M2 = p_K7 S_split^2
M3 = p_K7 S_split^3.
```

The Gate 731 expansion becomes:

```text
D_base
≈
p_K7 S_split
+
kappa_e p_K7 S_split^2
-
2p_K7^2 S_split^3.
```

Equivalently:

```text
D_base
≈
p_K7 S_split[1+kappa_e S_split-2p_K7 S_split^2].
```

This conditionally supports the active expansion as a raw scalar response function on `S_split`.

## Projector-power degeneracy

All powers remain supported on the same projector:

```text
R_wall^n ∝ P_K7.
```

Thus the powers do not create independent operator directions.  They only supply scalar powers of the boundary split coordinate.  The moment expansion is therefore source-typed as a scalar response function, not new operator geometry.

## Variance coordinate

The Bernoulli variance coordinate is:

```text
Var_wall = p_K7(1-p_K7)S_split^2
         ≈ 1.46612305634e-7.
```

The residual coefficient in variance units is:

```text
E_wall/Var_wall ≈ 0.00581522428.
```

This is less close to `kappa_e` than the raw second-moment coefficient:

```text
E_wall/M2_wall ≈ 0.00524985525.
```

Variance remains a typed fluctuation scale, but it is not selected as the active residual coordinate in the current ledger.

## Central third moment

The Bernoulli third central moment is:

```text
mu3_wall = p_K7(1-p_K7)(1-2p_K7)S_split^3.
```

Using the same typed cubic coefficient on `mu3_wall` leaves a much larger residual than the raw `M3` route:

```text
raw cubic residual      ≈ -3.88e-13
central moment residual ≈ -1.15e-11.
```

Therefore the Gate 730 compression is specifically raw-moment based.

## Source-type interpretation

Current best source type:

```text
leading:
  raw first expectation M1 = p_K7 S_split

quadratic:
  raw second moment M2 modulated by kappa_e

cubic:
  raw third moment M3 pulled by double event weight 2p_K7
```

Compactly:

```text
D_base
≈
M1 + kappa_e M2 - 2p_K7 M3.
```

## Coordinate-naturality firewall

Gate 732 preserves:

```text
raw moments are selected in the current bridge ledger,
but no native theorem proves that history response must use raw moments.
```

No variance, cumulant, central-moment, normalized-moment, scalar runtime, Higgs mass, or Yukawa theorem follows.

## Verdict

```text
PASS_GATE731_CUBIC_COEFFICIENT_SOURCE_INHERITED
PASS_RAW_MOMENT_RESPONSE_FUNCTION_REWRITTEN
PASS_PROJECTOR_POWER_DEGENERACY_RECORDED
PASS_VARIANCE_COORDINATE_AUDITED
PASS_CENTRAL_THIRD_MOMENT_AUDITED
PASS_RAW_VERSUS_CENTRAL_COMPARISON_AUDITED
PASS_SOURCE_TYPE_INTERPRETATION_RECORDED
PASS_COORDINATE_NATURALITY_FIREWALL_ENFORCED
CONDITIONAL_SUPPORT_ACTIVE_RESIDUAL_EXPANSION_LIVES_IN_RAW_RESPONSE_MOMENTS
CONDITIONAL_SUPPORT_MOMENT_EXPANSION_IS_SCALAR_RESPONSE_FUNCTION_ON_S_SPLIT
CONDITIONAL_SUPPORT_RAW_M3_COORDINATE_BEST_COMPRESSES_CURRENT_RESIDUAL
FAILED_ROUTE_VARIANCE_COORDINATE_NOT_ACTIVE
FAILED_ROUTE_CENTRAL_MOMENT_FORM_NOT_ACTIVE
FAILED_ROUTE_PROJECTOR_POWERS_DO_NOT_SUPPLY_INDEPENDENT_OPERATOR_DIRECTIONS
FAILED_ROUTE_NO_NATIVE_RAW_MOMENT_RESPONSE_COORDINATE_THEOREM
FAILED_ROUTE_NO_NATIVE_BOUNDARY_MOMENT_EXPANSION_THEOREM
FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM
FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM
FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM
FIREWALL_PRESERVED_GATE732_RAW_MOMENT_COORDINATE_BOUNDARY
```
