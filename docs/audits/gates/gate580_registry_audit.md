# Gate 580 — Koide Transport-Vector Decomposition Audit

## Purpose

Gate 580 continues the charged-lepton environmental flavor reduction from Gates 577–579.  Gate 577 identified the charged-lepton square-root Koide cone, Gate 578 computed the remaining azimuth, and Gate 579 showed that the pole-mass and `M_Z` Yukawa frames are projectively identical in v1 while the `Lambda_12` endpoint is slightly closer to `Q=2/3`.

Gate 580 asks the next typed question: how does the charged-lepton square-root vector move under the current v1 history transport when written in Koide coordinates?

This is a bridge-layer finite-difference audit.  It does **not** derive Koide, charged-lepton masses, Yukawa eigenvalues, CKM/PMNS data, generation hierarchy, flavor texture, or a native root-trace operator.

## Koide coordinate flow

Use the projective charged-lepton square-root vector:

```text
x_e(t) = rho_e(t) [ cos(theta_e(t)) n + sin(theta_e(t)) u(phi_e(t)) ]
```

where:

```text
n = (1,1,1)/sqrt(3)
u(phi) = cos(phi) e1 + sin(phi) e2
```

and:

```text
e1 = (1,-1,0)/sqrt(2)
e2 = (1,1,-2)/sqrt(6)
t = ln(mu)
```

The Koide cone is:

```text
theta_e = 45 degrees.
```

Gate 580 computes the finite-difference components:

```text
d ln rho_e / d ln mu,
d theta_e / d ln mu,
d phi_e / d ln mu
```

from `M_Z` to `Lambda_12`.

## Runtime endpoints

### `M_Z`

```text
rho(M_Z) = 0.103997928984285
Q(M_Z) = 0.666660511477385
deltaQ(M_Z) = -6.15518928115399e-06
theta(M_Z) = 44.9997354978 degrees
theta(M_Z)-45 = -0.000264502218037 degrees
phi(M_Z) = 257.267180033 degrees
```

### `Lambda_12`

```text
rho(Lambda_12) = 0.103860068494375
Q(Lambda_12) = 0.66666338118905
deltaQ(Lambda_12) = -3.28547761707654e-06
theta(Lambda_12) = 44.9998588163 degrees
theta(Lambda_12)-45 = -0.000141183696655 degrees
phi(Lambda_12) = 257.267382532 degrees
```

The transport interval is:

```text
Delta t = ln(Lambda_12/M_Z) = 27.6953098781871.
```

## Transport-vector decomposition

```text
Delta rho = -0.000137860489910338
Delta ln rho = -0.00132648742696639
d ln rho / dt = -4.7895742376623e-05

Delta theta = 0.000123318521382032 degrees
d theta / dt = 4.45268610188609e-06 degrees

Delta phi = 0.000202498653266048 degrees
d phi / dt = 7.31165869443969e-06 degrees
```

The projective angular displacement is tiny:

```text
Delta_projective = 3.29817195900213e-06 rad
                = 0.000188971333359217 degrees.
```

The radial-to-projective ratio is:

```text
|Delta ln rho| / Delta_projective = 402.18868011.
```

Thus, in the current v1 transport, the dominant motion of the charged-lepton square-root vector is radial rescaling, while the projective ray is nearly stable.

## Cone and azimuth interpretation

The cone residual improves:

```text
|Q(M_Z)-2/3|       = 6.15518928115399e-06
|Q(Lambda_12)-2/3| = 3.28547761707654e-06
improvement factor = 1.87345342095.
```

The angle `theta` moves toward the Koide cone:

```text
theta(M_Z)-45        = -0.000264502218037 degrees
theta(Lambda_12)-45 = -0.000141183696655 degrees.
```

The azimuth is nearly invariant:

```text
Delta phi = 0.000202498653266048 degrees.
```

Therefore the finite-difference evidence supports the following bridge-layer picture:

```text
rho_e: nonzero radial rescaling,
theta_e: slight motion toward the Koide cone,
phi_e: nearly invariant in v1.
```

But this does **not** certify a Koide attractor theorem.  A two-endpoint v1 finite difference is not a continuous beta function.

## Required missing theorem

Promotion beyond environmental seal would require:

```text
continuous Koide-coordinate RG equations for (rho, theta, phi),
with threshold/multi-loop control,
and a native root-trace or absolute-Dirac observable if ASHA-native promotion is attempted.
```

Gate 352 remains binding: no native root-trace operator is certified.

## Verdict

```text
PASS_GATE579_AND_HISTORY_TRANSPORT_RUNTIME_INHERITED
PASS_KOIDE_PROJECTIVE_COORDINATES_INHERITED
PASS_MZ_TO_LAMBDA12_LOG_TRANSPORT_INTERVAL_CERTIFIED
PASS_KOIDE_TRANSPORT_VECTOR_COMPONENTS_COMPUTED
PASS_CHARGED_LEPTON_TRANSPORT_DOMINATED_BY_RADIAL_RESCALING_IN_V1
CONDITIONAL_SUPPORT_THETA_COMPONENT_MOVES_TOWARD_KOIDE_CONE_IN_V1
PASS_PHI_COMPONENT_NEARLY_INVARIANT_IN_V1
CONDITIONAL_SUPPORT_CHARGED_LEPTON_PROJECTIVE_RAY_NEARLY_STABLE_IN_V1
FAILED_ROUTE_KOIDE_CONE_ATTRACTOR_NOT_CERTIFIED_BY_TWO_POINT_V1_FINITE_DIFFERENCE
FAILED_ROUTE_NO_CONTINUOUS_KOIDE_COORDINATE_BETA_FUNCTION_CERTIFIED
FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_OR_ABSOLUTE_DIRAC_TRANSPORT_OPERATOR
FAILED_ROUTE_NO_ASHA_NATIVE_CHARGED_LEPTON_FLAVOR_TRANSPORT_DERIVATION
FIREWALL_PRESERVED_NO_NEW_FLAVOR_CARRIER_OR_SELECTOR_INTRODUCED
FIREWALL_PRESERVED_OBSERVED_LEPTON_ENDPOINTS_REMAIN_HISTORY_DATA
FIREWALL_PRESERVED_KOIDE_TRANSPORT_DOES_NOT_DERIVE_TEXTURE_CKM_PMNS_OR_GENERATIONS
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_GATE580_KOIDE_TRANSPORT_VECTOR_BOUNDARY
```
