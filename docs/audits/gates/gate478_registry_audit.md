# Gate 478 Registry Audit — Lepton-Sector Observed Comparator Adapter / PMNS Airlock Non-Computation Audit

## Verdict

`FAILED_ROUTE_OBSERVED_LEPTON_DENU_NOT_COMPUTABLE_FROM_FILE`

Gate 478 reads `data/lepton_observed_ledger.json` through the lepton empirical airlock. The checked-in file contains charged-lepton, neutrino, and PMNS residual-target rows, but it does not contain explicit ASHA rank-complete bridge comparator values for `I_spec`, `I_K`, or branch tags. Therefore the adapter refuses to fabricate e/nu cylinder coordinates and leaves `d_eν` undefined.

## Inheritance

executed=true K=true triangle=true inverse=true branch_tags=true d_eν_socket=true lepton_airlock=true preflight=true native_clean=true verdict=CONDITIONAL_SUPPORT_GATE477_LEPTON_AIRLOCK_INHERITED

## Data-file import

executed=true loaded=true path=/mnt/data/asha_work_478/data/lepton_observed_ledger.json empirical_import=true bridge_only=true rows=6 accepted=6 rejected=0 charged_lepton_rows=3 neutrino_rows=2 comparator_rows=0 branch_rows=0 pmns_targets=1 metadata=true policies=true quarantined=true native_write_requested=false verdict=CONDITIONAL_SUPPORT_GATE478_AIRLOCK_ACCEPTED_QUARANTINED_LEPTON_ROWS reason=explicit lepton observed ledger rows entered only the bridge comparator airlock

## Parsed sector inputs

- sector=e I_spec=missing I_K=missing sigma_CP=missing n_C3=missing scale=declared lepton comparison convention; not a native ASHA scale scheme=declared bridge-only lepton/PMNS convention metadata=true policies=true bridge_only=true observed_rows=true pmns_rows=false lepton_no_IK=true
- sector=nu I_spec=missing I_K=missing sigma_CP=missing n_C3=missing scale=declared lepton comparison convention; not a native ASHA scale scheme=declared bridge-only lepton/PMNS convention metadata=true policies=true bridge_only=true observed_rows=true pmns_rows=false lepton_no_IK=true

## Observed lepton adapter

executed=true attempted=true ready=false d_eν_computed=false d_eν=undefined pmns_target_available=true target=0.707106781187 residual_computed=false residual=undefined alignment=false missing_I=true missing_branch=true missing_policies=false lepton_no_IK=true verdict=FAILED_ROUTE_OBSERVED_LEPTON_DENU_NOT_COMPUTABLE_FROM_FILE reason=the explicit lepton file was parsed, but it does not supply rank-complete ASHA bridge comparators; d_eν and PMNS residual remain undefined

```text
alpha = sqrt(3) I_K / sqrt(1-I_K^2)
cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)
d_eν = sqrt((alpha_ν-alpha_e)^2 + 4 sin^2((phi_ν-phi_e)/2))
Gate478 d_eν = undefined
observed bridge target θ23-like PMNS row = 0.707106781187
PMNS residual = undefined
```

## Firewall proof

executed=true rows_native=false coords_native=false d_eν_native=false pmns_native=false pmns_matrix=false pmns_entry=false pmns_as_ray=false native_write=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE478_LEPTON_DATA_FILE reason=Gate478 lepton data-file rows are quarantined bridge comparators; no row, coordinate, d_eν, residual, PMNS entry, or alignment flag writes to native law-space

No data-file row enters the native theorem registry. No lepton mass, neutrino row, PMNS value, `I_K`, `I_spec`, branch tag, cylinder coordinate, `d_eν`, residual, matrix entry, or alignment flag is exported as a native law.

## Structural equivalence to quark socket

Gate 478 uses the same cylinder metric as Gate 470/Gate 464, with the sector labels changed from `u,d` to `e,ν`. The only difference is the lepton preflight policy: neutrino ordering, absolute-scale, Majorana/Dirac phase, and eigenbasis conventions are mandatory metadata.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE477_LEPTON_AIRLOCK_INHERITED`
- `CONDITIONAL_SUPPORT_EXPLICIT_LEPTON_OBSERVED_LEDGER_LOADED`
- `CONDITIONAL_SUPPORT_GATE478_AIRLOCK_ACCEPTED_QUARANTINED_LEPTON_ROWS`
- `CONDITIONAL_SUPPORT_GATE478_OBSERVED_LEPTON_ADAPTER_ATTEMPTED`
- `FAILED_ROUTE_GATE478_MISSING_EXPLICIT_LEPTON_I_SPEC_I_K_VALUES`
- `FAILED_ROUTE_GATE478_MISSING_EXPLICIT_LEPTON_BRANCH_TAGS`
- `FAILED_ROUTE_LEPTON_MASS_PMNS_LEDGER_DOES_NOT_SUPPLY_ASHA_I_K_INVARIANT`
- `FAILED_ROUTE_OBSERVED_LEPTON_DENU_NOT_COMPUTABLE_FROM_FILE`
- `FAILED_ROUTE_PMNS_RESIDUAL_UNDEFINED_WITHOUT_DENU`
- `FAILED_ROUTE_PMNS_USED_AS_GATE478_LEPTON_RAY_INPUT_REJECTED`
- `FAILED_ROUTE_GATE478_LEPTON_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_GATE478_LEPTON_NATIVE_REGISTRY_WRITE_REJECTED`
- `FAILED_ROUTE_GATE478_PMNS_NATIVE_PREDICTION_REJECTED`
- `FAILED_ROUTE_GATE478_PMNS_MATRIX_EXPORT_REJECTED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE478_LEPTON_DATA_FILE`

## Next gate

Gate 479 — Lepton rank-complete external ledger acceptance test: Gate478 parsed the explicit lepton observed file and refused to fabricate ASHA comparators from masses or PMNS targets. Primary task: evaluate only a user-supplied rank-complete e/nu bridge ledger with I_spec, I_K, and branch tags; export bridge residuals only

## Truth statement

Gate 478 successfully reads the explicit observed lepton ledger through the empirical airlock, but the checked-in lepton/PMNS-style file does not contain explicit ASHA rank-complete comparators. Charged-lepton rows, neutrino rows, and PMNS target rows remain quarantined bridge data; d_eν and the PMNS residual are undefined until I_spec, I_K, and branch tags are explicitly supplied. The lepton socket is structurally identical to the quark socket and fails closed in the same way.
