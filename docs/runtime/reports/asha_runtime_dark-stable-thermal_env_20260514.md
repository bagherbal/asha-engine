# ASHA Runtime Board

- Runtime: `gate425-runtime-env-scenarios-20260514`
- Latest gate: `425`
- Scenario: `dark-stable-thermal`
- Source: `github.com/bagherbal/asha-engine`

## Dark-sector conditional scenarios

The runtime now computes the viable and rejected dark-sector paths. A stable thermal B-gap Majorana relic is ruled out by overclosure; a suppressed/nonthermal or decaying route remains a conditional cosmological-history bridge.

| Symbol | Value | Status | Formula / note |
|---|---:|---|---|
| `M_B` | 1467749.73718 GeV | `bridge-required` | B-gap heavy Majorana scale  |
| `Ω_DM h² target` | 0.12 | `environmental` | observational comparator  |
| `Y_required` | 2.97981078152e-16 | `bridge-required` | Ω h² ρ_c/(m s_0) yield needed if B-gap particle were all dark matter |
| `Y_thermal` | 0.00390149836766 | `bridge-required` | 135 ζ(3)/(8π⁴) · g/g_*S stable relativistic thermal abundance stress test |
| `Ω_thermal h²` | 1.57117293159e+12 | `failed-route` | m Y_thermal s_0/ρ_c  |
| `Ω_thermal/Ω_DM` | 1.30931077633e+13 | `failed-route` | stable thermal B-gap relic overclosure  |
| `Y_required/Y_thermal` | 7.63760612134e-14 | `bridge-required` | required suppression / dilution fraction viable only with nonthermal production, dilution, or decay history |

### Boundaries

- **B-gap Majorana stable thermal relic:** `Ω_candidate/Ω_DM ~ 1.3×10¹³` — simple stable thermal interpretation is rejected by overclosure.
- **decaying/portal heavy sector:** `Ω_heavy h² = 0 after decay assumptions` — allowed only with sealed decay/portal dynamics.
- **nonthermal/axion-like routes:** `requires production-history axiom` — not native ASHA output.

## Runtime checks

- ✅ **Clifford dimension** — dim Cℓ(1,7)=2^8
- ✅ **Exterior grade dimensions** — [1,8,28,56,70,56,28,8,1]
- ✅ **Boolean/G2 contact vacuum** — rank(P_B)=56 rank(P_G)=14 dim K=7
- ✅ **Scalar shape** — Tr(M_K^2)/Tr(M_K)^2
- ✅ **Hypercharge normalization** — k_Y=5/3
- ✅ **Boundary weak angle** — sin²θ*=3/8
- ✅ **Gauge/Higgs inventory** — U(1)_Y × SU(2)_L × SU(3)_C + one complex Higgs doublet
- ✅ **Pfaffian scale positive** — v_Pf computed from Planck mass bridge
- ✅ **Higgs tree proxy** — m_H^tree ≈ 124.925 GeV under project Planck convention
- ✅ **Majorana stable thermal relic rejected** — overcloses by ~1.3e13
- ✅ **Native charged flavor firewall** — dim M_charged^native=13
- ✅ **KMS family hierarchy capacity** — ρβ nontracial for β≠0
- ✅ **Noncommuting capacity** — K does not commute with shift/quadrature
- ✅ **CP capacity not CP prediction** — phase coefficients remain free
- ✅ **Latest gate marker** — gate=425
- ✅ **Stable thermal B-gap Majorana rejected** — overclosure ratio computed, not guessed
- ✅ **Suppressed/nonthermal target computed** — conditional yield fraction exists, production history remains sealed

## Verdict

PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.
