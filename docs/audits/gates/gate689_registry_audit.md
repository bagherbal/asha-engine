# Gate 689 — First-Trace Functional Selection and Spectral-Order Audit

## Registry target

```text
pkg/bridge/generation2firsttracefunctionalselectionandspectralorderaudit
```

Registered theorem:

```text
generation2firsttracefunctionalselectionandspectralorderaudit.Generation2FirstTraceFunctionalSelectionAndSpectralOrderAuditTheorem()
```

## Purpose

Gate 688 audited the support-selected response operator

```text
R_split = S_split P_K7
```

and recorded the trace-power cable

```text
Tr(R_split^n)=7 S_split^n, n>=1.
```

Gate 689 audits which scalar functional on `R_split` matches the active history defect coordinate

```text
D_base = kappa_lambda+kappa_e+lambda(Lambda_12).
```

This is a bridge-layer spectral-functional audit only.  It does not derive boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS, a native first-trace theorem, or a native `7/72` theorem.

## Inherited typed objects

```text
R_split = S_split P_K7
S_split = lambda(Lambda_12)+(R_3-1)
P_K7 = Boolean-octonionic intersection projector
H_72 = Lambda^4 R^8 ⊕ R^2_boundary
rank(P_K7)=7
dim(H_72)=72
D_base = kappa_lambda+kappa_e+lambda(Lambda_12)
```

Numerical bridge coordinates inherited from the current ledger:

```text
S_split ≈ 0.0012924448188162962
D_base  ≈ 0.0001256552099683575
```

## Candidate spectral functionals

Gate 689 audits six scalarizations of the same support-selected response operator.

| Candidate | Functional | Value | Classification |
|---|---:|---:|---|
| `F_1` | `Tr(R_split)/72` | `0.0001256543573849177` | active first ordinary trace |
| `F_2` | `Tr(R_split^2)/72` | `1.624013231638281e-7` | quadratic, inactive |
| `F_3` | `Tr(R_split^3)/72` | `2.098966032736838e-10` | cubic, inactive |
| `F_Frob` | `||R_split||_F^2/72` | `1.624013231638281e-7` | same as quadratic trace, inactive |
| `F_signed` | `Tr((P_+-P_-)R_split)/72` | `1.795062248355967e-5` | Hodge-signed polarity trace, inactive |
| `F_full` | `Tr(S_split I_H72)/72` | `0.0012924448188162962` | full identity trace, inactive |

The first trace residual against the active wall-defect coordinate is

```text
D_base - F_1 ≈ 8.525834398014336e-10.
```

The higher-power candidates are wrong by spectral order: `F_2` and `F_Frob` are order `S_split^2`, and `F_3` is order `S_split^3`.  The Hodge-signed trace is linear, but it uses the internal polarity imbalance `4-3=1` rather than the total selected support `4+3=7`.  The full identity trace is also linear, but it ignores the defect projector support.

## Linear order audit

The active defect coordinate is linear in wall-distance coordinates:

```text
D_base = kappa_lambda+kappa_e+lambda(Lambda_12).
```

The boundary split scalar is also linear:

```text
S_split = lambda(Lambda_12)+(R_3-1).
```

Therefore the matching scalar functional must be first order in `S_split`.  Quadratic trace powers and Frobenius norm are second-order response candidates, not the active bridge.

## Trace-type audit

On the Hodge-polarized contact carrier:

```text
K_7 = K_7^+ ⊕ K_7^-
dim K_7^+ = 4
dim K_7^- = 3
```

The ordinary total-support trace sees

```text
4+3 = 7,
```

while the Hodge-signed trace sees

```text
4-3 = 1.
```

The active bridge uses total support rank, not the signed polarity imbalance:

```text
Tr_H72(R_split)/72 = (7/72)S_split,
Tr_H72((P_+-P_-)R_split)/72 = (1/72)S_split.
```

## Verdict interpretation

Gate 689 conditionally supports:

```text
active bridge = first ordinary trace of the support-selected response operator.
```

It does not prove a native principle forcing that choice.  The sharpened missing theorem is:

```text
HistoryResponseFirstTraceTheorem.
```

A future theorem would need to explain why physical history uses the first ordinary total-support trace rather than higher spectral powers, Frobenius norm, signed polarity trace, or full identity trace.

## Expected status lines

```text
PASS_GATE688_RESPONSE_OPERATOR_SPECTRUM_INHERITED
PASS_SPECTRAL_FUNCTIONAL_CANDIDATES_DEFINED
PASS_FIRST_TRACE_RESPONSE_COMPUTED
PASS_HIGHER_TRACE_RESPONSES_COMPUTED
PASS_HODGE_SIGNED_TRACE_COMPUTED
PASS_RESIDUAL_COMPARISON_AUDITED
PASS_LINEAR_ORDER_MATCH_AUDITED
CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_SELECTS_FIRST_ORDER_ORDINARY_TRACE
CONDITIONAL_SUPPORT_DBASE_IS_LINEAR_WALL_RESPONSE_COORDINATE
FAILED_ROUTE_QUADRATIC_TRACE_OR_FROBENIUS_NORM_NOT_ACTIVE
FAILED_ROUTE_HODGE_SIGNED_TRACE_NOT_ACTIVE
FAILED_ROUTE_FULL_IDENTITY_TRACE_NOT_ACTIVE
FAILED_ROUTE_NO_NATIVE_FIRST_TRACE_RESPONSE_THEOREM
FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM
FIREWALL_PRESERVED_GATE689_FIRST_TRACE_SELECTION_BOUNDARY
```

## Validation command

Focused validation command used for this gate:

```text
go test -p=1 ./pkg/bridge/generation2firsttracefunctionalselectionandspectralorderaudit -count=1
```

`internal/app` was not tested directly.
