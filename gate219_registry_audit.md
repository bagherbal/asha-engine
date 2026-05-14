# Gate 219 Registry Audit — Input-sensitivity and bottom-tau-Yukawa completeness audit

## Gate identity

```text
Gate 219 — Input-sensitivity and bottom-tau-Yukawa completeness audit
Package: pkg/bridge/inputsensitivityaudit
Registry theorem: BRIDGE-INPUT-SENSITIVITY-BOTTOM-TAU-YUKAWA-COMPLETENESS
Status: CONDITIONAL_PHENOMENOLOGY_INPUT_SENSITIVITY_BOTTOM_TAU_COMPLETE_2LOOP
```

## Inherited state

Gate 218 introduced:

```text
MatchingCorrectionSeal
SEAL-MATCHING-CORRECTION-GATE218
```

and showed that the forced single-scale spectrum

```text
Dirac (1,3,Y=1) + Dirac (8,2,Y=1/2)
```

remains viable after top-Yukawa and Higgs-quartic running:

```text
M_B  ≈ 2.56883502e6 GeV
M_*  ≈ 1.72153998e17 GeV
max|δ_required| / ε_u ≈ 0.1344
```

Gate 218 deliberately omitted bottom and tau Yukawa running and did not propagate empirical input uncertainties. Gate 219 closes that narrow precision gap.

## Empirical input ledger

Gate 219 uses a quarantined phenomenological input ledger. None of these are finite-core derivations.

```text
M_Z              = 91.1876 GeV
α_em^-1(M_Z)     = 127.955       [held fixed]
sin²θ_W(M_Z)     = 0.23122       [held fixed]
α_s(M_Z)         = 0.1179 ± 0.0009
m_t              = 172.56 ± 0.70 GeV
m_H              = 125.20 ± 0.11 GeV
m_b              = 4.18 ± 0.03 GeV
m_τ              = 1.77686 ± 0.00012 GeV
v                = 246.21965 GeV
```

Tree-level seeds are used only as phenomenological starting values:

```text
y_t(M_Z)     = sqrt(2) m_t / v
y_b(M_Z)     = sqrt(2) m_b / v
y_τ(M_Z)     = sqrt(2) m_τ / v
λ(M_Z)       = m_H² / (2v²)
```

## Bottom/tau completeness upgrade

Gate 219 evolves:

```text
y_t
y_b
y_τ
λ
```

and upgrades the two-loop gauge equation to:

```text
du_i/dlnμ = -b_i/(8π²)
            - (Σ_j B_ij/u_j - c_i^t y_t² - c_i^b y_b² - c_i^τ y_τ²)/(128π⁴)
```

with Yukawa coefficients:

```text
c^t = (17/10, 3/2, 2)
c^b = (1/2,   3/2, 2)
c^τ = (3/2,   1/2, 0)
```

The lighter fermion Yukawas and heavy-sector Yukawas remain absent because the engine has not derived them and the prompt does not introduce a seal for them.

## Central bottom/tau-complete result

Forced degenerate threshold:

```text
M_B1 = M_B2 = M_B
```

Central fit:

```text
L_B  = 10.2460917
L_*  = 35.1743947
M_B  = 2.56895727e6 GeV
M_*  = 1.72179441e17 GeV
```

Boundary values:

```text
u_1(M_*) = 1.00083561056
u_2(M_*) = 0.999144875073
u_3(M_*) = 1.00085491722
```

Required matching residual:

```text
δ_required = (-0.000835610558, +0.000855124927, -0.000854917218)
```

Loop-factor comparison:

```text
ε_u = 1/(16π²) ≈ 0.00633257398
max|δ_required| / ε_u ≈ 0.135036
```

The bottom/tau upgrade shifts the Gate-218 central result only mildly and preserves the matching plausibility envelope.

## 1σ sensitivity scan

Gate 219 runs one-at-a-time ±1σ perturbations of every scan-enabled input:

```text
α_s(M_Z), m_t, m_H, m_b, m_τ
```

No input is tuned to reduce the residual. Each perturbation is independently propagated through the same forced single-scale two-loop solver.

Summary:

```text
cases audited:       11  (central + 10 boundary cases)
converged cases:     11
plausible cases:     11
broken-envelope:     0
```

Induced scale ranges:

```text
M_B range  = [2.46868509e6, 2.67089887e6] GeV
M_* range  = [1.66008302e17, 1.78344443e17] GeV
```

Central error bars from the bounded scan:

```text
M_B = 2.56895727e6 GeV  -1.00272e5 / +1.01942e5 GeV
M_* = 1.72179441e17 GeV -6.17114e15 / +6.16500e15 GeV
```

Residual-envelope range:

```text
min max|δ|/ε_u = 0.133839
max max|δ|/ε_u = 0.411919
```

Worst case:

```text
α_s(M_Z) - 1σ
max|δ|/ε_u ≈ 0.411919
```

Even the worst one-at-a-time 1σ input perturbation remains comfortably inside the loop-factor matching envelope.

## Dominant sensitivity

The scan identifies the strong coupling as the dominant driver:

```text
Dominant scale driver:    α_s(M_Z) +1σ
Dominant residual driver: α_s(M_Z) -1σ
```

Top-mass, bottom-mass, tau-mass, and Higgs-mass perturbations are subdominant in this gauge-focused two-loop audit. The Higgs mass changes λ, but at this approximation λ does not feed back into the two-loop gauge equations; therefore its direct effect on the matching residual is negligible.

## Verdict

```text
CONDITIONAL_PHENOMENOLOGY
```

The PeV-scale single-threshold hypothesis remains viable after:

```text
bottom-Yukawa completion
tau-Yukawa completion
±1σ empirical input propagation
```

The most important quantitative output is:

```text
M_B ≈ 2.57 PeV with ~4% one-at-a-time 1σ input sensitivity
M_* ≈ 1.72e17 GeV with ~3.6% one-at-a-time 1σ input sensitivity
matching residual remains within ε_u for all audited 1σ cases
```

## Firewalls preserved

Gate 219 does not claim:

```text
finite-derived α_s
finite-derived top mass
finite-derived bottom mass
finite-derived tau mass
finite-derived Higgs mass
finite-derived Yukawa matrices
finite-derived matching corrections
physical prediction
proton lifetime
contact-mode particle promotion
B-gap mass activation
```

The scan is an empirical uncertainty propagation under active seals:

```text
ThresholdSpectrumSeal
MatchingCorrectionSeal
EmpiricalCarrierSeal
LeptoquarkDynamicsSeal
```

## Next obligation

The next natural gate should move from internal precision viability to observability:

```text
Gate 220 — experimental observability / PeV-threshold indirect-signature audit
```

Possible branches:

```text
PeV-scale direct production obstruction
threshold-imprint observables
precision unification residue
cosmological stability / relic abundance firewall
neutrino or flavor constraints if couplings are sealed
```
