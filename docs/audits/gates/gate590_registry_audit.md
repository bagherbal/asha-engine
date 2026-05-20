# Gate 590 — Koide-Reactor-CKM Orientation Combination Audit

## Purpose

Gate 590 continues from Gate 589.  Gate 588 found the reactor-angle bridge

```text
kappa_e ≈ sin^2(theta13)/4.
```

Gate 589 showed that this relation belongs to the measured near-Koide charged-lepton ray, not to the exact `R=1` projected ratio-closure coordinate.  Gate 590 tests the sharper typed environmental candidate

```text
kappa_e ?= sin^2(theta13)/4 - J_CKM
```

equivalently

```text
epsilon_e ?= (1/(8*pi)) [1 - sin^2(theta13)/4 + J_CKM].
```

This is a bridge-layer environmental orientation audit only.  It does not derive Koide, PMNS, CKM, charged-lepton masses, neutrino parameters, or flavor texture.

## Inherited observed Koide wall coordinate

From Gate 589, the measured near-Koide charged-lepton ray gives:

```text
epsilon_obs = 0.039569756309433 rad
epsilon_obs = 2.26718003289167°
R_obs       = 0.999990767173456
Q_obs       = 0.666660511477386
kappa_obs   = 1 - 8*pi*epsilon_obs
            = 0.00550355419157456.
```

The small amplitude and cone residuals remain:

```text
R_obs - 1   = -9.23282654408109e-06
Q_obs - 2/3 = -6.15518928104297e-06.
```

## Input orientation data

The reactor angle is inherited from the NuFIT 6.0 Normal Ordering input used in Gates 587–589:

```text
sin^2(theta13) = 0.02215 +0.00056 -0.00058.
```

The CKM orientation input is the runtime history-transport Jarlskog invariant:

```text
J_CKM = 3.11699352875547e-05.
```

No CKM uncertainty is present in the runtime data file, so the uncertainty propagation can only use the reactor-angle one-sigma interval in this gate.

## Candidate comparison

Define:

```text
A = sin^2(theta13)/4
B = A - J_CKM.
```

Numerically:

```text
A = 0.0055375
B = 0.00550633006471245.
```

Against the observed deficit:

```text
kappa_obs = 0.00550355419157456.
```

The reactor-quarter candidate alone gives:

```text
A - kappa_obs       = +3.39458084254443e-05
relative residual   = +0.00616797931733.
```

The combined reactor-minus-CKM candidate gives:

```text
B - kappa_obs       = +2.77587313788925e-06
relative residual   = +0.000504378269254.
```

Thus the combined candidate improves the central residual by:

```text
|A-kappa_obs| / |B-kappa_obs| = 12.2288760109752.
```

## Epsilon prediction

The reactor-quarter-only prediction is:

```text
epsilon_A = (1/(8*pi))(1-A)
          = 0.039568405648631 rad
          = 2.26710264572816°.
```

Its residual is:

```text
epsilon_A - epsilon_obs = -1.35066080203528e-06 rad
                        = -0.0000773871635103761°.
```

The combined orientation prediction is:

```text
epsilon_B = (1/(8*pi))(1-B)
          = 0.0395696458609502 rad
          = 2.26717370465975°.
```

Its residual is:

```text
epsilon_B - epsilon_obs = -1.10448482824876e-07 rad
                        = -0.00000632823191948858°.
```

So the combined formula

```text
epsilon_e ≈ (1/(8*pi)) [1 - sin^2(theta13)/4 + J_CKM]
```

is the tightest typed environmental bridge found so far.

## Inverse theta13 prediction

If

```text
kappa_e = sin^2(theta13)/4 - J_CKM,
```

then

```text
sin^2(theta13)_pred = 4(kappa_obs + J_CKM)
                    = 0.0221388965074484.
```

This gives:

```text
theta13_pred = 8.55689599683003°.
```

NuFIT central and one-sigma interval:

```text
theta13_central = 8.55905763231384°
one-sigma range = [8.44542463585360°, 8.66740052566978°].
```

The residual is:

```text
theta13_pred - theta13_central = -0.00216163548381°.
```

The inverse prediction is well inside the NuFIT one-sigma interval.

## Uncertainty audit

With `J_CKM` held fixed because the runtime does not include its uncertainty, the one-sigma range of the combined candidate is:

```text
B_1sigma = [0.00536133006471245, 0.00564633006471245].
```

This range covers:

```text
kappa_obs = 0.00550355419157456.
```

However, full certification is blocked because the CKM `J` uncertainty is not present in the runtime artifact.

## Sector-lawfulness audit

The combined candidate is typed, but cross-sector:

```text
lepton reactor leakage: sin^2(theta13)/4
quark CP orientation:   J_CKM
charged-lepton wall:    kappa_e.
```

The current ASHA runtime contains no native theorem or operator providing:

```text
CKM orientation area -> charged-lepton Koide wall correction,
PMNS reactor angle -> charged-lepton Koide wall correction,
root-trace / absolute-Dirac observable -> epsilon_e,
cross-sector orientation intertwiner.
```

Therefore the numerical bridge is not promoted to native law.

## Residual control

The remaining combined residual is:

```text
B - kappa_obs = 2.77587313788925e-06.
```

It is smaller than both the observed Koide amplitude defect and the Koide-cone residual:

```text
1 - R_obs       = 9.23282654408109e-06
Q_obs - 2/3     = -6.15518928104297e-06.
```

Ratios:

```text
(B-kappa_obs)/|R_obs-1|   = 0.300652581811
(B-kappa_obs)/|Q_obs-2/3| = 0.450980954629.
```

No typed coefficient is certified from these residuals.

## Verdict

Gate 590 records:

```text
PASS_GATE589_NEAR_KOIDE_RAY_RESULT_INHERITED
PASS_NUFIT60_REACTOR_INPUT_INHERITED
PASS_RUNTIME_CKM_JARLSKOG_INHERITED
PASS_REACTOR_QUARTER_CANDIDATE_A_COMPUTED
PASS_REACTOR_MINUS_CKM_CANDIDATE_B_COMPUTED
PASS_REACTOR_MINUS_CKM_OUTPERFORMS_REACTOR_QUARTER_ALONE
PASS_COMBINED_ORIENTATION_CANDIDATE_COVERS_KAPPA_WITH_THETA13_ONE_SIGMA
PASS_COMBINED_INVERSE_THETA13_PREDICTION_COMPUTED
PASS_COMBINED_INVERSE_THETA13_PREDICTION_WITHIN_NUFIT_ONE_SIGMA
PASS_COMBINED_EPSILON_PREDICTION_COMPUTED
CONDITIONAL_SUPPORT_COMBINED_ORIENTATION_RESIDUAL_AT_FIVE_E_MINUS_FOUR_RELATIVE
CONDITIONAL_SUPPORT_CKM_J_UNCERTAINTY_NOT_PRESENT_IN_RUNTIME
CONDITIONAL_SUPPORT_AVAILABLE_UNCERTAINTY_DOMINATED_BY_THETA13_INPUT
FAILED_ROUTE_NO_CROSS_SECTOR_ORIENTATION_INTERTWINER
FAILED_ROUTE_NO_NATIVE_KOIDE_REACTOR_CKM_OPERATOR
FAILED_ROUTE_COMBINED_RESIDUAL_NOT_TYPED_R_DEFECT_OR_Q_RESIDUAL
FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL
FIREWALL_PRESERVED_NO_KOIDE_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION
FIREWALL_PRESERVED_REACTOR_CKM_AND_KOIDE_INPUTS_REMAIN_OBSERVED_DATA
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_GATE590_REACTOR_CKM_ORIENTATION_BOUNDARY
```

## Final statement

The best current bridge-layer environmental relation is:

```text
1 - 8*pi*epsilon_e ≈ sin^2(theta13)/4 - J_CKM.
```

Equivalently:

```text
epsilon_e ≈ (1/(8*pi)) [1 - sin^2(theta13)/4 + J_CKM].
```

It strongly improves the reactor-quarter relation and predicts `theta13` well within one sigma, but it remains a cross-sector environmental seal.  ASHA has not derived a native orientation intertwiner connecting CKM area, PMNS reactor leakage, and the charged-lepton Koide chamber wall.
