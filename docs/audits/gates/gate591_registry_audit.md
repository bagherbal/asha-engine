# Gate 591 — Koide-Reactor-CKM Residual Closure and Uncertainty Audit

## Purpose

Gate 591 continues from Gate 590.  Gate 590 found the strongest bridge-layer environmental relation so far:

```text
kappa_e ≈ sin²(theta13)/4 - J_CKM.
```

Gate 591 asks whether the remaining residual is statistically meaningful, whether it is smaller than the existing near-Koide defects, and whether a typed `R`- or `Q`-defect correction closes it.

This is an environmental bridge audit only.  It does not derive Koide, PMNS, CKM, charged-lepton masses, neutrino parameters, or flavor texture.

## Inherited relation

Observed near-Koide wall coordinate:

```text
epsilon_obs = 0.039569756309433 rad
kappa_obs   = 1 - 8*pi*epsilon_obs
            = 0.00550355419157456
R_obs       = 0.999990767173456
Q_obs       = 0.666660511477386
```

Combined Gate 590 candidate:

```text
B = sin²(theta13)/4 - J_CKM
  = 0.00550633006471245.
```

Residual:

```text
Delta_590 = B - kappa_obs
          = 2.77587313788925e-06
relative  = 5.04378269254e-04.
```

The equivalent epsilon residual is:

```text
epsilon_B - epsilon_obs = -1.10448482824876e-07 rad
                        = -0.00000632823191949°.
```

## Uncertainty input

The reactor-angle input remains NuFIT 6.0 Normal Ordering:

```text
sin²(theta13) = 0.02215 +0.00056 -0.00058.
```

Gate 591 adds a version-pinned CKM Jarlskog uncertainty for uncertainty propagation:

```text
J_CKM = (3.12 +0.13 -0.12) x 10^-5
```

The central relation still uses the runtime history-transport value:

```text
J_CKM(runtime) = 3.11699352875547e-05.
```

## Propagated one-sigma band

For:

```text
B = sin²(theta13)/4 - J_CKM,
```

the one-sigma interval is:

```text
B_1sigma = [0.00536003006471245, 0.00564753006471245].
```

This covers:

```text
kappa_obs = 0.00550355419157456.
```

The central residual occupies only about:

```text
1.90% to 1.97%
```

of the propagated one-sigma width.  The uncertainty is dominated by the `theta13` input; the CKM-J uncertainty is included but is roughly two orders of magnitude smaller than the reactor-angle contribution.

## Inverse prediction with CKM uncertainty

The inverse relation is:

```text
sin²(theta13)_pred = 4(kappa_obs + J_CKM).
```

Using the runtime central `J_CKM`:

```text
sin²(theta13)_pred = 0.0221388965074484
theta13_pred       = 8.55689599683003°.
```

Propagating CKM-J uncertainty gives:

```text
theta13_pred(J low)  = 8.55596136554858°
theta13_pred(J high) = 8.55790840231023°.
```

NuFIT one-sigma range:

```text
theta13 = [8.44542463585360°, 8.66740052566978°].
```

So the inverse prediction remains safely inside one sigma.

## Residual compared with Koide defects

The remaining residual is smaller than both near-Koide defects:

```text
Delta_590       = 2.77587313788925e-06
1 - R_obs       = 9.23282654408109e-06
|Q_obs - 2/3|   = 6.15518928104297e-06.
```

Ratios:

```text
Delta_590 / (1-R_obs)     = 0.300652581811
Delta_590 / |Q_obs-2/3|   = 0.450980954629.
```

Thus the residual is already beneath the near-Koide amplitude and cone defects.  It should not be treated as a new exact law until the near-Koide defects and input uncertainties are resolved.

## R/Q correction candidates

Gate 591 tests typed-looking corrections from the existing near-Koide defects.  To close the residual directly, one would need:

```text
Delta_590 = c_R (1-R_obs),      c_R = 0.300652581811
Delta_590 = c_Q (Q_obs-2/3),    c_Q = -0.450980954629.
```

Trial corrections were checked.  The best in the limited trial set is:

```text
(1-R_obs)/pi = 2.93889996640113e-06.
```

It leaves:

```text
corrected residual = -1.63026828512e-07
```

but this is not certified.  No ASHA theorem supplies the coefficient `1/pi` or maps the Koide amplitude/cone defect into the cross-sector orientation residual.

## Sector-lawfulness audit

The current runtime contains no native theorem or operator providing:

```text
CKM orientation area -> charged-lepton Koide wall correction,
PMNS reactor angle -> charged-lepton Koide wall correction,
R/Q Koide defect -> reactor-CKM residual correction,
root-trace / absolute-Dirac observable -> epsilon_e,
cross-sector orientation intertwiner.
```

## Verdict

```text
PASS_DELTA590_INSIDE_COMBINED_ONE_SIGMA_BAND
PASS_DELTA590_SMALLER_THAN_KOIDE_R_AND_Q_DEFECTS
CONDITIONAL_SUPPORT_DELTA590_INPUT_NOISE_LIMITED_BY_THETA13
CONDITIONAL_SUPPORT_CKM_UNCERTAINTY_SUBDOMINANT_TO_THETA13
CONDITIONAL_SUPPORT_R_DEFECT_OVER_PI_NUMERIC_CLOSURE_HINT_NOT_CERTIFIED
FAILED_ROUTE_NO_R_OR_Q_DEFECT_CORRECTION_CERTIFIED
FAILED_ROUTE_NO_CROSS_SECTOR_ORIENTATION_INTERTWINER
FAILED_ROUTE_DELTA590_REMAINS_ENVIRONMENTAL_RESIDUAL
FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL
FIREWALL_PRESERVED_GATE591_RESIDUAL_CLOSURE_UNCERTAINTY_BOUNDARY
```

Gate 591 therefore keeps the Gate 590 relation as the best current environmental bridge, but it does not promote the residual closure to native ASHA law.
