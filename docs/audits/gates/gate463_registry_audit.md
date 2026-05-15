# Gate 463 Registry Audit — Eigenbasis Convention Ledger / Mixing-Matrix Gauge Audit

## Verdict

`CONDITIONAL_SUPPORT_CKM_NULL_ADAPTER_PRECONDITION_SET`

Gate 463 defines the bridge-only eigenbasis convention ledger required before a u-d relative ray can enter any future CKM residual adapter. It computes no CKM or PMNS entries and imports no observed flavor data.

## Inheritance

executed=true K=true triangle=true inverse=true branch_tags=true multiplex=true sector_difference=true requires_eigenbasis=true rejects_observed=true rejects_native=true no_mixing_export=true no_observed=true verdict=CONDITIONAL_SUPPORT_GATE462_SECTOR_DIFFERENCE_INTERFACE_INHERITED

## Diagonalizer gauge audit

executed=true phase_gauge_per_sector=U(1)^3 permutation_sheets_per_sector=6 pair_phase_dim=6 pair_permutation_sheets=36 K_simple=true centralizer=centralizer_U(3)(K_gen)=U(1)^3 rephasings_only=true raw_not_observable=true permutations_not_native=true convention_bridge_fix=true convention_no_prediction=true verdict=CONDITIONAL_SUPPORT_DIAGONALIZER_GAUGE_AUDIT_COMPLETE reason=raw sector diagonalizers carry U(1)^3 phase gauge and S3 ordering ambiguity; a convention can fix bridge bookkeeping but cannot derive CKM/PMNS entries

```text
raw sector diagonalizer gauge: U(1)^3 x S3
u-d pair gauge before convention: (U(1)^3 x S3)_u x (U(1)^3 x S3)_d
pair phase-gauge dimension: 6
pair permutation sheets: 36
K_gen-preserving basis group: centralizer_U(3)(K_gen)=U(1)^3
```

## Required convention ledger

executed=true sectors=u,d K_basis=true ordering=true phase_gauge=true normalization=true degeneracy_policy=true branch_tag=true provenance=true bridge_only=true reject_raw=true reject_native_permutation=true reject_K_rotation=true reject_observed=true reject_native_mixing=true readiness_only=true verdict=CONDITIONAL_SUPPORT_EIGENBASIS_CONVENTION_SLOT_DEFINED reason=a future CKM residual adapter may receive only a bridge-only convention-ready u-d sector pair; no matrix element is exported here

```text
sector in {u,d}
K_gen basis declared
eigenvalue ordering declared as bridge convention, not native particle label
eigenvector phase gauge declared
unit normalization declared
degeneracy policy = fail closed
branch tag and provenance inherited from comparator ledger
bridge_only = true
```

## Sieve

executed=true accepted=1 rejected=10 valid=true missing_sector=true missing_convention=true raw_phase=true permutation=true degenerate=true K_rotation=true observed=true native_prediction=true native_eigenbasis=true matrix_export=true bridge_only=true no_matrix=true verdict=CONDITIONAL_SUPPORT_CKM_NULL_ADAPTER_PRECONDITION_SET reason=only the complete bridge-only u-d eigenbasis convention ledger is accepted; it exports readiness, not a CKM/PMNS matrix

| Case | Accepted | Verdict | Convention result | Reason |
|---|---|---|---|---|
| valid synthetic u-d convention ledger | true | `CONDITIONAL_SUPPORT_EIGENBASIS_CONVENTION_LEDGER_VALIDATED` | u->d complete=true bridge_only=true ready=true CKM_computed=false PMNS_computed=false native_export=false phase_fixed=true permutation_fixed=true degeneracy_rejected=true K_preserved=true verdict=CONDITIONAL_SUPPORT_EIGENBASIS_CONVENTION_LEDGER_VALIDATED | complete u-d eigenbasis conventions are present; the result is readiness for a later bridge residual adapter, not a CKM prediction |
| missing d sector convention | false | `FAILED_ROUTE_EIGENBASIS_REQUIRES_U_D_CONVENTIONS` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_EIGENBASIS_REQUIRES_U_D_CONVENTIONS | u and d sector conventions are both required and must be bridge-only |
| missing eigenvalue ordering | false | `FAILED_ROUTE_EIGENBASIS_CONVENTION_MISSING` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_EIGENBASIS_CONVENTION_MISSING | each sector must declare K_gen basis, ordering, phase gauge, normalization, degeneracy policy, branch tag, and provenance |
| raw diagonalizer phase gauge | false | `FAILED_ROUTE_RAW_DIAGONALIZERS_HAVE_PHASE_GAUGE` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_RAW_DIAGONALIZERS_HAVE_PHASE_GAUGE | raw diagonalizers retain U(1)^3 phase gauge per sector and cannot be used as observables |
| native eigenvalue permutation claim | false | `FAILED_ROUTE_EIGENVALUE_PERMUTATION_NOT_NATIVE` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_EIGENVALUE_PERMUTATION_NOT_NATIVE | eigenvalue ordering is a bridge convention; native geometry does not label mass-generation permutations |
| degenerate spectrum | false | `FAILED_ROUTE_DEGENERATE_SPECTRUM_REJECTED` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_DEGENERATE_SPECTRUM_REJECTED | degenerate spectra make eigenvectors non-unique, so the convention ledger must fail closed |
| K_gen basis rotation requested | false | `FAILED_ROUTE_KGEN_BASIS_ROTATION_REJECTED` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_KGEN_BASIS_ROTATION_REJECTED | a general family rotation would erase the native K_gen address and is not a convention gauge |
| observed CKM/PMNS import | false | `FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED | observed CKM/PMNS values are not accepted in the eigenbasis convention audit |
| native CKM/PMNS prediction claim | false | `FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED | a convention ledger cannot promote CKM/PMNS entries into native ASHA predictions |
| native eigenbasis promotion | false | `FAILED_ROUTE_EIGENBASIS_NATIVE_PROMOTION_REJECTED` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_EIGENBASIS_NATIVE_PROMOTION_REJECTED | sector eigenvectors are bridge gauge choices and cannot become native law-space |
| convention ledger tries to export CKM matrix | false | `FAILED_ROUTE_CONVENTION_LEDGER_CANNOT_EXPORT_CKM_MATRIX` | u->d complete=false bridge_only=false ready=false CKM_computed=false PMNS_computed=false native_export=false phase_fixed=false permutation_fixed=false degeneracy_rejected=false K_preserved=false verdict=FAILED_ROUTE_CONVENTION_LEDGER_CANNOT_EXPORT_CKM_MATRIX | Gate463 exports convention readiness only; CKM/PMNS matrix construction belongs to a later explicit bridge adapter |

## What Gate 463 does not do

Gate 463 does not diagonalize physical quark matrices, does not compute CKM entries, does not compute PMNS entries, does not choose particle labels natively, and does not turn eigenvectors into ASHA law. It only proves that a future residual adapter must carry an explicit bridge convention before matrix-like comparisons are meaningful.

## Result statuses

- `CONDITIONAL_SUPPORT_GATE462_SECTOR_DIFFERENCE_INTERFACE_INHERITED`
- `CONDITIONAL_SUPPORT_DIAGONALIZER_GAUGE_AUDIT_COMPLETE`
- `CONDITIONAL_SUPPORT_EIGENBASIS_CONVENTION_SLOT_DEFINED`
- `CONDITIONAL_SUPPORT_EIGENBASIS_CONVENTION_LEDGER_VALIDATED`
- `CONDITIONAL_SUPPORT_CKM_NULL_ADAPTER_PRECONDITION_SET`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED`
- `FAILED_ROUTE_EIGENBASIS_REQUIRES_U_D_CONVENTIONS`
- `FAILED_ROUTE_RAW_DIAGONALIZERS_HAVE_PHASE_GAUGE`
- `FAILED_ROUTE_EIGENVALUE_PERMUTATION_NOT_NATIVE`
- `FAILED_ROUTE_DEGENERATE_SPECTRUM_REJECTED`
- `FAILED_ROUTE_KGEN_BASIS_ROTATION_REJECTED`
- `FAILED_ROUTE_EIGENBASIS_CONVENTION_MISSING`
- `FAILED_ROUTE_OBSERVED_CKM_PMNS_IMPORT_REJECTED`
- `FAILED_ROUTE_CKM_PMNS_NATIVE_PREDICTION_REJECTED`
- `FAILED_ROUTE_EIGENBASIS_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_CONVENTION_LEDGER_CANNOT_EXPORT_CKM_MATRIX`

## Firewall

executed=true convention_defined=true null_adapter_may_proceed=true CKM_computed=false CKM_native=false PMNS_computed=false PMNS_native=false masses_imported=false yukawas_imported=false CKM_imported=false PMNS_imported=false GST_promoted=false K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED reason=Gate463 defines only a bridge eigenbasis convention ledger; CKM/PMNS values, masses, Yukawas, branch choices, and coefficients remain quarantined.

## Next gate

Gate 464 — CKM Null Residual Adapter / Convention-Ready Symbolic Map: Gate463 now supplies the bridge-only eigenbasis convention slot needed before relative u-d rays can be compared by any CKM-facing residual harness. Primary task: compose Gate462 relative-ray diagnostics with Gate463 eigenbasis conventions into a symbolic/null CKM residual adapter that still rejects observed CKM data and native-prediction claims

## Truth statement

Gate 463 proves that a CKM-facing bridge adapter needs an explicit u-d eigenbasis convention ledger. The ledger can fix bridge bookkeeping gauge, but it cannot create CKM/PMNS predictions or promote eigenvectors into native law-space.
