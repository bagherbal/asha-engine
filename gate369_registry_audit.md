# Gate 369 Registry Audit — Eta-Graded Left-Right Trace / Noncentral Hamiltonian Extraction Sieve

## Gate identity

- **Gate:** 369
- **Package:** `pkg/bridge/etagradedlrtrace`
- **Theorem:** `EtaGradedLeftRightTraceNoncentralHamiltonianExtractionSieveTheorem`
- **Audit ID:** `GATE369-ETA-GRADED-LEFT-RIGHT-TRACE-NONCENTRAL-HAMILTONIAN-EXTRACTION-SIEVE`
- **Layer:** Bridge / Phase-III Flow Extension
- **Purpose:** execute the exact eta-graded trace target isolated by Gate 368 and determine whether `tau_eta` emerges natively as a modular Hamiltonian.

## Files, folders, and active gate chain

| Region | Project objects | Gate-369 relevance |
|---|---|---|
| Core docs | `README.md`, `docs/architecture.md`, `GateResearcherMethod.md` | Ledger and method constraints. |
| Registry | `internal/app/app.go` | Gate 369 is registered after Gate 368 and before runtime cache. |
| Current package | `pkg/bridge/etagradedlrtrace` | Executes eta-graded trace sieve. |
| Gate 368 package | `pkg/bridge/bimodulemodularcurvature` | Supplies the target equation and circularity firewall. |
| Flow chain | Gates 362–369 | Static closure → modular flow → nontracial state → KMS → Hamiltonian origin → Lorentzian-time no-go → bimodule curvature → eta trace. |
| Bimodule/heavy sector | Gates 290–320, 347 | Morita trace capacity, doubled space, Majorana/heavy overlap, and flavor-invariance obstructions. |

## Inherited Gate-368 target

```text
executed=true
target=Pi_gen Tr_support^eta(C_LR) = a I_3 + b tau_eta, b != 0
circularity_firewall=true no_empirical_flavor=true
gate368_truth=Gate 368 audits whether finite Left-Right Morita bimodule curvature supplies the missing internal thermal-time Hamiltonian after Gate 367 ruled out ordinary Lorentzian time. Pure B-gap, pure Omega_Hsigma support, and ungraded Left-Right curvature project centrally on generation space. An eta/tau_eta-weighted lane has the correct noncentral KMS capacity, but the present bimodule audit does not derive tau_eta from the Left-Right contraction; inserting it would be circular. Therefore internal thermal time is not yet derived, and the 15 vacuum coordinates remain quarantined.
verdict=CONDITIONAL_SUPPORT_GATE368_TARGET_INHERITED;CONDITIONAL_SUPPORT_LEFT_RIGHT_CURVATURE_INHERITED
```

## Eta grading formalization

```text
executed=true
support_basis=[L-heavy support R/opposite Majorana support]
generation_basis=[g1 g2 g3]
native_support_eta=[1 -1] balanced_support_eta=[1 1] generation_eta_candidate=[2 -2 1] native_eta_generation=false
curvature=C_LR = Omega_Hsigma Omega_Hsigma^dagger - J_swap Omega_Hsigma^dagger Omega_Hsigma J_swap^{-1}
projection=K_eta = Pi_gen Tr_support(eta_support C_LR); generation-dependent eta requires a separate derivation
verdict=CONDITIONAL_SUPPORT_ETA_GRADING_OPERATOR_FORMALIZED;CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED
```

## Candidate lane table

| Lane | Candidate | Native? | Circular? | `K_eta` | Decomposition | Flavor action | Verdict |
|---|---|---:|---:|---|---|---|---|
| A | native support eta trace | true | false | `diag(2, 2, 2)` | `a=2, b=0, residual=0` | central=true, breaks=false | `CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED;CONDITIONAL_SUPPORT_TARGET_DECOMPOSITION_AUDITED;CONDITIONAL_TENSION_NATIVE_SUPPORT_ETA_TRACE_IS_GENERATION_CENTRAL;FAILED_ROUTE_ETA_GRADED_TRACE_REMAINS_GENERATION_CENTRAL;FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO` |
| B | balanced support cancellation trace | true | false | `diag(0, 0, 0)` | `a=0, b=0, residual=0` | central=true, breaks=false | `CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED;CONDITIONAL_TENSION_BALANCED_SUPPORT_TRACE_CAN_CANCEL_TO_ZERO;FAILED_ROUTE_ETA_GRADED_TRACE_REMAINS_GENERATION_CENTRAL;FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO` |
| C | B-gap coupled native support eta trace | true | false | `diag(0.204929842382, 0.204929842382, 0.204929842382)` | `a=0.204929842382, b=0, residual=0` | central=true, breaks=false | `CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED;CONDITIONAL_TENSION_BGAP_COUPLING_DOES_NOT_CREATE_FLAVOR_ASYMMETRY;FAILED_ROUTE_ETA_GRADED_TRACE_REMAINS_GENERATION_CENTRAL;FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO` |
| D | generation eta insertion capacity witness | false | true | `diag(0.204929842382, -0.204929842382, 0.102464921191)` | `a=0, b=0.102464921191, residual=0` | central=false, breaks=true | `CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED;CONDITIONAL_SUPPORT_THERMAL_TIME_ACTIVATION_SIEVE_EXECUTED;CONDITIONAL_SUPPORT_NONCENTRAL_CAPACITY_WITNESSED_UNDER_GENERATION_ETA_INSERTION;CONDITIONAL_TENSION_GENERATION_ETA_INSERTION_WOULD_BE_CIRCULAR;CONDITIONAL_TENSION_TAU_ETA_NOT_EXTRACTED_FROM_NATIVE_SUPPORT_TRACE;FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED;FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED` |

## Lane A — native support eta trace

```text
lane=A name=native support eta trace
formula=K = Pi_gen Tr_support(eta_LR C_LR)
eta_source=eta_LR=diag(+1,-1) on Left/Right support, tensor I_3 on generations
native=true circular=false
K=diag[2.000000000000 2.000000000000 2.000000000000]
self_adjoint=true central=true noncentral=false
decomposition: a=2.000000000000 b=0.000000000000 residual=0.000000000000e+00 exact=true nonzero_b=false target=false verdict=FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO
KMS: beta=1.000000000000 rho=[0.333333333333 0.333333333333 0.333333333333] faithful=true nontrivial=false
omega_1-2=0.000000000000 nonzero=false
omega_1-3=0.000000000000 nonzero=false
omega_2-3=0.000000000000 nonzero=false
[E_12,K]_norm=0.000000000000 nonzero=false
[E_13,K]_norm=0.000000000000 nonzero=false
[E_23,K]_norm=0.000000000000 nonzero=false
breaks_flavor=false promoted=false selects_vacuum=false
verdict=CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED;CONDITIONAL_SUPPORT_TARGET_DECOMPOSITION_AUDITED;CONDITIONAL_TENSION_NATIVE_SUPPORT_ETA_TRACE_IS_GENERATION_CENTRAL;FAILED_ROUTE_ETA_GRADED_TRACE_REMAINS_GENERATION_CENTRAL;FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO
```

## Lane B — balanced support cancellation trace

```text
lane=B name=balanced support cancellation trace
formula=K = Pi_gen Tr_support(eta_balanced C_LR)
eta_source=eta_balanced=diag(+1,+1) on symmetric support diagnostic
native=true circular=false
K=diag[0.000000000000 0.000000000000 0.000000000000]
self_adjoint=true central=true noncentral=false
decomposition: a=0.000000000000 b=0.000000000000 residual=0.000000000000e+00 exact=true nonzero_b=false target=false verdict=FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO
KMS: beta=1.000000000000 rho=[0.333333333333 0.333333333333 0.333333333333] faithful=true nontrivial=false
omega_1-2=0.000000000000 nonzero=false
omega_1-3=0.000000000000 nonzero=false
omega_2-3=0.000000000000 nonzero=false
[E_12,K]_norm=0.000000000000 nonzero=false
[E_13,K]_norm=0.000000000000 nonzero=false
[E_23,K]_norm=0.000000000000 nonzero=false
breaks_flavor=false promoted=false selects_vacuum=false
verdict=CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED;CONDITIONAL_TENSION_BALANCED_SUPPORT_TRACE_CAN_CANCEL_TO_ZERO;FAILED_ROUTE_ETA_GRADED_TRACE_REMAINS_GENERATION_CENTRAL;FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO
```

## Lane C — B-gap coupled native support eta trace

```text
lane=C name=B-gap coupled native support eta trace
formula=K = B_gap · Pi_gen Tr_support(eta_LR C_LR)
eta_source=native support eta with already-derived B_gap scalar coupling
native=true circular=false
K=diag[0.204929842382 0.204929842382 0.204929842382]
self_adjoint=true central=true noncentral=false
decomposition: a=0.204929842382 b=0.000000000000 residual=0.000000000000e+00 exact=true nonzero_b=false target=false verdict=FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO
KMS: beta=1.000000000000 rho=[0.333333333333 0.333333333333 0.333333333333] faithful=true nontrivial=false
omega_1-2=0.000000000000 nonzero=false
omega_1-3=0.000000000000 nonzero=false
omega_2-3=0.000000000000 nonzero=false
[E_12,K]_norm=0.000000000000 nonzero=false
[E_13,K]_norm=0.000000000000 nonzero=false
[E_23,K]_norm=0.000000000000 nonzero=false
breaks_flavor=false promoted=false selects_vacuum=false
verdict=CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED;CONDITIONAL_TENSION_BGAP_COUPLING_DOES_NOT_CREATE_FLAVOR_ASYMMETRY;FAILED_ROUTE_ETA_GRADED_TRACE_REMAINS_GENERATION_CENTRAL;FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO
```

## Lane D — generation eta insertion capacity witness

```text
lane=D name=generation eta insertion capacity witness
formula=K = B_gap · tau_eta
eta_source=eta_gen=tau_eta inserted on generation orbit rather than derived from support trace
native=false circular=true
K=diag[0.204929842382 -0.204929842382 0.102464921191]
self_adjoint=true central=false noncentral=true
decomposition: a=0.000000000000 b=0.102464921191 residual=0.000000000000e+00 exact=true nonzero_b=true target=true verdict=CONDITIONAL_SUPPORT_TAU_ETA_HAMILTONIAN_DERIVED_FROM_ETA_GRADED_LR_TRACE
KMS: beta=1.000000000000 rho=[0.276663070916 0.416822345111 0.306514583973] faithful=true nontrivial=true
omega_1-2=-0.409859684764 nonzero=true
omega_1-3=-0.102464921191 nonzero=true
omega_2-3=0.307394763573 nonzero=true
[E_12,K]_norm=0.579629124863 nonzero=true
[E_13,K]_norm=0.144907281216 nonzero=true
[E_23,K]_norm=0.434721843647 nonzero=true
breaks_flavor=true promoted=false selects_vacuum=false
verdict=CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED;CONDITIONAL_SUPPORT_THERMAL_TIME_ACTIVATION_SIEVE_EXECUTED;CONDITIONAL_SUPPORT_NONCENTRAL_CAPACITY_WITNESSED_UNDER_GENERATION_ETA_INSERTION;CONDITIONAL_TENSION_GENERATION_ETA_INSERTION_WOULD_BE_CIRCULAR;CONDITIONAL_TENSION_TAU_ETA_NOT_EXTRACTED_FROM_NATIVE_SUPPORT_TRACE;FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED;FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED
```

## Thermal-time activation sieve

```text
executed=true native_target=false circular_capacity=true promoted_native=false internal_time=false energy_constraint=false
answer=Native eta support traces execute, but they remain generation-central. The only noncentral target hit is the generation-eta insertion lane, which is circular and therefore cannot activate internal thermal time.
next=derive a representation theorem that maps support eta defects to generation-dependent weights, or prove that all native support eta contractions factor through I_3.
verdict=CONDITIONAL_SUPPORT_THERMAL_TIME_ACTIVATION_SIEVE_EXECUTED;CONDITIONAL_TENSION_NATIVE_SUPPORT_ETA_TRACE_IS_GENERATION_CENTRAL;CONDITIONAL_TENSION_TAU_ETA_NOT_EXTRACTED_FROM_NATIVE_SUPPORT_TRACE;CONDITIONAL_TENSION_INTERNAL_FLOW_NOT_ACTIVATED_BY_NATIVE_TRACE;FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED;FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED;FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_ETA_TRACE;FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_ETA_TRACE;FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_ETA_TRACE
```

## Landscape preservation

```text
executed=true weak=true quartic=true alphaGUT=true morita=true bgap=true omega=true no_empirical_flavor=true no_mass=true no_vacuum=true polluted=false
verdict=CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED
```

## Kinetic safety

```text
executed=true all_self_adjoint=true faithful_states=true no_rank_collapse=true no_ghost=true no_nonunitary_push=true
verdict=CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED
```

## Vacuum parameter census

```text
starting=15 reduction=0 remaining=15 seven_seal_target=false
verdict=CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED;FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED
```

## Status ledger

```text
CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED
CONDITIONAL_SUPPORT_ETA_GRADING_OPERATOR_FORMALIZED
CONDITIONAL_SUPPORT_GATE368_TARGET_INHERITED
CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED
CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED
CONDITIONAL_SUPPORT_LEFT_RIGHT_CURVATURE_INHERITED
CONDITIONAL_SUPPORT_NONCENTRAL_CAPACITY_WITNESSED_UNDER_GENERATION_ETA_INSERTION
CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED
CONDITIONAL_SUPPORT_TARGET_DECOMPOSITION_AUDITED
CONDITIONAL_SUPPORT_TAU_ETA_HAMILTONIAN_DERIVED_FROM_ETA_GRADED_LR_TRACE
CONDITIONAL_SUPPORT_THERMAL_TIME_ACTIVATION_SIEVE_EXECUTED
CONDITIONAL_TENSION_BALANCED_SUPPORT_TRACE_CAN_CANCEL_TO_ZERO
CONDITIONAL_TENSION_BGAP_COUPLING_DOES_NOT_CREATE_FLAVOR_ASYMMETRY
CONDITIONAL_TENSION_GENERATION_ETA_INSERTION_WOULD_BE_CIRCULAR
CONDITIONAL_TENSION_INTERNAL_FLOW_NOT_ACTIVATED_BY_NATIVE_TRACE
CONDITIONAL_TENSION_NATIVE_SUPPORT_ETA_TRACE_IS_GENERATION_CENTRAL
CONDITIONAL_TENSION_TAU_ETA_NOT_EXTRACTED_FROM_NATIVE_SUPPORT_TRACE
FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_ETA_TRACE
FAILED_ROUTE_ETA_GRADED_TRACE_REMAINS_GENERATION_CENTRAL
FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED
FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO
FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED
FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED
FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_ETA_TRACE
FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_ETA_TRACE
```

## Final truth statement

Gate 369 executes the eta-graded Left-Right trace target isolated by Gate 368. The lawful support eta gradings act on the heavy/Majorana Left-Right support uniformly over generations; their projected Hamiltonians are zero or proportional to I_3, so b=0 in aI_3+b tau_eta. A generation-eta insertion reproduces tau_eta and activates noncentral KMS capacity, but it is circular because the generation grading was assumed rather than extracted from the support trace. Therefore the eta-graded Left-Right trace does not yet derive internal thermal time, and the 15 vacuum coordinates remain quarantined.

## Next lawful theorem target

Gate 369 proves that the current native support eta trace is not enough. The next theorem must decide whether there exists a native representation map from support defects to generation-dependent weights, or whether every admissible eta-graded Left-Right contraction factors through `I_3`.
