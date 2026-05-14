# Gate 287 Registry Audit — Topological Action Variational Principle / `S_top = 8π²` Boundary Selector

## Gate

**Gate 287 — Topological Action Variational Principle / `S_top = 8π²` Boundary Selector Audit**

This gate implements the corrected dynamical pivot proposed after Gate 286. Instead of continuing to search for continuum Hopf/Chern-Simons structures, it promotes the exact finite topological action

```text
S_top = 8π²
```

to a proposed global spectral-action boundary constraint and audits whether the variational equations select the physical realization: the Gate-275 amplitude branch, physical `J`, cutoff moments, Higgs ratio, or B-gap instanton action.

## Inputs inherited

### Gate 286

Gate 286 established the correct finite NCG calculus route:

```text
δ(a) = [D_F,a]
A = Σ a_i[D_F,b_i]
F = [D_F,A] + A²
S_finite ≈ Tr(F†F)
```

It also proved that the local quaternionic diagnostic is non-vacuous but does not derive a non-trivial saddle or inverse-`B_gap` action.

### Gate 275 / 276

The scalar-Morita bridge supplies the finite moment model:

```text
κ_C : κ_Q = 1 : 3
r = |y/x|²
X = |x|²
Tr(D_F²) proxy = X(1+3r)
Tr(D_F⁴) proxy = X²(1+3r²)
```

and the exact contact shape constraint:

```text
(1+3r²)/(1+3r)² = 1197/4624
```

which gives:

```text
3099r² - 7182r + 3427 = 0
r = (3591 ± 136√123)/3099
```

Numerically:

```text
r_+ ≈ 1.645470463011191
|y/x|_+ ≈ 1.282758926303454

r_- ≈ 0.672051318208557
|y/x|_- ≈ 0.819787361581378
```

## New Gate-287 formalization

The proposed total spectral-action boundary is written as:

```text
S_total = F4 a0(D_F) + F2 a2(D_F) + F0 a4(D_F) = 8π²
```

where:

```text
F4 = f4 Λ⁴
F2 = f2 Λ²
F0 = f0
```

Using the finite moment proxy:

```text
S(r,X) = F4 a0 + F2 X(1+3r) + F0 X²(1+3r²)
```

The stationarity equation is:

```text
∂S/∂r = 3F2 X + 6F0 X²r
```

so:

```text
r_* = -F2/(2F0 X)
```

provided `F0·X ≠ 0`.

## Main finding

`S_top = 8π²` is a valid exact finite boundary datum, but it is not by itself a complete dynamical selector.

With free cutoff moments and unknown absolute Dirac scale, the stationarity equation can be tuned to any positive `r` if signs/ratios are left free. With standard positive cutoff moments, the stationarity equation does not select a positive nonzero `r`. Therefore the variational equation does not natively select `r_+` or `r_-`.

The scalar shape function itself has derivative:

```text
d/dr [(1+3r²)/(1+3r)²] = 6(r-1)/(1+3r)³
```

so the shape extremum occurs at:

```text
r = 1
λ = 1/4
```

not at the two Gate-275 branches. The branches solve the fixed contact-shape equation; they are not extrema of the shape functional.

## Constraint rank audit

The native equations currently available are:

```text
F4 a0 + F2 Tr(D_F²) + F0 Tr(D_F⁴) = 8π²
(1+3r²)/(1+3r)² = 1197/4624
∂S/∂r = 0, only after choosing admissible cutoff signs and scale
```

The unresolved unknowns include:

```text
r
X = |x|²
F0
F2
F4
a0 normalization
physical J
chiral/hypercharge representation
B_gap insertion/coupling map
```

Thus the proposed variational principle is underdetermined unless future gates derive additional constraints on cutoff moments, physical `D_F`, field normalizations, and the completed spectral triple.

## `J` as a consequence

The gate audits the proposal that physical `J` should be derived as the antiunitary symmetry preserving the selected action extremum.

This is structurally valid as a future criterion, but Gate 287 cannot execute it because no unique extremum/vacuum is selected. Therefore:

```text
J D_F(r_*) = D_F(r_*) J
Jγ = -γJ
```

remain target equations, not derived theorems.

## Four-over-pi test

The gate also audits whether the constrained action produces:

```text
S_inst = (4/π)/B_gap
```

It does not. `S_top` can encode the exact `4/π` volume identity inherited from Gate 283, but Gate 287 still lacks:

```text
B_gap as inverse coupling
non-perturbative finite spectral sector
saddle action gap ΔS
cutoff moment normalization
```

Therefore the `B_gap` instanton law remains blocked.

## Status ledger

```text
CONDITIONAL_SUPPORT_GATE286_NCG_SADDLE_BARRIER_INHERITED
CONDITIONAL_SUPPORT_S_TOP_BOUNDARY_ACTION_CONSTRAINT_FORMALIZED
CONDITIONAL_SUPPORT_SCALAR_MORITA_MOMENT_MODEL_INHERITED
CONDITIONAL_SUPPORT_VARIATIONAL_EQUATIONS_DERIVED
CONDITIONAL_SUPPORT_SHAPE_EXTREMUM_AUDIT_COMPLETED
CONDITIONAL_SUPPORT_CUTOFF_MOMENT_UNDERDETERMINATION_PROVED
CONDITIONAL_SUPPORT_PATH_B_C_CONVERGENCE_RECORDED
CONDITIONAL_SUPPORT_TOPOLOGICAL_ACTION_VARIATIONAL_FIREWALLS_PRESERVED
FAILED_ROUTE_S_TOP_VARIATION_DOES_NOT_SELECT_R_BRANCH
FAILED_ROUTE_VARIATIONAL_PRINCIPLE_UNDERDETERMINED_WITH_FREE_CUTOFF_MOMENTS
FAILED_ROUTE_PHYSICAL_J_NOT_DERIVED_AS_EXTREMUM_SYMMETRY
FAILED_ROUTE_CUTOFF_MOMENT_RATIOS_NOT_EXTRACTED
FAILED_ROUTE_FOUR_OVER_PI_INSTANTON_NOT_DERIVED_BY_VARIATION
FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED
```

## Final verdict

Gate 287 validates the new top-down idea as a legitimate future path:

```text
S_top = 8π²
```

is not merely a passive diagnostic. It can be formulated as a global finite spectral-action boundary constraint.

But the constraint is not yet a complete variational principle. It does not select `r_+` versus `r_-`, does not derive physical `J`, does not extract cutoff moment ratios, and does not produce the `4/π / B_gap` instanton action.

The correct next obligation is to derive either:

1. a native cutoff-moment/normalization theorem, or
2. a formal `TopologicalActionMomentSeal` specifying admissible spectral-action moments, or
3. the completed physical finite spectral triple whose symmetries reduce the variational degrees of freedom.

Until then, both Path B and Path C remain firewalled.
