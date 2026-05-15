# Gate 462 Registry Audit — Sector-Difference Invariant / CKM Interface Firewall Audit

## Verdict

`CONDITIONAL_SUPPORT_SECTOR_DIFFERENCE_CKM_INTERFACE_FIREWALL_VALIDATED`

Gate 462 isolates the relative u-d sector ray that a future CKM-facing bridge adapter may inspect. It does not compute CKM entries, PMNS entries, physical masses, Yukawa values, or native mixing predictions.

## Inheritance

executed=true K=true triangle=true inverse=true provenance=true branches=true residual=true multiplex=true independent=true native_universality_rejected=true contamination_rejected=true no_observed=true verdict=CONDITIONAL_SUPPORT_GATE461_SECTOR_MULTIPLEX_INHERITED

## Interface contract

executed=true require_u=true require_d=true provenance=true tags=true eigenbasis=true relative_dof=2 diagnostics_only=true reject_ckm_native=true reject_pmns_charged=true reject_observed=true reject_native_relative=true verdict=CONDITIONAL_SUPPORT_RELATIVE_RAY_LEDGER_DEFINED reason=a CKM-facing relative ray requires complete u and d bridge rays plus an explicit eigenbasis convention, and exports only relative diagnostics

```text
required charged sectors: u,d
relative ray: Delta_alpha = alpha_d-alpha_u
relative phase: Delta_phi = wrap_pi(phi_d-phi_u)
projective diagnostic: distance = sqrt(Delta_alpha^2 + (2 sin(Delta_phi/2))^2)
forbidden export: CKM_ij, PMNS_ij, masses, Yukawas, GST/Fritzsch relations, native coefficient values
```

## Sieve

executed=true accepted=1 rejected=8 valid_ud=true missing_sector=true missing_provenance=true missing_eigenbasis=true observed=true native_prediction=true native_relative=true lepton_misroute=true universality_native=true all_bridge=true no_native_mixing=true verdict=CONDITIONAL_SUPPORT_SECTOR_DIFFERENCE_CKM_INTERFACE_FIREWALL_VALIDATED reason=only the synthetic, fully provenanced u-d relative ray is accepted, and it exports diagnostics rather than CKM/PMNS observables

| Case | Accepted | Verdict | Relative diagnostic | Reason |
|---|---|---|---|---|
| valid synthetic u-d relative ray | true | `CONDITIONAL_SUPPORT_UD_SECTOR_DIFFERENCE_BRIDGE_ONLY_COMPUTED` | u->d delta_alpha=-0.3 delta_phi=0.55 phase_chord=0.543093873912 distance=0.620444160163 complete=true eigenbasis=true bridge=true exports_CKM=false exports_PMNS=false native_export=false verdict=CONDITIONAL_SUPPORT_UD_SECTOR_DIFFERENCE_BRIDGE_ONLY_COMPUTED reason=relative u-d coefficient-ray diagnostics were computed without exporting CKM/PMNS entries | relative u-d coefficient-ray diagnostics were computed without exporting CKM/PMNS entries |
| missing d sector | false | `FAILED_ROUTE_SECTOR_DIFFERENCE_REQUIRES_TWO_PROVENANCED_RAYS` | u->d delta_alpha=0 delta_phi=0 phase_chord=0 distance=0 complete=false eigenbasis=false bridge=false exports_CKM=false exports_PMNS=false native_export=false verdict=FAILED_ROUTE_SECTOR_DIFFERENCE_REQUIRES_TWO_PROVENANCED_RAYS reason=u and d rays must both be bridge-only, provenanced, and branch-tagged | u and d rays must both be bridge-only, provenanced, and branch-tagged |
| missing provenance | false | `FAILED_ROUTE_SECTOR_DIFFERENCE_REQUIRES_TWO_PROVENANCED_RAYS` | u->d delta_alpha=0 delta_phi=0 phase_chord=0 distance=0 complete=false eigenbasis=false bridge=false exports_CKM=false exports_PMNS=false native_export=false verdict=FAILED_ROUTE_SECTOR_DIFFERENCE_REQUIRES_TWO_PROVENANCED_RAYS reason=u and d rays must both be bridge-only, provenanced, and branch-tagged | u and d rays must both be bridge-only, provenanced, and branch-tagged |
| missing eigenbasis convention | false | `FAILED_ROUTE_UNLABELLED_EIGENBASIS_CONVENTION_REJECTED` | u->d delta_alpha=0 delta_phi=0 phase_chord=0 distance=0 complete=false eigenbasis=false bridge=false exports_CKM=false exports_PMNS=false native_export=false verdict=FAILED_ROUTE_UNLABELLED_EIGENBASIS_CONVENTION_REJECTED reason=a CKM-facing relative comparison requires explicit eigenvalue/eigenvector ordering and phase-gauge conventions | a CKM-facing relative comparison requires explicit eigenvalue/eigenvector ordering and phase-gauge conventions |
| observed CKM/PMNS import | false | `FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED` | u->d delta_alpha=0 delta_phi=0 phase_chord=0 distance=0 complete=false eigenbasis=false bridge=false exports_CKM=false exports_PMNS=false native_export=false verdict=FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED reason=observed CKM/PMNS values are not accepted in this native-boundary audit | observed CKM/PMNS values are not accepted in this native-boundary audit |
| native CKM prediction claim | false | `FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED` | u->d delta_alpha=0 delta_phi=0 phase_chord=0 distance=0 complete=false eigenbasis=false bridge=false exports_CKM=false exports_PMNS=false native_export=false verdict=FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED reason=observed mixing plus a native-prediction claim is forbidden | observed mixing plus a native-prediction claim is forbidden |
| native relative-ray promotion | false | `FAILED_ROUTE_RELATIVE_RAY_NATIVE_PROMOTION_REJECTED` | u->d delta_alpha=0 delta_phi=0 phase_chord=0 distance=0 complete=false eigenbasis=false bridge=false exports_CKM=false exports_PMNS=false native_export=false verdict=FAILED_ROUTE_RELATIVE_RAY_NATIVE_PROMOTION_REJECTED reason=relative rays are bridge diagnostics and cannot be promoted to native ASHA law | relative rays are bridge diagnostics and cannot be promoted to native ASHA law |
| lepton PMNS misrouted into charged CKM ledger | false | `FAILED_ROUTE_LEPTON_PMNS_SECTOR_MISRouted_TO_CHARGED_CKM_LEDGER` | u->d delta_alpha=0 delta_phi=0 phase_chord=0 distance=0 complete=false eigenbasis=false bridge=false exports_CKM=false exports_PMNS=false native_export=false verdict=FAILED_ROUTE_LEPTON_PMNS_SECTOR_MISRouted_TO_CHARGED_CKM_LEDGER reason=the charged CKM relative-ray ledger requires u and d sectors; lepton/PMNS records need a separate neutrino-sector interface | the charged CKM relative-ray ledger requires u and d sectors; lepton/PMNS records need a separate neutrino-sector interface |
| cross-sector universality as native law | false | `FAILED_ROUTE_CROSS_SECTOR_RAY_UNIVERSALITY_NOT_NATIVE` | u->d delta_alpha=0 delta_phi=0 phase_chord=0 distance=0 complete=false eigenbasis=false bridge=false exports_CKM=false exports_PMNS=false native_export=false verdict=FAILED_ROUTE_CROSS_SECTOR_RAY_UNIVERSALITY_NOT_NATIVE reason=Gate461 already rejected cross-sector coefficient-ray universality as native law | Gate461 already rejected cross-sector coefficient-ray universality as native law |

## Accepted relative-ray equation

For the accepted synthetic u-d bridge row only:

```text
Delta_alpha_ud = alpha_d - alpha_u
Delta_phi_ud   = wrap_pi(phi_d - phi_u)
d_ud           = sqrt(Delta_alpha_ud^2 + 4 sin^2(Delta_phi_ud/2))
```

This is a comparator diagnostic. It is not a CKM matrix element and not a native ASHA observable.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE461_SECTOR_MULTIPLEX_INHERITED`
- `CONDITIONAL_SUPPORT_RELATIVE_RAY_LEDGER_DEFINED`
- `CONDITIONAL_SUPPORT_UD_SECTOR_DIFFERENCE_BRIDGE_ONLY_COMPUTED`
- `CONDITIONAL_SUPPORT_SECTOR_DIFFERENCE_CKM_INTERFACE_FIREWALL_VALIDATED`
- `CONDITIONAL_SUPPORT_PMNS_INTERFACE_FIREWALL_VALIDATED_BY_REJECTION`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED`
- `FAILED_ROUTE_SECTOR_DIFFERENCE_REQUIRES_TWO_PROVENANCED_RAYS`
- `FAILED_ROUTE_UNLABELLED_EIGENBASIS_CONVENTION_REJECTED`
- `FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED`
- `FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED`
- `FAILED_ROUTE_RELATIVE_RAY_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_LEPTON_PMNS_SECTOR_MISRouted_TO_CHARGED_CKM_LEDGER`
- `FAILED_ROUTE_CROSS_SECTOR_RAY_UNIVERSALITY_NOT_NATIVE`

## Firewall

executed=true ray_may_feed_CKM_adapter=true CKM_computed=false CKM_native=false PMNS_computed=false PMNS_native=false requires_observed_import=true provenance=true eigenbasis=true K=true triangle=true Y_sealed=true coeffs_sealed=true no_masses=true no_yukawas=true no_CKM=true no_PMNS=true no_GST=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED reason=Gate462 creates only a relative-ray bridge diagnostic; CKM/PMNS entries and physical flavor data remain quarantined.

## CKM/PMNS boundary

A CKM-like adapter would need relative u-d diagonalization data plus explicit eigenbasis ordering and phase-gauge conventions. Gate 462 only constructs the relative-ray slot. PMNS requires a neutrino-sector ledger and is rejected in this charged-sector audit.

## Next gate

Gate 463 — Eigenbasis Convention Ledger / Mixing-Matrix Gauge Audit: Gate462 shows that a CKM-facing interface needs not only u-d relative rays but also explicit eigenvalue ordering and phase-gauge conventions. Primary task: formalize the bridge-only eigenbasis convention required before a relative u-d ray can be passed to any future CKM residual evaluator, and prove convention choices are not native predictions.

## Truth statement

Gate462 isolates the u-d sector-difference ray that a future CKM bridge adapter may inspect. The object is a labelled bridge diagnostic, not a CKM prediction: observed CKM/PMNS values, eigenbasis conventions, sector coefficients, Yukawas, and physical masses remain quarantined.
