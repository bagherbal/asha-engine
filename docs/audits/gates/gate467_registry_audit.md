# Gate 467 Registry Audit — Common-Scale Running Ledger / Coefficient-Ray Comparator Design

## Verdict

`CONDITIONAL_SUPPORT_COMMON_SCALE_COMPARATOR_DESIGN_BRIDGE_ONLY_VALIDATED`

Gate 467 does not compute `d_ud`. It defines the exact bridge-only input ledger that was missing in Gate 466. A future observed run must supply common-scale/common-scheme u and d sector spectra, the independent `I_K` comparator, complete `{sigma_CP,n_C3}` branch tags, provenance, and uncertainty propagation before the ASHA cylinder socket is allowed to evaluate coordinates.

## Inheritance

executed=true K=true triangle=true spectrum_rank_one=true inverse_two_scalars=true branch_tags=true ckm_socket=true airlock=true gate466_rows=true gate466_dud_undefined=true gate466_not_computable=true native_clean=true verdict=CONDITIONAL_SUPPORT_GATE466_OBSERVED_ADAPTER_OBSTRUCTION_INHERITED

## Required protocol

executed=true sectors=u,d common_scale=true common_scheme=true three_masses=true trace_zero=true I_spec=true I_K=true branch_tags=true metadata=true uncertainty=true dimensionless=true bridge_only=true reject_cabibbo_as_ray=true reject_native=true ray_dof=2 spectrum_rank=1 ray_scalars=2 branch_fields=2 schema_fields=12 verdict=CONDITIONAL_SUPPORT_COMMON_SCALE_COEFFICIENT_RAY_PROTOCOL_DEFINED reason=a future observed CKM bridge run needs common-scale sector spectra plus I_spec, I_K, and {sigma_CP,n_C3}; Cabibbo data may be a target residual only, never a ray-definition input

```text
I_spec = 2 cos(3 phi)/(alpha^2+3)^(3/2)
I_K    = alpha/sqrt(alpha^2+3)
branch = {sigma_CP, n_C3}
d_ud   = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))
```

## Schema audit

executed=true accepted=2 rejected=7 u=true d=true mixed_scale_rejected=true missing_IK=true missing_branch=true missing_uncertainty=true mass_only=true cabibbo_as_ray=true native_promotion=true both_ready=true dud_computable_if_numeric=true dud_computed_now=false verdict=CONDITIONAL_SUPPORT_COMMON_SCALE_COMPARATOR_DESIGN_BRIDGE_ONLY_VALIDATED reason=the schema accepts one complete u ledger and one complete d ledger, rejects every incomplete or unsafe variant, and still does not compute d_ud because Gate467 is a design/provenance contract, not an observed-value run

| Sector / probe | Accepted | Verdict | Ledger |
|---|---:|---|---|
| `u` | true | `CONDITIONAL_SUPPORT_COMMON_SCALE_RAY_LEDGER_SCHEMA_ACCEPTED` | sector=u masses=m_u(mu_star),m_c(mu_star),m_t(mu_star) scale=mu_star_common_bridge_scale scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty=covariance-or-interval-propagation-declared common_running=true trace_zero=true I_spec=true I_K=true cp_sign=true c3_sheet=true dimensionless=true bridge_only=true cabibbo_as_ray=false native_claim=false |
| `d` | true | `CONDITIONAL_SUPPORT_COMMON_SCALE_RAY_LEDGER_SCHEMA_ACCEPTED` | sector=d masses=m_d(mu_star),m_s(mu_star),m_b(mu_star) scale=mu_star_common_bridge_scale scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty=covariance-or-interval-propagation-declared common_running=true trace_zero=true I_spec=true I_K=true cp_sign=true c3_sheet=true dimensionless=true bridge_only=true cabibbo_as_ray=false native_claim=false |
| `u` | false | `FAILED_ROUTE_MIXED_SCALE_RUNNING_LEDGER_REJECTED` | sector=u masses=m_u(mu_star),m_c(mu_star),m_t(mu_star) scale=mixed: 2 GeV, mu=m_c, mu=m_t scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty=covariance-or-interval-propagation-declared common_running=false trace_zero=true I_spec=true I_K=true cp_sign=true c3_sheet=true dimensionless=true bridge_only=true cabibbo_as_ray=false native_claim=false |
| `u` | false | `FAILED_ROUTE_MISSING_IK_COMPARATOR_REJECTED` | sector=u masses=m_u(mu_star),m_c(mu_star),m_t(mu_star) scale=mu_star_common_bridge_scale scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty=covariance-or-interval-propagation-declared common_running=true trace_zero=true I_spec=true I_K=false cp_sign=true c3_sheet=true dimensionless=true bridge_only=true cabibbo_as_ray=false native_claim=false |
| `d` | false | `FAILED_ROUTE_MISSING_BRANCH_TAGS_REJECTED` | sector=d masses=m_u(mu_star),m_c(mu_star),m_t(mu_star) scale=mu_star_common_bridge_scale scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty=covariance-or-interval-propagation-declared common_running=true trace_zero=true I_spec=true I_K=true cp_sign=false c3_sheet=false dimensionless=true bridge_only=true cabibbo_as_ray=false native_claim=false |
| `d` | false | `FAILED_ROUTE_UNCERTAINTY_PROPAGATION_MISSING` | sector=d masses=m_u(mu_star),m_c(mu_star),m_t(mu_star) scale=mu_star_common_bridge_scale scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty= common_running=true trace_zero=true I_spec=true I_K=true cp_sign=true c3_sheet=true dimensionless=true bridge_only=true cabibbo_as_ray=false native_claim=false |
| `u` | false | `FAILED_ROUTE_MISSING_IK_COMPARATOR_REJECTED` | sector=u masses=m_u(mu_star),m_c(mu_star),m_t(mu_star) scale=mu_star_common_bridge_scale scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty=covariance-or-interval-propagation-declared common_running=true trace_zero=true I_spec=true I_K=false cp_sign=false c3_sheet=false dimensionless=true bridge_only=true cabibbo_as_ray=false native_claim=false |
| `u-d` | false | `FAILED_ROUTE_CABIBBO_USED_AS_RAY_INPUT_REJECTED` | sector=u-d masses=\|V_us\| scale=mu_star_common_bridge_scale scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty=covariance-or-interval-propagation-declared common_running=true trace_zero=true I_spec=true I_K=true cp_sign=true c3_sheet=true dimensionless=true bridge_only=true cabibbo_as_ray=true native_claim=false |
| `u` | false | `FAILED_ROUTE_COMMON_SCALE_LEDGER_NATIVE_PROMOTION_REJECTED` | sector=u masses=m_u(mu_star),m_c(mu_star),m_t(mu_star) scale=mu_star_common_bridge_scale scheme=MS-bar-or-explicit-running-scheme source=bridge-running-ledger-source version=versioned external running calculation uncertainty=covariance-or-interval-propagation-declared common_running=true trace_zero=true I_spec=true I_K=true cp_sign=true c3_sheet=true dimensionless=true bridge_only=true cabibbo_as_ray=false native_claim=true |

The two accepted ledgers are schemas only: they prove that a complete bridge data product can be represented. They do not contain numerical coordinates and do not evaluate `d_ud`. Every mass-only, mixed-scale, missing-`I_K`, missing-branch, missing-uncertainty, Cabibbo-as-coordinate, or native-promotion attempt fails closed.

## Native firewall proof

executed=true protocol_native=false coords_native=false dud_native=false ckm_native=false cabibbo_as_ray=false masses_theorem_input=false native_write=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED reason=Gate467 defines the missing empirical bridge schema only; it does not turn common-scale ledgers, I_K, branch tags, d_ud, or CKM residuals into native ASHA law-space

The common-scale ledger is an empirical bridge contract, not a native theorem. `K_gen` and `X_triangle` remain structural; `Y_phase`, sector coefficients, branch tags, and any future coordinates remain quarantined. Cabibbo/CKM data may only be a residual target, never an input used to define the ray.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE466_OBSERVED_ADAPTER_OBSTRUCTION_INHERITED`
- `CONDITIONAL_SUPPORT_COMMON_SCALE_COEFFICIENT_RAY_PROTOCOL_DEFINED`
- `CONDITIONAL_SUPPORT_COMMON_SCALE_RAY_LEDGER_SCHEMA_ACCEPTED`
- `CONDITIONAL_SUPPORT_RANK_COMPLETE_COMPARATOR_REQUIREMENTS_VALIDATED`
- `CONDITIONAL_SUPPORT_COMMON_SCALE_COMPARATOR_DESIGN_BRIDGE_ONLY_VALIDATED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED`
- `FAILED_ROUTE_MIXED_SCALE_RUNNING_LEDGER_REJECTED`
- `FAILED_ROUTE_MISSING_IK_COMPARATOR_REJECTED`
- `FAILED_ROUTE_MISSING_BRANCH_TAGS_REJECTED`
- `FAILED_ROUTE_UNCERTAINTY_PROPAGATION_MISSING`
- `FAILED_ROUTE_MASS_SPECTRA_ONLY_STILL_RANK_ONE`
- `FAILED_ROUTE_CABIBBO_USED_AS_RAY_INPUT_REJECTED`
- `FAILED_ROUTE_COMMON_SCALE_LEDGER_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_OBSERVED_CKM_NATIVE_PREDICTION_REJECTED`

## Next gate

Gate 468 — Common-Scale Synthetic Inversion Run / Uncertainty Propagation Harness: Gate467 defines the complete u/d comparator data product; the next safe step is to exercise the full inverse and d_ud formula on synthetic or redacted complete ledgers with interval propagation. Primary task: apply the Gate456 inverse and Gate464 d_ud socket to rank-complete synthetic common-scale u/d records, while keeping all numerical coordinates bridge-only

## Truth statement

Gate 467 defines the exact bridge-only data product missing from Gate 466: common-scale u/d running spectra, I_spec, I_K, complete branch tags, provenance, and uncertainty propagation. It proves that a future d_ud calculation is allowed only after this rank-complete ledger exists; Cabibbo data remains a target residual, not an input coordinate or native prediction.
