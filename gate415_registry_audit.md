# Gate 415 Registry Audit — Family Boundary Condition / Sector Source Axiom Minimality Sieve

## Claim tested

Gate 415 asks what minimal boundary/source axiom would be required after Gate 414 failed to derive family coefficients from trace or curvature functionals. It ranks candidate axioms without promoting any of them to native ASHA theorems and without importing observed Yukawa matrices.

## Prior boundary inherited

```text
executed=true gate413_capacity=true pair_not_native=true gate414_no_selector=true trace_curvature_exhausted=true coefficients_free=true charged_moduli=13 verdict=Gate 415 inherits the Gate-414 boundary: noncommuting texture capacity exists only conditionally, while native trace/curvature functionals leave the family coefficients free.
```

## Boundary/source arena

```text
executed=true family_pair="K_gen plus S_gen/X_gen family texture pair" sectors=up,down,charged-lepton,neutrino coeff_per_sector=2 baseline_rays=8 gauge_compatible=true native_selector=false empirical_yukawa=false verdict=boundary/source axiom arena formalized reason=The family pair is gauge-blind and can be coupled to SM sectors, but ASHA supplies no source selecting the sector coefficient rays.
```

## Minimality ranking

```text
executed=true ranked=universal family source(cost=1); charge-sector source boundary(cost=2); Z3 Weyl phase source(cost=2); flat U(3)_gen holonomy boundary(cost=3); modular KMS sector source(cost=3); observed Yukawa matrix source(cost=5) least_cost_name="charge-sector source boundary" least_cost=2 least_still_axiom=true least_ckm=true least_fixes_angles=false no_candidate_native=true verdict=least-cost CKM-capable source axiom identified but quarantined
```

## CKM/PMNS capacity audit

```text
executed=true conditional_ckm=true conditional_pmns=true fixes_angles=false native=false curve_fitting_candidate=true best_empirical_independent="charge-sector source boundary" required_extra_data="a sector-source/boundary rule selecting coefficient rays for K_gen and S_gen" verdict=capacity requires explicit boundary/source axiom reason=A minimal source can activate mixing without observed Yukawa matrices, but it still supplies new data outside current ASHA.
```

## Empirical firewall

```text
executed=true no_masses=true no_ckm=true no_pmns=true no_yukawa_matrices=true axioms_quarantined=true no_native_derivation=true verdict=all boundary/source candidates remain quarantined; empirical firewall preserved
```

## Candidate axiom ledger

```text
name="universal family source" executed=true kind="single coefficient ray shared by all sectors" cost=1 gauge=true JGamma=true empirical_independent=true imports_yukawa=false selects_ray=true fixes_values=false ckm=false pmns=false diagonal_only=true free_params=1 native=false promoted=false verdict=least cost but flavor-blind reason=A universal source is mathematically cheap and gauge compatible, but it aligns all sectors and cannot produce CKM/PMNS.
name="charge-sector source boundary" executed=true kind="sector-labelled source rays for up/down/lepton/neutrino channels" cost=2 gauge=true JGamma=true empirical_independent=true imports_yukawa=false selects_ray=true fixes_values=false ckm=true pmns=true diagonal_only=false free_params=8 native=false promoted=false verdict=least cost CKM-capable axiom reason=This is the smallest audited rule that can assign distinct K/S coefficients to sectors, but the coefficient values remain boundary data.
name="Z3 Weyl phase source" executed=true kind="roots-of-unity phase rule on family shift" cost=2 gauge=true JGamma=true empirical_independent=true imports_yukawa=false selects_ray=true fixes_values=false ckm=true pmns=true diagonal_only=false free_params=4 native=false promoted=false verdict=discrete phase capacity, angle underdetermined reason=The clock/shift algebra supplies phases, but roots of unity do not determine physical mixing magnitudes or CKM angles.
name="flat U(3)_gen holonomy boundary" executed=true kind="family connection/holonomy boundary condition" cost=3 gauge=true JGamma=true empirical_independent=true imports_yukawa=false selects_ray=true fixes_values=false ckm=true pmns=true diagonal_only=false free_params=4 native=false promoted=false verdict=mixing-capable but unconstrained holonomy reason=A family holonomy can encode mixing, but without a topological quantization or source equation it is exactly a new connection axiom.
name="modular KMS sector source" executed=true kind="nontracial density/Hamiltonian plus sector source" cost=3 gauge=true JGamma=true empirical_independent=true imports_yukawa=false selects_ray=true fixes_values=false ckm=true pmns=true diagonal_only=false free_params=6 native=false promoted=false verdict=hierarchy plus mixing capacity, still sourced reason=KMS structure gives hierarchy language, but the source Hamiltonian and sector splitting are additional axioms.
name="observed Yukawa matrix source" executed=true kind="full external Yukawa amplitude ledger" cost=5 gauge=true JGamma=true empirical_independent=false imports_yukawa=true selects_ray=true fixes_values=true ckm=true pmns=true diagonal_only=false free_params=13 native=false promoted=false verdict=rejected curve fitting reason=This is phenomenologically complete, but it imports the firewall data and is not a geometric explanation.
```

## Moduli impact

```text
start_dim=13 best_native_dim=13 native_reduction=false conditional_mixing=true coefficients_free=true firewall=true verdict=boundary/source axiom ledger compiled; no native moduli reduction
scenario="current ASHA through Gate 414" status=FIREWALL_PRESERVED_13_MODULI dim=13 masses3=true ckm=true pmns=true coefficients_fixed=false native_reduction=false conditional=true empirical_fitting=false reason=Noncommuting capacity exists only under quarantined K/S axioms; coefficients remain free.
scenario="universal family source" status=CONDITIONAL_UNIVERSAL_SOURCE_FLAVOR_BLIND dim=13 masses3=true ckm=false pmns=false coefficients_fixed=false native_reduction=false conditional=true empirical_fitting=false reason=One shared source gives hierarchy scale but aligns all sectors.
scenario="charge-sector source boundary" status=CONDITIONAL_SECTOR_SOURCE_CKM_CAPACITY dim=13 masses3=true ckm=true pmns=true coefficients_fixed=false native_reduction=false conditional=true empirical_fitting=false reason=Smallest CKM-capable source, but coefficient values are boundary data.
scenario="Z3 Weyl phase source" status=FAILED_ROUTE_DISCRETE_SYMMETRY_SOURCE_UNDERDETERMINES_ANGLES dim=13 masses3=true ckm=true pmns=true coefficients_fixed=false native_reduction=false conditional=true empirical_fitting=false reason=Roots of unity constrain phases but not physical angles/magnitudes.
scenario="observed Yukawa matrix source" status=FAILED_ROUTE_OBSERVED_YUKAWA_SOURCE_REJECTED_AS_CURVE_FITTING dim=0 masses3=true ckm=true pmns=true coefficients_fixed=true native_reduction=false conditional=false empirical_fitting=true reason=Fixes data by importing the target phenomenology; rejected as explanation.
```

## Statuses

```text
CONDITIONAL_SUPPORT_GATE414_COEFFICIENT_SELECTOR_BOUNDARY_INHERITED
CONDITIONAL_SUPPORT_FAMILY_BOUNDARY_SOURCE_AXIOM_LEDGER_COMPILED
CONDITIONAL_SUPPORT_BOUNDARY_SOURCE_MINIMALITY_RANKING_AUDITED
CONDITIONAL_SUPPORT_BOUNDARY_SOURCE_CKM_PMNS_CAPACITY_AUDITED
CONDITIONAL_SUPPORT_BOUNDARY_SOURCE_EMPIRICAL_INDEPENDENCE_AUDITED
CONDITIONAL_SUPPORT_LEAST_COST_BOUNDARY_AXIOM_IDENTIFIED
FAILED_ROUTE_NO_NATIVE_FAMILY_BOUNDARY_CONDITION
FAILED_ROUTE_COEFFICIENT_SELECTOR_REQUIRES_SOURCE_AXIOM
FAILED_ROUTE_DISCRETE_SYMMETRY_SOURCE_UNDERDETERMINES_ANGLES
FAILED_ROUTE_FLAT_CONNECTION_BOUNDARY_DIAGONAL_OR_UNCONSTRAINED
FAILED_ROUTE_OBSERVED_YUKAWA_SOURCE_REJECTED_AS_CURVE_FITTING
FAILED_ROUTE_NO_NATIVE_FLAVOR_MODULI_REDUCTION
FIREWALL_PRESERVED_13_MODULI
```

## Conclusion

Gate 415 compiles the minimal boundary/source axiom ledger after Gate 414's coefficient-selector failure. The least-cost CKM-capable option is a charge-sector source boundary that assigns coefficient rays to the K/S family pair, but it remains an explicit new axiom and does not fix physical mixing angles. Discrete roots of unity, flat holonomy, and modular KMS sources all have capacity but underdetermine coefficients; observed Yukawa matrices are rejected as curve-fitting. No native ASHA boundary condition is found, and dim M_charged remains 13.

## Next gate

```text
gate=416 title="Minimal Sector-Source Axiom Consistency / Parameter-Counting Sieve" reason=Gate 415 identifies the charge-sector source boundary as the least-cost CKM-capable axiom, but its coefficients remain free. The next gate should test this explicit axiom's consistency and exact parameter count without claiming prediction. primary_task=Formalize the minimal sector-source axiom and compute how many flavor parameters it leaves free under gauge, J, Γ, and first-order compatibility.
```
