# Gate 632 — Hodge-Star K7-to-W7 Leakage Rank Audit

## Purpose

Gate 631 sharpened the index-zero defect pairing problem to the explicit candidate

```text
Phi_* = P_W * |_{K_7}: K_7 -> W_7,
```

where

```text
H = Lambda^4 R^8,
U = Im(P_B),
V = Im(P_G),
K_7 = U ∩ V,
W_7 = (U+V)^perp.
```

Gate 632 performs the actual finite matrix rank test.  This is a native finite-linear-algebra audit only.  It does not derive boundary stress, scalar RG matching, flavor, Higgs mass, CKM/PMNS, gauge unification, or a native trace-weight theorem.

## Inherited data

```text
H = Lambda^4 R^8, dim H = 70
U = Im(P_B), dim U = 56
V = Im(P_G), dim V = 14
K_7 = U ∩ V, dim K_7 = 7
U+V has dim 63
W_7 = (U+V)^perp, dim W_7 = 7
```

Gate 631 exact sequence:

```text
0 -> K_7 -> U⊕V -> H -> W_7 -> 0.
```

Candidate pairing:

```text
Phi_* = P_W * |_{K_7}.
```

Status inherited:

```text
PASS_GATE631_PAIRING_PROBLEM_INHERITED
```

## A. Hodge-star matrix audit

Gate 632 constructs the Hodge star on the lexicographic oriented wedge basis of `Lambda^4 R^8`:

```text
*(e_I) = sgn(I,I^c) e_{I^c},
```

relative to

```text
e_0 ∧ e_1 ∧ ... ∧ e_7.
```

Computed certificate:

```text
dim(* matrix) = 70 x 70
*^2 residual = 0
trace(*) = 0
Lambda^4 R^8 = Lambda^4_+ ⊕ Lambda^4_-
dim Lambda^4_+ = 35
dim Lambda^4_- = 35
```

Status:

```text
PASS_HODGE_STAR_OPERATOR_TYPED_ON_LAMBDA4_R8
```

## B. Basis certificate

Gate 632 constructs orthonormal matrices:

```text
Q_K: columns span K_7
Q_W: columns span W_7
```

and verifies:

```text
Q_K^T Q_K ≈ I_7
Q_W^T Q_W ≈ I_7
Q_W^T Q_K ≈ 0
P_B Q_K ≈ Q_K
P_G Q_K ≈ Q_K
P_B Q_W ≈ 0
P_G Q_W ≈ 0
```

Computed residuals:

```text
||Q_K^T Q_K - I||_F ≈ 2.48e-16
||Q_W^T Q_W - I||_F ≈ 2.16e-15
||Q_W^T Q_K||_F ≈ 5.24e-14
||P_B Q_K - Q_K||_F ≈ 3.95e-14
||P_G Q_K - Q_K||_F = 0
||P_B Q_W||_F ≈ 1.76e-13
||P_G Q_W||_F ≈ 7.82e-14
```

Status:

```text
PASS_K7_AND_W7_BASES_CERTIFIED
```

## C. Leakage rank test

Gate 632 computes

```text
M_* = Q_W^T * Q_K.
```

Computed leakage table:

```text
size(M_*) = 7 x 7
rank(M_*) = 0
singular values ≈ [
  3.3152e-14,
  2.7184e-14,
  1.8878e-14,
  1.4669e-14,
  1.4027e-14,
  1.0408e-14,
  5.6382e-15
]
||M_*||_F ≈ 5.2406e-14
det(M_*) = 0
```

Therefore the Hodge-star route fails as a nondegenerate `K_7 -> W_7` pairing:

```text
FAILED_ROUTE_HODGE_STAR_DOES_NOT_PAIR_K7_TO_W7
FAILED_ROUTE_NO_CANONICAL_K7_W7_PAIRING_FOUND
```

## D. Image containment audit

Gate 632 checks where `*K_7` lands:

```text
||*K_7||_F ≈ 2.64575131106459
||P_W * K_7||_F ≈ 2.1961e-13
||P_{U+V} * K_7||_F ≈ 2.64575131106459
```

So, to numerical precision,

```text
*K_7 ⊂ U+V,
P_W * K_7 ≈ 0.
```

This kills the clean Hodge leakage path.

## E. Pairing metric audit

The induced metric

```text
G_* = M_*^T M_*
```

is degenerate:

```text
trace(G_*) ≈ 2.7464e-27
scale candidate ≈ 3.9235e-28
rank full = false
```

No conformal, isometric, or anisotropic nondegenerate `K_7 -> W_7` pairing is certified.

## F. Orientation audit

Since

```text
det(M_*) = 0,
```

Gate 632 has no orientation sign to promote.  No physical orientation or boundary orientation is inferred.

## G. Alternative star/projector composites

Gate 632 also tests typed variants:

```text
P_W * P_B |_{K_7}
P_W * P_G |_{K_7}
P_W [*,P_B] |_{K_7}
P_W [*,P_G] |_{K_7}
P_W (*P_B - *P_G) |_{K_7}
```

All tested variants have rank zero at the same tolerance.  This matches the structural obstruction:

```text
P_B k = k,
P_G k = k,
P_W(U+V)=0,
```

for `k in K_7`.

Status:

```text
FAILED_ROUTE_ALTERNATIVE_STAR_PROJECTOR_COMPOSITES_DO_NOT_PAIR_K7_TO_W7
```

## H. Boundary readiness audit

Because the Hodge-star pairing fails, Gate 632 does not even reach a certified `K_7 -> W_7` pairing.  A future route would still need one of:

```text
W_7 -> R^2_boundary,
K_7/W_7 defect trace -> R^2_boundary,
eta/signed operator on Lambda^4 R^8,
determinant-line or product-airlock bridge.
```

Status:

```text
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT_YET
```

## Final verdict

```text
PASS_GATE631_PAIRING_PROBLEM_INHERITED
PASS_HODGE_STAR_OPERATOR_TYPED_ON_LAMBDA4_R8
PASS_K7_AND_W7_BASES_CERTIFIED
PASS_HODGE_LEAKAGE_MATRIX_COMPUTED
FAILED_ROUTE_HODGE_STAR_DOES_NOT_PAIR_K7_TO_W7
FAILED_ROUTE_NO_CANONICAL_K7_W7_PAIRING_FOUND
FAILED_ROUTE_ALTERNATIVE_STAR_PROJECTOR_COMPOSITES_DO_NOT_PAIR_K7_TO_W7
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT_YET
FIREWALL_PRESERVED_GATE632_HODGE_PAIRING_BOUNDARY
```

## Interpretation

Gate 632 is a productive failure.  It does not leave the pairing problem vague.  It proves, by explicit finite matrix computation, that the most natural native middle-form operator does **not** send the `K_7` defect into the transverse `W_7` defect:

```text
rank(Q_W^T * Q_K)=0.
```

Thus the next missing object is not ordinary Hodge star.  The project must look for a different typed operator, an eta/signed structure on `Lambda^4 R^8`, a determinant-line pairing, or a product-airlock map before the `7/72` boundary pull can be promoted beyond a bridge-layer dimension-compression seal.
