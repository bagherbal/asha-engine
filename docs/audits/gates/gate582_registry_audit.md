# Gate 582 — Koide Fourier/Circulant Phase Audit

## Purpose

Gate 582 rewrites the charged-lepton square-root Yukawa ray from Gates 577–581 in the democratic plus Fourier-plane form

```text
x_j = A [ 1 + sqrt(2) R cos(delta + 2*pi*j/3) ], j=0,1,2.
```

The audit tests whether the remaining charged-lepton Koide azimuth is more naturally expressed as a Fourier/circulant phase `delta`, whether that phase is stable under v1 transport, and whether any simple rational/root-of-unity or native ASHA selector is certified.

This is an environmental-seal audit only. It does not derive charged-lepton masses, Yukawa eigenvalues, CKM/PMNS data, a generation hierarchy, flavor texture, or a native root-trace/absolute-Dirac observable.

## Formula certificate

For positive square-root Yukawa coordinates `x_j=sqrt(y_j)`, define

```text
A = (x_0+x_1+x_2)/3,
C = sum_j (x_j/A - 1) exp(-2*pi*i*j/3),
R = |C| / (3*sqrt(2)/2),
delta = arg(C).
```

Then

```text
x_j = A [ 1 + sqrt(2) R cos(delta + 2*pi*j/3) ].
```

The Koide cone is exactly the Fourier-amplitude condition

```text
Q = (x·x)/(sum_j x_j)^2 = (1+R^2)/3,
Q = 2/3 iff R = 1.
```

Status: `PASS_FOURIER_CIRCULANT_KOIDE_FRAME_DERIVED`; `PASS_KOIDE_CONE_EQUIVALENT_TO_FOURIER_PLANE_AMPLITUDE_ONE`.

## Runtime result

Using the canonical observed charged-lepton order `(e, mu, tau)`:

```text
M_Z:
  A = 0.04245717271902
  R = 0.999990767173456
  R - 1 = -9.23282654408109e-06
  Q = 0.666660511477385
  Q - 2/3 = -6.15518928115399e-06
  delta = 132.732819967108 degrees
  delta_turn = 0.368702277686412

Lambda_12:
  A = 0.0424007998905763
  R = 0.999995071771431
  R - 1 = -4.92822856912323e-06
  Q = 0.66666338118905
  Q - 2/3 = -3.28547761707654e-06
  delta = 132.732617468455 degrees
  delta_turn = 0.368701715190153
```

The phase drift is

```text
delta(Lambda_12)-delta(M_Z) = -0.000202498653266048 degrees.
```

The Fourier amplitude moves closer to `R=1`, matching Gate 579/580's observation that the v1 boundary frame is slightly closer to the exact Koide cone.

Status: `PASS_CANONICAL_CHARGED_LEPTON_FOURIER_PHASE_COMPUTED_AT_MZ`; `PASS_CANONICAL_CHARGED_LEPTON_FOURIER_PHASE_COMPUTED_AT_LAMBDA12`; `PASS_FOURIER_PHASE_STABLE_UNDER_V1_TRANSPORT`; `CONDITIONAL_SUPPORT_FOURIER_PHASE_EXPOSES_GENERATION_PLANE_ORIENTATION`.

## Relation to Gate 578 azimuth

With the Gate 578 basis

```text
n  = (1,1,1)/sqrt(3),
e1 = (1,-1,0)/sqrt(2),
e2 = (1,1,-2)/sqrt(6),
```

and the canonical Fourier convention above, the phase is just a rotated/reflected form of the same generation-plane azimuth:

```text
delta = pi/6 - phi  mod 2*pi.
```

Therefore the Fourier phase is useful because it exposes the circulant form of the Koide ray, but it is not an independent new datum beyond the projective orientation already seen in `phi_e`.

Status: `PASS_FOURIER_PHASE_RELATED_TO_GATE578_AZIMUTH_BY_DELTA_EQUALS_PI_OVER_SIX_MINUS_PHI`.

## Permutation and simple-phase audit

The phase depends on the chosen ordering and phase convention. For the six label orderings at `M_Z`, the audit finds:

```text
(e,mu,tau):  delta = 132.732819967108°, nearest 7/19 turn, residual 0.101241019739916°
(e,tau,mu):  delta = 227.267180032892°, nearest 12/19 turn, residual 0.101241019739888°
(mu,e,tau):  delta = 107.267180032892°, nearest 14/47 turn, residual 0.0331374797002582°
(mu,tau,e):  delta = 252.732819967108°, nearest 33/47 turn, residual 0.0331374797001729°
(tau,mu,e):  delta = 347.267180032892°, nearest 55/57 turn, residual 0.101241019739859°
(tau,e,mu):  delta = 12.7328199671083°, nearest 2/57 turn, residual 0.101241019739888°
```

The best rational candidate below denominator `72` is `33/47` turn under the `(mu,tau,e)` ordering, but its residual is

```text
0.0331374797001729 degrees,
```

while the drift-based certification threshold is

```text
100 * |delta drift| = 0.0202498653266048 degrees.
```

So no simple rational or root-of-unity phase is certified.

Status: `FAILED_ROUTE_FOURIER_PHASE_NOT_UNIQUE_UNDER_PERMUTATION_OR_PHASE_CONVENTION`; `FAILED_ROUTE_NO_SIMPLE_RATIONAL_FOURIER_PHASE_CERTIFIED`.

## Firewall

Gate 582 does not derive the charged-lepton projective ray. It only rewrites the observed environmental ray in Fourier/circulant coordinates.

A native promotion would require a root-trace/absolute-Dirac or circulant generation-plane operator that selects the ray, fixes the ordering/phase convention, and overcomes Gate 352's obstruction.

Status: `FAILED_ROUTE_NO_NATIVE_CIRCULANT_GENERATION_OPERATOR_OR_ROOT_TRACE_PHASE_SELECTOR`; `FAILED_ROUTE_NO_ASHA_NATIVE_CHARGED_LEPTON_PROJECTIVE_RAY_DERIVATION_FROM_FOURIER_PHASE`; `FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING`; `FIREWALL_PRESERVED_GATE582_KOIDE_FOURIER_PHASE_BOUNDARY`.

## Final verdict

```text
ChargedLeptonKoideFourierPhaseSeal:
  delta(M_Z) = 132.732819967108 degrees
  delta(Lambda_12) = 132.732617468455 degrees
  R(M_Z) = 0.999990767173456
  R(Lambda_12) = 0.999995071771431
  phase_stable_in_v1 = true
  simple_rational_certified = false
  native_selector_certified = false
```

Gate 582 certifies the Fourier/circulant coordinate form of the charged-lepton Koide environmental seal, but it does not decode the phase as a simple symbolic angle or derive it from ASHA-native law.
