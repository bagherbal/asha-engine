# Gate 620 — b/a² One-Third Rigidity and Spectral Quartic Proxy Audit

## Purpose

Gate 620 follows Gate 619 by separating the positive spectral/tree scalar proxy lane from the RG-transported runtime quartic lane.  Gate 619 showed that the direct identification

```text
lambda_runtime(Lambda_12) = c_lambda * b/a^2
```

fails by sign at `Lambda_12`, because `b/a^2 >= 0` while the v1 transported runtime quartic is negative.  Gate 620 audits the separate diagnostic clue that `b/a^2` is nearly frozen near `1/3`, and that the typed candidate

```text
lambda_proxy = (3/8)(b/a^2)
```

is a positive spectral/tree scalar quartic proxy close to the low-scale runtime quartic.

This is a symbolic/diagnostic audit only.  It does not derive Higgs mass, Higgs stability, a lambda-zero boundary, or a native scalar quartic theorem.

## Inherited runtime data

```text
M_Z:
  a = 2.8424095142339083
  b = 2.6910096440382287
  b/a^2 = 0.33307493962706697
  lambda_runtime(M_Z) = 0.1296525650504758

Lambda_12:
  a = 0.6941198223775996
  b = 0.16047699018700937
  b/a^2 = 0.3330764110541872
  lambda_runtime(Lambda_12) = -0.049700942077683274
```

## b/a² rigidity

Gate 620 compares `b/a^2` to `1/3`:

```text
M_Z:
  b/a^2 - 1/3 = -0.000258393706266348

Lambda_12:
  b/a^2 - 1/3 = -0.000256922279146099

Drift:
  b/a^2(Lambda_12) - b/a^2(M_Z) = 0.000001471427120230
```

The ratio is therefore nearly locked to `1/3` under the v1 visible Yukawa transport.  This is not a native theorem.

## Top/color dominance explanation

Using the trace definitions:

```text
a = Tr(Y_e†Y_e + Y_nu†Y_nu + 3Y_u†Y_u + 3Y_d†Y_d)
b = Tr((Y_e†Y_e)^2 + (Y_nu†Y_nu)^2 + 3(Y_u†Y_u)^2 + 3(Y_d†Y_d)^2)
```

observed top dominance gives the conditional approximation:

```text
a ≈ 3 y_t^2
b ≈ 3 y_t^4
b/a^2 ≈ 1/3
```

This explains the rigidity as an observed hierarchy effect, not as ASHA-native flavor derivation.

## c_lambda = 3/8 proxy

Gate 620 tests the typed candidate:

```text
c_lambda = sin^2(theta_*) = 3/8
lambda_proxy = (3/8)(b/a^2)
```

The results are:

```text
M_Z:
  lambda_proxy = 0.12490310236015
  lambda_runtime = 0.1296525650504758
  lambda_proxy - lambda_runtime = -0.0047494626903257

Lambda_12:
  lambda_proxy = 0.12490365414532
  lambda_runtime = -0.049700942077683274
  lambda_proxy - lambda_runtime = 0.174604596223003
```

Thus the proxy is close to the low-scale runtime quartic but cannot equal the negative high-scale runtime quartic.

## Low-scale Higgs proxy diagnostic

Using the runtime `v` ledger:

```text
v = 246.21965079413738 GeV
m_H_proxy = sqrt(2 lambda_proxy) v ≈ 123.079640045076 GeV
m_H_runtime = sqrt(2 lambda_runtime(M_Z)) v ≈ 125.38 GeV
```

This is a diagnostic of the positive spectral/tree proxy only.  It is not a Higgs mass derivation or pole-mass theorem.

## Runtime transport separation

Gate 620 separates two scalar lanes:

```text
spectral/tree scalar lane:
  lambda_proxy = (3/8)(b/a^2) > 0

runtime RG scalar lane:
  lambda_runtime(mu), transported by the v1 continuum RG ledger
```

Therefore the Gate 613 stress seal remains:

```text
S_boundary = (R_3 - 1, lambda_runtime(Lambda_12)) ≈ (+xi_boundary, -xi_boundary)
```

It cannot use `lambda_proxy(Lambda_12)` in place of the negative runtime quartic.

## Verdict

```text
PASS_GATE619_BA2_DIAGNOSTIC_INHERITED
PASS_B_OVER_A_SQUARED_NEAR_ONE_THIRD_AT_MZ_AND_LAMBDA12
CONDITIONAL_SUPPORT_BA2_ONE_THIRD_FROM_TOP_COLOR_DOMINANCE
PASS_C_LAMBDA_THREE_EIGHTHS_PROXY_COMPUTED
CONDITIONAL_SUPPORT_LAMBDA_PROXY_CLOSE_TO_RUNTIME_LAMBDA_AT_MZ
FAILED_ROUTE_LAMBDA_PROXY_DOES_NOT_EQUAL_NEGATIVE_RUNTIME_LAMBDA_AT_LAMBDA12
CONDITIONAL_SUPPORT_SPECTRAL_TREE_QUARTIC_AND_RUNTIME_RG_QUARTIC_MUST_BE_SEPARATED
FAILED_ROUTE_NO_NATIVE_BA2_ONE_THIRD_THEOREM
FAILED_ROUTE_NO_NATIVE_C_LAMBDA_THREE_EIGHTHS_SCALAR_THEOREM
FAILED_ROUTE_NO_NATIVE_PROXY_TO_RUNTIME_MATCHING_THEOREM
FIREWALL_PRESERVED_GATE620_BA2_QUARTIC_PROXY_BOUNDARY
```
