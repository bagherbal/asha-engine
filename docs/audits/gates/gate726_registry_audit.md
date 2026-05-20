# Gate 726 — Radial-Phase Hopf Fiber and Angular Complement Decomposition Audit

## Purpose

Gate 726 follows Gate 725 by auditing the finer geometry of the supplied radial projector once the twistor selector `n` is also supplied. Gate 725 gave the split:

```text
K7+ = K_rad ⊕ K_ang
```

with dimensions `1+3`. Gate 726 asks whether the chosen complex structure `J_H(n)` further splits the angular complement into a Hopf phase direction and a transverse pair.

## Registered theorem

```text
pkg/bridge/generation2radialphasehopffiberandangularcomplementdecompositionaudit
```

```text
generation2radialphasehopffiberandangularcomplementdecompositionaudit.Generation2RadialPhaseHopfFiberAndAngularComplementDecompositionAuditTheorem()
```

## Radial phase direction

Choose a unit vector `v_rad` spanning:

```text
K_rad = Im(P_rad).
```

Since `J_H(n)^2=-I` and `J_H(n)` is skew-orthogonal, one has:

```text
<v_rad, J_H(n)v_rad> = 0.
```

Therefore:

```text
J_H(n)v_rad ∈ K_ang.
```

Gate 726 defines:

```text
K_phase = span(J_H(n)v_rad)
P_phase = projector onto K_phase
```

with:

```text
rank(P_phase)=1
P_phase P_rad=0.
```

## Radial/phase/transverse decomposition

Define:

```text
P_trans = I_K7+ - P_rad - P_phase.
```

Then:

```text
rank(P_rad)=1
rank(P_phase)=1
rank(P_trans)=2
```

and:

```text
K7+ = K_rad ⊕ K_phase ⊕ K_trans
4 = 1 + 1 + 2.
```

Equivalently:

```text
K_ang = K_phase ⊕ K_trans
3 = 1 + 2.
```

## Hopf-fiber phase loop

The selected complex structure gives the circle orbit through the radial vector:

```text
v_rad(theta)=exp(theta J_H(n))v_rad.
```

This orbit lies in:

```text
span(v_rad,J_H(n)v_rad).
```

Thus Gate 726 sharpens Gate 724: the candidate `1/(2*pi)` phase-loop payoff can be read as the normalized Hopf-fiber unit through the supplied radial event, not merely an abstract phase line.

## Event weights

Under:

```text
rho_plus = I_K7+ / 4,
```

Gate 726 computes:

```text
Pr(radial)       = 1/4
Pr(phase)        = 1/4
Pr(transverse)   = 1/2
```

The active HistoryLoopUnit candidate remains:

```text
L = Pr(radial) * 1/(2*pi) = 1/(8*pi).
```

It uses the radial event weight, not the full phase-direction event weight.

## U(2) orbit interpretation

Gate 725 recorded:

```text
dim U(2)=4
stabilizer dimension=1
orbit dimension=3.
```

Gate 726 decomposes that local angular orbit as:

```text
3 = 1 phase-fiber direction + 2 projective/transverse directions.
```

This is the internal Hopf-style pattern:

```text
S^3 orbit -> CP1 base with S^1 fiber.
```

## Selector dependence

This decomposition requires both missing selectors:

```text
n      -> supplies J_H(n), hence the phase direction
P_rad  -> supplies the radial line, hence the Hopf fiber through it
```

Without `n`, there is no phase direction. Without `P_rad`, there is no radial event.

## Firewalls

Gate 726 does not certify physical Goldstone bosons, electroweak symmetry breaking, physical time, RG scale, a native HistoryLoopUnit source theorem, a Higgs mass theorem, or a Yukawa theorem.

## Verdict

```text
PASS_GATE725_RADIAL_GOLDSTONE_ORBIT_INHERITED
PASS_PHASE_DIRECTION_FROM_RADIAL_LINE_DEFINED
PASS_RADIAL_PHASE_TRANSVERSE_DECOMPOSITION_COMPUTED
PASS_HOPF_FIBER_PHASE_LOOP_AUDITED
PASS_RADIAL_PHASE_TRANSVERSE_EVENT_WEIGHTS_COMPUTED
PASS_U2_ORBIT_HOPF_STRUCTURE_AUDITED
PASS_SELECTOR_DEPENDENCE_AUDITED
PASS_PHYSICAL_FIREWALL_ENFORCED
CONDITIONAL_SUPPORT_K7_PLUS_DECOMPOSES_AS_RADIAL_PHASE_TRANSVERSE_AFTER_N_AND_P_RAD
CONDITIONAL_SUPPORT_ANGULAR_COMPLEMENT_HAS_1_PLUS_2_HOPF_STRUCTURE
CONDITIONAL_SUPPORT_ONE_OVER_TWO_PI_IS_PHASE_LOOP_UNIT_ON_RADIAL_HOPF_FIBER
CONDITIONAL_SUPPORT_L_EQUALS_RADIAL_EVENT_WEIGHT_TIMES_HOPF_PHASE_UNIT
FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR
FAILED_ROUTE_NO_NATIVE_TWISTOR_SELECTOR_N
FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM
FAILED_ROUTE_NO_PHYSICAL_GOLDSTONE_IDENTIFICATION
FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM
FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM
FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM
FIREWALL_PRESERVED_GATE726_RADIAL_PHASE_HOPF_DECOMPOSITION_BOUNDARY
```
