# Gate 764 — CP1 Gauge-Orbit Demotion and Radial Rank-Invariant HistoryLoop Audit

## Purpose

Gate 764 follows Gate 763 by asking whether the lack of a native `CP1` selector is a scalar-runtime defect or the expected behavior of a gauge orbit.

Gate 763 established:

```text
K7+_J(n) ~= C^2
CP1 = P(C^2)
Pi_vac_C remains ComplexVacuumLineSeal
no native CP1 selector functional is certified
```

Gate 764 audits the alternative interpretation:

```text
Pi_vac_C is a gauge/vacuum-orientation representative for scalar-runtime numerics,
while L_Hopf depends only on the rank-one real radial event weight.
```

This is a gauge-orbit and rank-invariance audit only. It does not derive electroweak symmetry breaking, scalar runtime lambda, Higgs mass, pole mass, Yukawa operators, CKM/PMNS, flavor hierarchy, or a native HistoryLoopUnit theorem.

## Implemented package

```text
pkg/bridge/generation2cp1gaugeorbitdemotionandradialrankinvarianthistoryloopaudit
```

Registered theorem:

```text
generation2cp1gaugeorbitdemotionandradialrankinvarianthistoryloopaudit.Generation2CP1GaugeOrbitDemotionAndRadialRankInvariantHistoryLoopAuditTheorem()
```

## Inherited Gate763 firewall

Gate 763 supplied the arena:

```text
K7+_J(n) ~= C^2
CP1 = P(C^2)
Pi_vac_C = complex vacuum line / CP1 point
P_rad = gauge-fixed real radial representative inside Pi_vac_C
rho_plus = I_K7+ / 4
L_Hopf = (1/(2*pi)) Tr(rho_plus P_rad)
```

and the selector-functional verdict:

```text
No native CP1 selector functional is certified.
A nonconstant CP1 functional would require a Hermitian SU(2) socket axis.
```

## CP1 as an SU(2) orbit

For the complex doublet carrier:

```text
K7+_J(n) ~= C^2,
```

the internal socket geometry has:

```text
CP1 ~= SU(2)/U(1).
```

The `SU(2)` action is transitive on complex lines. Therefore a fully `SU(2)`-invariant scalar functional is constant on `CP1` and cannot select a line. A line can be selected only if an anisotropic Hermitian axis is supplied.

Gate 764 therefore interprets the Gate 763 selector failure as expected gauge-orbit behavior at the internal socket layer, with the following firewall:

```text
FAILED_ROUTE_INTERNAL_C_NOT_CERTIFIED_AS_PHYSICAL_SU2L
```

The demotion is conditional on the electroweak gauge airlock. It is not promoted into a physical `SU(2)_L` theorem.

## Gauge-orbit demotion

Gate 764 reclassifies:

```text
Pi_vac_C in CP1
```

as:

```text
gauge/vacuum-orientation representative data for scalar-runtime numerics,
unless a physical anisotropy is supplied.
```

Thus the absence of a CP1 selector is not counted as a scalar-runtime numerical failure. It remains a physical vacuum-orientation / electroweak symmetry-breaking gap if the CP1 point is promoted to physical orientation.

## Radial rank-invariance audit

For any real rank-one projector `P_rad` inside `K7+`:

```text
Tr(rho_plus P_rad)
=
Tr((I_K7+/4)P_rad)
=
rank(P_rad)/4
=
1/4.
```

Therefore:

```text
L_Hopf
=
(1/(2*pi)) Tr(rho_plus P_rad)
=
(1/(2*pi))(1/4)
=
1/(8*pi).
```

This depends on:

```text
real radial event rank = 1
K7+ real dimension = 4
phase-loop payoff = 1/(2*pi)
```

not on the `CP1` position.

## Complex-line versus radial-event distinction

The full complex vacuum line has real rank two:

```text
rank_R(Pi_vac_C)=2.
```

So:

```text
Tr(rho_plus Pi_vac_C)=1/2
```

and:

```text
(1/(2*pi))Tr(rho_plus Pi_vac_C)=1/(4*pi).
```

This is not the active HistoryLoop unit.

The active unit uses the real radial amplitude event:

```text
rank_R(P_rad)=1,
Tr(rho_plus P_rad)=1/4,
L_Hopf=1/(8*pi).
```

Therefore the scalar-runtime source question is not primarily:

```text
Which CP1 point?
```

but:

```text
Why does HistoryLoop transport use a real rank-one radial amplitude event
rather than the full complex vacuum line?
```

That question remains open.

## Updated missing-object hierarchy

Before Gate 764:

```text
ComplexVacuumLineSeal looked like the primary missing scalar source.
```

After Gate 764:

```text
For scalar-runtime coefficient:
  CP1 location is gauge-representative data.
  The active source is rank-one radial event type.

Remaining scalar source problem:
  Why real rank-one radial amplitude event with phase-loop payoff 1/(2*pi)?

Remaining physical Higgs problem:
  native electroweak symmetry breaking / vacuum-orbit theorem.
```

## Verdict

Gate 764 records:

```text
PASS_GATE763_CP1_SELECTOR_FIREWALL_INHERITED
PASS_CP1_AS_SU2_ORBIT_AUDITED
PASS_GAUGE_ORBIT_DEMOTION_DEFINED
PASS_RADIAL_RANK_INVARIANCE_COMPUTED
PASS_COMPLEX_LINE_VERSUS_RADIAL_EVENT_DISTINCTION_AUDITED
PASS_UPDATED_MISSING_OBJECT_HIERARCHY_RECORDED
PASS_PHYSICAL_FIREWALLS_ENFORCED
CONDITIONAL_SUPPORT_CP1_SELECTOR_ABSENCE_IS_EXPECTED_FOR_GAUGE_ORBIT
CONDITIONAL_SUPPORT_PI_VAC_C_IS_GAUGE_REPRESENTATIVE_NOT_SCALAR_NUMERICAL_SOURCE
CONDITIONAL_SUPPORT_L_HOPF_DEPENDS_ON_RANK_ONE_RADIAL_EVENT_NOT_CP1_POSITION
CONDITIONAL_SUPPORT_NEXT_SCALAR_SOURCE_TARGET_IS_RADIAL_EVENT_TYPE_SELECTION
FAILED_ROUTE_INTERNAL_C_NOT_CERTIFIED_AS_PHYSICAL_SU2L
FAILED_ROUTE_NO_NATIVE_REASON_HISTORYLOOP_USES_REAL_RADIAL_EVENT
FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM
FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM
FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM
FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM
FIREWALL_PRESERVED_GATE764_CP1_GAUGE_ORBIT_RADIAL_RANK_BOUNDARY
```

## Final interpretation

Gate 764 does not derive the Higgs vacuum or the HistoryLoop unit. It sharpens the source problem.

The CP1 point is demoted from scalar-runtime source to internal gauge-orbit representative. The active scalar coefficient is controlled by a rank-one real radial event:

```text
L_Hopf = (1/(2*pi)) * (1/4) = 1/(8*pi).
```

The next scalar-source pressure point is therefore:

```text
Why the real rank-one radial event type?
```

The separate physical Higgs pressure point remains:

```text
native electroweak symmetry breaking / vacuum-orbit theorem.
```
