# Gate 223 Registry Audit — Colored-octet pure-SM portal search / Spectrum falsification audit

## Status

```text
CONDITIONAL_PHENOMENOLOGY_RELIC_DECAY_SEAL_GRANTED
COLORED_OCTET_DIM6_PORTAL_FOUND
RANK1_SPECTRUM_NOT_FALSIFIED_BY_RELIC_DECAY
```

Gate 223 resolves the Gate-222 colored relic obstruction at the level of sealed EFT phenomenology. It does **not** derive a finite-core decay theorem.

## Inherited state

Gate 222 left the PeV spectrum partially rescued:

```text
Dirac (1,3,Y=1)      triplet portal found
Dirac (8,2,Y=1/2)    colored octet portal missing
M_B ≈ 2.56895727e6 GeV
```

The following seals remain active:

```text
EmpiricalCarrierSeal
ThresholdSpectrumSeal
MatchingCorrectionSeal
LeptoquarkDynamicsSeal
```

The false mass-mixing shortcut remains rejected:

```text
(8,2,Y=1/2) ≠ Q=(3,2,Y=1/6)
```

## Tensor-search target

Using `bar(Ψ8)O_SM`, the heavy conjugate is:

```text
bar(Ψ8) = (8,2,Y=-1/2)
```

Therefore the searched pure-SM composite must satisfy:

```text
O_SM = (8,2,Y=1/2)
fermionic Lorentz parity
B = 0 under the baryon firewall
operator dimension ≤ 6 after including bar(Ψ8)
```

The audited SM alphabet is:

```text
Q, u^c, d^c, L, e^c, H, H†, G_{μν}, W_{μν}, B_{μν}
```

No new mediator, no dormant leptoquark propagator, and no SU(5)/SO(10) parent field is imported.

## Search result

```text
combinations scanned:      213
gauge matches:             5
Lorentz + gauge matches:   3
valid baryon-safe portals: 2
```

Best-ranked witness:

```text
O_SM = Q · u^c · e^c
color contains 8
weak contains 2
Y = 1/2
dim(O_SM) = 9/2
dim[bar(Ψ8)O_SM] = 6
B = 0
```

Operator:

```text
(c_8/Λ²) bar(Ψ8)^a_i (Q_i u^c e^c)^a + h.c.
```

Second witness:

```text
O_SM = e^c · H† · G_{μν}
color = 8
weak = 2
Y = 1/2
dim(O_SM) = 9/2
dim[bar(Ψ8)O_SM] = 6
B = 0
```

Operator:

```text
(c'_8/Λ²) bar(Ψ8)^a_i σ^{μν} e^c H†_i G^a_{μν} + h.c.
```

## BBN safety bound

The BBN filter is:

```text
τ < 1 second
Γ_required > 6.582119569e-25 GeV
```

For unit Wilson coefficient, Gate 223 records two standard EFT estimates:

```text
Γ3 ≈ |c|² M_B^5/(192π³Λ⁴)
Γ2 ≈ |c|² v² M_B^3/(8πΛ⁴) after H† obtains v
```

Numerical bounds:

```text
Λ_3body(unit Wilson)  < 1.29992096e13 GeV
Λ_dipole(unit Wilson) < 4.99261316e11 GeV
conservative Λ        < 4.99261316e11 GeV
```

For a reference `Λ=1e17 GeV`, the conservative Wilson bound is:

```text
|c_8| ≳ 4.01184518e10
```

This shows that a Planck/GUT-suppressed unit-coefficient operator would not be enough under the conservative dipole proxy. The seal therefore quarantines both `c_8` and `Λ`; it does not pretend they are finite-derived.

## RelicDecaySeal verdict

Gate 223 grants:

```text
RelicDecaySeal granted conditionally on EFT portals
```

Quarantined data:

```text
triplet Yukawa y_T
octet Wilson coefficient c_8
octet suppression scale Λ
flavor choice e^c/μ^c/τ^c
post-EWSB decay/cascade semantics
future relic Boltzmann history
```

The Rank-1 PeV spectrum is therefore **not falsified** by the colored relic problem at Gate 223.

## Firewalls preserved

Gate 223 does not claim:

```text
finite-derived decay operators
finite-derived Wilson coefficients
finite-derived EFT scale
new mediator particles
activated leptoquark dynamics
relic abundance
proton lifetime
flavor safety
physical discovery
```

## Next gate

Gate 224 should audit relic abundance, flavor constraints, and rare-decay safety for the sealed portals. In particular, it must check whether the large Wilson coefficient required for a very high EFT scale is phenomenologically acceptable, or whether the decay sector demands a lower portal scale near the PeV-to-intermediate regime.
