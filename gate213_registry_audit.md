# Gate 213 Registry Audit — ThresholdSpectrumSeal / matching-correction and two-loop stability preflight audit

## Package

`pkg/bridge/thresholdspectrumseal`

## Theorem

`BRIDGE-THRESHOLD-SPECTRUM-SEAL-MATCHING-CORRECTION-TWO-LOOP-PREFLIGHT-AUDIT`

## Status

`PHENOMENOLOGY`

Internal gate status:

```text
CONDITIONAL_PHENOMENOLOGY_ON_THRESHOLD_SPECTRUM_SEAL
```

Gate 213 does not claim a unique spectrum, corrected mass scale, finite matching correction, or physical prediction. It seals one Gate-211 witness as a conditional test subject and audits the next layer of theoretical uncertainty.

---

## Gate 212 inheritance

Gate 212 proved:

```text
FAILED_ROUTE_CANONICAL_THRESHOLD_UNIQUENESS
```

with:

```text
ordered Gate-211 witnesses:      44
unordered physical pair classes: 22
canonical finite selector:       none
ThresholdSpectrumSeal required:  yes
```

Gate 213 inherits this result and refuses to treat Gate-211 ranking as finite uniqueness.

---

## ThresholdSpectrumSeal

Gate 213 introduces:

```text
ThresholdSpectrumSeal
SEAL-THRESHOLD-SPECTRUM-GATE213
```

Purpose:

```text
quarantine the choice of one heavy threshold spectrum from the degenerate Gate-211 witness space
```

The seal permits one ranked witness to be used for a conditional preflight, while explicitly forbidding the following claims:

```text
unique finite-derived heavy spectrum: no
contact-mode carrier origin:          no
B-sector mass activation:             no
matching scheme derivation:           no
two-loop coefficients as finite core: no
physical prediction:                  no
```

---

## Sealed conditional test subject

The Gate-211 best-ranked topological witness is selected only under the seal:

| Quantity | Value |
|---|---:|
| Boundary branch | `u_topological = 1` |
| Row 1 | `Dirac fermion (1,3,Y=1)` |
| `Δb^(1)` | `(12/5,8/3,0)` |
| Row 2 | `Dirac fermion (8,2,Y=1/2)` |
| `Δb^(2)` | `(16/5,16/3,8)` |
| Total `Δb` | `(28/5,8,8)` |
| Total one-loop beta | `(9.7,4.83333333333,1)` |
| `L_B1` | `7.11786257508` |
| `M_B1` | `1.12508212505e5 GeV` |
| `L_B2` | `7.49883655200` |
| `M_B2` | `1.64679341123e5 GeV` |
| `ΔL` | `0.380973976917` |
| `L_*` | `34.3263534514` |
| `M_*` | `7.37363563442e16 GeV` |
| `α_GUT^-1` | `4π` |

These are one-loop Gate-211 reference values only.

---

## Matching-correction obstruction audit

The finite data available from earlier gates include:

```text
tau_eta scalar fundamental class: available
scalar trace support:             available
contact zeta traces:              available
```

But precision threshold matching requires more than finite support traces. Gate 213 finds the following structures still missing:

```text
spectral triple:                    not complete
heat-kernel matching map:           not derived
canonical subtraction scheme:       not derived
finite counterterm functional:      not derived
MSbar / dimensional regularization: not imported as finite theorem
threshold matching rows δ_i^match:  0
```

Therefore the matching branch records:

```text
FAILED_ROUTE_DERIVED_MATCHING_CORRECTIONS
```

Interpretation:

```text
δ_i^match remains scheme-dependent and sealed. The exact one-loop scales have theoretical uncertainty that Gate 213 does not pretend to remove.
```

---

## Two-loop coefficient preflight

Gate 213 uses the standard no-Yukawa two-loop gauge convention:

```text
dg_i/dlnμ = g_i^3 b_i/(16π²) + g_i^3 Σ_j B_ij g_j²/(16π²)²
```

with GUT-normalized `U(1)`.

For a sealed Dirac fermion representation, the preflight formula is:

```text
ΔB_ii = (20/3 C2(G_i) + 4 C2_i(R)) S_i(R)
ΔB_ij = 4 C2_j(R) S_i(R), i ≠ j
```

where `S_i(R)` is the Dynkin index for gauge factor `i`, including multiplicity under spectator gauge factors.

This is a standard-QFT preflight formula. It is **not** imported as a finite-core theorem.

### Carrier 1: Dirac `(1,3,Y=1)`

```text
ΔB_1 = [[108/25,72/5,0],
        [24/5,128/3,0],
        [0,0,0]]
```

### Carrier 2: Dirac `(8,2,Y=1/2)`

```text
ΔB_2 = [[36/25,36/5,144/5],
        [12/5,196/3,48],
        [18/5,18,192]]
```

### Heavy-induced two-loop matrix

```text
ΔB_heavy = [[144/25,108/5,144/5],
            [36/5,108,48],
            [18/5,18,192]]
```

### Total no-Yukawa two-loop matrix

Adding the standard SM no-Yukawa two-loop matrix:

```text
B_SM = [[199/50,27/10,44/5],
        [9/10,35/6,12],
        [11/10,9/2,-26]]
```

produces:

```text
B_total = [[487/50,243/10,188/5],
           [81/10,683/6,60],
           [47/10,45/2,166]]
```

---

## Two-loop stability preflight

Gate 213 audits the ratio of the two-loop contribution to the one-loop derivative in `u = 1/g²` space over the three one-loop segments:

| Segment | Active rows | Max ratio | Dominant gauge |
|---|---|---:|---|
| `0 ≤ L < L_B1` | SM | `0.0408914859` | `SU(2)_L` |
| `L_B1 ≤ L < L_B2` | SM + `(1,3,1)` | `0.365744349` | `SU(2)_L` |
| `L_B2 ≤ L ≤ L_*` | SM + both sealed carriers | `1.22345329` | `SU(3)_C` |

Verdict:

```text
TWO_LOOP_PREFLIGHT_WARNING_ONE_LOOP_STABILITY_NOT_PROVEN
```

This is not a proof of catastrophic failure. It is a warning that the one-loop witness is not perturbatively controlled by the one-loop equation alone near the high-scale segment. A full two-loop piecewise integration plus a matching-correction envelope is required before any precision scale claim.

---

## Firewall

```text
ThresholdSpectrumSeal introduced: yes
LeptoquarkDynamicsSeal inherited: yes
EmpiricalCarrierSeal inherited:   yes
Z-pole ledger quarantined:        yes
unique physical spectrum claimed: no
contact modes promoted:           no
B-sector gap promoted to mass:    no
matching corrections derived:     no
MSbar imported as finite core:     no
two-loop coefficients finite-core: no
two-loop corrected scales claimed: no
physical prediction claimed:       no
proton lifetime computed:          no
```

---

## Final statement

Gate 213 safely opens the next precision layer without violating the finite-core firewall.

It establishes:

```text
Gate-211 one-loop bridge survives as sealed phenomenological reference data.
Gate-212 degeneracy is quarantined by ThresholdSpectrumSeal.
Finite matching corrections remain obstructed.
Exact symbolic heavy two-loop coefficients can be computed as preflight data.
The selected witness is not proven two-loop stable because the high-scale SU(3) correction is not small.
```

Therefore the correct next obligation is a sealed two-loop RG integration and matching-envelope uncertainty audit, not a claim of final physical prediction.

## Validation

Focused tests:

```bash
go test -p=1 ./pkg/bridge/thresholdspectrumseal -count=1 -timeout=300s
```

Focused dependency batch:

```bash
go test -p=1 ./pkg/bridge/twothresholdminimality ./pkg/bridge/thresholdspectrumseal -count=1 -timeout=300s
```

Source-level wiring:

```bash
go list ./pkg/bridge/thresholdspectrumseal ./internal/app ./cmd/asha
```

Full historical theorem-ladder tests were not run.
