# Gate 459 Registry Audit — Oriented Comparator Branch Tag Sieve / CP-Sign Ledger

## Verdict

`CONDITIONAL_SUPPORT_ORIENTED_BRANCH_TAG_BRIDGE_ONLY_VALIDATED`

Gate 459 formalizes the orientation metadata required after the Gate 456 inverse and Gate 458 redacted evaluator. It proves that a cosine invariant alone gives six branches, a CP-odd sign gives three, and a complete bridge tag `{sigma_CP,n_C3}` gives one synthetic phase branch. None of these tags is native ASHA law.

## Inheritance

executed=true K=true triangle=true inverse=true branches=6 provenance=true harness=true observed_rejected=true bridge_only=true native_CP_selector_absent=true native_C3_selector_absent=true no_observed=true verdict=CONDITIONAL_SUPPORT_GATE458_REDACTED_HARNESS_INHERITED

## Branch ledger

executed=true cosine=true cp_sign=true c3_sheet=true cosine_branches=6 cp_sign_branches=3 complete_branches=1 bridge_only=true reject_CKM_PMNS=true reject_native=true native_CP_absent=true native_C3_absent=true verdict=CONDITIONAL_SUPPORT_ORIENTED_BRANCH_TAG_LEDGER_DEFINED reason=A cosine invariant fixes only cos(3phi). A CP-odd sign chooses one orientation of 3phi but leaves three cubic sheets. A complete bridge-only tag {sigma_CP,n_C3} is necessary and sufficient for a unique synthetic phase branch.

```text
C = cos(3 phi)
cosine-only branches: phi = (± arccos(C) + 2*pi*n)/3, n=0,1,2
sigma_CP = sign(sin(3 phi)) chooses the ± orientation
n_C3 in {0,1,2} chooses the residual cubic sheet
complete tag: phi = (sigma_CP arccos(C) + 2*pi*n_C3)/3
```

## Branch-tag sieve

executed=true accepted=2 rejected=5 cosine_only=true cp_only=true complete_pos=true complete_neg=true CKM_PMNS=true native_promotion=true invalid_tag=true all_bridge=true no_native_phase=true verdict=CONDITIONAL_SUPPORT_ORIENTED_BRANCH_TAG_BRIDGE_ONLY_VALIDATED reason=2 complete bridge tags selected unique synthetic branches; 5 incomplete, empirical-selector, or native-promotion routes failed closed.

| Record | C=cos(3phi) | sigma_CP | n_C3 | Accepted | Selected | Phase | Branches | Bridge-only | Verdict | Reason |
|---|---:|---:|---:|---|---|---:|---:|---|---|---|
| cosine-only Gate456 inverse result | 0.25 | ∅ | ∅ | false | false | 0 | 6 | true | `FAILED_ROUTE_COSINE_INVARIANT_RETURNS_SIX_BRANCHES` | cos(3phi) alone leaves the six Gate456 branches phi=(±arccos(C)+2πn)/3. |
| CP-odd sign without C3 sheet | 0.25 | 1 | ∅ | false | false | 0 | 3 | true | `FAILED_ROUTE_CP_ODD_SIGN_ONLY_LEAVES_C3_SHEETS` | the CP-odd sign chooses ±arccos(C) but still leaves the three C3 sheets n=0,1,2. |
| complete positive bridge branch tag | 0.25 | 1 | 2 | true | true | 4.6281622 | 1 | true | `CONDITIONAL_SUPPORT_COMPLETE_BRANCH_TAG_SELECTS_UNIQUE_PHASE` | complete bridge tag {sigma_CP,n_C3} selects one synthetic phase branch without importing a physical CP value. |
| complete negative bridge branch tag | 0.25 | -1 | 0 | true | true | 5.8438133 | 1 | true | `CONDITIONAL_SUPPORT_COMPLETE_BRANCH_TAG_SELECTS_UNIQUE_PHASE` | complete bridge tag {sigma_CP,n_C3} selects one synthetic phase branch without importing a physical CP value. |
| CKM or PMNS used as branch selector | 0.25 | 1 | 1 | false | false | 0 | 0 | true | `FAILED_ROUTE_CKM_PMNS_BRANCH_SELECTOR_REJECTED` | CKM/PMNS phases are physical comparator data and cannot be used as native or implicit branch selectors in this ledger. |
| branch tag attempts native phase promotion | 0.25 | 1 | 1 | false | false | 0 | 0 | false | `FAILED_ROUTE_BRANCH_TAG_NATIVE_PROMOTION_REJECTED` | branch tags are bridge metadata and cannot promote a CP phase, ray branch, or orientation to native law-space. |
| invalid C3 sheet rejected | 0.25 | 1 | 3 | false | false | 0 | 3 | true | `FAILED_ROUTE_INVALID_OR_INCOMPLETE_BRANCH_TAG` | C3 sheet tag must be one of n=0,1,2. |

## Result statuses

- `CONDITIONAL_SUPPORT_GATE458_REDACTED_HARNESS_INHERITED`
- `CONDITIONAL_SUPPORT_ORIENTED_BRANCH_TAG_LEDGER_DEFINED`
- `CONDITIONAL_SUPPORT_COMPLETE_BRANCH_TAG_SELECTS_UNIQUE_PHASE`
- `CONDITIONAL_SUPPORT_ORIENTED_BRANCH_TAG_BRIDGE_ONLY_VALIDATED`
- `CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED`
- `FAILED_ROUTE_COSINE_INVARIANT_RETURNS_SIX_BRANCHES`
- `FAILED_ROUTE_CP_ODD_SIGN_ONLY_LEAVES_C3_SHEETS`
- `FAILED_ROUTE_INVALID_OR_INCOMPLETE_BRANCH_TAG`
- `FAILED_ROUTE_CKM_PMNS_BRANCH_SELECTOR_REJECTED`
- `FAILED_ROUTE_BRANCH_TAG_NATIVE_PROMOTION_REJECTED`
- `FAILED_ROUTE_CP_SIGN_NOT_NATIVE`
- `FAILED_ROUTE_C3_SHEET_NOT_NATIVE`

## Firewall

executed=true no_muon=true no_charm=true no_yukawa=true no_ckm=true no_pmns=true no_GST=true no_ray=true no_CP=true no_curvefit=true K=true triangle=true Y_sealed=true coeffs_sealed=true native_dim=13 kxy_dim=9 verdict=CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED reason=Gate459 selects phase branches only when complete bridge tags are supplied; native CP-sign and C3-sheet selectors remain absent, so the 13-moduli firewall is unchanged.

## Native no-go boundary

- `FAILED_ROUTE_CP_SIGN_NOT_NATIVE`: no native law in the current atlas selects `sign(sin(3phi))`.
- `FAILED_ROUTE_C3_SHEET_NOT_NATIVE`: no native law in the current atlas selects `n_C3 in {0,1,2}`.
- CKM/PMNS phases are rejected as hidden branch selectors in this gate.

## Next gate

Gate 460 — Branch-Resolved Synthetic Texture Residual Harness / Null-Data Run: Gate459 can select a unique bridge phase branch when a complete symbolic branch tag is supplied, so the next audit can run branch-resolved residuals on synthetic/null data while still rejecting observed flavor imports by default Primary task: compose the Gate458 evaluator with the Gate459 branch ledger to produce branch-resolved, bridge-only texture residual records with no CKM/PMNS, Yukawa, or mass data imported

## Truth statement

Gate459 proves the exact metadata boundary for phase-branch resolution: cos(3phi) gives six branches, a CP-odd sign gives three, and only the pair {sigma_CP,n_C3} gives one. That pair is a bridge tag, not native geometry; CKM/PMNS phases and native-promotion attempts are rejected.
