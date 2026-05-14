# Gate 207 Registry Audit — Sealed-Threshold Prediction Stress Test / Experimental and Proton-Decay Firewall Audit

## Status

```text
FAILED_ROUTE_UNIVERSAL_COMPLETION_STRESS
```

Gate 207 consumes the Gate-206 `EmpiricalCarrierSeal` outputs as **conditional phenomenology only**. It does not promote the two sealed threshold scales into finite-core predictions. The gate stress-tests three branches:

1. direct collider reach;
2. proton-decay mediator support;
3. one-loop high-scale viability of the required external universal beta completion.

The first two branches pass conditionally. The third branch fails sharply: the required universal completion creates positive one-loop beta rows and formal sub-Planck Landau-pole/asymptotic-safety obstructions. Therefore the Gate-206 universal-completion scenario is rejected as a clean failed route unless a future gate derives a different UV completion, matching rule, or finite regulator mechanism.

---

## 1. Inherited Gate-206 sealed ledger

| Quantity | Dirac vectorlike quark doublet | Weyl `SU(2)L` adjoint fermion |
|---|---:|---:|
| Representation | `(3,2,1/6)` | `(1,3,0)` |
| Non-universal row | `(2/15,2,4/3)` | `(0,4/3,0)` |
| Required universal row `c_univ` | `7.65295390904` | `10.1497542656` |
| Sealed threshold `M_B` | `1.46774973718e6 GeV` | `8.19807624157e6 GeV` |
| Boundary scale `M_*` | `2.40099519719e15 GeV` | `2.42276543552e14 GeV` |
| Status | Conditional on `EmpiricalCarrierSeal` | Conditional on `EmpiricalCarrierSeal` |

Inherited firewalls:

```text
carrier seal explicit: true
carrier seal quarantined: true
native contact activation still obstructed: true
anomaly compatibility passed: true
universal beta source finite-derived: false
absolute mass prediction: false
physical unification claim: false
```

---

## 2. Quarantined experimental constraint ledger

Gate 207 uses external experimental facts only as stress-test inputs. They are **not** used to derive finite algebra.

| External item | Stress value used | Role |
|---|---:|---|
| Current LHC run energy marker | `13.6 TeV` | direct-production scale comparison |
| Conservative direct new-particle mass proxy | `5 TeV` | stronger-than-current stress marker for TeV searches |
| Conservative future-reach proxy | `100 TeV` | future collider stress marker, not an actual exclusion |
| Super-Kamiokande `p -> e+ pi0` lower lifetime | `2.4e34 yr` at `90% CL` | proton-decay warning marker |
| Hyper-Kamiokande long-exposure sensitivity proxy | `6e34 yr` | future proton-decay stress marker |

Reference anchors used for the external ledger:

- Super-Kamiokande reports `tau/B(p -> e+ pi0) > 2.4e34 yr` and `tau/B(p -> mu+ pi0) > 1.6e34 yr` at `90% CL` in the enlarged fiducial-volume analysis.
- Recent ATLAS/CMS vectorlike-quark searches remain TeV-scale direct searches; public summaries and proceedings typically discuss tested/excluded VLQ masses in the `~1–3 TeV` range, with some resonance searches extending to a few TeV.
- Hyper-Kamiokande projections target order-`1e35 yr` sensitivity for golden proton-decay channels over long exposures; this gate uses `6e34 yr` only as a conservative stress marker.

---

## 3. Collider constraints audit

The sealed threshold scale is PeV-scale:

| Carrier | `M_B` in GeV | `M_B` in TeV | Separation from `13.6 TeV` LHC marker | Separation from `5 TeV` direct proxy | Separation from `100 TeV` future proxy | Result |
|---|---:|---:|---:|---:|---:|---|
| Dirac vectorlike quark doublet | `1.46774973718e6` | `1467.74973718` | `107.923` | `293.550` | `14.6775` | Pass |
| Weyl `SU(2)L` adjoint fermion | `8.19807624157e6` | `8198.07624157` | `602.800` | `1639.615` | `81.9808` | Pass |

Verdict:

```text
CONDITIONAL_PASS: sealed PeV-scale carriers evade direct collider reach by orders of magnitude.
```

Firewall:

```text
No indirect collider, flavour, precision-electroweak, cosmology, lifetime, or decay-portal exclusion is claimed.
Those require portal couplings and decay operators that the engine has not derived.
```

---

## 4. Proton-decay audit

The boundary scale is low enough to trigger a naive GUT warning:

```text
min(M_*) = 2.42276543552e14 GeV
max(M_*) = 2.40099519719e15 GeV
```

A conventional unified gauge theory with `X/Y` bosons at this scale would require a serious proton-decay lifetime calculation. Gate 207 therefore audits whether ASHA currently derives the mediating channels.

### Engine-native mediator inventory

| Mediator / operator requirement | Derived by current engine? |
|---|---:|
| Full `SU(5)` gauge algebra | `false` |
| `SO(10)` gauge algebra | `false` |
| `X/Y` leptoquark gauge bosons | `false` |
| `B,L`-violating gauge curvature | `false` |
| Dimension-six proton-decay operator | `false` |
| Four-fermion `B,L`-violating local operator | `false` |
| Baryon-number violation | `false` |
| Lepton-number violation | `false` |

The current engine inventory remains:

```text
contact-preserving su(2)+u(1) seed
u(4)=central+color-su3+B-L+leptoquark matter-current inventory
no derived SU(5)/SO(10) unified gauge connection
no derived X/Y curvature sector
```

Verdict:

```text
CONDITIONAL_PASS_WITH_WARNING
```

The low `M_*` is dangerous for naive `SU(5)`-like dimension-six proton decay, but ASHA currently lacks the mediator/operator support required to instantiate that decay channel. This is a natural suppression firewall by mediator absence, not a proton lifetime prediction.

Required future gate:

```text
derive or seal a B/L-violating operator basis before computing proton lifetime
```

---

## 5. Universal completion viability audit

This is the decisive failed branch.

Gate 206 required a large real universal beta row. Above the sealed threshold, the total one-loop rows become:

| Carrier | `b1_total = 41/10 + Δb1_total` | `b2_total = -19/6 + Δb2_total` | `b3_total = -7 + Δb3_total` |
|---|---:|---:|---:|
| Dirac vectorlike quark doublet | `11.88628724237` | `6.48628724237` | `1.98628724237` |
| Weyl `SU(2)L` adjoint fermion | `14.2497542656` | `8.31642093223` | `3.1497542656` |

All three total rows are positive in both branches. The non-Abelian sectors therefore lose asymptotic freedom in the stress window.

Using the Gate-206 boundary condition `alpha^-1(M_*) = 4π`, the formal one-loop pole scale for any positive beta row is:

```text
M_pole = M_* · exp(8π² / b_total)
```

| Carrier | `U(1)` pole | `SU(2)` pole | `SU(3)` pole | Planck comparison |
|---|---:|---:|---:|---|
| Dirac vectorlike quark doublet | `1.8419242e18 GeV` | `4.64524381e20 GeV` | `4.4057126e32 GeV` | `U(1)` pole is sub-Planck |
| Weyl `SU(2)L` adjoint fermion | `6.17596741e16 GeV` | `3.21767041e18 GeV` | `1.86654873e25 GeV` | `U(1)` and `SU(2)` poles are sub-Planck |

Verdict:

```text
FAILED_ROUTE: the external universal beta completion is high-scale pathological at one loop.
```

This does not falsify the finite core, the Gate-204 row lattice, or the Gate-206 anomaly check. It falsifies this specific Gate-206 universal-completion scenario as a viable high-scale bridge unless a future theorem supplies a different UV completion, threshold matching rule, regulator trace, or finite spectral mechanism that changes the one-loop stress equation.

---

## 6. Final theorem classification

| Branch | Status |
|---|---|
| Gate-206 sealed scale inheritance | `CONDITIONAL_ON_CARRIER_SEAL` |
| Direct collider stress | `CONDITIONAL_PASS` |
| Proton-decay mediator firewall | `CONDITIONAL_PASS_WITH_WARNING` |
| Universal completion one-loop stress | `FAILED_ROUTE` |
| Overall Gate 207 theorem | `FAILED_ROUTE_UNIVERSAL_COMPLETION_STRESS` |

Truth statement:

```text
Gate 207 does not convert the sealed threshold scales into native predictions. It shows they survive a first direct-collider scale test, records a proton-decay warning for naive unified gauge theories, proves the current finite connection lacks the X/Y or B,L-violating gauge channels needed to realize that warning inside ASHA, and then rejects the external universal completion because its one-loop beta rows create sub-Planck Landau-pole/asymptotic-safety pathologies. The result is a clean FAILED_ROUTE for the Gate-206 universal-completion scenario, not an absolute phenomenology claim.
```

---

## 7. Next structural obligation

```text
Gate 208 — baryon/lepton violating operator basis audit / proton-decay channel construction obstruction
```

Rationale:

1. Gate 207 cannot legally compute proton lifetime because no `B/L`-violating local operator has been derived.
2. The low `M_*` cannot be interpreted using naive `SU(5)` formulas unless `X/Y`-like gauge bosons or equivalent dimension-six operators are present.
3. A future gate must audit whether ASHA's `u(4)` matter-current inventory and contact/electroweak carrier can generate, forbid, or seal any baryon/lepton-violating operator basis.
