# ASHA-HISTORY-TRANSPORT-END-CALCULATION-V1 Audit

## Scope

This audit implements the minimal Standard-Model history-transport calculation from ASHA boundary law to the measured endpoint at `mu0=M_Z`.

It is an observed-data airlock, not a theorem claiming native derivation of the measured universe.

## Certified ASHA boundary input

```text
k_Y = 5/3
g1(Lambda)=g2(Lambda)
g1^2 = (5/3) gY^2
sin^2(theta_*) = 3/8
A_F = C + H + M_3(C)
H_phi ~= C^2
```

## Output bundle

```text
history_transport/asha_history_transport_end_calculation_v1/01_inputs.yaml
history_transport/asha_history_transport_end_calculation_v1/02_end_vector.json
history_transport/asha_history_transport_end_calculation_v1/03_boundary_running.json
history_transport/asha_history_transport_end_calculation_v1/04_scalar_transport.json
history_transport/asha_history_transport_end_calculation_v1/05_flavor_transport.json
history_transport/asha_history_transport_end_calculation_v1/06_history_residual.json
history_transport/asha_history_transport_end_calculation_v1/07_summary.md
```

## Numeric core result

Using the pinned v1 inputs:

```text
v                     = 246.219650794 GeV
gY(M_Z)                = 0.350075688597
g1(M_Z)                = 0.451945770616
g2(M_Z)                = 0.652752123893
g3(M_Z)                = 1.21719969415
sin^2(theta_End)       = 0.223376644705
lambda(M_Z)            = 0.12965256505
Lambda_12              = 9.72424831265e13 GeV
g_star                 = 0.537781779093
g3(Lambda_12)          = 0.56520509342
Delta_3                = -0.327390433
R_3                    = 1.0509933869
Delta_sin2             = -0.151623355295
lambda(Lambda_12)      = -0.0497009420777
zero_crossing_v1       = 2.57592720461e6 GeV
J_CKM                  = 3.11699352876e-5
Koide_Qe               = 0.666661338648
```

## Status ledger

```text
PASS_END_VECTOR_MZ_REPRODUCIBLY_BUILT_FROM_PINNED_OBSERVED_INPUTS
PASS_G1_G2_BOUNDARY_SCALE_LAMBDA12_SOLVED_ONE_LOOP_SM
PASS_ASHA_BOUNDARY_WEAK_ANGLE_THREE_EIGHTHS_CERTIFIED
PASS_WEAK_ANGLE_TRANSPORT_RESIDUAL_VISIBLE
PASS_STRONG_COUPLING_MISMATCH_DELTA3_VISIBLE
CONDITIONAL_SUPPORT_SCALAR_TRANSPORT_COMPUTED_ONE_LOOP_TOP_DOMINANT_APPROXIMATION
CONDITIONAL_SUPPORT_FLAVOR_TRANSPORT_COMPUTED_DIAGONAL_ONE_LOOP_APPROXIMATION
CONDITIONAL_SUPPORT_PLANCK_LCDM_ENDPOINT_INCLUDED_AS_COSMOLOGY_SEAL
FIREWALL_PRESERVED_NO_OBSERVED_INPUT_IMPORTED_AS_ASHA_NATIVE_DERIVATION
FIREWALL_PRESERVED_G1_G2_TEST_ONLY_NO_FULL_PHYSICAL_UNIFICATION_CLAIM
FIREWALL_PRESERVED_THRESHOLDS_AND_SCHEMES_EXPLICITLY_LABELED
```

## Residual interpretation

The output residual vector is nonzero. That is the expected result.

```text
R_hist = (
  Delta_3,
  Delta_sin2,
  lambda(Lambda_12),
  Y(Lambda_12),
  J_CKM,
  Koide_Qe,
  Planck-LambdaCDM endpoint
)
```

The residual is the mathematical fingerprint of history transport.  It measures what remains after ASHA boundary normalization is mapped to measured low-energy endpoint data.

## Firewalls

The calculation does not claim:

- full gauge unification;
- threshold completion;
- native derivation of `m_W`, `m_Z`, `m_H`, `G_F`, `alpha_s`, fermion masses, CKM, or Planck cosmology;
- native derivation of Yukawa texture, generation hierarchy, CKM/PMNS, scalar stability, or observed flavor data;
- physical unification from a single `g1=g2` crossing.

## Reproduction

```bash
go test ./pkg/historytransport ./pkg/asha -count=1 -timeout=180s
go run ./cmd/asha-history-transport --out history_transport/asha_history_transport_end_calculation_v1
```

Do not run `internal/app` tests for this audit unless timeout risk is explicitly accepted.
