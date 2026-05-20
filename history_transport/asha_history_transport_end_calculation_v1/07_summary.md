# ASHA History Transport End Calculation v1

Task: `ASHA-HISTORY-TRANSPORT-END-CALCULATION-V1`

This is an observed-data transport calculation, not a native ASHA derivation.  It asks how the ASHA boundary normalization `k_Y=5/3`, `g1=g2`, and `sin^2(theta_*)=3/8` maps to the measured endpoint at `mu0=M_Z`.

## Status ledger

- `PASS_END_VECTOR_MZ_REPRODUCIBLY_BUILT_FROM_PINNED_OBSERVED_INPUTS`
- `PASS_G1_G2_BOUNDARY_SCALE_LAMBDA12_SOLVED_ONE_LOOP_SM`
- `PASS_ASHA_BOUNDARY_WEAK_ANGLE_THREE_EIGHTHS_CERTIFIED`
- `PASS_WEAK_ANGLE_TRANSPORT_RESIDUAL_VISIBLE`
- `PASS_STRONG_COUPLING_MISMATCH_DELTA3_VISIBLE`
- `CONDITIONAL_SUPPORT_SCALAR_TRANSPORT_COMPUTED_ONE_LOOP_TOP_DOMINANT_APPROXIMATION`
- `CONDITIONAL_SUPPORT_FLAVOR_TRANSPORT_COMPUTED_DIAGONAL_ONE_LOOP_APPROXIMATION`
- `CONDITIONAL_SUPPORT_PLANCK_LCDM_ENDPOINT_INCLUDED_AS_COSMOLOGY_SEAL`
- `FIREWALL_PRESERVED_NO_OBSERVED_INPUT_IMPORTED_AS_ASHA_NATIVE_DERIVATION`
- `FIREWALL_PRESERVED_G1_G2_TEST_ONLY_NO_FULL_PHYSICAL_UNIFICATION_CLAIM`
- `FIREWALL_PRESERVED_THRESHOLDS_AND_SCHEMES_EXPLICITLY_LABELED`

## Phase 1 — End vector at M_Z

- `v = 246.219650794 GeV`
- `gY = 0.350075688597`
- `g1 = 0.451945770616`
- `g2 = 0.652752123893`
- `g3 = 1.21719969415`
- `sin2_theta_End = 0.223376644705`
- `lambda(M_Z) = 0.12965256505`

## Phase 2 — Gauge boundary running

- `Lambda_12 = 9.72424831265e+13 GeV`
- `g_star = 0.537781779093`
- `g3(Lambda_12) = 0.56520509342`
- `Delta_3 = -0.327390433`
- `R_3 = 1.0509933869`
- interpretation: `threshold_needed`

## Phase 3 — Weak-angle transport

- `sin2_theta_boundary = 0.375`
- `sin2_theta_End = 0.223376644705`
- `Delta_sin2 = -0.151623355295`
- `transport_required = true`

## Phase 4 — Scalar transport

- `lambda(Lambda_12) = -0.0497009420777`
- `y_t(Lambda_12) = 0.480920030972`
- `beta_lambda(M_Z) = -0.0240692903178`
- zero crossing: `2575927.20461 GeV`
- status: `lambda_crosses_zero_before_lambda12_in_v1_approximation`

## Phase 5 — Flavor transport

- `J_CKM = 3.11699352876e-05`
- `Koide_Qe = 0.666660511477`
- convention: `Y_u=diag(y_u,y_c,y_t); Y_d=V_CKM diag(y_d,y_s,y_b); Y_e=diag(y_e,y_mu,y_tau); PMNS skipped in v1`
- Yukawa spectra remain sharply hierarchical; hierarchy is observed history data, not ASHA-native derivation.
- CKM phase and Jarlskog invariant are imported endpoint data and held fixed in v1 transport.
- Neutrino/PMNS sector skipped in v1 by explicit convention.
- Full matrix RGE with thresholds is required before any precision flavor-boundary claim.

## Phase 6 — History residual

The residual vector is nonzero and therefore records history seals:

```text
R_hist = (Delta_3=-0.327390433, Delta_sin2=-0.151623355295, lambda_Lambda12=-0.0497009420777, J_CKM=3.11699352876e-05, Koide_Qe=0.666660511477, Omega_c h^2=0.12, Omega_b h^2=0.0224)
```

## Firewall

The calculation preserves the ASHA boundary: it does not claim full gauge unification, does not hide thresholds, does not derive Yukawa/flavor data, does not derive Planck cosmology, and does not import observed masses as native finite algebra.
