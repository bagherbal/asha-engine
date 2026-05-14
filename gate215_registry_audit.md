# Gate 215 Registry Audit

## Gate

**Gate 215 — Single-scale degenerate-limit matching audit / global two-loop class scan**

Package:

```text
pkg/bridge/singlescalematchingaudit
```

Registry theorem:

```text
BRIDGE-SINGLE-SCALE-DEGENERATE-LIMIT-MATCHING-AUDIT
```

Status:

```text
CONDITIONAL_PHENOMENOLOGY_SINGLE_SCALE_MATCHING_AUDIT
```

This is a sealed numerical phenomenology scan. It does not derive finite matching corrections, heavy masses, or a unique physical spectrum.

---

## Purpose

Gate 214 showed that the separated two-loop solution for the Gate-211 ranked witness becomes nearly degenerate:

```text
ΔL ≈ 0.0498670021
```

Gate 215 tests the resulting hypothesis:

```text
Could the true heavy spectrum be single-scale, with the remaining closure defect supplied by finite threshold matching corrections?
```

Because Gate 213 and Gate 214 both obstructed derived `δ_i^match`, Gate 215 does not invent those corrections. It computes the **required** correction vector and compares its magnitude to the explicit Gate-214 loop-factor envelope:

```text
ε_u = 1/(16π²) = 0.00633257397765
```

---

## Inheritance

```text
gate214=true thresholdSeal=true envelope=true central=true ΔL=0.0498670021 ε=0.00633257397765 orderedPairs=44 unorderedClasses=22
```

The scan inherits:

- `ThresholdSpectrumSeal`
- `EmpiricalCarrierSeal`
- `LeptoquarkDynamicsSeal`
- quarantined Gate-200 Z-pole ledger
- Gate-214 no-Yukawa two-loop integration convention
- Gate-214 matching uncertainty envelope

---

## Method

For each of the 22 unordered Gate-211 viable pair classes, Gate 215 forces:

```text
L_B1 = L_B2 = L_B
M_B1 = M_B2 = M_B
```

It then integrates the two-loop system:

```text
du_i/dlnμ = -b_i/(8π²) - Σ_j B_ij/u_j /(128π⁴)
```

and optimizes over only two continuous parameters:

```text
(L_B, L_*)
```

The residual at the boundary is:

```text
r_i = u_i(L_*) - 1
```

The required threshold matching correction is recorded as:

```text
δ_i^req = -r_i
```

A class is marked plausible only if:

```text
max_i |r_i| ≤ ε_u
```

This is a plausibility filter, not a derivation of matching corrections.

---

## Global scan result

```text
classes=22 supported=22 converged=22 plausible=1 rejected=21 bestRank=1 best="(1,3,Y=1) [Dirac fermion] + (8,2,Y=1/2) [Dirac fermion]" bestMax=0.000561440698 bestOverε=0.0886592 MBRange=[806427.924,1.24954064e+14] verdict="1 of 22 unordered Gate-211 classes have forced single-scale two-loop residuals inside the ε=1/(16π²) matching envelope"
```

Main result:

```text
Only the Gate-211 ranked witness remains plausible in the forced single-scale two-loop limit under the loop-factor matching envelope.
```

---

## Best degenerate-limit witness

```text
rank=1 rows=[(1,3,Y=1) [Dirac fermion] + (8,2,Y=1/2) [Dirac fermion]] L=(LB=10.2609928,L*=35.1715498) M=(MB=2607524.25,M*=1.71690311e+17) U=(1.00056119,0.999438559,1.00056051) residual=(0.000561193804,-0.000561440698,0.000560508948) requiredδmatch=(-0.000561193804,0.000561440698,-0.000560508948) max=0.000561440698 rms=0.000561047955 overε=0.0886592 plausible=true ordered=true subPlanck=true positive=true noLandau=true status=MATCHING_RESIDUAL_WITHIN_LOOP_FACTOR_ENVELOPE
```

Interpretation:

| Quantity | Value |
|---|---:|
| `M_B` | `2.60752425e6 GeV` |
| `M_*` | `1.71690311e17 GeV` |
| `max |r_i|` | `0.000561440698` |
| `max |r_i| / ε_u` | `0.0886592` |
| Matching plausibility | `true` |

The required matching correction is under 9% of the Gate-214 loop-factor envelope. This is a strong conditional signal that the Gate-211 ranked witness can be interpreted as a single-scale spectrum **if** future finite spectral matching supplies a correction of the displayed form.

---

## Ranked 22-class scan

| Rank | Gate-211 class | Rows | `M_B` GeV | `M_*` GeV | `max |r|` | `max |r| / ε` | RMS residual | Plausible | No sub-Planck Landau |
|---:|---:|---|---:|---:|---:|---:|---:|---|---|
| `1` | `1` | `(1,3,Y=1) [Dirac fermion] + (8,2,Y=1/2) [Dirac fermion]` | `2.60752e+06` | `1.7169e+17` | `0.000561441` | `0.0886592` | `0.000561048` | `true` | `true` |
| `2` | `10` | `(8,2,Y=2/3) [Dirac fermion] + (8,3,Y=2/3) [complex scalar]` | `2.64071e+09` | `8.01326e+16` | `0.0105735` | `1.6697` | `0.0105715` | `false` | `true` |
| `3` | `2` | `(1,3,Y=0) [Dirac fermion] + (8,2,Y=2/3) [Dirac fermion]` | `2.23341e+06` | `1.15373e+17` | `0.0125651` | `1.98421` | `0.0125613` | `false` | `true` |
| `4` | `3` | `(1,3,Y=1/6) [Dirac fermion] + (8,2,Y=2/3) [Dirac fermion]` | `2.08906e+06` | `9.7106e+16` | `0.017018` | `2.68737` | `0.0170155` | `false` | `true` |
| `5` | `11` | `(8,2,Y=1/2) [Dirac fermion] + (8,3,Y=1) [complex scalar]` | `1.35938e+09` | `2.24115e+16` | `0.0181683` | `2.86903` | `0.0181681` | `false` | `true` |
| `6` | `4` | `(8,2,Y=2/3) [Dirac fermion] + (1,3,Y=1/3) [Dirac fermion]` | `1.71533e+06` | `5.83909e+16` | `0.0301528` | `4.76154` | `0.0301521` | `false` | `true` |
| `7` | `14` | `(8,2,Y=2/3) [Dirac fermion] + (3,3,Y=1/3) [Dirac fermion]` | `1.34394e+08` | `4.3613e+17` | `0.0349079` | `5.51244` | `0.0261951` | `false` | `false` |
| `8` | `5` | `(8,2,Y=2/3) [Dirac fermion] + (1,3,Y=1/2) [Dirac fermion]` | `1.24833e+06` | `2.57038e+16` | `0.0513228` | `8.10457` | `0.0513224` | `false` | `true` |
| `9` | `9` | `(8,2,Y=2/3) [Dirac fermion] + (8,3,Y=1/2) [complex scalar]` | `8.66398e+09` | `6.90994e+17` | `0.0637346` | `10.0646` | `0.0637341` | `false` | `true` |
| `10` | `6` | `(1,3,Y=2/3) [Dirac fermion] + (8,2,Y=2/3) [Dirac fermion]` | `806428` | `8.31581e+15` | `0.0804167` | `12.6989` | `0.0789653` | `false` | `true` |
| `11` | `8` | `(8,2,Y=2/3) [Dirac fermion] + (8,3,Y=1/3) [complex scalar]` | `1.43733e+10` | `2.3783e+18` | `0.127118` | `20.0737` | `0.10973` | `false` | `true` |
| `12` | `15` | `(8,2,Y=2/3) [Dirac fermion] + (3,3,Y=1/2) [Dirac fermion]` | `1.48415e+07` | `1.18718e+16` | `0.139196` | `21.9809` | `0.0962674` | `false` | `false` |
| `13` | `7` | `(8,3,Y=0) [Weyl fermion] + (8,1,Y=1) [Dirac fermion]` | `1.09392e+09` | `3.83247e+17` | `0.143796` | `22.7074` | `0.143796` | `false` | `true` |
| `14` | `17` | `(8,2,Y=1/2) [Dirac fermion] + (3,3,Y=1) [Dirac fermion]` | `1.00513e+08` | `2.10229e+14` | `0.194583` | `30.7274` | `0.19449` | `false` | `false` |
| `15` | `16` | `(3,3,Y=2/3) [Dirac fermion] + (8,2,Y=2/3) [Dirac fermion]` | `3.56559e+08` | `1.83794e+15` | `0.270744` | `42.7542` | `0.269613` | `false` | `false` |
| `16` | `20` | `(8,3,Y=2/3) [Dirac fermion] + (8,1,Y=2/3) [Dirac fermion]` | `9.80461e+11` | `2.99204e+16` | `0.38264` | `60.4241` | `0.381857` | `false` | `false` |
| `17` | `18` | `(8,3,Y=0) [Dirac fermion] + (8,1,Y=1) [Dirac fermion]` | `8.44812e+10` | `1.71536e+18` | `0.396056` | `62.5427` | `0.33714` | `false` | `false` |
| `18` | `13` | `(3,3,Y=1/6) [Dirac fermion] + (8,2,Y=2/3) [Dirac fermion]` | `7.12059e+09` | `2.98083e+17` | `0.453684` | `71.6429` | `0.453684` | `false` | `true` |
| `19` | `12` | `(3,3,Y=0) [Dirac fermion] + (8,2,Y=2/3) [Dirac fermion]` | `8.96804e+09` | `4.40186e+17` | `0.467811` | `73.8737` | `0.46781` | `false` | `true` |
| `20` | `22` | `(8,3,Y=1/2) [Dirac fermion] + (8,2,Y=2/3) [Dirac fermion]` | `1.24954e+14` | `1.02938e+18` | `0.506514` | `79.9855` | `0.506383` | `false` | `true` |
| `21` | `19` | `(8,1,Y=1) [Dirac fermion] + (8,3,Y=1/6) [Dirac fermion]` | `1.89281e+13` | `2.37829e+18` | `0.646103` | `102.029` | `0.592478` | `false` | `true` |
| `22` | `21` | `(8,2,Y=2/3) [Dirac fermion] + (8,3,Y=1/3) [Dirac fermion]` | `7.58203e+13` | `2.14298e+18` | `0.675917` | `106.737` | `0.574159` | `false` | `false` |

---

## Matching obstruction preserved

```text
gate214Envelope=true nativeRows=false heatKernel=false subtraction=false proxy=true promoted=false status=FAILED_ROUTE_DERIVED_MATCHING_CORRECTIONS_PRESERVED verdict="Gate 215 computes the required matching residual vector but does not derive δ_i^match; the loop-factor envelope remains a phenomenological proxy inherited from Gate 214"
```

The residual vector is a target for future finite matching theory. It is not a derived counterterm.

---

## Firewall audit

```text
thresholdSeal=true carrierSeal=true lqSeal=true ledger=true all22=true forcedFinite=false matchingDerived=false residualPromoted=false yukawa=false prediction=false lifetime=false next="Gate 216 — matching-residual structure audit / spectral heat-kernel coefficient search"
```

Preserved firewalls:

| Firewall | Status |
|---|---|
| `ThresholdSpectrumSeal` remains active | Preserved |
| `EmpiricalCarrierSeal` remains active | Preserved |
| `LeptoquarkDynamicsSeal` remains active | Preserved |
| Z-pole ledger remains phenomenological | Preserved |
| Single-scale threshold is not finite-core derived | Preserved |
| Required matching residual is not promoted to a finite theorem | Preserved |
| No SM Yukawa matrices imported | Preserved |
| No physical prediction claimed | Preserved |
| No proton lifetime computed | Preserved |

---

## Theorem statement

Gate 215 globally audits the forced single-scale, two-loop degenerate limit of all 22 unordered Gate-211 viable spectra. Exactly one class—the Gate-211 ranked witness `(1,3,Y=1)` Dirac fermion plus `(8,2,Y=1/2)` Dirac fermion—requires a matching residual smaller than the inherited loop-factor envelope. This supports the hypothesis that the near-degenerate Gate-214 two-threshold result may be interpretable as a single-scale spectrum once finite matching corrections are derived. The correction is not currently derived, so the result remains conditional phenomenology.

## Next structural obligation

```text
Gate 216 — matching-residual structure audit / spectral heat-kernel coefficient search
```

Minimum open requirements:

```text
derive a finite heat-kernel or spectral matching map before promoting δ_i^req
include SM Yukawa two-loop terms only under a separate empirical Yukawa seal
rerun the 22-class scan if the matching envelope is replaced by exact finite counterterms
preserve ThresholdSpectrumSeal until a native spectrum selector is derived
```
