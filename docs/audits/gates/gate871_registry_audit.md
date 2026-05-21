# Gate 871 — BoundaryExterior Degree-Target Selection Map Audit

## Purpose

Gate 871 follows Gate 870's reduced boundary-pair exterior response audit. Gate
870 sharpened the alpha response shape to the reduced exterior form:

```text
R_B(s) = (1+s b1)(1+s b2)-1
       = s(b1+b2)+s^2(b1 wedge b2)
```

This conditionally explains the absence of a constant term and the absence of
cubic/higher powers, because `Lambda^3 B_2=0`. Gate 871 audits the remaining
wound: whether the exterior degrees have lawful targets

```text
Lambda^1 B_2 -> Pi_top
Lambda^2 B_2 -> H_R^min
```

rather than being assigned by hand.

This is a degree-target selection audit only. It does not derive `alpha_B`, does
not certify a native reduced boundary response functional, does not prove a
BoundaryAlphaTransportMap, does not update `N_eff`, `C_Yukawa`, or `C_Higgs`,
and does not promote the branch to R3 or R4.

## Implemented package

```text
pkg/bridge/generation2boundaryexteriordegreetargetselectionmapaudit
```

Registered theorem:

```text
generation2boundaryexteriordegreetargetselectionmapaudit.Generation2BoundaryExteriorDegreeTargetSelectionMapAuditTheorem()
```

## Inherited objects

```text
R_B(s)=(1+s b1)(1+s b2)-1
      =s(b1+b2)+s^2(b1 wedge b2)
```

```text
Pi_top = e_+ tensor P_3
rank(Pi_top)=3
```

```text
H_R^min = (C_R^2 tensor W) minus (e_+ tensor P_1)
rank(H_R^min)=7
```

```text
H10 = H_R^ambient plus B_2
rank(H_R^ambient)=8
rank(B_2)=2
rank(H10)=10
```

```text
H72 = Lambda^4 V_8 plus B_2
rank(Lambda^4 V_8)=70
rank(B_2)=2
rank(H72)=72
```

## Candidate degree-target selection

Gate 871 audits the candidate assignments:

```text
Lambda^1 B_2 -> Pi_top
```

and:

```text
Lambda^2 B_2 -> H_R^min
```

These reconstruct:

```text
alpha_B = [rank(Pi_top)/10] s + [rank(H_R^min)/72] s^2
        = (3/10)s + (7/72)s^2
```

with:

```text
alpha_B = 0.0003878958469680527
```

## Cross-lane exclusions

The gate explicitly audits the missing exclusions:

```text
Lambda^1 B_2 -> H_R^min
```

and:

```text
Lambda^2 B_2 -> Pi_top
```

If these were allowed, extra terms such as `7s/72` or `3s^2/10` would appear.
Gate 871 therefore preserves the firewall:

```text
FAILED_ROUTE_NO_CROSS_LANE_EXCLUSION_THEOREM
```

## Verdict

Gate 871 conditionally supports the candidate target logic:

```text
CONDITIONAL_SUPPORT_DEGREE_ONE_TARGETS_DOMINANT_SOCKET_CANDIDATE
CONDITIONAL_SUPPORT_DEGREE_TWO_TARGETS_ACTIVE_PUNCTURED_RIGHT_DOMAIN_CANDIDATE
CONDITIONAL_SUPPORT_DEGREE_ONE_AS_SINGLE_BOUNDARY_DOMINANT_SOCKET_RESPONSE
CONDITIONAL_SUPPORT_DEGREE_TWO_AS_FULL_BOUNDARY_PAIR_ACTIVE_DOMAIN_RESPONSE
CONDITIONAL_SUPPORT_ALPHA_B_WOUND_REDUCED_TO_DEGREE_TARGET_SELECTION
```

but blocks native promotion:

```text
FAILED_ROUTE_NO_NATIVE_DEGREE_TARGET_SELECTION_MAP
FAILED_ROUTE_NO_NATIVE_LAMBDA1B2_TO_PI_TOP_MAP
FAILED_ROUTE_NO_NATIVE_LAMBDA2B2_TO_H_R_MIN_MAP
FAILED_ROUTE_NO_CROSS_LANE_EXCLUSION_THEOREM
FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_LINEAR_H_R_MIN_TERM
FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_QUADRATIC_PI_TOP_TERM
FAILED_ROUTE_RESPONSE_CHAMBER_TYPING_NOT_TARGET_SELECTION_THEOREM
FAILED_ROUTE_ALPHA_B_REMAINS_SEALED
FAILED_ROUTE_NOT_R3_BOUNDARY_EXTERIOR_DEGREE_TARGET_SELECTION_OBSTRUCTION
```

## Final classification

```text
R2+++++_BOUNDARY_EXTERIOR_DEGREE_TARGET_SELECTION_OBSTRUCTION
```

The reduced exterior response gives the correct `s+s^2` shape, but the branch
still lacks a native map assigning degree one to `Pi_top` and degree two to
`H_R^min`, as well as a theorem excluding the cross-lanes.
