# Gate 581 — Koide Coordinate Beta-Function Audit

## Purpose

Gate 581 continues the charged-lepton environmental flavor reduction from Gates 577–580. Gate 580 showed that the square-root charged-lepton vector moves almost entirely by radial rescaling from `M_Z` to `Lambda_12`, with a nearly frozen projective ray and a small motion toward the Koide cone.

Gate 581 asks the sharper continuous question: what do the v1 charged-lepton Yukawa beta functions say directly about the Koide coordinates?

This is a bridge-layer transport audit. It does **not** derive Koide, charged-lepton masses, Yukawa eigenvalues, CKM/PMNS data, generation hierarchy, flavor texture, or a native root-trace/absolute-Dirac observable.

## Coordinate beta functions

Let:

```text
x_i = sqrt(y_i)
x = rho s
|s| = 1
s = cos(theta)n + sin(theta)u(phi)
```

with:

```text
n = (1,1,1)/sqrt(3)
u(phi) = cos(phi)e1 + sin(phi)e2
```

and:

```text
e1 = (1,-1,0)/sqrt(2)
e2 = (1,1,-2)/sqrt(6)
t = ln(mu)
```

For charged-lepton Yukawa rates:

```text
r_i = d ln y_i / dt
```

we have:

```text
dx_i/dt = (1/2) r_i x_i.
```

Therefore:

```text
d ln rho/dt = (1/2) sum_i s_i^2 r_i
```

and:

```text
ds/dt = (1/2)(diag(r_i) - sum_j s_j^2 r_j I)s.
```

Projecting onto the Koide frame gives:

```text
d theta/dt = (ds/dt) · (-sin(theta)n + cos(theta)u(phi))
```

and:

```text
d phi/dt = ((ds/dt) · (-sin(phi)e1 + cos(phi)e2))/sin(theta).
```

## Common-rescaling theorem in v1

In the v1 diagonal charged-lepton RGE:

```text
r_i = A(t) + (3/2)y_i^2/(16*pi^2),
```

where `A(t)` contains the gauge and trace terms common to all three charged leptons.

If all `r_i=A(t)`, then:

```text
ds/dt = 0,
d theta/dt = 0,
d phi/dt = 0.
```

Thus common multiplicative charged-lepton running changes only `rho`. The projective motion is sourced only by the tiny family-dependent self terms `(3/2)y_i^2/(16*pi^2)` in this v1 model.

## Runtime local beta values

### `M_Z`

```text
rho = 0.103997928984285
theta = 44.9997354978 degrees
theta - 45 = -0.000264502218037 degrees
phi = 257.267180033 degrees

r_e   = 0.0090184960056189
r_mu  = 0.00901849950389275
r_tau = 0.00901948538416842
common rate = 0.00901849600553707
rate spread = 9.89378549523176e-07
relative spread/common = 0.000109705492902

d ln rho/dt = 0.00450971489855077
d theta/dt = 4.25133316926433e-06 degrees
d phi/dt   = 6.98104646216702e-06 degrees
projective speed = 6.51468445802022e-06 degrees
```

The local `theta` beta is positive while `theta<45°`, so the local v1 flow points toward the Koide cone at the runtime endpoint.

### `Lambda_12`

```text
rho = 0.103860068494375
theta = 44.9998588163 degrees
theta - 45 = -0.000141183696655 degrees
phi = 257.267382532 degrees

r_e   = -0.00384591001312992
r_mu  = -0.00384590653355637
r_tau = -0.00384492586709711
common rate = -0.00384591001321131
rate spread = 9.84146032812026e-07
relative spread/common = 0.000255894191344

d ln rho/dt = -0.00192249057935013
d theta/dt = 4.22880857135361e-06 degrees
d phi/dt   = 6.9440093918165e-06 degrees
projective speed = 6.48014940660288e-06 degrees
```

Again, the local `theta` beta points toward the Koide cone at the runtime endpoint.

## Exact-cone invariant test

The crucial test is not merely whether the runtime point moves toward the cone. The sharper question is whether the exact Koide cone is an invariant surface:

```text
theta = 45 degrees.
```

Gate 581 evaluates the same v1 beta function at exact `theta=45°`, holding the endpoint radius and azimuth fixed.

```text
M_Z exact-cone d theta/dt       = 4.25122615017455e-06 degrees
Lambda_12 exact-cone d theta/dt = 4.22875174961713e-06 degrees
```

These values are nonzero. Therefore:

```text
FAILED_ROUTE_KOIDE_CONE_NOT_RG_INVARIANT_IN_V1_COORDINATE_BETA
```

The v1 local flow points toward the cone from the current side, but the cone itself is not fixed. Hence no attractor theorem is certified.

## Interpretation

Gate 581 separates the transport mechanism from the environmental mystery:

```text
common running -> radial flow only,
family-dependent rate splitting -> tiny projective motion,
Koide cone placement -> still unexplained environmental ray.
```

Thus the near projective stability seen in Gate 580 is mostly expected from common multiplicative running. The remaining mystery is why the charged-lepton projective ray is already so close to the Koide cone.

## Required missing theorem

Promotion beyond environmental seal would require:

```text
full charged-lepton matrix/flavor-threshold RGE with uncertainty control,
plus a native root-trace or absolute-Dirac observable
if the Koide projective ray is to be promoted beyond an environmental seal.
```

Gate 352 remains binding: no native root-trace operator is certified.

## Verdict

```text
PASS_GATE580_KOIDE_TRANSPORT_VECTOR_INHERITED
PASS_KOIDE_COORDINATE_BETA_FUNCTIONS_DERIVED_FROM_DIAGONAL_YUKAWA_RATES
PASS_COMMON_MULTIPLICATIVE_CHARGED_LEPTON_RUNNING_CANCELS_PROJECTIVE_MOTION
PASS_KOIDE_COORDINATE_BETA_COMPUTED_AT_MZ
PASS_KOIDE_COORDINATE_BETA_COMPUTED_AT_LAMBDA12
PASS_PROJECTIVE_MOTION_SOURCED_ONLY_BY_FAMILY_DEPENDENT_RATE_SPLITTING_IN_V1
CONDITIONAL_SUPPORT_LOCAL_THETA_BETA_POINTS_TOWARD_KOIDE_CONE_AT_RUNTIME_ENDPOINTS
PASS_LOCAL_PHI_BETA_IS_SMALL_IN_V1
FAILED_ROUTE_KOIDE_CONE_NOT_RG_INVARIANT_IN_V1_COORDINATE_BETA
FAILED_ROUTE_KOIDE_CONE_ATTRACTOR_NOT_CERTIFIED_BY_V1_BETA_FUNCTION
FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_OR_ABSOLUTE_DIRAC_KOIDE_BETA_OPERATOR
FAILED_ROUTE_NO_ASHA_NATIVE_CHARGED_LEPTON_FLAVOR_DERIVATION_FROM_KOIDE_BETA
FIREWALL_PRESERVED_NO_NEW_FLAVOR_CARRIER_OR_SELECTOR_INTRODUCED
FIREWALL_PRESERVED_CHARGED_LEPTON_INPUTS_REMAIN_HISTORY_ENDPOINT_DATA
FIREWALL_PRESERVED_KOIDE_BETA_DOES_NOT_DERIVE_TEXTURE_CKM_PMNS_OR_GENERATIONS
FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING
FIREWALL_PRESERVED_GATE581_KOIDE_COORDINATE_BETA_BOUNDARY
```
