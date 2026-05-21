# Gate 870 — Reduced BoundaryPair Exterior Response and Degree-Target Selection Audit

## Purpose

Gate 870 follows Gate 869's boundary-pair exterior truncation obstruction. Gate
869 showed that `Lambda^bullet B_2` has the right finite degree shape for the
`alpha_B` power pattern, because `Lambda^3 B_2=0`. But the full exterior algebra
still contains the zero-degree lane `Lambda^0 B_2`, so exterior truncation alone
cannot explain why the alpha response has no constant term.

Gate 870 audits the sharper candidate: use the reduced boundary-pair exterior
response

```text
R_B(s) = (1+s b1)(1+s b2)-1
       = s(b1+b2)+s^2(b1 wedge b2).
```

This suppresses the zero-order term by construction and still has no cubic or
higher term because the boundary carrier has rank two.

This is a response-shape and degree-target audit only. It does not derive
`alpha_B`, certify a native boundary response functional, prove degree-target
selection, update `N_eff`, `C_Yukawa`, or `C_Higgs`, or promote the branch to R3
or R4.

## Inherited data

```text
s = S_split = 0.0012924448188162962
alpha_B = 0.0003878958469680527

Pi_top = e_+ tensor P_3
rank(Pi_top) = 3

H_R^ambient = C_R^2 tensor W
rank(H_R^ambient) = 8
H10 = H_R^ambient plus B_2
rank(H10)=10

H_R^min = (C_R^2 tensor W) minus (e_+ tensor P_1)
rank(H_R^min)=7

H72 = Lambda^4 V_8 plus B_2
rank(H72)=70+2=72
```

## Boundary-pair exterior response

For a rank-two boundary pair `B_2` with generators `b1,b2`:

```text
Lambda^0 B_2 dim = 1
Lambda^1 B_2 dim = 2
Lambda^2 B_2 dim = 1
Lambda^3 B_2 dim = 0
```

The reduced response is:

```text
R_B(s) = (1+s b1)(1+s b2)-1
       = s(b1+b2)+s^2(b1 wedge b2).
```

This conditionally supports:

```text
CONDITIONAL_SUPPORT_ZERO_ORDER_TERM_SUPPRESSED_BY_REDUCED_RESPONSE
CONDITIONAL_SUPPORT_NO_CUBIC_OR_HIGHER_TERMS_FROM_LAMBDA3_B2_EQUALS_ZERO
CONDITIONAL_SUPPORT_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_HAS_EXACT_S_PLUS_S_SQUARED_SHAPE
```

## Formal alpha reconstruction

With candidate target maps:

```text
Lambda^1 B_2 -> Pi_top in H10
Lambda^2 B_2 -> H_R^min in H72
```

Gate 870 reconstructs:

```text
alpha_B = (1/10) Tr_H10(Pi_top sI)
        + (1/72) Tr_H72(P_HRmin s^2 I)

        = (3/10)s + (7/72)s^2.
```

This exactly reproduces the active `alpha_B` value, but only as a reduced exterior response candidate.

## Remaining obstruction

Gate 870 does not certify the target maps:

```text
FAILED_ROUTE_NO_TYPED_DEGREE_ONE_TO_PI_TOP_MAP
FAILED_ROUTE_NO_TYPED_DEGREE_TWO_TO_H_R_MIN_MAP
FAILED_ROUTE_NO_DEGREE_TARGET_SELECTION_THEOREM
```

It also does not certify a native reduced response functional:

```text
FAILED_ROUTE_NO_NATIVE_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_FUNCTIONAL
FAILED_ROUTE_REDUCED_RESPONSE_NOT_NATIVE_WITHOUT_BOUNDARY_FUNCTIONAL
```

Cross-lane exclusions remain unproved:

```text
FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_LINEAR_H_R_MIN_TERM
FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_QUADRATIC_PI_TOP_TERM
```

## Verdict

Gate 870 improves the Gate 869 obstruction:

```text
constant term: conditionally explained by R_B(s)=E_B(s)-1
cubic and higher terms: conditionally explained by Lambda^3 B_2=0
degree-target selection: still blocked
native boundary functional: still blocked
alpha_B: still sealed
R3/R4: blocked
```

Final classification:

```text
R2+++++_REDUCED_BOUNDARY_PAIR_EXTERIOR_RESPONSE_OBSTRUCTION
```

The next wound is not the alpha polynomial shape. It is the missing native map:

```text
Lambda^1 B_2 -> Pi_top
Lambda^2 B_2 -> H_R^min
```

and the missing boundary response functional that would make those target maps lawful rather than declared.
