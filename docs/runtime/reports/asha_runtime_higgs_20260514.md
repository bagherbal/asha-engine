--- scenario higgs text ---
ASHA runtime board — gate425-runtime-20260514 (Gate 425)
Scenario: higgs
Source: github.com/bagherbal/asha-engine

== Higgs and coefficient bridge ==
One-form edge measure and Pfaffian scale lane yield the tree-level Higgs proxy; pole-mass and RG thresholds remain bridge work.
SYMBOL            VALUE              STATUS           FORMULA/NOTE
f₂(Λ/M_P)²        π²/8               bridge-required  f₂(Λ/M_P)² = π²/8 CCM/Einstein normalization bridge
8π                25.1327412287      bridge-required  (π²/8)/(π/64)=8π correction of earlier coefficient route
v_Pf              247.151135557 GeV  bridge-required  v_Pf = M_P 2^(3/2) exp(-4π²) 
(e/a²)_node       1197/4624          native/audited   1197/4624 
(e/a²)_edge       0.181206747405     bridge-required  (7/10)(1197/4624) 
λ_H               0.12774563655      bridge-required  π²(1197/4624)/20 
m_H^tree          124.925370288 GeV  bridge-required  v_Pf sqrt(2λ_H) tree-level proxy, not pole-mass theorem
M_B               1467749.73718 GeV  bridge-required  sealed B-gap Majorana ledger 
Ω_candidate/Ω_DM  1.3e+13            failed-route     stable thermal B-gap Majorana relic overclosure 
BOUNDARY: Higgs pole mass | m_H^tree + RG + thresholds + self-energy | runtime reports tree proxy only

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
