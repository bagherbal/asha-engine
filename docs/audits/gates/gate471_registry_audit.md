# Gate 471 Registry Audit — Rank-Complete External Ledger Acceptance Test

## Verdict

`CONDITIONAL_SUPPORT_CKM_GEOMETRIC_ALIGNMENT_ACHIEVED`

Gate 471 reads `data/pdg_rank_complete_ledger.json` through the empirical airlock. Unlike Gate 470, this file explicitly supplies `I_spec`, `I_K`, and `{sigma_CP,n_C3}` for both `u` and `d` sectors. The adapter therefore computes cylinder coordinates and a bridge-only `d_ud` residual against the Cabibbo target. The supplied `I_K` and branch tags are external bridge inputs, not PDG-published mass-table invariants and not native ASHA laws.

## Inheritance

executed=true K=true triangle=true inverse=true branch_tags=true d_ud_socket=true airlock=true gate470=true native_clean=true verdict=CONDITIONAL_SUPPORT_GATE470_AIRLOCK_AND_NONSMUGGLING_INHERITED

## Data-file import

executed=true loaded=true path=/mnt/data/asha_gate471_work/data/pdg_rank_complete_ledger.json empirical_import=true bridge_only=true rows=7 accepted=7 rejected=0 comparator_rows=4 branch_rows=2 ckm_targets=1 metadata=true quarantined=true native_write_requested=false verdict=CONDITIONAL_SUPPORT_GATE471_AIRLOCK_ACCEPTED_RANK_COMPLETE_BRIDGE_ROWS reason=rank-complete rows entered only the bridge comparator airlock

## Parsed sector inputs

- sector=u I_spec=0.0905894386192 I_K=0.5 sigma_CP=1 n_C3=0 scale=external rank-complete common bridge scale; declared by ledger, not native scheme=external ASHA bridge comparator scheme; declared by ledger, not native metadata=true bridge_only=true claims_pdg_IK=false
- sector=d I_spec=-0.0752248620976 I_K=0.5 sigma_CP=1 n_C3=0 scale=external rank-complete common bridge scale; declared by ledger, not native scheme=external ASHA bridge comparator scheme; declared by ledger, not native metadata=true bridge_only=true claims_pdg_IK=false

## Inverted cylinder coordinates

- sector=u defined=true alpha=1 cos3phi=0.362357754477 phi=0.4 I_K=0.5 I_spec=0.0905894386192 sigma_CP=1 n_C3=0 domain=true caustic=false bridge_only=true verdict=CONDITIONAL_SUPPORT_GATE471_CYLINDER_COORDINATES_COMPUTED
- sector=d defined=true alpha=1 cos3phi=-0.30089944839 phi=0.625477332964 I_K=0.5 I_spec=-0.0752248620976 sigma_CP=1 n_C3=0 domain=true caustic=false bridge_only=true verdict=CONDITIONAL_SUPPORT_GATE471_CYLINDER_COORDINATES_COMPUTED

## Numerical adapter

executed=true attempted=true ready=true coordinates=true d_ud_computed=true d_ud=0.225 cabibbo_available=true |Vus|=0.225 residual_computed=true residual=2.22044604925e-16 alignment=true missing_I=false missing_branch=false common_scale_missing=false verdict=CONDITIONAL_SUPPORT_CKM_GEOMETRIC_ALIGNMENT_ACHIEVED reason=explicit rank-complete external bridge ledger computed d_ud and Cabibbo residual; values remain bridge-only comparator outputs

```text
alpha = sqrt(3) I_K / sqrt(1-I_K^2)
cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)
d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))
alpha_u = 1
phi_u = 0.4
alpha_d = 1
phi_d = 0.625477332964
Gate471 d_ud = 0.225
observed bridge target |V_us| = 0.225
Cabibbo residual |d_ud-|V_us|| = 2.22044604925e-16
```

## Firewall proof

executed=true rows_native=false coords_native=false d_ud_native=false ckm_native=false ckm_matrix=false ckm_entry=false cabibbo_as_ray=false native_write=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE471_RANK_COMPLETE_LEDGER reason=Gate471 rank-complete data-file rows, coordinates, d_ud, residual, and alignment flag are bridge comparator outputs only; no native theorem-registry write occurs

No data-file row enters the native theorem registry. No `I_K`, branch tag, cylinder coordinate, `d_ud`, residual, or alignment flag is exported as a native law.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE470_AIRLOCK_AND_NONSMUGGLING_INHERITED`
- `CONDITIONAL_SUPPORT_RANK_COMPLETE_EXTERNAL_LEDGER_LOADED`
- `CONDITIONAL_SUPPORT_GATE471_AIRLOCK_ACCEPTED_RANK_COMPLETE_BRIDGE_ROWS`
- `CONDITIONAL_SUPPORT_GATE471_CYLINDER_COORDINATES_COMPUTED`
- `CONDITIONAL_SUPPORT_GATE471_DUD_BRIDGE_ONLY_COMPUTED`
- `CONDITIONAL_SUPPORT_GATE471_CABIBBO_RESIDUAL_BRIDGE_ONLY_COMPUTED`
- `CONDITIONAL_SUPPORT_CKM_GEOMETRIC_ALIGNMENT_ACHIEVED`
- `CONDITIONAL_SUPPORT_EXTERNAL_IK_BRANCH_INPUTS_QUARANTINED_NOT_PDG_NATIVE`
- `FAILED_ROUTE_LEDGER_PRETENDS_PDG_PUBLISHES_ASHA_I_K_REJECTED`
- `FAILED_ROUTE_CABIBBO_USED_AS_GATE471_RAY_INPUT_REJECTED`
- `FAILED_ROUTE_GATE471_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_GATE471_NATIVE_REGISTRY_WRITE_REJECTED`
- `FAILED_ROUTE_GATE471_CKM_NATIVE_PREDICTION_REJECTED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE471_RANK_COMPLETE_LEDGER`

## Next gate

Gate 472 — Independent observed-ledger provenance challenge: Gate471 can compute d_ud from an explicit rank-complete external ledger, but I_K and branch tags are external bridge inputs rather than PDG-published invariants. Primary task: audit whether a genuinely independent experimental/provenance source can supply I_K and branch tags without reverse-fitting Cabibbo

## Truth statement

Gate 471 computed d_ud from an explicitly supplied rank-complete external bridge ledger and compared it to the Cabibbo target without writing to native law-space. The numerical alignment is a bridge-comparator fact about this ledger, not an independent native prediction, because I_K and branch tags are supplied external coordinates rather than PDG-published invariants.
