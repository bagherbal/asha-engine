# Gate 468 Registry Audit — Common-Scale Synthetic Inversion Run / Uncertainty Propagation Harness

## Verdict

`CONDITIONAL_SUPPORT_COMMON_SCALE_SYNTHETIC_INVERSION_VALIDATED`

Gate 468 exercises the full ASHA cylinder socket on synthetic rank-complete u/d comparator ledgers. It computes bridge-only `(alpha,phi)` rays, propagates input uncertainty boxes, and returns a synthetic `d_ud` interval. It deliberately imports no observed PDG, CKM, PMNS, Yukawa, or mass values.

## Inheritance

executed=true K=true triangle=true inverse=true branch_tags=true socket=true airlock=true gate466_mass_only_obstruction=true gate467_ledger=true requires_I_spec_I_K=true requires_branch=true requires_uncertainty=true computable_if_numeric=true did_not_compute=true native_clean=true verdict=CONDITIONAL_SUPPORT_GATE467_COMMON_SCALE_LEDGER_INHERITED

## Formulae executed

```text
alpha = sqrt(3) I_K / sqrt(1-I_K^2)
cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)
phi = (sigma_CP arccos(cos(3phi)) + 2pi n_C3)/3
d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))
```

## Harness

executed=true accepted=1 rejected=9 valid_dud=true uncertainty=true observed_rejected=true missing_rank=true projective_rejected=true phase_rejected=true caustic_rejected=true branch_rejected=true uncertainty_missing=true cabibbo_as_ray=true native_promotion=true no_ckm_matrix=true no_ckm_entry=true no_native=true bridge_synthetic=true verdict=CONDITIONAL_SUPPORT_COMMON_SCALE_SYNTHETIC_INVERSION_VALIDATED reason=one complete synthetic u-d ledger computes bridge-only alpha/phi rays, propagates comparator uncertainties to a d_ud interval, and every observed/incomplete/unsafe/native-promotion route fails closed

| Case | Accepted | Verdict | `d_ud` / reason |
|---|---:|---|---|
| valid synthetic rank-complete u-d ledger | true | `CONDITIONAL_SUPPORT_SYNTHETIC_DUD_BRIDGE_ONLY_COMPUTED` | Delta_alpha=0.22 Delta_phi=0.33 d_ud=0.395367314 interval=[0.369461432,0.422279903] uncertainty=true bridge_only=true synthetic=true cabibbo_compared=false ckm_matrix=false ckm_entry=false native=false verdict=CONDITIONAL_SUPPORT_SYNTHETIC_DUD_BRIDGE_ONLY_COMPUTED |
| observed row rejected | false | `FAILED_ROUTE_OBSERVED_DATA_REJECTED_IN_SYNTHETIC_INVERSION` | one or both synthetic sector ledgers failed the rank-complete bridge-only preconditions |
| missing IK rejected | false | `FAILED_ROUTE_SYNTHETIC_INVERSION_REQUIRES_RANK_COMPLETE_LEDGER` | one or both synthetic sector ledgers failed the rank-complete bridge-only preconditions |
| projective domain rejected | false | `FAILED_ROUTE_SYNTHETIC_INVERSION_PROJECTIVE_DOMAIN_REJECTED` | Gate456 inverse rejected at least one sector |
| phase domain rejected | false | `FAILED_ROUTE_SYNTHETIC_INVERSION_PHASE_DOMAIN_REJECTED` | Gate456 inverse rejected at least one sector |
| caustic rejected | false | `FAILED_ROUTE_SYNTHETIC_INVERSION_CAUSTIC_REJECTED` | Gate456 inverse rejected at least one sector |
| branch tag rejected | false | `FAILED_ROUTE_SYNTHETIC_INVERSION_BRANCH_TAG_REJECTED` | one or both synthetic sector ledgers failed the rank-complete bridge-only preconditions |
| missing uncertainty rejected | false | `FAILED_ROUTE_SYNTHETIC_INVERSION_UNCERTAINTY_MISSING` | one or both synthetic sector ledgers failed the rank-complete bridge-only preconditions |
| Cabibbo as ray input rejected | false | `FAILED_ROUTE_CABIBBO_USED_AS_SYNTHETIC_RAY_INPUT_REJECTED` | one or both synthetic sector ledgers failed the rank-complete bridge-only preconditions |
| native promotion rejected | false | `FAILED_ROUTE_SYNTHETIC_DUD_NATIVE_PROMOTION_REJECTED` | one or both synthetic sector ledgers failed the rank-complete bridge-only preconditions |

## Accepted synthetic dry-run detail

- U ray: sector=u alpha=1 cos3phi=0.362357754 phi=0.4 alpha_interval=[0.989375699,1.01070964] phi_interval=[0.394608542,0.405267518] domain=true caustic=false bridge_only=true synthetic=true native_ray=false verdict=CONDITIONAL_SUPPORT_SYNTHETIC_RAY_INVERSION_EXECUTED
- D ray: sector=d alpha=1.22 cos3phi=-0.580386863 phi=0.73 alpha_interval=[1.20425296,1.23595187] phi_interval=[0.721291499,0.739122419] domain=true caustic=false bridge_only=true synthetic=true native_ray=false verdict=CONDITIONAL_SUPPORT_SYNTHETIC_RAY_INVERSION_EXECUTED
- Distance: Delta_alpha=0.22 Delta_phi=0.33 d_ud=0.395367314 interval=[0.369461432,0.422279903] uncertainty=true bridge_only=true synthetic=true cabibbo_compared=false ckm_matrix=false ckm_entry=false native=false verdict=CONDITIONAL_SUPPORT_SYNTHETIC_DUD_BRIDGE_ONLY_COMPUTED

## Native firewall proof

executed=true coords_native=false dud_native=false ckm_native=false ckm_matrix=false ckm_entry=false observed_masses=false observed_ckm=false cabibbo_as_ray=false native_write=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED reason=Gate468 computes only a synthetic bridge coordinate and interval; it writes no native ray, no CKM entry, and no theorem-registry observable

The computed synthetic `d_ud` is a bridge socket validation only. It is not `V_us`, not a CKM entry, not a CKM matrix, not a physical prediction, and not a native ASHA theorem. Observed values remain outside this gate.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE467_COMMON_SCALE_LEDGER_INHERITED`
- `CONDITIONAL_SUPPORT_SYNTHETIC_COMMON_SCALE_LEDGER_ACCEPTED`
- `CONDITIONAL_SUPPORT_SYNTHETIC_RAY_INVERSION_EXECUTED`
- `CONDITIONAL_SUPPORT_UNCERTAINTY_PROPAGATION_EXECUTED`
- `CONDITIONAL_SUPPORT_SYNTHETIC_DUD_BRIDGE_ONLY_COMPUTED`
- `CONDITIONAL_SUPPORT_COMMON_SCALE_SYNTHETIC_INVERSION_VALIDATED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED`
- `FAILED_ROUTE_OBSERVED_DATA_REJECTED_IN_SYNTHETIC_INVERSION`
- `FAILED_ROUTE_SYNTHETIC_INVERSION_REQUIRES_RANK_COMPLETE_LEDGER`
- `FAILED_ROUTE_SYNTHETIC_INVERSION_PROJECTIVE_DOMAIN_REJECTED`
- `FAILED_ROUTE_SYNTHETIC_INVERSION_PHASE_DOMAIN_REJECTED`
- `FAILED_ROUTE_SYNTHETIC_INVERSION_CAUSTIC_REJECTED`
- `FAILED_ROUTE_SYNTHETIC_INVERSION_BRANCH_TAG_REJECTED`
- `FAILED_ROUTE_SYNTHETIC_INVERSION_UNCERTAINTY_MISSING`
- `FAILED_ROUTE_CABIBBO_USED_AS_SYNTHETIC_RAY_INPUT_REJECTED`
- `FAILED_ROUTE_SYNTHETIC_DUD_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_SYNTHETIC_RESIDUAL_IS_NOT_CKM_PREDICTION`

## Next gate

Gate 469 — Observed Complete Comparator Dry-Run / Airlock Numerical Trial: Gate468 proves the socket works for rank-complete synthetic ledgers; the next empirical step may run only if a real common-scale u/d ledger supplies I_spec, I_K, branch tags, and uncertainties without using CKM as an input. Primary task: admit a fully provenanced observed comparator ledger through the Gate465 airlock, compute d_ud as bridge-only, and compare to Cabibbo solely as a residual target

## Truth statement

Gate 468 proves the Gate467 data product is sufficient in principle: a rank-complete synthetic u/d ledger can be inverted to bridge-only cylinder coordinates and propagated to a d_ud interval. This is a socket validation, not a CKM prediction; observed masses, CKM entries, and native coefficient values remain absent.
