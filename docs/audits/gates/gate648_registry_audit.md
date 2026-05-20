# Gate 648 — Cubic Slot Multiplicity versus Negative-Sector Dimension Audit

## Purpose

Gate 647 showed that the admissible `S_K`-twisted Hitchin metric has the finite block ray

```text
g_twist ∝ P_+ - 3P_-.
```

Gate 647 also observed that `3 = dim(K_7^-)`.  Gate 648 corrects the source typing: the contraction ledger directly exposes **three ordered cubic Hitchin channels**.  Since the Hitchin metric is cubic, the immediate coefficient source may be the cubic slot count rather than a general negative-sector dimension theorem.

This is an internal tensor-source audit only.  It does not derive split-G2, boundary stress, scalar/flavor transport, physical spacetime, Higgs mass, CKM/PMNS, gauge unification, or a native `7/72` theorem.

## Gate 647 inheritance

Inherited finite data:

```text
K_7 = K_7^+ ⊕ K_7^-

dim K_7^+ = p = 4
dim K_7^- = q = 3

cubic Hitchin degree = 3

g_twist ∝ P_+ - 3P_-.
```

The Gate 647 ordered ledger contains one positive channel

```text
Omega++- × Omega++- × Omega++-
```

and three negative channels

```text
Omega++- × Omega++- × Omega---
Omega++- × Omega--- × Omega++-
Omega--- × Omega++- × Omega++-
```

## Per-direction and total-trace audit

Gate 648 separates per-direction coefficients from total sector trace.

For each route,

```text
c_+ = (1/p) Tr(P_+ g_twist P_+)
c_- = (1/q) Tr(P_- g_twist P_-)
```

satisfies

```text
c_- / c_+ = -3.
```

But the total sector trace ratio is

```text
Tr(P_- g_twist P_-) / Tr(P_+ g_twist P_+)
= -(3q)/p
= -9/4.
```

This separates the slot coefficient `-3` from the total size of the negative sector.

## Ordered-slot contribution audit

For each of the repeated routes

```text
omega_1_alt
omega_2_alt
omega_B_alt
```

the finite ledger shows that each negative ordered channel contributes exactly one unit of negative block weight relative to the positive unit:

```text
Omega++- × Omega++- × Omega--- -> -1 unit
Omega++- × Omega--- × Omega++- -> -1 unit
Omega--- × Omega++- × Omega++- -> -1 unit
```

Therefore

```text
-1 - 1 - 1 = -3.
```

Gate 648 classifies this as the direct certified finite source:

```text
CONDITIONAL_SUPPORT_MINUS_THREE_ARISES_FROM_CUBIC_SLOT_MULTIPLICITY.
```

## Negative-index contribution audit

Each negative basis direction receives the same three unit channel contributions.  The negative-sector dimension `q=3` multiplies the total negative trace, but the per-direction coefficient is sourced by the ordered cubic slot ledger.

Thus the finite result is more precisely:

```text
per negative direction: -3 units
negative sector total:  q × (-3 units).
```

## Dimension formula versus slot formula

Gate 646 used the candidate general form

```text
G_dim ∝ P_+ - qP_-
||G_dim||^2 = p + q^3.
```

Gate 648 records the more conservative slot-supported form

```text
G_slot ∝ P_+ - 3P_-
||G_slot||^2 = p + 9q.
```

For ASHA,

```text
p = 4,
q = 3,
```

so the two formulae coincide:

```text
p + q^3 = 4 + 27 = 31,
p + 9q = 4 + 27 = 31.
```

The final normalized ray cannot distinguish the two.  The contribution ledger selects the slot-source interpretation as the directly witnessed mechanism.

## Synthetic/ablative diagnostics

Gate 648 records diagnostic-only source probes:

```text
remove one ordered negative channel  -> coefficient -2
remove two ordered negative channels -> coefficient -1
remove Omega--- family              -> negative block source removed
remove one negative basis direction  -> total trace changes, per-surviving-direction coefficient remains -3
```

These are not promoted to native deformations.  They are diagnostic evidence that the observed `-3` is slot-driven at the finite ledger level.

## Refined theorem target

The old target was:

```text
HitchinMetric(Omega_twist) ∝ P_+ - qP_-,
q = dim(K_7^-).
```

Gate 648 refines the current honest target to:

```text
HitchinMetric(Omega_twist) ∝ P_+ - 3P_-
```

because the cubic Hitchin contraction has one positive `AAA` channel and three ordered `AAB` negative channels.  In this ASHA carrier,

```text
cubic degree = dim(K_7^-) = 3.
```

A future symbolic theorem may explain why this coincidence is forced, but Gate 648 does not claim it.

## Verdict

```text
PASS_GATE647_CONTRACTION_LEDGER_INHERITED
PASS_PER_DIRECTION_AND_TOTAL_TRACE_AUDIT_COMPUTED
PASS_ORDERED_SLOT_CONTRIBUTIONS_COMPUTED
PASS_NEGATIVE_INDEX_CONTRIBUTIONS_COMPUTED
PASS_DIMENSION_VERSUS_SLOT_FORMULA_DISAMBIGUATION_AUDITED
PASS_SYNTHETIC_ABLATIVE_DIAGNOSTICS_COMPUTED
CONDITIONAL_SUPPORT_MINUS_THREE_ARISES_FROM_CUBIC_SLOT_MULTIPLICITY
CONDITIONAL_SUPPORT_DIM_K7_MINUS_EQUALS_CUBIC_DEGREE_IN_ASHA_CARRIER
CONDITIONAL_SUPPORT_HITCHIN_MULTIPLICITY_THEOREM_REFINED
FAILED_ROUTE_NO_GENERAL_P_Q_DIMENSION_THEOREM_YET
FAILED_ROUTE_NO_FULL_SYMBOLIC_HITCHIN_MULTIPLICITY_THEOREM
FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE
FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT
FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM
FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM
FAILED_ROUTE_HITCHIN_CONTRACTION_METRIC_IS_NOT_PHYSICAL_METRIC
FIREWALL_PRESERVED_GATE648_CUBIC_SLOT_MULTIPLICITY_BOUNDARY
```

## Final classification

Gate 648 prevents over-generalization.  The safest current statement is:

```text
g_twist ∝ P_+ - 3P_-
```

where `-3` is directly witnessed as the three ordered cubic Hitchin negative channels.  The equality

```text
3 = dim(K_7^-)
```

is true and important, but in this gate it is recorded as an ASHA carrier coincidence with the cubic Hitchin degree, not as a general `p,q` theorem.
