# Gate 636 — K7 Split-Signature Hodge Bilinear Audit

## Purpose

Gate 635 correctly preserved the carrier firewall:

```text
K_7^+ ⊕ K_7^- = 4|3
not=>
W = C^4 with B-L selector 4=1+3.
```

Gate 636 therefore keeps the next calculation native to `K_7`.  It defines the Hodge bilinear form

```text
B_K(x,y) = <x,S_*y>|_{K_7}
```

or, in the Gate634 orthonormal `K_7` frame,

```text
B_K = S_K = Q_K^T S_* Q_K.
```

This is a finite bilinear-geometry audit only.  It does not derive a Fock selector, split-G2 structure, physical spacetime metric, boundary stress, scalar RG matching, Higgs mass, flavor, CKM/PMNS, gauge unification, or a native `7/72` trace theorem.

## Gate635 inheritance

Gate 636 inherits the Gate635 firewall:

```text
K_7 has native Hodge polarity 4|3,
no typed Theta:K_7 <-> W=C^4 is certified,
no native 4=1+3 refinement inside K_7^+ is certified,
tr(S_K)=+1 is a Hodge imbalance, not a distinguished line.
```

Verdict:

```text
PASS_GATE635_HODGE_POLARITY_FIREWALL_INHERITED
```

## B_K definition

Using the Euclidean metric inherited from `Lambda^4 R^8` on the orthonormal `K_7` frame,

```text
g_K(x,y)=<x,y>,
B_K(x,y)=g_K(x,S_K y).
```

Since Gate634 certified `S_K` as symmetric, orthogonal, and involutive, `B_K` is symmetric and nondegenerate.

Verdict:

```text
PASS_BK_HODGE_BILINEAR_DEFINED_ON_K7
PASS_BK_METRIC_CONVERSION_OPERATOR_CERTIFIED
```

## Signature certificate

Gate634 supplied

```text
Spec(S_K) = {+1,+1,+1,+1,-1,-1,-1}.
```

Therefore the inertia of `B_K` is

```text
inertia(B_K) = (4,3,0),
tr(B_K) = +1,
det(B_K) = -1.
```

Because both positive and negative directions exist, the bilinear form has a real null cone.  It is neither positive-definite nor negative-definite.

Verdict:

```text
PASS_BK_SIGNATURE_4_3_CERTIFIED
CONDITIONAL_SUPPORT_K7_CARRIES_NATIVE_SPLIT_SIGNATURE_STRUCTURE
```

## Orthogonality of the Hodge sectors

The Gate634 projectors

```text
P_K^+ = (I+S_K)/2,
P_K^- = (I-S_K)/2
```

split

```text
K_7 = K_7^+ ⊕ K_7^-.
```

Since `S_K` is symmetric, the two eigenspaces are orthogonal under `g_K`.  Under `B_K`, the plus sector carries `+g_K`, the minus sector carries `-g_K`, and cross terms vanish.

Verdict:

```text
PASS_K7_PLUS_MINUS_ORTHOGONALITY_CERTIFIED_FOR_GK_AND_BK
CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_IS_BILINEAR_NOT_SELECTOR
```

## Split-octonionic compatibility audit

The signature `(4,3)` is the correct bilinear signature type for a split-signature seven-carrier candidate.  However, a split-G2 or split-octonionic structure requires more than a bilinear form.  It requires a compatible native object such as

```text
Omega_K,
```

a stable 3-form, cross product, or calibration on `K_7` compatible with `B_K`.

Gate 636 certifies `B_K` only.  It does not certify `Omega_K`, a cross product, a calibration, or a split-G2 preservation theorem.

Verdict:

```text
FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE_YET
FAILED_ROUTE_NO_NATIVE_OMEGA_K_THREE_FORM_CERTIFIED
```

## Stabilizer audit

The bilinear form suggests the abstract stabilizer lane

```text
O(4,3),
SO(4,3) after orientation restriction.
```

A split-G2 subgroup would require the missing compatible `Omega_K`:

```text
G2_split ⊂ SO(4,3).
```

Gate 636 therefore records the stabilizer candidate without promoting a split-G2 theorem.

Verdict:

```text
PASS_SPLIT_SIGNATURE_STABILIZER_CANDIDATE_AUDITED
FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE_YET
```

## Firewalls

Gate 636 explicitly blocks the following jumps:

```text
(4,3) Hodge bilinear ≠ physical spacetime metric,
(4,3) Hodge bilinear ≠ Fock 1+3 selector,
(4,3) Hodge bilinear ≠ split-G2 without Omega_K,
(4,3) Hodge bilinear ≠ boundary stress,
(4,3) Hodge bilinear ≠ native 7/72 trace theorem.
```

The next missing object is therefore not the Fock selector map yet, and not a boundary map.  It is:

```text
Omega_K:
compatible native 3-form / cross product / calibration on (K_7,B_K).
```

## Final verdict

```text
PASS_GATE635_HODGE_POLARITY_FIREWALL_INHERITED
PASS_BK_HODGE_BILINEAR_DEFINED_ON_K7
PASS_BK_SIGNATURE_4_3_CERTIFIED
PASS_BK_METRIC_CONVERSION_OPERATOR_CERTIFIED
PASS_K7_PLUS_MINUS_ORTHOGONALITY_CERTIFIED_FOR_GK_AND_BK
CONDITIONAL_SUPPORT_K7_CARRIES_NATIVE_SPLIT_SIGNATURE_STRUCTURE
CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_IS_BILINEAR_NOT_SELECTOR
PASS_SPLIT_SIGNATURE_STABILIZER_CANDIDATE_AUDITED
FAILED_ROUTE_NO_TYPED_K7_TO_FOCK_SELECTOR_MAP
FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE_YET
FAILED_ROUTE_NO_NATIVE_OMEGA_K_THREE_FORM_CERTIFIED
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_K7_SPLIT_SIGNATURE_NOT_PHYSICAL_SPACETIME_METRIC
FIREWALL_PRESERVED_GATE636_SPLIT_SIGNATURE_IS_NATIVE_NOT_PHYSICAL
```
