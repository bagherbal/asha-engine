# Gate 728 — Dual Event-Expectation Scalar Runtime Transport Assembly Audit

## Purpose

Gate 728 follows Gate 700 and Gate 727 by assembling two conditional event-expectation bridge laws into the active scalar runtime transport form.

Gate 700 closed the boundary/history response:

```text
D_base ≈ Tr(rho_72 sigma_boundary P_K7)
       = (7/72)S_split.
```

Gate 727 closed the Radial-Hopf source law for the HistoryLoopUnit:

```text
L = Tr[rho_plus (1/(2*pi))P_rad]
  = 1/(8*pi).
```

Gate 728 audits whether these two bridge laws coherently assemble into:

```text
lambda_runtime ≈ lambda_proxy[1+L(1-W_72+kappa_e)].
```

This is a bridge-layer scalar-runtime assembly and residual-propagation audit only. It does not derive scalar runtime lambda, Higgs mass, radial selector, twistor selector, Yukawa operators, CKM/PMNS, flavor hierarchy, or a native HistoryLoopUnit theorem.

## Registered theorem

```text
pkg/bridge/generation2dualeventexpectationscalarruntimetransportassemblyaudit
```

```text
generation2dualeventexpectationscalarruntimetransportassemblyaudit.Generation2DualEventExpectationScalarRuntimeTransportAssemblyAuditTheorem()
```

## Boundary/history substitution

Gate 700 gives:

```text
D_base = kappa_lambda+kappa_e+lambda
D_base ≈ p_K7 S_split
p_K7 = 7/72.
```

Therefore:

```text
kappa_lambda ≈ p_K7 S_split - kappa_e - lambda.
```

Equivalently:

```text
kappa_lambda ≈ W_72-kappa_e
```

with:

```text
W_72 = |lambda| + p_K7 S_split
     = (65/72)|lambda| + (7/72)(R_3-1).
```

The exact relation retains the inherited wall residual:

```text
kappa_lambda = W_72-kappa_e+E_wall.
```

## Radial-Hopf substitution

Gate 727 gives:

```text
L = Tr[rho_plus (1/(2*pi))P_rad].
```

Substituting this into the scalar transport lane gives:

```text
lambda_runtime
≈ lambda_proxy[
  1 + Tr(rho_plus R_Hopf)(1-W_72+kappa_e)
].
```

## Dual event-expectation form

Using the positive-distance boundary wound observable:

```text
W_boundary = |lambda|I_H72 + S_split P_K7,
```

with:

```text
Tr(rho_72 W_boundary)=W_72,
```

Gate 728 assembles the conditional bridge form:

```text
lambda_runtime
≈
lambda_proxy{
  1+
  Tr[rho_plus (1/(2*pi))P_rad]
  [1-Tr(rho_72 W_boundary)+kappa_e]
}.
```

This is the current scalar runtime bridge as a dual event-expectation form:

```text
K7 boundary/history event expectation
+
Radial-Hopf scalar loop event expectation.
```

## Residual propagation

The wall residual is:

```text
E_wall = D_base - (7/72)S_split ≈ 8.525834398e-10.
```

Using `W_72-kappa_e` instead of the exact `kappa_lambda` shifts the scalar runtime prediction by:

```text
Delta_lambda_pred = lambda_proxy * L * E_wall.
```

With inherited values:

```text
lambda_proxy ≈ 0.12490310236015
L            ≈ 0.0397887357729738
E_wall       ≈ 8.525834398e-10
```

this gives:

```text
Delta_lambda_pred ≈ 4.237e-12.
```

So the tiny scalar-runtime residual is typed as propagated wall-balance residual, not as a new independent theorem.

## Noncircularity boundary

Gate 728 preserves the fact that `kappa_lambda` was originally defined from `lambda_runtime`, `lambda_proxy`, and `L`. Therefore the assembled runtime form is a bridge consistency closure, not an independent prediction.

## Seal dependence

The assembled formula depends on sealed or conditional objects:

```text
n:        twistor selector for J_H
P_rad:    radial projector / scalar vacuum direction
rho_plus: no-bias state on K7+
rho_72:   no-bias state on H72
P_K7:     Boolean-octonionic support-selected event
kappa_e:  flavor wall deficit / OrientationBalance input
```

These premises remain non-native at this gate.

## Verdict

```text
PASS_GATE700_BOUNDARY_HISTORY_RESPONSE_INHERITED
PASS_GATE727_RADIAL_HOPF_HISTORYLOOP_INHERITED
PASS_BOUNDARY_HISTORY_RESPONSE_SUBSTITUTED_INTO_KAPPA_LAMBDA
PASS_RADIAL_HOPF_L_SUBSTITUTED_INTO_SCALAR_TRANSPORT
PASS_DUAL_EVENT_EXPECTATION_FORM_ASSEMBLED
PASS_WALL_RESIDUAL_PROPAGATION_COMPUTED
PASS_NONCIRCULARITY_AUDITED
PASS_SEAL_DEPENDENCE_AUDITED
PASS_FIREWALLS_ENFORCED
CONDITIONAL_SUPPORT_SCALAR_RUNTIME_BRIDGE_AS_DUAL_EVENT_EXPECTATION_FORM
CONDITIONAL_SUPPORT_RUNTIME_RESIDUAL_IS_PROPAGATED_HISTORY_WALL_RESIDUAL
CONDITIONAL_SUPPORT_RADIAL_HOPF_LOOP_UNIT_AND_K7_BOUNDARY_RESPONSE_COMBINE_IN_SCALAR_TRANSPORT
FAILED_ROUTE_ASSEMBLED_RUNTIME_FORM_NOT_INDEPENDENT_PREDICTION
FAILED_ROUTE_PREMISES_NOT_NATIVELY_DERIVED
FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_TO_RUNTIME_THEOREM
FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM
FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR
FAILED_ROUTE_NO_NATIVE_BOUNDARY_HISTORY_RESPONSE_THEOREM
FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM
FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM
FIREWALL_PRESERVED_GATE728_DUAL_EVENT_EXPECTATION_RUNTIME_BOUNDARY
```
