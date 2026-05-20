# Gate 644 — HodgeProjector Plane MetricRatio Audit

## Purpose

Gate 643 showed that the compact/split residual tensor is not off-sector.  It is same-sector Hodge-diagonal inside the projector plane

```text
span{P_{K7+}, P_{K7-}}.
```

Gate 644 reconstructs the normalized twist metric itself.  Instead of continuing to audit only the scalar angle

```text
cos(theta_twist)=13/sqrt(217),
sin(theta_twist)=4*sqrt(3)/sqrt(217),
```

it asks whether the repeated split-twist metrics collapse to the projector-plane ratio

```text
G_hat = (P_{K7+} - 3 P_{K7-}) / sqrt(31),
```

while the Hodge bilinear ray is

```text
B_hat = (P_{K7+} - P_{K7-}) / sqrt(7).
```

This is an internal finite-geometry audit only.  It does not derive split-G2 structure, boundary stress, scalar/flavor transport, physical spacetime, Higgs mass, CKM/PMNS, gauge unification, or a native `7/72` theorem.

## Gate 643 inheritance

Gate 643 certified the projective residual tensor

```text
R_hat = [G_hat - <G_hat,B_hat>_F B_hat] / rho_twist
```

and found the repeated same-sector block profile

```text
||R_++||_F^2  = 3/7,
||R_--||_F^2  = 4/7,
2||R_+-||_F^2 = 0.
```

Equivalently, in the Gate634 Hodge-polarity frame,

```text
R_hat ≈ -sqrt(3/28) P_{K7+} - sqrt(4/21) P_{K7-}
```

up to route orientation conventions.

## Projector-plane reconstruction

Gate 644 reconstructs

```text
G_hat = cos(theta) B_hat + rho_twist R_hat.
```

Using

```text
cos(theta)=13/sqrt(217),
rho_twist=4*sqrt(3)/sqrt(217),
B_hat=(P_{K7+}-P_{K7-})/sqrt(7),
R_hat=-sqrt(3/28)P_{K7+}-sqrt(4/21)P_{K7-},
```

the diagonal coefficients reduce to

```text
G_hat = (P_{K7+} - 3P_{K7-}) / sqrt(31).
```

The denominator collapse is

```text
217 = 7*31.
```

## Route certificate

Gate 644 repeats the check for the three Gate638/Gate639 routes:

```text
omega_1_alt,
omega_2_alt,
omega_B_alt.
```

After projective sign alignment with `B_hat`, every route certifies the same Hodge-sector metric ratio:

```text
K_7^+ eigenvalue:  +1/sqrt(31),
K_7^- eigenvalue:  -3/sqrt(31),
K_7^+ x K_7^- block: 0.
```

Thus the observed ratio is

```text
K_7^+ : K_7^- = 1 : -3.
```

## Angle consequence

The Gate642 angle now follows from two diagonal projector-plane rays:

```text
B_hat = (P_{K7+}-P_{K7-})/sqrt(7),
G_hat = (P_{K7+}-3P_{K7-})/sqrt(31).
```

Their Frobenius inner product is

```text
<G_hat,B_hat>_F
= [4*(1)(1) + 3*(-3)(-1)] / sqrt(31*7)
= (4+9)/sqrt(217)
= 13/sqrt(217).
```

Therefore

```text
1 - <G_hat,B_hat>_F^2 = 48/217.
```

Gate 644 turns the `169:48:217` obstruction into a simpler projector-plane comparison:

```text
(1,-1) versus (1,-3)
```

on the native Hodge sectors `(4|3)`.

## Remaining source pressure

The new missing theorem is no longer only a trace identity for `169:48:217`.  It is sharper:

```text
Why does the admissible split-twist metric produce the -3 weight on K_7^-?
```

The leading candidate is

```text
-3 = -dim(K_7^-),
```

but Gate 644 does not certify a native source theorem for this weight.

## Verdict

```text
PASS_GATE643_RESIDUAL_TENSOR_BLOCK_STRUCTURE_INHERITED
PASS_GHAT_RECONSTRUCTED_FROM_BHAT_AND_RHAT
PASS_HODGE_PROJECTOR_PLANE_METRICS_DEFINED
PASS_ROUTE_METRIC_RATIOS_COMPUTED
CONDITIONAL_SUPPORT_GTWIST_HAS_HODGE_DIAGONAL_RATIO_1_TO_MINUS_3
CONDITIONAL_SUPPORT_PROJECTIVE_ANGLE_DERIVES_FROM_PROJECTOR_PLANE_GEOMETRY
CONDITIONAL_SUPPORT_MINUS_THREE_WEIGHT_MATCHES_NEGATIVE_HODGE_SECTOR_DIMENSION_CANDIDATE
FAILED_ROUTE_NO_NATIVE_SOURCE_FOR_MINUS_THREE_WEIGHT_YET
FAILED_ROUTE_NO_NATIVE_TRACE_IDENTITY_FOR_PROJECTOR_PLANE_RATIO_YET
FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_PROJECTOR_PLANE_ANGLE_IS_NOT_PHYSICAL_ANGLE
FAILED_ROUTE_PROJECTOR_PLANE_METRIC_IS_NOT_PHYSICAL_METRIC_THEOREM
FIREWALL_PRESERVED_GATE644_PROJECTOR_PLANE_RATIO_IS_INTERNAL_ONLY
```

## Final classification

Gate 644 upgrades the residual-tensor result into a projector-plane metric-ratio audit.  The compact/split obstruction is now expressed as a clean internal comparison between

```text
B_hat: (1,-1) on K_7^+ | K_7^-,
G_hat: (1,-3) on K_7^+ | K_7^-.
```

This explains the projective angle algebraically inside the Hodge projector plane, but it still does not derive why the twist metric has the `-3` negative-sector weight.  All split-G2, boundary-stress, scalar/flavor, physical-geometry, and native `7/72` firewalls remain intact.
