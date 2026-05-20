# Gate 588 — Koide Loop-Deficit Reactor-Angle Audit

## Purpose

Gate 588 continues from Gate 587.  Gate 586 wrote the charged-lepton Koide electron-wall offset as

```text
epsilon_e = (1/(8*pi))(1-kappa_e)
```

with

```text
kappa_e = 0.00550355419157456.
```

Gate 587 tested direct PMNS orientation through `J_PMNS`, but the central PMNS Jarlskog scale was too large and no PMNS candidate certified.  Gate 588 tests the sharper lepton-sector reactor-angle candidate

```text
kappa_e ?= sin^2(theta13)/4
```

using the same NuFIT 6.0 Normal Ordering input inherited from Gate 587.

## Reactor-angle input

The audit inherits the version-pinned NuFIT 6.0 input used in Gate 587:

```text
source:       NuFIT 6.0 global three-neutrino oscillation analysis
variant:      IC24 with SK atmospheric data
ordering:     Normal Ordering best fit
data through: September 2024
```

The reactor angle is:

```text
sin^2(theta13) = 0.02215 +0.00056 -0.00058.
```

Converted to an angle,

```text
theta13 central = 8.55905763231384°
theta13 1σ range = [8.44542463585360°, 8.66740052566978°].
```

## Reactor-quarter candidate

Gate 588 computes:

```text
candidate = sin^2(theta13)/4
          = 0.0055375.
```

Compared with

```text
kappa_e = 0.00550355419157456,
```

this gives:

```text
signed residual   = +0.0000339458084254443
relative residual = +0.00616797931733138.
```

The one-sigma reactor-quarter range is:

```text
[0.0053925, 0.0056775],
```

so it covers `kappa_e`.

This is the best central PMNS-sector clue so far.  It improves on the previous Gate 587 PMNS-assisted weak-coupling candidate

```text
alpha_2(M_Z)/(2*pi*c13)
relative residual = -0.00842061869716227,
```

and improves on the Gate 586 CKM orientation clue

```text
sqrt(J_CKM)
relative residual = +0.0144361172907456.
```

However, the CKM/weak midpoint remains numerically closer:

```text
0.5*(sqrt(J_CKM)+alpha_2/(2*pi))
relative residual = -0.00251386870274715.
```

That midpoint is still not lawful because no typed lepton-sector map connects it to the charged-lepton Koide wall deficit.

## Inverse prediction

If the candidate relation is inverted,

```text
sin^2(theta13)_pred = 4*kappa_e
                    = 0.0220142167662982.
```

Then:

```text
theta13_pred = arcsin(sqrt(4*kappa_e))
             = 8.53258678608598°.
```

The residual from the NuFIT central value is:

```text
theta13_pred - theta13_central = -0.0264708462278609°.
```

This lies inside the one-sigma interval:

```text
8.44542463585360° < 8.53258678608598° < 8.66740052566978°.
```

## Full epsilon prediction

Using

```text
L = 1/(8*pi)
```

and the candidate

```text
kappa_candidate = sin^2(theta13)/4,
```

Gate 588 predicts

```text
epsilon_pred = L*(1 - sin^2(theta13)/4)
             = 0.039568405648631 rad
             = 2.26710264572816°.
```

The observed charged-lepton wall offset is

```text
epsilon_e = 0.039569756309433 rad
          = 2.26718003289167°.
```

So:

```text
epsilon_pred - epsilon_e = -0.00000135066080203528 rad
                          = -0.0000773871635103761°.
```

The target epsilon lies inside the one-sigma epsilon range induced by `theta13`.

## Interpretation of the factor one quarter

The factor `1/4` is recorded only as a weak-normalization clue.  It is not promoted to a theorem.  To certify the relation, ASHA would need a native operator or bridge theorem connecting:

```text
PMNS reactor angle / lepton orientation
weak-doublet normalization
root-trace or absolute-Dirac Koide wall coordinate
```

to the charged-lepton loop-angle deficit.  No such theorem is currently present.

## Verdict

Gate 588 records:

```text
PASS_GATE587_LOOP_DEFICIT_PMNS_RUNTIME_INHERITED
PASS_NUFIT60_REACTOR_ANGLE_INPUT_INHERITED
PASS_REACTOR_QUARTER_CANDIDATE_COMPUTED
CONDITIONAL_SUPPORT_REACTOR_QUARTER_BEATS_PREVIOUS_PMNS_CANDIDATES
PASS_KAPPA_WITHIN_THETA13_ONE_SIGMA_REACTOR_QUARTER_RANGE
PASS_INVERSE_THETA13_PREDICTION_COMPUTED
PASS_INVERSE_THETA13_PREDICTION_WITHIN_NUFIT_ONE_SIGMA
PASS_FULL_EPSILON_PREDICTION_FROM_REACTOR_QUARTER_COMPUTED
CONDITIONAL_SUPPORT_FACTOR_ONE_QUARTER_WEAK_NORMALIZATION_CLUE_ONLY
CONDITIONAL_SUPPORT_REACTOR_QUARTER_COVERS_KAPPA_BUT_NOT_CERTIFIED
CONDITIONAL_SUPPORT_CKM_ALPHA2_MIDPOINT_REMAINS_CLOSER_NUMERIC_CLUE
FAILED_ROUTE_NO_NATIVE_LEPTON_ORIENTATION_WEAK_DOUBLET_ROOT_TRACE_OPERATOR
FAILED_ROUTE_THETA13_NOT_DERIVED_FROM_KOIDE_DEFICIT
FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL
FIREWALL_PRESERVED_NO_KOIDE_CHARGED_LEPTON_PMNS_NEUTRINO_OR_FLAVOR_DERIVATION
FIREWALL_PRESERVED_NUFIT_THETA13_REMAINS_VERSION_PINNED_OBSERVED_INPUT
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_GATE588_REACTOR_ANGLE_BOUNDARY
```

## Final statement

The reactor-angle quarter relation is the strongest PMNS-sector clue so far:

```text
kappa_e ≈ sin^2(theta13)/4.
```

It matches at the `0.617%` central level, covers `kappa_e` within the current NuFIT one-sigma reactor-angle uncertainty, and its inverse predicts `theta13` inside one sigma.  But it is still not a native ASHA derivation.  Without a lepton-orientation / weak-doublet / root-trace operator, `kappa_e` remains the charged-lepton Koide loop-angle deficit environmental seal.
