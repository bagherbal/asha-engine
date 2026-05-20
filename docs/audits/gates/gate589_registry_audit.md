# Gate 589 — Koide-Reactor Relation Robustness and R-Defect Sensitivity Audit

## Purpose

Gate 589 continues from Gate 588.  Gate 588 found the strongest lepton-sector environmental clue so far:

```text
kappa_e ≈ sin^2(theta13)/4.
```

Gate 589 asks whether this relation belongs to the measured near-Koide charged-lepton ray, the exact `R=1` projected ratio-closure ray from Gate 584, or a two-variable correction involving the Koide amplitude defect `R_e-1`.

This is a robustness and sensitivity audit only.  It does not derive Koide, PMNS, neutrino parameters, charged-lepton masses, or flavor texture.

## Inherited wall coordinates

The observed near-Koide wall coordinate is:

```text
epsilon_obs = 0.039569756309433 rad
R_obs       = 0.999990767173456
kappa_obs   = 1 - 8*pi*epsilon_obs
            = 0.00550355419157456.
```

The exact-`R=1` ratio-closure coordinate from Gate 584 is:

```text
epsilon_R1 = 0.039577340701281 rad
kappa_R1   = 1 - 8*pi*epsilon_R1
           = 0.00531293763388241.
```

Thus:

```text
kappa_obs - kappa_R1 = 0.000190616557693035.
```

## Reactor-quarter candidate

Using the NuFIT 6.0 Normal Ordering reactor-angle input inherited from Gate 588,

```text
sin^2(theta13) = 0.02215 +0.00056 -0.00058,
```

the candidate is:

```text
sin^2(theta13)/4 = 0.0055375
one-sigma range  = [0.0053925, 0.0056775].
```

## Observed near-Koide ray comparison

For the measured near-Koide ray:

```text
kappa_obs = 0.00550355419157456
candidate - kappa_obs = +0.0000339458084254443
relative residual      = +0.00616797931733.
```

The inverse prediction is:

```text
sin^2(theta13)_pred_obs = 4*kappa_obs
                         = 0.0220142167662982

theta13_pred_obs = 8.53258678608598°.
```

This lies inside the NuFIT one-sigma interval:

```text
[8.44542463585360°, 8.66740052566978°].
```

## Exact-R=1 projected comparison

For the exact-`R=1` ratio-closure coordinate:

```text
kappa_R1 = 0.00531293763388241
candidate - kappa_R1 = +0.000224562366117591
relative residual     = +0.0422670811502625.
```

The inverse prediction is:

```text
sin^2(theta13)_pred_R1 = 4*kappa_R1
                       = 0.0212517505355296

theta13_pred_R1 = 8.38243836864531°.
```

This is below the NuFIT one-sigma range, so the exact-`R=1` projection weakens the reactor-quarter relation.

## R-defect sensitivity

The observed Koide amplitude defect is:

```text
dR = 1 - R_obs
   = 0.00000923282654397006.
```

To map the exact-`R=1` deficit to the observed deficit by a linear correction,

```text
kappa_obs = kappa_R1 + c*dR,
```

the required coefficient is:

```text
c = (kappa_obs-kappa_R1)/(1-R_obs)
  = 20.6455256996.
```

Typed trial coefficients do not certify:

```text
1, 2, sqrt(2), sqrt(3), 2*pi, 8*pi.
```

The nearest tested typed coefficient is `8*pi`, but its residual is still about `21.7%` of the required shift.  Therefore no simple typed `R`-defect correction is certified.

## Shift control

The kappa shift is exactly controlled by the epsilon projection shift because

```text
kappa = 1 - 8*pi*epsilon.
```

Indeed:

```text
epsilon_R1 - epsilon_obs = 0.00000758439184800341
8*pi*(epsilon_R1-epsilon_obs) = 0.000190616557693071
kappa_obs-kappa_R1            = 0.000190616557693035.
```

But this is a definitional relationship, not a native source theorem.

The same shift is not directly supplied by a typed `R` or `Q-2/3` correction:

```text
R defect:          1-R_obs       = 0.00000923282654397006
Q residual:        Q_obs-2/3     = -0.00000615518928093195
required c_R:      20.6455256996
required c_|Qres|: 30.9684315125.
```

No typed coefficient is certified.

## Interpretation

Gate 589 shows that the reactor-quarter relation belongs to the measured near-Koide environmental ray, not to the exact-`R=1` projected ratio-closure model.  The `R` defect is required for the best reactor match, but no lawful typed linear correction from `R_obs-1` or `Q_obs-2/3` has been found.

## Verdict

Gate 589 records:

```text
PASS_GATE588_REACTOR_QUARTER_RESULT_INHERITED
PASS_OBSERVED_NEAR_KOIDE_WALL_COORDINATE_INHERITED
PASS_EXACT_R1_RATIO_CLOSURE_COORDINATE_INHERITED
PASS_NUFIT60_REACTOR_QUARTER_CANDIDATE_INHERITED
PASS_REACTOR_QUARTER_MATCHES_OBSERVED_EPSILON_BETTER_THAN_EXACT_R1_EPSILON
PASS_OBSERVED_EPSILON_INVERSE_THETA13_PREDICTION_WITHIN_NUFIT_ONE_SIGMA
FAILED_ROUTE_EXACT_R1_INVERSE_THETA13_PREDICTION_OUTSIDE_NUFIT_ONE_SIGMA
CONDITIONAL_SUPPORT_REACTOR_RELATION_BELONGS_TO_MEASURED_NEAR_KOIDE_RAY_NOT_EXACT_R1_PROJECTION
PASS_R_DEFECT_AND_KAPPA_SHIFT_COMPUTED
PASS_REQUIRED_R_DEFECT_LINEAR_COEFFICIENT_COMPUTED
FAILED_ROUTE_NO_TYPED_SIMPLE_R_DEFECT_CORRECTION_CERTIFIED
PASS_KAPPA_SHIFT_EXACTLY_CONTROLLED_BY_EPSILON_PROJECTION_SHIFT
FAILED_ROUTE_R_DEFECT_ALONE_DOES_NOT_FIX_REACTOR_RELATION_WITH_TYPED_COEFFICIENT
FAILED_ROUTE_Q_MINUS_TWO_THIRDS_DOES_NOT_SUPPLY_TYPED_KAPPA_CORRECTION
FAILED_ROUTE_NO_NATIVE_KOIDE_REACTOR_R_DEFECT_OPERATOR
FAILED_ROUTE_KOIDE_REACTOR_RELATION_REMAINS_ENVIRONMENTAL_HISTORY_SEAL
FIREWALL_PRESERVED_NO_KOIDE_THETA13_PMNS_NEUTRINO_OR_FLAVOR_DERIVATION
FIREWALL_PRESERVED_REACTOR_AND_KOIDE_INPUTS_REMAIN_VERSION_PINNED_OBSERVED_DATA
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_GATE589_R_DEFECT_SENSITIVITY_BOUNDARY
```

## Final statement

The relation

```text
kappa_e ≈ sin^2(theta13)/4
```

is strongest for the observed near-Koide ray.  It does not survive as cleanly after exact-`R=1` projection.  Therefore the current bridge is a relation between the measured environmental charged-lepton ray and the observed PMNS reactor angle, with the small Koide amplitude defect still present.  No ASHA-native operator or theorem derives it.
