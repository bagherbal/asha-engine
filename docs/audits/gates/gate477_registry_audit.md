# Gate 477 Registry Audit — Lepton-Sector Empirical Import Switch / PMNS Data Firewall

## Verdict

`CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_IMPORT_SWITCH_VALIDATED`

Gate 477 opens the lepton observed-data airlock. It does not evaluate observed PMNS data, does not construct `U_PMNS`, does not infer `I_K`, and does not write observed values into the native theorem registry. It proves only that fully metadated lepton rows can enter a quarantined comparator ledger when `empirical_import=true`.

## Inheritance

executed=true K=true triangle=true preflight=true gate476=true synthetic_only=true observed_pmns_reject=true pmns_ray_reject=true native_prediction_reject=true matrix_export_reject=true diagnostic_only=true native_registry_clean=true verdict=CONDITIONAL_SUPPORT_GATE476_PMNS_NULL_SOCKET_INHERITED

## Airlock policy

executed=true state=empirical_import default=false explicit_true=true source=true scale=true scheme=true uncertainty=true bridge_quarantine=true ordering_policy=true absolute_scale_policy=true majorana_dirac_policy=true pmns_target=true pmns_as_ray=false ledger=lepton-sector-comparator-ledger reject_native_prediction=true reject_native_law=true reject_native_registry=true reject_theorem_input=true metadata_count=4 verdict=CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_IMPORT_AIRLOCK_DEFINED reason=external lepton/PMNS rows may pass only through empirical_import=true into a bridge-only comparator ledger with source, scale, scheme, uncertainty, and neutrino-policy metadata

```text
empirical_import default = false
empirical_import must be true before external lepton rows are accepted
required metadata = {source, scale, scheme, uncertainty}
required lepton policies = {neutrino_ordering, absolute_neutrino_scale, majorana_dirac_phase, eigenbasis_convention}
allowed target = lepton-sector-comparator-ledger
PMNS allowed role = residual target only
forbidden roles = alpha/phi/I_K source, native theorem registry, native_prediction, native_law
```

## Sieve

executed=true accepted=3 rejected=14 closed_switch=true missing_metadata=true missing_uncertainty=true missing_bridge=true unsupported_ledger=true missing_policies=true pmns_as_ray=true native_promotion=true native_registry=true pmns_native=true theorem_input=true charged_lepton=true neutrino=true pmns_target=true quarantined=true no_native_write=true no_pmns_ray=true verdict=CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_IMPORT_SWITCH_VALIDATED reason=the lepton airlock accepts only explicitly switched-on, fully metadated, bridge-only lepton/PMNS-residual rows into the comparator ledger and rejects every coordinate-smuggling or native-registry route

| case | accepted | verdict | reason |
|---|---:|---|---|
| valid charged-lepton comparator import | true | `CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_LEDGER` | lepton record imported into the quarantined comparator ledger only; native theorem registry remains untouched |
| valid neutrino comparator import | true | `CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_LEDGER` | lepton record imported into the quarantined comparator ledger only; native theorem registry remains untouched |
| valid PMNS residual-target import | true | `CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_LEDGER` | lepton record imported into the quarantined comparator ledger only; native theorem registry remains untouched |
| switch disabled | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_IMPORT_SWITCH_DISABLED` | empirical_import must be explicitly true before external lepton or PMNS data can enter the comparator ledger |
| missing source metadata | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_METADATA_INCOMPLETE` | lepton empirical records require source, scale, scheme, and uncertainty metadata |
| missing scale metadata | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_METADATA_INCOMPLETE` | lepton empirical records require source, scale, scheme, and uncertainty metadata |
| missing scheme metadata | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_METADATA_INCOMPLETE` | lepton empirical records require source, scale, scheme, and uncertainty metadata |
| missing uncertainty metadata | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_UNCERTAINTY_MISSING` | lepton empirical records must declare uncertainty before import |
| missing bridge-only quarantine | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_MISSING_BRIDGE_ONLY_QUARANTINE` | lepton empirical records must carry bridge_only=true quarantine metadata |
| unsupported target ledger | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_UNSUPPORTED_LEDGER_REJECTED` | lepton empirical data may target only the quarantined lepton-sector comparator ledger |
| missing neutrino ordering policy | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_NEUTRINO_POLICIES_MISSING` | lepton empirical records require neutrino ordering, absolute neutrino scale, Majorana/Dirac phase, and eigenbasis-convention policies |
| PMNS as lepton ray input | false | `FAILED_ROUTE_PMNS_USED_AS_EMPIRICAL_LEPTON_RAY_INPUT_REJECTED` | PMNS values may be residual targets, not alpha/phi/I_K ray-coordinate sources |
| native prediction promotion | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED` | lepton empirical records cannot be logged as native_prediction or native_law |
| native law promotion | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED` | lepton empirical records cannot be logged as native_prediction or native_law |
| native registry write | false | `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_NATIVE_REGISTRY_WRITE_REJECTED` | the lepton empirical airlock has no write path into the native theorem registry |
| PMNS native prediction | false | `FAILED_ROUTE_PMNS_NATIVE_PREDICTION_REJECTED` | PMNS values may be bridge residual targets only and cannot become native predictions |
| observed data as theorem input | false | `FAILED_ROUTE_LEPTON_OBSERVED_DATA_AS_THEOREM_INPUT_REJECTED` | observed lepton/PMNS data cannot be used as an input premise for native theorem verification |

## Firewall

executed=true airlock_open=true imported_rows=3 all_quarantined=true empirical_in_native=false native_prediction=false native_law=false theorem_input=false PMNS_native=false PMNS_constructed=false PMNS_entry=false lepton_mass_native=false neutrino_mass_native=false IK_native=false d_e_nu_observed=false d_e_nu_native=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_LEPTON_AIRLOCK_OPEN reason=even with empirical_import=true, lepton and PMNS residual rows are quarantined comparator inputs only; no observed lepton quantity can become native law or compute d_e_nu without a later rank-complete evaluator

## Status ledger

- `CONDITIONAL_SUPPORT_GATE476_PMNS_NULL_SOCKET_INHERITED`
- `CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_IMPORT_AIRLOCK_DEFINED`
- `CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_DATA_IMPORTED_TO_QUARANTINED_LEDGER`
- `CONDITIONAL_SUPPORT_LEPTON_EMPIRICAL_IMPORT_SWITCH_VALIDATED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_LEPTON_AIRLOCK_OPEN`
- `FAILED_ROUTE_LEPTON_EMPIRICAL_IMPORT_SWITCH_DISABLED`
- `FAILED_ROUTE_LEPTON_EMPIRICAL_METADATA_INCOMPLETE`
- `FAILED_ROUTE_LEPTON_EMPIRICAL_UNCERTAINTY_MISSING`
- `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_MISSING_BRIDGE_ONLY_QUARANTINE`
- `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_UNSUPPORTED_LEDGER_REJECTED`
- `FAILED_ROUTE_LEPTON_EMPIRICAL_NEUTRINO_POLICIES_MISSING`
- `FAILED_ROUTE_PMNS_USED_AS_EMPIRICAL_LEPTON_RAY_INPUT_REJECTED`
- `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_LEPTON_EMPIRICAL_DATA_NATIVE_REGISTRY_WRITE_REJECTED`
- `FAILED_ROUTE_PMNS_NATIVE_PREDICTION_REJECTED`
- `FAILED_ROUTE_LEPTON_OBSERVED_DATA_AS_THEOREM_INPUT_REJECTED`

## Next

Gate 478 — Observed lepton rank-complete ledger / PMNS non-computation audit: Gate477 opens the lepton empirical airlock but does not evaluate observed PMNS data or infer I_K. Primary task: ingest an explicit observed lepton ledger and fail closed unless e and nu sectors supply rank-complete I_spec/I_K/branch tags, neutrino policies, and bridge-only provenance

## Truth statement

Gate 477 validates the lepton-sector empirical airlock: empirical_import=true can admit fully metadated charged-lepton, neutrino, and PMNS residual-target rows into a quarantined bridge ledger, while PMNS-as-ray, native_prediction, native_law, theorem-input, and native-registry routes fail closed.
