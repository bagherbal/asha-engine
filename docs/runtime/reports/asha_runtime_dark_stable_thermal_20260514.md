--- scenario dark-stable-thermal text ---
ASHA runtime board — gate425-runtime-20260514 (Gate 425)
Scenario: dark-stable-thermal
Source: github.com/bagherbal/asha-engine

== Dark-sector scenarios ==
Heavy finite sectors are scenario-classified. Stable thermal B-gap Majorana relics are rejected by overclosure; decay/nonthermal routes require extra history.
SYMBOL            VALUE              STATUS           FORMULA/NOTE
M_B               1467749.73718 GeV  bridge-required   
Ω_candidate/Ω_DM  1.3e+13            failed-route      
BOUNDARY: B-gap Majorana stable thermal relic | Ω_candidate/Ω_DM ~ 1.3×10¹³ | simple stable thermal interpretation is rejected by overclosure.
BOUNDARY: decaying/portal heavy sector | Ω_heavy h² = 0 after decay assumptions | allowed only with sealed decay/portal dynamics.
BOUNDARY: nonthermal/axion-like routes | requires production-history axiom | not native ASHA output.

Checks:
- PASS Clifford dimension: dim Cℓ(1,7)=2^8
- PASS Exterior grade dimensions: [1,8,28,56,70,56,28,8,1]
- PASS Boolean/G2 contact vacuum: rank(P_B)=56 rank(P_G)=14 dim K=7
- PASS Scalar shape: Tr(M_K^2)/Tr(M_K)^2
- PASS Hypercharge normalization: k_Y=5/3
- PASS Boundary weak angle: sin²θ*=3/8
- PASS Gauge/Higgs inventory: U(1)_Y × SU(2)_L × SU(3)_C + one complex Higgs doublet
- PASS Pfaffian scale positive: v_Pf computed from Planck mass bridge
- PASS Higgs tree proxy: m_H^tree ≈ 124.925 GeV under project Planck convention
- PASS Majorana stable thermal relic rejected: overcloses by ~1.3e13
- PASS Native charged flavor firewall: dim M_charged^native=13
- PASS KMS family hierarchy capacity: ρβ nontracial for β≠0
- PASS Noncommuting capacity: K does not commute with shift/quadrature
- PASS CP capacity not CP prediction: phase coefficients remain free
- PASS Latest gate marker: gate=425

Verdict: PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data.
