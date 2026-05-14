# Gate 214 Registry Audit — Sealed two-loop RG integration / matching-correction uncertainty envelope audit

## Verdict

```text
CONDITIONAL_PHENOMENOLOGY_ON_THRESHOLD_SPECTRUM_AND_MATCHING_ENVELOPE
```

Gate 214 performs the first full numerical two-loop integration for the Gate-213 sealed spectrum. It does **not** replace the finite algebra with a continuum assumption. It inherits the `ThresholdSpectrumSeal`, keeps the `EmpiricalCarrierSeal` and `LeptoquarkDynamicsSeal` active, and treats the Z-pole ledger as quarantined phenomenological input.

The result is a conditional two-loop numerical fit, plus a theory-uncertainty envelope for the still-un-derived threshold matching corrections.

---

## Inherited Gate-213 state

Gate 213 introduced:

```text
ThresholdSpectrumSeal
SEAL-THRESHOLD-SPECTRUM-GATE213
```

and selected the Gate-211 ranked witness only as a sealed test subject:

| Carrier | Representation | One-loop row |
|---|---:|---:|
| row 1 | Dirac `(1,3,Y=1)` | `(12/5, 8/3, 0)` |
| row 2 | Dirac `(8,2,Y=1/2)` | `(16/5, 16/3, 8)` |

Gate 213 also proved:

```text
FAILED_ROUTE_DERIVED_MATCHING_CORRECTIONS
TWO_LOOP_PREFLIGHT_WARNING_ONE_LOOP_STABILITY_NOT_PROVEN
```

So Gate 214 is allowed to run a numerical two-loop phenomenology test, but it is **not** allowed to claim finite-derived matching corrections, physical masses, or precision predictions.

---

## Two-loop equation integrated

Gate 214 integrates in `u = 1/g²` coordinates using the no-Yukawa two-loop gauge equation:

```text
du_i/dlnμ = -b_i/(8π²) - Σ_j B_ij / u_j /(128π⁴)
```

with piecewise threshold activation at two continuous logarithmic scales:

```text
L_B1 = ln(M_B1/M_Z)
L_B2 = ln(M_B2/M_Z)
L_*  = ln(M_*/M_Z)
```

The numerical method is:

```text
piecewise fixed-step RK4 + damped Newton finite-difference solve
```

No SM Yukawa matrices are included yet. No matching corrections are used in the central solve.

---

## Central two-loop corrected solution

The central target is:

```text
u_*(M_*) = (1, 1, 1)
```

The two-loop solve converges to:

| Quantity | Value |
|---|---:|
| `L_B1` for Dirac `(1,3,Y=1)` | `10.3098091` |
| `M_B1` | `2.73797183e6 GeV` |
| `L_B2` for Dirac `(8,2,Y=1/2)` | `10.2599421` |
| `M_B2` | `2.60478578e6 GeV` |
| `L_*` | `35.1875394` |
| `M_*` | `1.74457638e17 GeV` |
| `|L_B1-L_B2|` | `0.0498670021` |
| residual norm | `8.39e-14` |

The two-loop correction changes the Gate-211 one-loop picture substantially:

| Scale | Gate 211 / Gate 213 one-loop reference | Gate 214 central two-loop fit |
|---|---:|---:|
| `M_B1` | `1.12508213e5 GeV` | `2.73797183e6 GeV` |
| `M_B2` | `1.64679341e5 GeV` | `2.60478578e6 GeV` |
| `M_*` | `7.37363563e16 GeV` | `1.74457638e17 GeV` |

The threshold order also flips relative to the Gate-211 ordered witness:

```text
(8,2,Y=1/2) activates before (1,3,Y=1)
```

This is allowed in Gate 214 because the `ThresholdSpectrumSeal` quarantines the spectrum, and the two logarithmic scales are treated as continuous phenomenological fit parameters. It is not a finite-derived ordering theorem.

---

## Positivity and Landau safety

The central two-loop solution passes the numerical safety checks:

| Check | Result |
|---|---:|
| `L_B1 > 0`, `L_B2 > 0` | pass |
| `L_* > max(L_B1,L_B2)` | pass |
| `L_* < 37.8` prompt Planck-log firewall | pass |
| distinct thresholds | pass |
| positive couplings up to `M_*` | pass |
| no Landau pole below prompt Planck-log firewall | pass |

These are viability checks only. They are not a UV-complete proof.

---

## Matching-correction uncertainty envelope

Gate 213 proved that exact threshold matching corrections are not derived. Gate 214 therefore introduces an explicit phenomenological proxy:

```text
MatchingUncertaintyEnvelope
ε_u = 1/(16π²) = 0.00633257397765
```

The envelope is a deterministic corner scan over:

```text
δu_i ∈ { -ε_u, +ε_u }
```

for the three gauge directions at the boundary target. This is a standard loop-factor uncertainty proxy. It is **not** a finite-core theorem and does not replace a real matching calculation.

All 8 corner solves converge.

| Scale | Central | Envelope minimum | Envelope maximum |
|---|---:|---:|---:|
| `M_B1` | `2.73797183e6 GeV` | `1.57840858e6 GeV` | `4.74995204e6 GeV` |
| `M_B2` | `2.60478578e6 GeV` | `2.41692805e6 GeV` | `2.80741458e6 GeV` |
| `M_*` | `1.74457638e17 GeV` | `1.45661625e17 GeV` | `2.08954763e17 GeV` |

The large `M_B1` spread is a warning: the sealed two-threshold spectrum is sensitive to unknown finite matching data.

---

## Matching-correction obstruction remains active

Available finite data remains only support data:

```text
tau_eta scalar fundamental class
scalar traces
contact zeta traces
```

Still missing:

```text
spectral triple
heat-kernel matching map
canonical subtraction scheme
finite counterterm functional
δ_i^match rows
```

Therefore:

```text
FAILED_ROUTE_DERIVED_MATCHING_CORRECTIONS
```

remains active.

---

## Firewalls

Gate 214 does **not** claim:

```text
finite-derived heavy masses
finite-derived threshold spectrum
finite-derived matching corrections
unique physical spectrum
SM Yukawa-corrected two-loop running
physical prediction
proton lifetime
contact/B-sector particle promotion
```

The output is conditional on:

```text
ThresholdSpectrumSeal
EmpiricalCarrierSeal
LeptoquarkDynamicsSeal
quarantined Z-pole ledger
no-Yukawa standard-QFT two-loop equations
phenomenological matching-envelope proxy
```

---

## Registry theorem

Package:

```text
pkg/bridge/twoloopintegration
```

Theorem:

```text
SealedTwoLoopRGIntegrationMatchingEnvelopeTheorem
```

Status:

```text
PHENOMENOLOGY
```

The theorem is successful only in the conditional sense: it computes a sealed two-loop fit and an explicit uncertainty envelope while preserving all firewalls.

---

## Next obligation

Gate 214 shows that the one-loop witness survives as a two-loop fit, but the uncertainty envelope is not small enough to support a final precision claim. The next legal directions are:

1. audit the same two-loop envelope over all 22 unordered Gate-211 spectra;
2. add SM Yukawa/scalar effects only under a new seal or finite derivation;
3. derive a finite heat-kernel/spectral matching map to replace the loop-factor proxy.
