# Gate 603 — Charged-Lepton Sigma Degeneracy Gauge-or-Orientation Audit

## Purpose

Gate 603 continues the Gate 602 branch-row selector result. Gate 602 showed that the environmental balance

```text
B_flav(sigma,alpha,i,s_J)
= 1 - 8*pi*epsilon_{sigma,alpha}(H_e)
  - (1/4)Tr(P_alpha P_i^nu)
  + s_J J_CKM
```

selects the electron row, the third neutrino projector, and the positive CKM orientation sign, but leaves all six charged-lepton sigma/cyclic orderings degenerate. Gate 603 audits whether that remaining sigma degeneracy is physical branch data or Fourier-coordinate redundancy.

## Result

The audit identifies the degeneracy source: once the physical electron wall `alpha=e` has been selected, the balance uses an unsigned electron-wall distance. That distance is invariant across the six Fourier presentations of the charged-lepton root vector. Therefore `B_flav` does not see the cyclic sigma orientation.

The remaining sixfold degeneracy is therefore gauge-like for the balance functional itself. If full charged-lepton cyclic/order orientation is considered physical, the missing object is a signed discriminant/Vandermonde orientation seal:

```text
Delta_e = prod_{i<j}(lambda_i-lambda_j)^2
V_e     = prod_{i<j}(lambda_j-lambda_i)
V_x     = prod_{i<j}(x_j-x_i)
```

The native trace ring supplies the symmetric discriminant `Delta_e`, but not the signed orientation `sgn(V_e)` or `sgn(V_x)`.

## Interpretation

Gate 603 separates two statements:

1. For `B_flav`, sigma is a Fourier-coordinate redundancy once `alpha=e`, `P_3^nu`, and `+J_CKM` are selected.
2. For a full charged-lepton branch theorem, an additional `ChargedLeptonDiscriminantOrientationSeal` would be required.

No native ASHA theorem currently supplies this signed-discriminant orientation.

## Verdict

- `PASS_SIGMA_DEGENERACY_SOURCE_IDENTIFIED`
- `PASS_S3_ACTION_ON_KOIDE_FOURIER_COORDINATES_AUDITED`
- `PASS_SIGNED_DISCRIMINANT_AND_VANDERMONDE_ORIENTATION_AUDITED`
- `PASS_FOURIER_CYCLIC_ORIENTATION_AUDITED`
- `CONDITIONAL_SUPPORT_SIGMA_IS_FOURIER_COORDINATE_REDUNDANCY_FOR_B_FLAV`
- `CONDITIONAL_SUPPORT_CHARGED_LEPTON_DISCRIMINANT_ORIENTATION_SEAL_REQUIRED_FOR_FULL_ORDER_SELECTION`
- `FAILED_ROUTE_B_FLAV_DOES_NOT_SEE_CYCLIC_SIGMA`
- `FAILED_ROUTE_NO_NATIVE_SIGNED_DISCRIMINANT_ORIENTATION_THEOREM`
- `FAILED_ROUTE_NO_NATIVE_SIGMA_SELECTION_THEOREM`
- `FIREWALL_PRESERVED_GATE603_SIGMA_GAUGE_ORIENTATION_BOUNDARY`
