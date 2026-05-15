# Gate 470 Registry Audit — Observed Numerical d_ud Adapter / Explicit Data-File Run

## Verdict

`FAILED_ROUTE_OBSERVED_NUMERICAL_DUD_NOT_COMPUTABLE_FROM_FILE`

Gate 470 reads `data/pdg_observed_ledger.json` through the empirical airlock. The checked-in file contains PDG-style quark mass rows and a Cabibbo target row, but it does not contain explicit ASHA rank-complete bridge comparator values for `I_spec`, `I_K`, or branch tags. Therefore the adapter refuses to fabricate cylinder coordinates and leaves `d_ud` undefined.

## Inheritance

executed=true K=true triangle=true inverse=true branch_tags=true d_ud_socket=true airlock=true preflight=true native_clean=true verdict=CONDITIONAL_SUPPORT_GATE469_PREFLIGHT_INHERITED

## Data-file import

executed=true loaded=true path=/mnt/data/asha_work470/data/pdg_observed_ledger.json empirical_import=true bridge_only=true rows=13 accepted=13 rejected=0 mass_rows=6 comparator_rows=4 branch_rows=2 ckm_targets=1 metadata=true quarantined=true native_write_requested=false verdict=CONDITIONAL_SUPPORT_GATE470_AIRLOCK_ACCEPTED_QUARANTINED_ROWS reason=explicit ledger rows entered only the bridge comparator airlock

## Parsed sector inputs

- sector=u I_spec=missing I_K=missing sigma_CP=missing n_C3=missing scale=not supplied by PDG mass rows; required for numerical d_ud scheme=MS-bar where applicable; top mass convention requires explicit bridge normalization before use metadata=true bridge_only=true pdg_mass_rows=true pdg_no_IK=true
- sector=d I_spec=missing I_K=missing sigma_CP=missing n_C3=missing scale=not supplied by PDG mass rows; required for numerical d_ud scheme=MS-bar where applicable; top mass convention requires explicit bridge normalization before use metadata=true bridge_only=true pdg_mass_rows=true pdg_no_IK=true

## Numerical adapter

executed=true attempted=true ready=false d_ud_computed=false d_ud=undefined cabibbo_available=true |Vus|=0.225 residual_computed=false residual=undefined alignment=false missing_I=true missing_branch=true common_scale_missing=true pdg_no_IK=true verdict=FAILED_ROUTE_OBSERVED_NUMERICAL_DUD_NOT_COMPUTABLE_FROM_FILE reason=the explicit file was parsed, but it does not supply rank-complete ASHA bridge comparators; d_ud and Cabibbo residual remain undefined

```text
alpha = sqrt(3) I_K / sqrt(1-I_K^2)
cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)
d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))
Gate470 d_ud = undefined
observed bridge target |V_us| = 0.225
Cabibbo residual = undefined
```

## Firewall proof

executed=true rows_native=false coords_native=false d_ud_native=false ckm_native=false ckm_matrix=false ckm_entry=false cabibbo_as_ray=false native_write=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE470_DATA_FILE reason=Gate470 data-file rows are quarantined bridge comparators; no row, coordinate, d_ud, residual, CKM entry, or alignment flag writes to native law-space

No data-file row enters the native theorem registry. No mass, CKM value, `I_K`, `I_spec`, branch tag, cylinder coordinate, `d_ud`, residual, or alignment flag is exported as a native law.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE469_PREFLIGHT_INHERITED`
- `CONDITIONAL_SUPPORT_EXPLICIT_PDG_OBSERVED_LEDGER_LOADED`
- `CONDITIONAL_SUPPORT_GATE470_AIRLOCK_ACCEPTED_QUARANTINED_ROWS`
- `CONDITIONAL_SUPPORT_GATE470_NUMERICAL_ADAPTER_ATTEMPTED`
- `FAILED_ROUTE_GATE470_MISSING_EXPLICIT_I_SPEC_I_K_VALUES`
- `FAILED_ROUTE_GATE470_MISSING_EXPLICIT_BRANCH_TAGS`
- `FAILED_ROUTE_GATE470_COMMON_SCALE_SCHEME_NOT_SUPPLIED`
- `FAILED_ROUTE_PDG_MASS_LEDGER_DOES_NOT_SUPPLY_ASHA_I_K_INVARIANT`
- `FAILED_ROUTE_OBSERVED_NUMERICAL_DUD_NOT_COMPUTABLE_FROM_FILE`
- `FAILED_ROUTE_CABIBBO_RESIDUAL_UNDEFINED_WITHOUT_DUD`
- `FAILED_ROUTE_CABIBBO_USED_AS_GATE470_RAY_INPUT_REJECTED`
- `FAILED_ROUTE_GATE470_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_GATE470_NATIVE_REGISTRY_WRITE_REJECTED`
- `FAILED_ROUTE_GATE470_CKM_NATIVE_PREDICTION_REJECTED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE470_DATA_FILE`

## Next gate

Gate 471 — Rank-Complete External Ledger Acceptance Test: Gate470 parsed the explicit data file and refused to fabricate missing ASHA comparators; the next possible step is to supply a genuinely rank-complete external bridge ledger with I_spec, I_K, and branch tags. Primary task: evaluate a user-supplied rank-complete observed bridge ledger, never PDG masses alone, and export only bridge residuals

## Truth statement

Gate 470 successfully reads the explicit observed ledger through the empirical airlock, but the checked-in PDG-style file does not contain explicit ASHA rank-complete comparators. PDG mass rows and the Cabibbo target remain quarantined bridge data; d_ud and the Cabibbo residual are undefined until I_spec, I_K, and branch tags are explicitly supplied. The geometry has not been numerically matched to CKM by this run.
