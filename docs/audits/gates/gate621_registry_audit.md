# Gate 621 — Scalar Tree-Proxy to Runtime Matching Gap Audit

## Purpose

Gate 620 separated the positive spectral/tree scalar proxy lane from the RG-transported runtime Standard Model quartic lane:

```text
lambda_proxy = (3/8)(b/a^2)
lambda_runtime(mu) = canonical SM quartic transported by the v1 RG ledger
```

Gate 621 audits the low-scale matching gap between these two scalar lanes at `M_Z`.  It does not derive a Higgs mass, scalar stability, a lambda-zero boundary, gauge unification, or a native scalar matching theorem.

## Inherited values

```text
b/a^2(M_Z) = 0.33307493962706697
lambda_proxy(M_Z) = 0.12490310236015
lambda_runtime(M_Z) = 0.1296525650504758

b/a^2(Lambda_12) = 0.3330764110541872
lambda_proxy(Lambda_12) = 0.12490365414532
lambda_runtime(Lambda_12) = -0.049700942077683274

v = 246.21965079413738 GeV
```

## Matching gap

```text
Delta lambda_match
= lambda_runtime(M_Z) - lambda_proxy(M_Z)
= 0.0047494626903257
```

Relative gap:

```text
Delta lambda_match / lambda_proxy = 0.0380251779225709
Delta lambda_match / lambda_runtime = 0.0366322308276489
```

The correction has positive sign: `lambda_proxy(M_Z)` is below `lambda_runtime(M_Z)`.

## Effective c_lambda diagnostic

```text
c_proxy = 3/8 = 0.375
c_needed(M_Z) = lambda_runtime(M_Z)/(b/a^2)
              = 0.389259441720964

Delta c = c_needed - c_proxy
        = 0.0142594417209637

Delta c / c_proxy = 0.0380251779225699
```

This is a bridge diagnostic only.  It does not certify a new scalar convention coefficient.

## Higgs proxy gap diagnostic

Using the same runtime `v` ledger:

```text
m_H_proxy = sqrt(2 lambda_proxy) v
          = 123.062099940214 GeV

m_H_runtime = sqrt(2 lambda_runtime(M_Z)) v
            ≈ 125.380000000000 GeV

Delta m_H = m_H_runtime - m_H_proxy
          ≈ 2.317900059786 GeV
```

This is only a tree-level diagnostic translation through the runtime `v` ledger.  It is not a Higgs pole-mass theorem.

## Typed source candidates

| candidate | sign-compatible? | native? | status |
|---|---:|---:|---|
| pole/MSbar matching | yes | no | bridge slot only |
| one-loop scalar threshold correction | yes | no | bridge slot only |
| top-loop correction | yes | no | bridge slot only |
| gauge-loop correction | yes | no | bridge slot only |
| two-loop RG improvement | yes | no | bridge slot only |
| scalar field normalization convention | yes | no | `c_lambda` still convention-sealed |
| missing neutrino/Yukawa trace contribution | yes | no | no values inserted |
| spectral-action convention coefficient `c_lambda` | yes | no | no certified value |

## Runtime transport chain

Gate 621 records the scalar chain:

```text
lambda_proxy
+
Delta lambda_match
->
lambda_runtime(M_Z)
->
RG transport
->
lambda_runtime(Lambda_12)
```

This separates the positive low-scale proxy from the negative high-scale runtime quartic.

## Stress-seal impact

The matching-gap audit improves the scalar lane architecture, but does not change the Gate 613 stress seal:

```text
S_boundary = (R_3 - 1, lambda_runtime(Lambda_12)) ≈ (+xi_boundary, -xi_boundary)
```

The stress seal cannot replace `lambda_runtime(Lambda_12)` with `lambda_proxy(Lambda_12)`.

## Verdict

```text
PASS_GATE620_PROXY_RUNTIME_SEPARATION_INHERITED
PASS_LOW_SCALE_MATCHING_GAP_COMPUTED
PASS_EFFECTIVE_C_LAMBDA_CORRECTION_COMPUTED
PASS_HIGGS_PROXY_GAP_DIAGNOSTIC_COMPUTED
CONDITIONAL_SUPPORT_LAMBDA_PROXY_CLOSE_TO_RUNTIME_LAMBDA_AT_MZ
CONDITIONAL_SUPPORT_POSITIVE_MATCHING_CORRECTION_REQUIRED
CONDITIONAL_SUPPORT_PROXY_TO_RUNTIME_CHAIN_DEFINED
FAILED_ROUTE_NO_NATIVE_MATCHING_CORRECTION_THEOREM
FAILED_ROUTE_NO_NATIVE_C_LAMBDA_THREE_EIGHTHS_THEOREM
FAILED_ROUTE_NO_NATIVE_HIGGS_MASS_OR_POLE_THEOREM
FAILED_ROUTE_NO_NATIVE_PROXY_TO_RUNTIME_MATCHING_THEOREM
FAILED_ROUTE_NO_NATIVE_NEUTRINO_TRACE_COMPLETION_THEOREM
FIREWALL_PRESERVED_GATE621_SCALAR_PROXY_MATCHING_BOUNDARY
```
