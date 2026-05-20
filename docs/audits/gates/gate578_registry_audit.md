# Gate 578 — Charged-Lepton Koide Azimuth Environmental Orientation Audit

## Purpose

Gate 578 continues the Gate 577 reduction.  Gate 577 showed that the charged-lepton square-root Yukawa vector lies extremely close to the Koide cone.  Gate 578 audits the remaining angular datum on that cone: the azimuth around the democratic axis.

This gate is strictly bridge-layer.  It does **not** derive charged-lepton masses, Yukawa eigenvalues, CKM/PMNS, generation hierarchy, flavor texture, ASHA projective phase, or observed data from native ASHA algebra.

## Azimuth frame

Use the orthonormal frame:

```text
n  = (1,1,1)/sqrt(3)
e1 = (1,-1,0)/sqrt(2)
e2 = (1,1,-2)/sqrt(6)
```

For:

```text
x_e=(sqrt(y_e),sqrt(y_mu),sqrt(y_tau))
```

write:

```text
x_e = (x_e·n)n + x_perp
phi_e = atan2(x_perp·e2, x_perp·e1).
```

Gate 578 certifies that this frame is orthonormal and right-handed with respect to `n`.

## Runtime azimuth certificate

At `M_Z`:

```text
Q_e = 0.666660511477385
Q_e - 2/3 = -6.15518928115e-06
theta_e = 44.999735497782 degrees
phi_e = -102.732819967108 degrees
      = 257.267180032892 degrees
```

At `Lambda_12` under v1 transport:

```text
Q_e(Lambda_12) = 0.66666338118905
theta_e(Lambda_12) = 44.999858816303 degrees
phi_e(Lambda_12) = -102.732617468455 degrees
                 = 257.267382531545 degrees
```

The azimuth drift is tiny in the current approximation:

```text
Delta phi_e = 0.000202498653266 degrees.
```

Thus the charged-lepton environmental geometry is now represented as:

```text
Y_e -> (rho_e, phi_e) on the Koide cone.
```

## Candidate phase audit

Gate 578 tests simple candidate identifications but does not force them.

The nearest simple rational turn with denominator up to 72 is:

```text
5/7 turn = 257.142857142857 degrees
```

with residual:

```text
|phi_e - 5/7 turn| = 0.124322890035 degrees.
```

The v1 transport drift scale is:

```text
0.000202498653266 degrees.
```

The gate uses a conservative certification threshold of `100 × drift`, namely:

```text
0.0202498653266 degrees.
```

Therefore the `5/7` proximity is recorded but **not certified**.  It is about six times too far even under the conservative threshold, and hundreds of times larger than the drift itself.

Root-of-unity grids also fail:

```text
nearest 120-degree grid: 240 degrees, residual 17.267180032892 degrees
nearest 90-degree grid:  270 degrees, residual 12.732819967108 degrees
nearest 60-degree grid:  240 degrees, residual 17.267180032892 degrees
nearest 45-degree grid:  270 degrees, residual 12.732819967108 degrees
```

CKM phase comparisons also fail in v1:

```text
delta_CKM = 65.718259101505 degrees, residual 168.451079068614 degrees
180+delta_CKM = 245.718259101505 degrees, residual 11.548920931386 degrees
360-delta_CKM = 294.281740898495 degrees, residual 37.014560865603 degrees
```

No PMNS runtime input is available for a lawful comparison in this gate.

## Environmental azimuth seal

Gate 578 defines the bridge-only seal:

```text
ChargedLeptonKoideAzimuthSeal
```

with carrier:

```text
positive charged-lepton square-root Yukawa cone around the democratic axis.
```

The sealed data are:

```text
rho_e,
phi_e,
Q_e≈2/3.
```

This is a compression of observed endpoint data, not a native mass theorem.

## Gate 352 inheritance

Gate 352 remains binding.  A native derivation would require a root-trace/phase operator or functor capable of producing both the Koide cone and the azimuth while preserving all flavor firewalls.  No such operator is certified here.

## Verdict

```text
PASS_GATE577_AND_HISTORY_TRANSPORT_RUNTIME_INHERITED
PASS_DEMOCRATIC_AXIS_AZIMUTH_FRAME_ORTHONORMAL_CERTIFIED
PASS_CHARGED_LEPTON_KOIDE_AZIMUTH_COMPUTED_AT_MZ
PASS_CHARGED_LEPTON_KOIDE_AZIMUTH_COMPUTED_AT_LAMBDA12
PASS_KOIDE_AZIMUTH_STABLE_UNDER_V1_TRANSPORT
PASS_CHARGED_LEPTON_YE_REDUCED_TO_RADIUS_PLUS_AZIMUTH_ON_KOIDE_CONE
CONDITIONAL_SUPPORT_NEAREST_SIMPLE_RATIONAL_IS_FIVE_SEVENTHS_TURN_BUT_NOT_CERTIFIED
CONDITIONAL_SUPPORT_CHARGED_LEPTON_KOIDE_AZIMUTH_ENVIRONMENTAL_ORIENTATION_SEAL_CANDIDATE
FAILED_ROUTE_NO_SIMPLE_RATIONAL_OR_ROOT_OF_UNITY_PHASE_MATCH_CERTIFIED
FAILED_ROUTE_KOIDE_AZIMUTH_NOT_IDENTIFIED_WITH_CKM_PHASE_OR_JARLSKOG_ORIENTATION
FAILED_ROUTE_NO_PMNS_RUNTIME_INPUT_FOR_KOIDE_AZIMUTH_IDENTIFICATION
FAILED_ROUTE_GATE352_ROOT_TRACE_OBSTRUCTION_STILL_BLOCKS_NATIVE_KOIDE_AZIMUTH_OPERATOR
FAILED_ROUTE_NO_ASHA_NATIVE_KOIDE_AZIMUTH_DERIVATION
FIREWALL_PRESERVED_KOIDE_AZIMUTH_DOES_NOT_DERIVE_FLAVOR_TEXTURE_CKM_PMNS_OR_GENERATIONS
FIREWALL_PRESERVED_KOIDE_AZIMUTH_REMAINS_OBSERVED_HISTORY_ENDPOINT_ORIENTATION
FIREWALL_PRESERVED_GATE578_KOIDE_AZIMUTH_ENVIRONMENTAL_SEAL_BOUNDARY
```
