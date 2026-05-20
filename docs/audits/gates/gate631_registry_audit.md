# Gate 631 — Orthogonal Cokernel Representative and K7 Defect Pairing Audit

## Purpose

Gate 630 constructed the square finite addition map

```text
A: U⊕V -> Lambda^4 R^8,
A(u,v)=u+v,
```

with

```text
ker(A) ≅ K_7,
dim ker(A)=7,
dim coker(A)=7,
index(A)=0.
```

Gate 631 sharpens the cokernel and pairing problem.  It asks whether the cokernel has a canonical orthogonal representative

```text
W_7 = (U+V)^perp
```

inside

```text
H = Lambda^4 R^8,
```

and whether any typed native operator on `H` gives a nondegenerate map

```text
Phi: K_7 -> W_7.
```

This is a finite-operator pairing audit only.  It does not derive boundary stress, scalar RG matching, Higgs mass, flavor, CKM/PMNS, gauge unification, or a native trace theorem.

## Inherited data

```text
H = Lambda^4 R^8, dim H = 70
U = Im(P_B), dim U = 56
V = Im(P_G), dim V = 14
K_7 = U ∩ V, dim K_7 = 7

dim(U+V)=56+14-7=63

A: U⊕V -> H,
A(u,v)=u+v

ker(A)=K_7 diagonal anti-copy
coker(A)=H/(U+V), dim coker(A)=7
index(A)=0
```

The Gate626/Gate628 bridge chamber remains

```text
Lambda^4 R^8 ⊕ R^2_boundary,
```

with boundary pair

```text
(|lambda(Lambda_12)|, R_3-1),
```

but this pair is still bridge/environmental, not native finite algebra.

## A. Orthogonal cokernel representative

Using the ambient metric on `H`, define

```text
P_{U+V} = orthogonal projector onto U+V,
P_W = I - P_{U+V},
W_7 = im(P_W) = (U+V)^perp.
```

Since

```text
dim H = 70,
dim(U+V)=63,
```

we get

```text
dim W_7 = 70-63 = 7.
```

The representative satisfies

```text
W_7 ⟂ U,
W_7 ⟂ V,
H = (U+V) ⊕ W_7,
H/(U+V) ≅ W_7.
```

Status:

```text
PASS_ORTHOGONAL_COKERNEL_REPRESENTATIVE_DEFINED
CONDITIONAL_SUPPORT_COKERNEL_REPRESENTED_BY_W7_ORTHOGONAL_COMPLEMENT
```

The conditional wording is important: the metric gives a clean representative of the quotient, but not yet a canonical `K_7 -> W_7` pairing.

## B. Exact defect sequence

Gate 631 writes the exact finite sequence

```text
0 -> K_7 -> U⊕V -> H -> W_7 -> 0
```

where

```text
K_7 -> U⊕V:  k -> (k,-k),
U⊕V -> H:   A(u,v)=u+v,
H -> W_7:   P_W.
```

Exactness checks:

```text
ker(A) = im(K_7 -> U⊕V),
im(A)=U+V=ker(P_W),
im(P_W)=W_7.
```

Dimension check:

```text
dim K_7 - dim(U⊕V) + dim H - dim W_7
= 7 - 70 + 70 - 7
= 0.
```

Status:

```text
PASS_EXACT_DEFECT_SEQUENCE_WRITTEN
```

## C. Hodge-star pairing audit

Candidate:

```text
Phi_* = P_W * |_{K_7},
```

where `*` is the Hodge star on `Lambda^4 R^8`.

The Hodge star is a typed operation on middle-degree forms:

```text
*: Lambda^4 R^8 -> Lambda^4 R^8.
```

However, Gate 631 does not have an explicit basis-level rank computation proving that

```text
rank(P_W * |_{K_7}) = 7.
```

Nor does it prove that `*` preserves or exchanges the `U,V,W_7` decomposition in the required way.

Status:

```text
CONDITIONAL_SUPPORT_HODGE_STAR_PAIRING_REQUIRES_EXPLICIT_RANK_TEST
```

## D. Projector algebra pairing audit

Simple projector candidates fail:

```text
P_W P_B |_{K_7}
P_W P_G |_{K_7}
P_W(P_B-P_G)|_{K_7}
P_W[P_B,P_G]|_{K_7}
P_W(P_B+P_G)|_{K_7}
```

Reason: for `k in K_7`,

```text
P_B k = k,
P_G k = k,
k in U+V,
P_W k = 0.
```

So the simple projector algebra either keeps `k` inside `U+V`, annihilates it before projection, or has no certified transverse rank-seven component.

Status:

```text
FAILED_ROUTE_PROJECTOR_ALGEBRA_DOES_NOT_PAIR_K7_TO_W7
```

## E. Eta / signed pairing audit

Candidate:

```text
Phi_eta = P_W eta |_{K_7}.
```

ASHA has eta/signed-trace lanes elsewhere, but Gate 631 does not have a typed `eta` operator acting on

```text
H = Lambda^4 R^8
```

with certified rank from `K_7` to `W_7`.

Status:

```text
FAILED_ROUTE_NO_CANONICAL_K7_TO_W7_PAIRING_YET
```

## F. Determinant-line audit

The exact sequence gives a determinant-line relation:

```text
det(K_7) ⊗ det(H) ≅ det(U⊕V) ⊗ det(W_7).
```

This is useful volume/orientation bookkeeping and may become relevant for a normalized trace or orientation constraint.  But it is not a pointwise vector-space isomorphism

```text
K_7 -> W_7,
```

and it does not by itself justify the `7/72` boundary weight.

Status:

```text
CONDITIONAL_SUPPORT_K7_W7_PAIRING_PROBLEM_SHARPENED
```

## G. Boundary readiness audit

Even if a future gate certifies a nondegenerate pairing

```text
K_7 -> W_7,
```

one more bridge map is still required:

```text
W_7 -> R^2_boundary
```

or

```text
K_7/W_7 defect trace -> R^2_boundary.
```

Without this map, Gate626's boundary-stress pull remains a bridge-layer diagnostic, not a native transport theorem.

Status:

```text
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT_YET
```

## Final verdict

```text
PASS_GATE630_INDEX_ZERO_OPERATOR_INHERITED
PASS_ORTHOGONAL_COKERNEL_REPRESENTATIVE_DEFINED
PASS_EXACT_DEFECT_SEQUENCE_WRITTEN
CONDITIONAL_SUPPORT_COKERNEL_REPRESENTED_BY_W7_ORTHOGONAL_COMPLEMENT
CONDITIONAL_SUPPORT_K7_W7_PAIRING_PROBLEM_SHARPENED
FAILED_ROUTE_NO_CANONICAL_K7_TO_W7_PAIRING_YET
FAILED_ROUTE_PROJECTOR_ALGEBRA_DOES_NOT_PAIR_K7_TO_W7
CONDITIONAL_SUPPORT_HODGE_STAR_PAIRING_REQUIRES_EXPLICIT_RANK_TEST
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT_YET
FIREWALL_PRESERVED_GATE631_K7_COKERNEL_PAIRING_BOUNDARY
```

## Missing object

Gate 630's missing map

```text
Phi: ker(A) -> coker(A)
```

is sharpened to

```text
Phi: K_7 -> W_7 = (U+V)^perp,
```

or, more operationally,

```text
Phi_O = P_W O|_{K_7}: K_7 -> W_7
```

for a typed operator `O` on `Lambda^4 R^8` with certified rank seven.

The most natural next candidate is the Hodge-star route:

```text
Phi_* = P_W * |_{K_7},
```

but it requires an explicit rank test before it can be promoted.
