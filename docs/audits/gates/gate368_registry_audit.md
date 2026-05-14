# Gate 368 Registry Audit — Bimodule Modular Curvature / Internal Thermal Time Origin Sieve

## Gate identity

- **Gate:** 368
- **Package:** `pkg/bridge/bimodulemodularcurvature`
- **Theorem:** `BimoduleModularCurvatureInternalThermalTimeOriginSieveTheorem`
- **Audit ID:** `GATE368-BIMODULE-MODULAR-CURVATURE-INTERNAL-THERMAL-TIME-ORIGIN-SIEVE`
- **Layer:** Bridge / Phase-III Flow Extension
- **Purpose:** test whether finite Left-Right bimodule curvature derives the internal modular Hamiltonian after ordinary Lorentzian time failed on flavor space.

## Inherited theorem chain

| Gate | Inherited fact | Gate-368 consequence |
|---:|---|---|
| 319/320 | Heavy-light `Omega_Hsigma` support index exists conditionally with `Omega=1`; explicit matrix promotion remains firewalled. | The overlap may be tested only as a support ingredient, not as an already-derived generation Hamiltonian. |
| 347 | Standard Majorana/Dirac cross-terms and pure `Omega_Hsigma` are flavor-unitary invariant. | Pure support overlap is expected to be central unless an eta-graded projection is derived. |
| 359 | Flavor-orientation templates remain quarantined unless a native noncentral operator is derived. | No CKM/PMNS/Yukawa data may enter this gate. |
| 362 | Phase III demands flow-based vacuum selection. | Static texture search is forbidden. |
| 363 | Tracial modular state gives `Delta=I`. | A nontracial source is required. |
| 364–366 | `tau_eta` has modular capacity but is not selected as energy. | `tau_eta` may not be silently chosen. |
| 367 | `e0/gamma0` is physical time but flavor-central. | The next source must be internal finite curvature. |

## Formalization

```text
executed=true
left=rho_L(A_F) on Standard Model generation carriers
right=rho_R(A_F^op)=J_swap rho_L(A_F*) J_swap^{-1} on the doubled/opposite Majorana sector
tomita=sigma_t(a)=Delta^{it} a Delta^{-it}; finite candidate K=-log(rho) must be self-adjoint and noncentral on generation space
curvature=C_LR = Omega_Hsigma Omega_Hsigma^dagger - J_swap Omega_Hsigma^dagger Omega_Hsigma J_swap^{-1}
projection=K_LR = Pi_gen Tr_support^eta(C_LR)
needs_eta=true forbids_manual_tau=true
verdict=CONDITIONAL_SUPPORT_BIMODULE_MODULAR_CURVATURE_FORMALIZED;CONDITIONAL_SUPPORT_LEFT_RIGHT_COMMUTANT_FRAMEWORK_AUDITED;CONDITIONAL_SUPPORT_HEAVY_LIGHT_OVERLAP_INHERITED
```

## Candidate lane table

| Lane | Candidate | Generation result | Flavor action | Verdict |
|---|---|---|---|---|
| A | pure B-gap scalar lane | `diag(0.102464921191, 0.102464921191, 0.102464921191)` | central=true, noncentral=false, breaks=false | `CONDITIONAL_SUPPORT_BGAP_THERMAL_COUPLING_AUDITED;CONDITIONAL_TENSION_PURE_BGAP_IS_FLAVOR_CENTRAL;FAILED_ROUTE_BIMODULE_MODULAR_CURVATURE_NOT_NONCENTRAL` |
| B | pure Omega_Hsigma support lane | `diag(1, 1, 1)` | central=true, noncentral=false, breaks=false | `CONDITIONAL_SUPPORT_HEAVY_LIGHT_OVERLAP_INHERITED;CONDITIONAL_TENSION_PURE_OMEGA_OVERLAP_IS_SUPPORT_INDEX_NOT_GENERATION_HAMILTONIAN;FAILED_ROUTE_BIMODULE_MODULAR_CURVATURE_NOT_NONCENTRAL` |
| C | ungraded Left-Right commutant curvature lane | `diag(0, 0, 0)` | central=true, noncentral=false, breaks=false | `CONDITIONAL_SUPPORT_GENERATION_PROJECTION_EXECUTED;CONDITIONAL_TENSION_LR_CURVATURE_REQUIRES_ETA_GRADED_PROJECTION;FAILED_ROUTE_BIMODULE_MODULAR_CURVATURE_NOT_NONCENTRAL` |
| D | eta-weighted triality curvature capacity lane | `diag(0.204929842382, -0.204929842382, 0.102464921191)` | central=false, noncentral=true, breaks=true | `CONDITIONAL_SUPPORT_GENERATION_PROJECTION_EXECUTED;CONDITIONAL_SUPPORT_FLAVOR_COMMUTATOR_SIEVE_EXECUTED;CONDITIONAL_SUPPORT_KMS_RECONSTRUCTION_EXECUTED;CONDITIONAL_SUPPORT_NONTRIVIAL_MODULAR_CAPACITY_WITNESSED_UNDER_ETA_INSERTION;CONDITIONAL_TENSION_TAU_ETA_INSERTION_WOULD_BE_CIRCULAR;FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED;FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED` |

## Lane A — pure B-gap scalar lane

```text
lane=A name=pure B-gap scalar lane
formula=K = B_gap · I_3
source=B_gap trace capacity alone
K=diag[0.102464921191 0.102464921191 0.102464921191] beta=1.000000000000
self_adjoint=true central=true noncentral=false native=true tau_derived=false tau_inserted=false projected=true
rho=[0.333333333333 0.333333333333 0.333333333333]
omega_1-2=0.000000000000 nonzero=false
omega_1-3=0.000000000000 nonzero=false
omega_2-3=0.000000000000 nonzero=false
[E_12,K]_norm=0.000000000000 nonzero=false
[E_13,K]_norm=0.000000000000 nonzero=false
[E_23,K]_norm=0.000000000000 nonzero=false
breaks_flavor=false selects_vacuum=false
verdict=CONDITIONAL_SUPPORT_BGAP_THERMAL_COUPLING_AUDITED;CONDITIONAL_TENSION_PURE_BGAP_IS_FLAVOR_CENTRAL;FAILED_ROUTE_BIMODULE_MODULAR_CURVATURE_NOT_NONCENTRAL
```

## Lane B — pure Omega_Hsigma support lane

```text
lane=B name=pure Omega_Hsigma support lane
formula=K = Tr_support(Omega_Hsigma^dagger Omega_Hsigma) · I_3
source=Gate-319/320 heavy-light support index Omega_Hsigma=1
K=diag[1.000000000000 1.000000000000 1.000000000000] beta=0.102464921191
self_adjoint=true central=true noncentral=false native=true tau_derived=false tau_inserted=false projected=true
rho=[0.333333333333 0.333333333333 0.333333333333]
omega_1-2=0.000000000000 nonzero=false
omega_1-3=0.000000000000 nonzero=false
omega_2-3=0.000000000000 nonzero=false
[E_12,K]_norm=0.000000000000 nonzero=false
[E_13,K]_norm=0.000000000000 nonzero=false
[E_23,K]_norm=0.000000000000 nonzero=false
breaks_flavor=false selects_vacuum=false
verdict=CONDITIONAL_SUPPORT_HEAVY_LIGHT_OVERLAP_INHERITED;CONDITIONAL_TENSION_PURE_OMEGA_OVERLAP_IS_SUPPORT_INDEX_NOT_GENERATION_HAMILTONIAN;FAILED_ROUTE_BIMODULE_MODULAR_CURVATURE_NOT_NONCENTRAL
```

## Lane C — ungraded Left-Right commutant curvature lane

```text
lane=C name=ungraded Left-Right commutant curvature lane
formula=K = Pi_gen Tr_support(C_LR)
source=ungraded LR difference after J_swap pairing
K=diag[0.000000000000 0.000000000000 0.000000000000] beta=1.000000000000
self_adjoint=true central=true noncentral=false native=true tau_derived=false tau_inserted=false projected=true
rho=[0.333333333333 0.333333333333 0.333333333333]
omega_1-2=0.000000000000 nonzero=false
omega_1-3=0.000000000000 nonzero=false
omega_2-3=0.000000000000 nonzero=false
[E_12,K]_norm=0.000000000000 nonzero=false
[E_13,K]_norm=0.000000000000 nonzero=false
[E_23,K]_norm=0.000000000000 nonzero=false
breaks_flavor=false selects_vacuum=false
verdict=CONDITIONAL_SUPPORT_GENERATION_PROJECTION_EXECUTED;CONDITIONAL_TENSION_LR_CURVATURE_REQUIRES_ETA_GRADED_PROJECTION;FAILED_ROUTE_BIMODULE_MODULAR_CURVATURE_NOT_NONCENTRAL
```

## Lane D — eta-weighted triality curvature capacity lane

```text
lane=D name=eta-weighted triality curvature capacity lane
formula=K = B_gap · tau_eta, tested only as eta-inserted capacity witness
source=requires eta/tau_eta insertion not derived by this LR contraction
K=diag[0.204929842382 -0.204929842382 0.102464921191] beta=1.000000000000
self_adjoint=true central=false noncentral=true native=false tau_derived=false tau_inserted=true projected=false
rho=[0.276663070916 0.416822345111 0.306514583973]
omega_1-2=-0.409859684764 nonzero=true
omega_1-3=-0.102464921191 nonzero=true
omega_2-3=0.307394763573 nonzero=true
[E_12,K]_norm=0.579629124863 nonzero=true
[E_13,K]_norm=0.144907281216 nonzero=true
[E_23,K]_norm=0.434721843647 nonzero=true
breaks_flavor=true selects_vacuum=false
verdict=CONDITIONAL_SUPPORT_GENERATION_PROJECTION_EXECUTED;CONDITIONAL_SUPPORT_FLAVOR_COMMUTATOR_SIEVE_EXECUTED;CONDITIONAL_SUPPORT_KMS_RECONSTRUCTION_EXECUTED;CONDITIONAL_SUPPORT_NONTRIVIAL_MODULAR_CAPACITY_WITNESSED_UNDER_ETA_INSERTION;CONDITIONAL_TENSION_TAU_ETA_INSERTION_WOULD_BE_CIRCULAR;FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED;FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED
```

## KMS reconstruction

```text
executed=true best=eta-weighted triality curvature capacity lane
rho=[0.276663070916 0.416822345111 0.306514583973]
omega_1-2=-0.409859684764 nonzero=true
omega_1-3=-0.102464921191 nonzero=true
omega_2-3=0.307394763573 nonzero=true
nontrivial=true promoted_native=false energy_constraint=false
verdict=CONDITIONAL_SUPPORT_KMS_RECONSTRUCTION_EXECUTED;CONDITIONAL_TENSION_TAU_ETA_INSERTION_WOULD_BE_CIRCULAR;FAILED_ROUTE_MODULAR_ENERGY_CONSTRAINT_NOT_DERIVED;FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED
```

## Landscape preservation

```text
executed=true weak=true quartic=true alphaGUT=true morita=true bgap=true omega=true no_empirical_flavor=true no_mass=true no_vacuum=true polluted=false
verdict=CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED
```

## Kinetic safety

```text
executed=true all_self_adjoint=true faithful_state=true no_rank_collapse=true no_ghost=true no_nonunitary_push=true
verdict=CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED
```

## Flow verdict

```text
executed=true native_noncentral=false tau_eta_derived=false nontrivial_capacity=true promoted_native=false selects_vacuum=false
answer=The Left-Right bimodule ingredients are real and safe, but their native scalar/support/ungraded projections remain flavor-central. The only noncentral lane requires eta/tau_eta insertion, so it witnesses capacity but does not derive the internal thermal-time origin.
next=derive an eta-graded Left-Right trace theorem that produces Pi_gen Tr_support^eta(C_LR)=aI_3+b tau_eta with b != 0, or prove this route impossible.
verdict=CONDITIONAL_SUPPORT_FLAVOR_COMMUTATOR_SIEVE_EXECUTED;CONDITIONAL_TENSION_PURE_BGAP_IS_FLAVOR_CENTRAL;CONDITIONAL_TENSION_PURE_OMEGA_OVERLAP_IS_SUPPORT_INDEX_NOT_GENERATION_HAMILTONIAN;CONDITIONAL_TENSION_LR_CURVATURE_REQUIRES_ETA_GRADED_PROJECTION;CONDITIONAL_TENSION_TAU_ETA_INSERTION_WOULD_BE_CIRCULAR;CONDITIONAL_TENSION_INTERNAL_FLOW_NONTRIVIAL_BUT_NOT_VACUUM_SELECTING;FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED;FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED;FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_BIMODULE_FLOW;FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_BIMODULE_FLOW;FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_BIMODULE_FLOW
```

## Vacuum parameter census

```text
starting=15 reduction=0 remaining=15 seven_seal_target=false
verdict=CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED;FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED
```

## Status ledger

```text
CONDITIONAL_SUPPORT_BGAP_THERMAL_COUPLING_AUDITED
CONDITIONAL_SUPPORT_BIMODULE_MODULAR_CURVATURE_FORMALIZED
CONDITIONAL_SUPPORT_FLAVOR_COMMUTATOR_SIEVE_EXECUTED
CONDITIONAL_SUPPORT_GENERATION_PROJECTION_EXECUTED
CONDITIONAL_SUPPORT_HEAVY_LIGHT_OVERLAP_INHERITED
CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED
CONDITIONAL_SUPPORT_KMS_RECONSTRUCTION_EXECUTED
CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED
CONDITIONAL_SUPPORT_LEFT_RIGHT_COMMUTANT_FRAMEWORK_AUDITED
CONDITIONAL_SUPPORT_NONTRIVIAL_MODULAR_CAPACITY_WITNESSED_UNDER_ETA_INSERTION
CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED
CONDITIONAL_TENSION_INTERNAL_FLOW_NONTRIVIAL_BUT_NOT_VACUUM_SELECTING
CONDITIONAL_TENSION_LR_CURVATURE_REQUIRES_ETA_GRADED_PROJECTION
CONDITIONAL_TENSION_PURE_BGAP_IS_FLAVOR_CENTRAL
CONDITIONAL_TENSION_PURE_OMEGA_OVERLAP_IS_SUPPORT_INDEX_NOT_GENERATION_HAMILTONIAN
CONDITIONAL_TENSION_TAU_ETA_INSERTION_WOULD_BE_CIRCULAR
FAILED_ROUTE_BIMODULE_MODULAR_CURVATURE_NOT_NONCENTRAL
FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_BIMODULE_FLOW
FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED
FAILED_ROUTE_MODULAR_ENERGY_CONSTRAINT_NOT_DERIVED
FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED
FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED
FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_BIMODULE_FLOW
FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_BIMODULE_FLOW
```

## Final truth statement

Gate 368 audits whether finite Left-Right Morita bimodule curvature supplies the missing internal thermal-time Hamiltonian after Gate 367 ruled out ordinary Lorentzian time. Pure B-gap, pure Omega_Hsigma support, and ungraded Left-Right curvature project centrally on generation space. An eta/tau_eta-weighted lane has the correct noncentral KMS capacity, but the present bimodule audit does not derive tau_eta from the Left-Right contraction; inserting it would be circular. Therefore internal thermal time is not yet derived, and the 15 vacuum coordinates remain quarantined.

The exact success target remains:

```text
Pi_gen Tr_support^eta(C_LR) = aI_3 + b tau_eta, b != 0
```

Gate 368 does not prove that target. It proves that the route is now sharply localized: either the eta-graded Left-Right trace derives the noncentral part, or the bimodule modular-curvature origin route must be rejected.
