# Gate 634 — K7 Hodge-Signature Stabilizer Audit

## Purpose

Gate 633 proved the native stabilizer result:

```text
*K_7 = K_7.
```

Gate 634 therefore stops treating Hodge star as a leakage map and restricts it to the contact carrier itself.  The central object is the internal endomorphism

```text
S_K = Q_K^T S_* Q_K : K_7 -> K_7,
```

where `S_*` is the Hodge-star matrix on `Lambda^4 R^8` and `Q_K` is an orthonormal basis for the seven-dimensional contact carrier.  This gate asks whether `K_7` is fully self-dual, fully anti-self-dual, or mixed under the restricted Hodge action.

This is a native finite-linear-algebra audit only.  It does not derive boundary stress, `7/72`, scalar RG matching, Higgs mass, flavor, CKM/PMNS, physical Lorentzian orientation, or gauge unification.

## Inherited data

```text
H = Lambda^4 R^8, dim H = 70
K_7 = U ∩ V, dim K_7 = 7
S_*^2 = I on Lambda^4 R^8
*K_7 = K_7
```

Gate 633 blocks the earlier routes:

```text
*K_7 != W_7,
*K_7 != V_0 = V⊖K_7,
*K_7 != U_0 = U⊖K_7.
```

## Restricted Hodge operator

Gate 634 computes:

```text
S_K = Q_K^T S_* Q_K.
```

The certificate is:

```text
S_K^T S_K ≈ I_7,
S_K^2 ≈ I_7,
S_K^T ≈ S_K,
tr(S_K) ≈ +1,
det(S_K) ≈ -1.
```

Verdict:

```text
PASS_RESTRICTED_HODGE_OPERATOR_SK_DEFINED
PASS_SK_ORTHOGONAL_SYMMETRIC_INVOLUTIVE
```

## Hodge spectrum and signature

The computed spectrum is:

```text
Spec(S_K) = {+1,+1,+1,+1,-1,-1,-1}.
```

Therefore:

```text
dim K_7^+ = rank((I+S_K)/2) = 4,
dim K_7^- = rank((I-S_K)/2) = 3,
K_7 = K_7^+ ⊕ K_7^-.
```

So `K_7` is neither fully self-dual nor fully anti-self-dual.  It is Hodge-stable with a mixed native polarity:

```text
(n_+, n_-) = (4,3).
```

Verdict:

```text
PASS_K7_HODGE_SPECTRUM_COMPUTED
PASS_K7_HAS_MIXED_HODGE_SIGNATURE_4_PLUS_3_MINUS
FAILED_ROUTE_K7_NOT_FULLY_SELF_DUAL
FAILED_ROUTE_K7_NOT_FULLY_ANTI_SELF_DUAL
```

## Ambient self-dual / anti-self-dual projection audit

Using the ambient projectors

```text
P_+ = (I+S_*)/2,
P_- = (I-S_*)/2,
```

Gate 634 confirms the expected ambient split:

```text
Lambda^4 R^8 = Lambda^4_+ ⊕ Lambda^4_-,
dim Lambda^4_+ = 35,
dim Lambda^4_- = 35.
```

Restricted to `K_7`, the projection weights are:

```text
||P_+ Q_K||_F^2 = 4,
||P_- Q_K||_F^2 = 3,
||P_+ Q_K||_F^2 / 7 = 4/7,
||P_- Q_K||_F^2 / 7 = 3/7.
```

Verdict:

```text
PASS_AMBIENT_SELF_ANTI_SELF_DUAL_PROJECTIONS_COMPUTED
PASS_K7_SELF_ANTI_SELF_DUAL_PROJECTORS_CERTIFIED
```

## Consequence for previous routes

Gate 634 strengthens Gate 633 by showing that the contact carrier is not merely Hodge-stable; it carries an internal Hodge polarity.  But it does not reopen the failed maps:

```text
K_7 -> W_7,
K_7 -> V_0,
K_7 -> R^2_boundary.
```

The new native object is:

```text
S_*|_{K_7} with signature (4,3).
```

The missing object remains a typed bridge from this polarity, or from a future defect trace, into the boundary-stress pair.

## Final verdict

```text
PASS_GATE633_HODGE_STABILITY_INHERITED
PASS_RESTRICTED_HODGE_OPERATOR_SK_DEFINED
PASS_SK_ORTHOGONAL_SYMMETRIC_INVOLUTIVE
PASS_K7_HODGE_SPECTRUM_COMPUTED
PASS_K7_HAS_MIXED_HODGE_SIGNATURE_4_PLUS_3_MINUS
PASS_AMBIENT_SELF_ANTI_SELF_DUAL_PROJECTIONS_COMPUTED
PASS_K7_SELF_ANTI_SELF_DUAL_PROJECTORS_CERTIFIED
FAILED_ROUTE_K7_NOT_FULLY_SELF_DUAL
FAILED_ROUTE_K7_NOT_FULLY_ANTI_SELF_DUAL
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FIREWALL_PRESERVED_GATE634_K7_HODGE_SIGNATURE_BOUNDARY
```
