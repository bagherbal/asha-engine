# Gate 469 Registry Audit — Observed Rank-Complete Comparator Preflight / Airlock Non-Computation Audit

## Verdict

`CONDITIONAL_SUPPORT_OBSERVED_RANK_COMPLETE_PREFLIGHT_VALIDATED`

Gate 469 validates the exact observed-data preflight required after Gate 468. It does not fetch, invent, or evaluate PDG values. It proves that observed ledgers must carry common scale/scheme, `I_spec`, `I_K`, complete branch tags, uncertainty metadata, and bridge-only quarantine before any future `d_ud` adapter may run.

## Inheritance

executed=true K=true triangle=true inverse=true branch=true dud_socket=true airlock=true schema=true synthetic_socket=true native_clean=true verdict=CONDITIONAL_SUPPORT_GATE468_SYNTHETIC_SOCKET_INHERITED

## Preflight summary

executed=true accepted_schema=1 rejected=9 ready_numeric=0 dud_computed=false switch_closed=true missing_I_spec=true missing_I_K=true missing_branch=true mixed_scale=true missing_uncertainty=true missing_numeric=true cabibbo_as_ray=true native_promotion=true ckm_native=true bridge_observed=true verdict=CONDITIONAL_SUPPORT_OBSERVED_RANK_COMPLETE_PREFLIGHT_VALIDATED reason=observed ledgers may pass only through the empirical airlock as bridge-only records; redacted ledgers do not compute d_ud

| Case | Accepted | Ready for numerical `d_ud` | Verdict | Reason |
|---|---:|---:|---|---|
| complete observed schema, redacted numeric values | true | false | `CONDITIONAL_SUPPORT_OBSERVED_RANK_COMPLETE_SCHEMA_ACCEPTED` | schema is complete but numeric I_spec/I_K values are redacted, so d_ud is not computed |
| switch closed rejected | false | false | `FAILED_ROUTE_OBSERVED_PREFLIGHT_EMPIRICAL_SWITCH_CLOSED` | observed preflight rejected one or both sector ledgers |
| missing I_spec rejected | false | false | `FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_I_SPEC` | observed preflight rejected one or both sector ledgers |
| missing I_K rejected | false | false | `FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_I_K` | observed preflight rejected one or both sector ledgers |
| missing branch tags rejected | false | false | `FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_BRANCH_TAGS` | observed preflight rejected one or both sector ledgers |
| mixed scale rejected | false | false | `FAILED_ROUTE_OBSERVED_PREFLIGHT_MIXED_SCALE_SCHEME_REJECTED` | observed preflight rejected one or both sector ledgers |
| missing uncertainty rejected | false | false | `FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_UNCERTAINTY` | observed preflight rejected one or both sector ledgers |
| Cabibbo as ray input rejected | false | false | `FAILED_ROUTE_CABIBBO_USED_AS_OBSERVED_RAY_INPUT_REJECTED` | observed preflight rejected one or both sector ledgers |
| native promotion rejected | false | false | `FAILED_ROUTE_OBSERVED_PREFLIGHT_NATIVE_PROMOTION_REJECTED` | observed preflight rejected one or both sector ledgers |
| CKM native prediction rejected | false | false | `FAILED_ROUTE_OBSERVED_PREFLIGHT_CKM_NATIVE_PREDICTION_REJECTED` | observed preflight rejected one or both sector ledgers |

## Formula socket held in reserve

```text
alpha = sqrt(3) I_K / sqrt(1-I_K^2)
cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)
phi = (sigma_CP arccos(cos(3phi)) + 2pi n_C3)/3
d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))
```

Gate 469 does not evaluate this socket for redacted observed rows. A numeric-ready row is only declared eligible for a later bridge-only adapter; it is not promoted to native law and it is not interpreted as `V_us`.

## Native firewall proof

executed=true observed_native=false dud_native=false ckm_native=false ckm_matrix=false ckm_entry=false cabibbo_as_ray=false native_write=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_OBSERVED_PREFLIGHT reason=Gate469 performs observed-ledger preflight only; no observed comparator is written to native registry and no CKM entry is constructed

No observed comparator row can become a native prediction, native law, CKM matrix element, PMNS value, Yukawa value, or coefficient selector. Cabibbo remains a residual target only, never a coordinate input.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE468_SYNTHETIC_SOCKET_INHERITED`
- `CONDITIONAL_SUPPORT_OBSERVED_PREFLIGHT_POLICY_DEFINED`
- `CONDITIONAL_SUPPORT_OBSERVED_RANK_COMPLETE_SCHEMA_ACCEPTED`
- `CONDITIONAL_SUPPORT_OBSERVED_RANK_COMPLETE_PREFLIGHT_VALIDATED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_OBSERVED_PREFLIGHT`
- `FAILED_ROUTE_OBSERVED_PREFLIGHT_EMPIRICAL_SWITCH_CLOSED`
- `FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_I_SPEC`
- `FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_I_K`
- `FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_BRANCH_TAGS`
- `FAILED_ROUTE_OBSERVED_PREFLIGHT_MIXED_SCALE_SCHEME_REJECTED`
- `FAILED_ROUTE_OBSERVED_PREFLIGHT_MISSING_UNCERTAINTY`
- `FAILED_ROUTE_OBSERVED_DUD_NOT_COMPUTED_WITHOUT_ACTUAL_I_SPEC_I_K_VALUES`
- `FAILED_ROUTE_CABIBBO_USED_AS_OBSERVED_RAY_INPUT_REJECTED`
- `FAILED_ROUTE_OBSERVED_PREFLIGHT_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_OBSERVED_PREFLIGHT_CKM_NATIVE_PREDICTION_REJECTED`

## Next gate

Gate 470 — Observed Numerical d_ud Adapter / Explicit Data-File Run: Gate469 validates the observed schema and proves redacted rows cannot compute d_ud; a later run may evaluate only with explicit rank-complete I_spec/I_K values and branch tags. Primary task: read a user-supplied observed comparator ledger, evaluate d_ud as bridge-only, and compare to Cabibbo only as a residual target

## Truth statement

Gate 469 preflight failed.
