# Gate 211 Registry Audit

## Gate

**Gate 211 — Two-threshold rational lattice viability filter / scale-ordered Landau safety audit**

Package:

```text
pkg/bridge/twothresholdviability
```

Registry theorem:

```text
BRIDGE-TWO-THRESHOLD-RATIONAL-LATTICE-VIABILITY-FILTER
```

Status:

```text
CONDITIONAL_VIABLE_TWO_THRESHOLD_LATTICE
```

This is a conditional phenomenological viability result, not a finite-core prediction. The Z-pole ledger remains quarantined, the `LeptoquarkDynamicsSeal` remains active, and no observed mass is used as finite algebraic data.

---

## Purpose

Gate 210 proved that a single rational threshold row cannot exactly close the mismatch triangle. Gate 211 uses the critical dimension-counting pivot: two independent rational threshold rows plus `b_SM` generically span `R^3`, so exact closure is a 3×3 linear solve. The scientific content is therefore the physical filter applied after solving.

The solved u-space system is:

```text
A_i = u_target + [(b_i + Δb_i^(1) + Δb_i^(2))/(8π²)] L*
      - [Δb_i^(1)/(8π²)] L_B1 - [Δb_i^(2)/(8π²)] L_B2
```

where `A_i = α_i^-1/(4π)` is inherited from the quarantined Gate-200 Z-pole ledger.

---

## Inputs and firewalls

```text
Gate-200 quarantined Z-pole empirical comparison ledger MZ=91.1876 α^-1=(59.02154694,29.5857551,8.48176420696) u=(4.69678547222,2.35435958464,0.674957349838) bSM=(4.1,-3.16666666667,-7) planck=1.2209e+19 Lbound=37.8 quarantined=true finiteUse=false universalAllowed=false
```

Generator basis:

```text
sourceUnique=158 safe=108 expectedGate210=108 anomaly=108 seal=108 zeroExcluded=true universalRow=false realCoeff=false inheritsGate210=true
```

The filter inherits the 108 anomaly-safe, leptoquark-compatible nonzero rational rows from Gate 210. It inserts no universal beta row and no arbitrary real coefficient.

---

## Boundary target audit

### u_topological

```text
target=u_topological u=1 orderedPairs=11556 pairIndependent=11350 invertible=11350 scaleOrdered=518 distinct=518 subPlanck=110 positive=518 noLandau=44 viable=44 dominant=scale-ordering (10832) best={target=u_topological u=1 rows=[Dirac fermion (1,3,Y=1) (1,3,Y=1); Dirac fermion (8,2,Y=1/2) (8,2,Y=1/2)] Δb1=(12/5,8/3,0) Δb2=(16/5,16/3,8) ΔbTot=(28/5,8,8) L=(L*=34.3263535,LB1=7.11786258,LB2=7.49883655) M=(M*=7.37363563e+16,MB1=112508.213,MB2=164679.341) totalBeta=(9.7,4.83333333333,1) closureU=4.44e-16 exactClosure=true minU=0.573255799 pole=none ordered=true distinct=true subPlanck=true positive=true noPole=true anomaly=true seal=true gate201=false contactMatch=false AF(SU2,SU3)=(false,false)}
```

Interpretation: quarantined instanton/topological branch.

Binding constraints:

| Constraint | Count |
|---|---:|
| `viable` | `44` |
| `scale-ordering` | `10832` |
| `sub-planck-bound` | `408` |
| `sub-planck-landau-pole` | `66` |
| `singular-or-dependent-3x3-system` | `206` |

### u_centroid

```text
target=u_centroid u=3.33 orderedPairs=11556 pairIndependent=11350 invertible=11350 scaleOrdered=0 distinct=0 subPlanck=0 positive=0 noLandau=0 viable=0 dominant=scale-ordering (11350) best={none}
```

Interpretation: Gate-200 SM-only mismatch-triangle centroid comparison branch.

Binding constraints:

| Constraint | Count |
|---|---:|
| `scale-ordering` | `11350` |
| `singular-or-dependent-3x3-system` | `206` |

---

## Viable solutions

The topological branch has `44` ordered viable pairs. Reverse orderings are listed separately because Gate 211 solves ordered threshold labels `(B1,B2)`. The first 20 ranked witnesses are stored in the theorem summary; all viable ordered pairs are listed below.

Ranking rule: fewest carriers, then smallest total `Δb` norm, then closeness of `M*` to the conventional `10^15–10^16 GeV` range. All entries have two carriers, exact linear closure, ordered distinct thresholds, `L* < 37.8`, positive couplings, no sub-Planck Landau pole, anomaly compatibility, and leptoquark-seal compatibility.

| # | Row 1 | Row 2 | Δb1 | Δb2 | LB1 | LB2 | L* | MB1 GeV | MB2 GeV | M* GeV | b_total | AF SU2/SU3 |
|---:|---|---|---|---|---:|---:|---:|---:|---:|---:|---|---|
| `1` | `(1,3,Y=1)` | `(8,2,Y=1/2)` | `(12/5,8/3,0)` | `(16/5,16/3,8)` | `7.11786` | `7.49884` | `34.3264` | `112508` | `164679` | `7.37364e+16` | `(9.7,4.83333333333,1)` | `false/false` |
| `2` | `(8,2,Y=1/2)` | `(1,3,Y=1)` | `(16/5,16/3,8)` | `(12/5,8/3,0)` | `7.49884` | `7.11786` | `34.3264` | `164679` | `112508` | `7.37364e+16` | `(9.7,4.83333333333,1)` | `false/false` |
| `3` | `(1,3,Y=0)` | `(8,2,Y=2/3)` | `(0,8/3,0)` | `(256/45,16/3,8)` | `6.86478` | `7.47859` | `34.1644` | `87352.1` | `161379` | `6.27103e+16` | `(9.78888888889,4.83333333333,1)` | `false/false` |
| `4` | `(8,2,Y=2/3)` | `(1,3,Y=0)` | `(256/45,16/3,8)` | `(0,8/3,0)` | `7.47859` | `6.86478` | `34.1644` | `161379` | `87352.1` | `6.27103e+16` | `(9.78888888889,4.83333333333,1)` | `false/false` |
| `5` | `(1,3,Y=1/6)` | `(8,2,Y=2/3)` | `(1/15,8/3,0)` | `(256/45,16/3,8)` | `6.55022` | `7.45343` | `33.9631` | `63776.7` | `157368` | `5.12752e+16` | `(9.85555555556,4.83333333333,1)` | `false/false` |
| `6` | `(8,2,Y=2/3)` | `(1,3,Y=1/6)` | `(256/45,16/3,8)` | `(1/15,8/3,0)` | `7.45343` | `6.55022` | `33.9631` | `157368` | `63776.7` | `5.12752e+16` | `(9.85555555556,4.83333333333,1)` | `false/false` |
| `7` | `(1,3,Y=1/3)` | `(8,2,Y=2/3)` | `(4/15,8/3,0)` | `(256/45,16/3,8)` | `5.59069` | `7.37666` | `33.349` | `24431` | `145740` | `2.77464e+16` | `(10.0555555556,4.83333333333,1)` | `false/false` |
| `8` | `(8,2,Y=2/3)` | `(1,3,Y=1/3)` | `(256/45,16/3,8)` | `(4/15,8/3,0)` | `7.37666` | `5.59069` | `33.349` | `145740` | `24431` | `2.77464e+16` | `(10.0555555556,4.83333333333,1)` | `false/false` |
| `9` | `(1,3,Y=1/2)` | `(8,2,Y=2/3)` | `(3/5,8/3,0)` | `(256/45,16/3,8)` | `3.93657` | `7.24433` | `32.2903` | `4672.7` | `127676` | `9.62603e+15` | `(10.3888888889,4.83333333333,1)` | `false/false` |
| `10` | `(8,2,Y=2/3)` | `(1,3,Y=1/2)` | `(256/45,16/3,8)` | `(3/5,8/3,0)` | `7.24433` | `3.93657` | `32.2903` | `127676` | `4672.7` | `9.62603e+15` | `(10.3888888889,4.83333333333,1)` | `false/false` |
| `11` | `(1,3,Y=2/3)` | `(8,2,Y=2/3)` | `(16/15,8/3,0)` | `(256/45,16/3,8)` | `1.49789` | `7.04924` | `30.7296` | `407.813` | `105046` | `2.02125e+15` | `(10.8555555556,4.83333333333,1)` | `false/false` |
| `12` | `(8,2,Y=2/3)` | `(1,3,Y=2/3)` | `(256/45,16/3,8)` | `(16/15,8/3,0)` | `7.04924` | `1.49789` | `30.7296` | `105046` | `407.813` | `2.02125e+15` | `(10.8555555556,4.83333333333,1)` | `false/false` |
| `13` | `(8,1,Y=1)` | `(8,3,Y=0)` | `(32/5,0,4)` | `(0,32/3,6)` | `10.9504` | `14.2137` | `34.4732` | `5.19552e+06` | `1.35796e+08` | `8.54027e+16` | `(10.5,7.5,3)` | `false/false` |
| `14` | `(8,3,Y=0)` | `(8,1,Y=1)` | `(0,32/3,6)` | `(32/5,0,4)` | `14.2137` | `10.9504` | `34.4732` | `1.35796e+08` | `5.19552e+06` | `8.54027e+16` | `(10.5,7.5,3)` | `false/false` |
| `15` | `(8,3,Y=1/3)` | `(8,2,Y=2/3)` | `(8/15,16/3,3)` | `(256/45,16/3,8)` | `17.4961` | `15.5124` | `37.7309` | `3.61734e+09` | `4.97624e+08` | `2.21941e+18` | `(10.3222222222,7.5,4)` | `false/false` |
| `16` | `(8,2,Y=2/3)` | `(8,3,Y=1/3)` | `(256/45,16/3,8)` | `(8/15,16/3,3)` | `15.5124` | `17.4961` | `37.7309` | `4.97624e+08` | `3.61734e+09` | `2.21941e+18` | `(10.3222222222,7.5,4)` | `false/false` |
| `17` | `(8,3,Y=1/2)` | `(8,2,Y=2/3)` | `(6/5,16/3,3)` | `(256/45,16/3,8)` | `15.4366` | `15.5746` | `36.3105` | `4.61294e+08` | `5.29527e+08` | `5.36293e+17` | `(10.9888888889,7.5,4)` | `false/false` |
| `18` | `(8,2,Y=2/3)` | `(8,3,Y=1/2)` | `(256/45,16/3,8)` | `(6/5,16/3,3)` | `15.5746` | `15.4366` | `36.3105` | `5.29527e+08` | `4.61294e+08` | `5.36293e+17` | `(10.9888888889,7.5,4)` | `false/false` |
| `19` | `(8,2,Y=2/3)` | `(8,3,Y=2/3)` | `(256/45,16/3,8)` | `(32/15,16/3,3)` | `15.6685` | `12.3247` | `34.1644` | `5.81656e+08` | `2.05346e+07` | `6.27103e+16` | `(11.9222222222,7.5,4)` | `false/false` |
| `20` | `(8,3,Y=2/3)` | `(8,2,Y=2/3)` | `(32/15,16/3,3)` | `(256/45,16/3,8)` | `12.3247` | `15.6685` | `34.1644` | `2.05346e+07` | `5.81656e+08` | `6.27103e+16` | `(11.9222222222,7.5,4)` | `false/false` |
| `21` | `(8,2,Y=1/2)` | `(8,3,Y=1)` | `(16/5,16/3,8)` | `(24/5,16/3,3)` | `15.7696` | `8.97299` | `31.8529` | `6.43555e+08` | `719208` | `6.21519e+15` | `(12.1,7.5,4)` | `false/false` |
| `22` | `(8,3,Y=1)` | `(8,2,Y=1/2)` | `(24/5,16/3,3)` | `(16/5,16/3,8)` | `8.97299` | `15.7696` | `31.8529` | `719208` | `6.43555e+08` | `6.21519e+15` | `(12.1,7.5,4)` | `false/false` |
| `23` | `(3,3,Y=0)` | `(8,2,Y=2/3)` | `(0,8,2)` | `(256/45,16/3,8)` | `25.2689` | `10.3222` | `35.8169` | `8.59202e+12` | `2.77201e+06` | `3.2737e+17` | `(9.78888888889,10.1666666667,3)` | `false/false` |
| `24` | `(8,2,Y=2/3)` | `(3,3,Y=0)` | `(256/45,16/3,8)` | `(0,8,2)` | `10.3222` | `25.2689` | `35.8169` | `2.77201e+06` | `8.59202e+12` | `3.2737e+17` | `(9.78888888889,10.1666666667,3)` | `false/false` |
| `25` | `(3,3,Y=1/6)` | `(8,2,Y=2/3)` | `(1/5,8,2)` | `(256/45,16/3,8)` | `24.9926` | `10.3067` | `35.5914` | `6.51771e+12` | `2.72935e+06` | `2.61264e+17` | `(9.98888888889,10.1666666667,3)` | `false/false` |
| `26` | `(8,2,Y=2/3)` | `(3,3,Y=1/6)` | `(256/45,16/3,8)` | `(1/5,8,2)` | `10.3067` | `24.9926` | `35.5914` | `2.72935e+06` | `6.51771e+12` | `2.61264e+17` | `(9.98888888889,10.1666666667,3)` | `false/false` |
| `27` | `(3,3,Y=1/3)` | `(8,2,Y=2/3)` | `(4/5,8,2)` | `(256/45,16/3,8)` | `24.1475` | `10.2592` | `34.9015` | `2.79941e+12` | `2.60292e+06` | `1.31058e+17` | `(10.5888888889,10.1666666667,3)` | `false/false` |
| `28` | `(8,2,Y=2/3)` | `(3,3,Y=1/3)` | `(256/45,16/3,8)` | `(4/5,8,2)` | `10.2592` | `24.1475` | `34.9015` | `2.60292e+06` | `2.79941e+12` | `1.31058e+17` | `(10.5888888889,10.1666666667,3)` | `false/false` |
| `29` | `(3,3,Y=1/2)` | `(8,2,Y=2/3)` | `(9/5,8,2)` | `(256/45,16/3,8)` | `22.6826` | `10.177` | `33.7057` | `6.46938e+11` | `2.39749e+06` | `3.96382e+16` | `(11.5888888889,10.1666666667,3)` | `false/false` |
| `30` | `(8,2,Y=2/3)` | `(3,3,Y=1/2)` | `(256/45,16/3,8)` | `(9/5,8,2)` | `10.177` | `22.6826` | `33.7057` | `2.39749e+06` | `6.46938e+11` | `3.96382e+16` | `(11.5888888889,10.1666666667,3)` | `false/false` |
| `31` | `(3,3,Y=2/3)` | `(8,2,Y=2/3)` | `(16/5,8,2)` | `(256/45,16/3,8)` | `20.5041` | `10.0547` | `31.9273` | `7.32406e+10` | `2.12157e+06` | `6.69544e+15` | `(12.9888888889,10.1666666667,3)` | `false/false` |
| `32` | `(8,2,Y=2/3)` | `(3,3,Y=2/3)` | `(256/45,16/3,8)` | `(16/5,8,2)` | `10.0547` | `20.5041` | `31.9273` | `2.12157e+06` | `7.32406e+10` | `6.69544e+15` | `(12.9888888889,10.1666666667,3)` | `false/false` |
| `33` | `(8,2,Y=1/2)` | `(3,3,Y=1)` | `(16/5,16/3,8)` | `(36/5,8,2)` | `10.1649` | `22.4662` | `33.529` | `2.36855e+06` | `5.21076e+11` | `3.32208e+16` | `(14.5,10.1666666667,3)` | `false/false` |
| `34` | `(3,3,Y=1)` | `(8,2,Y=1/2)` | `(36/5,8,2)` | `(16/5,16/3,8)` | `22.4662` | `10.1649` | `33.529` | `5.21076e+11` | `2.36855e+06` | `3.32208e+16` | `(14.5,10.1666666667,3)` | `false/false` |
| `35` | `(8,1,Y=1)` | `(8,3,Y=0)` | `(32/5,0,4)` | `(0,64/3,12)` | `10.9504` | `24.3435` | `34.4732` | `5.19552e+06` | `3.40549e+12` | `8.54027e+16` | `(10.5,18.1666666667,9)` | `false/false` |
| `36` | `(8,3,Y=0)` | `(8,1,Y=1)` | `(0,64/3,12)` | `(32/5,0,4)` | `24.3435` | `10.9504` | `34.4732` | `3.40549e+12` | `5.19552e+06` | `8.54027e+16` | `(10.5,18.1666666667,9)` | `false/false` |
| `37` | `(8,1,Y=1)` | `(8,3,Y=1/6)` | `(32/5,0,4)` | `(8/15,64/3,12)` | `11.0818` | `23.9763` | `34.042` | `5.92498e+06` | `2.35891e+12` | `5.54888e+16` | `(11.0333333333,18.1666666667,9)` | `false/false` |
| `38` | `(8,3,Y=1/6)` | `(8,1,Y=1)` | `(8/15,64/3,12)` | `(32/5,0,4)` | `23.9763` | `11.0818` | `34.042` | `2.35891e+12` | `5.92498e+06` | `5.54888e+16` | `(11.0333333333,18.1666666667,9)` | `false/false` |
| `39` | `(8,1,Y=2/3)` | `(8,3,Y=2/3)` | `(128/45,0,4)` | `(128/15,64/3,12)` | `11.0445` | `24.0805` | `34.1644` | `5.70818e+06` | `2.61793e+12` | `6.27103e+16` | `(15.4777777778,18.1666666667,9)` | `false/false` |
| `40` | `(8,3,Y=2/3)` | `(8,1,Y=2/3)` | `(128/15,64/3,12)` | `(128/45,0,4)` | `24.0805` | `11.0445` | `34.1644` | `2.61793e+12` | `5.70818e+06` | `6.27103e+16` | `(15.4777777778,18.1666666667,9)` | `false/false` |
| `41` | `(8,2,Y=2/3)` | `(8,3,Y=1/3)` | `(256/45,16/3,8)` | `(32/15,64/3,12)` | `15.5124` | `32.6722` | `37.7309` | `4.97624e+08` | `1.41018e+16` | `2.21941e+18` | `(11.9222222222,23.5,13)` | `false/false` |
| `42` | `(8,3,Y=1/3)` | `(8,2,Y=2/3)` | `(32/15,64/3,12)` | `(256/45,16/3,8)` | `32.6722` | `15.5124` | `37.7309` | `1.41018e+16` | `4.97624e+08` | `2.21941e+18` | `(11.9222222222,23.5,13)` | `false/false` |
| `43` | `(8,2,Y=2/3)` | `(8,3,Y=1/2)` | `(256/45,16/3,8)` | `(24/5,64/3,12)` | `15.5746` | `31.0921` | `36.3105` | `5.29527e+08` | `2.90433e+15` | `5.36293e+17` | `(14.5888888889,23.5,13)` | `false/false` |
| `44` | `(8,3,Y=1/2)` | `(8,2,Y=2/3)` | `(24/5,64/3,12)` | `(256/45,16/3,8)` | `31.0921` | `15.5746` | `36.3105` | `2.90433e+15` | `5.29527e+08` | `5.36293e+17` | `(14.5888888889,23.5,13)` | `false/false` |

---

## Best ranked witness

```text
target=u_topological u=1 rows=[Dirac fermion (1,3,Y=1) (1,3,Y=1); Dirac fermion (8,2,Y=1/2) (8,2,Y=1/2)] Δb1=(12/5,8/3,0) Δb2=(16/5,16/3,8) ΔbTot=(28/5,8,8) L=(L*=34.3263535,LB1=7.11786258,LB2=7.49883655) M=(M*=7.37363563e+16,MB1=112508.213,MB2=164679.341) totalBeta=(9.7,4.83333333333,1) closureU=4.44e-16 exactClosure=true minU=0.573255799 pole=none ordered=true distinct=true subPlanck=true positive=true noPole=true anomaly=true seal=true gate201=false contactMatch=false AF(SU2,SU3)=(false,false)
```

Notably, the best witnesses do **not** preserve non-Abelian asymptotic freedom in the strict `b_total < 0` sense; they survive the weaker and explicitly requested one-loop Landau-safety filter up to the prompt Planck-log bound. This distinction is recorded instead of hidden.

---

## B-sector/contact match audit

```text
bgap=true contact=true numericMatch=false semantics=false promoted=false
```

No viable row is identified with the B-sector spectral gap or seven contact partial-overlap modes. Gate 205's carrier-activation obstruction remains active: contact data still lack charge, spin-statistics, mass activation, and decoupling semantics.

---

## Baryon, anomaly, and seal audit

```text
sealInherited=true allRowsAnomaly=true allRowsSeal=true viableAnomaly=true viableSeal=true protonOperator=false lifetime=false
```

The `LeptoquarkDynamicsSeal` remains active. No dormant `u(4)` leptoquark slot is used as a mediator, propagator, coefficient, or proton-decay channel. No proton lifetime is computed.

---

## Firewall audit

```text
gate210=true lqSeal=true ledger=true observedFinite=false universal=false realCoeff=false physicalPrediction=false finiteMass=false protonSeal=false lifetime=false matching=false conditionalOnly=true next=Gate 212 — two-threshold solution minimality / finite-origin and matching-correction preflight audit
```

The numerical scales are conditional outputs of a sealed phenomenological two-threshold solve. They are not finite-derived masses, not physical predictions, and not threshold-corrected observables. Finite matching corrections and two-loop stability remain open.

---

## Theorem statement

Gate 211 confirms the dimension-counting pivot: two independent rational threshold rows make exact closure a linear solve, so the scientific content is the physical viability filter. At least one ordered two-threshold pair survives scale-ordering, prompt Planck, positivity, Landau, anomaly, and leptoquark-seal constraints (u_topological:44). These are CONDITIONAL_VIABLE phenomenological completions only; no finite carrier origin, matching correction, or physical prediction is claimed.

## Next structural obligation

Gate 212 should audit minimality, degeneracy, and finite-origin/matching-correction preflight for the viable two-threshold witnesses before any stronger phenomenological interpretation is allowed.
