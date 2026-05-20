# Gate 584 — Koide Wall-Offset One-Parameter Ratio Closure Audit

Gate 584 continues the charged-lepton environmental geometry sequence from Gates 577–583.  Gate 583 identified the charged-lepton hierarchy as a near-electron-wall point in the canonical positive Koide chamber.  Gate 584 asks the next narrower question:

```text
If R=1 is imposed exactly, does one charged-lepton square-root ratio determine the wall offset epsilon and predict the other independent ratio?
```

This is a ratio-closure audit.  It is not a native ASHA derivation of `epsilon_e`, Koide, charged-lepton masses, CKM/PMNS, or generation hierarchy.

## Exact R=1 wall model

In the canonical `(e,mu,tau)` chamber, write

```text
delta = 135° - epsilon,
0° < epsilon < 30°.
```

On the exact Koide circle `R=1`, the normalized square-root components are:

```text
E(epsilon) = x_e/A
            = 1 - cos(epsilon) + sin(epsilon)

M(epsilon) = x_mu/A
            = 1 - ((sqrt(3)-1)/2) cos(epsilon)
                - ((sqrt(3)+1)/2) sin(epsilon)

T(epsilon) = x_tau/A
            = 1 + ((sqrt(3)+1)/2) cos(epsilon)
                + ((sqrt(3)-1)/2) sin(epsilon).
```

Thus the two independent square-root ratios are functions of one parameter:

```text
r_e_mu(epsilon)  = E(epsilon)/M(epsilon)
r_mu_tau(epsilon)= M(epsilon)/T(epsilon).
```

So if exact `R=1` and the chamber are accepted, one ratio should solve `epsilon`, and the other ratio should be predicted.

Status: `PASS_EXACT_KOIDE_R1_WALL_RATIO_MODEL_DEFINED`.

## M_Z ratio closure

Observed runtime ratios from Gate 583 at `M_Z` are:

```text
x_e/x_mu   = 0.0695437394192847
x_mu/x_tau = 0.24385145943446
R(M_Z)     = 0.999990767173456
```

Solving the exact `R=1` wall model from `x_e/x_mu` gives:

```text
epsilon_from_e_mu = 2.26761458653473°
```

Then the model predicts:

```text
predicted x_mu/x_tau = 0.243843978487768
observed  x_mu/x_tau = 0.24385145943446
residual              = -7.48094669203447e-06
relative residual     = -3.06782937013552e-05
```

The squared mass-ratio prediction is:

```text
predicted (m_mu/m_tau) = 0.0594598858447431
observed  (m_mu/m_tau) = 0.0594635342683162
residual                = -3.64842357304662e-06
```

Conversely, solving from `x_mu/x_tau` gives:

```text
epsilon_from_mu_tau = 2.26689961821991°
predicted x_e/x_mu  = 0.0695193793739238
observed  x_e/x_mu  = 0.0695437394192847
residual             = -2.43600453609522e-05
```

Both closures pass the v1 tolerance.  The small residuals are expected because the runtime point has `R=0.999990767173456`, not exactly `R=1`.

Status: `PASS_ELECTRON_MUON_RATIO_SOLVES_UNIQUE_EPSILON_IN_POSITIVE_CHAMBER`; `PASS_ELECTRON_MUON_SOLVED_EPSILON_PREDICTS_MUON_TAU_RATIO`; `PASS_MUON_TAU_RATIO_SOLVES_UNIQUE_EPSILON_IN_POSITIVE_CHAMBER`; `PASS_MUON_TAU_SOLVED_EPSILON_PREDICTS_ELECTRON_MUON_RATIO`.

## Lambda_12 ratio closure

At `Lambda_12`, Gate 583 gives:

```text
x_e/x_mu   = 0.0695437358909112
x_mu/x_tau = 0.243847972677041
R          = 0.999995071771431
```

Solving from `x_e/x_mu` gives:

```text
epsilon_from_e_mu = 2.26761448298036°
```

and predicts:

```text
predicted x_mu/x_tau = 0.243843979571286
observed  x_mu/x_tau = 0.243847972677041
residual              = -3.99310575513456e-06
relative residual     = -1.63753904176318e-05
```

The closure residual improves relative to the `M_Z` endpoint because the `Lambda_12` point is closer to `R=1`.

Status: `PASS_ONE_PARAMETER_RATIO_CLOSURE_CERTIFIED_IN_CHARGED_LEPTON_SECTOR`; `PASS_RATIO_CLOSURE_STABLE_BETWEEN_MZ_AND_LAMBDA12_IN_V1`.

## Interpretation

Gate 584 verifies the most important consequence of the chamber-wall picture:

```text
exact Koide circle + canonical chamber + one ratio
  -> epsilon_e
  -> the other ratio.
```

The charged-lepton hierarchy is therefore one-parameter inside the exact Koide wall model.  In compressed form:

```text
Y_e -> A, R≈1, epsilon_e, chamber
```

and if `R=1` is imposed:

```text
Y_e -> A, epsilon_e, chamber.
```

The remaining small parameter is still:

```text
epsilon_e ≈ 2.26718°.
```

This is a genuine environmental compression, not a native law-space derivation.

Status: `CONDITIONAL_SUPPORT_EXACT_R1_WALL_MODEL_REDUCES_TWO_RATIOS_TO_ONE_EPSILON`; `CONDITIONAL_SUPPORT_RATIO_RESIDUALS_CONTROLLED_BY_R_MINUS_ONE_AND_ENDPOINT_PRECISION`.

## Quark comparison

Gate 584 inherits the Gate 583 quark audit:

```text
R_up   = 1.27683615501823
R_down = 1.10716260048739
```

Because neither sector is on the `R=1` Koide circle in v1, the one-parameter Koide wall-ratio closure is not certified for quarks.

Status: `FAILED_ROUTE_NO_QUARK_ONE_PARAMETER_KOIDE_WALL_RATIO_CLOSURE_IN_V1`.

## Firewalls

Gate 584 does not derive:

- Koide itself;
- the wall offset `epsilon_e`;
- charged-lepton masses;
- Yukawa eigenvalues;
- CKM or PMNS data;
- flavor texture;
- generation hierarchy;
- a new ASHA carrier or selector.

Gate 352 remains binding: no native root-trace, absolute-Dirac, or circulant phase-selection operator is supplied.

Status: `FAILED_ROUTE_EPSILON_NOT_DERIVED_NATIVE_FROM_RATIO_CLOSURE`; `FAILED_ROUTE_NO_NATIVE_ROOT_TRACE_OR_CIRCULANT_RATIO_OPERATOR`; `FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING`; `FIREWALL_PRESERVED_GATE584_RATIO_CLOSURE_BOUNDARY`.

## Final verdict

Gate 584 certifies the bridge-layer environmental seal:

```text
ChargedLeptonKoideWallOffsetOneParameterRatioSeal
```

The exact `R=1` wall model turns the two charged-lepton hierarchy ratios into a one-parameter closure.  One ratio solves the wall offset and predicts the other with residuals at `10^-5` root-ratio scale in v1.  The hierarchy is therefore geometrically compressed to a chamber-wall distance, while the distance itself remains the unresolved environmental seal.

Status: `PASS_ONE_PARAMETER_RATIO_CLOSURE_CERTIFIED_IN_CHARGED_LEPTON_SECTOR`; `CONDITIONAL_SUPPORT_EXACT_R1_WALL_MODEL_REDUCES_TWO_RATIOS_TO_ONE_EPSILON`; `FAILED_ROUTE_EPSILON_NOT_DERIVED_NATIVE_FROM_RATIO_CLOSURE`; `FIREWALL_PRESERVED_GATE584_RATIO_CLOSURE_BOUNDARY`.
