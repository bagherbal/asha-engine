# Gate 635 — K7 Hodge Polarity and Projective Selector Alignment Audit

## Purpose

Gate 634 proved the native Hodge polarity of the contact carrier:

```text
K_7 = K_7^+ ⊕ K_7^-,
dim K_7^+ = 4,
dim K_7^- = 3,
tr(S_K) = +1,
det(S_K) = -1.
```

Gate 635 asks whether this native `4|3` Hodge split has any typed relation to the previously certified Witt/Fock projective selector lane:

```text
W = C^4,
B-L = diag(-1, 1/3, 1/3, 1/3),
4 = 1 + 3,
CP^0 | CP^2.
```

This is a carrier-comparison audit only.  It does not derive flavor, CKM/PMNS, scalar RG matching, Higgs mass, gauge unification, boundary stress, a native `7/72` trace theorem, or a physical orientation theorem.

## Inherited K7 polarity

Gate 635 inherits the Gate634 restricted Hodge operator:

```text
S_K = Q_K^T S_* Q_K : K_7 -> K_7.
```

The certified split is:

```text
K_7^+ = im((I+S_K)/2), dim K_7^+ = 4,
K_7^- = im((I-S_K)/2), dim K_7^- = 3.
```

Verdict:

```text
PASS_GATE634_K7_HODGE_SIGNATURE_INHERITED
PASS_K7_PLUS_MINUS_SUBSPACES_DEFINED
```

## Projective selector reference

Gate 572 supplies the comparison carrier:

```text
W = C^4,
CP^3 = P(C^4),
B-L = diag(-1, 1/3, 1/3, 1/3).
```

Its selector geometry gives:

```text
CP^0 | CP^2,
1 + 3 eigenspace split before projectivization,
stabilizer U(1) × U(3),
Lie algebra u(1)+u(3).
```

Verdict:

```text
PASS_PROJECTIVE_SELECTOR_1_PLUS_3_INHERITED
```

## 4|3 versus 4=1+3 alignment audit

The dimension pattern is genuinely suggestive:

```text
K_7 Hodge polarity:        4 | 3
Fock selector reference:   4 = 1 + 3
```

But the carriers are different:

```text
K_7 ⊂ Lambda^4 R^8,
W = C^4,
CP^3 = P(C^4).
```

Gate 572 already preserved the firewall that no CP3-to-K7 functor is certified.  Gate 635 therefore records the resemblance only as a candidate alignment:

```text
CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_RESEMBLES_PROJECTIVE_SELECTOR_SPLIT
FAILED_ROUTE_NO_TYPED_K7_TO_FOCK_SELECTOR_MAP_YET
```

## K7+ refinement audit

The four-dimensional self-dual sector is real and native:

```text
K_7^+, dim K_7^+ = 4.
```

But Hodge polarity alone acts as the identity on `K_7^+`.  It supplies a four-plane, not an internal rank-one line plus a complementary three-plane.  Therefore the selector refinement

```text
4 = 1 + 3
```

is not recovered inside `K_7^+` without an additional typed selector.

Verdict:

```text
FAILED_ROUTE_NO_NATIVE_4_EQUALS_1_PLUS_3_REFINEMENT_INSIDE_K7_PLUS
```

## K7− triplet audit

The anti-self-dual sector has dimension three:

```text
dim K_7^- = 3.
```

This matches the spatial block dimension in the `B-L` selector lane, but no typed map identifies `K_7^-` with the Fock/projective `CP^2` spatial eigenspace.

Verdict:

```text
CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_RESEMBLES_PROJECTIVE_SELECTOR_SPLIT
FAILED_ROUTE_NO_TYPED_K7_TO_FOCK_SELECTOR_MAP_YET
```

## Trace imbalance audit

Gate 634 found:

```text
tr(S_K) = 4 - 3 = +1.
```

Gate 635 classifies this as a signed Hodge imbalance.  It is not a rank-one projector and does not define a distinguished line inside `K_7^+`.

Verdict:

```text
FAILED_ROUTE_TRACE_PLUS_ONE_IS_HODGE_IMBALANCE_NOT_DISTINGUISHED_LINE
```

## Missing carrier map

The sharpened missing object is:

```text
Theta:
K_7^+ ⊕ K_7^-
<->
W = C^4 with B-L 1+3 selector data.
```

or a proof that no such typed carrier comparison map exists.

At Gate 635, no such `Theta` is certified.

## Boundary and 7/72 readiness

Even a future `Theta` map would not by itself assign boundary stress.  A separate lawful map into the boundary-stress pair would still be required:

```text
K_7 polarity / selector comparison
-> R^2_boundary.
```

Therefore Gate 635 does not promote:

```text
boundary stress,
7/72,
scalar RG matching,
flavor,
CKM/PMNS,
gauge unification,
physical orientation.
```

## Final verdict

```text
PASS_GATE634_K7_HODGE_SIGNATURE_INHERITED
PASS_K7_PLUS_MINUS_SUBSPACES_DEFINED
PASS_PROJECTIVE_SELECTOR_1_PLUS_3_INHERITED
PASS_4_PLUS_3_POLARITY_AUDITED
CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_RESEMBLES_PROJECTIVE_SELECTOR_SPLIT
FAILED_ROUTE_NO_TYPED_K7_TO_FOCK_SELECTOR_MAP_YET
FAILED_ROUTE_NO_NATIVE_4_EQUALS_1_PLUS_3_REFINEMENT_INSIDE_K7_PLUS
FAILED_ROUTE_TRACE_PLUS_ONE_IS_HODGE_IMBALANCE_NOT_DISTINGUISHED_LINE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FIREWALL_PRESERVED_GATE635_HODGE_POLARITY_SELECTOR_BOUNDARY
```
