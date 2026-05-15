# Gate 481 Registry Audit — Null-Baseline Perturbation Ledger / Sector Transport Audit

## Verdict

`CONDITIONAL_SUPPORT_NULL_BASELINE_PERTURBATION_LEDGER_DEFINED`

Gate 481 accepts the Gate 480 result `alpha_vac=1`, `I_K,vac=0.5` only as a **common null-vacuum baseline**. It then asks whether this baseline can be transported into physical sector coordinates. The answer is negative at the native-law level: the transport chart can be defined, but the sector perturbations remain bridge-only.

## Transport chart

```text
alpha_s = alpha_vac + delta_alpha_s
phi_s   = phi_vac   + delta_phi_s
alpha_vac = 1
I_K,vac = 1/2
```

The chart is useful, but it is not a sector-coordinate theorem. No native rule in the current atlas fixes `delta_alpha_s` or `delta_phi_s`.

## Baseline cancellation proof

Starting from the cylinder distance:

```text
d_st=sqrt((alpha_t-alpha_s)^2+4 sin^2((phi_t-phi_s)/2))
alpha_s=alpha_vac+delta_alpha_s, phi_s=phi_vac+delta_phi_s
d_st=sqrt((delta_alpha_t-delta_alpha_s)^2+4 sin^2((delta_phi_t-delta_phi_s)/2))
```

Because `alpha_vac` and `phi_vac` are common-mode baseline data, they cancel. Relative mixing diagnostics are controlled by sector perturbations, not by the null baseline itself.

## Synthetic bridge-only dry run

| sector | delta_alpha | delta_phi | alpha | phi | I_K | bridge_only |
|---|---:|---:|---:|---:|---:|---|
| u | 0.02 | 0.1 | 1.02 | 0.1 | 0.507443846646 | true |
| d | -0.03 | -0.12 | 0.97 | -0.12 | 0.48862313618 | true |
| e | -0.05 | 0.2 | 0.95 | 0.2 | 0.480897090375 | true |
| nu | 0.04 | 0.62 | 1.04 | 0.62 | 0.514775795811 | true |

| pair | delta_alpha | delta_phi | distance | native_prediction |
|---|---:|---:|---:|---|
| synthetic quark null-baseline residual | -0.05 | -0.22 | 0.225177932619 | false |
| synthetic lepton null-baseline residual | 0.09 | 0.42 | 0.426523292887 | false |

These values are synthetic diagnostics proving the socket mechanics only. They are not CKM/PMNS predictions.

## Rejected promotions

```text
FAILED_ROUTE_NULL_TO_SECTOR_TRANSPORT_NOT_NATIVE
FAILED_ROUTE_NULL_BASELINE_DOES_NOT_FIX_SECTOR_PERTURBATIONS
FAILED_ROUTE_I_K_VAC_HALF_CANNOT_REPLACE_SECTOR_I_K
FAILED_ROUTE_NULL_BASELINE_TRANSPORT_AS_CKM_PMNS_PREDICTION_REJECTED
FAILED_ROUTE_NULL_PERTURBATION_LEDGER_NATIVE_PROMOTION_REJECTED
```

## Numerical output

```text
I_K,vac = 0.500000000000
synthetic d_ud  = 0.225177932619
synthetic d_eν  = 0.426523292887
physical d_ud  = undefined
physical d_eν  = undefined
CKM/PMNS       = not constructed
```

## Next step

Gate 482 — Null-baseline sector deformation source search: audit whether finite algebraic orientation, chirality, or Higgs-edge operators can natively source the sector perturbations delta_alpha and delta_phi without observed CKM/PMNS data
