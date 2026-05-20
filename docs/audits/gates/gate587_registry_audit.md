# Gate 587 — Koide Loop-Deficit PMNS Orientation Audit

## Purpose

Gate 587 continues from Gate 586.  The charged-lepton Koide electron-wall offset has been compressed to

```text
epsilon_e = (1/(8*pi))(1-kappa_e)
```

with

```text
kappa_e = 0.00550355419157456.
```

Gate 586 found percent-level clues from `sqrt(J_CKM)` and `alpha_2(M_Z)/(2*pi)`, but CKM is a quark-sector orientation invariant.  Gate 587 therefore imports a version-pinned PMNS data set and asks whether the loop-angle deficit is more naturally lepton-sector orientation data.

## PMNS input

The audit uses NuFIT 6.0, global three-neutrino oscillation analysis, data through September 2024, with the `IC24 with SK atmospheric data` variant and Normal Ordering best fit.  The convention is the standard PMNS parametrization

```text
J_PMNS = c12 c23 c13^2 s12 s23 s13 sin(delta_CP).
```

The pinned central values and one-sigma intervals are:

```text
sin^2(theta12) = 0.308 +0.012 -0.011
theta12        = 33.68° +0.73 -0.70

sin^2(theta23) = 0.470 +0.017 -0.013
theta23        = 43.3° +1.0 -0.8

sin^2(theta13) = 0.02215 +0.00056 -0.00058
theta13        = 8.56° +0.11 -0.11

delta_CP       = 212° +26 -41
```

The CP phase is explicitly treated as uncertain and non-Gaussian; the gate uses a conservative one-sigma corner scan only as an audit sieve, not as a precision likelihood calculation.

## PMNS orientation invariants

From the central values, Gate 587 computes:

```text
s13        = 0.148828760661372
c13        = 0.988862983430971
J_PMNS     = -0.0177698631165826
|J_PMNS|   = 0.0177698631165826
sqrt(|J_PMNS|) = 0.133303650049736
```

Direct PMNS orientation scales are therefore much larger than `kappa_e`:

```text
sqrt(|J_PMNS|) / kappa_e - 1 = 23.2213750259445
|J_PMNS| / kappa_e - 1       = 2.22879770018193.
```

Thus the direct PMNS Jarlskog scale does not explain the charged-lepton loop-angle deficit.

## PMNS-assisted correction candidates

The closest PMNS-assisted typed candidate is:

```text
alpha_2(M_Z)/(2*pi*c13) = 0.00545721086024814.
```

Compared with `kappa_e`:

```text
signed residual   = -0.0000463433313264185
relative residual = -0.00842061869716227.
```

This is closer than the Gate 586 quark-sector clue

```text
sqrt(J_CKM) = 0.0055830041454001
relative residual = +0.0144361172907456,
```

but it is still not certified.  Its one-sigma range from the NuFIT `theta13` uncertainty is

```text
[0.00545559314027467, 0.00545877416306492],
```

which does not cover `kappa_e`.

## Uncertainty audit

The broad one-sigma corner scan for direct PMNS area gives:

```text
|J_PMNS| range = [0.00511683221330893, 0.0291209215394155].
```

This range can cross `kappa_e`, but this is not a certification: the central value is far away, the crossing is driven by the poorly constrained `delta_CP`, and no ASHA operator maps a PMNS area into the charged-lepton Koide wall-deficit coordinate.

The direct square-root invariant remains far away even under the same scan:

```text
sqrt(|J_PMNS|) range = [0.0715320362726305, 0.170648532192385].
```

## CKM comparison

Gate 586's `sqrt(J_CKM)` clue survives only as a numerical clue, not a source theorem.  A midpoint of the quark-sector orientation clue and the weak coupling correction is numerically even closer:

```text
0.5*(sqrt(J_CKM)+alpha_2/(2*pi)) = 0.00548971897893848
relative residual = -0.00251386870274715.
```

This is the closest numerical candidate in the audit, but it is not certified because it is a midpoint of two unrelated bridge quantities without a typed lepton-sector map.

## Verdict

Gate 587 records:

```text
PASS_GATE586_LOOP_ANGLE_DEFICIT_INHERITED
PASS_NUFIT60_PMNS_DATASET_IMPORTED
PASS_PMNS_ORIENTATION_INVARIANTS_COMPUTED
PASS_PMNS_KAPPA_CANDIDATE_SET_DEFINED
PASS_PMNS_UNCERTAINTY_PROPAGATED
CONDITIONAL_SUPPORT_ABS_J_PMNS_CAN_COVER_KAPPA_WITH_CP_UNCERTAINTY_BUT_NOT_CERTIFIED
CONDITIONAL_SUPPORT_BEST_PMNS_ASSISTED_CANDIDATE_IS_ALPHA2_OVER_2PI_DIV_C13
CONDITIONAL_SUPPORT_PMNS_ASSISTED_COUPLING_CANDIDATE_BEATS_SQRT_J_CKM_BUT_NOT_CERTIFIED
FAILED_ROUTE_DIRECT_PMNS_ORIENTATION_INVARIANTS_TOO_LARGE_FOR_KAPPA
FAILED_ROUTE_NO_PMNS_CANDIDATE_CERTIFIED_WITH_UNCERTAINTIES
CONDITIONAL_SUPPORT_CKM_ALPHA2_MIDPOINT_NUMERIC_CLUE_SURVIVES_BUT_NOT_LAWFUL
FAILED_ROUTE_CKM_CLUE_STILL_NOT_LEPTON_SOURCE_WITHOUT_INTERTWINER
FAILED_ROUTE_NO_LEPTON_ORIENTATION_TO_KOIDE_DEFICIT_INTERTWINER
FAILED_ROUTE_KAPPA_E_REMAINS_ENVIRONMENTAL_HISTORY_SEAL
FIREWALL_PRESERVED_NO_KOIDE_PMNS_NEUTRINO_OR_FLAVOR_DERIVATION
FIREWALL_PRESERVED_PMNS_DATA_REMAINS_VERSION_PINNED_OBSERVED_INPUT
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_GATE587_PMNS_ORIENTATION_BOUNDARY
```

## Final statement

PMNS data does not certify the Koide loop-angle deficit.  Direct PMNS orientation invariants are too large.  The best PMNS-assisted coupling correction, `alpha_2/(2*pi*c13)`, is closer than `sqrt(J_CKM)` but still misses by about `0.842%` and is not covered by the `theta13` uncertainty.  The CKM/alpha midpoint remains a stronger numerical coincidence, but it is not lawful.  Therefore `kappa_e` remains the charged-lepton Koide loop-angle deficit environmental seal.
