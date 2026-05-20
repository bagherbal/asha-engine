# Gate 727 — Conditional Radial-Hopf HistoryLoopUnit Law and Premise-Minimality Audit

## Purpose

Gate 727 follows Gate 726 by closing the radial-Hopf source-type chain for the HistoryLoopUnit candidate:

```text
L = 1/(8*pi)
```

as the expectation value of a radial Hopf payoff observable:

```text
L = Tr(rho_plus R_Hopf)
R_Hopf = (1/(2*pi))P_rad
rho_plus = I_K7+ / 4.
```

This is a bridge-layer closure and premise-minimality audit only. It does not derive the radial selector, twistor selector, scalar runtime lambda, Higgs mass, Yukawa operators, CKM/PMNS, flavor hierarchy, or a native HistoryLoopUnit theorem.

## Registered theorem

```text
pkg/bridge/generation2conditionalradialhopfhistoryloopunitlawandpremiseminimalityaudit
```

```text
generation2conditionalradialhopfhistoryloopunitlawandpremiseminimalityaudit.Generation2ConditionalRadialHopfHistoryLoopUnitLawAndPremiseMinimalityAuditTheorem()
```

## Conditional Radial-Hopf law

Gate 726 supplied the Hopf-fiber phase loop through the radial event:

```text
v_rad(theta)=exp(theta J_H(n))v_rad.
```

Gate 727 defines the payoff observable:

```text
R_Hopf = (1/(2*pi))P_rad.
```

Then:

```text
Tr(rho_plus R_Hopf)
= Tr((I_K7+/4)(1/(2*pi))P_rad)
= (1/4)(1/(2*pi))
= 1/(8*pi).
```

Thus:

```text
A_loop = L - Tr(rho_plus R_Hopf) = 0
```

inside the conditional source-type ledger.

## Premise ladder

The law depends on five typed premises:

1. `dim_R K7+=4`;
2. `rho_plus=I_K7+/4`;
3. a rank-one radial projector `P_rad`;
4. a twistor-selected complex structure `J_H(n)` defining the Hopf fiber through the radial event;
5. first expectation of the phase payoff `Tr(rho_plus[(1/(2*pi))P_rad])`.

Each premise has a distinct structural role.

## Premise-removal audit

Gate 727 verifies:

```text
remove rho_plus  -> radial event weight is not fixed to 1/4
remove P_rad     -> no radial event or Hopf fiber through a radial point
remove n/J_H(n)  -> no selected Hopf phase direction
remove 1/(2*pi)  -> only the event weight 1/4 remains
rank-two event   -> 1/(4*pi), not 1/(8*pi)
full K7+ event   -> 1/(2*pi), not 1/(8*pi)
quadratic moment -> wrong spectral order
```

Therefore every premise is doing real work.

## Non-tautology boundary

The equality is conditionally exact after the premises are supplied. It is not a native theorem because the decisive objects are not derived:

```text
P_rad is not natively selected
n is not natively selected
history transport is not proven to use Hopf phase payoff
rho_plus is a no-bias state choice, not a physical history theorem
```

## Scalar transport placement

Gate 727 preserves the Gate 722/Gate 723 placement:

```text
sealed Higgs socket
-> finite Higgs one-form lane
-> lambda_proxy
-> HistoryLoopUnit transport
```

The radial-Hopf law supplies a source-type candidate for `L` only after the scalar transport lane is active. It does not prove the native runtime transport formula.

## Relation to `7/72`

Gate 727 records the analogy:

```text
7/72      = Tr(rho_72 P_K7)
1/(8*pi) = Tr(rho_plus[(1/(2*pi))P_rad])
```

Both are event-expectation bridge forms, but they belong to different lanes:

```text
7/72      -> boundary/history response
1/(8*pi) -> scalar/runtime HistoryLoop transport
```

Neither derives the other.

## Verdict

```text
PASS_GATE726_RADIAL_PHASE_HOPF_DECOMPOSITION_INHERITED
PASS_RADIAL_HOPF_PAYOFF_OBSERVABLE_DEFINED
PASS_CONDITIONAL_HISTORYLOOP_FUNCTIONAL_DEFINED
PASS_EXPECTATION_REPRODUCES_ONE_OVER_8PI
PASS_PREMISE_LADDER_CONSTRUCTED
PASS_PREMISE_REMOVAL_AUDIT_COMPUTED
PASS_NON_TAUTOLOGY_AUDITED
PASS_SCALAR_TRANSPORT_PLACEMENT_PRESERVED
PASS_EVENT_WEIGHT_ANALOGY_TO_7_OVER_72_AUDITED
CONDITIONAL_SUPPORT_CURRENT_PREMISES_FORM_COMPLETE_CONDITIONAL_HISTORYLOOPUNIT_SOURCE_LAW
CONDITIONAL_SUPPORT_L_IS_RADIAL_HOPF_EXPECTATION_VALUE
CONDITIONAL_SUPPORT_EACH_PREMISE_HAS_NONREDUNDANT_STRUCTURAL_ROLE
FAILED_ROUTE_PREMISES_NOT_NATIVELY_DERIVED
FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR
FAILED_ROUTE_NO_NATIVE_TWISTOR_SELECTOR_N
FAILED_ROUTE_NO_NATIVE_REASON_HISTORY_TRANSPORT_USES_HOPF_PHASE_PAYOFF
FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM
FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_TO_RUNTIME_THEOREM
FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM
FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM
FIREWALL_PRESERVED_GATE727_CONDITIONAL_RADIAL_HOPF_HISTORYLOOP_BOUNDARY
```
